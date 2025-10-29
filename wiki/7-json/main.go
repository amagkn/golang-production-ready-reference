package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Email     *string           `json:"email,omitempty"` // Пропускается если nil (или поля не было в json)
	Roles     []string          `json:"roles"`
	CreatedAt time.Time         `json:"created_at"`
	Password  string            `json:"-"`              // Не попадает в JSON
	Metadata  map[string]string `json:"meta,omitempty"` // Пропускается если nil (или поля не было в json)
}

func main() {
	{ // 1) Маршалинг структуры в JSON
		user := User{
			ID:        1,
			Name:      "Alice",
			Email:     nil, // будет пропущено в JSON из-за "omitempty"
			Roles:     []string{"admin", "user"},
			CreatedAt: time.Now().UTC().Round(0),
			Password:  "super-secret", // будет исключён потому тег "-"
			Metadata:  map[string]string{"lang": "ru"},
		}

		b, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b)) // красиво отформатированный JSON
	}

	os.Exit(42)

	{ // 2) Анмаршалинг JSON в структуру
		input := `{"id":2,"name":"Bob","email":"bob@example.com","roles":["user"],"created_at":"2025-09-25T10:20:30Z","meta":{"city":"Moscow"}}`

		var user User

		err := json.Unmarshal([]byte(input), &user)
		if err != nil {
			panic(err)
		}

		fmt.Println(user)
	}

	{ // 3) Строгий анмаршалинг: запрет неизвестных полей
		strictJSON := `{"id":3,"name":"Eve","roles":[],"created_at":"2025-01-02T03:04:05Z","unknown":123}`
		dec := json.NewDecoder(strings.NewReader(strictJSON))
		dec.DisallowUnknownFields()

		var user User

		err := dec.Decode(&user)
		if err != nil {
			fmt.Println("Строгий анмаршалинг (ошибка из-за неизвестного поля):", err)
		} else {
			fmt.Println(user)
		}
	}
}
