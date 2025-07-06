# My Service Template

Template oficial para microserviços Go seguindo Clean Architecture e melhores práticas da empresa.

## Requisitos
- Go 1.21+
- Docker (para testes de integração)

## Configuração

1. **Criar arquivo .env** (baseado no config.yaml):
```sh
cp config/config.yaml .env
```

2. **Ajustar configurações** no arquivo `.env` conforme seu ambiente.

## Build
```sh
go build -o bin/myservice ./cmd/myservice
```

## Execução
```sh
go run ./cmd/myservice
```

## Testes
```sh
go test -v ./...
```

## Migrações
```sh
make migrate
```

## Estrutura
- `cmd/myservice/`: entrypoint
- `internal/`: domínio, casos de uso, interfaces, infraestrutura
- `config/`: configuração via Viper
- `pkg/`: utilitários
- `scripts/`: migrações, inicializações
- `test/`: testes unitários, integração, E2E

## Tecnologias
- **Echo** - Web framework
- **GORM** - ORM com multitenancy
- **Zap** - Logging estruturado
- **Viper** - Configuração (.env, .yaml, env vars, flags)
- **OpenTelemetry** - Tracing e métricas
- **Keycloak** - Autenticação JWT
- **MinIO** - Armazenamento de evidências
- **Outbox Pattern** - Eventos assíncronos
