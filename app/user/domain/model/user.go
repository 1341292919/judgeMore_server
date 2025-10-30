package model

type User struct {
	Uid      int64
	UserName string
	Email    string
	Password string
	College  string
	Major    string
	Grade    string
	Status   int
	Role     string
}
type EmailAuth struct {
	Code  string
	Email string
	Uid   int64
	Time  int64 //时间戳
}
