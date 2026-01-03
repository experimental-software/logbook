package core

type LogbookEntry struct {
	DateTime  string `json:"dateTime"`
	Title     string `json:"title"`
	Directory string `json:"directory"`
}
