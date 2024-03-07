package app

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

func BuatFolder(filePath, folderPath string) error {
	// Gunakan path absolut untuk file Excel
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("error mendapatkan path : %w", err)
	}

	// Gunakan path absolut untuk folder tujuan
	absFolderPath, err := filepath.Abs(folderPath)
	if err != nil {
		return fmt.Errorf("error mendapatkan path : %w", err)
	}

	// Buka file Excel
	f, err := excelize.OpenFile(absFilePath)
	if err != nil {
		return fmt.Errorf("error saat buka file excel: %w", err)
	}

	// Baca baris dari file Excel
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return fmt.Errorf("error saat baca baris: %w", err)
	}

	// Iterasi setiap baris dan buat folder
	for _, row := range rows {
		if len(row) > 0 {
			folderName := row[0]
			fullFolderPath := filepath.Join(absFolderPath, folderName)
			err := os.MkdirAll(fullFolderPath, 0755)
			if err != nil {
				fmt.Println("")
				return fmt.Errorf("gagal membuat folder %s: %w", fullFolderPath, err)
			} else {
				fmt.Println("")
				fmt.Println("Berhasil membuat folder", fullFolderPath)
			}
			
		}
	}

	return nil
}
