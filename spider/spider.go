package spider

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/net/html"

	"github.com/Rx947getrexp/ArachneXGo/iface"
)

type Spider struct {
	seedURL    string              // 种子URL
	queue      chan string         // URL队列
	visitedURL map[string]struct{} // 已访问过的URL
	mu         sync.Mutex          // 互斥锁
}

func NewSpider(seedURL string, concurrency int) iface.ISpider {
	return &Spider{
		seedURL:    seedURL,
		queue:      make(chan string),
		visitedURL: make(map[string]struct{}),
	}
}

// 启动爬虫框架
func (s *Spider) Run() {
	s.queue <- s.seedURL
	for link := range s.queue {
		if _, exists := s.visitedURL[link]; !exists {
			go s.Fetch(link)
		}
	}
}

// 发起 HTTP请求并解析页面
func (s *Spider) Fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求 %s 失败: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("解析 %s 内容失败: %v\n", url, err)
		return
	}

	s.ParseLinks(doc)
}

func (s *Spider) ParseLinks(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link := a.Val
				s.queue <- link
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s.ParseLinks(c)
	}
}

func (s *Spider) visited(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.visitedURL[url]; ok {
		return true
	}
	s.visitedURL[url] = struct{}{}
	return false
}

var _ iface.ISpider = (*Spider)(nil)
