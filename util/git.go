package util

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetRepoPath() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get repo path: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

func GetGitHash(repoPath string, file string) (string, error) {
	fullPath := filepath.Join(repoPath, file)

	cmd := exec.Command("git", "log", "-n", "1", "--pretty=format:%h", "--", fullPath)
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get Git hash for %s: %w", fullPath, err)
	}

	hash := strings.TrimSpace(string(output))
	if len(hash) >= 7 {
		return hash[:7], nil
	}
	return hash, nil
}

func GetChanges(repoPath string, file string) ([]string, error) {

	unstagedCmd := exec.Command("git", "diff", "--name-only")
	unstagedCmd.Dir = repoPath
	unstagedOutput, err := unstagedCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get unstaged changes: %w", err)
	}

	stagedCmd := exec.Command("git", "diff", "--name-only", "--cached")
	stagedCmd.Dir = repoPath
	stagedOutput, err := stagedCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get staged changes: %w", err)
	}

	changes := []string{}
	changes = append(changes, strings.Split(string(unstagedOutput), "\n")...)
	changes = append(changes, strings.Split(string(stagedOutput), "\n")...)

	return changes, nil
}
