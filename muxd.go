package muxd

type Mux struct  {
	pTree *pathTree
	filters *filterChain
}

func New() *Mux {
	mux := Mux{}
	pTree := pathTree{}
	nodes := make(map[string]pathNode)
	pTree.nodes = &nodes
	mux.pTree = &pTree
	filters := filterChain{chain:make([]filter,0,8)}
	mux.filters = &filters
	return &mux
}

func (self *Mux) Get(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Post(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Put(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Delete(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Head(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Options(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) Trace(path string, handleFunc handleFunc) *Mux {

	return self
}

func (self *Mux) AddFilter(path string, handleFunc handleFunc) *Mux {

	return self
}
