package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// lengte van wachtwoord opgeven

// keuze uit tekens? getallen, letters, tekens, symbolen

// genereer uit gekozen tekens en lengte

// optioneel gebruiker kan wachtwoord invoeren en controleer of dit goed is
const (
	DefaultLength    = 16
	DefaultLetterSet = "abcdefghijklmnopqrstuvwxyz"
	DefaultNumberSet = "0123456789"
	DefaultSymbolSet = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
)

var (
	Defaultconfig = Config{
		Length:                  DefaultLength,
		IncludeNumbers:          true,
		IncludeSymbols:          true,
		IncludeLowercaseLetters: true,
		IncludeUppercaseLetters: true,
	}
)

type Generator struct {
	*Config
}

func New(config *Config) (*Generator, error) {
	if config == nil {
		fmt.Println("Geen config opgegeven, default config wordt gebruikt")
		config = &Defaultconfig
	}

	if config.Characterset == "" {
		config.Characterset = buildCharacterSet(config)
	}

	return &Generator{Config: config}, nil
}

func buildCharacterSet(config *Config) string {
	var characterSet string

	if config.IncludeNumbers {
		characterSet += DefaultNumberSet
	}

	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
	}

	if config.IncludeLowercaseLetters {
		characterSet += DefaultLetterSet
	}

	if config.IncludeUppercaseLetters {
		characterSet += strings.ToUpper(DefaultLetterSet)
	}

	return characterSet
}

func (g Generator) Generate() (*string, error) {
	var generated string

	characterSet := strings.Split(g.Config.Characterset, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

func (g Generator) GenerateMany(amount int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}
	return generated, nil
}
