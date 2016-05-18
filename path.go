package muxd

type handleFunc func(c *Context)

type pathTree struct  {
	nodes *map[string]pathNode
}

type pathNode struct  {
	part string
	next *map[string]pathNode
	hasNext bool
	handleFuncMap *map[string]handleFunc
}

func (self *pathTree) add(path string, method string, handleFunc handleFunc)  {

}
