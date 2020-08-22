# Как запускать
```sh
git clone https://github.com/rsaddatimov/messenger-server.git
cd messenger-server
docker-compose build
docker-compose up
```
# Архитектура
### База данных
В качесиве хранилища данных выбрал PostgreSQL 10. БД поднимается как контейнер из под `docker-compose`, инициализируется табличками из файла `init.sql`:
```sql
CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(32),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Chats(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64),
    users INTEGER[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    lastUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Messages(
    id SERIAL PRIMARY KEY,
    chat INTEGER,
    author INTEGER,
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
У таблицы `Chats` есть ещё поле `lastUpdated`, который соответствует времени последнего обновления чата (посылка сообщения/создание чата). В случае падения/завершения работы сервера/контейнера данные БД не теряются, а сохраняются в папочке `./pgdata` (при следующем старте БД подтянет оттуда данные)
### Сервер
Сервер написал на Go. В качестве фреймфорка для работы с HTTP выбрал встроенный `net/http` модуль. Перед записью в БД входные данные проходят некоторые проверки - существование пользователя/чата или право на отправление сообщения пользователем в чат (должен состоять в этом чате) и т.д