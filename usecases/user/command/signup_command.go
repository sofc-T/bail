package usercmd



// SignUpCommand represents the command to register a new user with necessary details.
type SignUpCommand struct {
	name                string
	email               string
	salary              float64
	age                 int
	role                string
	coSignerName        string
	codeNumber          string
	coSignerDocument    []byte 
	educationalDocument []byte
	password            string
}

// NewSignUpCommand creates a new SignUpCommand instance with the provided user details.
func NewSignUpCommand(name string, email string, salary float64, age int, role string, coSignerName string, codeNumber string, coSignerDocument []byte, educationalDocument []byte,password string) *SignUpCommand {
	return &SignUpCommand{
		name:                name,
		email:               email,
		salary:              salary,
		age:                 age,
		role:                role,
		coSignerName:        coSignerName,
		codeNumber:          codeNumber,
		coSignerDocument:    coSignerDocument,
		educationalDocument: educationalDocument,
		password:            password,
	}
}
