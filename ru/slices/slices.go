// Package срезы — русские имена для slices.
// Обобщённые функции нельзя присвоить переменной без инстанцирования,
// поэтому здесь обёртки, а не var.
package срезы

import (
	"cmp"
	"iter"
	"slices"
)

func Содержит[S ~[]E, E comparable](с S, э E) bool { return slices.Contains(с, э) }
func СодержитФунк[S ~[]E, E any](с S, ф func(E) bool) bool {
	return slices.ContainsFunc(с, ф)
}
func Индекс[S ~[]E, E comparable](с S, э E) int         { return slices.Index(с, э) }
func ИндексФунк[S ~[]E, E any](с S, ф func(E) bool) int { return slices.IndexFunc(с, ф) }
func Равны[S ~[]E, E comparable](а, б S) bool           { return slices.Equal(а, б) }

func Мин[S ~[]E, E cmp.Ordered](с S) E  { return slices.Min(с) }
func Макс[S ~[]E, E cmp.Ordered](с S) E { return slices.Max(с) }

func Сортировать[S ~[]E, E cmp.Ordered](с S) { slices.Sort(с) }
func СортироватьФунк[S ~[]E, E any](с S, ср func(а, б E) int) {
	slices.SortFunc(с, ср)
}
func СортироватьСтабильно[S ~[]E, E any](с S, ср func(а, б E) int) {
	slices.SortStableFunc(с, ср)
}
func Отсортирован[S ~[]E, E cmp.Ordered](с S) bool { return slices.IsSorted(с) }
func ДвоичныйПоиск[S ~[]E, E cmp.Ordered](с S, э E) (int, bool) {
	return slices.BinarySearch(с, э)
}

func Развернуть[S ~[]E, E any](с S)                    { slices.Reverse(с) }
func Клонировать[S ~[]E, E any](с S) S                 { return slices.Clone(с) }
func Сжать[S ~[]E, E comparable](с S) S                { return slices.Compact(с) }
func Вставить[S ~[]E, E any](с S, и int, э ...E) S     { return slices.Insert(с, и, э...) }
func Удалить[S ~[]E, E any](с S, от, до int) S         { return slices.Delete(с, от, до) }
func УдалитьЕсли[S ~[]E, E any](с S, ф func(E) bool) S { return slices.DeleteFunc(с, ф) }
func Повторить[S ~[]E, E any](с S, раз int) S          { return slices.Repeat(с, раз) }
func Склеить[S ~[]E, E any](срезы ...S) S              { return slices.Concat(срезы...) }

// Итераторы.
func Все[S ~[]E, E any](с S) iter.Seq2[int, E]   { return slices.All(с) }
func Значения[S ~[]E, E any](с S) iter.Seq[E]    { return slices.Values(с) }
func Задом[S ~[]E, E any](с S) iter.Seq2[int, E] { return slices.Backward(с) }
func Собрать[E any](посл iter.Seq[E]) []E        { return slices.Collect(посл) }
func Отсортированный[E cmp.Ordered](посл iter.Seq[E]) []E {
	return slices.Sorted(посл)
}
func Куски[S ~[]E, E any](с S, размер int) iter.Seq[S] {
	return slices.Chunk(с, размер)
}
