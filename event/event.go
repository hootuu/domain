package event

import (
	"errors"
	"github.com/hootuu/domain/chain"
	"github.com/hootuu/domain/howm"
	"github.com/hootuu/domain/inject"
	"github.com/hootuu/domain/what"
	"github.com/hootuu/domain/where"
	"github.com/hootuu/domain/who"
	"sync"
)

type Act = string

type Event struct {
	Scope   chain.Lead   `bson:"scope" json:"scope"`
	Who     who.Who      `bson:"who" json:"who"`
	Where   where.Where  `bson:"where" json:"where"`
	Act     Act          `bson:"act" json:"act"`
	What    what.What    `bson:"what" json:"what"`
	HowMuch howm.HouMuch `bson:"how_much" json:"how_much"`
}

func (e Event) Inscribe() (chain.Cid, *chain.Lead, error) {
	ce := &ChainEvent{
		Type:    chain.EVENT,
		Scope:   e.Scope,
		Act:     e.Act,
		HowMuch: e.HowMuch,
	}

	hasErr := false
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		var whoErr error
		ce.Who, _, whoErr = e.Who.Inscribe()
		if whoErr != nil {
			hasErr = true
		}
	}()
	go func() {
		defer wg.Done()
		var whatErr error
		ce.What, _, whatErr = e.What.ToChainWhat()
		if whatErr != nil {
			hasErr = true
		}
	}()
	go func() {
		defer wg.Done()
		var whereErr error
		ce.Where, _, whereErr = e.Where.ToChainWhere()
		if whereErr != nil {
			hasErr = true
		}
	}()
	wg.Wait()
	if hasErr {
		return chain.NilCid, nil, errors.New("put data failed")
	}

	cid, err := inject.GetDataContainer().Put(ce)
	if err != nil {
		return chain.NilCid, nil, err
	}

	lead, err := inject.GetLinker().Append(e.Scope.VN, chain.KeyOf(e.Scope.Scope, chain.EVENT), cid)
	if err != nil {
		return chain.NilCid, nil, err
	}

	return cid, lead, nil
}

type ChainEvent struct {
	Type    chain.Type   `bson:"t" json:"t"`
	Scope   chain.Lead   `bson:"scope" json:"scope"`
	Who     chain.Cid    `bson:"who" json:"who"`
	Where   chain.Cid    `bson:"where" json:"where"`
	Act     Act          `bson:"act" json:"act"`
	What    chain.Cid    `bson:"what" json:"what"`
	HowMuch howm.HouMuch `bson:"how_much" json:"how_much"`
}

func (c ChainEvent) GetType() chain.Type {
	return c.Type
}
