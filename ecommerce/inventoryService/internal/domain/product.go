package domain

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id int) (*Product, error)
	Update(product *Product) error
	Delete(id int) error
	List() ([]Product, error)
}

type ProductUsecase interface {
	Create(product *Product) error
	GetByID(id int) (*Product, error)
	Update(product *Product) error
	Delete(id int) error
	List() ([]Product, error)
}
