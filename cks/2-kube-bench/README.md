
准备：
```
docker run --rm -v `pwd`:/host aquasec/kube-bench:latest install
#仅对 master 节点检查
./kube-bench master 
./kube-bench node
```



做题：
```
$ ssh root@kscs00201-master
$ vim /etc/kubernetes/manifests/kube-apiserver.yaml #- --authorization-mode=AlwaysAllow
- --authorization-mode=Node,RBAC
#- --insecure-bind-address=0.0.0.0
- --insecure-port=0
$ vim /var/lib/kubelet/config.yaml anonymous:
enabled: false authorization:
mode: Webhook
#master 和 node 节点都需要改配置
$ vim /etc/kubernetes/manifests/etcd.yaml - --client-cert-auth=true
$ systemctl daemon-reload $ systemctl restart kubelet
```
