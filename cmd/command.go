package cmd

import (
	serve "github.com/pius-microservices/piopos-auth-service/api"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend",
	Long: "Go backend service",
}

func init() {
	initCommand.AddCommand(serve.ServeCMD)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}