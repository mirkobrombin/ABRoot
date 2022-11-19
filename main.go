package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/ABRoot/cmd"
)

var (
	Version = "0.0.1"
)

func help(cmd *cobra.Command, args []string) {
	fmt.Print(`Usage: 
abroot [options] [command]

Options:
	--help/-h		show this message
	--verbose/-v		show more verbosity
	--version/-V		show version

Commands:
	_sync-future		sync the future root with the present root
	_update-boot		update the boot partition
	get			get the present or future root partition
	shell			enter a transactional shell in the future root and switch root on next boot
	exec			execute a command in a transactional shell in the future root and switch to it on next boot
`)
}

func newABRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "abroot",
		Short:   "ABRoot provides full immutability and atomicity by performing transactions between 2 root partitions (A<->B).",
		Version: Version,
	}
}

func main() {
	rootCmd := newABRootCommand()
	rootCmd.AddCommand(cmd.NewSyncCommand())
	rootCmd.AddCommand(cmd.NewUpdateBootCommand())
	rootCmd.AddCommand(cmd.NewGetCommand())
	rootCmd.AddCommand(cmd.NewShellCommand())
	rootCmd.AddCommand(cmd.NewExecCommand())
	rootCmd.SetHelpFunc(help)
	rootCmd.Execute()
}