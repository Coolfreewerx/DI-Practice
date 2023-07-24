# Dependency injection practice
learn how to use dependency injection and design pattern
## How to Run

1. Clone the Repository
```
git clone https://github.com/Coolfreewerx/DI-Practice.git
```
2. Change branch main to second_di_pattern
```
git checkout second_di_pattern
```
3. Copy .env.example file 
```
cp .env.example .env
```
4. Install PostgreSQL
```
docker-compose up -d
```
5. Run program
```
go run main.go
```
or
```
air
```

# API Requests and Response

Get Posts data

METHOD GET `http://localhost:4040/api/posts`

## Response

```
[
    {
        "userId": 1,
        "id": 1,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    },
    {
        "userId": 1,
        "id": 2,
        "title": "qui est esse",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    }
    ...
]
```