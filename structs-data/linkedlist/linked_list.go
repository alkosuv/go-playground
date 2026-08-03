package linkedlist

import (
	"errors"
)

//go:generate go test -v -run=.

// Односвязный спиское
//
// Правила:
// 1. Тип данных: Список работает строго с целыми числами (`int`)
// 2. Структура: Односвязный список. Каждый узел (`node`) знает только свой `next`.
// 3. Оптимизация: Структура `linkedList` хранит `head`, `tail` и `size`.
//    - Добавление в конец (`Append`) должно работать за O(1) благодаря указателю `tail`.
// 4. Дубликаты: Разрешены (например, 10 -> 20 -> 20).
//    - Поиск и удаление всегда работают с ПЕРВЫМ вхождением элемента.

var (
	ErrIndexOutOfBound = errors.New("index out of bounds")
	ErrorNodeNil       = errors.New("node is nil")
)

type node struct {
	next  *node
	value int
}

type linkedList struct {
	head *node // первый элемент (аналог index = 0)
	tail *node // последний элемент (аналог index = size-1)
	size int
}

func New() *linkedList {
	return new(linkedList)
}

// Get – возращает значение по индексу
// Сложность алгоритма O(N)
func (l *linkedList) Get(index int) (int, error) {
	node, err := l.getNodeAt(index)
	if err != nil {
		return 0, err
	}

	if node == nil {
		panic(ErrorNodeNil)
	}

	return node.value, nil
}

// Find – ищет значение в списке и возвращает индекс
// Если значение не найдено, то индекс -1
// Сложность алгоритма O(N)
func (l *linkedList) Find(value int) int {
	_, index := l.getNode(value)
	return index
}

// Contains – проверяет если ли указаное значение в списке
// TODO
func (l *linkedList) Contains(valut int) bool {
	// TODO: Проверить что в списке есть значения
	return false
}

// Append – встаялет значение в конец списка (после tail)
// Сложность алгоритма O(1)
func (l *linkedList) Append(value int) {
	// 1. Создать новуй узел и добавить туда value
	// 2. Увеличить size на 1
	// 3. Проверить что head и tail равны nil (IsEmpty()) если так, то устанавливем новый узел
	// 4. Получить текущую голову и добавить в next новую узел
	// 5. Установить в head новый узел

	l.size++
	newNode := &node{value: value}

	if l.IsEmpty() {
		l.head, l.tail = newNode, newNode
		return
	}

	l.tail.next, l.tail = newNode, newNode
}

// Prepend – встаялет значение в начало списка (перед head)
// Сложность алгоритма O(1)
func (l *linkedList) Prepend(value int) {
	// 1. Создать новые узел
	// 2. Увеличить значение size на 1
	// 3. Проверить что head и tail равны nil (IsEmpty()) если так, то устанавливем новый узел
	// 4. В новый узел в next делаем указание на tail
	// 5. В tail пишем новый узел

	l.size++
	newNode := &node{value: value}

	if l.IsEmpty() {
		l.tail, l.head = newNode, newNode
		return
	}

	newNode.next, l.head = l.head, newNode
}

// Insert – вставляет элемент по конкретному порядковому номеру
// Сложность алгоритма O(N)
func (l *linkedList) Insert(index int, value int) error {
	// 1. Проверить что index не выходит за пределы массива
	// 2. Найти нужно найти index-1, так как нужно будет обновить next в index-1
	//    и прописать корретный next в новом узле.
	//    (очень похоже на операцию get, можно сделать метод для поиска новы и использовать его)
	// 3. Обновить связи
	// 4. Если index=0, то значение вставляет в самое начало (Prepend)
	// 5. Индекс равень size

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	prevNode, err := l.getNodeAt(index - 1)
	if err != nil {
		return err
	}

	if prevNode == nil {
		panic(ErrorNodeNil)
	}

	if index == l.size {
		l.Append(value)
		return nil
	}

	newNode := &node{
		value: value,
		next:  prevNode.next,
	}
	prevNode.next = newNode
	l.size++

	return nil
}

// MoveToFront – перемещает узел в начала списка
// TODO
func (l *linkedList) MoveToFront(n *node) {}

// MoveToBack – перемещает узел в конец списка
// TODO
func (l *linkedList) MoveToBack(n *node) {}

// Remove – удаляет указаное значение из списка
func (l *linkedList) Remove(value int) bool {
	if l.IsEmpty() {
		return false
	}

	var (
		isFind   bool
		index    int
		prevNode *node
		findNode = l.head
	)

	for range l.size {
		if findNode.value == value {
			isFind = true
			break
		}

		prevNode = findNode
		findNode = findNode.next
		index++
	}

	if !isFind {
		return false
	}

	l.size--

	// Сценарий 1: Удаляем самый первый элемент (head)
	if prevNode == nil {
		l.head = findNode.next
		findNode.next = nil
		if l.head == nil {
			l.tail = nil
		}
		return true
	}

	// Сценарий 2: Удаляем самый последний элемент (tail)
	// size зарение уменьше на 1, поэтому просто сравниваем с index
	if index == l.size {
		l.tail = prevNode
		prevNode.next = nil
		return true
	}

	// Сценарий 3: Удаляем элемент из середины
	prevNode.next = findNode.next
	findNode.next = nil
	return true
}

// RemoveAt – удаляет значение по index
// TODO
func (l *linkedList) RemoveAt(index int) (int, error) {
	// TODO: Проверить что в списке есть значения
	return 0, nil
}

// Clear – удаляет все элементы в списке
// TODO
func (l *linkedList) Clear() {
}

// Size – возращет количество элементов в списке
func (l *linkedList) Size() int {
	return l.size
}

// IsEmpty – возвращает true, если список не содержит элементов
func (l *linkedList) IsEmpty() bool {
	if l.tail == nil && l.head == nil {
		return true
	}
	return false
}

// getNode – возращает узел по индексу
// Сложность алгоритма O(N)
func (l *linkedList) getNode(value int) (*node, int) {
	findNode := l.head
	for i := range l.size {
		if findNode.value == value {
			return findNode, i
		}
		findNode = findNode.next
	}

	return nil, -1
}

// getNodeAt – возращает узел по индексу
// Сложность алгоритма O(N)
func (l *linkedList) getNodeAt(index int) (*node, error) {
	// 1. Проверить что index не выходи за границы массива
	// 2. Перемещаться по узлам (head -> tail), пока не найдём нужный индекс
	if index >= l.size || index < 0 {
		return nil, ErrIndexOutOfBound
	}

	findNode := l.head
	for range index {
		findNode = findNode.next
	}

	return findNode, nil
}
