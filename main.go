package main

import (
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/rancher/go-rancher/v2"
	"github.com/rancherlabs/kattle/events"
	"github.com/rancherlabs/kattle/handlers"
	"github.com/rancherlabs/kattle/hostname"
	"github.com/rancherlabs/kattle/watch"
	"github.com/urfave/cli"
)

const (
	cattleURLEnv       = "CATTLE_URL"
	cattleAccessKeyEnv = "CATTLE_ACCESS_KEY"
	cattleSecretKeyEnv = "CATTLE_SECRET_KEY"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "kattle"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "kubernetes-master",
		},
		cli.StringFlag{
			Name: "username",
		},
		cli.StringFlag{
			Name: "password",
		},
		cli.StringFlag{
			Name: "token",
		},
	}
	app.Action = action
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func action(c *cli.Context) error {
	rancherClient, err := createRancherClient()
	if err != nil {
		return err
	}

	// TODO: this is bad
	hostname.RancherClient = rancherClient

	kubernetesURL := c.String("kubernetes-master")
	username := c.String("username")
	password := c.String("password")
	token := c.String("token")
	clientset, err := createKubernetesClient(kubernetesURL, username, password, token)
	if err != nil {
		return err
	}

	watchClient := watch.NewClient(clientset)
	watchClient.Start()

	time.Sleep(5 * time.Second)

	handlers.WatchClient = watchClient
	handlers.Clientset = clientset

	cattleURL := os.Getenv(cattleURLEnv)
	cattleAccessKey := os.Getenv(cattleAccessKeyEnv)
	cattleSecretKey := os.Getenv(cattleSecretKeyEnv)

	return events.Listen(cattleURL, cattleAccessKey, cattleSecretKey, 250)
}

func createRancherClient() (*client.RancherClient, error) {
	url, err := client.NormalizeUrl(os.Getenv(cattleURLEnv))
	if err != nil {
		return nil, err
	}
	return client.NewRancherClient(&client.ClientOpts{
		Url:       url,
		AccessKey: os.Getenv(cattleAccessKeyEnv),
		SecretKey: os.Getenv(cattleSecretKeyEnv),
	})
}

func createKubernetesClient(url, username, password, token string) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(&rest.Config{
		Host:        url,
		Username:    username,
		Password:    password,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	})
}
