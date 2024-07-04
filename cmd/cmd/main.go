package main

import (
	"fmt"
	argparser "webpcdn/internal/adapters/primary/arg"
	cmdadapter "webpcdn/internal/adapters/primary/cmd"
	filecacher "webpcdn/internal/adapters/secondary/cache/file"
	"webpcdn/internal/adapters/secondary/storage/filestorage"
	converterservice "webpcdn/internal/core/services/converter"
)

func main() {
	adapter := cmdadapter.New(
		argparser.New(),
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
