package howm

type HouMuch struct {
	Amount   int64              `bson:"amount" json:"amount"`
	Currency string             `bson:"currency" json:"currency"`
	Ex       map[string]HouMuch `bson:"ex,omitempty" json:"ex,omitempty"`
}

func (h HouMuch) Put(key string, hm HouMuch) {
	if len(h.Ex) == 0 {
		h.Ex = make(map[string]HouMuch)
	}
	h.Ex[key] = hm
}

func Of(amount int64, currency string) HouMuch {
	return HouMuch{
		Amount:   amount,
		Currency: currency,
	}
}
