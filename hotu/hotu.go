package hotu

import (
	"github.com/hootuu/domain/vn"
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/sys"
	"github.com/rs/xid"
)

type IHotu interface {
	GetID() string
	GetLead() vn.Lead
}

type NilHotu struct {
	id   string
	lead vn.Lead
}

func NewNilHotu() *NilHotu {
	return &NilHotu{
		id: xid.New().String(),
		lead: vn.Lead{
			VN:    configure.GetString("hotu.vn", "hotu.vn"),
			Scope: configure.GetString("hotu.scope", "hotu.scope"),
		},
	}
}

func (n *NilHotu) GetID() string {
	return n.id
}

func (n *NilHotu) GetLead() vn.Lead {
	return n.lead
}

var gHotu IHotu

func Inject(ht IHotu) {
	gHotu = ht
}

func Hotu() IHotu {
	if gHotu == nil {
		if sys.RunMode.IsLocal() {
			gHotu = NewNilHotu()
			return gHotu
		}
		sys.Error("Must Inject HOTU First")
		return nil
	}
	return gHotu
}
