## Модули ##

Модули традиционно вызывают много вопросов у слушателей курса. Авторы Go написали серию статей по модулям: https://go.dev/blog/using-go-modules, однако в этой заметке я постараюсь привести базовую инструкцию по правильному использованию модулей во время обучения.  

Модуль - это проект. Для нас проектом является весь курс, поэтому в своих репозиториях мы должны создать только один модуль в корне каталога нашего репозитория, там же где и каталог `.git`. 

Ваша структура каталогов домашних заданий должна примерно повторять структуру этого репозитория. То есть требуется создать корневой каталог (например `go-core-4`), в нём создать модуль с помощью, например, `go mod init go-core-4`. Внутри этого каталога нужно создавать отдельные каталоги для каждого урока.
Пример структуры каталогов:  
```
├───00-docs
├───01-intro
│   ├───demoapp
│   │   ├───cmd
│   │   │   └───app
│   │   └───pkg
│   │       └───stringutils
│   └───hello
├───02-syntax
│
│   go.mod
│   README.md
```
При создании модуля создаётся файл go.mod, в котором указано имя модуля, который мы ввели как аргумент команды `go mod init`. Например, `go.mod` этого репозитория начинается с: `module go-core-4`.  

Название модуля является псевдонимом корневого каталога проекта и именно с этого названия начинаются пути импорта пакетов, находящихся в каталогах внутри каталога модуля. После имени модуля для каждого пакета идёт структура каталогов для данного пакета.  
Например, в первом занятии у нас есть путь импорта: `"go-core-4/01-intro/demoapp/pkg/stringutils"`. Этот путь начинается с имени модуля, после чего идет путь до каталога с пакетом. То есть этот пакет находится в каталоге `/01-intro/demoapp/pkg/stringutils`
