package db

type Product struct {
	Name  string
	Price int
}

func Products() ([]Product, error) {
	data := []Product{
		{
			Name:  "Пепелац",
			Price: 5, // в Кэцэ
		},
	}
	return data, nil
}

func NewProduct(Product) error {
	return nil
}

func UpdateProduct(Product) error {
	return nil
}

func DeleteProduct(Product) error {
	return nil
}
