package service

import (
	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/model"
	"github.com/dullaz/freshdocs/util"
)

type Validator struct {
	config *config.FreshConfig
}

func NewValidator(config *config.FreshConfig) *Validator {
	return &Validator{config: config}
}

func (v *Validator) Validate(document *model.Document) (*model.ValidateResults, error) {
	results := &model.ValidateResults{}

	for _, annotation := range document.Annotations {
		latestHash, err := util.GetGitHash(v.config.Repositories[annotation.Repo].Path, annotation.RepoFilePath)
		if err != nil {
			results.AddFileMissingResult(document, annotation)
			continue
		}

		if annotation.DocumentHash == "" {
			results.AddInvalidResult(document, annotation, "no hash")
			continue
		}

		if annotation.DocumentHash != latestHash {
			results.AddStaleResult(document, annotation, latestHash)
		}
	}

	return results, nil
}
