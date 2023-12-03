
准备：
```

```



做题：
```
$ kubectl create sa backend-sa -n qa --dry-run=client -o yaml > sa.yaml $ vim sa.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
name: backend-sa
namespace: qa automountServiceAccountToken: false
$ kubectl apply -f sa.yaml
#禁止访问任何凭据
#考试会自动有 backend 的 pod，不需要自己创建 $ vim /cks/9/pod9.yaml
apiVersion: v1
kind: Pod
metadata: labels:
run: backend name: backend namespace: qa
spec:
serviceAccountName: backend-sa containers:
- image: nginx:1.9
name: backend
resources: {} dnsPolicy: ClusterFirst restartPolicy: Always
$kubectl apply -f /cks/9/pod9.yaml
可以看看 qa 名称空间有哪些 sa，把除了 backend-sa 的 sa 都删除 kubectl delete sa sa 名字 -n qa

```

代替：
```

```