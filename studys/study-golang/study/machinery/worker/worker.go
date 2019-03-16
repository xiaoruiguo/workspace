//create: 2017/11/06 13:44:20 change: 2017/11/06 14:32:42 author:lijiao
package main

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
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

	return "Hi, i'm worker@localhost, just receive task: " + arg, nil
}

func main() {
	//init server
	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}

	//regist task
	server.RegisterTask("HelloWorld", HelloWorld)

	//create worker
	worker := server.NewWorker("worker@localhost", 10)
	err = worker.Launch()
	if err != nil {
		log.Fatal(err)
	}
}
