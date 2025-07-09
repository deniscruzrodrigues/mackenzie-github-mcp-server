# Guia de Build - GitHub MCP Server

Este documento explica como compilar e desenvolver o GitHub MCP Server.

## Pré-requisitos

- Go 1.23.7 ou superior
- Git
- Make (opcional, mas recomendado)

## Instalação do Go

### Linux/macOS
```bash
# Baixar e instalar Go
wget https://go.dev/dl/go1.23.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Adicionar ao .bashrc/.zshrc para persistir
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```

### Windows
Baixe o instalador do [site oficial do Go](https://golang.org/dl/).

## Build Rápido

### Usando Make (Recomendado)
```bash
# Compilar todos os executáveis
make build

# Compilar apenas o servidor principal
make github-mcp-server

# Ver todos os comandos disponíveis
make help
```

### Usando Go diretamente
```bash
# Instalar dependências
go mod download

# Compilar servidor principal
go build -o bin/github-mcp-server ./cmd/github-mcp-server

# Compilar analisador de código
go build -o bin/code-analyzer-server ./cmd/code-analyzer-server

# Compilar cliente curl MCP
go build -o bin/mcpcurl ./cmd/mcpcurl
```

## Desenvolvimento

### Configurar ambiente de desenvolvimento
```bash
make dev-setup
```

### Executar testes
```bash
# Testes básicos
make test

# Testes com cobertura
make test-coverage

# Testes detalhados
make test-verbose
```

### Verificar qualidade do código
```bash
# Executar linter
make lint

# Formatar código
make fmt

# Executar lint + testes
make check
```

### Executar o servidor
```bash
# Compilar e mostrar ajuda
make run-github-mcp-server

# Executar diretamente
./bin/github-mcp-server stdio --help
```

## Build para Produção

### Build otimizado
```bash
make release-build
```

Isso criará executáveis otimizados para:
- Linux (amd64)
- macOS (amd64) 
- Windows (amd64)

### Docker
```bash
make docker-build
```

## Estrutura dos Executáveis

- **github-mcp-server**: Servidor principal MCP para GitHub
- **code-analyzer-server**: Servidor de análise de código
- **mcpcurl**: Cliente de linha de comando para MCP

## Limpeza

```bash
# Remover arquivos de build
make clean
```

## Solução de Problemas

### Erro "go: command not found"
Certifique-se de que o Go está instalado e no PATH.

### Erro de dependências
```bash
go mod tidy
go mod download
```

### Problemas de permissão
```bash
chmod +x bin/*
```

## Contribuindo

1. Fork o repositório
2. Crie uma branch para sua feature
3. Execute `make check` antes de commitar
4. Faça commit das mudanças
5. Abra um Pull Request

## Comandos Make Disponíveis

| Comando | Descrição |
|---------|-----------|
| `make build` | Compila todos os executáveis |
| `make test` | Executa testes |
| `make lint` | Executa linter |
| `make clean` | Remove arquivos de build |
| `make dev-setup` | Configura ambiente de desenvolvimento |
| `make help` | Mostra todos os comandos |

