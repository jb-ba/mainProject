# Configure Node


## Labels

### Check
To show all labels of a node
```
kubectl get nodes --show-labels
```

### Add
To add a label use:
```
kubectl label nodes <your-node-name> disktype=ssd
```
E.g. kubectl label nodes piwifi pi=hello

