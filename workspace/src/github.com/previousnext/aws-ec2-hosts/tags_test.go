package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/stretchr/testify/assert"
)

func TestFlatTags(t *testing.T) {
	tags := []*ec2.Tag{
		{
			Key:   aws.String("FOO"),
			Value: aws.String("bar"),
		},
		{
			Key:   aws.String("BAR"),
			Value: aws.String("baz"),
		},
	}

	want := []string{
		"FOO=bar",
		"BAR=baz",
	}

	assert.Equal(t, want, flatTags(tags))
}

func TestTagValue(t *testing.T) {
	tags := []*ec2.Tag{
		{
			Key:   aws.String("FOO"),
			Value: aws.String("bar"),
		},
		{
			Key:   aws.String("BAR"),
			Value: aws.String("baz"),
		},
	}

	value, err := tagValue(tags, "FOO")
	assert.Nil(t, err)
	assert.Equal(t, "bar", value)
}
