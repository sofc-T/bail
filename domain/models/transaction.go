package models


type Transaction struct {
	file []byte 
	sheet int
	PaidIn float64
	Balance float64
	Withdrawal float64
	Agent string


}