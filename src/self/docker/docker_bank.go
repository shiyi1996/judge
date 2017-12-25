package docker

//
//import (
//	"github.com/samalba/dockerclient"
//	log "github.com/sirupsen/logrus"
//)
//
//type DockerCli struct {
//	*dockerclient.DockerClient
//}
//
//func NewDockerClient() DockerCli {
//	docker, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
//	if err != nil {
//		panic(err)
//	}
//	return DockerCli{docker}
//}
//
//func (docker DockerCli) CompileInContainer(imageName string, workDir string, cmd []string) <-chan dockerclient.WaitResult {
//	containerConfig := &dockerclient.ContainerConfig{
//		Image: imageName,
//		Cmd:   cmd,
//	}
//	containerId, err := docker.CreateContainer(containerConfig, "", nil)
//	if err != nil {
//		panic(err)
//	}
//
//	binds := []string{
//		workDir + ":/workspace",
//	}
//
//	hostConfig := &dockerclient.HostConfig{
//		Binds: binds,
//	}
//	err = docker.StartContainer(containerId, hostConfig)
//	if err != nil {
//		panic(err)
//	}
//
//	return docker.Wait(containerId)
//}
//
//func (docker DockerCli) ListImages() {
//	// Get only running containers
//	containers, err := docker.ListContainers(false, false, "")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, c := range containers {
//		log.Println(c.Id, c.Names)
//	}
//
//	// Inspect the first container returned
//	if len(containers) > 0 {
//		id := containers[0].Id
//		info, _ := docker.InspectContainer(id)
//		log.Println(info)
//	}
//}
