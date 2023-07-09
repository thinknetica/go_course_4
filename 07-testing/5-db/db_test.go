package db

import (
	"testing"
)

func TestProductCRUD(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	var item = Product{
		Name:  "Пепелац",
		Price: 5, // в Кэцэ
	}

	err := NewProduct(item)
	if err != nil {
		t.Error(err)
	}

	data, err := Products()
	if err != nil {
		t.Error(err)
	}
	found := false
	for _, p := range data {
		if p.Name == item.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("не найден продукт: %v", item)
	}

	err = UpdateProduct(item)
	if err != nil {
		t.Error(err)
	}
	data, err = Products()
	if err != nil {
		t.Error(err)
	}
	found = false
	for _, p := range data {
		if p.Name == item.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("не найден продукт: %v", item)
	}

	err = DeleteProduct(item)
	if err != nil {
		t.Error(err)
	}
	data, err = Products()
	if err != nil {
		t.Error(err)
	}
	found = false
	for _, p := range data {
		if p.Name == item.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("не удалён продукт: %v", item)
	}
}
