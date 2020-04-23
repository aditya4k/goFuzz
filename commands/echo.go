package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	echoCmd.Flags().IntVarP(&times, "times", "t", 2, "times to echo")

	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
				        Echo echo’s.
					    `,
	Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
	for i := 0; i < times; i++ {
		fmt.Println(strings.Join(args, " "))
	}
}
