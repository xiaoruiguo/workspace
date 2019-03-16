// Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/kubernetes/pkg/kubelet/cadvisor"
	"time"
)

type CmdLine struct {
	ContainerName string
}

var cmdline CmdLine

func init() {
	flag.StringVar(&cmdline.ContainerName, "name", "/", "container name")
}

func main() {
	flag.Parse()

	ContainerRuntime := "docker"
	RemoteRuntimeEndpoint := ""
	imagefs := cadvisor.NewImageFsInfoProvider(ContainerRuntime, RemoteRuntimeEndpoint)
	usingLegacyStats := cadvisor.UsingLegacyCadvisorStats(ContainerRuntime, RemoteRuntimeEndpoint)
	ca, err := cadvisor.New("127.0.0.1", 4000, imagefs, "/var/lib/kubelet", usingLegacyStats)
	if err != nil {
		glog.Error(err)
	}

	if err := ca.Start(); err != nil {
		glog.Error(err)
	}

	options := cadvisorapiv2.RequestOptions{
		IdType:    cadvisorapiv2.TypeName,
		Count:     2, // 2 samples are needed to compute "instantaneous" CPU
		Recursive: true,
	}

	machineInfo, err := ca.MachineInfo()
	if err != nil {
		glog.Error(err.Error())
	}
	bytes, err := json.Marshal(&machineInfo)
	fmt.Printf("machineInfo: %s\n", bytes)

	rootfsInfo, err := ca.RootFsInfo()
	if err != nil {
		glog.Error(err.Error())
	}
	bytes, err = json.Marshal(&rootfsInfo)
	fmt.Printf("rootfsInfo: %s\n", bytes)

	infoMap, err := ca.ContainerInfoV2(cmdline.ContainerName, options)
	if err != nil {
		glog.Error(err.Error())
	}

	if infoMap == nil {
		glog.Error("infoMap is null")
	}

	for name, value := range infoMap {
		fmt.Printf("%s: %s  %s\n", name, value.Spec.Namespace, value.Spec.Image)
	}

	time.Sleep(600 * time.Second)
}
