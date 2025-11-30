package services

type SignUpCodeSvc interface {
	CreateAndSendSignUpCode(email string, password string) (err error)
	
}