package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/Musa-Labs/Spark/cmd/deploy"
	"github.com/Musa-Labs/Spark/cmd/new"
)

//go:embed assets/*
var embeddedFiles embed.FS

var (
	outputDir string
	verbose   bool
)

var rootCmd = &cobra.Command{
	Use:   "spark",
	Short: "N8N Development Toolkit",
	Long:  `N8N Development Toolkit`,
}

func init() {
	rootCmd.AddCommand(new.NewCmd())
	rootCmd.AddCommand(deploy.DeployCmd())
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
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
