package types

import "github.com/TimeleapLabs/unchained/internal/config"

type Token struct {
	ID     *string  `json:"id"`
	Chain  string   `json:"chain"`
	Name   string   `json:"name"`
	Pair   string   `json:"pair"`
	Unit   string   `json:"unit"`
	Symbol string   `json:"symbol"`
	Delta  int64    `json:"delta"`
	Invert bool     `json:"invert"`
	Store  bool     `json:"store"`
	Send   bool     `json:"send"`
	Cross  []string `json:"cross"`
}

func (t Token) GetCrossTokenKeys(crossTokens map[string]TokenKey) TokenKeys {
	var cross []TokenKey

	for _, id := range t.Cross {
		cross = append(cross, crossTokens[id])
	}

	return cross
}

func NewTokensFromCfg(input []config.Token) []Token {
	result := []Token{}
	for _, t := range input {
		result = append(result, NewTokenFromCfg(t))
	}

	return result
}

func NewTokenFromCfg(input config.Token) Token {
	return Token{
		Chain:  input.Chain,
		Name:   input.Name,
		Pair:   input.Pair,
		Unit:   input.Unit,
		Delta:  input.Delta,
		Invert: input.Invert,
		Send:   input.Send,
	}
}
