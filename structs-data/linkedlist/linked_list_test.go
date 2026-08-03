package linkedlist

import (
	"testing"
)

func TestAppendAndSize(t *testing.T) {
	t.Run("Empty list state", func(t *testing.T) {
		list := New()

		if !list.IsEmpty() {
			t.Errorf("Ожидался пустой список")
		}
		if list.Size() != 0 {
			t.Errorf("У пустого списка размер должен быть 0, получили %d", list.Size())
		}
		if list.head != nil || list.tail != nil {
			t.Error("У пустого списка head и tail должны быть nil")
		}
	})

	t.Run("Append to empty list", func(t *testing.T) {
		list := New()
		list.Append(10)

		if list.Size() != 1 {
			t.Errorf("Ожидался размер 1, получили %d", list.Size())
		}
		if list.IsEmpty() {
			t.Errorf("Список не должен быть пустым")
		}
		// Важно: когда элемент один, head и tail смотрят на него вместе
		if list.head == nil || list.tail == nil || list.head != list.tail {
			t.Error("При одном элементе head и tail должны указывать на один и тот же узел")
		}
		if list.head.value != 10 {
			t.Errorf("Ожидалось значение головы 10, получили %d", list.head.value)
		}
	})

	t.Run("Multiple appends", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)
		list.Append(30)

		if list.Size() != 3 {
			t.Errorf("Ожидался размер 3, получили %d", list.Size())
		}
		if list.head.value != 10 {
			t.Errorf("Голова списка должна быть 10, получили %d", list.head.value)
		}
		if list.tail == nil || list.tail.value != 30 {
			t.Errorf("Хвост списка должен быть 30, получили %v", list.tail)
		}
		if list.tail.next != nil {
			t.Error("Указатель next у хвоста должен быть nil")
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("Prepend to empty list", func(t *testing.T) {
		list := New()
		list.Prepend(10)

		if list.Size() != 1 {
			t.Errorf("Ожидался размер 1, получили %d", list.Size())
		}
		// Как и в Append, один элемент инициализирует и голову, и хвост
		if list.head == nil || list.tail == nil || list.head != list.tail {
			t.Error("При первом Prepend head и tail должны указывать на один узел")
		}
	})

	t.Run("Multiple prepends order", func(t *testing.T) {
		list := New()
		list.Prepend(10)
		list.Prepend(20) // Теперь 20 должно быть в начале
		list.Prepend(30) // Теперь 30 должно быть в начале (30 -> 20 -> 10)

		if list.Size() != 3 {
			t.Errorf("Ожидался размер 3, получили %d", list.Size())
		}

		// Проверяем правильность головы и хвоста
		if list.head.value != 30 {
			t.Errorf("Голова списка должна быть 30, получили %d", list.head.value)
		}
		if list.tail == nil || list.tail.value != 10 {
			t.Errorf("Хвост списка должен остаться 10, получили %v", list.tail)
		}

		// Проверяем последовательность через Get
		expected := []int{30, 20, 10}
		for i, exp := range expected {
			val, err := list.Get(i)
			if err != nil {
				t.Fatalf("Ошибка при получении индекса %d: %v", i, err)
			}
			if val != exp {
				t.Errorf("На индексе %d ожидалось %d, получили %d", i, exp, val)
			}
		}
	})
}

func TestInsert(t *testing.T) {
	// Сценарий 1: Вставка в абсолютно пустой список (index 0)
	t.Run("Insert into empty list", func(t *testing.T) {
		list := New()
		err := list.Insert(0, 100)
		if err != nil {
			t.Fatalf("Не ожидали ошибку при вставке в пустой список: %v", err)
		}

		if list.Size() != 1 {
			t.Errorf("Ожидали размер 1, получили %d", list.Size())
		}

		val, _ := list.Get(0)
		if val != 100 {
			t.Errorf("Ожидали 100 на индексе 0, получили %d", val)
		}

		// Проверяем, что head и tail смотрят на один и тот же узел
		if list.head == nil || list.tail == nil || list.head != list.tail {
			t.Error("Head и Tail должны указывать на один и тот же единственный узел")
		}
	})

	// Сценарий 2: Вставка в начало (index 0) уже непустого списка
	t.Run("Insert at the beginning (index 0)", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)

		err := list.Insert(0, 5)
		if err != nil {
			t.Fatalf("Ошибка при вставке в начало: %v", err)
		}

		if list.Size() != 3 {
			t.Errorf("Ожидали размер 3, получили %d", list.Size())
		}

		val, _ := list.Get(0)
		if val != 5 {
			t.Errorf("Новый head должен быть 5, получили %d", val)
		}

		if list.head.value != 5 {
			t.Errorf("Поле head.value должно быть 5, получили %d", list.head.value)
		}
	})

	// Сценарий 3: Вставка в самый конец (index == size)
	t.Run("Insert at the end (index == size)", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)

		err := list.Insert(2, 30) // индекс равен l.size (2)
		if err != nil {
			t.Fatalf("Ошибка при вставке в конец: %v", err)
		}

		if list.Size() != 3 {
			t.Errorf("Ожидали размер 3, получили %d", list.Size())
		}

		val, _ := list.Get(2)
		if val != 30 {
			t.Errorf("Последний элемент должен быть 30, получили %d", val)
		}

		if list.tail == nil || list.tail.value != 30 {
			t.Errorf("Указатель tail должен обновиться на 30, получили %v", list.tail)
		}

		if list.tail.next != nil {
			t.Error("Поле next у хвоста должно быть nil")
		}
	})

	// Сценарий 4: Вставка в середину списка (ваш базовый сценарий)
	t.Run("Insert in the middle", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(30)

		err := list.Insert(1, 20)
		if err != nil {
			t.Fatalf("Ошибка при вставке в середину: %v", err)
		}

		if list.Size() != 3 {
			t.Errorf("Ожидали размер 3, получили %d", list.Size())
		}

		// Проверяем всю цепочку значений по индексам
		expected := []int{10, 20, 30}
		for i, exp := range expected {
			val, _ := list.Get(i)
			if val != exp {
				t.Errorf("На индексе %d ожидали %d, получили %d", i, exp, val)
			}
		}
	})

	// Сценарий 5: Проверка невалидных индексов (границы)
	t.Run("Insert out of bounds", func(t *testing.T) {
		list := New()
		list.Append(10) // size = 1

		// Отрицательный индекс
		if err := list.Insert(-1, 99); err == nil {
			t.Error("Ожидали ошибку для отрицательного индекса, но её нет")
		}

		// Слишком большой индекс
		if err := list.Insert(2, 99); err == nil {
			t.Error("Ожидали ошибку для индекса 2 (при размере 1), но её нет")
		}
	})
}

