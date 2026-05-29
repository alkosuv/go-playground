# Context


## Когда использовать context?

### WithTimeout / WithDeadline / CancelFunc / WithCancel / WithCancelCause
- Если Go routine'а может зависнуть на долго, то лучше использовать context для прирывания работы  Go routine'ы.
- Если метод/функция ходит по сети, то такой первым параметром надо передавать context. Это необходимо для прирывания 
долгих запросов во избижания зависания приложения.

### WithValue

- В редких случая можено испольховать. Чаще всего если нет легального способа передать данные через сторонний пакет

## Context

### Background() Context

Возвращает пустой контекст, который никогда не отменялся, не имеет знечений и не имеет крайнего срока действия. 
Обычно вызывается в функции main, в функциях инициализации и в тестах и возвращает контекст верхнего уровня для входящих запросов.

### TODO() Context

Возвращает пустой контекст, но предназначенный для использования в качестве заполнителя, когда не ясно, какой контекст 
использовать, или когда родительский контекст еще недоступен.

###  WithCancel(Context) (Context, CancelFunc)
 Ничего не принимает и возвращает только функцию, которую можно вызвать, чтобы явно отменить действие контекста.

### WithDeadline(Context, time.Time) (Context, CancelFunc)

Принимает интервал времени, по истечении которого действие контекста будет прекращено, а канал Done – закрыт.
Внутри вызывает WithCancel

```
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// The current deadline is already sooner than the new one.
		return WithCancel(parent)
	}

    // continue func .... 
}
```

### WithTimeout(Context, time.Duration) (Context, CancelFunc)

Принимает интервал времени, по истечении которого действие контекста будет прекращено, а канал Done – закрыт.
В реализации вызвает WithDeadline

```
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}
```



## Link

- [Package context](https://pkg.go.dev/context)
- [YT Defer panic | Context](https://www.youtube.com/watch?v=Fjkckov4F38)