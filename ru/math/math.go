// Package матем — русские имена для math.
package матем

import "math"

const (
	Пи = math.Pi
	Е  = math.E

	МаксЦелое          = math.MaxInt
	МинЦелое           = math.MinInt
	МаксЦелое32        = math.MaxInt32
	МаксЦелое64        = math.MaxInt64
	МаксВещественное64 = math.MaxFloat64
)

var (
	Модуль     = math.Abs
	Корень     = math.Sqrt
	Куб        = math.Cbrt
	Степень    = math.Pow
	Экспонента = math.Exp
	Лог        = math.Log
	Лог2       = math.Log2
	Лог10      = math.Log10
	Гипотенуза = math.Hypot

	Пол       = math.Floor
	Потолок   = math.Ceil
	Округлить = math.Round
	Отсечь    = math.Trunc
	Остаток   = math.Mod

	Синус       = math.Sin
	Косинус     = math.Cos
	Тангенс     = math.Tan
	Арксинус    = math.Asin
	Арккосинус  = math.Acos
	Арктангенс  = math.Atan
	Арктангенс2 = math.Atan2

	Бесконечность    = math.Inf
	ЭтоБесконечность = math.IsInf
	НеЧисло          = math.NaN
	ЭтоНеЧисло       = math.IsNaN
)
