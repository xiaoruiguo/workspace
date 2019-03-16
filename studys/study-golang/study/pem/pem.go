//create: 2015/09/09 11:16:51 change: 2015/09/09 11:33:37 author:lijiao
package main

import(
	"encoding/pem"
	"io/ioutil"
	"log"
	"fmt"
)

func init(){
	log.SetFlags(log.Lshortfile)
}

func display(p *pem.Block){
	fmt.Printf("%s\n", p.Type)
	for key,val := range p.Headers{
		fmt.Printf("%s\n\t%s\n", key, val)
	}
}

func main() {
	f,err := ioutil.ReadFile("cas.pem")
	if err != nil || len(f) ==  0{
		log.Fatal(err)
	}

	var p *pem.Block

	for len(f) > 0{
		p, f =  pem.Decode(f)
		if p == nil{
			log.Fatal(err)
		}else{
			display(p)
		}
	}
}
