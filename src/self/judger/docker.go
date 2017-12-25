/**
 * Created by shiyi on 2017/12/17.
 * Email: shiyi@fightcoder.com
 */

package judger

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerCli struct {
	cli *client.Client
}

func NewDockerCli() DockerCli {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return DockerCli{cli}
}

func (this *DockerCli) RunContainer(imageName string, cmd []string, workDir string) (code int64, msg string) {
	ctx := context.Background()

	containerBody, err := this.cli.ContainerCreate(ctx,
		&container.Config{
			Image: imageName,
			Cmd:   cmd,
			User:  "root",
		}, &container.HostConfig{
			Binds: []string{
				workDir + ":/workspace",
			},
		}, nil, "")
	if err != nil {
		panic(err)
	}

	if err := this.cli.ContainerStart(ctx, containerBody.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("container ", containerBody.ID, "has start")
	statusCh, errCh := this.cli.ContainerWait(ctx, containerBody.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case status := <-statusCh:
		code = status.StatusCode
		if status.Error != nil {
			msg = status.Error.Message
		}
		return code, msg
	}

	return code, msg
}

func (this *DockerCli) KillContainer(containerId string) {
	ctx := context.Background()

	err := this.cli.ContainerKill(ctx, containerId, "SIGKILL")
	if err != nil {
		panic(err)
	}
}

func (this *DockerCli) ListContainers() []types.Container {
	ctx := context.Background()

	containers, err := this.cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	//fmt.Println(len(containers))
	//
	//for _, container := range containers {
	//	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	//}

	return containers
}

func (this *DockerCli) ListImages() []types.ImageSummary {
	ctx := context.Background()

	images, err := this.cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//for _, image := range images {
	//	fmt.Println(image.ID)
	//}

	return images
}

func (this *DockerCli) PullImage(imageName string) {
	ctx := context.Background()

	out, err := this.cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(os.Stdout, out)
}

func (this *DockerCli) RemoveContainer(containerID string, force bool, removeVolumes bool, removeLinks bool) {
	ctx := context.Background()

	options := types.ContainerRemoveOptions{Force: force, RemoveVolumes: removeVolumes, RemoveLinks: removeLinks}
	if err := this.cli.ContainerRemove(ctx, containerID, options); err != nil {
		panic(err)
	}
}

func (this *DockerCli) PrintLogContainer(containerID string) {
	ctx := context.Background()

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := this.cli.ContainerLogs(ctx, containerID, options)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(os.Stdout, out)
}

func (this *DockerCli) StartContainer(containerID string) {
	err := this.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	//fmt.Println("容器", containerID, "启动成功")
}

// 停止
func (this *DockerCli) StopContainer(containerID string) {
	timeout := time.Second * 10
	if err := this.cli.ContainerStop(context.Background(), containerID, &timeout); err != nil {
		panic(err)
	}

	//fmt.Printf("容器%s已经被停止\n", containerID)
}
