package controller

import (
	"flover-market/model"
	"fmt"
	"time"
)

type Product struct {
	ProductId    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	ImgSrc       string `json:"img_src"`
	Diameter     string `json:"diameter"`
	Height       int    `json:"height"`
	Price        int    `json:"price"`
	Count        int    `json:"count"`
}

type ProductFilter struct {
	CategoryId int `json:"category_id"`
	SortBy     int `json:"sort_by"`
}

func GetProduct(productId int) (Product, error) {
	productModel, err := model.GetProduct(productId)
	if err != nil {
		fmt.Println(err)
		return Product{}, err
	}
	price, err := model.GetLastPriceChange(productModel.ProductId)
	if err != nil {
		fmt.Println(err)
		return Product{}, err
	}

	product := convertProductFromModel(productModel)
	product.Price = price.NewPrice

	category, err := model.GetCategory(product.CategoryId)
	if err != nil {
		fmt.Println(err)
		return Product{}, err
	}
	product.CategoryName = category.CategoryName

	return product, nil
}

func AddProduct(product Product) (err error) {
	if err = validateProduct(&product); err != nil {
		fmt.Println(err)
		return err
	}
	product.ProductId, err = model.AddProduct(convertProductToModel(product))
	if err != nil {
		return err
	}
	err = model.AddPriceChange(model.PriceChange{
		ProductId:       product.ProductId,
		DatePriceChange: int(time.Now().Unix()),
		NewPrice:        product.Price,
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateProduct
func UpdateProduct(productId int, product Product) {
	// TODO
}

func DeleteProduct(productId int) {
	// TODO
}

func Filter(filter ProductFilter) []Product {
	// TODO
	return nil
}

func GetAllProducts() ([]Product, error) {
	productsModel, err := model.GetProducts()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	products := make([]Product, len(productsModel))
	for i := range productsModel {
		product := convertProductFromModel(productsModel[i])

		price, err := model.GetLastPriceChange(product.ProductId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		product.Price = price.NewPrice
		category, err := model.GetCategory(product.CategoryId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		product.CategoryName = category.CategoryName
		products[i] = product
	}
	return products, nil
}

func validateProduct(product *Product) error {

	if product.ProductName == "" {
		return fmt.Errorf("имя продукта не может быть пустым")
	}
	if product.Description == "" {
		return fmt.Errorf("описание продукта не может быть пустым")
	}
	//if product.CategoryId == 0 {
	//	return fmt.Errorf("категория продукта не может быть пустой")
	//}

	if product.ImgSrc == "" {
		return fmt.Errorf("ссылка на изображение продукта не может быть пустой")
	}
	//if product.Diameter == "" {
	//	return fmt.Errorf("диаметр продукта не может быть пустым")
	//}
	//if product.Height == 0 {
	//	return fmt.Errorf("высота продукта не может быть пустой")
	//}
	if product.Price == 0 {
		return fmt.Errorf("цена продукта не может быть пустой")
	}
	return nil
}

func convertProductToModel(product Product) model.Product {
	return model.Product{
		ProductId:   product.ProductId,
		ProductName: product.ProductName,
		CategoryId:  product.CategoryId,
		Description: product.Description,
		ImgSrc:      product.ImgSrc,
		Diameter:    product.Diameter,
		Height:      product.Height,
	}
}

func convertProductFromModel(product model.Product) Product {
	return Product{
		ProductId:   product.ProductId,
		ProductName: product.ProductName,
		CategoryId:  product.CategoryId,
		Description: product.Description,
		ImgSrc:      product.ImgSrc,
		Diameter:    product.Diameter,
		Height:      product.Height,
	}
}
