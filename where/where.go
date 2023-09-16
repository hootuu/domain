package where

import (
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/inject"
	"github.com/hootuu/domain/ref"
)

type Addr = string

type Where struct {
	Ref ref.Ref `bson:"ref" json:"ref"`
}

func (w Where) ToChainWhere() (chain.Cid, *ChainWhere, error) {
	cw := &ChainWhere{
		Type: chain.Types.Where,
		Code: w.Ref.Code,
		ID:   w.Ref.ID,
		Tag:  w.Ref.Tag,
		Attr: w.Ref.Attr,
	}
	cid, err := inject.GetDataContainer().Put(cw)
	if err != nil {
		return chain.NilCid, nil, err
	}
	return cid, cw, nil
}

type ChainWhere struct {
	Type chain.Type        `bson:"t" json:"t"`
	Code string            `bson:"code" json:"code"`
	ID   string            `bson:"id" json:"id"`
	Tag  []string          `bson:"tag,omitempty" json:"tag,omitempty"`
	Attr map[string]string `bson:"attr,omitempty" json:"attr,omitempty"`
}

func (c ChainWhere) GetType() chain.Type {
	return c.Type
}
