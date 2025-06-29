package processor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dullaz/freshdocs/config"
)

// Processor handles the core logic for FreshDocs
type Processor struct {
	config *config.Config
}

// Document represents a processed document with its annotations
type Document struct {
	Path        string
	AffectedBy  string
	Annotations []Annotation
}

// Annotation represents a link between a document and code
type Annotation struct {
	Repository string
	Path       string
	Line       int
	Hash       string // Git commit hash of the linked file
}

// NewProcessor creates a new processor instance
func NewProcessor(cfg *config.Config) *Processor {
	return &Processor{
		config: cfg,
	}
}

// Validate checks which documents are affected by code changes
func (p *Processor) Validate() ([]Document, error) {
	var affected []Document

	for _, group := range p.config.DocumentGroups {
		docs, err := p.processDocumentGroup(group)
		if err != nil {
			return nil, err
		}

		for _, doc := range docs {
			for _, ann := range doc.Annotations {
				if p.isAnnotationStale(ann) {
					doc.AffectedBy = ann.Path
					affected = append(affected, doc)
					break // Only add document once even if multiple annotations are stale
				}
			}
		}
	}

	return affected, nil
}

// Check performs a quick check for stale documentation
func (p *Processor) Check() ([]Document, error) {
	var affected []Document

	for _, group := range p.config.DocumentGroups {
		docs, err := p.processDocumentGroup(group)
		if err != nil {
			return nil, err
		}

		for _, doc := range docs {
			for _, ann := range doc.Annotations {
				if p.isAnnotationAffectedByGitChanges(ann) {
					doc.AffectedBy = ann.Path
					affected = append(affected, doc)
					break // Only add document once even if multiple annotations are affected
				}
			}
		}
	}

	return affected, nil
}

// Update updates hashes for all documents
func (p *Processor) Update() error {
	for _, group := range p.config.DocumentGroups {
		docs, err := p.processDocumentGroup(group)
		if err != nil {
			return err
		}

		for _, doc := range docs {
			if err := p.updateDocumentHashes(doc); err != nil {
				return err
			}
		}
	}

	return nil
}

// UpdateFile updates hash for a specific document
func (p *Processor) UpdateFile(filePath string) error {
	// Find the document group that contains this file
	for _, group := range p.config.DocumentGroups {
		if strings.HasPrefix(filePath, group.Path) {
			doc := Document{Path: filePath}
			return p.updateDocumentHashes(doc)
		}
	}

	return fmt.Errorf("file %s not found in any document group", filePath)
}

// Find finds documents linked to a specific code file
func (p *Processor) Find(codeFile string) ([]string, error) {
	var linkedDocs []string

	for _, group := range p.config.DocumentGroups {
		docs, err := p.processDocumentGroup(group)
		if err != nil {
			return nil, err
		}

		for _, doc := range docs {
			for _, ann := range doc.Annotations {
				if ann.Path == codeFile {
					linkedDocs = append(linkedDocs, doc.Path)
					break
				}
			}
		}
	}

	return linkedDocs, nil
}

// processDocumentGroup processes all documents in a group
func (p *Processor) processDocumentGroup(group config.DocumentGroup) ([]Document, error) {
	var docs []Document

	err := filepath.Walk(group.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, group.Ext) {
			doc, err := p.processDocument(path)
			if err != nil {
				return err
			}
			docs = append(docs, doc)
		}

		return nil
	})

	return docs, err
}

// processDocument processes a single document
func (p *Processor) processDocument(path string) (Document, error) {
	doc := Document{Path: path}

	// Read file content
	content, err := os.ReadFile(path)
	if err != nil {
		return doc, err
	}

	// Parse annotations
	annotations, err := p.parseAnnotations(string(content))
	if err != nil {
		return doc, err
	}
	doc.Annotations = annotations

	return doc, nil
}

