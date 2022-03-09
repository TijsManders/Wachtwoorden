package generator

type Config struct {
	Length                  int
	Characterset            string
	IncludeNumbers          bool
	IncludeSymbols          bool
	IncludeLowercaseLetters bool
	IncludeUppercaseLetters bool
}
