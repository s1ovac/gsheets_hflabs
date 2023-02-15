package config

type SheetConfig struct {
	sheetId       int
	spreadsheetID string
	credPath      string
	url           string
}

func NewSheetConfig() SheetConfig {
	return SheetConfig{
		sheetId:       0,
		spreadsheetID: "1c9rQsoZddsCylo8LYvT8Qj28XmofNOGXGdLIbgA7ANw",
		credPath:      "credentials.json",
		url:           "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999",
	}
}

func (s SheetConfig) SheetID() int {
	return s.sheetId
}

func (s SheetConfig) SpreadSheetID() string {
	return s.spreadsheetID
}

func (s SheetConfig) CredPath() string {
	return s.credPath
}

func (s SheetConfig) URL() string {
	return s.url
}
