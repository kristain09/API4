package handler

type ExpectedRespond struct {
	ID    uint   `json:"id"`
	Judul string `json:"judul"`
	Tahun string `json:"tahun"`
	Nama  string `json:"nama"`
}
