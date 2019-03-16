//create: 2015/06/02 17:22:01 change: 2015/06/02 17:59:55 author:lijiao
package main

import(
	"github.com/codegangsta/cli"
	"os"
)


func main() {
	app := cli.NewApp()
	app.Name ="TestCli"
	app.Usage = "Show you how to use TestCli"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang,l",          //--lang, -l
			Value: "english",
			Usage: "choose a language",
			EnvVar: "LANG_ENV,LANG",  //set the default value by $LANG_ENV, if not exist,then $LANG
		},
	}
	app.Action = func(c *cli.Context){
		println("TestCli is working!")
		print("You have choose english",c.String("lang"),"\n")
	}

	app.Commands = []cli.Command{
		{
			Name:    "SubCmd",
			Aliases: []string{"a"},
			Usage:   "This is a SubCmd",
			Action: func(c *cli.Context){
				println("This is a SubCmd")
			},
		},
	}
	app.Run(os.Args)
}
