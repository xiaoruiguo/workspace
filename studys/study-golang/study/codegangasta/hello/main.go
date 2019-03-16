package main

import(
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Hello"
	app.Usage = "Welcome use this!"
	app.Action = func(c *cli.Context){
		println("Hello world!")
	}
	app.Run(os.Args)
}
