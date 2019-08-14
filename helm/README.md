# Helm

## create new chart
Use
```
helm create <mychart>
```

## Check a helm chart
Use
`helm lint <dirOfChart>`

## install
Inside the main helm directory of a project execute
```
helm install . --namespace picoap-namespace
```

## upgrade
Inside the main helm directory of a project execute
```
helm upgrade <currentHelmChartName> .
```