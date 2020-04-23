package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	url        string
	postParam  string
	cookie     string
	attackType string
	delay      int
	outputFile string
)

func init() {
	fuzzCmd.Flags().StringVarP(&url, "url", "u", "'", "target url")
	fuzzCmd.Flags().StringVarP(&cookie, "cookie", "c", "'", "target url")
	fuzzCmd.Flags().StringVarP(&attackType, "type", "t", "'", "target url")
	fuzzCmd.Flags().StringVarP(&outputFile, "outfile", "o", "'", "target url")
	fuzzCmd.Flags().StringVarP(&postParam, "post", "p", "'", "Send fuzz through post parameter")
	fuzzCmd.Flags().IntVarP(&delay, "delay", "d", 0, "delay between each request")
	fuzzCmd.MarkFlagRequired("url")
	fuzzCmd.MarkFlagRequired("type")
	rootCmd.AddCommand(fuzzCmd)
}

var fuzzCmd = &cobra.Command{
	Use:   "fuzz [url, type, cookie, outfile, post, delay]",
	Short: "check which symbols and keywords are allowed by the WAF.",
	Long: `You can fuzz the webapp or website and get the keywords that
		  can be send to these websites to bypass the Web Application Firewall.`,
	Run: echoRun,
}

func echoRun(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
