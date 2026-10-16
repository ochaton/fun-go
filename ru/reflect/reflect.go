// Package отражение — русские имена для reflect.
// Методы Type/Value остаются английскими (Kind, Elem, Field, ...).
package отражение

import "reflect"

type (
	Тип      = reflect.Type
	Значение = reflect.Value
	Вид      = reflect.Kind
	Поле     = reflect.StructField
	Метод    = reflect.Method
)

var (
	ТипИз          = reflect.TypeOf
	ЗначениеИз     = reflect.ValueOf
	ГлубокоРавны   = reflect.DeepEqual
	Нуль           = reflect.Zero
	Новое          = reflect.New
	СоздатьСрез    = reflect.MakeSlice
	СоздатьКарту   = reflect.MakeMapWithSize
	СоздатьКанал   = reflect.MakeChan
	СоздатьФункцию = reflect.MakeFunc
	Косвенно       = reflect.Indirect
)

func ТипДля[T any]() Тип { return reflect.TypeFor[T]() }

// Виды.
const (
	Недопустимый   = reflect.Invalid
	Булево         = reflect.Bool
	Целое          = reflect.Int
	Целое8         = reflect.Int8
	Целое16        = reflect.Int16
	Целое32        = reflect.Int32
	Целое64        = reflect.Int64
	Бцелое         = reflect.Uint
	Бцелое8        = reflect.Uint8
	Бцелое16       = reflect.Uint16
	Бцелое32       = reflect.Uint32
	Бцелое64       = reflect.Uint64
	Вещественное32 = reflect.Float32
	Вещественное64 = reflect.Float64
	Комплексное64  = reflect.Complex64
	Комплексное128 = reflect.Complex128
	Массив         = reflect.Array
	Канал          = reflect.Chan
	Функция        = reflect.Func
	Интерфейс      = reflect.Interface
	Карта          = reflect.Map
	Указатель      = reflect.Pointer
	Срез           = reflect.Slice
	Строка         = reflect.String
	Структура      = reflect.Struct
)
