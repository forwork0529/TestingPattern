package repository

import(
	"project/pkg/models"
)

type InterfaceRepository interface {
	GetOrders ()([]models.Order, error)
	GetOrderById (id int)models.Order
	SetOrder(o models.Order) error
}


type Repository struct{
	Name string
}

func New(name string)*Repository{
	return &Repository{
		Name : name,
	}
}

func (r *Repository) GetOrders ()([]models.Order, error){
	orders := make([]models.Order, 3)
	orders = append(orders, models.Order{Id: 777, Name: "Cool", Price: 500, CreatedAt: 466464})
	return orders, nil
}

func (r *Repository) GetOrderById (id int)models.Order{
		return models.Order{Id: 888, Name: "VeryCool", Price: 800, CreatedAt: 4768869}
}


func (r *Repository)SetOrder(o models.Order) error{
	return nil
}
