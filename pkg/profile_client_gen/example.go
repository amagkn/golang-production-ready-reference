package profile_client_gen

import (
	"context"
	"fmt"
)

// Для запуска этого примера:
// - Переместить в корень проекта
// - Добавить импорт: . "github.com/amagkn/golang-production-ready-reference/pkg/profile_client_gen"
// - Переименовать функцию Example в main
// - Переименовать название пакета (в самом верху этого файла): package main
// - Запустить сервер: make up, make migrate-up, make run
// - Запустить клиента (этот пример): go run example.go

func Example() { //nolint:funlen
	profile, err := New(Config{Host: "localhost", Port: "8080"})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	id, err := profile.Create(ctx, "John", 25, "john@gmail.com", "+73003002020")
	if err != nil {
		panic(err)
	}

	p, err := profile.GetProfile(ctx, id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	var (
		name  = "John Doe"
		age   = 26
		email = "new-john@gmail.com"
		phone = "+73003004000"
	)

	err = profile.Update(ctx, id.String(), &name, &age, &email, &phone)
	if err != nil {
		panic(err)
	}

	p, err = profile.GetProfile(ctx, id.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(p.ID)
	fmt.Println(p.Age)
	fmt.Println(p.Name)
	fmt.Println(p.Contacts.Email)
	fmt.Println(p.Contacts.Phone)

	err = profile.Delete(ctx, id.String())
	if err != nil {
		panic(err)
	}

	_, err = profile.GetProfile(ctx, id.String())

	fmt.Println("Get request:", err)
}
