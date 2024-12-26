package sheet1_controller

type Sheet1Dto struct {
	
	File []byte `json:"file" binding:"required"`
	Sheet int `json:"sheet" binding:"required"`
}

