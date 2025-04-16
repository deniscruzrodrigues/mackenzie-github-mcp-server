package main

import (
	"context"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	codeanalyzer "github.com/github/github-mcp-server/pkg/codenalyzer"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "code-analyzer-server",
		Short: "Code Analyzer MCP Server",
		Long:  `Code Analyzer MCP Server is a server that provides code analysis tools using the MCP protocol.`,
	}

	stdioCmd = &cobra.Command{
		Use:   "stdio",
		Short: "Start a stdio server",
		Long:  `Start a stdio server thats communicates via stdio streams using JSON-RPC messages`,
		Run: func(_ *cobra.Command, _ []string) {
			if err := runStdioServer(); err != nil {
				stdlog.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(stdioCmd)
}

func runStdioServer() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	codeAnalyzerServer := codeanalyzer.NewCodeAnalyzerServer()
	stdioServer := server.NewStdioServer(codeAnalyzerServer)

	errC := make(chan error, 1)
	go func() {
		in, out := io.Reader(os.Stdin), io.Writer(os.Stdout)
		errC <- stdioServer.Listen(ctx, in, out)
	}()

	_, _ = fmt.Fprintf(os.Stderr, "Code Analyzer MCP Server running on stdio...\n")

	select {
	case <-ctx.Done():
		fmt.Println("Received shutdown signal, shutting down...")
	case err := <-errC:
		if err != nil {
			return fmt.Errorf("error running stdio server: %w", err)
		}
	}

	return nil

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
