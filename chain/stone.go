package chain

import "github.com/hootuu/utils/sys"

type Stone interface {
	Inscribe(data Data) (Cid, error)
	Get(cid Cid) (interface{}, error)
}

var gStone Stone

func InjectStone(s Stone) {
	gStone = s
}

func GetStone() Stone {
	if gStone == nil {
		sys.Error("must inject chain.InjectStone first")
	}
	return gStone
}
