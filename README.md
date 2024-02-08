# codegen_go_talk

Реализует простейший веб-сервер для демонстрации работы кодогенерации из OpenAPI и GRPC спецификаций.

## REST

Установка:

```bash
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Генерация:

```bash
oapi-codegen -package v1 \
    -generate server,types \
    api/rest/coverage/v1/service.yaml > pkg/api/rest/coverage/v1/service.gen.go
```

## GRPC

Установка:

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Генерация:

```bash
protoc --go_out=pkg --go_opt=paths=source_relative \
    --go-grpc_out=pkg --go-grpc_opt=paths=source_relative \
    api/grpc/coverage/v1/service.proto
```