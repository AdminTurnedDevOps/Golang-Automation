package main

import (
	"fmt"
	//"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	//instanceID := os.Args[1]
	listInstances()
}

func listInstances() {
	fmt.Println("Listing EC2 Instances...")

	awsConnect := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})

	ec2sess := ec2.New(awsConnect)

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String("i-0858366964463b4db")},
	}
	fmt.Println(ec2sess.DescribeInstances(input))
}