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

.PHONY: setup-gateway
setup-gateway:
	kubectl apply -f ./services/gateway/k8s/deployment.yaml
	kubectl apply -f ./services/gateway/k8s/service.yaml

.PHONY: setup-k8s-root
setup-k8s-root: ## Применить корневые манифесты приложения
	kubectl apply -f ./k8s/ingress.yaml

.PHONY: setup-k8s
setup-k8s: setup-k8s-services setup-k8s-root setup-gateway## Применить все доступные манифесты

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

.PHONY: build-gateway
build-gateway:
	docker build -t gateway:latest -f ./services/gateway/Dockerfile ./services/gateway


.PHONY: build-images ## Собрать образы всех приложений
build-images: build-auth build-user build-message build-gateway

.PHONY: proto
proto:
	@protoc --proto_path=proto/api/auth --go_out=./services/auth/proto --go_opt=paths=source_relative --go-grpc_out=./services/auth/proto --go-grpc_opt=paths=source_relative auth.proto
	@protoc --proto_path=proto/api/auth --go_out=./services/gateway/proto/auth --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/auth --go-grpc_opt=paths=source_relative auth.proto
	@protoc --proto_path=proto/api/user --go_out=./services/user/proto --go_opt=paths=source_relative --go-grpc_out=./services/user/proto --go-grpc_opt=paths=source_relative user.proto
	@protoc --proto_path=proto/api/user --go_out=./services/gateway/proto/user --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/user --go-grpc_opt=paths=source_relative user.proto
	@protoc --proto_path=proto/api/message --go_out=./services/message/proto --go_opt=paths=source_relative --go-grpc_out=./services/message/proto --go-grpc_opt=paths=source_relative message.proto
	@protoc --proto_path=proto/api/message --go_out=./services/gateway/proto/message --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/message --go-grpc_opt=paths=source_relative message.proto

