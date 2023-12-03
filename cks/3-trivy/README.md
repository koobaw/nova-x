
准备：
```
docker run -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/Library/Caches:/root/.cache/ aquasec/trivy:0.20.2 registry.aliyuncs.com/google_containers/coredns:1.7.0 | grep -i "high"
```



做题：
```
$ ssh root@kssc00401-master
$ kubectl get po -n yavin -oyaml | grep ‘image’
$ trivy image --skip-update ‘刚才搜到的镜像’ | egrep -i "High|Critical" 
$ 把检测出来漏洞的镜像对应的 pod 删除
#--skip-update #跳过镜像更新，考试可以直接跳过。
```


代替：
```
kubectl get nodes
ssh root@master
kubectl get pods -n kube-system -oyaml | grep "image"
trivy image --skip-update 'registry.aliyuncs.com/google_containers/coredns:1.7.0' | egrep -i "High|Critical"
kubectl get pod <pod-name> -o yaml
kubectl delete pods ""
```