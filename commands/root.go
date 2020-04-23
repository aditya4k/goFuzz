package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	userLicense string
	verbose     bool
	times       int
)

var rootCmd = &cobra.Command{
	Use:   "fuzzme [fuzz, bypass, insert-fuzz, insert-bypass]",
	Short: "fuzzme is used for fuzzing and bypassing the WAF",
	Long: `fuzzme is a client app which can be used by the user to get the
		payload that can bypass the WAF. You can fuzz the webapp or website
		and get the payload that can be send to these websites to bypass the
		Web Application Firewall.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Welcome to FuzzMe - A dear friend of Pentester!!")
	},
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}
