package cmd

import (
	"github.com/spf13/cobra"
	"go-chart/internal"
	"log"
)

// barChartCmd represents the barChart command
var barChartCmd = &cobra.Command{
	Use:   "barChart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("csv")
		if err != nil {
			log.Fatalln(err)
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalln(err)
		}
		data, err := internal.ParseCsv(path)
		if err != nil {
			log.Fatalln(err)
		}
		internal.RenderChart(data, output, 1)
	},
}

func init() {
	chartCmd.AddCommand(barChartCmd)
	barChartCmd.Flags().StringP("csv", "f", "./data.csv", "Specify the path of csv file")
	barChartCmd.Flags().StringP("output", "o", "./", "Specify the output path")
}

