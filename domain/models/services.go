package models

import "time"


type ResetCode struct {
	Code int     // The reset code
	Expr time.Time // Expiration time of the reset code
}
