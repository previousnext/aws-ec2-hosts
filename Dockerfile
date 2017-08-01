FROM golang:1.8
ADD workspace /go
RUN go get github.com/mitchellh/gox
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/aws-ec2-hosts_linux_amd64 /usr/local/bin/aws-ec2-hosts
CMD ["aws-ec2-hosts"]
