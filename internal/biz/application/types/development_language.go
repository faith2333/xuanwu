package types

import "strings"

const (
	DevelopmentLanguageJAVA       DevelopmentLanguage = "JAVA"
	DevelopmentLanguageCLANG      DevelopmentLanguage = "CLANG"
	DevelopmentLanguageGOLANG     DevelopmentLanguage = "GOLANG"
	DevelopmentLanguagePYTHON     DevelopmentLanguage = "PYTHON"
	DevelopmentLanguageTYPESCRIPT DevelopmentLanguage = "TYPESCRIPT"
	DevelopmentLanguageJAVASCRIPT DevelopmentLanguage = "JAVASCRIPT"
)

var AllDevelopmentLanguages = []DevelopmentLanguage{
	DevelopmentLanguageJAVA,
	DevelopmentLanguageCLANG,
	DevelopmentLanguageGOLANG,
	DevelopmentLanguagePYTHON,
	DevelopmentLanguageTYPESCRIPT,
	DevelopmentLanguageJAVASCRIPT,
}

type DevelopmentLanguage string

func (dl DevelopmentLanguage) String() string {
	return string(dl)
}

func (dl DevelopmentLanguage) Upper() DevelopmentLanguage {
	return DevelopmentLanguage(strings.ToUpper(dl.String()))
}

func (dl DevelopmentLanguage) IsSupported() bool {
	for _, sLang := range AllDevelopmentLanguages {
		if sLang == dl.Upper() {
			return true
		}
	}
	return false
}

func (dl DevelopmentLanguage) IsJava() bool {
	return dl.Upper() == DevelopmentLanguageJAVA
}

func (dl DevelopmentLanguage) IsGolang() bool {
	return dl.Upper() == DevelopmentLanguageGOLANG
}

func (dl DevelopmentLanguage) IsClang() bool {
	return dl.Upper() == DevelopmentLanguageCLANG
}

func (dl DevelopmentLanguage) IsTypeScript() bool {
	return dl.Upper() == DevelopmentLanguageTYPESCRIPT
}

func (dl DevelopmentLanguage) IsJavaScript() bool {
	return dl.Upper() == DevelopmentLanguageJAVASCRIPT
}
