package inject

import (
	"fmt"
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/tools"
	"github.com/hootuu/domain/vn"
)

type DataContainer interface {
	Put(data chain.Data) (chain.Cid, error)
	Get(cid chain.Cid) (chain.Data, error)
}

type nilDataContainer struct {
	db map[chain.Cid]chain.Data
}

func (n nilDataContainer) Put(data chain.Data) (chain.Cid, error) {
	cid := chain.CidOf(tools.Md5(fmt.Sprintf("%v", data)))
	n.db[cid] = data
	return cid, nil
}

func (n nilDataContainer) Get(cid chain.Cid) (chain.Data, error) {
	d, ok := n.db[cid]
	if !ok {
		return nil, nil
	}
	return d, nil
}

type Linker interface {
	Append(vnCid chain.Cid, chainKey chain.Key, cid chain.Cid) (*chain.Lead, error)
}

type nilLinker struct {
	link map[vn.Addr]map[chain.Key]map[chain.Cid]*chain.Node
}

func (n nilLinker) Append(vnAddr vn.Addr, chainKey chain.Key, cid chain.Cid) (*chain.Lead, error) {
	if n.link == nil {
		n.link = make(map[vn.Addr]map[chain.Key]map[chain.Cid]*chain.Node)
	}
	vnLink, ok := n.link[vnAddr]
	if !ok {
		vnLink = make(map[chain.Key]map[chain.Cid]*chain.Node)
		n.link[vnAddr] = vnLink
	}
	chainLink, ok := vnLink[chainKey]
	if !ok {
		chainLink = make(map[chain.Cid]*chain.Node)
		vnLink[chainKey] = chainLink
	}
	thisNode := &chain.Node{
		Pre: nil,
		Cid: cid,
		Nxt: nil,
	}
	chainLink[cid] = thisNode
	lead := &chain.Lead{}
	for _, v := range chainLink {
		if v.Pre == nil {
			lead.Head = v.Cid
		}
		if v.Nxt == nil {
			v.Nxt = thisNode
			thisNode.Pre = v
		}
	}
	lead.Tail = thisNode.Cid
	return lead, nil
}
