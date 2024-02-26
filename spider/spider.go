package spider

import (
	"sync"
)

type Spider struct {
	seedURL    string //种子url
	queue      chan string
	visitedURL map[string]struct{}
	mu         sync.Mutex
}

func NewSpider(seedURL string, concurrency int) {
	return &Spider{
		seedURL:    seedURL,
		queue:      make(chan string),
		visitedURL: make(map[string]struct{}),
	}
}

func (s *Spider) Run() {
	
}

func (s *Spider) Fetch(url string) {

}

var _ iface.ISpider = (*Spider)(nil)
