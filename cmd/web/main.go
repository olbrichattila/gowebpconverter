package main

import (
	"fmt"
	webadapter "webpcdn/internal/adapters/primary/web"
	filecacher "webpcdn/internal/adapters/secondary/cache/file"
	"webpcdn/internal/adapters/secondary/storage/filestorage"
	converterservice "webpcdn/internal/core/services/converter"
)

func main() {
	adapter := webadapter.New(
		converterservice.New(
			filestorage.New(),
		),
		filecacher.New(
			filestorage.New(),
		),
	)
	err := adapter.Run()
	if err != nil {
		fmt.Println(err)
	}
}
