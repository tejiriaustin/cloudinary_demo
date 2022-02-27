package excelservice

import (
	"cloudinary_demo/model"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strconv"
)

func OpenAndRead() ([][]string, error) {
	fileName := os.Getenv("CLOUDINARY_URL")

	f, err := excelize.OpenFile(fileName)
	if err != nil {

		log.Fatal(err)
	}
	defer f.Close()

	rows, err := f.GetRows("RESULTS")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func SaveToFile(dt model.Data) error {
	fileName := os.Getenv("CLOUDINARY_URL")

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := checkForLastItemPosition()
	f.SetCellValue("RESULTS", s, "")
	return nil
}

func toInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return r
}

func checkForLastItemPosition() string {
	fileName := os.Getenv("CLOUDINARY_URL")

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows, err := f.GetRows("RESULTS")
	for _, row := range rows {
		for i, col := range row {
			if col == "" {
				return string(i)
			}
		}
	}
	return ""
}
