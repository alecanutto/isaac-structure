package services

import (
	"github.com/alecanutto/isaac-structure/src/interfaces"
	"github.com/alecanutto/isaac-structure/src/repositories"
	"github.com/alecanutto/isaac-structure/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repos repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repos),
	}
}
