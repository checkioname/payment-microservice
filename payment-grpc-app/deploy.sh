set -ex


# Atualizando as imagens
cd api--order-service
sudo docker build -t lucas10sh/order-service:latest .
sudo docker push lucas10sh/order-service:latest
cd ..

cd api--payment-service
sudo docker build -t lucas10sh/payment-service:latest .
sudo docker push lucas10sh/payment-service:latest
cd ..


cd api--inventory-service
sudo docker build -t lucas10sh/inventory-service:latest .
sudo docker push lucas10sh/inventory-service:latest
cd ..

cd api--invoice-service
sudo docker build -t lucas10sh/invoice-service:latest .
sudo docker push lucas10sh/invoice-service:latest
cd ..

cd api--shipping-service
sudo docker build -t lucas10sh/shipping-service:latest .
sudo docker push lucas10sh/shipping-service:latest
cd ..



echo "Matando port-forwards antigos..."
pkill -f "kubectl port-forward" || true
sleep 2

## Subir o banco
docker-compose up -d


# instalar crd's
echo "Instalando CRDs do Prometheus Operator (idempotente)..."
#kubectl apply --server-side -f https://github.com/prometheus-operator/prometheus-operator/releases/latest/download/bundle.yaml

kubectl apply -R -f monitoring/prometheus

# Subir grafana, prometheus e tempo
cd terraform/
terraform init
terraform apply -auto-approve
cd ..

# garantir que o servico ja nao foi criado antes
kubectl delete deployment order-service --ignore-not-found=true
kubectl delete deployment payment-service --ignore-not-found=true
kubectl delete deployment inventory-service --ignore-not-found=true
kubectl delete deployment invoice-service --ignore-not-found=true
kubectl delete deployment shipping-service --ignore-not-found=true


kubectl apply -R -f api--order-service/k8s
kubectl rollout status deployment/order-service --timeout=90s

kubectl apply -R -f api--payment-service/k8s
kubectl rollout status deployment/payment-service --timeout=90s

kubectl apply -R -f api--inventory-service/k8s
kubectl rollout status deployment/inventory-service --timeout=90s

kubectl apply -R -f api--invoice-service/k8s
kubectl rollout status deployment/invoice-service --timeout=90s

kubectl apply -R -f api--shipping-service/k8s
kubectl rollout status deployment/shipping-service --timeout=90s

#todos os port forward
kubectl port-forward svc/order-service 8007:8007 &

kubectl port-forward svc/payment-service 8082:8082 &
kubectl port-forward svc/payment-service 8008:8008 &

kubectl port-forward svc/tempo 4318 -n monitoring &
kubectl port-forward svc/grafana 3000:80 -n monitoring &
kubectl port-forward svc/prometheus-stack-kube-prom-prometheus 9090:9090 -n monitoring &

echo
echo "---------"
echo "Back-office:"
echo "  - Grafana:    http://localhost:3000"
echo "  - Prometheus: http://localhost:9090"
echo "  - Tempo:      http://localhost:4318"
echo "  - API Métricas:  http://localhost:8082/metrics"
echo "---------"