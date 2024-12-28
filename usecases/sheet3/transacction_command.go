package transaction_cmd

type Sheet3Command struct {
	file []byte 
	sheet int
}

// NewLoginCommand creates a new instance of Sheet1 Command.
func NewSheet3Command(file []byte, sheet int) *Sheet3Command {
	return &Sheet3Command{
		file: file,
		sheet: sheet,
	}
}
