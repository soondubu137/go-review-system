package model

type Resolution struct {
	AppealID   string
	operatorID string
	Reason     string
	Content    string
	Pictures   []string
	Videos     []string
	Status     string
}
