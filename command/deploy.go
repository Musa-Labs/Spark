package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func DeployCmd() *cobra.Command {
	var commands []string
	var verbose bool
	var deployCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Execute deployment commands",
		Long:  `Execute a series of OS commands for deployment.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(commands) == 0 {
				return fmt.Errorf("no commands specified. Use --cmd flag to specify commands")
			}

			for _, cmdStr := range commands {
				if verbose {
					fmt.Printf("Executing: %s\n", cmdStr)
				}

				cmd := exec.Command("bash", "-c", cmdStr)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("failed to execute command '%s': %w", cmdStr, err)
				}

				if verbose {
					fmt.Printf("Successfully executed: %s\n", cmdStr)
				}
			}

			return nil
		},
	}
	deployCmd.Flags().StringArrayVarP(&commands, "cmd", "c", []string{}, "Commands to execute (can be specified multiple times)")
	deployCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	return deployCmd
}
