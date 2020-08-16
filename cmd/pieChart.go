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
		renderChart(path, output)
	},
}

func init() {
	chartCmd.AddCommand(pieChartCmd)
	pieChartCmd.Flags().StringP("csv", "f", "./data.csv", "Specify the path of csv file")
	pieChartCmd.Flags().StringP("output", "o", "./", "Specify the output path")
}

func parseCsv(path string) (chart.Values, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	csv := csv2.NewReader(file)

	var data []chart.Value

	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		value, _ := strconv.ParseFloat(row[1], 64)
		data = append(data, chart.Value{Label: row[0], Value: value})
	}
	return data, nil
}
func renderChart(path string, output string)  {
	data, err := parseCsv(path)
	if err != nil {
		log.Fatalln(err)
	}
	pieChart := chart.PieChart{
		Width: 512,
		Height: 512,
		Values: data,
	}
	out, _ := os.Create(output + "/output.png")
	defer out.Close()
	err = pieChart.Render(chart.PNG, out)
	if err != nil {
		log.Fatalln(err)
	}
}