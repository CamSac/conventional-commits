package commands

import (
    "fmt"
    "log"
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
        log.Fatalf("Prompt failed: %v", err)
    }

    var commitMessage string
    textPrompt := &survey.Input{
        Message: "Commit Message",
    }

    err = survey.AskOne(textPrompt, &commitMessage, survey.WithValidator(survey.Required))
    if err != nil {
        log.Fatalf("Prompt failed: %v", err)
    }

    gitCmd := exec.Command("bash", "-c", fmt.Sprintf("git commit -m \"%s: %s\" ", selectedWord, commitMessage))
    gitCmd.Stderr = os.Stderr
    output, err := gitCmd.Output()
    if err != nil {
        log.Fatalf("Error echoing: %v", err)
    }

    fmt.Println("commit message:", string(output))
}
