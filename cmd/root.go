package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emberd",
	Short: "emberd — manage systemd services efficiently",
	Long:  "emberd is a small CLI to start/stop/restart/reload and view logs for systemd units.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// ensure an explicit help command is available
	rootCmd.SetHelpCommand(cobra.HelpCommand())

	// version subcommand
	var version = "v0.1.0"
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the emberd version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	rootCmd.AddCommand(startCmd, stopCmd, restartCmd, reloadCmd, statusCmd, logsCmd, versionCmd)
}
