// Package сравнение — русские имена для cmp.
package сравнение

import "cmp"

type Упорядочимый = cmp.Ordered

// Сравнить возвращает -1, 0 или +1.
func Сравнить[T Упорядочимый](а, б T) int { return cmp.Compare(а, б) }

func Меньше[T Упорядочимый](а, б T) bool { return cmp.Less(а, б) }

// Или возвращает первое ненулевое значение.
func Или[T comparable](значения ...T) T { return cmp.Or(значения...) }
