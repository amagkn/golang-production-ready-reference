package tests_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест текущей попки
// go test .

// Тест текущей и вложеных
// go test ./...

// -count=1 - отключить кэширование
// go test -count=1 ./...

// -v - выводит подробную информацию (verbose)
// go test -v ./...

// -run - запуск конкретного теста
// go test ./... -run TestSomething

// Пакет Assert не останавливается при ошибке
func TestAssert(t *testing.T) {
	// Проверка на равенство
	assert.Equal(t, 123, 123)

	// Проверка на неравенство
	assert.NotEqual(t, 123, 456)

	// Проверка на Nil
	var err = errors.New("my-error")
	assert.NotNil(t, err)

	// Когда должен быть Nil
	var object *struct{}
	assert.Nil(t, object)
}

// Пакет Require останавливается при первой ошибке
func TestRequire(t *testing.T) {
	// Проверка на равенство
	require.Equal(t, 123, 123)

	// Проверка на неравенство
	require.NotEqual(t, 123, 456)

	// Проверка на Nil
	var err = errors.New("my-error")
	require.NotNil(t, err)

	// Когда должен быть Nil
	var object *struct{}
	require.Nil(t, object)
}
