package edeeste

type DataCategory struct {
	Category Category `json:"category"`
	Files    []File   `json:"files"`
}

type Category struct {
	Access      int64  `json:"access"`
	Count       int64  `json:"count"`
	Description string `json:"description"`
	Filter      string `json:"filter"`
	Name        string `json:"name"`
	TermID      int64  `json:"term_id"`
	TermGroup   int64  `json:"term_group"`
}

type File struct {
	ID           int64  `json:"ID"`
	CatName      string `json:"catname"`
	LinkDownload string `json:"linkdownload"`
}
