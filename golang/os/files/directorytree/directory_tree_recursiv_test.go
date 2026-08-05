package directorytree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func BenchmarkDirectoryTreeRecursiv(b *testing.B) {
	// Будем тестировать функцию на разной глубине вложенности папок
	depths := []int{10, 50, 100}

	for _, depth := range depths {
		// Используем b.Run для создания под-бенчмарков для каждого уровня глубины
		b.Run(fmt.Sprintf("Depth-%d", depth), func(b *testing.B) {

			// 1. Создаем временную папку для текущего под-бенчмарка
			tmpDir := b.TempDir()

			// 2. Генерируем структуру папок заданной глубины
			// Передаем b вместо t, так как наш генератор принимает интерфейс testing.TB
			_ = generateDeepDirStructure(b, tmpDir, depth)

			// 3. Сбрасываем таймер бенчмарка, чтобы время генерации папок
			// на жестком диске/SSD не учитывалось в результатах теста функции
			b.ResetTimer()

			// 4. Основной цикл бенчмарка Go
			for i := 0; i < b.N; i++ {
				_, err := DirectoryTreeRecursiv(tmpDir)
				if err != nil {
					b.Fatalf("Ошибка во время бенчмарка: %v", err)
				}
			}
		})
	}
}

func TestDirectoryTreeRecursiv(t *testing.T) {
	// 1. Создаем изолированную временную папку
	tmpDir := t.TempDir()

	// 2. Генерируем структуру папок глубиной в 15 уровней
	expected := generateDeepDirStructure(t, tmpDir, 15)

	// 3. Вызываем тестируемую функцию
	actual, err := DirectoryTreeRecursiv(tmpDir)
	if err != nil {
		t.Fatalf("DirectoryTreeRecursiv вернула ошибку: %v", err)
	}

	// 4. Сортируем результаты для корректного сравнения путей
	sort.Strings(expected)
	sort.Strings(actual)

	// 5. Проверяем количество
	if len(actual) != len(expected) {
		t.Fatalf("Несовпадение количества элементов. Ожидалось: %d, Получено: %d", len(expected), len(actual))
	}

	// 6. Поэлементная проверка
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("На позиции %d\nОжидалось: %q\nПолучено:  %q", i, expected[i], actual[i])
		}
	}
}

// generateDeepDirStructure создает глубокую структуру папок для тестов и бенчмарков.
// Возвращает список всех абсолютно созданных путей (без файлов, так как функция падает на файлах).
func generateDeepDirStructure(t testing.TB, rootDir string, maxDepth int) []string {
	t.Helper() // Указывает go test, что эта функция является вспомогательной

	var createdPaths []string
	currentPath := rootDir

	for i := 1; i <= maxDepth; i++ {
		dirName := fmt.Sprintf("level%d", i)
		currentPath = filepath.Join(currentPath, dirName)
		createdPaths = append(createdPaths, currentPath)

		// Добавляем боковое ответвление на середине пути для проверки ветвления
		if i == maxDepth/2 {
			sidePath := filepath.Join(filepath.Dir(currentPath), "side_level")
			createdPaths = append(createdPaths, sidePath)

			if err := os.MkdirAll(sidePath, 0755); err != nil {
				t.Fatalf("Генератор: не удалось создать боковую ветку: %v", err)
			}
		}
	}

	// Физически создаем всю цепочку папок
	if err := os.MkdirAll(currentPath, 0755); err != nil {
		t.Fatalf("Генератор: не удалось создать глубокую структуру папок: %v", err)
	}

	return createdPaths
}
