// Package преобр — русские имена для strconv.
package преобр

import "strconv"

var (
	ВСтроку  = strconv.Itoa
	ИзСтроки = strconv.Atoi

	РазобратьЦелое        = strconv.ParseInt
	РазобратьБцелое       = strconv.ParseUint
	РазобратьВещественное = strconv.ParseFloat
	РазобратьБулево       = strconv.ParseBool

	ФорматЦелое        = strconv.FormatInt
	ФорматБцелое       = strconv.FormatUint
	ФорматВещественное = strconv.FormatFloat
	ФорматБулево       = strconv.FormatBool

	ВКавычки  = strconv.Quote
	ИзКавычек = strconv.Unquote

	ОшибкаСинтаксиса = strconv.ErrSyntax
	ОшибкаДиапазона  = strconv.ErrRange
)

type ОшибкаЧисла = strconv.NumError
