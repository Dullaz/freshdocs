package service

import (
	"fmt"
	"log"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/model"
	"github.com/dullaz/freshdocs/util"
)

type Checker struct {
	changesByRepo map[string][]string
	config        *config.FreshConfig
}

func NewChecker(config *config.FreshConfig) *Checker {
	changesByRepo := make(map[string][]string)
	for repoName, repo := range config.Repositories {
		changes, err := util.GetChanges(repo.Path, "")
		if err != nil {
			log.Fatalf("failed to get changes for repo %s: %v", repo.Path, err)
		}
		changesByRepo[repoName] = changes
	}
	return &Checker{config: config, changesByRepo: changesByRepo}
}

func (c *Checker) Check(document *model.Document) ([]model.CheckResult, error) {

	results := []model.CheckResult{}
	for _, annotation := range document.Annotations {
		changes, ok := c.changesByRepo[annotation.Repo]
		if !ok {
			continue
		}

		if annotation.DocumentHash == "" {
			results = append(results, model.CheckResult{
				DocumentPath: document.Path,
				Line:         annotation.Line,
				Repo:         annotation.Repo,
				RepoFilePath: annotation.RepoFilePath,
				State:        model.CheckStateNeedsUpdate,
			})
			continue
		}

		for _, change := range changes {
			ok, err := doublestar.PathMatch(annotation.RepoFilePath, change)
			if err != nil {
				return nil, fmt.Errorf("failed to match path %s with %s: %w", annotation.RepoFilePath, change, err)
			}
			if ok {
				results = append(results, model.CheckResult{
					DocumentPath: document.Path,
					Line:         annotation.Line,
					Repo:         annotation.Repo,
					RepoFilePath: change,
					State:        model.CheckStateAffected,
				})
			}
		}
	}

	return results, nil
}
