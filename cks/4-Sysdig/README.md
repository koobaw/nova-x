
准备：
```
【伯乐大典】CKS考试真题讲解12
https://www.bilibili.com/video/BV1ag411x7iw/?spm_id_from=333.337.search-card.all.click&vd_source=159ce9d2ab22bb84cfb4ceac4a8f857c
CKS第十二题：Sysdig & Falco
https://www.bilibili.com/video/BV12L411d7ZP/?spm_id_from=333.788&vd_source=159ce9d2ab22bb84cfb4ceac4a8f857c
```



做题：
```

```


代替：
```
sysdig -l 
susdig -h
$ docker ps | grep redis  这里应该可以grep到container.id
$ sysdig -l | grep time
$ sysdig -l | grep uid
$ sysdig -l | grep proc
kubectl get pod tomcat123 -o yaml   这里应该可以grep到container.name
可以看到里面的container name
sysdig -M 30 -p "*%evt.time,%user.uid,%proc.name" container.name=tomcat123 > /opt/2/report
```