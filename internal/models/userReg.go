package models

type UserReg struct {
	FullName    *string
	DateOfBirth *string
	PhoneNumber *string
	ContactLink *string
	State       int
	StateReg    bool
	StateFlag   int
}
