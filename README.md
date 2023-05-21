<h1 align="center"> :smiley_cat: Cat Shelter CRM (backend) </h1>

<p align="center"> Проект в рамках курса "Программирование на основе Классов и Шаблонов" (МГТУ им. Н. Э. Баумана, ИУ5, 2 семестр) </p>
<hr>

:heavy_exclamation_mark: **Репозиторий с кодом десктопного клиента: https://github.com/Yu-Leo/bmstu-cat-shelter-crm-desktop**

## Навигация

* [Описание проекта](#chapter-0)
* [API](#chapter-1)
* [Запуск](#chapter-2)
* [Исходный код](#chapter-3)
* [Авторы](#chapter-4)

<a id="chapter-0"></a>

## :page_facing_up: Описание проекта

CRM-система для управления внутренней деятельностью кошачьего приюта.

<a id="chapter-1"></a>

## :pushpin: API

OpenAPI спецификация:
- [`docs/swagger.json`](./docs/swagger.json)
- [`docs/swagger.yaml`](./docs/swagger.yaml)

Можно использовать: [визуализация файла OpenAPI спецификации](https://editor.swagger.io).

После запуска сервиса можно использовать Swagger UI: [`http://127.0.0.1:9000/swagger/index.html`](http://127.0.0.1:9000/swagger/index.html).

<a id="chapter-2"></a>

## :zap: Запуск
0. Инициализация БД
```bash
make init-db
```

1. Запуск локально на машине
```bash 
make run
```

2. Запуск Docker-контейнере
```bash
make d-run
```

<a id="chapter-3"></a>

## :computer: Исходный код

Структура проекта основана на [go-clean-template](https://github.com/evrone/go-clean-template).

### Конфигурация

Структура настроек описана в  [`config/config.go`](./config/config.go).

Значения параметров задаются в [`config/config.yaml`](./config/config.yaml) и в переменных окружения.

### Технологии

- СУБД: **SQLite3**
- Язык программирования: **Go (1.20)**
- Фреймворки и библиотеки:
    - [`gin`](https://github.com/gin-gonic/gin) - HTTP веб-фреймворк
    - [`swag`](https://github.com/swaggo/swag) - автоматическая генерация RESTful API документации с Swagger 2.0
    - [`cleanenv`](http://github.com/ilyakaznacheev/cleanenv) - минималистичный конфигуратор настроек
    - [`logrus`](http://github.com/sirupsen/logrus) - логгер
- Инструменты
    - **Docker**
    - **make**

<a id="chapter-4"></a>

## :smile: Авторы
- [Ювенский Лев](https://github.com/Yu-Leo)
- [Беспалова Виктория](https://github.com/victobes)
