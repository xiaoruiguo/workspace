//create: 2015/08/24 16:27:03 change: 2015/09/14 14:54:43 author:lijiao
package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)
/*
Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version
*/

func main() {
	  usage := `Naval Fate.
Usage:
	dcopt --speed=SPEED --sxx=XSS --flag=FLAG
Options:
	--speed=SPEED`
	
	  arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
	  fmt.Println(arguments)
}


