package vn

type Addr = string
type Code = string
type Number = int64

type ValuableNet struct {
	Addr    Addr   `bson:"addr" json:"addr"`
	Code    Code   `bson:"code" json:"code"`
	Number  Number `bson:"number" json:"number"`
	Founder string `bson:"founder" json:"founder"`
}

func Of(code string, number int64, founder string) ValuableNet {
	return ValuableNet{
		Addr:    "",
		Code:    code,
		Number:  number,
		Founder: founder,
	}
}

type Scope struct {
	VN    Addr   `bson:"vn" json:"vn"`
	Scope string `bson:"scope" json:"scope"`
}
