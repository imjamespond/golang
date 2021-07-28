package model

import "github.com/go-kit/kit/endpoint"

type Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type Request map[string]interface{}

type Middleware func(endpoint.Endpoint) endpoint.Endpoint
