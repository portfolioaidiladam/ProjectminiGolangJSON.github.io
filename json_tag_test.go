package belajar_golang_json

import (
	"encoding/json" // Import package encoding/json untuk proses encoding dan decoding JSON
	"fmt"           // Import package fmt untuk output ke console
	"testing"       // Import package testing untuk unit test
)

// Product adalah struct yang merepresentasikan data produk
// Tag json digunakan untuk mapping nama field pada JSON
// Id akan di-encode/decode sebagai "id"
// Name akan di-encode/decode sebagai "name"
// ImageURL akan di-encode/decode sebagai "image_url"
type Product struct {
	Id       string `json:"id"`         // ID produk
	Name     string `json:"name"`       // Nama produk
	ImageURL string `json:"image_url"`  // URL gambar produk
}

// TestJSONTag menguji proses encoding struct Product ke JSON
func TestJSONTag(t *testing.T) {
	// Membuat data produk
	product := Product{
		Id:       "P0001", // ID produk
		Name:     "Apple Mac Book Pro", // Nama produk
		ImageURL: "http://example.com/image.png", // URL gambar
	}

	// Mengubah struct product menjadi JSON
	bytes, _ := json.Marshal(product)
	// Menampilkan hasil encoding JSON ke console
	fmt.Println(string(bytes))
}

// TestJSONTagDecode menguji proses decoding JSON ke struct Product
func TestJSONTagDecode(t *testing.T) {
	// JSON string yang akan didecode
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString) // Mengubah string ke byte array

	product := &Product{} // Membuat pointer ke struct Product
	// Melakukan decoding JSON ke struct Product
	json.Unmarshal(jsonBytes, product)
	// Menampilkan hasil decoding ke console
	fmt.Println(product)
}
