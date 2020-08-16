package cmd

import (
	"github.com/spf13/cobra"
)

// chartCmd represents the chart command
var chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Charting tool",
	Long: `Using chart command it is possible to create different kind of charts.
Use one of the available charts below to create now!`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(chartCmd)
}
