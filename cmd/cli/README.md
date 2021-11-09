# A Simple Command-line Tool

## Golang standard libraries: flag

Go provides us a simple way to build a commad-line tool (CLI) using only standard libraries (flag). I will help you step by step walk through the process to make a CLI.

In this exxample, we will create a CLI tool XXX that will ... We will describe in detail about the function of the tool in detail below:

1. First, please take a look with our example.
   The command bellow will run and test our example.

```bash
$ go build
$ ./main
```

2. Let 's me explain you about the code

## Third party libraries:

### urfave/cli:

@https://github.com/urfave/cli

Such library helps developers building command line apps easier by many uasage functions.
Some of them are making help doccumentation,

** Getting started **

```bash
$ go get "github.com/urfave/cli/v2"
```

** Examples **

1. Make help doccument

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

Cli generates help text bellow:

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
