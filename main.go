package main

import (

	"fmt"
	"os"
	"banyak_folder/app"
)

func main() {
	var pilihan int

	for {
		fmt.Println("")
		fmt.Println("---- App Banyak Folder ----")
		fmt.Println("Menu: ")
		fmt.Println(" 1. Buat file")
		fmt.Println(" 2. Buat folder")
		fmt.Println(" 3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println("")

		switch pilihan {
		case 1:
			var filePath string
			var namaFileX string
			fmt.Println("---- Menu Buat File ----")
			fmt.Print("Masukkan nama file Excel: ")
			fmt.Scanln(&namaFileX)
			filePath = "file/" + namaFileX + ".xlsx"

			var jumlahFolder int
			fmt.Print("Masukkan jumlah folder: ")
			fmt.Scanln(&jumlahFolder)

			namaFolder := make([]string, jumlahFolder)
			for i := 0; i < jumlahFolder; i++ {
				fmt.Printf("Masukkan nama folder %d: ", i+1)
				fmt.Scanln(&namaFolder[i])
			}

			err := app.BuatFileExcel(filePath, namaFolder)
			if err != nil {
				fmt.Println("")
				fmt.Println("Pesan error:", err)
				fmt.Println("")
			} else {
				fmt.Println("")
				fmt.Println("File Excel berhasil dibuat:", filePath)
				fmt.Println("")
		
			}
		case 2:
			var folderPath string
			var namaFile string
			fmt.Println("---- Menu Buat Folder ----")
			fmt.Print("Masukan nama file(contoh: list_folder) : ")
			fmt.Scanln(&namaFile)
			fmt.Print("Masukan Path dimana folder akan di buat (Contoh: D:/karyawan/ ) : ")
			fmt.Scanln(&folderPath)

			filePath := "file/" + namaFile +".xlsx" 

			err := app.BuatFolder(filePath, folderPath)
			if err != nil {
				fmt.Println("Pesan error :", err)
				fmt.Println("")
			} else {
				fmt.Println("Proses selesai .")
				fmt.Println("")
			}
		case 3:
			fmt.Println("Keluar dari program.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
			fmt.Println("")
		}
	}
}
