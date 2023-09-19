package main

import (
	"github.com/Wazven/SOLID-Principle-Golang/OCP/after"
)
func main() {
	gudang := &after.Gudang{}
	logger := &after.Logger{}

	//logger catat informasi
	inputItem(gudang, logger, "Item 1")
	inputItem(gudang, logger, "Item 2")

	gudang.ListItems()
}

func inputItem(penyimpanan after.Penyimpanan, logger *after.Logger, item string) {
    penyimpanan.AddItem(item)
    logger.Log("Item " + item + " ditambahkan ke gudang")
}