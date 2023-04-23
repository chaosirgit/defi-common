package structs

type ReadyTrade struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	DefiAddress string `json:"defi_address"`
	ChainId     string `json:"chain_id"`
	PrivateKey  string `json:"private_key"`
	Amount      string `json:"amount"`
	Account     string `json:"account"`
}
