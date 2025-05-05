DIR=$(shell pwd)

.PHONY: build-auth
build-auth: ## Сборка сервиса auth
	@docker build -t auth:latest -f $(DIR)/services/auth/Dockerfile $(DIR)
	

.PHONY: restart-auth
restart-auth: ## Перезапуск пода авторизации
	@kubectl delete pod -l app=auth -n auth

.PHONY: rebuild-auth
rebuild-auth: build-auth restart-auth ## Пересборка образа авторизации и перезапуск пода

.PHONY: setup-k8s-auth
setup-k8s-auth: ## Применить к8 манифесты
	@kubectl apply -f $(DIR)/services/auth/k8s/namespace.yaml
	@kubectl apply -f $(DIR)/services/auth/k8s/deployment.yaml
	@kubectl apply -f $(DIR)/services/auth/k8s/service.yaml
