DIR=$(shell pwd)

.PHONY: build-gateway
build-gateway:
	@docker build -t gateway:latest -f $(DIR)/services/gateway/Dockerfile $(DIR)

.PHONY: restart-gateway
restart-gateway: ## Перезаппуск пода гейтвея
	@kubectl delete pod -l app=gateway-svc

.PHONY: rebuild-gateway
rebuild-gateway: build-gateway restart-gateway ## Пересборка образа гейтвея и перезапуск пода

.PHONY: setup-k8s-gateway
setup-k8s-gateway: ## Применить к8 манифесты
	@kubectl apply -f $(DIR)/services/gateway/k8s/deployment.yaml
	@kubectl apply -f $(DIR)/services/gateway/k8s/service.yaml
