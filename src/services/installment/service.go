package installment

import (
	"errors"
	"time"

	"github.com/alecanutto/isaac-structure/src/interfaces"
	"github.com/alecanutto/isaac-structure/src/repositories"
	"github.com/alecanutto/isaac-structure/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repos repositories.RepositoryContainer) InstallmentService {
	return InstallmentService{
		InstallmentRepository: repos.InstallmentRepository,
	}
}

func (is InstallmentService) Create(installment structs.Installments) error {

	if installment.DueDay < time.Now().Day() {
		return errors.New("Parcela vencida")
	}

	return is.InstallmentRepository.Create(installment)
}
