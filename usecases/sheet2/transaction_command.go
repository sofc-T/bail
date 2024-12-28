package transaction_cmd

type Sheet2Command struct {
	file []byte 
	sheet int
}

// NewLoginCommand creates a new instance of Sheet2 Command.
func NewSheet2Command(file []byte, sheet int) *Sheet2Command {
	return &Sheet2Command{
		file: file,
		sheet: sheet,
	}
}
