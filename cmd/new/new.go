package new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var verbose bool
	var actionName string

	var newCmd = &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new project",
		Long:  `Create a new project with the specified name.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			projectDir := filepath.Join(".", projectName)

			if actionName != "" {
				// Check if the 'actions' directory exists
				_, err := os.Stat("actions")
				if os.IsNotExist(err) {
					return fmt.Errorf("the 'actions' directory does not exist")
				} else if err != nil {
					return fmt.Errorf("error checking 'actions' directory: %v", err)
				}

				projectDir = filepath.Join(".", "actions", actionName)
			}

			if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
				return fmt.Errorf("directory %s already exists", projectDir)
			}

			if err := os.MkdirAll(projectDir, 0755); err != nil {
				return fmt.Errorf("error creating project directory: %v", err)
			}

			if verbose {
				fmt.Printf("Created project directory: %s\n", projectDir)
			}

			return nil
		},
	}
	newCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	newCmd.Flags().StringVar(&actionName, "action", "", "Create a new action within the 'actions' directory")
	return newCmd
}
