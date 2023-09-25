package chain

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
)

type IChain interface {
	OnHappen()
	Write()
	Consensus()
}

type Category = string

type Chain struct {
	VN       Cid      `bson:"v" json:"v"`
	Scope    Cid      `bson:"s" json:"s"`
	Category Category `bson:"c" json:"c"`

	_serialize string
}

func (c *Chain) Same(oChn *Chain) bool {
	return c.VN == oChn.VN &&
		c.Scope == oChn.Scope &&
		c.Category == oChn.Category
}

func (c *Chain) SerializeString() string {
	if len(c._serialize) == 0 {
		str := fmt.Sprintf("v=%s|s=%s|c=%s", c.VN, c.Scope, c.Category)
		c._serialize = strs.MD5(str)
	}
	return c._serialize
}

type Block struct {
	Chain    Chain `bson:"c" json:"c"`
	Numb     int64 `bson:"n" json:"n"`
	Data     Cid   `bson:"d" json:"d"`
	Previous Cid   `bson:"p" json:"p"`

	_serialize string
}

func (b *Block) Next(data Cid, thisCid Cid) *Block {
	return &Block{
		Chain:    b.Chain,
		Numb:     b.Numb + 1,
		Data:     data,
		Previous: thisCid,
	}
}

func (b *Block) SerializeString() string {
	//if len(b._serialize) == 0 {
	//	str := fmt.Sprintf("c=%s|n=%d|d=%s|p=%s", b.Chain.SerializeString(),
	//		b.Numb, b.Data, b.Previous)
	//	fmt.Println("str===", str)
	//	b._serialize = strs.MD5(str)
	//}
	//return b._serialize
	str := fmt.Sprintf("c=%s|n=%d|d=%s|p=%s", b.Chain.SerializeString(),
		b.Numb, b.Data, b.Previous)
	//fmt.Println(str)
	return strs.MD5(str)
}

func (b *Block) GetType() Type {
	return Types.Block
}

func (b *Block) GetVn() Cid {
	return b.Chain.VN
}

func (b *Block) GetScope() Cid {
	return b.Chain.Scope
}

const (
	HeadBlockNumb     = 0
	HeadBlockData     = "YI"
	HeadBlockPrevious = "YI"
)

func NewHeadBlock(chn Chain) (Cid, *Block, *errors.Error) {
	b := &Block{
		Chain:    chn,
		Numb:     HeadBlockNumb,
		Data:     HeadBlockData,
		Previous: HeadBlockPrevious,
	}
	bCid, err := GetStone().Inscribe(b)
	if err != nil {
		return NilCid, nil, errors.Sys("e", err)
	}
	return bCid, b, nil
}
