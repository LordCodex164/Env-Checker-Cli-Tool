package main

import (
	"env-checker/cmd"
	"fmt"
	"github.com/lpernett/godotenv"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "envcheck",
	Short: "validate environmental variables against a schema",
	Long: `envcheck validates environment variables against a schema file.

	It helps catch configuration errors before your application runs,
	ensuring all required variables are set with correct types.

	Perfect for:
	• Local development environment validation
	• CI/CD pipeline checks
	• Production deployment verification
	• Team onboarding (document required env vars)`,

		Example: `  # Create a schema template
	envcheck create

	# Validate current environment
	envcheck validate

	# Validate with custom schema
	envcheck validate --schema prod-schema.yaml

	# Strict mode: fail on unexpected variables
	envcheck validate --strict

	# Verbose output for debugging
	envcheck validate --verbose
	
	# Use config file (.envcheck.yaml)
	envcheck validate`,
	PersistentPreRunE: func(c *cobra.Command, args []string) error {
		return cmd.InitConfig()
	},
}

func init() {
	//initialize config

	//cobra.OnInitialize(cmd.InitConfig)

	rootCmd.AddCommand(
		cmd.CreateCmd(),
		cmd.ValidateCmd(),
		cmd.VersionCmd(),
		cmd.InitCmd(),
	)
}

func main() {
	godotenv.Load()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
