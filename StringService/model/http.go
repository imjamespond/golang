package model

type Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type Request struct {
	Body string `json:"body"`
}
