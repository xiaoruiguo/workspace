//create: 2015/08/22 19:05:35 change: 2015/08/22 20:02:14 author:lijiao
package main

import(
	"fmt"
	"time"
)

func main() {
	select{
		default:
			fmt.Println("1")
			time.Sleep(time.Second)
	}

	var NeverStop <-chan struct{} = make(chan struct{})
	select{
		case <-NeverStop:
			fmt.Println("Never")
		default:
			fmt.Println("3")
			time.Sleep(time.Second)
	}

	var c chan int
	c = make(chan int)
	go func(){
		for{
			select{
				case c<-8:
					fmt.Println("Send 8")
			}
			time.Sleep(time.Second)
		}
	}()
	for{
		select{
			case a:=<-c:
				fmt.Println(a)
		}
		time.Sleep(time.Second)
	}
}


