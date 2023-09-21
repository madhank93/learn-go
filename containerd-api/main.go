// package main

// import (
// 	"context"

// 	"github.com/containerd/containerd"
// 	"github.com/containerd/containerd/namespaces"
// 	"github.com/containerd/containerd/oci"
// )

// func main() {
//     // Connect to the containerd daemon
//     client, err := containerd.New("/run/containerd/containerd.sock")
//     if err != nil {
//         panic(err)
//     }
//     defer client.Close()

//     // Define a namespace for the container
//     ctx := namespaces.WithNamespace(context.Background(), "my-namespace")

//     // Pull an OCI image (replace "your-image-name" with the actual image name)
//     imageRef := "nginx"
//     image, err := client.Pull(ctx, imageRef, containerd.WithPullUnpack)
//     if err != nil {
//         panic(err)
//     }

//     // Create a new container
//     containerID := "my-container"
//     container, err := client.NewContainer(ctx, containerID,
//         containerd.WithNewSpec(oci.WithImageConfig(image)),
//     )
//     if err != nil {
//         panic(err)
//     }

//     // Create a new task (process) within the container
//     task, err := container.NewTask(ctx, nil)
//     if err != nil {
//         panic(err)
//     }

//     // Start the task
//     err = task.Start(ctx)
//     if err != nil {
//         panic(err)
//     }

//     // Optionally, you can execute commands in the container
// }

package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
        if err := redisExample(); err != nil {
                log.Fatal(err)
        }
}

func redisExample() error {
        client, err := containerd.New("/run/containerd/containerd.sock")
        if err != nil {
                return err
        }
        defer client.Close()

        ctx := namespaces.WithNamespace(context.Background(), "example")
        image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
        if err != nil {
                return err
        }
        log.Printf("Successfully pulled %s image\n", image.Name())

        return nil
}