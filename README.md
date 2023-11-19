<h1 align="center"> :smiley_cat: Cat Shelter CRM (backend) </h1>

<p align="center">
  <a href="https://github.com/Yu-Leo/bmstu-cat-shelter-crm-back/actions/workflows/go.yml" target="_blank"> <img alt="license" src="https://img.shields.io/github/actions/workflow/status/Yu-Leo/bmstu-cat-shelter-crm-back/go.yml?style=for-the-badge&logo=github"></a>
  <a href="https://goreportcard.com/report/github.com/Yu-Leo/bmstu-cat-shelter-crm-back" target="_blank"> <img alt="license" src="https://goreportcard.com/badge/github.com/Yu-Leo/bmstu-cat-shelter-crm-back?style=for-the-badge"></a>
  <a href="https://github.com/Yu-Leo/bmstu-cat-shelter-crm-back/commits/main" target="_blank"> <img alt="last commit" src="https://img.shields.io/github/last-commit/Yu-Leo/bmstu-cat-shelter-crm-back?style=for-the-badge"></a>
  <a href="https://github.com/Yu-Leo/bmstu-cat-shelter-crm-back/graphs/contributors" target="_blank"> <img alt="commit activity" src="https://img.shields.io/github/commit-activity/m/Yu-Leo/bmstu-cat-shelter-crm-back?style=for-the-badge"></a>
</p>

<p align="center"> Проект в рамках курсов "Программирование на основе Классов и Шаблонов" и "Парадигмы и Конструкции Языков Программирования" (МГТУ им. Н. Э. Баумана, ИУ5, 2 и 3 семестры) </p>
<hr>

## :handshake: Клиенты

- Desktop: https://github.com/Yu-Leo/bmstu-cat-shelter-crm-desktop
- Mobile: https://github.com/Yu-Leo/bmstu-cat-shelter-crm-mobile

## :world_map: Навигация

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

1.1 Запуск локально на машине
```bash 
make run
```

1.2 Запуск Docker-контейнере
```bash
make d-run
```

<a id="chapter-3"></a>

## :computer: Исходный код

Структура проекта основана на [go-clean-template](https://github.com/evrone/go-clean-template).

### Make-команды

- `make build` - сборка
- `make init-db` - инициализация файла с БД (SQLite3)
- `make run` - локальный запуск
- `make d-run` - запуск в Docker-контейнере
- `make lint` - запуск линтера
- `make swag-init` - обновление Swagger-документации
- `make mocks` - генерация моков для unit-тестов
- `make test` - запуск unit-тестов
- `make gotools` - установка вспомогательных инструментов (golangci-lint и mockery)

### Конфигурация

Структура настроек описана в  [`config/config.go`](./config/config.go).

Значения параметров задаются в [`config/config.yaml`](./config/config.yaml) и в переменных окружения.

### Линтер

В качестве линтера используется [golangci-lint](https://golangci-lint.run/). Конфиг задается в файле [./.golangci.yml](./.golangci.yml).

Линтер запускается при помощи команды:

```bash
make lint
```

### Unit-тесты

Для генерации моков используется [mockery](https://vektra.github.io/mockery/latest/).

Unit-тесты запускаются при помощи команды:

```bash
make test
```

P.S. Поскольку на данном этапе развития проекта в нём отсутсвет как таковая бизнес-логика, которую необходимо было бы покрыть unit-тестами,
написание необльшого кол-ва unit-тестов необходимо для выполнения учебных задач - изучения mockery и запуска тестов в GitHub Actions.

### E2E-тесты

End-to-end тесты написаны на Python. Располагаются в директории [`./e2e-tests`](./e2e-tests).

В пайплайне запускаются в Docker-контейнерах при помощи Docker Compose ([конфиг](./docker-compose.ci.yaml)).

### CI/CD

В качестве инструмента для CI/CD используется GitHub Actions.

Инструкции описаны в файле [`./.github/workflows/go.yml`](./.github/workflows/go.yml).

На каждый `push` в любой ветке запускается пайплайн, состоящий из следующих этапов:

1. Сборка проекта
2. Запуск линтера
3. Запуск тестов
   - Запуск unit-тестов
   - Запуск e2e-тестов

### Используемые технологии

- Язык программирования: **Go (1.20)**
- Фреймворки и библиотеки:
    - [`gin`](https://github.com/gin-gonic/gin) - HTTP веб-фреймворк
    - [`swag`](https://github.com/swaggo/swag) - автоматическая генерация RESTful API документации с Swagger 2.0
    - [`cleanenv`](http://github.com/ilyakaznacheev/cleanenv) - минималистичный конфигуратор настроек
    - [`logrus`](http://github.com/sirupsen/logrus) - логгер
- СУБД: **SQLite3**
- Инструменты
    - **Docker** и **Docker compose**
    - **make**
    - **GitHub Actions**
 
<a id="chapter-4"></a>

## :smile: Авторы

- [Ювенский Лев](https://github.com/Yu-Leo)
- [Беспалова Виктория](https://github.com/victobes)
