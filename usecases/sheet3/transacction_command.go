package transaction_cmd

import "time"

type Sheet3Command struct {
	file []byte 
	sheet int
	date time.Time
}

// NewLoginCommand creates a new instance of Sheet1 Command.
func NewSheet3Command(file []byte, sheet int, date time.Time) *Sheet3Command {
	return &Sheet3Command{
		file: file,
		sheet: sheet,
		date: date,
	}
}
