package trade



type transaction struct {
}

type quote struct {
	currency string
	ammount float
}

type buy struct {
	timestamp string
	exchange string
	account string
	asset string
	price quote
	ammount float
	fees quote
	cost quote
	total quote

}

func (transaction) buyBTC() string {
	return "btc"
}

func (transaction) buyUSD() string {
	return "btc"
}

func (transaction) sellBTC() string {
	return "btc"
}

func (transaction) sellUSD() string {
	return "btc"
}

func (transaction) transfert() string {
	return "btc"
}