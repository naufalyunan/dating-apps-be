package entities

type ActivityLog struct {
	ID            uint
	UserID        uint
	ActionType    string
	ActionDetails string
}
