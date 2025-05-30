Steps to execute the project

1º Run the Db (execute the following command on the root dir)

> docker-compose up 

2º Start the grafana pods
You also need to start minikube

> minikube start
> terraform apply

3º Redirect grafana ports and monitor pods 

> kubectl port-forward svc/tempo 4318 -n monitoring
>
> kubectl port-forward svc/grafana 3000:80 -n monitoring
> 
4º Add the data sources
First add the prometheus datasource
> http://prometheus-server.monitoring.svc.cluster.local

Then add the tempo data source (for tracing)
> http://tempo.monitoring:3000




-----

Para config do Kubernets (Localmente com Kind)

Kind create cluster (cria cluster local)

kubectl config use-context kind-kind (Usa contexto default do cluster)

Agora precisa baixar os crd's

helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
kubectl create namespace monitoring
helm install prometheus-stack prometheus-community/kube-prometheus-stack -n monitoring


