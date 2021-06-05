docker build --no-cache -t yersultan/simplefullstackapp:v1 .

docker push yersultan/simpleFullstackApp:v1

kubectl delete deployment simpleFullstackApp

kubectl apply -f manifest/deployment.yaml

kubectl apply -f manifest/service.yaml