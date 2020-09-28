package appsflyer_sdk

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"strings"
)

type CSVParser struct {
}

func (p *CSVParser) getReader(body string) *csv.Reader {
	reader := csv.NewReader(strings.NewReader(strings.TrimSuffix(body, "\n")))
	reader.FieldsPerRecord = -1
	return reader
}

func (p *CSVParser) Write(body string, file *os.File) error {
	reader := p.getReader(body)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func (p *CSVParser) Parse(body string, v interface{}, f func(result interface{}), tags ...string) error {
	reader := p.getReader(body)

	headers, err := reader.Read()
	if err != nil {
		return err
	}

	dataType := reflect.TypeOf(v)
	newData := reflect.New(dataType).Elem()

	if len(tags) == 0 {
		tags = []string{"csv"}
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		for i := 0; i < dataType.NumField(); i++ {
			var (
				f     = dataType.Field(i)
				index = -1
			)
			fieldName := f.Tag.Get(tags[0])
			for k, v := range headers {
				if v == fieldName {
					index = k
				}
			}
			if index == -1 {
				continue
			}
			newField := newData.FieldByName(f.Name)
			if !newField.IsValid() || !newField.CanSet() {
				continue
			}
			newField.Set(reflect.ValueOf(row[index]))
		}
		f(newData.Interface())
	}
	return nil
}
