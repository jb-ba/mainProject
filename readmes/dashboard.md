# The K8s Dashboard

1. Make sure kubectl is installed locally
2. Copy `/etc/kubernetes/admin.conf` to local computer
3. Next create a sample user https://github.com/kubernetes/dashboard/wiki/Creating-sample-user
4. Create a service account:
``` 
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
 ```
5. Create ClusterRoleBinding
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kube-system
```
6. Create a bearer token <br />
`kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')`
7. Run on local machine: <br/>
`sudo kubectl --kubeconfig=./admin.conf proxy --address='0.0.0.0' --accept-hosts='^*$'`
8. Go to http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/overview?namespace=default
9. Paste bearer token from 6.


