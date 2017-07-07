package model

/*
Asset BO
*/
type Asset struct {
	ID      int    `json:"id"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}
