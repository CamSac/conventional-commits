package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/CamSac/gi/commands"
)

var rootCmd = &cobra.Command{
    Use:   "word-selector",
    Short: "A CLI tool to select a word from a list and enter some text",
    Run:   commands.Commit,
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
