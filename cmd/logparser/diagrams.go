package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateBarChart(info int, error int, warn int) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Log Informations",
		Subtitle: "Log Informations",
	}))

	bar.SetXAxis([]string{"INFO", "ERROR", "WARN"})

	bar.AddSeries("Category", generateBarItems(info, error, warn))

	f, _ := os.Create("bar.html")
	bar.Render(f)

}

func generateBarItems(info int, error int, warn int) []opts.BarData {
	items := make([]opts.BarData, 0)
	items = append(items, opts.BarData{Value: info, ItemStyle: &opts.ItemStyle{Color: "green"}})
	items = append(items, opts.BarData{Value: error, ItemStyle: &opts.ItemStyle{Color: "red"}})
	items = append(items, opts.BarData{Value: warn, ItemStyle: &opts.ItemStyle{Color: "yellow"}})
	return items
}
