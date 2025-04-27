//go:generate go run -tags mcpgen ./cmd/gen/main.go
package main

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/joho/godotenv"
	"github.com/ktr0731/go-mcp"
	"github.com/miyamo2/hotpepper-gourmet-mcp-server/internal/configs/di"
	"golang.org/x/exp/jsonrpc2"
	"log/slog"
	"os"
	"os/signal"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	handler := di.GetHandler()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	srv, err := jsonrpc2.Serve(mcp.NewStdioTransport(ctx, handler, nil))
	if err != nil {
		panic(err)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Wait()
	}()

	select {
	case <-ctx.Done():
		slog.Error("received interrupt signal")
		slog.Error("shutting down server")
		if err := ctx.Err(); err != nil {
			if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
				slog.Error("occurred error", slog.String("error", err.Error()))
			}
		}
	case err := <-errCh:
		if err != nil {
			slog.Error("occurred error", slog.String("error", err.Error()))
		}
	}
}
