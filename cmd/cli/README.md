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

1. urfave/cli
   @https://github.com/urfave/cli

Such library helps developers building command line apps easier by many uasage functions.
Some of them are making help doccumentation,

- Getting started

```bash
$ go get "github.com/urfave/cli/v2"
```

- Example
  > Make help doccument

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

## License

[MIT](https://www.gomatching.org/license)
