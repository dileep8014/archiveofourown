package service

// Subsection request
type SubsectionRequest struct {
	Name      string `json:"name" binding:"maxlength=20"`
	Introduce string `json:"introduce" binding:"maxlength=200"`
}

// Subsection response
type SubsectionResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	Seq       int    `json:"seq"`
	WorkNum   int    `json:"workNum"`
}
