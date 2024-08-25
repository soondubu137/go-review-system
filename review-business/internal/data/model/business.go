package model

type Reply struct {
	ReviewID int64
	UserID   int64
	Content  string
	Pictures []string
	Videos   []string
}
