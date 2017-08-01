package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Helper function to query for our /etc/hosts instances.
//  * Query instances with tag X
//  * Group by Y tag
//  * Only return a list that contains 1 instance per group
func queryInstanaces(region, tagFilter, tagGroup string) ([]host, error) {
	var instances []host

	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))

	// Look up all the Instances from the provided query.
	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		return instances, err
	}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if *instance.State.Name != ec2.InstanceStateNameRunning {
				continue
			}

			// Check if we can filter this instance out.
			if !contains(flatTags(instance.Tags), tagFilter) {
				continue
			}

			// Check that this reservation has our tag.
			name, err := tagValue(instance.Tags, tagGroup)
			if err != nil {
				continue
			}

			instances = append(instances, host{
				Name: name,
				IP:   *instance.PrivateIpAddress,
			})
		}
	}

	return instances, nil
}
