DIRS := auth message user
FILES := namespace deployment service

.DEFAYLT_GOAL: help

.PHONY: help
help:
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: setup-k8s-services
setup-k8s-services: ## Применить манифесты всех сервисов
	for dir in $(DIRS); do \
  		for file in $(FILES); do \
			kubectl apply -f ./services/$$dir/k8s/$$file.yaml; \
		done; \
  	done;

.PHONY: setup-k8s-root
setup-k8s-root: ## Применить корневые манифесты приложения
	kubectl apply -f ./k8s/auth-svc-bridge.yaml
	kubectl apply -f ./k8s/user-svc-bridge.yaml
	kubectl apply -f ./k8s/message-svc-bridge.yaml
	kubectl apply -f ./k8s/ingress.yaml

.PHONY: setup-k8s
setup-k8s: setup-k8s-services setup-k8s-root ## Применить все доступные манифесты

.PHONY: up
up: build-images setup-k8s ## Запустить приложение

.PHONY: setup-nginx
setup-nginx: ## Установка ingress-nginx если ранее не была настроена
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

.PHONY: build-user
build-user: ## Сборка сервиса user
	docker build -t user:latest -f ./services/user/Dockerfile ./services/user

.PHONY: build-message
build-message: ## Сборка сервиса message
	docker build -t message:latest -f ./services/message/Dockerfile ./services/message

.PHONY: build-auth
build-auth: ## Сборка сервиса auth
	docker build -t auth:latest -f ./services/auth/Dockerfile ./services/auth

.PHONY: build-images ## Собрать образы всех приложений
build-images: build-auth build-user build-message