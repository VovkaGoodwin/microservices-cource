Сборка и запуск приложеня тестировались в следуюущей среде 
    - Docker version 27.4.0, build bde2b89
    - Kubernetes Client Version: v1.30.5
      Kustomize Version: v5.0.4-0.20230601165947-6ce0bf390ce3
      Server Version: v1.30.5

Для запуска приложения

```shell 
make up 
```

Для получения справки по доступным командам 
```shell
make help
```

Если есть проблемы доступом к кластеру то вероятнее всего необходимо установить ingress-nginx
```shell
make setup-nginx
```

Доступные маршруты
[Auth service](http://localhost/auth/healthcheck)
[User service](http://localhost/user/ping)
[Message service](http://localhost/message/ping)
