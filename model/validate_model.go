package model

import "fmt"

type ValidateResult struct {
	DocumentPath string
	Line         int
	Repo         string
	RepoFilePath string
}

type StaleResult struct {
	ValidateResult
	DocumentHash string
	LatestHash   string
}

func (sr *StaleResult) String() string {
	return fmt.Sprintf("%s:%s:%d is stale! %s -> %s",
		sr.Repo, sr.DocumentPath, sr.Line, sr.DocumentHash, sr.LatestHash)
}

type FileMissingResult struct {
	ValidateResult
	LastSeenHash string
}

func (fmr *FileMissingResult) String() string {
	return fmt.Sprintf("%s:%s:%d is missing! last seen hash: %s",
		fmr.Repo, fmr.DocumentPath, fmr.Line, fmr.LastSeenHash)
}

type InvalidResult struct {
	ValidateResult
	Message string
}

func (ir *InvalidResult) String() string {
	return fmt.Sprintf("%s:%s:%d is invalid! %s",
		ir.Repo, ir.DocumentPath, ir.Line, ir.Message)
}

type ValidateResults struct {
	StaleResults       []*StaleResult
	FileMissingResults []*FileMissingResult
	InvalidResults     []*InvalidResult
}

func (vr *ValidateResults) Append(result *ValidateResults) {
	vr.StaleResults = append(vr.StaleResults, result.StaleResults...)
	vr.FileMissingResults = append(vr.FileMissingResults, result.FileMissingResults...)
	vr.InvalidResults = append(vr.InvalidResults, result.InvalidResults...)
}

func (vr *ValidateResults) AddStaleResult(document *Document, annotation *Annotation, latestHash string) {
	vr.StaleResults = append(vr.StaleResults, &StaleResult{
		ValidateResult: ValidateResult{
			DocumentPath: document.Path,
			Line:         annotation.Line,
			Repo:         annotation.Repo,
			RepoFilePath: annotation.RepoFilePath,
		},
		DocumentHash: annotation.DocumentHash,
		LatestHash:   latestHash,
	})
}

func (vr *ValidateResults) AddFileMissingResult(document *Document, annotation *Annotation) {
	vr.FileMissingResults = append(vr.FileMissingResults, &FileMissingResult{
		ValidateResult: ValidateResult{
			DocumentPath: document.Path,
			Line:         annotation.Line,
			Repo:         annotation.Repo,
			RepoFilePath: annotation.RepoFilePath,
		},
		LastSeenHash: annotation.DocumentHash,
	})
}

func (vr *ValidateResults) AddInvalidResult(document *Document, annotation *Annotation, message string) {
	vr.InvalidResults = append(vr.InvalidResults, &InvalidResult{
		ValidateResult: ValidateResult{
			DocumentPath: document.Path,
			Line:         annotation.Line,
			Repo:         annotation.Repo,
			RepoFilePath: annotation.RepoFilePath,
		},
		Message: message,
	})
}
