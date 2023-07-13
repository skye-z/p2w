/*
子命令 - 转PDF

BetaX Page to what
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Convert web page to PDF",
	Long:  "Convert web page to PDF",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(pdfCmd)
}
