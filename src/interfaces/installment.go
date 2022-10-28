package interfaces

import "github.com/alecanutto/isaac-structure/src/structs"

type InstallmentService interface {
	Create(installment structs.Installments) error
	GetLastItem() (structs.Installments, error)
}

type InstallmentRepository interface {
	Create(installment structs.Installments) error
	GetLastItem() (structs.Installments, error)
}
