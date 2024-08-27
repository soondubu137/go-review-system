package model

type Resolution struct {
	AppealID   string
	OperatorID string
	Reason     string
	Content    string
	Pictures   []string
	Videos     []string
	Status     string
}
