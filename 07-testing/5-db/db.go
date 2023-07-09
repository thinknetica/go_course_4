package db

type Product struct {
	Name  string
	Price int
}

func Products() ([]Product, error) {

	i := 4
	_ = i

	s := "ABC"
	_ = s

	p := Product{
		Name: "A",
	}
	_ = p

	var m map[string]bool = map[string]bool{"s": true}
	_ = m

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
