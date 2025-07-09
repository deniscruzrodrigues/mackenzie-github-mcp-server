# Makefile para GitHub MCP Server

.PHONY: all build clean test lint help install-deps

# Variáveis
BINARY_DIR := bin
GITHUB_MCP_SERVER := $(BINARY_DIR)/github-mcp-server
CODE_ANALYZER_SERVER := $(BINARY_DIR)/code-analyzer-server
MCPCURL := $(BINARY_DIR)/mcpcurl

# Comandos Go
GO := go
GOFMT := gofmt
GOLINT := golangci-lint

# Targets principais
all: build

help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

install-deps: ## Instala dependências do projeto
	$(GO) mod download
	$(GO) mod tidy

build: install-deps ## Compila todos os executáveis
	@echo "Compilando executáveis..."
	@mkdir -p $(BINARY_DIR)
	$(GO) build -o $(GITHUB_MCP_SERVER) ./cmd/github-mcp-server
	$(GO) build -o $(CODE_ANALYZER_SERVER) ./cmd/code-analyzer-server
	$(GO) build -o $(MCPCURL) ./cmd/mcpcurl
	@echo "Build concluído! Executáveis em $(BINARY_DIR)/"

github-mcp-server: install-deps ## Compila apenas o github-mcp-server
	@mkdir -p $(BINARY_DIR)
	$(GO) build -o $(GITHUB_MCP_SERVER) ./cmd/github-mcp-server

code-analyzer-server: install-deps ## Compila apenas o code-analyzer-server
	@mkdir -p $(BINARY_DIR)
	$(GO) build -o $(CODE_ANALYZER_SERVER) ./cmd/code-analyzer-server

mcpcurl: install-deps ## Compila apenas o mcpcurl
	@mkdir -p $(BINARY_DIR)
	$(GO) build -o $(MCPCURL) ./cmd/mcpcurl

test: ## Executa todos os testes
	$(GO) test ./...

test-verbose: ## Executa testes com saída detalhada
	$(GO) test -v ./...

test-coverage: ## Executa testes com cobertura
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

lint: ## Executa linter no código
	$(GOLINT) run

fmt: ## Formata o código
	$(GOFMT) -s -w .

clean: ## Remove arquivos de build
	@echo "Limpando arquivos de build..."
	rm -rf $(BINARY_DIR)
	rm -f coverage.out coverage.html
	rm -f github-mcp-server code-analyzer-server mcpcurl

run-github-mcp-server: github-mcp-server ## Compila e executa o github-mcp-server
	./$(GITHUB_MCP_SERVER) --help

docker-build: ## Constrói imagem Docker
	docker build -t github-mcp-server .

# Targets de desenvolvimento
dev-setup: install-deps ## Configura ambiente de desenvolvimento
	@echo "Configurando ambiente de desenvolvimento..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "Instalando golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2; \
	fi
	@echo "Ambiente configurado!"

check: lint test ## Executa lint e testes

release-build: ## Build para release (otimizado)
	@mkdir -p $(BINARY_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags="-w -s" -o $(GITHUB_MCP_SERVER)-linux-amd64 ./cmd/github-mcp-server
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -ldflags="-w -s" -o $(GITHUB_MCP_SERVER)-darwin-amd64 ./cmd/github-mcp-server
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -ldflags="-w -s" -o $(GITHUB_MCP_SERVER)-windows-amd64.exe ./cmd/github-mcp-server

