// Package строки — русские имена для strings.
package строки

import "strings"

// Поиск.
var (
	Содержит         = strings.Contains
	СодержитЛюбой    = strings.ContainsAny
	СодержитРуну     = strings.ContainsRune
	СодержитФунк     = strings.ContainsFunc
	Количество       = strings.Count
	НачинаетсяС      = strings.HasPrefix
	ЗаканчиваетсяНа  = strings.HasSuffix
	Индекс           = strings.Index
	ИндексРуны       = strings.IndexRune
	ИндексЛюбого     = strings.IndexAny
	ПоследнийИндекс  = strings.LastIndex
	РавныБезРегистра = strings.EqualFold
	Сравнить         = strings.Compare
)

// Разбиение и сборка.
var (
	Разделить        = strings.Split
	РазделитьН       = strings.SplitN
	РазделитьПосле   = strings.SplitAfter
	Поля             = strings.Fields
	ПоляФунк         = strings.FieldsFunc
	Разрезать        = strings.Cut
	РазрезатьПрефикс = strings.CutPrefix
	РазрезатьСуффикс = strings.CutSuffix
	Соединить        = strings.Join
	Повторить        = strings.Repeat

	// Итераторы.
	Строки        = strings.Lines
	РазделитьПосл = strings.SplitSeq
	ПоляПосл      = strings.FieldsSeq
)

// Преобразование.
var (
	Заменить        = strings.Replace
	ЗаменитьВсе     = strings.ReplaceAll
	ВНижний         = strings.ToLower
	ВВерхний        = strings.ToUpper
	Обрезать        = strings.Trim
	ОбрезатьСлева   = strings.TrimLeft
	ОбрезатьСправа  = strings.TrimRight
	ОбрезатьПробелы = strings.TrimSpace
	ОбрезатьПрефикс = strings.TrimPrefix
	ОбрезатьСуффикс = strings.TrimSuffix
	Отобразить      = strings.Map
	Клонировать     = strings.Clone
)

// Типы.
type (
	Читатель   = strings.Reader
	Заменитель = strings.Replacer
)

var (
	НовыйЧитатель   = strings.NewReader
	НовыйЗаменитель = strings.NewReplacer
)

// Строитель — strings.Builder с русскими методами.
// Английские методы (WriteString, String, ...) тоже на месте,
// поэтому *Строитель остаётся io.Writer и fmt.Stringer.
type Строитель struct {
	strings.Builder
}

func (с *Строитель) Записать(текст string) *Строитель {
	с.WriteString(текст)
	return с
}
func (с *Строитель) ЗаписатьРуну(р rune) *Строитель {
	с.WriteRune(р)
	return с
}
func (с *Строитель) ЗаписатьБайт(б byte) *Строитель {
	с.WriteByte(б)
	return с
}
func (с *Строитель) Строка() string  { return с.String() }
func (с *Строитель) Длина() int      { return с.Len() }
func (с *Строитель) Ёмкость() int    { return с.Cap() }
func (с *Строитель) Нарастить(н int) { с.Grow(н) }
func (с *Строитель) Сбросить()       { с.Reset() }
