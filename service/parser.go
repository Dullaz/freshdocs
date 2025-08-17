package service

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/model"
)

type Parser struct {
	config       *config.FreshConfig
	onlyDocument string
}

func NewParser(config *config.FreshConfig) *Parser {
	return &Parser{config: config}
}

func (p *Parser) SetOnlyDocument(onlyDocument string) {
	p.onlyDocument = onlyDocument
}

func (p *Parser) Parse() ([]*model.Document, error) {
	documents := []*model.Document{}

	for _, documentGroup := range p.config.DocumentGroups {

		groupDocuments, err := p.parseDocumentGroup(documentGroup)
		if err != nil {
			return nil, err
		}
		documents = append(documents, groupDocuments...)
	}

	return documents, nil
}

func (p *Parser) parseDocumentGroup(documentGroup config.DocumentGroup) ([]*model.Document, error) {
	var documents []*model.Document

	var filepathFilter []string

	if p.onlyDocument != "" {
		matches, err := filepath.Glob(p.onlyDocument)
		if err != nil {
			return nil, err
		}
		for _, match := range matches {
			absPath, err := filepath.Abs(match)
			if err != nil {
				return nil, err
			}
			filepathFilter = append(filepathFilter, absPath)
		}
	}

	err := filepath.Walk(documentGroup.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, documentGroup.Ext) {
			return nil
		}

		if p.onlyDocument != "" {
			absDocument, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if !slices.Contains(filepathFilter, absDocument) {
				return nil
			}
		}

		doc, err := p.parseDocument(path)
		if err != nil {
			return err
		}

		documents = append(documents, doc)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return documents, err
}

func (p *Parser) parseDocument(path string) (*model.Document, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	doc := &model.Document{
		Path: path,
	}

	lines := strings.Split(string(content), "\n")

	doc.Annotations = ParseAnnotations(&lines, path)

	return doc, nil
}

func ParseAnnotations(lines *[]string, documentPath string) []*model.Annotation {
	annotations := []*model.Annotation{}

	for idx, line := range *lines {
		matches := model.ANNOTATION_REGEX.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}

		match := matches[0]

		if len(match) >= 3 {
			repo := match[1]
			path := match[2]
			var hash string
			if len(match) >= 4 {
				hash = match[3]
			}

			annotations = append(annotations, &model.Annotation{
				DocumentPath: documentPath,
				Line:         idx,
				Repo:         repo,
				RepoFilePath: path,
				DocumentHash: hash,
			})
		}
	}

	return annotations
}
