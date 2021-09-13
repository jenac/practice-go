package exjson

type ResultForJSON struct {
	Data struct {
		Directors []string `json:"directors"`
		Rate      string   `json:"rate"`
		Cover     int      `json:"cover_x"`
		Star      string   `json:"start"`
		Title     string   `json:"title"`
		URL       string   `json:"url"`
		Casts     []string `json:"casts"`
	} `json:"data"`
}
