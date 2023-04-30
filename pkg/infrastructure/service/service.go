package service

import (
	"errors"
	"project/pkg/models"
	"project/pkg/repository"
)

// Напоминалка для команды генерирующей моки:
// mockgen -source=/c/Go_Homework/AnyProject/pkg/interfaces/interfaces.go -destination=/c/Go_Homework/AnyProject/pkg/interfaces/mockInterfaces/mockInterfaces.go

type Interface interface{
	GetOrders (password int)([]models.Order, error)
	GetOrderById (password, id int)(models.Order, error)
	SetOrder(password int, o models.Order) error
}


type Service struct{
	Name string
	Repo repository.InterfaceRepository
}

func New(name string, repo repository.InterfaceRepository)*Service{
	return &Service{
		Name : "MyService",
		Repo: repo,
	}
}


// Тестов столько сколько возможных источников ошибок + 1 на нормальное выполнение функции

func (s *Service) GetOrders (password int)([]models.Order, error){

	if password != 123{
		return nil, errors.New("wrong password") // здесь возврат конкретной ошибки его удобно теситровать
	}

	res, err := s.Repo.GetOrders()
	if err != nil{
		return nil, err  // Здесб предполагается возврат широкого круга ошибок из вызываемого сервиса.. протестируем общей функцией
	}
	if len(res) == 0{
		return nil, errors.New("len repository error") // здесь возврат конкретной ошибки его удобно теситровать
	}

	return res, nil
}

func (s *Service) GetOrderById (password, id int)(models.Order, error){
	return models.Order{}, nil
}


func (s *Service)SetOrder(password int, o models.Order) error{
	return nil
}


