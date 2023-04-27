package structs

type PairPrice struct {
	TokenAddress string `json:"token_address"`
	BaseAddress  string `json:"base_address"`
	Price        string `json:"price"`
}
