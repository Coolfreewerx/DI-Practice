package progresql

import (
	"api/posts/ent"
	"context"
	"log"
	"os"

	// "github.com/joho/godotenv"
)

func InitDatabase() (*ent.Client, error) {
	ctx := context.Background()

	// init Database
	clientDB, err := ent.Open("postgres", os.Getenv("PSQL_DB_CONNECT"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil, err
	}

	// Run the auto migration tool.

	if err := clientDB.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return clientDB, err
	}

	// CreateUser & Post
	createUser(ctx, clientDB)
	createPost(ctx, clientDB)

	return clientDB, nil
}

func CloseDatabase(client *ent.Client) {
	client.Close()
}

func createUser(ctx context.Context, client *ent.Client) {
	if (client.User.Query().CountX(ctx) != 0 ) {
		return
	}
	err := client.User.Create().
				SetName("Dew").
				Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
func createPost(ctx context.Context, client *ent.Client) {
	if (client.Post.Query().CountX(ctx) != 0) {
		return
	}
	err := client.Post.Create().
				SetUserId(1).
				SetBody("est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla").
				SetTitle("sunt aut facere repellat provident occaecati excepturi optio reprehenderit").
				Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Post.Create().
				SetUserId(1).
				SetBody("quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto").
				SetTitle("qui est esse").
				Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}