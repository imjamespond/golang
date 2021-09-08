## docker run
```
docker run --rm -p 8080:8080 -it testgo:v0.1 foo -bar
```

## docker export
```
docker save -o testgo.tar testgo:v0.1 
```

## import
``find `pwd`| grep ".tar"| xargs -n1 ctr  -n k8s.io i import``

## run
```
kubectl apply -f dp.yaml
kubectl delete deployment testgo-deployment
kubectl get pod
kubectl get service
```

## service
``kubectl expose deployment testgo-deployment --type=NodePort --name=testgo-svc``