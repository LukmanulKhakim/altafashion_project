package repository

import (
	"altafashion_be/features/product/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Image       string
	Name        string
	Description string
	Category    string
	Qty         int
	Price       int
	UserID      uint
}

func FromDomain(dc domain.Core) Product {
	return Product{
		Model:       gorm.Model{ID: dc.ID},
		Image:       dc.Image,
		Name:        dc.Name,
		Description: dc.Description,
		Category:    dc.Category,
		Qty:         dc.Qty,
		Price:       dc.Price,
		UserID:      dc.UserID,
	}
}

func ToDomain(p Product) domain.Core {
	return domain.Core{
		ID:          p.ID,
		Image:       p.Image,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Qty:         p.Qty,
		Price:       p.Price,
		UserID:      p.UserID,
	}
}

func ToDomainArray(listProduct []Product) []domain.Core {
	var res []domain.Core
	for _, val := range listProduct {
		res = append(res, domain.Core{
			ID:          val.ID,
			Image:       val.Image,
			Name:        val.Name,
			Description: val.Description,
			Category:    val.Category,
			Qty:         val.Qty,
			Price:       val.Price,
			UserID:      val.UserID,
		})
	}

	return res
}
