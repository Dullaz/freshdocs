package model

import (
	"regexp"
)

var ANNOTATION_REGEX = regexp.MustCompile(`(?m)^\s*<!---\s*fresh:file\s+(\w+):([^\s]+)(?:\s+([a-f0-9]+))?\s*-->`)

type Document struct {
	Path        string
	Annotations []*Annotation
}

type Annotation struct {
	DocumentPath string
	Line         int
	Repo         string
	RepoFilePath string
	DocumentHash string
}
