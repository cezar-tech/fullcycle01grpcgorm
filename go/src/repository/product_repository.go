package repository

import (
	"context"
	"fmt"

	"log"

	fullcycle01grpcgorm "github.com/cezar-tech/fullcycle01grpcgorm/go/src/proto"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

type ProductDAO struct {
	ID          int32 `gorm:"primary_key"`
	Name        string
	Description string
	Price       float32
}

func (p *ProductRepository) CreateProduct(ctx context.Context, in *fullcycle01grpcgorm.CreateProductRequest) (*fullcycle01grpcgorm.CreateProductResponse, error) {

	prod := ProductDAO{
		Price:       in.GetPrice(),
		Description: in.GetDescription(),
		Name:        in.GetName(),
	}

	err := p.Db.Save(&prod).Error
	if err != nil {
		fmt.Printf("Error to save data %v", err)
		return nil, err
	}

	return &fullcycle01grpcgorm.CreateProductResponse{
		Product: &fullcycle01grpcgorm.Product{
			Price:       prod.Price,
			Description: prod.Description,
			Name:        prod.Name,
			Id:          prod.ID,
		},
	}, nil
}

func (p *ProductRepository) FindProducts(ctx context.Context, in *fullcycle01grpcgorm.FindProductsRequest) (*fullcycle01grpcgorm.FindProductsResponse, error) {
	var results []ProductDAO
	res := p.Db.Find(&results, "")
	err := res.Error

	if err != nil {
		log.Printf("Error to find products %v", err)
		return nil, err
	}

	returnVal := make([]*fullcycle01grpcgorm.Product, 0, len(results))
	for _, each := range results {
		returnVal = append(returnVal, &fullcycle01grpcgorm.Product{
			Id:          each.ID,
			Description: each.Description,
			Name:        each.Name,
			Price:       each.Price,
		})
	}
	return &fullcycle01grpcgorm.FindProductsResponse{Products: returnVal}, nil
}
