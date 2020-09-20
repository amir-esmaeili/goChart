package internal

import (
	csv2 "encoding/csv"
	"github.com/wcharczuk/go-chart"
	"io"
	"os"
	"strconv"
)

func ParseCsv(path string) (chart.Values, error) {
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

func RenderChart(chartData chart.Values, output string, chartType int) {
	switch chartType {
	case 0:
		pieChart(chartData, output)
		break
	case 1:
		barChart(chartData, output)
		break
	}

}

func pieChart(chartData chart.Values, output string) {
	plot := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: chartData,
	}
	out, _ := os.Create(output + "/output.png")
	defer out.Close()
	plot.Render(chart.PNG, out)
}

func barChart(chartData chart.Values, output string) {
	barChart := chart.BarChart{
		Title:      "Quarterly Sales",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 100,
			},
		},
		Width:      810,
		Height:     500,
		XAxis:      chart.StyleShow(),
		BarSpacing: 50,
		Bars: chartData,
	}
	out, _ := os.Create(output + "/output.png")
	defer out.Close()
	barChart.Render(chart.PNG, out)
}
