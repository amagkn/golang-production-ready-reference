package tests_test

import "testing"

// Функция, которую мы будем тестировать
func Add(a, b int) int {
	return a + b
}

// Тестовая функция должна начинаться с префикса "Test"
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; ожидается %d", result, expected)
	}
}

// Табличный тест
func TestAddTable(t *testing.T) {
	// Таблица с тест кейсами
	var tests = []struct {
		FirstArgument  int
		SecondArgument int
		Expected       int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{1, -1, 0},
		{-1, -1, -2},
		{1, 1, 2},
		{100, 100, 200},
		{1_000, 1_000, 2_000},
		{10_000, 10_000, 20_000},
		{100_000, 10_0000, 200_000},
	}

	// Перебор всех тест кейсов
	for _, tt := range tests {
		result := Add(tt.FirstArgument, tt.SecondArgument)

		// Сравнение результата с ожидаемым значением
		if result != tt.Expected {
			t.Errorf("Add(%d, %d) = %d; а ожидается %d",
				tt.FirstArgument,
				tt.SecondArgument,
				tt.Expected,
				result,
			)
		}
	}
}
