package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Result struct {
	Name  string `col:"who"`
	Point int
	Age   int `col:"-"` // ignore
}

const COLTAG = "col"

func main() {

	x := []Result{Result{"sam", 100, 24}, Result{"tom", 0, 100025}}

	file, err := os.OpenFile("aaaa.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	WriteCSV(file, x)
}

// Convert interface{} to []interface{}
func toSlice(src interface{}) []interface{} {

	ret := []interface{}{}

	if v := reflect.ValueOf(src); v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			ret = append(ret, v.Index(i).Interface())
		}
	} else {
		ret = append(ret, v.Interface())
	}

	return ret
}

// Generate csv rows including header from interface{} slice or object
func genRows(src interface{}) [][]string {

	sl := toSlice(src)
	rows := make([][]string, 1)
	ignoreColIndex := map[int]bool{}

	for n, d := range sl {
		rows = append(rows, []string{})
		v := reflect.ValueOf(d)

		for i := 0; i < v.NumField(); i++ {
			if n == 0 {
				// Header
				colName := v.Type().Field(i).Tag.Get(COLTAG)
				if colName == "" {
					colName = strings.ToLower(v.Type().Field(i).Name)
				} else if colName == "-" {
					ignoreColIndex[i] = true
					continue
				}
				rows[0] = append(rows[0], colName)
			}

			if !ignoreColIndex[i] {
				rows[len(rows)-1] = append(rows[len(rows)-1], fmt.Sprint(v.Field(i).Interface()))
			}
		}
	}
	return rows
}

// Write csv file to path
func WriteCSV(file *os.File, data interface{}) {
	rows := genRows(data)

	writer := csv.NewWriter(file)
	for _, row := range rows {
		writer.Write(row)
	}
	writer.Flush()
}
