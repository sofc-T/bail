package sheet3_controller

type sheet3Dto struct {
	
	File []byte `json:"file" binding:"required"`
	Sheet int `json:"sheet" binding:"required"`
}

