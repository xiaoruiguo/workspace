//create: 2015/09/11 13:26:08 change: 2015/09/15 15:06:34 author:lijiao
package main

import(
	"github.com/garyburd/redigo/redis"
	"github.com/docopt/docopt-go"
	"log"
	"fmt"
)

func main() {
	  usage := `Naval Fate.
Usage:
	dcopt --saddr=SADDR --daddr=DADDR --hash=HASH
Options:
	--saddr=SADDR    ip:port
	--daddr=DADDR    ip:port
	--hash=HASH      hashkey
	`
	arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
	fmt.Println(arguments)

	saddr, _  := arguments["--saddr"].(string)
	daddr, _  := arguments["--daddr"].(string)
	hashkey,_ := arguments["--hash"].(string)

	log.SetFlags(log.Lshortfile)
	sc,err := redis.Dial("tcp", saddr)
	if err != nil{
		log.Fatal(err)
	}

	dc,err := redis.Dial("tcp", daddr)
	if err != nil{
		log.Fatal(err)
	}

	values, err := redis.Strings(sc.Do("HGETALl",hashkey))
	if err != nil{
		log.Fatal(err)
	}
	length := len(values)
	fmt.Printf("Len:%d\n",length)
	if length%2 != 0{
		log.Fatal("Length is wrong!")
	}

	for i:=0;i<length;i=i+2{
		fmt.Printf("%s:  %s\n", values[i],values[i+1])
		_, err = dc.Do("HSET", hashkey, values[i], values[i+1])
		if err != nil{
			log.Fatal(err)
		}
	}

	sc.Close()
	dc.Close()
}
