package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Plotter struct {
	p *plot.Plot
}

func NewPlotter(title string) *Plotter {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	return &Plotter{p: p}
}

func (p *Plotter) Plot(X, Y []float64, lineName string) error {
	if len(X) != len(Y) {
		return fmt.Errorf("len(X) != len(Y)")
	}

	XYs := make(plotter.XYs, len(X))

	for i := range X {
		XYs[i].X = X[i]
		XYs[i].Y = Y[i]
	}

	err := plotutil.AddLinePoints(p.p, lineName, XYs)
	if err != nil {
		return fmt.Errorf("error while plotting: %v", err)
	}

	return nil
}

func (p *Plotter) Save(path string) error {
	err := p.p.Save(4*vg.Inch, 4*vg.Inch, path)
	if err != nil {
		return fmt.Errorf("error while saving: %v", err)
	}

	return nil
}
