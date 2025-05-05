include services/user/.meta/Makefile
include services/message/.meta/Makefile
include services/gateway/.meta/Makefile
include services/auth/.meta/Makefile

.DEFAULT_GOAL: help

.PHONY: help
help:
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) |  sed -E 's/^[^:]+://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup-k8s-services
setup-k8s-services: setup-k8s-auth setup-k8s-gateway setup-k8s-message setup-k8s-user ## Применить манифесты всех сервисов

.PHONY: setup-k8s-root
setup-k8s-root: ## Применить корневые манифесты приложения
	kubectl apply -f ./k8s/ingress.yaml

.PHONY: setup-k8s
setup-k8s: setup-k8s-services setup-k8s-root ## Применить все доступные манифесты

.PHONY: up
up: build-images setup-k8s ## Запустить приложение

.PHONY: setup-nginx
setup-nginx: ## Установка ingress-nginx если ранее не была настроена
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

.PHONY: build-images
build-images: build-auth build-user build-message build-gateway ## Собрать образы всех приложений

.PHONY: proto
proto: ## Пересборка .proto файлов
	@protoc --proto_path=proto/api/auth --go_out=./services/auth/proto --go_opt=paths=source_relative --go-grpc_out=./services/auth/proto --go-grpc_opt=paths=source_relative auth.proto
	@protoc --proto_path=proto/api/auth --go_out=./services/gateway/proto/auth --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/auth --go-grpc_opt=paths=source_relative auth.proto
	@protoc --proto_path=proto/api/user --go_out=./services/user/proto --go_opt=paths=source_relative --go-grpc_out=./services/user/proto --go-grpc_opt=paths=source_relative user.proto
	@protoc --proto_path=proto/api/user --go_out=./services/gateway/proto/user --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/user --go-grpc_opt=paths=source_relative user.proto
	@protoc --proto_path=proto/api/message --go_out=./services/message/proto --go_opt=paths=source_relative --go-grpc_out=./services/message/proto --go-grpc_opt=paths=source_relative message.proto
	@protoc --proto_path=proto/api/message --go_out=./services/gateway/proto/message --go_opt=paths=source_relative --go-grpc_out=./services/gateway/proto/message --go-grpc_opt=paths=source_relative message.proto

