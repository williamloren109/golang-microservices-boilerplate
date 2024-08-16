// Package medicine provides the use case for medicine
package medicine

import (
	domainMedicine "github.com/williamloren109/golang-microservices-boilerplate/src/domain/medicine"
)

func (n *NewMedicine) toDomainMapper() *domainMedicine.Medicine {
	return &domainMedicine.Medicine{
		Name:        n.Name,
		Description: n.Description,
		EanCode:     n.EANCode,
		Laboratory:  n.Laboratory,
	}
}
