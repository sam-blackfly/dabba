package main

import (
	"github.com/sam-blackfly/dabba/internal/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dabba",
	Short: "dabba is a dummy containerization platform",
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(cmd.RunCmd)
}

func main() {
	RootCmd.Execute()
}
