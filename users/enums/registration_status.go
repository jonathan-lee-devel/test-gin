package enums

type RegistrationStatus int64

const (
	Success           RegistrationStatus = 0
	Failure           RegistrationStatus = 1
	UserAlreadyExists RegistrationStatus = 2
)
