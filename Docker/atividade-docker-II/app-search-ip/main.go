package main

import (
	"app-search-ip/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	app := app.Generate()
	error := app.Run(os.Args)

	if error != nil {
		log.Fatal(error)
	}
}
