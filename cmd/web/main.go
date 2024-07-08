package main

import (
	"fmt"
	webadapter "webpcdn/internal/adapters/primary/web"
	filecacher "webpcdn/internal/adapters/secondary/cache/file"
	storerfactory "webpcdn/internal/adapters/secondary/storage/factory"
	"webpcdn/internal/adapters/secondary/storage/filestorage"
	"webpcdn/internal/core/domain/config"
	converterservice "webpcdn/internal/core/services/converter"
)

func main() {
	config := config.New()
	stype := config.GetStorageType()
	cacheStorer, err := storerfactory.New(stype)
	if err != nil {
		fmt.Println(err)
		return
	}

	adapter := webadapter.New(
		converterservice.New(
			filestorage.New(),
		),
		filecacher.New(
			cacheStorer,
		),
	)
	err = adapter.Run()
	if err != nil {
		fmt.Println(err)
	}
}
