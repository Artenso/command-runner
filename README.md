# Command-Runner
Command-Runner - приложение для запуска bash скриптов.

## Запуск
Приложение запускается командой:
```
docker-compose up -d
```

## Описание
#### Приложение поддерживает gRPC и REST API через gRPC-Gateway
Для этого был написан proto файл со следующими ручками:
- `AddCommand` - добавляет и запускает команды. Команды передаются в тело запроса в виде строки.
- `GetCommand` - получает информацию о команде (текст команды, статус выполнения, вывод команды и ее pid в системе)
- `ListCommand`- получает информацию о списке команд (текст команды, статус выполнения и ее pid в системе). В запрос передаются query-параметры limit и offset для пагинации
- `StopCommand` - останавливает запущенные команды
#### Приложение работает с базой данных Postgres
Для взаимодействия были выбраны библиотеки goose, pgx, scanny и squirrel, как наиболее популярные и удобные.
Бд состоит из одной таблицы "commands"
```sql
CREATE TABLE commands (
    -- Уникальный идентификатор команды в бд
    id BIGSERIAL PRIMARY KEY ,
    -- Текст команды
    command TEXT NOT NULL DEFAULT '',
    -- Статус выполнения команды
    status status NOT NULL DEFAULT 'UNSPECIFIED',
    -- Идентификатор процесса в системе
    pid BIGINT,
    -- Вывод команды
    output TEXT
);
```
Для статуса выполнения команды создан отдельный тип
```sql
CREATE TYPE status AS ENUM ('UNSPECIFIED', 'NEW', 'IN_PROGRESS', 'DONE', 'FAILED', 'STOPPED');
```
#### Приложение контейнеризовано с помощью Docker и Docker-Compose
Написаны Docker файлы для приложения и для миграций бд, а также файл docker-compose для запуска всех контейнеров.
#### Логгирование
Написан логгер на базе uber zap, с возможностью устанавливать уровень логирования и указывать путь вывода логов.
#### Конфигурция
При запуске приложение читает файл конфигурации (config.json), где содержатся такие параметры как:
- gRPC и HTTP порты
- уровень логирования и путь к файлу с логами
- период обновления базы данных (сек)
- базовый размер буфера для записи вывода команд
#### Примеры использования
- Добавление команды: метод `POST`, адрес `localhost:8000/command/add`, запрос
  ```json
  {
    "command":"ping -c 2 ya.ru"
  }
  ```
  ответ
  ```json
  {
    "id": "1"
  }
  ```
- Получение информации о команде: метод `GET`, адрес `localhost:8000/command/1`, ответ до завершения выполнения команды
  ```json
  {
    "command": "ping -c 2 ya.ru",
    "status": "IN_PROGRESS",
    "output": "PING ya.ru (77.88.44.242): 56 data bytes\n64 bytes from 77.88.44.242: seq=0 ttl=55 time=13.089 ms",
    "pid": "15"
  }
  ```
  ответ после завершения выполнения команды
  ```json
  {
    "command": "ping -c 2 ya.ru",
    "status": "DONE",
    "output": "PING ya.ru (77.88.44.242): 56 data bytes\n64 bytes from 77.88.44.242: seq=0 ttl=55 time=13.089 ms\n64 bytes from 77.88.44.242: seq=1 ttl=55 time=13.154 ms\n\n--- ya.ru ping statistics ---\n2 packets transmitted, 2 packets received, 0% packet loss\nround-trip min/avg/max = 13.089/13.121/13.154 ms\n",
    "pid": "15"
  }
  ```
- Получение списка команд: метод `GET`, адрес `localhost:8000/command/list?limit=10&offset=0`, ответ
  ```json
  {
    "commands": [
        {
            "command": "ping -c 2 ya.ru",
            "status": "DONE",
            "pid": "14"
        },
        {
            "command": "ping -c 2 ya.ru",
            "status": "DONE",
            "pid": "15"
        },
        {
            "command": "ping -c 2 ya.ru",
            "status": "DONE",
            "pid": "16"
        },

    ]
  }
  ```
- Остановка выполнения команды: метод `PUT`, адрес `localhost:8000/command/2/stop`, ответ
  ```json
  {}
  ```
  если после этого вызвать получение команды ответ будет со статусом `STOPPED`
  ```json
  {
    "command": "ping -c 2 ya.ru",
    "status": "STOPPED",
    "output": "PING ya.ru (5.255.255.242): 56 data bytes\n64 bytes from 5.255.255.242: seq=0 ttl=53 time=14.114 ms",
    "pid": "18"
  }
  ```
