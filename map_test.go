package belajar_golang_json

import (
	"encoding/json" // Import package encoding/json untuk proses encoding dan decoding JSON
	"fmt"           // Import package fmt untuk output ke console
	"testing"       // Import package testing untuk unit test
)

// TestMapDecode menguji proses decoding JSON ke map di Go
func TestMapDecode(t *testing.T) {
	// JSON string yang akan didecode
	jsonString := `{"id":"p0001", "name":"Apple Mac Book Pro", "price": 20000000}`
	jsonBytes := []byte(jsonString) // Mengubah string ke byte array

	var result map[string]interface{} // Membuat map dengan key string dan value tipe data apapun
	json.Unmarshal(jsonBytes, &result) // Melakukan decoding JSON ke map

	fmt.Println(result)         // Menampilkan seluruh map hasil decoding
	fmt.Println(result["id"])   // Menampilkan value dengan key "id"
	fmt.Println(result["name"]) // Menampilkan value dengan key "name"
	fmt.Println(result["price"])// Menampilkan value dengan key "price"
}

// TestMapEncode menguji proses encoding map ke JSON di Go
func TestMapEncode(t *testing.T) {
	// Membuat map dengan key string dan value tipe data apapun
	product := map[string]interface{}{
		"id" : "P0001",           // ID produk
		"name" : "Apple Mac Book Pro", // Nama produk
		"price" : 20000000,       // Harga produk
	}

	bytes, _ := json.Marshal(product) // Mengubah map menjadi JSON
	fmt.Println(string(bytes))        // Menampilkan hasil encoding JSON ke console
}
