//create: 2015/07/16 10:29:14 change: 2015/07/21 09:54:42 author:lijiao
package dockeropt

import(
	"testing"
	"log"
	"os"
	"github.com/fsouza/go-dockerclient"
	"time"
	"fmt"
)

var (
	err    error
	client *docker.Client
)
func TestDefaultConConfig(t *testing.T){
	config := DefaultConConfig()
	config.SetClient(client)
	config.SetImage("192.168.202.240:5000/liangqiushi/redis",
					"192.168.202.240:5000", "latest")
	config.SetName("redis")
	config.SetAuth("", "", "", "")
	config.SetMemLimit(0,-1)
	config.SetNetworkMode("host")
	config.SetCmds("--port", "7000")

	err = config.Create()
	if err != nil{
		t.Fatal(err)
	}

	t.Log(config.Container)
	
	err = config.Start()
	if err != nil{
		t.Error(err)
	}

	conlist,err := config.ListAll()
	if err != nil{
		t.Error(err)
	}
	for _,v := range conlist{
		log.Println(v)
	}

	names,err := config.ListAllName()
	if err != nil{
		t.Error(err)
	}
	log.Println(names)

	err = config.Remove()
	if err != nil{
		t.Log(err)
	}
}

func TestRemoveByName(t *testing.T){
	config := DefaultConConfig()
	config.SetClient(client)
	config.SetImage("192.168.202.240:5000/liangqiushi/redis",
					"192.168.202.240:5000", "latest")
	config.SetNetworkMode("host")
	config.SetCmds("--port", "7000")
	config.SetName("redis-remove")

	err = config.Create()
	if err != nil{
		t.Fatal(err)
	}

	err = config.Start()
	if err != nil{
		t.Error(err)
	}

	fmt.Println("Test RemoveByName Start sleep: 15 sec")
	time.Sleep(15*time.Second)

	err := RemoveByName(client, "redis-remove", true, true)
	if err != nil{
		t.Error(err)
	}
}

func TestMain(m *testing.M){
	client,err = docker.NewClient("tcp://192.168.183.59:2000")
	if err != nil{
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}
