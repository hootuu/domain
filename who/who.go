package who

import (
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/inject"
	"github.com/hootuu/domain/ref"
	"github.com/hootuu/domain/scope"
)

type Addr = string

type Who struct {
	Scope scope.Lead `bson:"scope" json:"scope"`
	Addr  Addr       `bson:"addr" json:"addr"`
	Ref   ref.Ref    `bson:"ref" json:"ref"`
}

func (w Who) Inscribe() (chain.Cid, *chain.Lead, error) {
	cw := &ChainWho{
		Type:  chain.Types.Who,
		Scope: w.Scope,
		Addr:  w.Addr,
		Code:  w.Ref.Code,
		ID:    w.Ref.ID,
		Tag:   w.Ref.Tag,
		Attr:  w.Ref.Attr,
	}
	cid, err := inject.GetDataContainer().Put(cw)
	if err != nil {
		return chain.NilCid, nil, err
	}
	lead, err := inject.GetLinker().Append(w.Scope.VN, chain.KeyOf(w.Scope.Scope, chain.WHO), cid)
	if err != nil {
		return chain.NilCid, nil, err
	}
	return cid, lead, nil
}

type ChainWho struct {
	Type  chain.Type        `bson:"t" json:"t"`
	Scope scope.Lead        `bson:"scope" json:"scope"`
	Addr  Addr              `bson:"addr" json:"addr"`
	Code  string            `bson:"code" json:"code"`
	ID    string            `bson:"id" json:"id"`
	Tag   []string          `bson:"tag,omitempty" json:"tag,omitempty"`
	Attr  map[string]string `bson:"attr,omitempty" json:"attr,omitempty"`
}

func (c ChainWho) GetType() chain.Type {
	return c.Type
}
