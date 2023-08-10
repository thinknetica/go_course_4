/*
    Первичное наполнение БД данными.
    Часто требуется для заполнения таблиц, на которые ссылаются вторичные ключи,
    имеющие значения по умолчанию.
*/

INSERT INTO authors (id, first_name) VALUES (0, 'автор не указан');
INSERT INTO authors (id, first_name) VALUES (1, 'автор неизвестен');
-- поскольку id установлен принудительно, то необходимо изменить начало послеовательности
ALTER SEQUENCE authors_id_seq RESTART WITH 100;

INSERT INTO publishers (id, name) VALUES (0, 'издательство не указано');
ALTER SEQUENCE publishers_id_seq RESTART WITH 100;

-- insert into books (title, year) VALUES ('Book', 2050); - что будет в результате выполнения?