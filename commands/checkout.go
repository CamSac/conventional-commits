package commands

import (
	"os"
    "os/exec"
	"strings"

    "github.com/spf13/cobra"
	"github.com/AlecAivazis/survey/v2"
)

func Checkout(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
		branchCmd := exec.Command("git", "branch")
		output, err := branchCmd.CombinedOutput()
		if err != nil {
			os.Exit(1)
		}
		var branches []string
		for _, line := range strings.Split(string(output), "\n") {
			branch := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			if branch != "" {
				branches = append(branches, branch)
			}
		}
		var selectedWord string
		wordPrompt := &survey.Select{
			Message: "Select a commit type",
			Options: branches,
		}
		err = survey.AskOne(wordPrompt, &selectedWord, survey.WithValidator(survey.Required))
		if err != nil {
			os.Exit(1)
		}
		checkoutCmd := exec.Command("git", "checkout", selectedWord)
		// checkoutCmd.Stdout = os.Stdout
		checkoutCmd.Stderr = os.Stderr
		err = checkoutCmd.Run()
		if err != nil {
			os.Exit(1)
		}
	}
}
