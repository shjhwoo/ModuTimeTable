package repo

import (
	"fmt"
	"reflect"
	"strings"
)

func GetInsertColumnsAndValues(entity any) ([]string, []any) {
	e := reflect.ValueOf(entity)

	if e.Kind() == reflect.Ptr {
		e = e.Elem()
	}

	var columns []string
	var values []any

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		if !f.IsZero() {
			sf := e.Type().Field(i)
			column := sf.Tag.Get("db")
			if column == "" || column == "Id" {
				continue
			}

			var value any
			typ := sf.Type.Kind()
			if typ == reflect.Slice || typ == reflect.Struct {
				continue
			}

			if typ == reflect.Ptr {
				if ptrType := f.Elem().Kind(); ptrType == reflect.Slice || ptrType == reflect.Struct {
					continue
				}
			}

			value = f.Interface()
			if value != nil {
				columns = append(columns, column)
				values = append(values, value)
			}
		}
	}

	return columns, values
}

func GetUpdateColumnsAndValues(entity any) ([]string, []any) {
	e := reflect.ValueOf(entity)

	if e.Kind() == reflect.Ptr {
		e = e.Elem()
	}

	var columns []string
	var values []any

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		if !f.IsZero() {
			sf := e.Type().Field(i)
			column := sf.Tag.Get("db")
			if column == "" || column == "Id" {
				continue
			}

			var value any
			typ := sf.Type.Kind()
			if typ == reflect.Slice || typ == reflect.Struct {
				continue
			}

			if typ == reflect.Ptr {
				if ptrType := f.Elem().Kind(); ptrType == reflect.Slice || ptrType == reflect.Struct {
					continue
				}
			}

			value = f.Interface()
			if value != nil {
				columns = append(columns, fmt.Sprintf("%s = ?", column))
				values = append(values, value)
			}
		}
	}

	return columns, values
}

func BuildPlaceHolders(n int) string {
	if n <= 0 {
		return ""
	}

	var placeHolders []string
	for i := 0; i < n; i++ {
		placeHolders = append(placeHolders, "?")
	}

	return fmt.Sprintf("(%s)", strings.Join(placeHolders, ", "))
}
