package scope

import (
	"github.com/hootuu/domain/ref"
	"github.com/hootuu/utils/errors"
)

type Addr string
type Code string
type Number int64

type Lead struct {
	VN    string `bson:"vn" json:"vn"`
	Scope string `bson:"scope" json:"scope"`
}

func (l *Lead) Verify() *errors.Error {
	//todo
	return nil
}

type Scope struct {
	VN      string  `bson:"vn" json:"vn"`
	Addr    Addr    `bson:"addr" json:"addr"`
	Code    Code    `bson:"code" json:"code"`
	Number  Number  `bson:"number" json:"number"`
	Founder string  `bson:"founder" json:"founder"`
	Ref     ref.Ref `bson:"ref" json:"ref"`
}

//func (e Scope) Inscribe() (string, *Lead, error) {
//	cs := &ChainScope{
//		Type:    chain.Types.Scope,
//		VN:      e.VN,
//		Addr:    "",
//		Code:    e.Code,
//		Number:  e.Number,
//		Founder: e.Founder,
//		Ref:     e.Ref,
//	}
//
//	cid, err := inject.GetDataContainer().Put(cs)
//	if err != nil {
//		return chain.NilCid, nil, err
//	}
//
//	lead, err := inject.GetLinker().Append(cs.VN, chain.KeyOf(SCOPE), cid)
//	if err != nil {
//		return chain.NilCid, nil, err
//	}
//
//	return cid, lead, nil
//}

type ChainScope struct {
	Type    int32   `bson:"t" json:"t"`
	VN      string  `bson:"vn" json:"vn"`
	Addr    Addr    `bson:"addr" json:"addr"`
	Code    Code    `bson:"code" json:"code"`
	Number  Number  `bson:"number" json:"number"`
	Founder string  `bson:"founder" json:"founder"`
	Ref     ref.Ref `bson:"ref" json:"ref"`
}

func (c ChainScope) GetType() int32 {
	return c.Type
}
