package main

import (
	"context"
	"errors"

	"github.com/ServiceWeaver/weaver"
)

type LocalCache interface {
	Get(ctx context.Context, key string) (value string, err error)
	Set(ctx context.Context, key string, value string) (err error)
}

type implLocalCache struct {
	weaver.Implements[LocalCache]
	scaler weaver.Ref[ImageScaler]
}

func (e *implLocalCache) Init(ctx context.Context) error {
	return nil
}

func (e *implLocalCache) Get(ctx context.Context, key string) (value string, err error) {
	_ = e.scaler.Get()
	return "", errors.New("not implemented")
}

func (e *implLocalCache) Set(ctx context.Context, key string, value string) (err error) {
	return nil
}
