package cmd

import (
	"github.com/spf13/cobra"
	"go-chart/internal"
	"log"
)

// pieChartCmd represents the pieChart command
var pieChartCmd = &cobra.Command{
	Use:   "pieChart",
	Short: "Creates a pie chart for provided csv file",
	Long: `It is a command for goChart program to produce a 
			pie chart from csv file.`,
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
		internal.RenderChart(data, output, 0)
	},
}

func init() {
	chartCmd.AddCommand(pieChartCmd)
	pieChartCmd.Flags().StringP("csv", "f", "./data.csv", "Specify the path of csv file")
	pieChartCmd.Flags().StringP("output", "o", "./", "Specify the output path")
}

