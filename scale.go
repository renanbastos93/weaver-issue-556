package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type ImageScaler interface {
	Scale(ctx context.Context, image string) (string, error)
}

type scaler struct {
	weaver.Implements[ImageScaler]
	weaver.Ref[LocalCache]
}

func (e *scaler) Init(ctx context.Context) error {
	return nil
}

func (e *scaler) Scale(ctx context.Context, image string) (string, error) {
	return "", nil
}
