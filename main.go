package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/CamSac/gi/commands"
)

// var rootCmd = &cobra.Command{
    // Use:   "word-selector",
    // Short: "A CLI tool to select a word from a list and enter some text",
    // Run:   commands.Commit,
// }

func main() {
    var rootCmd = &cobra.Command{Use: "mytool"}

    var addCmnd = &cobra.Command{
		Use:   "add",
		Short: "Add files to the repository",
		Run: commands.Add,
	}

	var commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "Commit changes to the repository",
		Run: commands.Commit,
	}

    rootCmd.AddCommand(addCmnd)
	rootCmd.AddCommand(commitCmd)

    if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
