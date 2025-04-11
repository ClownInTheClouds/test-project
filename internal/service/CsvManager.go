package service

import (
	"encoding/csv"
	"os"
)

func CreateCsvReader(file *os.File) *csv.Reader {
	var reader = csv.NewReader(file)
	reader.Comma = ';'
	reader.Comment = '#'
	return reader
}

func CreateCsvWriter(file *os.File) *csv.Writer {
	var writer = csv.NewWriter(file)
	writer.Comma = ';'
	writer.UseCRLF = true
	return writer
}
