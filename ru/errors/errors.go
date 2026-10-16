// Package ошибки — русские имена для errors.
package ошибки

import "errors"

var (
	Новая      = errors.New
	Является   = errors.Is
	Как        = errors.As
	Развернуть = errors.Unwrap
	Объединить = errors.Join

	ОшибкаНеПоддерживается = errors.ErrUnsupported
)

// КакТип — типобезопасный errors.As: ош, ок := ошибки.КакТип[*МояОшибка](е).
func КакТип[E error](е error) (E, bool) { return errors.AsType[E](е) }
