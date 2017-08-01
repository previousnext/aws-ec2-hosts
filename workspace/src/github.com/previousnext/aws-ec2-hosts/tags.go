package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Flatten tags down to a string, ideal for comparison operations.
func flatTags(tags []*ec2.Tag) []string {
	var newTags []string

	for _, tag := range tags {
		newTags = append(newTags, fmt.Sprintf("%s=%s", *tag.Key, *tag.Value))
	}

	return newTags
}

func tagValue(tags []*ec2.Tag, key string) (string, error) {
	for _, tag := range tags {
		if *tag.Key == key {
			return *tag.Value, nil
		}
	}

	return "", fmt.Errorf("")
}
