// Package синхр — русские имена для sync.
// Типы встраивают оригиналы: русские методы добавлены,
// английские (Lock, Wait, ...) работают как раньше.
package синхр

import "sync"

type Мьютекс struct{ sync.Mutex }

func (м *Мьютекс) Захватить()                 { м.Lock() }
func (м *Мьютекс) Освободить()                { м.Unlock() }
func (м *Мьютекс) ПопробоватьЗахватить() bool { return м.TryLock() }

type ЧЗМьютекс struct{ sync.RWMutex }

func (м *ЧЗМьютекс) Захватить()        { м.Lock() }
func (м *ЧЗМьютекс) Освободить()       { м.Unlock() }
func (м *ЧЗМьютекс) ЗахватитьЧтение()  { м.RLock() }
func (м *ЧЗМьютекс) ОсвободитьЧтение() { м.RUnlock() }

type ГруппаОжидания struct{ sync.WaitGroup }

func (г *ГруппаОжидания) Добавить(н int)     { г.Add(н) }
func (г *ГруппаОжидания) Готово()            { г.Done() }
func (г *ГруппаОжидания) Ждать()             { г.Wait() }
func (г *ГруппаОжидания) Запустить(ф func()) { г.Go(ф) }

type Однократно struct{ sync.Once }

func (о *Однократно) Выполнить(ф func()) { о.Do(ф) }

// Карта — потокобезопасная sync.Map.
type Карта = sync.Map

var (
	ОднократноФунк = sync.OnceFunc
)

func ОднократноЗначение[T any](ф func() T) func() T { return sync.OnceValue(ф) }
