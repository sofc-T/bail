package sheet2_controller

type Sheet2Dto struct {
	
	File []byte `json:"file" binding:"required"`
	Sheet int `json:"sheet" binding:"required"`
}

