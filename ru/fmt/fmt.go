// Package формат — русские имена для fmt.
package формат

import "fmt"

// Вывод в stdout.
var (
	Вывести       = fmt.Print
	ВывестиСтроку = fmt.Println
	ВывестиФ      = fmt.Printf
)

// Вывод в io.Writer.
var (
	Записать       = fmt.Fprint
	ЗаписатьСтроку = fmt.Fprintln
	ЗаписатьФ      = fmt.Fprintf
)

// В строку.
var (
	Строка       = fmt.Sprint
	СтрокаСтрока = fmt.Sprintln
	СтрокаФ      = fmt.Sprintf
	Добавить     = fmt.Append
	ДобавитьФ    = fmt.Appendf
)

// Ошибки.
var ОшибкаФ = fmt.Errorf

// Ввод.
var (
	Считать       = fmt.Scan
	СчитатьСтроку = fmt.Scanln
	СчитатьФ      = fmt.Scanf
	РазобратьФ    = fmt.Sscanf
)

// Интерфейсы. Имя метода (String, Format) менять нельзя —
// по нему fmt и находит реализацию.
type (
	Строкер  = fmt.Stringer
	Форматер = fmt.Formatter
)
