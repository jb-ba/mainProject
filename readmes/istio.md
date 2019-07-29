curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.2 sh -


## Kiali
For forwarding use
```
kubectl -n istio-system port-forward --address 0.0.0.0 $(kubectl -n istio-system get pod -l app=kiali -o jsonpath='{.items[0].metadata.name}') 20001:20001
```