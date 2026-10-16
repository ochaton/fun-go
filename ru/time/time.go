// Package время — русские имена для time.
// Методы Time/Duration остаются английскими.
package время

import "time"

type (
	Время             = time.Time
	Продолжительность = time.Duration
	Месяц             = time.Month
	ДеньНедели        = time.Weekday
	Таймер            = time.Timer
	Тикер             = time.Ticker
	Зона              = time.Location
)

const (
	Наносекунда  = time.Nanosecond
	Микросекунда = time.Microsecond
	Миллисекунда = time.Millisecond
	Секунда      = time.Second
	Минута       = time.Minute
	Час          = time.Hour
)

const (
	Январь   = time.January
	Февраль  = time.February
	Март     = time.March
	Апрель   = time.April
	Май      = time.May
	Июнь     = time.June
	Июль     = time.July
	Август   = time.August
	Сентябрь = time.September
	Октябрь  = time.October
	Ноябрь   = time.November
	Декабрь  = time.December
)

const (
	Воскресенье = time.Sunday
	Понедельник = time.Monday
	Вторник     = time.Tuesday
	Среда       = time.Wednesday
	Четверг     = time.Thursday
	Пятница     = time.Friday
	Суббота     = time.Saturday
)

// Шаблоны форматирования.
const (
	ФорматRFC3339   = time.RFC3339
	ФорматДата      = time.DateOnly
	ФорматВремя     = time.TimeOnly
	ФорматДатаВремя = time.DateTime
)

var (
	Сейчас                     = time.Now
	Спать                      = time.Sleep
	ПрошлоС                    = time.Since
	ОсталосьДо                 = time.Until
	Через                      = time.After
	Тик                        = time.Tick
	НовыйТаймер                = time.NewTimer
	НовыйТикер                 = time.NewTicker
	ПослеФунк                  = time.AfterFunc
	Дата                       = time.Date
	Разобрать                  = time.Parse
	РазобратьПродолжительность = time.ParseDuration
	ИзЮникс                    = time.Unix
	ЗагрузитьЗону              = time.LoadLocation

	UTC     = time.UTC
	Местное = time.Local
)
