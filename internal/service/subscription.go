package service

// Subscription request
type SubscriptionRequest struct {
	ObjType int   `json:"objType" binding:"required,one of 0 1 2"`
	ObjID   int64 `json:"objId" binding:"required"`
}
