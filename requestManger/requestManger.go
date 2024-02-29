package requestmanger

import (
	"container/list"
	"errors"
	"sync"
	"time"

	"github.com/Rx947getrexp/ArachneXGo/iface"
)

type (
	RequestManger struct {
		UrlList     *list.List
		UsedUrlList map[string]struct{}
		Headers     map[string]string
		timeout     time.Duration
		RetryTimes  uint8
		IPAgent     string
		general     sync.RWMutex
		addLock     sync.Mutex
		takeLock    sync.Mutex
	}

	request struct {
		url chan string
	}
)

func NewRequestManger() iface.IRequestManger {
	return &RequestManger{
		UrlList:     new(list.List),
		UsedUrlList: make(map[string]struct{}),
		Headers:     make(map[string]string),
		timeout:     1 << 5 * time.Second,
		RetryTimes:  1<<2 - 1,
	}
}

// AddRequest implements iface.IRequestManger.
func (r *RequestManger) AddRequest(uri iface.IUri) error {
	address := uri.GetAddress()
	r.addLock.Lock()
	defer r.addLock.Unlock()

	if _, exist := r.UsedUrlList[address]; exist {
		return errors.New(address + " is used.")
	}

	r.UrlList.PushBack(uri)
	return nil
}

// GetNextRequest implements iface.IRequestManger.
func (r *RequestManger) GetNextRequest() iface.IUri {
	r.takeLock.Lock()
	defer r.takeLock.Unlock()

	if element := r.UrlList.Front(); element != nil {
		return element.Value.(iface.IUri)
	}
	return nil
}

// MarkURLAsProcess implements iface.IRequestManger.
func (r *RequestManger) MarkURLAsProcess(url string) {
	r.general.RLock()
	defer r.general.RUnlock()

	r.UsedUrlList[url] = struct{}{}
}

// SetRequestTimeout implements iface.IRequestManger.
func (r *RequestManger) SetRequestTimeout(time time.Duration) iface.IRequestManger {
	r.timeout = time
	return r
}

// SetRequestheaders implements iface.IRequestManger.
func (r *RequestManger) SetRequestheaders(key string, value string) iface.IRequestManger {
	r.Headers[key] = value
	return r
}

var _ iface.IRequestManger = (*RequestManger)(nil)
