package domain

type Core struct {
	ID          uint
	IdUser      uint
	IdProduct   uint
	ProductQty  uint
	ProductName string
	Price       int
	Image       string
}

type Repository interface {
	Insert(NewCart Core) (Core, error)
	Get(id uint) ([]Core, error)
	Update(NewCart Core) (Core, error)
	Delete(id uint) error
}

type Service interface {
	AddCart(NewCart Core) (Core, error)
	GetMyCart(id uint) ([]Core, error)
	UpdateQty(NewCart Core) (Core, error)
	DeleteonCart(id uint) error
}
