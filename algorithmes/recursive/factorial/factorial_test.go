package factorial

import (
	"errors"
	"fmt"
	"testing"
)

// Универсальный хелпер для запуска бенчмарка
func BenchmarkFactorialRecursive(b *testing.B) {
	nums := []int64{1, 5, 10, 15, 20}

	for _, num := range nums {
		b.Run(fmt.Sprintf("Size_%.2d", num), func(b *testing.B) {
			b.ResetTimer()

			for b.Loop() {
				FactorialRecursive(num)
			}
		})
	}
}

func TestFactorialRecursive(t *testing.T) {
	// Описываем тест-кейсы
	tests := []struct {
		name    string
		input   int64
		want    int64
		wantErr error
	}{
		{
			name:    "Факториал 0",
			input:   0,
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Факториал 1",
			input:   1,
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Обычное число (5!)",
			input:   5,
			want:    120,
			wantErr: nil,
		},
		{
			name:    "Максимально возможное число для int64 (20!)",
			input:   20,
			want:    2432902008176640000,
			wantErr: nil,
		},
		{
			name:    "Отрицательное число",
			input:   -5,
			want:    0,
			wantErr: ErrNegativNumber, // Проверяем конкретную ошибку
		},
		{
			name:    "Минимальное переполнение (21!)",
			input:   21,
			want:    0,
			wantErr: ErrOverflow,
		},
		{
			name:    "Сильное переполнение (100!)",
			input:   100,
			want:    0,
			wantErr: ErrOverflow,
		},
		// Такой тест не возможно провести с рекурсией
		// {
		// 	name:    "Максимальное значение int64",
		// 	input:   math.MaxInt64,
		// 	want:    0,
		// 	wantErr: ErrOverflow,
		// },
	}

	// Запускаем каждый тест-кейс
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FactorialRecursive(tt.input)

			// Проверяем ошибку с помощью errors.Is (лучшая практика Go)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("FactorialRecursive() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Проверяем результат
			if got != tt.want {
				t.Errorf("FactorialRecursive() got = %v, want %v", got, tt.want)
			}
		})
	}
}
