package model

type Reply struct {
	ReviewID string
	UserID   string
	Content  string
	Pictures []string
	Videos   []string
}

type Appeal struct {
	ReviewID string
	UserID   string
	Reason   string
	Content  string
	Pictures []string
	Videos   []string
}
