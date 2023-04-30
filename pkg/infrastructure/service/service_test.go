package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"project/pkg/models"
	mocked "project/pkg/repository/mocked"
	"testing"
)



// Состоит из: названия теста,  всех принимаемых значений ,
// всех возвращаемых значений и непосредственно самого тестируемого обьекта

type TestCase struct {
	NameOfTest    string     // в это поле предлагаю копировать название функции создающий TestCase
	InputPassword int				// одно из принимаемых значений тестируемого метода
	ExpectedRes   []models.Order	// один из ожидаемых результатов
	ExpectError bool 				// переменная говорящая тесту ожидается ли в нём ошибка
	ConcreteError error				// поле содержащее конкретную ошибку
	ObjectToTest  *Service			// поле содержащее непосредственно тестируемый объект
}


// Задача: затестить гет ордерс структуры сервис! - произведём функцией ниже итерирующейся по TestCase ам
func TestService_GetOrders(t *testing.T) {

	// Создаём контроллер
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	// Создаём массив тестов

	// Массив предполагает длину равную сумме всех return тестируемого метода + 1 на тест одного успешного выполнениния
	// Именно столько функций генерирующих тест кейзы и должно быть составлено В ТОМ СЛУЧАЕ ЕСЛИ ТЕСТИРУЕМЫЙ МЕТОД ВОЗРАЩАЕТ
	// РАЗУМНОЕ КОЛИЧЕСТВО КОНКРЕТНЫХ ОШИБОК (если нет неконкретные ошибки тестируем общим методом)

	tests := make([]*TestCase,0)

	tests = append(tests, GetOrdersPasswordError(ctl))
	tests = append(tests, GetOrdersLenRepoError(ctl))
	tests = append(tests, GetOrdersCommonError(ctl))
	tests = append(tests, GetOrdersSuccessful(ctl))

	// В цикле добавляем в tests наши TestCase


	for _, tc := range tests {
		t.Run(tc.NameOfTest, func(t *testing.T) {
			gotRes, err := tc.ObjectToTest.GetOrders(tc.InputPassword)
			// конструкция с if else для отличия кейза для общей или для конкретной ошибки..
			if tc.ExpectError{
				if tc.ConcreteError != nil{
					assert.Equal(t, tc.ConcreteError, err)
				}else{
					assert.NotEqual(t, nil, err)
				}
			}else{
				assert.Equal(t, nil, err )
			}

			assert.Equal(t, tc.ExpectedRes, gotRes)

		})
	}
}

// В функции подготовки TestCase задача:  создать и сконфигурировать обьект для тестирования хранить
// передваемые и ожидаемые значения

func GetOrdersPasswordError(ctl *gomock.Controller) *TestCase {  // Функция готовящая кейз ориетированный на конкретную ошибку

	// 1. подготовка , конфигурируем моки и реальные обьекты
	repoMock := mocked.NewMockInterfaceRepository(ctl)
	result := []models.Order{{Id: 777, Name: "Cool", Price: 500, CreatedAt: 466464},{Id: 777, Name: "Cool", Price: 500, CreatedAt: 466464}}
	// Мок конфигурирую всегда даже если предполагаю что до него не дойдёт..
	repoMock.EXPECT().GetOrders().Return(result, nil).Times(0)
	// 2. создаём обьект для тестирования
	service := New("service", repoMock)

	// 3. заполняем структуру значениями
	return &TestCase{
		NameOfTest:    "GetOrdersPasswordError",
		InputPassword: 321,
		ExpectedRes: nil ,
		ExpectError: true,
		ConcreteError: errors.New("wrong password"),
		ObjectToTest:  service,
	}
}

// Далее по коду метода GetOrders идёт обращение к репозиторию
// но так как нет понимания какая кокретно ошибка вернётся
// этот возврат будет проверен функцией GetOrdersCommonError
// А сайчас напишу тест к следующей возвращаемой конкретной ошибке:

func GetOrdersLenRepoError(ctl *gomock.Controller) *TestCase {  // Функция готовящая кейз ориетированный на конкретную ошибку

	// 1. подготовка , конфигурируем моки и реальные обьекты
	repoMock := mocked.NewMockInterfaceRepository(ctl)
	result := make([]models.Order,0)
	// Мок конфигурирую всегда даже если предполагаю что до него не дойдёт..
	repoMock.EXPECT().GetOrders().Return(result, nil)
	// 2. создаём обьект для тестирования
	service := New("service", repoMock)

	// 3. заполняем структуру значениями
	return &TestCase{
		NameOfTest:    "GetOrdersLenRepoError",
		InputPassword: 123,
		ExpectedRes: nil ,
		ExpectError: true,
		ConcreteError: errors.New("len repository error"),
		ObjectToTest:  service,
	}
}

// Функция возвращающая TestCase для общего тестирования:
func GetOrdersCommonError(ctl *gomock.Controller) * TestCase{
	// 1.
	repoMock := mocked.NewMockInterfaceRepository(ctl)
	repoMock.EXPECT().GetOrders().Return(nil, errors.New("any mystery error"))

	// 2.
	service := New("service", repoMock)

	// 3.
	return &TestCase{
		NameOfTest:    "GetOrdersCommonError",
		InputPassword: 123,
		ExpectedRes:   nil,
		ExpectError: true,
		ObjectToTest:  service,
	}

}

// Ну и конечно кейз с успешным проходом по телу метода:

func GetOrdersSuccessful(ctl *gomock.Controller) * TestCase {
	// 1. подготовка , конфигурируем моки и реальные обьекты
	repoMock := mocked.NewMockInterfaceRepository(ctl)
	result := []models.Order{{Id: 777, Name: "Cool", Price: 500, CreatedAt: 466464},{Id: 777, Name: "Cool", Price: 500, CreatedAt: 466464}}
	// Мок конфигурирую всегда даже если предполагаю что до него не дойдёт..
	repoMock.EXPECT().GetOrders().Return(result, nil)
	// 2. создаём обьект для тестирования
	service := New("service", repoMock)

	// 3. заполняем структуру значениями
	return &TestCase{
		NameOfTest:    "GetOrdersSuccessful",
		InputPassword: 123,   // в конкретном случае правильный пароль
		ExpectedRes: result ,
		ExpectError: false,
		ConcreteError: nil,
		ObjectToTest:  service,
	}
}

