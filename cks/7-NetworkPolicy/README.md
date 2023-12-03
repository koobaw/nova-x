
准备：
```
```



做题：
```
$ vim /cks/15/p1.yaml apiVersion: networking.k8s.io/v1 kind: NetworkPolicy
metadata:
name: "denynetwork"
namespace: "development" spec:
podSelector: {}
policyTypes:
- Ingress #拒绝所有入栈流量
$ kubectl apply -f /cks/15/p1.yaml
```
