package scope

import (
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/inject"
	"github.com/hootuu/domain/ref"
)

type Addr string
type Code string
type Number int64

type Lead struct {
	VN    chain.Cid `bson:"vn" json:"vn"`
	Scope chain.Cid `bson:"scope" json:"scope"`
}

type Scope struct {
	VN      chain.Cid `bson:"vn" json:"vn"`
	Addr    Addr      `bson:"addr" json:"addr"`
	Code    Code      `bson:"code" json:"code"`
	Number  Number    `bson:"number" json:"number"`
	Founder string    `bson:"founder" json:"founder"`
	Ref     ref.Ref   `bson:"ref" json:"ref"`
}

func (e Scope) Inscribe() (chain.Cid, *chain.Lead, error) {
	cs := &ChainScope{
		Type:    chain.SCOPE,
		VN:      e.VN,
		Addr:    "",
		Code:    e.Code,
		Number:  e.Number,
		Founder: e.Founder,
		Ref:     e.Ref,
	}

	cid, err := inject.GetDataContainer().Put(cs)
	if err != nil {
		return chain.NilCid, nil, err
	}

	lead, err := inject.GetLinker().Append(cs.VN, chain.KeyOf(chain.SCOPE), cid)
	if err != nil {
		return chain.NilCid, nil, err
	}

	return cid, lead, nil
}

type ChainScope struct {
	Type    chain.Type `bson:"t" json:"t"`
	VN      chain.Cid  `bson:"vn" json:"vn"`
	Addr    Addr       `bson:"addr" json:"addr"`
	Code    Code       `bson:"code" json:"code"`
	Number  Number     `bson:"number" json:"number"`
	Founder string     `bson:"founder" json:"founder"`
	Ref     ref.Ref    `bson:"ref" json:"ref"`
}

func (c ChainScope) GetType() chain.Type {
	return c.Type
}
