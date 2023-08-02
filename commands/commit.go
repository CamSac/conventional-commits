package commands

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
    "github.com/AlecAivazis/survey/v2"
)

func Commit(cmd *cobra.Command, args []string) {
    conventionalCommits := []string{"fix", "feat", "BREAKING CHANGE", "build", "chore", "ci", "docs", "style", "refactor", "perf", "test"}

    var selectedWord string
    wordPrompt := &survey.Select{
        Message: "Select a commit type",
        Options: conventionalCommits,
    }

    err := survey.AskOne(wordPrompt, &selectedWord, survey.WithValidator(survey.Required))
    if err != nil {
        os.Exit(1)
    }

    var commitMessage string
    textPrompt := &survey.Input{
        Message: "Commit Message",
    }

    err = survey.AskOne(textPrompt, &commitMessage, survey.WithValidator(survey.Required))
    if err != nil {
        os.Exit(1)
    }

    gitCmd := exec.Command("bash", "-c", fmt.Sprintf("git commit -m \"%s: %s\" ", selectedWord, commitMessage))
    gitCmd.Stderr = os.Stderr
    err = gitCmd.Run()
    if err != nil {
        os.Exit(1)
    }

}
