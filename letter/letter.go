package letter

import (
	"encoding/json"
	"github.com/hootuu/domain/hotu"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"github.com/hootuu/utils/sys"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"time"
)

type ILetter interface {
	GetTopic() string
	GetFrom() string
	GetId() string
	GetTimestamp() int64
	GetPayload() string
}

type Letter struct {
	topic     string
	id        string
	timestamp int64
	payload   string
}

func NewLetter(topic string, payload interface{}) *Letter {
	l := &Letter{
		topic:     topic,
		id:        xid.New().String(),
		timestamp: time.Now().UnixMilli(),
	}
	jsonByte, _ := json.Marshal(payload)
	l.payload = string(jsonByte)
	return l
}

func (l *Letter) GetTopic() string {
	return l.topic
}

func (l *Letter) GetFrom() string {
	return hotu.Hotu().GetID()
}

func (l *Letter) GetId() string {
	return l.id
}

func (l *Letter) GetTimestamp() int64 {
	return l.timestamp
}

func (l *Letter) GetPayload() string {
	return l.payload
}

type IListener interface {
	GetTopic() []string
	Deal(ltr ILetter) *errors.Error
}

type IPostOffice interface {
	Broadcast(ltr ILetter)
	Register(listener IListener)
}

type NilPostOffice struct {
	listeners []IListener
}

func NewNilPostOffice() *NilPostOffice {
	return &NilPostOffice{listeners: []IListener{}}
}

func (n *NilPostOffice) Broadcast(ltr ILetter) {
	//logger.Logger.Info("Broadcast letter:",
	//	zap.String("letter.topic", ltr.GetTopic()),
	//	zap.String("letter.id", ltr.GetId()),
	//	zap.String("letter.from", ltr.GetFrom()),
	//	zap.String("letter.payload", ltr.GetPayload()))
	for _, iListener := range n.listeners {
		topics := iListener.GetTopic()
		for _, t := range topics {
			if t == ltr.GetTopic() {
				err := iListener.Deal(ltr)
				if err != nil {
					logger.Logger.Error("handle letter failed",
						zap.String("letter.topic", ltr.GetTopic()),
						zap.String("letter.id", ltr.GetId()),
						zap.String("letter.from", ltr.GetFrom()),
						zap.String("letter.payload", ltr.GetPayload()))
				}
			}
		}
	}
}

func (n *NilPostOffice) Register(listener IListener) {
	n.listeners = append(n.listeners, listener)
}

var gPostOffice IPostOffice

func Inject(pf IPostOffice) {
	gPostOffice = pf
}

func PostOffice() IPostOffice {
	if gPostOffice == nil {
		if sys.RunMode.IsLocal() {
			gPostOffice = NewNilPostOffice()
			return gPostOffice
		}
		sys.Error("Must Inject Post Office First")
		return nil
	}
	return gPostOffice
}
