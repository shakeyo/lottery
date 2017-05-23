package model

type IMGroup struct {
	ID          string
	Name        string
	Description string
	Avatar      string
	Owner       int
	Members     []int
	Admins      []int
	Settings    int
}
