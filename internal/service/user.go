package service

// User response example
type UserResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Root      bool   `json:"root"`
	WorksNums int64  `json:"worksNums"`
	WorkDay   int64  `json:"workDay"`
	Words     int64  `json:"words"`
	FansNums  int64  `json:"fansNums"`
}
