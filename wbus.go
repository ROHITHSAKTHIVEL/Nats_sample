package wbus

import (
	"context"

	"github.com/ROHITHSAKTHIVEL/Nats_sample/dto"
	"github.com/ROHITHSAKTHIVEL/Nats_sample/nats"
)

func Publish(ctx context.Context, req *dto.Request) {
	nats.Publish(ctx, req)
}

func Subscribe(ctx context.Context, req *dto.Request) {
	nats.Subscribe(ctx, req)
}
