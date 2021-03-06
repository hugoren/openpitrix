#!/bin/bash
git clone https://github.com/openpitrix/openpitrix.git /opt/openpitrix
# Back to the root of the project
cd $(dirname $0)
cd ../..

kubectl create secret generic mysql-pass --from-file=./devops/kubernetes/password.txt -n default
kubectl apply -f ./devops/kubernetes/db
kubectl apply -f ./devops/kubernetes/etcd
kubectl apply -f ./devops/kubernetes/openpitrix