func TestGet(t *testing.T) {
	list := New()
	list.Append(10)

	val, err := list.Get(0)
	if err != nil || val != 10 {
		t.Errorf("Ошибка при получении элемента: %v", err)
	}

	_, err = list.Get(1)
	if err == nil {
		t.Errorf("Ожидалась ошибка выхода за границы списка")
	}
}

func TestFind(t *testing.T) {
	// Сценарий 1: Поиск в пустом списке
	t.Run("Find in empty list", func(t *testing.T) {
		list := New()

		index := list.Find(10)
		if index != -1 {
			t.Errorf("Ожидали -1 для пустого списка, получили %d", index)
		}
	})

	// Сценарий 2: Обычный поиск элементов
	t.Run("Find existing elements", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)
		list.Append(30)

		// Ищем первый элемент (head)
		if idx := list.Find(10); idx != 0 {
			t.Errorf("Ожидали индекс 0 для значения 10, получили %d", idx)
		}

		// Ищем элемент в середине
		if idx := list.Find(20); idx != 1 {
			t.Errorf("Ожидали индекс 1 для значения 20, получили %d", idx)
		}

		// Ищем последний элемент (tail)
		if idx := list.Find(30); idx != 2 {
			t.Errorf("Ожидали индекс 2 для значения 30, получили %d", idx)
		}
	})

	// Сценарий 3: Поиск элемента, которого нет в списке
	t.Run("Find non-existing element", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)

		if idx := list.Find(99); idx != -1 {
			t.Errorf("Ожидали -1 для отсутствующего значения 99, получили %d", idx)
		}
	})

	// Сценарий 4: Правило №4 — Поиск первого вхождения при наличии дубликатов
	t.Run("Find first occurrence with duplicates", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20) // Вот первое вхождение числа 20 (индекс 1)
		list.Append(30)
		list.Append(20) // Вот второе вхождение числа 20 (индекс 3)

		if idx := list.Find(20); idx != 1 {
			t.Errorf("По правилу №4 ожидали индекс ПЕРВОГО вхождения (1), но получили %d", idx)
		}
	})
}

