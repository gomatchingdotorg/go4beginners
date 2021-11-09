# A Simple Command-line Tool

## Golang standard libraries: flag

Go provides us a simple way to build a command-line tool (CLI) using only standard libraries (flag). I will help you step by step walk through the process to make a CLI.

In this example, we will create a CLI tool that will ... We will describe in detail the function of the tool in detail below:

First, please take a look at our example.

```bash
	textPtr := flag.String("text", "", "Hello! Welcome to gomatching.org")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
	flag.Parse()

	fmt.Printf("textPtr: %s,  uniquePtr: %t\n", *textPtr, *uniquePtr)
```

The command below will run and test our example.

```bash
$ go build
$ ./main
```

The result is:

```
textPtr: ,  uniquePtr: false
```

```bash
$ go build
$ ./main  --text "hello gomatching"
```

The result is:

```
textPtr: hello gomatching,  uniquePtr: false
```

2. Let's me explain to you about the code

## Third-party libraries:

### urfave/cli:

@https://github.com/urfave/cli

This library helps developers build command-line apps easier by many useful functions. Some of them are making help documentation,

** Getting started **

```bash
$ go get "github.com/urfave/cli/v2"
```

** Examples **

1. Make help document

```bash
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
```

Run our CLI

```bash
$ go build
$ ./cli help
```

Cli generates help text below:

```
NAME:
   Hello - Say Hello!

USAGE:
   cli [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

2. Using arguments
   Using Args function con cli.Context to get the arguments:

```bash
fmt.Printf("Hello %q", c.Args().Get(0))
```

Run our CLI

```bash
$ go build
$ ./cli Tian
```

The result is:

```bash*** GOMATCHING ***
Hello "Tian"%
```

3. Flags
   Using cli.Context to query flags:

```bash
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
```

The result is:
./cli

```bash
	Register new database service node.gomatching.org
```

./cli chat

```bash
	Register new database service chat
```

## License

[MIT](https://www.gomatching.org/license)
[SOURCE](https://www.thienhang.com/2021/11/go-for-beginners-simple-command-line.html)
