package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//HowToUseUrfaveV2_Hello()
	HowToUseUrfaveV2_Flag()
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

func HowToUseUrfaveV2_Flag() {
	var nodeType string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "node",
				Value:       "cluster.gomatching.org",
				Usage:       "register new node",
				Destination: &nodeType,
			},
		},
		Action: func(c *cli.Context) error {
			name := "node.gomatching.org"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("node") == "chat" {
				fmt.Println("Register new chat service: ", name)
			} else {
				fmt.Println("Register new database service", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