func TestRemove(t *testing.T) {
	// Сценарий 1: Удаление из пустого списка
	t.Run("Remove from empty list", func(t *testing.T) {
		list := New()

		removed := list.Remove(10)
		if removed {
			t.Error("Ожидали false при удалении из пустого списка")
		}
		if list.Size() != 0 {
			t.Errorf("Размер должен остаться 0, получили %d", list.Size())
		}
	})

	// Сценарий 2: Удаление единственного элемента в списке
	t.Run("Remove single element", func(t *testing.T) {
		list := New()
		list.Append(10)

		removed := list.Remove(10)
		if !removed {
			t.Error("Ожидали true при удалении существующего элемента")
		}
		if list.Size() != 0 {
			t.Errorf("После удаления единственного элемента размер должен быть 0, получили %d", list.Size())
		}
		// Критично: head и tail должны очиститься!
		if list.head != nil || list.tail != nil {
			t.Error("После удаления единственного элемента head и tail должны стать nil")
		}
	})

	// Сценарий 3: Удаление первого элемента (Головы / Head)
	t.Run("Remove head element", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)
		list.Append(30)

		removed := list.Remove(10)
		if !removed {
			t.Error("Ожидали true при удалении головы")
		}
		if list.Size() != 2 {
			t.Errorf("Ожидали размер 2, получили %d", list.Size())
		}
		// Проверяем, что голова сместилась на следующий узел
		if list.head == nil || list.head.value != 20 {
			t.Errorf("Новая голова должна быть 20, получили %v", list.head)
		}
	})

	// Сценарий 4: Удаление последнего элемента (Хвоста / Tail)
	t.Run("Remove tail element", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)
		list.Append(30)

		removed := list.Remove(30)
		if !removed {
			t.Error("Ожидали true при удалении хвоста")
		}
		if list.Size() != 2 {
			t.Errorf("Ожидали размер 2, получили %d", list.Size())
		}
		// Критично: tail должен сместиться назад на число 20, а его next стать nil!
		if list.tail == nil || list.tail.value != 20 {
			t.Errorf("Новый хвост должен быть 20, получили %v", list.tail)
		}
		if list.tail.next != nil {
			t.Error("У нового хвоста поле next должно быть nil")
		}
	})

	// Сценарий 5: Удаление элемента из середины
	t.Run("Remove middle element", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)
		list.Append(30)

		removed := list.Remove(20)
		if !removed {
			t.Error("Ожидали true при удалении элемента из середины")
		}
		if list.Size() != 2 {
			t.Errorf("Ожидали размер 2, получили %d", list.Size())
		}

		// Проверяем, что связи восстановились корректно (10 -> 30)
		val0, _ := list.Get(0)
		val1, _ := list.Get(1)
		if val0 != 10 || val1 != 30 {
			t.Errorf("Ожидали структуру списка 10 -> 30, получили %d -> %d", val0, val1)
		}
	})

	// Сценарий 6: Проверка правила №4 (Удаление только ПЕРВОГО вхождения дубликата)
	t.Run("Remove first occurrence of duplicate", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20) // Первое вхождение (индекс 1) - должно быть удалено
		list.Append(30)
		list.Append(20) // Второе вхождение (индекс 3) - должно остаться

		removed := list.Remove(20)
		if !removed {
			t.Error("Ожидали true")
		}
		if list.Size() != 3 {
			t.Errorf("Ожидали размер 3, получили %d", list.Size())
		}

		// Проверяем, что порядок теперь 10 -> 30 -> 20
		expected := []int{10, 30, 20}
		for i, exp := range expected {
			val, _ := list.Get(i)
			if val != exp {
				t.Errorf("На индексе %d ожидали %d, получили %d", i, exp, val)
			}
		}
	})

	// Сценарий 7: Удаление несуществующего элемента
	t.Run("Remove non-existing element", func(t *testing.T) {
		list := New()
		list.Append(10)
		list.Append(20)

		removed := list.Remove(99)
		if removed {
			t.Error("Ожидали false при удалении отсутствующего значения")
		}
		if list.Size() != 2 {
			t.Errorf("Размер не должен был измениться, получили %d", list.Size())
		}
	})
}

func TestRemoveAt(t *testing.T) {
	list := New()
	list.Append(10)
	list.Append(20)

	val, err := list.RemoveAt(0)
	if err != nil || val != 10 {
		t.Errorf("Ожидалось удаление элемента 10, получили %v (ошибка: %v)", val, err)
	}

	if list.Size() != 1 {
		t.Errorf("Ожидался размер 1, получили %d", list.Size())
	}
}

