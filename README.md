# Kubectl Node Allocaions

[![CircleCI](https://circleci.com/gh/wshihadeh/kubectl-counts/tree/master.svg?style=svg)](https://circleci.com/gh/wshihadeh/kubectl-counts/tree/master)
[![GoDoc](https://godoc.org/github.com/wshihadeh/kubectl-counts?status.svg)](https://godoc.org/github.com/wshihadeh/kubectl-counts)
[![Go Report Card](https://goreportcard.com/badge/github.com/wshihadeh/kubectl-counts)](https://goreportcard.com/report/github.com/wshihadeh/kubectl-counts)
[![GitHub release](https://img.shields.io/github/release/wshihadeh/kubectl-counts.svg)](https://github.com/wshihadeh/kubectl-counts/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wshihadeh/kubectl-counts)](https://github.com/wshihadeh/kubectl-counts/blob/master/go.mod)

Filter Kubernetes resources by matching their names

# Requirements

- Kubernetes 1.17.0+
- Kubectl 1.17.0+
- Krew 0.4.0+


# Installation

Installation via [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)

    $ kubectl krew version # make sure you are running 0.4.0+
    $ kubectl krew install counts
    $ kubectl krew upgrade

Manual Installation

    $ export KUBECTL_GREP_VERSION=$(curl -s https://api.github.com/repos/wshihadeh/kubectl-counts/releases/latest | jq -r .tag_name)
    $ curl -L -O https://github.com/wshihadeh/kubectl-counts/releases/download/${KUBECTL_GREP_VERSION}/kubectl-counts-$(uname -s)-$(uname -m).tar.gz
    $ tar zxvf kubectl-counts-$(uname -s)-$(uname -m).tar.gz
    $ mv kubectl-counts /usr/local/bin
    $ chmod +x /usr/local/bin/kubectl-counts

# Examples

Get pod counts
```$ kubectl counts pod --by-ns```
```$ kubectl counts pod --by-node```
```$ kubectl counts pod --by-restarts```
```$ kubectl counts pod --by-status```

Get service counts
```$ kubectl counts services --by-ns```
```$ kubectl counts services --by-type```

Get jobs counts
```$ kubectl counts jobs --by-ns```
```$ kubectl counts jobs --by-container```
```$ kubectl counts jobs --by-image```

Get deployment counts
```$ kubectl counts deployments --by-ns```

Get ingress counts
```$ kubectl counts ingresses --by-ns```
```$ kubectl counts ingresses --by-class```
```$ kubectl counts ingresses --by-address```


# How to get developer build?

    $ go get -u github.com/wshihadeh/kubectl-counts

    $ cd ${GOPATH}/src/github.com/wshihadeh/kubectl-counts

    $ make all

# Reference

- [Kubectl Plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)

# License

[Apache-2.0](LICENSE)
