Репозиторий создан как попытка создать единый стиль написания тестов с использованием gomock
В пакете infrastructure/service имеетстя структура с зависимостью которую будем тестировать

Нпоминалка команды для кодогенерации моков:
mockgen -source=/c/Go_Homework/GoMock/pkg/interfaces/interfaces.go -destination=/c/Go_Homework/GoMock/pkg/interfaces/mockInterfaces/mockInterfaces.go