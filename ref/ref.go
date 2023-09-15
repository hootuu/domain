package ref

import (
	"errors"
	"github.com/hootuu/domain/tools"
)

type Code = string
type ID = string

type Ref struct {
	Code Code              `bson:"code" json:"code"`
	ID   ID                `bson:"id" json:"id"`
	Tag  []string          `bson:"tag,omitempty" json:"tag,omitempty"`
	Attr map[string]string `bson:"attr,omitempty" json:"attr,omitempty"`
}

func (r Ref) Verify() error {
	if err := tools.NormalCodeVerify(r.Code); err != nil {
		return errors.New("invalid code: " + err.Error())
	}
	if err := tools.NormalIDVerify(r.ID); err != nil {
		return errors.New("invalid id: " + err.Error())
	}
	return nil
}

func (r Ref) PutAttribute(key string, val string) {
	if len(r.Attr) == 0 {
		r.Attr = make(map[string]string)
	}
	r.Attr[key] = val
}

func Of(code string, id string, tag ...string) Ref {
	return Ref{
		Code: Code(code),
		ID:   ID(id),
		Tag:  tag,
	}
}
