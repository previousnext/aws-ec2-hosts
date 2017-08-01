AWS: EC2 Hosts
==============

[![CircleCI](https://circleci.com/gh/previousnext/aws-ec2-hosts.svg?style=svg)](https://circleci.com/gh/previousnext/aws-ec2-hosts)

**Maintainer**: Nick Schuch

Creates `/etc/hosts` records based on EC2 instance tags.

## Development

### Principles

* Code lives in the `workspace` directory

### Tools

* **Dependency management** - https://getgb.io
* **Build** - https://github.com/mitchellh/gox
* **Linting** - https://github.com/golang/lint

### Workflow

(While in the `workspace` directory)

**Installing a new dependency**

```bash
gb vendor fetch github.com/foo/bar
```

**Running quality checks**

```bash
make lint test
```

**Building binaries**

```bash
make build
```
