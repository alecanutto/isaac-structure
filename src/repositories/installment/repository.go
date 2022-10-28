package installment

import (
	"github.com/alecanutto/isaac-structure/src/structs"
	"gorm.io/gorm"
)

type InstallmentRepository struct {
	db *gorm.DB
}

func NewInstallmentRepository(db *gorm.DB) InstallmentRepository {
	return InstallmentRepository{
		db: db,
	}
}

func (ir InstallmentRepository) Create(installment structs.Installments) error {
	return ir.db.Create(&installment).Error
}

func (ir InstallmentRepository) GetLastItem() (structs.Installments, error) {
	var item structs.Installments

	err := ir.db.Model(&structs.Installments{}).Last(&item).Error

	return item, err
}
