kubectl create configmap --dry-run go-rest-api-crt-config --from-file=./certs/go-rest-api.crt --output yaml | tee ./k8s/api/crt-configmap.tmp
kubectl patch --local -f ./k8s/api/crt-configmap.tmp --type=json -p='[{"op": "remove", "path": "/metadata/creationTimestamp"}]' -o yaml > ./k8s/api/crt-configmap.yaml
rm -rf ./k8s/api/crt-configmap.tmp

kubectl create configmap --dry-run go-rest-api-key-config --from-file=./certs/go-rest-api.key --output yaml | tee ./k8s/api/key-configmap.tmp
kubectl patch --local -f ./k8s/api/key-configmap.tmp --type=json -p='[{"op": "remove", "path": "/metadata/creationTimestamp"}]' -o yaml > ./k8s/api/key-configmap.yaml
rm -rf ./k8s/api/key-configmap.tmp