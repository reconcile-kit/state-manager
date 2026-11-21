// Package redisdiag logs Redis failures without recording URLs, credentials or command arguments.
package redisdiag

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log/slog"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
)

type Hook struct {
	Logger *slog.Logger
}

func errorClass(err error) string {
	var dns *net.DNSError
	var cert *tls.CertificateVerificationError
	var unknownCA x509.UnknownAuthorityError
	var record tls.RecordHeaderError
	var network net.Error
	switch {
	case errors.As(err, &dns):
		return "dns"
	case errors.As(err, &cert), errors.As(err, &unknownCA):
		return "tls_certificate"
	case errors.As(err, &record):
		return "tls_protocol"
	case errors.Is(err, context.DeadlineExceeded):
		return "deadline"
	case errors.Is(err, context.Canceled):
		return "canceled"
	case err.Error() == "redis: connection pool timeout":
		return "pool_timeout"
	case errors.As(err, &network) && network.Timeout():
		return "network_timeout"
	default:
		return "redis_or_network"
	}
}

func (h Hook) failure(ctx context.Context, stage string, start time.Time, err error, fields ...any) {
	if err == nil || errors.Is(err, redis.Nil) {
		return
	}
	args := []any{"stage", stage, "duration_ms", time.Since(start).Milliseconds(), "error_class", errorClass(err), "error", err.Error()}
	args = append(args, fields...)
	h.Logger.ErrorContext(ctx, "redis operation failed", args...)
}

func (h Hook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		start := time.Now()
		conn, err := next(ctx, network, addr)
		h.failure(ctx, "dial", start, err, "network", network, "address", addr)
		return conn, err
	}
}

func (h Hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		err := next(ctx, cmd)
		h.failure(ctx, "command", start, err, "command", cmd.Name())
		return err
	}
}

func (h Hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		start := time.Now()
		err := next(ctx, cmds)
		h.failure(ctx, "pipeline", start, err, "command_count", len(cmds))
		return err
	}
}
