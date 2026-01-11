package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize .envcheck.yaml config file",
		Long:  `Create a .envcheck.yaml configuration file in the current directory.`,
		Example: `  # Create config file with defaults
	envcheck init
	
	# Then customize it
	vim .envcheck.yaml`,

		RunE: func(cmd *cobra.Command, args []string) error {
			return createConfigFile()
		},
	}

	func InitCmd() *cobra.Command {
		return initCmd
	}

	func createConfigFile() error {
		configContent := `# envcheck configuration file
	# 
	# Schema file to validate against
	schema: schema.yaml

	# Strict mode: fail on unexpected APP_* variables
	strict: false

	# Verbose output: show detailed validation info
	verbose: false
`

	filename := ".envcheck.yaml"
	
	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("config file already exists: %s", filename)
	}

	if err := os.WriteFile(filename, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	fmt.Printf("✓ Created config file: %s\n", filename)
	fmt.Println("\nEdit this file to customize envcheck behavior.")
	fmt.Println("Flags will override config file settings.")
	
	return nil
}