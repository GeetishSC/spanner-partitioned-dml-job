## Purpose
The purpose of this job is to solely delete records on ChatMessage table using PartitionedDML queries


## Pre-Requisite
golang >= 1.19

## Install Golang 1.19 on a VM
```shell
wget  https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
sudo tar -xvf go1.20.2.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
echo "export GOPATH=$HOME/go" >> ~/.bashrc 
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc 
```

## Improvements
Make it more generic. Its pretty static right now
