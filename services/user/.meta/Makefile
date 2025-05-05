DIR=$(shell pwd)

.PHONY: restart-user
restart-user: ## Перезапуск пода с юзерами
	@kubectl delete pod -l app=user -n user

.PHONY: build-user
build-user: ## Сборка сервиса user
	@docker build -t user:latest -f $(DIR)/services/user/Dockerfile $(DIR)

.PHONY: rebuild-user
rebuild-user: build-user restart-user ## Пересборка образа юзеров и перезапуск пода

.PHONY: setup-k8s-user
setup-k8s-user: ## Применить к8 манифесты
	@kubectl apply -f $(DIR)/services/user/k8s/namespace.yaml
	@kubectl apply -f $(DIR)/services/user/k8s/deployment.yaml
	@kubectl apply -f $(DIR)/services/user/k8s/service.yaml
