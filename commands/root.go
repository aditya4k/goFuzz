package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fuzzWAF [fuzz, bypass, insert-fuzz, insert-bypass]",
	Short: "fuzzWAF is used for fuzzing and bypassing the WAF",
	Long: `fuzzWAF is a client cli which can be used by the user to get the
		payload that can bypass the WAF. You can fuzz the webapp or website
		and get the keywords that can be send to these websites to bypass the
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
