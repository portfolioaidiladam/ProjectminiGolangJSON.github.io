package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Address merepresentasikan struktur alamat
// Setiap field akan otomatis menjadi key pada JSON
// Field harus diawali huruf kapital agar bisa diakses oleh package lain (termasuk encoding/json)
type Address struct {
	Street     string // Nama jalan
	Country    string // Nama negara
	PostalCode string // Kode pos
}

// Customer merepresentasikan struktur data customer
// Field dengan tipe data slice akan di-encode sebagai array di JSON
type Customer struct {
	FirstName  string    // Nama depan
	MiddleName string    // Nama tengah
	LastName   string    // Nama belakang
	Age        int       // Umur
	Married    bool      // Status menikah
	Hobbies    []string  // Hobi (array of string)
	Addresses  []Address // Daftar alamat (array of Address)
}

// TestJSONObject adalah unit test untuk encode struct Customer ke JSON
func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Aidil",         // Set nama depan
		MiddleName: "Adam",   // Set nama tengah
		LastName:   "BaikHati",    // Set nama belakang
		Age:        36,             // Set umur
		Married:    true,           // Set status menikah
	}

	bytes, _ := json.Marshal(customer)      // Encode struct customer ke JSON (dalam bentuk []byte)
	fmt.Println(string(bytes))              // Cetak hasil JSON ke console
}
