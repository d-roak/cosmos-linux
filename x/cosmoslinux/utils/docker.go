package utils

import (
	"bufio"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetDockerContainer(name string) (types.Container, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
        All: true,
    })
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
        for _, containerName := range container.Names {
            if containerName == "/" + name {
		    	return container, nil
            }
		}
	}

	return types.Container{}, errors.New("Container not found")
}

func ExecuteDockerContainer(machineId string, commands []string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	_, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

    commands_str := strings.Join(commands[:], "&&")
    commands_hash := fmt.Sprintf("%x", md5.Sum([]byte(commands_str)))

    cont, err := GetDockerContainer(machineId + "_" + commands_hash)
    if err == nil {
	    out, err := cli.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{ShowStdout: true})
	    if err != nil {
	    	panic(err)
	    }
        
	    outString := ""
	    s := bufio.NewScanner(out)
	    for s.Scan() {
	    	outString += fmt.Sprintf("%s\n", s.Text())
	    }
        
	    return outString, nil
    }
	
    resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:           "alpine",
		Cmd:             []string{"sh", "-c", commands_str},
		Tty:             true,
		NetworkDisabled: true,
	}, nil, nil, nil, machineId + "_" + commands_hash)
	if err != nil {
		panic(err)
	}


	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	outString := ""
	s := bufio.NewScanner(out)
	for s.Scan() {
		outString += fmt.Sprintf("%s\n", s.Text())
	}

	return outString, nil
}
