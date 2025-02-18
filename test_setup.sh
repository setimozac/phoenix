#!/bin/bash

# exec > ~/instalation_log.log 2>&1
cd ~/
export CURRENT_USER=$(whoami)
sudo yum -y install docker
sudo usermod -a -G docker $CURRENT_USER
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
go install sigs.k8s.io/kind@v0.26.0
sudo cp ./go/bin/kind /usr/local/bin
echo "create cluster"


newgrp docker <<EOF
kind create cluster --name phoenix-dev
kubectl cluster-info --context kind-phoenix-dev

curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

newgrp docker
EOF



#docker build -t phoenix-operator:latest .
#kind load docker-image phoenix-operator:latest --name phoenix-dev

# sudo yum -y install git
# docker system prune -af
# docker volume prune -f


