package excelservice

import (
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

func Open() string {
	fileName := os.Getenv("CLOUDINARY_URL")

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cell, err := f.GetCellValue("", "")
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
