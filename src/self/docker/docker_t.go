package docker

//
//import (
//	"context"
//	"io"
//	"os"
//	"time"
//
//	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
//	"github.com/docker/docker/api/types/mount"
//	"github.com/docker/docker/client"
//)
//
//func NewDockerCli() *client.Client {
//	cli, err := client.NewEnvClient()
//	if err != nil {
//		panic(err)
//	}
//
//	return cli
//}
//
//func PullImage(cli *client.Client, imageName string) {
//	ctx := context.Background()
//
//	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
//	if err != nil {
//		panic(err)
//	}
//	defer out.Close()
//
//	//io.Copy(os.Stdout, out)
//}
//
//func CreateContainer(cli *client.Client, imageName string, cmd []string) string {
//	ctx := context.Background()
//
//	resp, err := cli.ContainerCreate(ctx,
//		&container.Config{
//			Image: imageName,
//			Cmd:   cmd,
//			User:  "root",
//		}, &container.HostConfig{
//
//			Mounts: []mount.Mount{
//				{
//					Type:   mount.TypeVolume,
//					Source: "Users",
//					Target: "/appa",
//				},
//			},
//		}, nil, "")
//	if err != nil {
//		panic(err)
//	}
//
//	return resp.ID
//}
//
//func RemoveContainer(cli *client.Client, containerID string, force bool, removeVolumes bool, removeLinks bool) {
//	ctx := context.Background()
//
//	options := types.ContainerRemoveOptions{Force: force, RemoveVolumes: removeVolumes, RemoveLinks: removeLinks}
//	if err := cli.ContainerRemove(ctx, containerID, options); err != nil {
//		panic(err)
//	}
//}
//
//func PrintLogContainer(cli *client.Client, containerID string) {
//	ctx := context.Background()
//
//	options := types.ContainerLogsOptions{ShowStdout: true}
//	out, err := cli.ContainerLogs(ctx, containerID, options)
//	if err != nil {
//		panic(err)
//	}
//	defer out.Close()
//
//	io.Copy(os.Stdout, out)
//}
//
//func StartContainer(cli *client.Client, containerID string) {
//	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
//	if err != nil {
//		panic(err)
//	}
//
//	//fmt.Println("容器", containerID, "启动成功")
//}
//
//// 停止
//func StopContainer(cli *client.Client, containerID string) {
//	timeout := time.Second * 10
//	if err := cli.ContainerStop(context.Background(), containerID, &timeout); err != nil {
//		panic(err)
//	}
//
//	//fmt.Printf("容器%s已经被停止\n", containerID)
//}
//
//func ListContainers(cli *client.Client) []types.Container {
//	ctx := context.Background()
//
//	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
//	if err != nil {
//		panic(err)
//	}
//
//	//fmt.Println(len(containers))
//	//
//	//for _, container := range containers {
//	//	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
//	//}
//
//	return containers
//}
//
//func ListImages(cli *client.Client) []types.ImageSummary {
//	ctx := context.Background()
//
//	images, err := cli.ImageList(ctx, types.ImageListOptions{})
//	if err != nil {
//		panic(err)
//	}
//
//	//for _, image := range images {
//	//	fmt.Println(image.ID)
//	//}
//
//	return images
//}
//
////-----TODO Run参数
//
//func Run(cli *client.Client, containerID string) {
//	ctx := context.Background()
//
//	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
//		panic(err)
//	}
//
//	statusCh, errCh := cli.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
//	select {
//	case err := <-errCh:
//		if err != nil {
//			panic(err)
//		}
//	case <-statusCh:
//	}
//}
//
//func RunInBackground(cli *client.Client, containerID string) {
//	ctx := context.Background()
//
//	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
//		panic(err)
//	}
//}
