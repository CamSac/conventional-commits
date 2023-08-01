package commands

import (
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

func Log(cmd *cobra.Command, args []string) {
    gitCmd := exec.Command("git", "--no-pager", "log", "--graph", "--abbrev-commit", "--decorate", "--format=format:%C(bold blue)%h%C(reset) - %C(bold cyan)%aD%C(reset) %C(bold green)(%ar)%C(reset)%C(bold yellow)%d%C(reset)%n''          %C(white)%s%C(reset) %C(dim white)- %an%C(reset)", "--all")

    // Set the GIT_PAGER environment variable to enable color output
    gitCmd.Env = append(os.Environ(), "GIT_PAGER=less -R")

    gitCmd.Stdout = os.Stdout
    gitCmd.Stderr = os.Stderr

    // Run the command
    err := gitCmd.Run()
    if err != nil {
        log.Fatalf("Error executing git log: %v", err)
    }
}
