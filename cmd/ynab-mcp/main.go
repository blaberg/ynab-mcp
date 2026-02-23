package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blaberg/ynab-mcp/internal/tools"
	"github.com/blaberg/ynab-mcp/internal/ynab"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"ynab",
		"0.1.0",
		server.WithToolCapabilities(false),
	)

	tools.RegisterTools(s)

	transport := os.Getenv("TRANSPORT")
	if transport == "" {
		transport = "stdio"
	}

	switch transport {
	case "stdio":
		serveStdio(s)
	case "http":
		serveHTTP(s)
	default:
		log.Fatalf("unknown transport %q: must be \"stdio\" or \"http\"", transport)
	}
}

func serveStdio(s *server.MCPServer) {
	token := os.Getenv("YNAB_API_TOKEN")
	if token == "" {
		log.Fatal("YNAB_API_TOKEN environment variable is required for stdio transport")
	}

	client := ynab.NewClient(token)

	if err := server.ServeStdio(s, server.WithStdioContextFunc(func(ctx context.Context) context.Context {
		return ynab.NewContext(ctx, client)
	})); err != nil {
		log.Fatalf("stdio server error: %v", err)
	}
}

func serveHTTP(s *server.MCPServer) {
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithHTTPContextFunc(func(ctx context.Context, r *http.Request) context.Context {
			token := r.Header.Get("X-YNAB-Token")
			if token == "" {
				return ctx
			}
			client := ynab.NewClient(token)
			return ynab.NewContext(ctx, client)
		}),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	fmt.Fprintf(os.Stderr, "YNAB MCP server listening on %s\n", addr)
	if err := httpServer.Start(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
