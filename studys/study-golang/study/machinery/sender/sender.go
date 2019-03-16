//create: 2017/11/06 13:44:20 change: 2017/11/06 14:49:02 author:lijiao
package main

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
)

var cnf = &config.Config{
	Broker:        "amqp://guest:guest@localhost:5672/",
	DefaultQueue:  "machinery_tasks",
	ResultBackend: "redis://127.0.0.1:6379",
	AMQP: &config.AMQPConfig{
		Exchange:     "machinery_exchange",
		ExchangeType: "direct",
		BindingKey:   "machinery_task",
	},
}

func HelloWorld(arg string) (string, error) {
	str := "Hi, i'm worker@localhost， i has received : " + arg
	return str, nil
}

func main() {
	//init server
	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}

	//regist task
	server.RegisterTask("HelloWorld", HelloWorld)

	//task signature
	signature := &tasks.Signature{
		Name: "HelloWorld",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "task1",
			},
		},
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Fatal(err)
	}

	res, err := asyncResult.Get(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", res)
}
