package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed assets/*
var embeddedFiles embed.FS

var (
	outputDir string
	verbose   bool
	commands  []string
)

var rootCmd = &cobra.Command{
	Use:   "spark",
	Short: "N8N Development Toolkit",
	Long:  `N8N Development Toolkit`,
}

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new project",
	Long:  `Create a new project with the specified name and extract embedded files into it.`,
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

		return extractFiles(projectDir)
	},
}

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

			// Split command string into command and arguments
			cmdParts := strings.Fields(cmdStr)
			if len(cmdParts) == 0 {
				continue
			}

			command := exec.Command(cmdParts[0], cmdParts[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			if err := command.Run(); err != nil {
				return fmt.Errorf("command failed: %s: %v", cmdStr, err)
			}

			if verbose {
				fmt.Printf("Successfully executed: %s\n", cmdStr)
			}
		}

		return nil
	},
}

func init() {
	// Add commands
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(deployCmd)

	// Add flags to new command
	newCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Add flags to deploy command
	deployCmd.Flags().StringArrayVarP(&commands, "cmd", "c", []string{}, "Commands to execute (can be specified multiple times)")
	deployCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

func extractFiles(targetDir string) error {
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %v", err)
	}

	return fs.WalkDir(embeddedFiles, "assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		content, err := embeddedFiles.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading embedded file %s: %v", path, err)
		}

		relPath, err := filepath.Rel("assets", path)
		if err != nil {
			return fmt.Errorf("error getting relative path for %s: %v", path, err)
		}
		outPath := filepath.Join(targetDir, relPath)

		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return fmt.Errorf("error creating directories for %s: %v", outPath, err)
		}

		if err := os.WriteFile(outPath, content, 0644); err != nil {
			return fmt.Errorf("error writing file %s: %v", outPath, err)
		}

		if verbose {
			fmt.Printf("Extracted: %s\n", outPath)
		}
		return nil
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
