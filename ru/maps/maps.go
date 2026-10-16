// Package карты — русские имена для maps.
package карты

import (
	"iter"
	"maps"
)

func Ключи[M ~map[K]V, K comparable, V any](к M) iter.Seq[K]    { return maps.Keys(к) }
func Значения[M ~map[K]V, K comparable, V any](к M) iter.Seq[V] { return maps.Values(к) }
func Все[M ~map[K]V, K comparable, V any](к M) iter.Seq2[K, V]  { return maps.All(к) }
func Собрать[K comparable, V any](посл iter.Seq2[K, V]) map[K]V {
	return maps.Collect(посл)
}
func Клонировать[M ~map[K]V, K comparable, V any](к M) M { return maps.Clone(к) }
func Копировать[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](куда M1, откуда M2) {
	maps.Copy(куда, откуда)
}
func Равны[M1, M2 ~map[K]V, K, V comparable](а M1, б M2) bool { return maps.Equal(а, б) }
func УдалитьЕсли[M ~map[K]V, K comparable, V any](к M, ф func(K, V) bool) {
	maps.DeleteFunc(к, ф)
}
