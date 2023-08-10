/*
    Схема учебной БД "Книги". Запрос в фомате MySQL.
*/

/*
    Удаляем таблицы, если они существуют.
    Удаление производится в обратном относительно создания порядке.
*/ 
DROP TABLE IF EXISTS books_authors;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS publishers;

/*
    Создаём таблицы БД.
    Сначала создаются таблицы, на которые ссылаются вторичные ключи.
*/
-- authors - писатели
CREATE TABLE authors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    first_name VARCHAR(20) NOT NULL DEFAULT '',
    last_name VARCHAR(20) NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

-- publishers - издатели
CREATE TABLE publishers (
    id SERIAL PRIMARY KEY, -- первичный ключ
    name VARCHAR(50) NOT NULL,
    website VARCHAR(100) NOT NULL DEFAULT ''
);

-- books - книги
CREATE TABLE books (
    id SERIAL PRIMARY KEY, -- первичный ключ
    title VARCHAR(50) NOT NULL, -- название
    year INTEGER DEFAULT 0, -- год выпуска (максимум текущий + 10)
    public_domain BOOLEAN DEFAULT FALSE, -- является ли общественным достоянием
    publisher_id INT DEFAULT 0 REFERENCES publishers(id) ON DELETE CASCADE ON UPDATE CASCADE 
);
-- индекс на базе бинарного дерева для быстрого поиска по названию книг
CREATE INDEX books_title_idx ON books(title) USING btree;

-- связь между книжками и писателями
-- (у одной книги может быть несколько авторов)
CREATE TABLE books_authors (
    id SERIAL PRIMARY KEY, -- первичный ключ
    book_id BIGINT NOT NULL REFERENCES books(id),
    author_id INTEGER NOT NULL REFERENCES authors(id),
    UNIQUE(book_id, author_id)
);

/*
    Первичное наполнение БД данными.
    Часто требуется для заполнения таблиц, на которые ссылаются вторичные ключи,
    имеющие значения по умолчанию.
*/

INSERT INTO authors (id, first_name) VALUES (0, 'автор не указан');
ALTER TABLE authors AUTO_INCREMENT = 100;

INSERT INTO publishers (id, name) VALUES (0, 'издательство не указано');
ALTER TABLE publishers AUTO_INCREMENT = 100;

INSERT INTO books (title) VALUES ('The Lord Of The Rings');
INSERT INTO books (title) VALUES ('1984');
ALTER TABLE books AUTO_INCREMENT = 100;