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

4. Call Swagger api:
```
curl localhost:1150/swagger/doc.json
```

5. Open SwaggerUI in browser by using URL:
```
http://localhost:1150/swagger/index.html
```

## API Requests and Response

### Request

- METHOD POST `http://localhost:1150/create-post`
```
{
    "userId" : 1,
    "title": "New Post",
    "body": "this is first post"
}
```

- METHOD GET `http://localhost:1150/posts`

### Response

- Post to create into database.

status : `201 Created`

```
{
    "userId": 1,
    "id": 1,
    "title": "New Post",
    "body": "This is first post"
}
```
- Get From Database or Jsonplaceholder.

status : `200 OK`

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
or 

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