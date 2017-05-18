package requestmodel

type CheckIn struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Message string `json:"message"`
}
