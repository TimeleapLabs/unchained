package model

import "github.com/TimeleapLabs/unchained/internal/config"

type Token struct {
	ID     *string  `mapstructure:"id"`
	Chain  string   `mapstructure:"chain"`
	Name   string   `mapstructure:"name"`
	Pair   string   `mapstructure:"pair"`
	Unit   string   `mapstructure:"unit"`
	Symbol string   `mapstructure:"symbol"`
	Delta  int64    `mapstructure:"delta"`
	Invert bool     `mapstructure:"invert"`
	Store  bool     `mapstructure:"store"`
	Send   bool     `mapstructure:"send"`
	Cross  []string `mapstructure:"cross"`
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
