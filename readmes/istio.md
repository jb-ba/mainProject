# Istio

## install via
curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.2 sh -


## Kiali
For forwarding use
```
kubectl -n istio-system port-forward --address 0.0.0.0 $(kubectl -n istio-system get pod -l app=kiali -o jsonpath='{.items[0].metadata.name}') 20001:20001
```

## ARM
The envoy proxy is not available on ARM yet. As a R_Pi is used on the edge this makes Istio impossible to use.<br />
There is work on porting envoy for ARM so maybe in the future.