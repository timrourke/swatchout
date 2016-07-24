package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/timrourke/swatchout/kmeans"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

type rgbaOutput struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

func main() {
	// Take in `-numcolors={int} "{path_to_file.jpg}"` command-line args
	numColors := flag.Int("numcolors", 5, "Number of colors to generate")
	flag.Parse()
	filePath := flag.Args()
	if len(filePath) == 0 {
		log.Fatal("Must specify a filename.")
	}

	// Read in image
	reader, err := os.Open(filePath[0])
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// Decode image
	originalImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Resize image to fit within 200x200 pixels (larger images do not improve
	// the selection of colors for a palette significantly)
	m := resize.Resize(200, 200, originalImage, resize.NearestNeighbor)
	bounds := m.Bounds()

	// Construct a slice containing rgba values for every pixel in the image
	lengthX := bounds.Max.X - bounds.Min.X
	lengthY := bounds.Max.Y - bounds.Min.Y
	dimensions := lengthX * lengthY
	pixels := make([][]float64, dimensions, dimensions)
	pixelNum := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			rF := float64(r)
			gF := float64(g)
			bF := float64(b)
			aF := float64(a)
			pixels[pixelNum] = []float64{rF, gF, bF, aF}
			pixelNum++
		}
	}

	// Run a k-means algorithm to group colors into clusters
	swatch, _ := kmeans.Cluster(pixels, *numColors)

	// Build structs for json output
	output := make([]rgbaOutput, *numColors, *numColors)
	for idx, rgba := range swatch {
		color := rgbaOutput{
			R: rgba[0] / 256,
			G: rgba[1] / 256,
			B: rgba[2] / 256,
			A: (rgba[3] / 255) - 256,
		}
		output[idx] = color
	}

	// Encode json
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(jsonOutput)
}
