package cmd

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"github.com/strawberry-tools/valerie/valerie/lib"
)

var linksCmd = &cobra.Command{
	Use:          "links",
	Short:        "Check links",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			fmt.Println("At least one argument is required.")
			return nil
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()

		/*if resp.StatusCode != 200 {
			return fmt.Errorf("Failed to retrieve URL.")
		}*/

		doc, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			return err
		}

		links := make(map[string]struct{})

		doc.Find("a, link").Each(func(i int, s *goquery.Selection) {

			val, ok := s.Attr("href")

			if ok {
				links[val] = struct{}{}
			}
		})

		iLinks, eLinks := lib.PrepLinks(links)

		lib.CheckLinks(eLinks)

		var failedCount int

		for _, link := range eLinks {
			if link.Failed {
				fmt.Printf("Link failed: %s Code: %d\n", link.TheURL.String(), link.StatusCode)
				failedCount++
			}
		}

		fmt.Println("")
		fmt.Println("Stats:")
		fmt.Println("=======")
		fmt.Printf("Internal links: %d\n", len(iLinks))
		fmt.Printf("External links: %d\n", len(eLinks))
		fmt.Printf("Links failed: %d\n", failedCount)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(linksCmd)
}
