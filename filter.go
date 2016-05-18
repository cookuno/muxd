package muxd

import "strings"

type filterFunc func()

type filterChain struct  {
	chain []filter
}

type filter struct  {
	method string
	path string
	filterFunc filterFunc
}

func (self *filterChain) Add(filterFunc filterFunc, path string, methods ...string)  {
	path = strings.TrimSpace(strings.Replace(path, "*", "", -1))
	if path == "" {
		path = "/"
	}
	methodLen := len(methods)
	if methodLen == 0 {
		methods = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "TRACE"}
	}
	methodLen = len(methods)
	for i := 0 ; i < methodLen ; i ++ {
		method := methods[i]
		self.chain = append(self.chain, filter{method:method, path:path, filterFunc:filterFunc})
	}
}