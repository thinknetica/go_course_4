/*
    Схема учебной БД "Книги".
    Имена таблиц во множественном числе.
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
    first_name TEXT NOT NULL DEFAULT '',
    last_name TEXT NOT NULL DEFAULT '',
    year_of_birth INTEGER NOT NULL DEFAULT 0
);

-- publishers - издатели
CREATE TABLE publishers (
    id SERIAL PRIMARY KEY, -- первичный ключ
    name TEXT NOT NULL,
    website TEXT NOT NULL DEFAULT ''
);

-- books - книги
CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    isbn BIGINT UNIQUE, -- ISBN
    title TEXT NOT NULL, -- название
    year INTEGER DEFAULT 0, -- год выпуска (максимум текущий + 10)
    public_domain BOOLEAN DEFAULT FALSE, -- является ли общественным достоянием
    publisher_id INTEGER REFERENCES publishers(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    price INTEGER DEFAULT 0 CHECK (price >= 0),
    genres TEXT[] DEFAULT '{"не указано"}', -- жанры
    info JSONB DEFAULT '{}' -- сведения: оглавление, описание и пр.
);
-- индекс на базе бинарного дерева для быстрого поиска по названию книг
CREATE INDEX IF NOT EXISTS books_title_idx ON books USING btree (lower(title));

-- связь между книжками и писателями
-- (у одной книги может быть несколько авторов)
CREATE TABLE books_authors (
    id BIGSERIAL PRIMARY KEY, -- первичный ключ
    book_id BIGINT NOT NULL REFERENCES books(id),
    author_id INTEGER NOT NULL REFERENCES authors(id),
    UNIQUE(book_id, author_id)
);

-- функция-триггер для проверки года выпуска книги
CREATE OR REPLACE FUNCTION check_book_year()
  RETURNS TRIGGER AS $$
BEGIN
    IF NEW.year < (SELECT (extract(year from current_date) - 10)) AND
        NEW.year > (SELECT (extract(year from current_date) + 10))
        THEN RETURN NEW;
        ELSE RAISE EXCEPTION 'Invalid book year'; --RETURN NULL;
    END IF;
END;
$$ LANGUAGE plpgsql;
-- регистрация тригера для таблицы
CREATE OR REPLACE TRIGGER check_book_year BEFORE INSERT OR UPDATE ON books 
FOR EACH ROW EXECUTE PROCEDURE check_book_year();