// parseAnnotations extracts FreshDocs annotations from document content
func (p *Processor) parseAnnotations(content string) ([]Annotation, error) {
	var annotations []Annotation

	// Regex to match FreshDocs annotations that start at the beginning of a line
	// <!--- fresh:file repo:path [hash] -->
	re := regexp.MustCompile(`(?m)^\s*<!---\s*fresh:file\s+(\w+):([^\s]+)(?:\s+([a-f0-9]+))?\s*-->`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			repo := match[1]
			path := match[2]
			hash := ""
			if len(match) >= 4 {
				hash = match[3]
			}

			// Validate repository exists in config
			if _, exists := p.config.Repositories[repo]; !exists {
				return nil, fmt.Errorf("unknown repository: %s", repo)
			}

			annotations = append(annotations, Annotation{
				Repository: repo,
				Path:       path,
				Hash:       hash,
			})
		}
	}

	return annotations, nil
}

// getGitHash gets the current Git commit hash for a file
func (p *Processor) getGitHash(repoName, filePath string) (string, error) {
	repo, exists := p.config.Repositories[repoName]
	if !exists {
		return "", fmt.Errorf("unknown repository: %s", repoName)
	}

	// Construct full path to the file
	fullPath := filepath.Join(repo.Path, filePath)

	// Get Git hash for the file
	cmd := exec.Command("git", "rev-parse", "HEAD:"+filePath)
	cmd.Dir = repo.Path

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get Git hash for %s: %w", fullPath, err)
	}

	// Return first 7 characters (short hash)
	hash := strings.TrimSpace(string(output))
	if len(hash) >= 7 {
		return hash[:7], nil
	}
	return hash, nil
}

// isAnnotationStale checks if an annotation is stale by comparing hashes
func (p *Processor) isAnnotationStale(ann Annotation) bool {
	// If no hash is set, it's considered stale (first time annotation)
	if ann.Hash == "" {
		return true
	}

	// Get current hash for the file
	currentHash, err := p.getGitHash(ann.Repository, ann.Path)
	if err != nil {
		// If we can't get the hash, assume it's stale
		return true
	}

	// Compare hashes
	return ann.Hash != currentHash
}

// isAnnotationAffectedByGitChanges checks if an annotation is affected by unstaged or staged changes
func (p *Processor) isAnnotationAffectedByGitChanges(ann Annotation) bool {
	repo, exists := p.config.Repositories[ann.Repository]
	if !exists {
		return false
	}

	// Check if the file has unstaged changes
	cmd := exec.Command("git", "diff", "--name-only", ann.Path)
	cmd.Dir = repo.Path
	unstagedOutput, err := cmd.Output()
	if err == nil && len(strings.TrimSpace(string(unstagedOutput))) > 0 {
		return true
	}

	// Check if the file has staged changes
	cmd = exec.Command("git", "diff", "--cached", "--name-only", ann.Path)
	cmd.Dir = repo.Path
	stagedOutput, err := cmd.Output()
	if err == nil && len(strings.TrimSpace(string(stagedOutput))) > 0 {
		return true
	}

	return false
}

// updateDocumentHashes updates the hashes in a document file
func (p *Processor) updateDocumentHashes(doc Document) error {
	// Read current content
	content, err := os.ReadFile(doc.Path)
	if err != nil {
		return err
	}

	contentStr := string(content)

	// Update each annotation with current hash
	for _, ann := range doc.Annotations {
		currentHash, err := p.getGitHash(ann.Repository, ann.Path)
		if err != nil {
			return err
		}

		// Create old and new annotation strings
		oldAnnotation := fmt.Sprintf("<!--- fresh:file %s:%s %s -->", ann.Repository, ann.Path, ann.Hash)
		newAnnotation := fmt.Sprintf("<!--- fresh:file %s:%s %s -->", ann.Repository, ann.Path, currentHash)

		// If no hash was present, handle that case
		if ann.Hash == "" {
			oldAnnotation = fmt.Sprintf("<!--- fresh:file %s:%s -->", ann.Repository, ann.Path)
		}

		// Replace in content
		contentStr = strings.Replace(contentStr, oldAnnotation, newAnnotation, 1)
	}

	// Write updated content back to file
	return os.WriteFile(doc.Path, []byte(contentStr), 0644)
}
