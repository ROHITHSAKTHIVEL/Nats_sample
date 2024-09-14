package wbus

import (
	"context"
	"project/dto"
	"project/nats"
)

func Publish(ctx context.Context, req *dto.Request) {
	nats.Publish(ctx, req)
}

func Subscribe(ctx context.Context,req *dto.Request) {
	nats.Subscribe(ctx,req )
}
