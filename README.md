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

---

## 🚀 Como Usar Este Template

1. **Clone este repositório:**
   ```sh
git clone <url-do-template> nome-do-seu-servico
cd nome-do-seu-servico
```

2. **Renomeie o módulo Go:**
   - Edite o `go.mod` e altere `module my-service` para o nome do seu serviço (ex: `module user-service`).
   - Atualize os imports nos arquivos Go, se necessário.

3. **Configure o serviço:**
   - Ajuste o `.env` e/ou `config/config.yaml` para o contexto do seu microserviço.
   - Altere o nome do binário e diretórios se desejar.

4. **Implemente seu domínio:**
   - Crie entidades em `internal/domain`.
   - Implemente casos de uso em `internal/usecase`.
   - Adicione handlers HTTP em `internal/interface/http`.
   - Configure integrações externas em `internal/infra`.

5. **Siga os padrões:**
   - Use sempre Viper para configuração.
   - Siga o style guide e mantenha o projeto formatado com `gofmt`.
   - Escreva testes unitários e de integração.

6. **Build e execução:**
   ```sh
make build
make run
```

7. **Deploy:**
   - O binário estará em `bin/`.
   - Use Docker/Kubernetes conforme padrão da empresa.

---

## 💡 Exemplos de Projetos Possíveis
- **user-service**: Gerenciamento de usuários, autenticação, perfis.
- **order-service**: Processamento de pedidos, integração com pagamentos.
- **evidence-service**: Armazenamento e consulta de evidências (MinIO).
- **audit-service**: Log imutável de operações (Rekor).
- **notification-service**: Envio de notificações assíncronas (Outbox).

---

## ⚠️ Pontos Importantes para Desenvolvedores
- **Nunca use valores hardcoded**: Sempre utilize Viper/configuração.
- **Inclua os campos obrigatórios nas entidades** (ver guia).
- **Implemente middlewares de observabilidade** (tracing, métricas, logs).
- **Proteja rotas sensíveis com autenticação JWT/Keycloak**.
- **Use o padrão Outbox para eventos assíncronos**.
- **Documente endpoints e integrações**.
- **Mantenha o projeto atualizado com as melhores práticas do Go**.
- **Faça code review e utilize os linters do projeto**.

---

## 📚 Referências
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
