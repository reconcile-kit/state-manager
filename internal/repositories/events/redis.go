package events

import (
	"context"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisClient {
	return &RedisClient{client: redis}
}

func (rc *RedisClient) Add(ctx context.Context, shardID, messageType string, resource dto.ResourceID) error {
	stream := shardID + "_stream"
	if err := rc.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]any{
			"resource_group": resource.ResourceGroup,
			"kind":           resource.Kind,
			"namespace":      resource.Namespace,
			"name":           resource.Name,
			"type":           messageType,
		},
	}).Err(); err != nil {
		return err
	}
	return nil
}
