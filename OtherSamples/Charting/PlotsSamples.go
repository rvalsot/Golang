package main

import (
	"math/rand"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func main() {
	rand.Seed(int64(0))

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	plt.Title.Text = "Some action random values"
	plt.X.Label.Text = "Windows"
	plt.Y.Label.Text = "Price"

	err = plotutil.AddLinePoints(plt,
		"First", randomPoints(18),
		"Second", randomPoints(18),
		"Third", randomPoints(18))

	if err != nil {
		panic(err)
	}

	if err := plt.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

}

func randomPoints(n int) plotter.XYs {

	pts := make(plotter.XYs, n)

	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
