package transaction_cmd

type Sheet1Command struct {
	file []byte 
	sheetname string 
}

// NewLoginCommand creates a new instance of LoginCommand.
func NewLoginCommand(email, password string) *Sheet1Command {
	return &Sheet1Command{
		email: email,
		password: password,
	}
}

