package new

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var verbose bool
	var newCmd = &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new project",
		Long:  `Create a new project with the specified name.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			projectDir := filepath.Join(".", projectName)

			if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
				return fmt.Errorf("directory %s already exists", projectName)
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
	return newCmd
}
