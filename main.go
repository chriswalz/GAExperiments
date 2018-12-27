package main

import (
	"fmt"
	"github.com/chriswalz/GAExperiments/genomes"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/MaxHalford/eaopt"
)


func main() {
	ga := setup()

	// Find the minimum
	err := ga.Minimize(genomes.RandomImagePerformance)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setup() *eaopt.GA {
	// setup out folders for image output
	err := os.RemoveAll("out/raw")
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll("out/raw",os.FileMode(int(0777)))
	if err != nil {
		panic(err)
	}

	const popSize = 2000
	// Instantiate a GA with a GAConfig
	ga, err := eaopt.GAConfig{
		NPops:        1,
		PopSize:      popSize,
		HofSize:      popSize * .1, // determines how many of the best individuals should be recorded.
		NGenerations: 2000,
		Model: eaopt.ModGenerational{
			Selector: eaopt.SelTournament{
				NContestants: uint(1.0 * 100),
			},
			MutRate:   0.5,
			CrossRate: 0.7,
		},
		ParallelEval: false,
	}.NewGA()
	if err != nil {
		panic(err)
	}

	const theoreticalMax = 325000.0 // max possible value for current evaluate function
	fitnesses := make([]float64, 0, 1000)
	ga.Callback = func(ga *eaopt.GA) {
		i := int(ga.Generations)
		path := "out/raw/" + strconv.Itoa(i) + "deleteme.jpg"

		fmt.Printf("Best fitness at generation %d: %f\n", i, ga.HallOfFame[0].Fitness)

		// save top Individual after each generation (overwrites file each time) 
		ga.HallOfFame[0].Genome.(genomes.ImageRandomGenomePerformance).SaveImage(path, "Generation: " + strconv.Itoa(i))

		// save graph of results after each generation (overwrites file each time)
		fitnesses = append(fitnesses, ga.HallOfFame[0].Fitness/theoreticalMax)
		title := fmt.Sprintf("Genome:%s,Model:%s,Popsize:%d", reflect.TypeOf(ga.HallOfFame[0].Genome), reflect.TypeOf(ga.Model), popSize)
		createLinePlot("out/graphs/", title, ga.Age,fitnesses)
		ga.HallOfFame[0].Genome.(genomes.ImageRandomGenomePerformance).SaveImage("out/graphs/" + title + "Winner.jpg", "")
	}
	count := 0
	prevFitness := 0.0
	ga.EarlyStop = func(ga *eaopt.GA) bool {
		currFitness := ga.HallOfFame[0].Fitness
		if currFitness != prevFitness {
			count = 0
			prevFitness = currFitness
			return false
		}
		prevFitness = currFitness
		count++
		if count > 5 {

			return true
		}
		return false
	}
	return ga
}

func createLinePlot(path string, title string, runTime time.Duration, series []float64)  {

	plt, err := plot.New()
	if err != nil {
		panic(err)
	}

	plt.Title.Text = fmt.Sprintf("%s,Time:%.0fs", title, runTime.Seconds())
	plt.X.Label.Text = "Generations"
	plt.Y.Label.Text = "Fitness"

	plt.Y.Max = 0
	plt.Y.Min = -1.0

	// create plotterXY points
	pts := make(plotter.XYs, len(series))
	for i := range pts {
		if i == 0 {
			pts[i].X = float64(i)
			pts[i].Y = series[i]
		} else {
			pts[i].X = float64(i)
			pts[i].Y = series[i]
		}
	}
	// end create plotterXY points


	err = plotutil.AddLinePoints(plt,
		"First", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := plt.Save(8*vg.Inch, 8*vg.Inch, path + title + ".png"); err != nil {
		panic(err)
	}

}