func TestContainsAndFind(t *testing.T) {
	list := New()
	list.Append(10)
	list.Append(20)

	if !list.Contains(20) {
		t.Errorf("Список должен содержать 20")
	}

	idx := list.Find(20)
	if idx != 1 {
		t.Errorf("Ожидался индекс 1 для значения 20, получили %d", idx)
	}

	idx = list.Find(99)
	if idx != -1 {
		t.Errorf("Для несуществующего элемента ожидался индекс -1, получили %d", idx)
	}
}

func TestClear(t *testing.T) {
	list := New()
	list.Append(10)
	list.Append(20)

	list.Clear()

	if !list.IsEmpty() {
		t.Errorf("Список должен быть пуст после Clear")
	}
}

// Тест для метода MoveToFront
func TestMoveToFront(t *testing.T) {
	t.Run("Перемещение из середины в начало", func(t *testing.T) {
		l := New()
		l.Append(10)
		l.Append(20) // Этот узел будем перемещать
		l.Append(30)

		n2 := findNodeByValue(t, l, 20)
		if n2 == nil {
			t.Fatal("Тестовый узел со значением 20 не найден в списке")
		}

		l.MoveToFront(n2)

		if l.head != n2 {
			t.Errorf("Ожидалась новая голова со значением %v, получили %v", 20, l.head.value)
		}
		if l.head.next == nil || l.head.next.value != 10 {
			t.Error("После новой головы должен идти элемент 10")
		}
		if l.tail == nil || l.tail.value != 30 {
			t.Error("Хвост списка не должен был измениться и должен быть равен 30")
		}
	})

	t.Run("Перемещение хвоста в начало", func(t *testing.T) {
		l := New()
		l.Append(10)
		l.Append(20)
		l.Append(30) // Хвост, который будем перемещать

		n3 := findNodeByValue(t, l, 30)

		l.MoveToFront(n3)

		if l.head != n3 {
			t.Errorf("Хвост не стал головой")
		}
		if l.tail == nil || l.tail.value != 20 {
			t.Errorf("Новым хвостом должен был стать элемент 20")
		}
		if l.tail.next != nil {
			t.Error("У нового хвоста next должен быть nil")
		}
	})

	t.Run("Перемещение узла, который уже в начале", func(t *testing.T) {
		l := New()
		l.Append(10)
		l.Append(20)

		n1 := findNodeByValue(t, l, 10)

		l.MoveToFront(n1)

		if l.head != n1 {
			t.Error("Голова списка изменилась, хотя не должна была")
		}
	})
}

// Тест для метода MoveToBack
func TestMoveToBack(t *testing.T) {
	t.Run("Перемещение из середины в конец", func(t *testing.T) {
		l := New()
		l.Append(10)
		l.Append(20) // Этот узел будем перемещать
		l.Append(30)

		n2 := findNodeByValue(t, l, 20)

		l.MoveToBack(n2)

		if l.tail != n2 {
			t.Errorf("Ожидался новый хвост со значением %v, получили %v", 20, l.tail.value)
		}
		if l.head == nil || l.head.value != 10 {
			t.Error("Голова списка не должна была измениться и должна быть равна 10")
		}
		if n2.next != nil {
			t.Error("У нового хвоста next должен быть nil")
		}
	})

	t.Run("Перемещение головы в конец", func(t *testing.T) {
		l := New()
		l.Append(10) // Голова, которую будем перемещать
		l.Append(20)
		l.Append(30)

		n1 := findNodeByValue(t, l, 10)

		l.MoveToBack(n1)

		if l.tail != n1 {
			t.Errorf("Голова не стала хвостом")
		}
		if l.head == nil || l.head.value != 20 {
			t.Errorf("Новой головой должен был стать элемент 20")
		}
	})

	t.Run("Перемещение узла, который уже в конце", func(t *testing.T) {
		l := New()
		l.Append(10)
		l.Append(20)

		n2 := findNodeByValue(t, l, 20)

		l.MoveToBack(n2)

		if l.tail != n2 {
			t.Error("Хвост списка изменился, хотя не должен был")
		}
	})
}

// findNodeByValue – Вспомогательная функция для тестов, чтобы найти узел по значению.
// Так как ваш Append не возвращает узел, мы найдем его перебором прямо в тесте.
func findNodeByValue(t *testing.T, l *linkedList, value int) *node {
	t.Helper()

	curr := l.head
	for curr != nil {
		if curr.value == value {
			return curr
		}
		curr = curr.next
	}
	return nil
}
