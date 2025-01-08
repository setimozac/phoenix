#!/bin/bash

exec > ~/instalation_log.log 2>&1
cd ~/
export CURRENT_USER=$(whoami)
sudo yum -y install docker
sudo usermod -a -G docker $CURRENT_USER
newgrp docker
sudo systemctl enable docker.service
sudo systemctl start docker.service
echo "Docker installation is done."
echo "Installing Go"
sudo yum -y install go
echo "install kubectl"
curl -o kubectl https://s3.us-west-2.amazonaws.com/amazon-eks/1.30.0/2024-05-12/bin/linux/amd64/kubectl
chmod +x kubectl
sudo mv kubectl /usr/local/bin
echo "install kind"
go install sign.k8s.io/kind@v0.26.0
sudo cp ./go/bin/kind /usr/local/bin
echo "create cluster"
kind create cluster --name phoenix-dev
kubectl cluster-info --context kind-operator-dev


