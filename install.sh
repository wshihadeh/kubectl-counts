#!/bin/bash


export KUBECTL_COUNTS_VERSION=$(curl -s https://api.github.com/repos/wshihadeh/kubectl-counts/releases/latest | jq -r .tag_name)
curl -fsL -O https://github.com/wshihadeh/kubectl-counts/releases/download/${KUBECTL_COUNTS_VERSION}/kubectl-counts-$(uname -s)-$(uname -m).tar.gz
tar zxvf kubectl-counts-$(uname -s)-$(uname -m).tar.gz
mv kubectl-counts /usr/local/bin
chmod +x /usr/local/bin/kubectl-counts
