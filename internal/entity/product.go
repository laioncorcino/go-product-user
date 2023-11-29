package entity

import (
	"errors"
	"github.com/laioncorcino/go-product-user/pkg"
	"time"
)

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ProductID string    `json:"product_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     int       `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewProduct(name string, price int) (*Product, error) {
	p := &Product{
		ProductID: pkg.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
