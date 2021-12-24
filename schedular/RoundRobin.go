package schedular

import (
	"errors"
	"net/url"
	"sync/atomic"
)

var ErrorServerNotExist = errors.New("server does not exist")

type roundRobin struct {
	urls []*url.URL
	next uint64
}

type RoundRobin interface {
	Next() *url.URL
}

func New(urls ...*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrorServerNotExist
	}
	return &roundRobin{urls: urls}, nil
}

func (r *roundRobin) Next() *url.URL {
	n := atomic.AddUint64(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}
