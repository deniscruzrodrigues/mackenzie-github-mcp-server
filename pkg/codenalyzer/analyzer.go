package codeanalyzer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewCodeAnalyzerServer cria um novo server MCP com ferramentas de análise de código
func NewCodeAnalyzerServer() *server.MCPServer {
	s := server.NewMCPServer(
		"code-analyzer-mcp-server",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	//Adiciona as ferramentas de análise de código
	s.AddTool(analyzeCode())
	s.AddTool(suggestImprovements())
	s.AddTool(explainCode())

	return s
}

// analyzeCode cria uma ferramenta para analisar código fonte
func analyzeCode() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("analyze_code",
			mcp.WithDescription("Analise código fonte e identifica problemas"),
			mcp.WithString("file_path", mcp.Required(), mcp.Description("Caminho do arquivo a ser analisado")),
			mcp.WithString("language", mcp.Required(), mcp.Description("Linguagem do arquivo a ser analisado")),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			//Implementação da análise de código
			filePath, _ := request.Params.Arguments["file_path"].(string)
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Erro ao ler arquivo: %v", err)), nil
			}

			language, _ := request.Params.Arguments["language"].(string)
			if language == "" {
				ext := filepath.Ext(filePath)
				language = detectLanguage(ext)
			}

			// Simulação de análise
			analysis := map[string]interface{}{
				"language":      language,
				"lines_of_code": len(strings.Split(string(content), "\n")),
				"issues": []map[string]interface{}{
					{"type": "error", "description": "Variável não inicializada", "line": 5},
					{"type": "warning", "description": "Uso de variável não utilizada", "line": 10},
				},
				"complexity_score": 7.5,
			}

			result, _ := json.Marshal(analysis)
			return mcp.NewToolResultText(string(result)), nil
		}
}

// Implementações simplificadas das outras ferramentas
func suggestImprovements() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	//Implementação similar ao analyzeCode, mas focada em sugestões
	//Código omitido por brevidade
	return mcp.NewTool("suggest_improvements",
			mcp.WithDescription("Sugere melhorias no código fonte"),
			mcp.WithString("file_path", mcp.Required(), mcp.Description("Caminho do arquivo a ser analisado")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return mcp.NewToolResultText("Sugestões de melhorias: ..."), nil
		}
}

func explainCode() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	//Implementação para explicar o código
	//Código omitido por brevidade
	return mcp.NewTool("explain_code",
			mcp.WithDescription("Explica o código fonte"),
			mcp.WithString("code", mcp.Required(), mcp.Description("Código a ser explicado")),
			mcp.WithString("language", mcp.Required(), mcp.Description("Linguagem do código")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return mcp.NewToolResultText("Explicação do código: ..."), nil
		}
}

// Função auxiliar para detectar linguagem com base na extensão do arquivo
func detectLanguage(extension string) string {
	extension = strings.ToLower(extension)
	switch extension {
	case ".go":
		return "Go"
	case ".py":
		return "Python"
	case ".js":
		return "JavaScript"
	case ".java":
		return "Java"
	case ".cpp":
		return "C++"
	case ".c":
		return "C"
	case ".html":
		return "HTML"
	default:
		return "Unknown"
	}
}
