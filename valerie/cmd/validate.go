package cmd

import (
	"errors"
	"fmt"

	"github.com/antchfx/htmlquery"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:          "validate",
	Short:        "Test getting Open Graph stuff",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		doc, _ := htmlquery.LoadURL(args[0])

		fmt.Println("Validating Open Graph tags for \"" + args[0] + "\":\n")
		failed := false

		// og:title
		tags := htmlquery.Find(doc, "//head/meta[@property='og:title']")
		if len(tags) > 0 {
			fmt.Println("og:title - yes")
		} else {
			fmt.Println("og:title - no")
			failed = true
		}

		// og:type
		tags = htmlquery.Find(doc, "//head/meta[@property='og:type']")
		if len(tags) > 0 {
			fmt.Println("og:type - yes")
		} else {
			fmt.Println("og:type - no")
			failed = true
		}

		// og:type
		tags = htmlquery.Find(doc, "//head/meta[@property='og:image']")
		if len(tags) > 0 {
			fmt.Println("og:image - yes")
		} else {
			fmt.Println("og:image - no")
			failed = true
		}

		// og:url
		tags = htmlquery.Find(doc, "//head/meta[@property='og:url']")
		if len(tags) > 0 {
			fmt.Println("og:url - yes")
		} else {
			fmt.Println("og:url - no")
			failed = true
		}

		if failed {
			return errors.New("One or more Open Graph tags are missing.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
