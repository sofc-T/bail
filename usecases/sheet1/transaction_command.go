package transaction_cmd

type Sheet1Command struct {
	file []byte 
	sheet int 
}

// NewLoginCommand creates a new instance of LoginCommand.
func NewLoginCommand(file []byte, sheet int) *Sheet1Command {
	return &Sheet1Command{
		file: file,
		sheet: sheet,
	}
}

