package database

import (
	"github.com/laioncorcino/go-product-user/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) ProductQuery {
	return &ProductDB{
		DB: db,
	}
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.
			Limit(limit).
			Offset((page - 1) * limit).
			Order("created_at " + sort).
			Find(&products).
			Error

	} else {
		err = p.DB.
			Order("created_at " + sort).
			Find(&products).
			Error
	}

	return products, err
}

func (p *ProductDB) FindByID(productId string) (*entity.Product, error) {
	var product *entity.Product
	err := p.DB.First(&product, "product_id = ?", productId).Error
	return product, err
}

func (p *ProductDB) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ProductID)
	if err != nil {
		return err
	}

	err = p.DB.
		Model(&entity.Product{}).
		Where("product_id = ?", product.ProductID).
		Updates(product).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductDB) Delete(productId string) error {
	product, err := p.FindByID(productId)
	if err != nil {
		return err
	}

	return p.DB.Delete(product, "product_id = ?", productId).Error
}
