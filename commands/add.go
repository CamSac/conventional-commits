package commands

import (
    "fmt"
	"log"
    "os"
    "os/exec"
	"path/filepath"

    "github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) {
	// Find the root folder of the Git repository.
	gitRoot, err := findGitRoot()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Move to the root folder if it is different from the current working directory.
	if gitRoot != "" {
		if err := os.Chdir(gitRoot); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	if len(args) == 0 {
		// If no arguments are provided, run 'git add .' to add all changes.
		cmd := exec.Command("git", "add", ".")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		// If arguments are provided, run 'git add' with the specified files.
		args = append([]string{"add"}, args...)
		cmd := exec.Command("git", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}

func findGitRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse up the directory tree until the root folder (identified by the .git directory) is found.
	for {
		gitDir := filepath.Join(cwd, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return cwd, nil
		}

		// Move one directory up.
		parent := filepath.Dir(cwd)
		if parent == cwd {
			break
		}
		cwd = parent
	}

	return "", fmt.Errorf("not in a Git repository")
}
