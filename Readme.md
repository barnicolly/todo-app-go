# REST API приложение TODO List на Go

### Технологический стек:

- go 1.14
- postgres (в докере) + миграции БД
- фреймворк gin-gonic
- конфигурация из env (godotenv) и конфига (spf13/viper)
- работа с бд через sqlx

### Todo:

- [ ] Покрыть тестами
- [ ] Описание Swagger
- [ ] Завернуть go полностью в docker

### Список команд (makefile):

* `make up` - для поднятия окружения
* `make run-app ` - скомпилировать и запустить go пакет

Для выполнения миграций должен быть установлен golang migrate https://github.com/golang-migrate/migrate или выполнить напрямую к базе.