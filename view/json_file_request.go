package view

type WriteJsonFileRequest struct {
	Filename string                 `json:"filename"`
	Content  map[string]interface{} `json:"content"`
}