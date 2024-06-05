package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
)

// type localCache struct {
// 	weaver.Implements[LocalCache]
// 	mu    sync.Mutex
// 	test  weaver.Ref[ImageScaler]
// 	cache *lru.Cache[string, string]
// 	// TODO: Eviction policy.
// }

// type scaler struct {
// 	weaver.Implements[ImageScaler]
// 	weaver.Ref[LocalCache]
// }

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
