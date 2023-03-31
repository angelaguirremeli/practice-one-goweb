package products

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load product from a json file

type Product struct {
	Id 		  	int		 `json:"id"`
	Name 	  	string	 `json:"name"`
	Quantity  	int		 `json:"quantity"`
	CodeValue 	string	 `json:"code_value"`
	IsPublished bool	 `json:"is_published"`
	Expiration 	string	 `json:"expiration"`
	Price 		float64  `json:"price"`
}

func LoadFromJsonFile(fileName string) ([]Product, error ){

    jsonFile, err := os.Open(fileName)

    if err != nil {
		return nil, err
    }
	defer jsonFile.Close()

	byteData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
    }

	var products []Product

	json.Unmarshal(byteData, &products)

	return products, nil

}

func SearchProductById(id int, products []Product) *Product  {
	for i := range products {
		if products[i].Id == id {
			return &products[i]
		}
	}
	return nil
}

func SearchProductByPrice(price float64, products []Product) []Product {
	productsFiltered := make([]Product, 0)

	for i := range products {
		if products[i].Price > price {
			productsFiltered = append(productsFiltered, products[i])
		}
	}
	return productsFiltered
}