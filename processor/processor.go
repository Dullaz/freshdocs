package processor

import (
	"crypto/sha256"
	"fmt"
	"os"
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
	Hash        string
}

// Annotation represents a link between a document and code
type Annotation struct {
	Repository string
	Path       string
	Line       int
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
			if p.isDocumentStale(doc) {
				affected = append(affected, doc)
			}
		}
	}

	return affected, nil
}

// Check performs a quick check for stale documentation
func (p *Processor) Check() ([]Document, error) {
	// Similar to Validate but with different output format
	return p.Validate()
}

// Update updates hashes for all documents
func (p *Processor) Update() error {
	for _, group := range p.config.DocumentGroups {
		docs, err := p.processDocumentGroup(group)
		if err != nil {
			return err
		}

		for _, doc := range docs {
			if err := p.updateDocumentHash(doc); err != nil {
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
			return p.updateDocumentHash(doc)
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

	// Calculate hash
	doc.Hash = p.calculateHash(string(content))

	return doc, nil
}

// parseAnnotations extracts FreshDocs annotations from document content
func (p *Processor) parseAnnotations(content string) ([]Annotation, error) {
	var annotations []Annotation

	// Regex to match FreshDocs annotations: <!--- fresh:file repo:path -->
	re := regexp.MustCompile(`<!---\s*fresh:file\s+(\w+):([^\s]+)\s*-->`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			repo := match[1]
			path := match[2]

			// Validate repository exists in config
			if _, exists := p.config.Repositories[repo]; !exists {
				return nil, fmt.Errorf("unknown repository: %s", repo)
			}

			annotations = append(annotations, Annotation{
				Repository: repo,
				Path:       path,
			})
		}
	}

	return annotations, nil
}

// calculateHash calculates SHA256 hash of content
func (p *Processor) calculateHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return fmt.Sprintf("%x", hash)
}

// isDocumentStale checks if a document is stale by comparing hashes
func (p *Processor) isDocumentStale(doc Document) bool {
	// For now, we'll consider a document stale if it has annotations
	// In a real implementation, you'd compare against stored hashes
	return len(doc.Annotations) > 0
}

// updateDocumentHash updates the stored hash for a document
func (p *Processor) updateDocumentHash(doc Document) error {
	// In a real implementation, you'd store the hash in a database or file
	// For now, we'll just recalculate it
	content, err := os.ReadFile(doc.Path)
	if err != nil {
		return err
	}

	doc.Hash = p.calculateHash(string(content))
	// Store the hash somewhere (database, file, etc.)

	return nil
}
