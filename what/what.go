package what

import (
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/inject"
	"github.com/hootuu/domain/ref"
)

type Addr = string

type What struct {
	Scope chain.Lead `bson:"scope" json:"scope"`
	Ref   ref.Ref    `bson:"ref" json:"ref"`
}

func (w What) Inscribe() (chain.Cid, *chain.Lead, error) {
	cw := &ChainWhat{
		Type: chain.WHAT,
		Code: w.Ref.Code,
		ID:   w.Ref.ID,
		Tag:  w.Ref.Tag,
		Attr: w.Ref.Attr,
	}
	cid, err := inject.GetDataContainer().Put(cw)
	if err != nil {
		return chain.NilCid, nil, err
	}
	lead, err := inject.GetLinker().Append(w.Scope.VN, chain.KeyOf(w.Scope.Scope, chain.WHAT, w.Ref.Code), cid)
	if err != nil {
		return chain.NilCid, nil, err
	}
	return cid, lead, nil
}

type ChainWhat struct {
	Type  chain.Type        `bson:"t" json:"t"`
	Scope chain.Lead        `bson:"scope" json:"scope"`
	Code  string            `bson:"code" json:"code"`
	ID    string            `bson:"id" json:"id"`
	Tag   []string          `bson:"tag,omitempty" json:"tag,omitempty"`
	Attr  map[string]string `bson:"attr,omitempty" json:"attr,omitempty"`
}

func (c ChainWhat) GetType() chain.Type {
	return c.Type
}
