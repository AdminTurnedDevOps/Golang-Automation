// Example: go run main.go i-0858366964463b4db

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	instanceID := os.Args[1]

	listInstances(instanceID)
}

func listInstances(instanceID string) {
	awsConnect, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	if err != nil {
		fmt.Println(err)
	}

	ec2sess := ec2.New(awsConnect)

	instanceInfo := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}
	fmt.Println("Listing EC2 Instance info...")

	fmt.Println(ec2sess.DescribeInstances(instanceInfo))
}
