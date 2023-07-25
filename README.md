# DI-Practice

Learn how to use dependency injection and design pattern.

## How To Run 
1. Clone the repository:
```git
git clone https://github.com/Coolfreewerx/DI.git
```

2. Install PostgreSQL with Docker:
```bash
docker compose up -d
```

3. Run Program:
```bash
go run main.go 
```
or 
```
air 
```

## API Requests and Response

METHOD GET `http://localhost:1150/posts`

### Response

- Get From Database
```
[
    {
        "userId": 1,
        "id": 1,
        "title": "Name",
        "body": "Oscar"
    },
    {
        "userId": 1,
        "id": 2,
        "title": "Location",
        "body": "Australia"
    },
    ...
]

```

- Get From Jsonplaceholder

```
[
    {
        "userId": 1,
        "id": 1,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    },
    {
        "userId": 1,
        "id": 2,
        "title": "qui est esse",
        "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    },
    ...
]
```

