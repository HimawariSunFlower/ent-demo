package main

import (
	"context"
	"ent-demo/ent"
	"ent-demo/ent/pet"
	"ent-demo/ent/player"
	"ent-demo/ent/schema"
	"ent-demo/ent/user"
	"entgo.io/ent/entc/integration/ent/migrate"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	client := DB()
	defer client.Close()
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		u, err := client.User.Create().
			SetUsername(name).
			SetAge(123).
			SetName(name).
			Save(ctx)

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = w.Write([]byte(fmt.Sprintf("user: %v", u)))
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	mux.HandleFunc("/pet", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		user, err := client.User.Query().Where(user.Username("aaa")).First(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}

		p, err := client.Pet.Create().SetName(name).SetMaster(user).SetAge(123).Save(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte(p.String()))
	})

	mux.HandleFunc("/ans", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		u, err := client.User.Query().Where(user.Username(name)).First(ctx)

		if err != nil {
			fmt.Println(err)
			return
		}

		pets := client.Pet.Query().Where(pet.HasMasterWith(user.ID(u.ID))).AllX(ctx)
		pl := client.Player.Query().Where(player.HasUserWith(user.ID(u.ID))).AllX(ctx)

		w.Write([]byte(fmt.Sprintf("user: %v /n", u)))
		w.Write([]byte(fmt.Sprintf("pets: %v /n", pets)))
		w.Write([]byte(fmt.Sprintf("players: %v /n", pl)))
	})

	mux.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		eq := &schema.Equip{
			Id:    1,
			Attr:  "1:100|2:200",
			Attrs: map[int]int64{1: 100, 2: 200},
		}
		p, err := client.Player.Create().SetUsername("aaa").SetEquip(eq).SetUserID(1).Save(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte(p.String()))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func DB() *ent.Client {
	client, err := ent.Open("mysql", "aaa:123456@tcp(localhost:3306)/ent-demo?parseTime=True")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
