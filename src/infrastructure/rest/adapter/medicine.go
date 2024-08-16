// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	medicineService "github.com/williamloren109/golang-microservices-boilerplate/src/application/usecases/medicine"
	medicineRepository "github.com/williamloren109/golang-microservices-boilerplate/src/infrastructure/repository/medicine"
	medicineController "github.com/williamloren109/golang-microservices-boilerplate/src/infrastructure/rest/controllers/medicine"
	"gorm.io/gorm"
)

// MedicineAdapter is a function that returns a medicine controller
func MedicineAdapter(db *gorm.DB) *medicineController.Controller {
	mRepository := medicineRepository.Repository{DB: db}
	service := medicineService.Service{MedicineRepository: mRepository}
	return &medicineController.Controller{MedicineService: service}
}
