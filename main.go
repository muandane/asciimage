package main

import (
	"fmt"
	_ "image/jpeg"
	"runtime"

	paintbrush "github.com/jordanella/go-ansi-paintbrush"
)

func main() {

	// Create a new AnsiArt instance
	canvas := paintbrush.New()

	err := canvas.LoadImage("examples/windowsxp.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add more characters and adjust weights as desired
	weights := map[rune]float64{
		'': .95,
		'': .95,
		'▁': .9,
		'▂': .9,
		'▃': .9,
		'▄': .9,
		'▅': .9,
		'▆': .85,
		'█': .85,
		'▊': .95,
		'▋': .95,
		'▌': .95,
		'▍': .95,
		'▎': .95,
		'▏': .95,
		'●': .95,
		'◀': .95,
		'▲': .95,
		'▶': .95,
		'▼': .9,
		'○': .8,
		'◉': .95,
		'◧': .9,
		'◨': .9,
		'◩': .9,
		'◪': .9,
	}

	canvas.Weights = weights

	// use all available CPU cores for rendering
	canvas.SetThreads(runtime.NumCPU())

	// Start the rendering process
	canvas.Paint()

	fmt.Printf("\r%s", canvas.Result)
}
