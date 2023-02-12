package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/edualb-challenge/treebabel/internal/apps"
)

func main() {
	var file string
	var wSize uint64

	flag.StringVar(&file, "input_file", "", "set the input file to process the average delivery time")
	flag.Uint64Var(&wSize, "window_size", 0, "set the window size is the size in minutes to be considered")
	flag.Parse()

	app, err := apps.NewTreeBabel(file, wSize)
	if err != nil {
		log.Fatal(err)
	}
	avgDeliveryTime, err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

	for _, adt := range avgDeliveryTime {
		b, err := json.Marshal(adt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
