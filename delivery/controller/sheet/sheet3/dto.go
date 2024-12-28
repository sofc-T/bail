package sheet3_controller

import "time"

type sheet3Dto struct {
	
	File []byte `json:"file" binding:"required"`
	Sheet int `json:"sheet" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
}

