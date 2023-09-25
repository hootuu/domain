package chain

import (
	"encoding/json"
	"github.com/hootuu/utils/sys"
)
import "github.com/hootuu/utils/strs"

type Stone interface {
	Inscribe(data Data) (Cid, error)
	Get(cid Cid) (interface{}, error)
}

type nilStone struct {
	dict map[string]string
}

func (n *nilStone) Inscribe(data Data) (Cid, error) {
	jsonStr, _ := json.Marshal(data)
	id := strs.MD5(string(jsonStr))
	n.dict[id] = string(jsonStr)
	return id, nil
}

func (n *nilStone) Get(cid Cid) (interface{}, error) {
	d, ok := n.dict[cid]
	if !ok {
		return nil, nil
	}
	return d, nil
}

var gStone Stone

func InjectStone(s Stone) {
	gStone = s
}

func GetStone() Stone {
	if gStone == nil {
		if sys.RunMode.IsLocal() {
			gStone = &nilStone{dict: make(map[string]string)}
			return gStone
		}
		sys.Error("must inject chain.InjectStone first")
	}
	return gStone
}
