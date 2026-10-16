// Package ос — русские имена для os.
package ос

import "os"

// Аргументы — функция, а не переменная: os.Args можно менять в рантайме.
func Аргументы() []string { return os.Args }

// Процесс и окружение.
var (
	Выход           = os.Exit
	Среда           = os.Getenv
	НайтиВСреде     = os.LookupEnv
	УстановитьСреду = os.Setenv
	УбратьИзСреды   = os.Unsetenv
	ВсяСреда        = os.Environ
	ИмяХоста        = os.Hostname
	ИДПроцесса      = os.Getpid
)

// Файлы и каталоги.
var (
	ЧитатьФайл       = os.ReadFile
	ЗаписатьФайл     = os.WriteFile
	Открыть          = os.Open
	ОткрытьФайл      = os.OpenFile
	Создать          = os.Create
	Удалить          = os.Remove
	УдалитьВсё       = os.RemoveAll
	Переименовать    = os.Rename
	СоздатьКаталог   = os.Mkdir
	СоздатьКаталоги  = os.MkdirAll
	ВременныйКаталог = os.MkdirTemp
	ЧитатьКаталог    = os.ReadDir
	РабочийКаталог   = os.Getwd
	СменитьКаталог   = os.Chdir
	Инфо             = os.Stat
)

// Стандартные потоки.
var (
	Ввод        = os.Stdin
	Вывод       = os.Stdout
	ВыводОшибок = os.Stderr
)

// Ошибки.
var (
	ОшибкаНеСуществует = os.ErrNotExist
	ОшибкаСуществует   = os.ErrExist
	ОшибкаДоступа      = os.ErrPermission
)

// Типы.
type (
	Файл           = os.File
	РежимФайла     = os.FileMode
	ИнфоФайла      = os.FileInfo
	ЗаписьКаталога = os.DirEntry
)

// Флаги для ОткрытьФайл.
const (
	ТолькоЧтение = os.O_RDONLY
	ТолькоЗапись = os.O_WRONLY
	ЧтениеЗапись = os.O_RDWR
	Дописывать   = os.O_APPEND
	Создавать    = os.O_CREATE
	Обнулять     = os.O_TRUNC
	Эксклюзивно  = os.O_EXCL
)
