# Реализация приложения ToDoList, с применением SDUI на GOLANG

## Используемые библиотеки:
1) gorilla mux;
1) gorilla websocket.
1) gorm

Сначала необходимо запустить сервер:

`$ go run *.go`

_Для ленивых, можно запустить файл `sdui (linux, osx)`_

В папке `/client-side` есть страница `index.html`. Ее можно просто в браузере открыть.

Ради примера, будут показаны json-ответы, которые сервер присылает клиенту:

## Главный экран:

```
    {
		"components": [
			{
				"type" : "title",
				"content" : "My ToDo list App"
			},
			{
				"type" : "button",
				"content" : "Создать задачу",
				"link" : "createscreen"
			},
			
			{
				"type" : "element",
				"id" : "2",
				"title" : "Покупки",
				"description" : "Купить черную дыру",
				"checked" : "1"
			},
			{
				"type" : "element",
				"id" : "3",
				"title" : "Уборка",
				"description" : "Убрать в ванной",
				"checked" : "1"
			},
			{
				"type" : "footer"
			}
		]
	}
```

## Экран создания задачи:

```
    {
		"components" : [
			{
				"type" : "text",
				"content" : "Название задачи"
			},
			{
				"type" : "input",
				"name" : "title"
			},
			{
				"type" : "text",
				"content" : "Описание задачи"
			},
			{
				"type" : "input",
				"name" : "description"
			},
			{
				"type" : "submit",
				"content" : "Создать",
				"inputs" : ["title", "description"],
				"link" : "create"
			},
			{
				"type" : "button",
				"link" : "gettodos",
				"content" : "Выход"
			}

		]
	}
```