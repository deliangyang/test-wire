package main

import (
	"log"

	"github.com/deliangyang/test-wire/internal/app"
)

func init() {
	if err := app.InitApplication(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	test := app.GetApplication().Test
	test.Print()
}
