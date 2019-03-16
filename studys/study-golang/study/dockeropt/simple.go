//create: 2015/07/20 17:59:57 change: 2015/07/21 13:28:24 author:lijiao
package dockeropt

import(
	"github.com/fsouza/go-dockerclient"
	"log"
	"fmt"
	"time"
)

func GoClientExample() error{

	cmds := make([]string, 0)
	cmds = append(cmds, "--port", "7000")

	volumes := make(map[string]struct{})
	volumes["/export"]=struct{}{}

	image := docker.PullImageOptions{
		Repository:     "192.168.202.240:5000/liangqiushi/redis"   ,
		Registry:       "192.168.202.240:5000"  ,
		Tag:            "latest"      ,
	}

	auth := docker.AuthConfiguration{
		Username:       ""        ,
		Password:       ""        ,
		Email:          ""        ,
		ServerAddress:  ""        ,
	}

	config := docker.Config{
		Memory:         30*1024*1024*1024   ,
		Cmd:            cmds                ,
		Volumes:        volumes             ,
		Image:          "192.168.202.240:5000/liangqiushi/redis:latest",
	}

	binds := make([]string,0)
	binds = append(binds, "/export/Data/:/export")
	hostconfig := docker.HostConfig{
		Binds:          binds                  ,
		NetworkMode:    "host"                 ,
		Memory:         30*1024*1024*1024      ,
	}

	createoption := docker.CreateContainerOptions{
		Name:         "ip_port"      ,
		Config:       &config         ,
		HostConfig:   &hostconfig     ,
	}

	client,err := docker.NewClient("tcp://192.168.192.37:2375")

	if err != nil{
		log.Fatalln(err)
	}

	var con  *docker.Container

	con, err = client.CreateContainer(createoption)

	if err == docker.ErrNoSuchImage{
		err = client.PullImage(image, auth)
		if err != nil{
			return err
		}
		con, err = client.CreateContainer(createoption)
	}

	if err != nil{
		return err
	}

	err = client.StartContainer(con.ID, createoption.HostConfig)

	if err != nil{
		return err
	}

	fmt.Println("Test RemoveByName Start sleep")
	fmt.Println("You have 30 sec to check the containe created by this test")
	time.Sleep(30*time.Second)

	removeoption := docker.RemoveContainerOptions{
		ID: "ip_port",     //Both "Name" and "ID" can work
		RemoveVolumes:true,
		Force: true,
	}

	err = client.RemoveContainer(removeoption)

	if err != nil{
		return err
	}

	return nil
}

