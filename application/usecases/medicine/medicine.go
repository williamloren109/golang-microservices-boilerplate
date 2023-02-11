package medicine

import (
	medicineDomain "github.com/gbrayhan/microservices-go/domain/medicine"
	medicineRepository "github.com/gbrayhan/microservices-go/infrastructure/repository/medicine"
)

type Service struct {
	MedicineRepository medicineRepository.Repository
}

func (s *Service) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	all, err := s.MedicineRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultMedicine{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

func (s *Service) GetById(id int) (*medicineDomain.Medicine, error) {
	return s.MedicineRepository.GetByID(id)
}

func (s *Service) Create(medicine *NewMedicine) (*medicineDomain.Medicine, error) {
	medicineModel := medicine.toDomainMapper()
	return s.MedicineRepository.Create(medicineModel)
}

func (s *Service) GetByMap(medicineMap map[string]interface{}) (*medicineDomain.Medicine, error) {
	return s.MedicineRepository.GetOneByMap(medicineMap)
}

func (s *Service) Delete(id int) error {
	return s.MedicineRepository.Delete(id)
}

func (s *Service) Update(id int, medicineMap map[string]interface{}) (*medicineDomain.Medicine, error) {
	return s.MedicineRepository.Update(id, medicineMap)
}
