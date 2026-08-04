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

// linkedList – односвязанные список.
// Не подходит работы в конкурентной среде.
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
	node, _, err := l.getNodeAt(index)
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
	_, _, index := l.getNode(value)
	return index
}

// Contains – проверяет если ли указаное значение в списке
func (l *linkedList) Contains(value int) bool {
	node, _, _ := l.getNode(value)
	if node != nil {
		return true
	}

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

	prevNode, _, err := l.getNodeAt(index - 1)
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
func (l *linkedList) MoveToFront(index int) (bool, error) {
	// Если перемещаемый индекс указывает на голову, то перемезать не нужно
	if index == 0 {
		return false, nil
	}

	findNode, _, err := l.getNodeAt(index)
	if err != nil {
		return false, err
	}

	if _, err := l.RemoveAt(index); err != nil {
		return false, err
	}

	l.Prepend(findNode.value)

	return true, nil
}

// MoveToBack – перемещает узел в конец списка
func (l *linkedList) MoveToBack(index int) (bool, error) {
	// Если перемещаемый индекс указывает на хвост, то перемезать не нужно
	if index == l.size-1 {
		return false, nil
	}

	findNode, _, err := l.getNodeAt(index)
	if err != nil {
		return false, err
	}

	if _, err := l.RemoveAt(index); err != nil {
		return false, err
	}

	l.Append(findNode.value)

	return true, nil
}

// Remove – удаляет указаное значение из списка
func (l *linkedList) Remove(value int) bool {
	findNode, prevNode, index := l.getNode(value)
	if index == -1 {
		return false
	}

	l.size--

	// Сценарий 1: Удаляем самый первый элемент (head)
	if prevNode == nil {
		l.head, findNode.next = findNode.next, nil

		if l.head == nil {
			l.tail = nil
		}

		return true
	}

	// Сценарий 2: Удаляем самый последний элемент (tail)
	// size зарение уменьше на 1, поэтому просто сравниваем с index
	if index == l.size {
		l.tail, prevNode.next = prevNode, nil
		return true
	}

	// Сценарий 3: Удаляем элемент из середины
	prevNode.next, findNode.next = findNode.next, nil
	return true
}

// RemoveAt – удаляет значение по index
// Возравщает значание которое было в узле или ошибку
func (l *linkedList) RemoveAt(index int) (int, error) {
	findNode, prevNode, err := l.getNodeAt(index)
	if err != nil {
		return 0, err
	}

	l.size--
	value := findNode.value

	if index == 0 {
		if l.size == 0 {
			l.head, l.tail = nil, nil
			return value, nil
		}

		node := l.head
		l.head, node = node.next, nil
		return value, nil
	}

	if index == l.size {
		if l.size == 0 {
			l.head, l.tail = nil, nil
			return value, nil
		}

		l.tail, prevNode.next = prevNode, nil
		return value, nil
	}

	prevNode.next, findNode = findNode.next, nil

	return value, nil
}

// Clear – удаляет все элементы в списке
func (l *linkedList) Clear() {
	for l.head != nil {
		l.RemoveAt(0)
	}
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

// getNode – возращает узел по значению
// Возращабтся следующие параметры искомый узел, узел перед искомым, индекс узла в списке
// Узел перед искаомым нужен в некоторсы методов, чтобы в них не дублировать поиск
// Сложность алгоритма O(N)
func (l *linkedList) getNode(value int) (*node, *node, int) {
	if l.IsEmpty() {
		return nil, nil, -1
	}

	var (
		prevNode *node
		findNode = l.head
	)

	for i := range l.size {
		if findNode.value == value {
			return findNode, prevNode, i
		}
		prevNode, findNode = findNode, findNode.next
	}

	return nil, nil, -1
}

// getNodeAt – ищет узел по индекс
// Возращабтся следующие параметры искомый узел, узел перед искомым, ошибка
// Узел перед искаомым нужен в некоторсы методов, чтобы в них не дублировать поиск
// Сложность алгоритма O(N)
func (l *linkedList) getNodeAt(index int) (*node, *node, error) {
	if index >= l.size || index < 0 {
		return nil, nil, ErrIndexOutOfBound
	}

	if l.IsEmpty() {
		return nil, nil, nil
	}

	var (
		prevNode *node
		findNode = l.head
	)

	for range index {
		prevNode, findNode = findNode, findNode.next
	}

	return findNode, prevNode, nil
}
