DIR=$(shell pwd)

.PHONY: build-message
build-message: ## Сборка сервиса message
	@docker build -t message:latest -f $(DIR)/services/message/Dockerfile $(DIR)

.PHONY: restart-message
restart-message: ## Перезапуск пода с юзерами
	@kubectl delete pod -l app=message -n message

.PHONY: rebuild-message
rebuild-message: build-message restart-message ## Пересборка образа юзеров и перезапуск пода

.PHONY: setup-k8s-message
setup-k8s-message: ## Применить к8 манифесты
	@kubectl apply -f $(DIR)/services/message/k8s/namespace.yaml
	@kubectl apply -f $(DIR)/services/message/k8s/deployment.yaml
	@kubectl apply -f $(DIR)/services/message/k8s/service.yaml
