package main

import (
	"fmt"
	"github.com/mthsrocha/GO/kakashi/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Kakashi aplication...")

	kakashi_aplication := app.Generate()
	if err := kakashi_aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
