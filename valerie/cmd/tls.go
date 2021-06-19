package cmd

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var days int

var tlsCmd = &cobra.Command{
	Use:          "tls",
	Short:        "Check if TLS certificate is expired or will expire in 14 days.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		conf := &tls.Config{
			InsecureSkipVerify: true,
		}

		conn, err := tls.Dial("tcp", args[0]+":443", conf)
		if err != nil {
			return err
		}
		defer conn.Close()

		cutoffDate := time.Now().AddDate(0, 0, days)
		expDate := conn.ConnectionState().PeerCertificates[0].NotAfter

		if time.Now().After(expDate) {
			return fmt.Errorf("The TLS certificate has expired.")
		} else if cutoffDate.After(expDate) {
			return fmt.Errorf("The TLS certificate will expire in %d days.", days)
		} else {
			fmt.Println("The TLS certificate is valid.")
		}

		return nil
	},
}

func init() {

	tlsCmd.Flags().IntVar(&days, "days", 14, "days within range to fail")
	rootCmd.AddCommand(tlsCmd)
}
