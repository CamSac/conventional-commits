package main

import (
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
    var rootCmd = &cobra.Command{Use: "gi"}

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

    var logCmd = &cobra.Command{
		Use:   "log",
		Short: "Print repository history",
		Run: commands.Log,
	}

	var checkoutCmd = &cobra.Command{
		Use:   "checkout",
		Short: "Checkout branches of the repository",
		Run: commands.Checkout,
	}

    rootCmd.AddCommand(addCmnd)
	rootCmd.AddCommand(commitCmd)
    rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(checkoutCmd)

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
	}
}
