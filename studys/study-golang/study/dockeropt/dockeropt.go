//create: 2015/07/15 11:57:52 change: 2015/07/20 17:59:55 author:lijiao
package dockeropt

import(
	"github.com/fsouza/go-dockerclient"
)

type ConConfig struct{
	CreateOptions docker.CreateContainerOptions
	RemoveOptions docker.RemoveContainerOptions
	ImageOptions  docker.PullImageOptions
	Auth          docker.AuthConfiguration
	Container     *docker.Container
	Client        *docker.Client
}

func (con *ConConfig) SetClient(client *docker.Client) {
	con.Client = client
}

func (con *ConConfig) SetImage(respos, registry, tag string){
	con.ImageOptions = docker.PullImageOptions{
		Repository:     respos    ,
		Registry:       registry  ,
		Tag:            tag       ,
		OutputStream:   nil       ,
		RawJSONStream:  false     ,
	}
	con.CreateOptions.Config.Image=respos+":"+tag
}

func (con *ConConfig) SetName(name string){
	con.CreateOptions.Name=name
}

func (con *ConConfig) SetAuth(user, pass, mail, addr string){
	con.Auth = docker.AuthConfiguration{
		Username:       user        ,
		Password:       pass        ,
		Email:          mail        ,
		ServerAddress:  addr        ,
	}
}

func (con *ConConfig) SetMemLimit(mem, swap int64){
	con.CreateOptions.HostConfig.Memory = mem
	con.CreateOptions.HostConfig.MemorySwap = swap
	con.CreateOptions.Config.Memory = mem
	con.CreateOptions.Config.MemorySwap = swap
}

func (con *ConConfig) SetNetworkMode(mode string){
	con.CreateOptions.HostConfig.NetworkMode = mode
}

func (con *ConConfig) SetPortSpec(ports ... string){
	con.CreateOptions.Config.PortSpecs = ports
}

func (con *ConConfig) SetCmds(cmd ... string){
	con.CreateOptions.Config.Cmd = cmd
}

func (con *ConConfig) PullImage() error{
	return con.Client.PullImage(con.ImageOptions, con.Auth)
}

func (con *ConConfig) Create() error{
	var err error
	con.Container, err = con.Client.CreateContainer(con.CreateOptions)
	if err == docker.ErrNoSuchImage{
		err = con.Client.PullImage(con.ImageOptions, con.Auth)
		if err != nil{
			return err
		}
		con.Container, err = con.Client.CreateContainer(con.CreateOptions)
	}
	if err != nil{
		return err
	}
	con.RemoveOptions.ID = con.Container.ID
	return nil
}

func (con *ConConfig) Remove()error{
	return con.Client.RemoveContainer(con.RemoveOptions)
}

func (con *ConConfig) Start()error{
	return con.Client.StartContainer(con.Container.ID, con.CreateOptions.HostConfig)
}

func (con *ConConfig) ListAll() ([]docker.APIContainers, error){
	cons, err := con.Client.ListContainers(docker.ListContainersOptions{All:true})
	return cons, err
}

func (con *ConConfig) ListAllName() ([]string, error){
	cons, err := con.Client.ListContainers(docker.ListContainersOptions{All:true})
	if err != nil{
		return nil, err
	}
	names := make([]string, 0)
	for _,c := range cons{
		detail, err := con.Client.InspectContainer(c.ID)
		if err != nil{
			return names, err
		}
		names = append(names, detail.Name)
	}
	return names, nil
}

func DefaultConConfig() *ConConfig {

	restartpolicy:=docker.RestartPolicy{
		Name:"no",
		MaximumRetryCount:0,
	}

	logconfig:=docker.LogConfig{
		Type:"",
		Config:nil,
	}

	config := docker.Config{
		Hostname:       "localhost"         ,
		Domainname:     ""                  ,
		User:           ""                  ,
		Memory:         30*1024*1024*1024   ,
		MemorySwap:     -1                  ,
		CPUShares:      10                  ,
		CPUSet:         ""                  ,
		AttachStdin:    false               ,
		AttachStdout:   false               ,
		AttachStderr:   false               ,
		PortSpecs:      nil                 ,
		ExposedPorts:   nil                 ,
		Tty:            false               ,
		OpenStdin:      false               ,
		StdinOnce:      false               ,
		Env:            nil                 ,
		Cmd:            nil                 ,
		DNS:            nil                 ,
		Image:          ""                  ,
		Volumes:        nil                 ,
		VolumesFrom:    ""                  ,
		WorkingDir:     ""                  ,
		MacAddress:     ""                  ,
		Entrypoint:     nil                 ,
		NetworkDisabled:false               ,
		SecurityOpts:   nil                 ,
		OnBuild:        nil                 ,
		Labels:         nil                 ,
	}

	hostconfig := docker.HostConfig{
		Binds:          nil                    ,
		CapAdd:         nil                    ,
		CapDrop:        nil                    ,
		ContainerIDFile:""                     ,
		LxcConf:        nil                    ,
		Privileged:     false                  ,
		PortBindings:   nil                    ,
		Links:          nil                    ,
		PublishAllPorts:false                  ,
		DNS:            nil                    ,
		DNSSearch:      nil                    ,
		ExtraHosts:     nil                    ,
		VolumesFrom:    nil                    ,
		NetworkMode:    ""                     ,
		IpcMode:        ""                     ,
		PidMode:        ""                     ,
		UTSMode:        ""                     ,
		RestartPolicy:  restartpolicy          ,
		Devices:        nil                    ,
		LogConfig:      logconfig              ,
		ReadonlyRootfs: false                  ,
		SecurityOpt:    nil                    ,
		CgroupParent:   ""                     ,
		Memory:         30*1024*1024*1024      ,
		MemorySwap:     -1                     ,
		CPUShares:      10                     ,
		CPUSet:         ""                     ,
		CPUQuota:       10                     ,
		CPUPeriod:      100000                 ,
		Ulimits:        nil                    ,
	}


	return &ConConfig{
				CreateOptions:docker.CreateContainerOptions{
					Name:"",
					Config: &config,
					HostConfig: &hostconfig,
				},
				RemoveOptions:docker.RemoveContainerOptions{
					ID:"",
					RemoveVolumes:true,
					Force:true,
				},
				Container:nil,
			}
}

func RemoveByName(client *docker.Client, name string, volume, force bool) error{
	op := docker.RemoveContainerOptions{
		ID:name,
		RemoveVolumes:volume,
		Force: force,
	}
	return client.RemoveContainer(op)
}
