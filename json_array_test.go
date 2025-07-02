package belajar_golang_json

import (
	"encoding/json" // Import package encoding/json untuk proses encoding dan decoding JSON
	"fmt"           // Import package fmt untuk output ke console
	"testing"       // Import package testing untuk unit test
)

// TestJSONArray menguji proses encoding struct Customer ke JSON
func TestJSONArray(t *testing.T) {

	// Membuat data customer dengan beberapa hobi
	customer := Customer{
		FirstName:  "Aidil", // Nama depan
		MiddleName: "Adam",  // Nama tengah
		LastName:   "BaikHati", // Nama belakang
		Hobbies:    []string{"Gaming", "Reading", "Coding"}, // Slice string hobi
	}

	// Mengubah struct customer menjadi JSON
	bytes, _ := json.Marshal(customer)
	// Menampilkan hasil encoding JSON ke console
	fmt.Println(string(bytes))

}

// TestJSONArrayDecode menguji proses decoding JSON ke struct Customer
func TestJSONArrayDecode(t *testing.T) {
	// JSON string yang akan didecode
	jsonString := `{"FirstName":"Aidil","MiddleName":"Adam","LastName":"BaikHati","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString) // Mengubah string ke byte array

	customer := &Customer{} // Membuat pointer ke struct Customer
	// Melakukan decoding JSON ke struct Customer
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err) // Jika error, hentikan test
	}
	// Menampilkan hasil decoding ke console
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

// TestJSONArrayComplex menguji encoding struct Customer yang memiliki slice Address
func TestJSONArrayComplex(t *testing.T) {
	// Membuat data customer dengan beberapa alamat
	customer := Customer{
		FirstName: "Aidil",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada", // Nama jalan
				Country:    "Indonesia",      // Negara
				PostalCode: "9999",           // Kode pos
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
		},
	}

	// Mengubah struct customer menjadi JSON
	bytes, _ := json.Marshal(customer)
	// Menampilkan hasil encoding JSON ke console
	fmt.Println(string(bytes))
}

// TestJSONArrayComplexDecode menguji decoding JSON ke struct Customer dengan slice Address
func TestJSONArrayComplexDecode(t *testing.T) {
	// JSON string yang akan didecode
	jsonString := `{"FirstName":"Aidil","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"}]}`
	jsonBytes := []byte(jsonString) // Mengubah string ke byte array

	customer := &Customer{} // Membuat pointer ke struct Customer
	// Melakukan decoding JSON ke struct Customer
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err) // Jika error, hentikan test
	}
	// Menampilkan hasil decoding ke console
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

// TestOnlyJSONArrayComplexDecode menguji decoding JSON array ke slice Address
func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	// JSON array string yang akan didecode
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonString) // Mengubah string ke byte array

	addresses := &[]Address{} // Membuat pointer ke slice Address
	// Melakukan decoding JSON ke slice Address
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err) // Jika error, hentikan test
	}

	// Menampilkan hasil decoding ke console
	fmt.Println(addresses)
}

// TestOnlyJSONArrayComplex menguji encoding slice Address ke JSON array
func TestOnlyJSONArrayComplex(t *testing.T) {

	// Membuat slice Address
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
	}

	// Mengubah slice Address menjadi JSON array
	bytes, _ := json.Marshal(addresses)
	// Menampilkan hasil encoding JSON ke console
	fmt.Println(string(bytes))
}
