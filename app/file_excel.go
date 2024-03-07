package app

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func BuatFileExcel(filePath string, folderNames []string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// Set nama folder ke dalam cell A
	for i, folderName := range folderNames {
		cell := fmt.Sprintf("A%d", i+1)
		f.SetCellValue(sheetName, cell, folderName)
	}

	// Simpan ke file Excel
	err := f.SaveAs(filePath)
	if err != nil {
		return fmt.Errorf("error saat menyimpan file Excel: %w", err)
	}

	return nil
}
