package models

import "time"

type Request struct {
	CustomerName string  `json:"customerName"`
	Items        []Items `json:"items"`
}

type Response struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Items   `json:"items"`
}
