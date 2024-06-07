package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
)

type implBFF struct {
	weaver.Implements[weaver.Main]
	scaler weaver.Ref[ImageScaler]
}

func Serve(ctx context.Context, e *implBFF) error {
	_ = e.scaler.Get()
	return nil
}

func main() {
	if err := weaver.Run(context.Background(), Serve); err != nil {
		log.Fatal(err)
	}
}
