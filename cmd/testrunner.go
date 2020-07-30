package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func setupTestCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Run test suite for the integration",
		Long:  "Use test runners to verify if the integration collect logs are metrics properly.",
		RunE:  testCommandAction,
	}
}

func testCommandAction(cmd *cobra.Command, args []string) error {
	fmt.Println("test is not implemented yet.")
	return nil
}
