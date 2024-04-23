package command

import (
	"fmt"
	"ginl"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

func Init() {
	var commands []*cobra.Command

	var cmd *cobra.Command

	cmd = &cobra.Command{
		Use:   "version",
		Short: "Version Information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Golang Version: ", runtime.Version())
			os.Exit(0)
		},
	}
	commands = append(commands, cmd)
	cmd = &cobra.Command{
		Use:   "server",
		Short: "server run",
		Run: func(cmd *cobra.Command, args []string) {
			ginl.NewServer()
		},
	}
	commands = append(commands, cmd)

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(commands...)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
