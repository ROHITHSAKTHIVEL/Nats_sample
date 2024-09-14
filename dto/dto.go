package dto

type Request struct {
	Name  string `json:"name",omitempty`
	Data  []byte `json:"data",omitempty`
	Topic string `json:"topic",omitempty`
}
