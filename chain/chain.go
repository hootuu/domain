package chain

import (
	"strings"
)

type Type int32

type types struct {
	ValuableNet Type
	Scope       Type
	Who         Type
	What        Type
	Where       Type
	Event       Type
	Link        Type
}

var Types = types{
	ValuableNet: 3,
	Scope:       14,
	Who:         15,
	What:        92,
	Where:       65,
	Event:       35,
	Link:        89,
}

type Data interface {
	GetType() Type
	GetVn() Cid
	GetScope() Cid
}

type Key = string

func KeyOf(arr ...string) Key {
	return strings.Join(arr, "_")
}

type Node struct {
	Pre *Node `bson:"pre" json:"pre"`
	Cid Cid   `bson:"cid" json:"cid"`
	Nxt *Node `bson:"nxt" json:"nxt"`
}

type Lead struct {
	Head Cid `bson:"h" json:"h"`
	Tail Cid `bson:"t" json:"t"`
}

type Cid = string

const NilCid Cid = ""

func CidOf(cidStr string) Cid {
	return Cid(cidStr)
}
