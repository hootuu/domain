package chain

import "log/slog"

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
		slog.Error("must inject chain.InjectStone first")
	}
	return gStone
}
