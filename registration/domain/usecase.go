package domain

type Usecase interface {
	AddRegister(request *AddRegisterRequest) (error, int, *Register)
	GetRegisters() (error, int, *[]Register)
}
