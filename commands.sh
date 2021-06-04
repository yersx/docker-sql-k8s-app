docker build --no-cache -t yersultan/simplegolangk8sdocker:v1 .

docker push yersultan/simplegolangk8sdocker:v1

kubectl delete deployment simplegolangk8sdocker

kubectl apply -f devops/deployment.yaml

kubectl apply -f devops/service.yaml