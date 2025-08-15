package model

import "fmt"

// Represents any annotations that
// 1. will be affected by staged/unstaged changes
// 2. are invalid
// 3. are new and need a hash
type CheckResult struct {
	DocumentPath string
	Line         int
	Repo         string
	RepoFilePath string
	DocumentHash *string
	State        CheckState
}

func (r CheckResult) String() string {
	if r.State == CheckStateNeedsUpdate {
		return fmt.Sprintf("%s:%d has no hash!",
			r.DocumentPath, r.Line)
	}
	if r.State == CheckStateInvalid {
		return fmt.Sprintf("%s:%d, repo: %s, repo file path: %s is invalid!",
			r.DocumentPath, r.Line, r.Repo, r.RepoFilePath)
	}
	return fmt.Sprintf("%s:%d affected by %s:%s",
		r.DocumentPath, r.Line, r.Repo, r.RepoFilePath)
}

type CheckState int

const (
	CheckStateAffected CheckState = iota
	CheckStateInvalid
	CheckStateNeedsUpdate
)

var checkStateNames = map[CheckState]string{
	CheckStateAffected:    "affected",
	CheckStateInvalid:     "invalid",
	CheckStateNeedsUpdate: "needs update",
}

func (s CheckState) String() string {
	return checkStateNames[s]
}
