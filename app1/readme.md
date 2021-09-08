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
kubectl rollout restart deployment testgo-deployment
kubectl get pod -o wide
kubectl describe svc testgo-svc
for i in {1..50}; do curl http://k8s3:port/; done
```

## service
``kubectl expose deployment testgo-deployment --type=NodePort --name=testgo-svc``  
[dns-pod-service](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/)

## [pull-image-private-registry](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/) 
```
docker login 192.168.1.50:5000
enter foo bar
kubectl create secret generic regcred \
    --from-file=.dockerconfigjson=./config.json \
    --type=kubernetes.io/dockerconfigjson

kubectl get secret
```