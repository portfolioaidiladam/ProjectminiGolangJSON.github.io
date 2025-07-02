package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// logJson adalah fungsi untuk mengubah data Go menjadi JSON dan menampilkannya ke console
func logJson(data interface{}) {
	bytes, err := json.Marshal(data) // Mengubah data Go menjadi byte JSON
	if err != nil {                 // Jika terjadi error saat encoding
		panic(err)                    // Hentikan eksekusi dan tampilkan error
	}
	fmt.Println(string(bytes))       // Cetak hasil encoding JSON ke console
}

// TestEncode adalah unit test untuk fungsi logJson dengan berbagai tipe data
func TestEncode(t *testing.T) {
	logJson("Aidil")                                 // Encode string
	logJson(1)                                     // Encode integer
	logJson(true)                                  // Encode boolean
	logJson([]string{"Aidil", "Adam", "BaikHati"}) // Encode slice of string
}
