package todo

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
	Done  bool   `json:"done,omitempty"`
}
