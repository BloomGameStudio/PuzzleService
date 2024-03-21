package models

type ERC20 struct {
	TokenName   string `json:"token_name"`
	TokenSymbol string `json:"token_symbol"`
	Decimals    int    `json:"decimals"`
	TotalSupply int    `json:"total_supply"`
}