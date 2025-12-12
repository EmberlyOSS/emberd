package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/EmberlyOSS/emberd/internal/systemd"
)

var startCmd = &cobra.Command{
	Use:   "start UNIT",
	Short: "Start a systemd unit",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		out, err := systemd.RunSystemctl("start", unit)
		fmt.Print(string(out))
		return err
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop UNIT",
	Short: "Stop a systemd unit",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		out, err := systemd.RunSystemctl("stop", unit)
		fmt.Print(string(out))
		return err
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart UNIT",
	Short: "Restart a systemd unit",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		out, err := systemd.RunSystemctl("restart", unit)
		fmt.Print(string(out))
		return err
	},
}

var reloadCmd = &cobra.Command{
	Use:   "reload UNIT",
	Short: "Reload a systemd unit if it supports reload",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		out, err := systemd.RunSystemctl("reload", unit)
		fmt.Print(string(out))
		if err != nil {
			return fmt.Errorf("reload failed: %w", err)
		}
		return nil
	},
}

var statusCmd = &cobra.Command{
	Use:   "status UNIT",
	Short: "Show status for a systemd unit",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		out, err := systemd.RunSystemctl("status", unit)
		fmt.Print(string(out))
		return err
	},
}

var logsCmd = &cobra.Command{
	Use:   "logs UNIT",
	Short: "Show logs for a systemd unit (uses journalctl)",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		unitFlag, _ := cmd.Flags().GetString("unit")
		unit, err := systemd.UnitFromArgs(args, unitFlag)
		if err != nil {
			return err
		}
		follow, _ := cmd.Flags().GetBool("follow")
		lines, _ := cmd.Flags().GetInt("lines")
		return systemd.StreamJournal(unit, follow, lines)
	},
}

func init() {
	// common flag to allow `-u` instead of passing as arg
	for _, c := range []*cobra.Command{startCmd, stopCmd, restartCmd, reloadCmd, statusCmd, logsCmd} {
		c.Flags().StringP("unit", "u", "", "systemd unit name (fallback if not provided as arg)")
	}
	logsCmd.Flags().BoolP("follow", "f", false, "Follow the journal (like -f)")
	logsCmd.Flags().IntP("lines", "n", 100, "Number of initial log lines to show")
}
