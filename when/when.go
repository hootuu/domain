package when

import (
	"errors"
	"time"
)

type When struct {
	Loc  string `bson:"loc" json:"loc"`
	Unix int64  `bson:"unix" json:"unix"`
}

func Now() When {
	t := time.Now()
	return When{
		Loc:  t.Location().String(),
		Unix: t.UnixMilli(),
	}
}

func Of(loc string, unix int64) (*When, error) {
	if len(loc) == 0 {
		return nil, errors.New("require loc")
	}
	if unix < 0 {
		return nil, errors.New("unix < 0")
	}
	tLoc, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	t := time.UnixMilli(unix)
	tInLoc := t.In(tLoc)
	return &When{
		Loc:  tInLoc.Location().String(),
		Unix: tInLoc.UnixMilli(),
	}, nil
}
