package store

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func mergeSeatJSON(ctx context.Context, p *Postgres, area, seatID string, update map[string]any) (string, bool) {
	raw, err := p.cache.HGet(ctx, "seats:"+area, seatID).Result()
	if err == redis.Nil {
		return "", false
	}
	item := map[string]any{}
	if err != nil || json.Unmarshal([]byte(raw), &item) != nil {
		return "", false
	}
	item["selectedBy"] = update["selectedBy"]
	item["updatedAt"] = update["updatedAt"]
	item["expirationTime"] = update["expirationTime"]
	item["reservedUserId"] = update["reservedUserId"]
	out, _ := json.Marshal(item)
	return string(out), true
}
