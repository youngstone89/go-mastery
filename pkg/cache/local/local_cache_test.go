package local

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	ttlcache "github.com/jellydator/ttlcache/v3"
)

func Test(t *testing.T) {
	cache := ttlcache.New(
		ttlcache.WithTTL[string, any](1 * time.Second),
	)
	fnOnInsert := func(ctx context.Context, i *ttlcache.Item[string, any]) {
		fmt.Printf("key: %v, value: %v \n", i.Key(), i.Value())
	}
	cache.OnInsertion(fnOnInsert)

	fnOnEvict := func(ctx context.Context, er ttlcache.EvictionReason, i *ttlcache.Item[string, any]) {
		fmt.Printf("reason: %v,key: %v, value: %v \n", er, i.Key(), i.Value())
	}
	cache.OnEviction(fnOnEvict)
	uid := uuid.NewString()
	cache.Set(uid, "i am here..", ttlcache.DefaultTTL)
	go cache.Start() // starts automatic expired item deletion

	item := cache.Get(uid)
	fmt.Println(item.Value(), item.ExpiresAt())

	time.Sleep(1 * time.Second)

	item = cache.Get(uid)
	if item == nil {
		fmt.Println("expired..")
	} else {
		fmt.Println(item.Value(), item.ExpiresAt())
	}

	cache.Delete(uid)

	item = cache.Get(uid)
	if item == nil {
		fmt.Println("expired..")
	} else {
		fmt.Println(item.Value(), item.ExpiresAt())
	}

}
