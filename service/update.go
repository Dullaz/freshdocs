package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dullaz/freshdocs/config"
	"github.com/dullaz/freshdocs/model"
	"github.com/dullaz/freshdocs/util"
)

type Updater struct {
	config *config.FreshConfig
}

func NewUpdater(config *config.FreshConfig) *Updater {
	return &Updater{config: config}
}

func (u *Updater) Update(documentGroup config.DocumentGroup) error {
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

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		lines := strings.Split(string(content), "\n")
		annotations := ParseAnnotations(&lines, path)

		u.updateLines(&lines, annotations)

		err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func (u *Updater) updateLines(lines *[]string, annotations []*model.Annotation) {
	for _, annotation := range annotations {
		latestHash, err := util.GetGitHash(u.config.Repositories[annotation.Repo].Path, annotation.RepoFilePath)
		if err != nil {
			log.Fatalf("failed to get latest hash for %s:%s: %v", annotation.Repo, annotation.RepoFilePath, err)
		}
		(*lines)[annotation.Line] = fmt.Sprintf("<!--- fresh:file %s:%s %s -->", annotation.Repo, annotation.RepoFilePath, latestHash)
	}
}
