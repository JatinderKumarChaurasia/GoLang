package main

import (
	"encoding/json"
	"fmt"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

type Value struct {
	Message      int     `json:"messageId"`
	Temperature  float64 `json:"temperature"`
	EnqueuedTime string  `json:"enqueuedTime"`
}

type Values struct {
	Name           string  `json:"name"`
	TemperatureMin float64 `json:"tempMin"`
	TemperatureMax float64 `json:"tempMax"`
	Interval       int     `json:"interval"`
	Values         []Value `json:"values"`
}

type reading struct {
	hour       int
	normal     float64
	outOfRange float64
}

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error Opening File")
		log.Fatal("Error File Opening .............. ")
	}
	defer jsonFile.Close()
	byteData, _ := ioutil.ReadAll(jsonFile)
	var v Values
	_ = json.Unmarshal(byteData, &v)
	temperatureMap := make(map[int][]float64)
	for _, value := range v.Values {
		enqueueTime, err := time.Parse("2006-01-02 15:04:05", value.EnqueuedTime)
		if err != nil {
			log.Fatal("Error Parsing Enqueued Time")
		}
		hour := enqueueTime.Hour()
		temperatureMap[hour] = append(temperatureMap[hour], value.Temperature)
	}

	var normal, outOfRange float64
	var readings []reading
	for hourIndex, temperatureValues := range temperatureMap {
		normal, outOfRange = 0.0, 0.0
		for _, temp := range temperatureValues {
			if temp >= v.TemperatureMin && temp <= v.TemperatureMax {
				normal++
			} else {
				outOfRange++
			}
		}
		read := reading{
			normal: normal, outOfRange: outOfRange, hour: hourIndex,
		}
		readings = append(readings, read)
	}

	sort.Slice(readings, func(i, j int) bool {
		return readings[i].hour < readings[j].hour
	})

	printTable(readings)
	printChart(readings)
}

func printTable(readings []reading) {
	fmt.Printf(" Hour \t Total \t Normal \t Out Of Range \t Percent\n")
	fmt.Println("---------------------------------------------------------")
	for _, value := range readings {
		total := value.normal + value.outOfRange
		percent := value.outOfRange / total * 100
		fmt.Printf(" %v \t %v \t %v \t\t %5v \t\t %5.1f\n", value.hour, total, value.normal, value.outOfRange, percent)
	}

}

func printChart(readings []reading) {
	var bars []chart.StackedBar
	for _, val := range readings {
		msg := fmt.Sprintf("Hour %d", val.hour)
		bar := chart.StackedBar{
			Name:  msg,
			Width: 0,
			Values: []chart.Value{
				{Style: chart.Style{
					Show:     true,
					DotWidth: 2.5,
				}, Label: "Green", Value: val.normal},
				{Style: chart.Style{
					Show:                true,
					StrokeWidth:         0,
					StrokeColor:         drawing.Color{},
					DotColor:            drawing.Color{},
					DotWidth:            1.5,
					DotWidthProvider:    nil,
					DotColorProvider:    nil,
					FontSize:            20,
					FontColor:           drawing.Color{R: 0, G: 0, B: 255},
					Font:                nil,
					TextHorizontalAlign: chart.TextHorizontalAlignCenter,
					TextVerticalAlign:   chart.TextVerticalAlignTop,
					TextWrap:            chart.TextWrapWord,
					TextLineSpacing:     2,
					TextRotationDegrees: 25.0,
				}, Label: "Yellow", Value: val.outOfRange},
			},
		}
		bars = append(bars, bar)
	}
	stackedBarChart := chart.StackedBarChart{
		Title:        "Sensor Temperature Details",
		TitleStyle:   chart.StyleTextDefaults(),
		ColorPalette: chart.ColorPalette(chart.AlternateColorPalette),
		Width:        2000,
		Height:       500,
		Bars:         bars,
	}
	f, err := os.Create("sensorOutput.png")
	if err != nil {
		fmt.Println("Error Creating Output Image")
	}
	//defer f.Close() // resolving the unhandled error exception
	defer func() {
		corr := f.Close()
		if err == nil {
			err = corr
		}
	}()

	_ = stackedBarChart.Render(chart.PNG, f)
}
