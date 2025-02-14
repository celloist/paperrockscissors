package models


type BaseUserInterface interface {
	IncrementScore()
	GetUsername() string
	GetScore() int
}

// props
// using lower case makes baseUser private
type baseUserModel struct {
	username string
	score    int
}

// use a pointer to our struct to modify the value and not create a new one
func (base *baseUserModel) IncrementScore() {
	base.score++
}

func (base baseUserModel) GetUsername() string {
	return base.username
}

func (base baseUserModel) GetScore() int {
	return base.score
}
