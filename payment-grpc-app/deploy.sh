set -ex

echo "Matando port-forwards antigos..."
pkill -f "kubectl port-forward" || true
sleep 2

## Subir o banco
docker-compose up -d

# Subir grafana, prometheus e tempo
cd terraform/
terraform init
terraform apply
cd ..


# garantir que o servico ja nao foi criado antes
kubectl delete deployment payment-service --ignore-not-found=true

kubectl apply -R -f api--payment-service/k8s
kubectl rollout status deployment/payment-service --timeout=90s

#todos os port forward
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