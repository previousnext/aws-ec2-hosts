#!/usr/bin/make -f

VERSION=$(shell git describe --tags --always)
IMAGE=previousnext/aws-ec2-hosts

release: build push

build:
	docker build -t ${IMAGE}:${VERSION} .

push:
	docker push ${IMAGE}:${VERSION}
