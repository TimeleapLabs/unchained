package runtime

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"log"
)

func RunDocker(ctx context.Context, image string, params []byte) ([]byte, error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// Define the container configuration
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{"3", "5"},
		Tty:   false,
	}, nil, nil, nil, fmt.Sprintf("unchained-%s", image))
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
	}

	// Start the container
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Error starting container: %v", err)
	}

	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("Error waiting for container: %v", err)
		}
	case <-statusCh:
	}

	// Get the container logs
	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
	if err != nil {
		log.Fatalf("Error getting logs: %v", err)
	}

	// Remove the container
	if err := cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
		log.Fatalf("Error removing container: %v", err)
	}

	result, err := io.ReadAll(out)
	if err != nil {
		log.Fatalf("Error reading logs: %v", err)
	}

	return result[7:], nil
}
