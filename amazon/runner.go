package amazon

import (
	"log"
	"time"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

func Run() {
	instances := startServer()
	log.Println("waiting 2 minutes")
	time.Sleep(2 * time.Minute)
	stopInstances(instances)
}

func startServer() []string {
	log.Println("Starting Instance")

	auth := aws.Auth{"", "", ""}
	e := ec2.New(auth, aws.EUWest)
	options := ec2.RunInstances{
		ImageId:      "ami-f0b11187",
		InstanceType: "t2.micro",
	}
	resp, err := e.RunInstances(&options)
	if err != nil {
		panic(err.Error())
	}

	var instances []string

	for _, instance := range resp.Instances {
		instances = append(instances, instance.InstanceId)
	}

	return instances
}

func stopInstances(instances []string) {
	log.Println("Stopping instance")

	auth := aws.Auth{"AKIAINQEUXY3PYAB632Q", "tkG5HBRypOTuRP/xAZyyPyGYtf+w3kFXXg0LNfAk", ""}
	e := ec2.New(auth, aws.EUWest)

	_, err := e.TerminateInstances(instances)
	if err != nil {
		panic(err.Error())
	}
}
