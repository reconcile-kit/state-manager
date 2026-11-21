package redisdiag

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"net"
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestDialFailurePreservesErrorAndReportsStage(t *testing.T) {
	var out bytes.Buffer
	h := Hook{Logger: slog.New(slog.NewJSONHandler(&out, nil))}
	dial := h.DialHook(func(context.Context, string, string) (net.Conn, error) {
		return nil, context.DeadlineExceeded
	})
	_, err := dial(context.Background(), "tcp", "redis.example:6379")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("lost original error: %v", err)
	}
	for _, want := range []string{`"stage":"dial"`, `"address":"redis.example:6379"`, `"error_class":"deadline"`} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("missing %s in %s", want, out.String())
		}
	}
}

func TestCommandFailureDoesNotLogArguments(t *testing.T) {
	var out bytes.Buffer
	h := Hook{Logger: slog.New(slog.NewJSONHandler(&out, nil))}
	want := errors.New("NOPERM this user has no permissions to run the 'auth' command")
	process := h.ProcessHook(func(context.Context, redis.Cmder) error { return want })
	err := process(context.Background(), redis.NewCmd(context.Background(), "auth", "private-user", "private-password"))
	if !errors.Is(err, want) {
		t.Fatal("lost original error")
	}
	if strings.Contains(out.String(), "private-user") || strings.Contains(out.String(), "private-password") {
		t.Fatalf("command arguments leaked: %s", out.String())
	}
	if !strings.Contains(out.String(), `"command":"auth"`) {
		t.Fatalf("missing command name: %s", out.String())
	}
}
