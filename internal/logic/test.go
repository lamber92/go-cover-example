package logic

import (
	"context"

	"github.com/lamber92/go-cover-example/internal/model"
)

func Test1(ctx context.Context) (string, error) {
	return model.Test1(ctx)
}

func Test2(ctx context.Context, flag string) (string, error) {
	if len(flag) == 0 {
		return model.Test2(ctx)
	}
	return "flag_test_2", nil
}
