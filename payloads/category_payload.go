package payloads

type Category []struct {
	Code  []string `json:"c"`
	Score float32  `json:"s"`
}
