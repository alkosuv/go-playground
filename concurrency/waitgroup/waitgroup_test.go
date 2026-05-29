package waitgroup_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type WaitGroup struct {
	count atomic.Int64
	err   atomic.Value
}

func New() *WaitGroup {
	return &WaitGroup{}
}

func (w *WaitGroup) Add(delta int64) {
	w.count.Add(delta)
}

func (w *WaitGroup) Down() {
	w.count.Add(-1)
}

func (w *WaitGroup) Wait() error {
	for w.count.Load() > 0 {
		// Цикл будет крутиться, пока count не станет 0
	}

	if err := w.err.Load(); err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func (w *WaitGroup) Go(fn func()) {
	w.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				w.err.Store(err)
			}

			w.Down()
		}()

		fn()
	}()
}

func TestWaitGroup_Success(t *testing.T) {
	wg := New()

	// Запускаем 5 горутин
	for range 5 {
		wg.Go(func() {
			// Имитируем какую-то работу
			time.Sleep(10 * time.Millisecond)
			// В реальном коде тут была бы критическая секция,
			// но для теста просто увеличим счетчик
			// (тут мы не используем атомарный счетчик, чтобы проверить сам факт вызова)
		})
	}

	// Ждем завершения всех
	err := wg.Wait()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestWaitGroup_Panic(t *testing.T) {
	wg := New()

	wg.Go(func() {
		time.Sleep(10 * time.Millisecond)
		panic("something went wrong!")
	})

	// Ждем завершения
	err := wg.Wait()

	if err == nil {
		t.Error("expected error from panic, but got nil")
	} else {
		expectedErr := "something went wrong!"
		if err.Error() != expectedErr {
			t.Errorf("expected error %q, got %q", expectedErr, err.Error())
		}
	}
}

func TestWaitGroup_ConcurrencyStress(t *testing.T) {
	wg := New()
	numGoroutines := 100

	// Используем атомарный счетчик для проверки реального выполнения задач
	var completedTasks atomic.Int64

	for i := range numGoroutines {
		wg.Go(func() {
			time.Sleep(time.Duration(i%10) * time.Millisecond)
			completedTasks.Add(1)
		})
	}

	err := wg.Wait()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if completedTasks.Load() != int64(numGoroutines) {
		t.Errorf("expected %d completed tasks, but got %d", numGoroutines, completedTasks.Load())
	}
}

func TestWaitGroup_AddAndDownManual(t *testing.T) {
	// Проверка ручного управления Add и Down (как в оригинальном sync.WaitGroup)
	wg := New()
	wg.Add(2)

	go func() {
		time.Sleep(10 * time.Millisecond)
		wg.Down()
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		wg.Down()
	}()

	err := wg.Wait()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
