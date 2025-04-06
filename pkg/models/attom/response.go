package models

import "time"

type Response struct {
	Status   Status     `json:"status"`
	Property []Property `json:"property"`
}

type Status struct {
	Version          string    `json:"version"`
	Code             int       `json:"code"`
	Msg              string    `json:"msg"`
	Total            int       `json:"total"`
	Page             int       `json:"page"`
	Pagesize         int       `json:"pagesize"`
	ResponseDateTime time.Time `json:"responseDateTime"`
	TransactionID    string    `json:"transactionID"`
	AttomID          *int      `json:"attomId"`
}
