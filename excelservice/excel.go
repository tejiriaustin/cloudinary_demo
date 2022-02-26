package excelservice

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func open() string {
	f, err := excelize.OpenFile("")
	if err != nil {
		log.Fatal(err)
	}

	cell, err := f.GetCellValue("value", "")
	if err != nil {
		log.Fatal(err)
	}

	return cell
}

func save() error {
	f, err := excelize.OpenFile("")
	if err != nil {
		return err
	}

	f.SetCellValue("", "", "")
	return nil
}
