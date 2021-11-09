package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	HowToUseUrfaveV2_Hello()
}

func HowToUseFlag() {
	textPtr := flag.String("text", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}

func HowToUseUrfaveV2_Hello() {
	app := &cli.App{
		Name:  "Hello",
		Usage: "Say Hello!",
		Action: func(c *cli.Context) error {
			fmt.Println("*** GOMATCHING ***")
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
