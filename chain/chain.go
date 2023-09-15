package chain

import (
	"encoding/hex"
	"strings"
)

type Type string

var SCOPE = TypeOf("SCOPE")
var WHO = TypeOf("WHO")
var WHAT = TypeOf("WHAT")
var WHERE = TypeOf("WHERE")
var EVENT = TypeOf("EVENT")

func TypeOf(strType string) Type {
	hexStr := hex.EncodeToString([]byte(strType))
	return Type(hexStr)
}

type Data interface {
	GetType() Type
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
