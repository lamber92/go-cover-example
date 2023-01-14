package logic

import (
	"context"

	"github.com/lamber92/go-cover-example/internal/model"
)

func Test1(ctx context.Context) (string, error) {
	return model.Test1(ctx)
}

func Test2(ctx context.Context) (string, error) {
	return model.Test2(ctx)
}
