package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var nbCmp int

type MyFloat struct {
	value float64
}

func (a MyFloat) Less(b MyFloat) bool {
	nbCmp++
	return a.value < b.value
}

// bubbleSort implementation
func bubbleSort(arr []MyFloat) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].Less(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// quickSort function
func quickSort(arr []MyFloat) {
	if len(arr) <= 1 {
		return
	}
	pivotIndex := partition(arr)
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

// partition function for quickSort
func partition(arr []MyFloat) int {
	pivot := arr[len(arr)-1]
	i := -1
	for j := 0; j < len(arr)-1; j++ {
		if arr[j].Less(pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}



//Change your algorithm here
func main() {
	rand.Seed(time.Now().UnixNano())

	var complexities []opts.LineData
	var xs []opts.LineData

	for n := 1; n <= 500; n++ {
		nbCmp = 0
		vals := make([]MyFloat, n)
		for i := range vals {
			vals[i] = MyFloat{value: float64(rand.Intn(101))}
		}

		//Change your algorithm here
		quickSort(vals)
		complexities = append(complexities, opts.LineData{Value: nbCmp})
		xs = append(xs, opts.LineData{Value: n})
	}

	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Quick sort complexity",
	}))

	line.SetXAxis(xs).
		AddSeries("Linéaire", xs).
		AddSeries("n*log(n)", func() []opts.LineData {
			data := make([]opts.LineData, len(xs))
			for i, x := range xs {
				value := x.Value.(int)
				data[i] = opts.LineData{Value: float64(value) * math.Log(float64(value))}
			}
			return data
		}()).
		AddSeries("Nombre de Comparaisons", complexities)

	//Algorithms name
	f, err := os.Create("bubbleSort_complexity.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = line.Render(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Graphique généré: quick_sort_complexity.html")
}
