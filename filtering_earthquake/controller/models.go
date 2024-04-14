package controller

import (
	"sync"

	"github.com/alexandrevicenzi/go-sse"
)

type IFilteringService interface {
	HashEarthquakes() error
}

type FilteringController struct {
	Srv   IFilteringService
	SseSv *sse.Server
}

var (
	OldKeys = make(map[string]int)
	mutex sync.Mutex
)