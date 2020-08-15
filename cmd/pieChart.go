package cmd

import (
	csv2 "encoding/csv"
	"github.com/spf13/cobra"
	"github.com/wcharczuk/go-chart"
	"io"
	"log"
	"os"
	"strconv"
)

// pieChartCmd represents the pieChart command
var pieChartCmd = &cobra.Command{
	Use:   "pieChart",
	Short: "Creates a pie chart for provided csv file",
	Long: `It is a command for goChart program to produce a 
			pie chart from csv file.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, errCsv := cmd.Flags().GetString("csv")
		output, errOut := cmd.Flags().GetString("output")
		if errCsv != nil {
			log.Fatalln(errCsv)
		}
		if errOut != nil {
			log.Fatalln(errOut)
		}
		parseCsv(path, output)
	},
}

func init() {
	rootCmd.AddCommand(pieChartCmd)
	pieChartCmd.Flags().StringP("csv", "f", "./data.csv", "Specify the path of csv file")
	pieChartCmd.Flags().StringP("output", "o", "./", "Specify the output path")
}

func parseCsv(path string, output string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	csv := csv2.NewReader(file)

	var data []chart.Value

	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		value, _ := strconv.ParseFloat(row[1], 64)
		data = append(data, chart.Value{Label: row[0], Value: value})
	}
	renderChart(data, output)

}
func renderChart(data []chart.Value, output string)  {
	pieChart := chart.PieChart{
		Width: 512,
		Height: 512,
		Values: data,
	}
	out, _ := os.Create(output + "/output.png")
	defer out.Close()
	err := pieChart.Render(chart.PNG, out)
	if err != nil {
		log.Fatalln(err)
	}
}