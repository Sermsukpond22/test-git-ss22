package main

import (
	"fmt"
	"log"
	"net/http"
)

func headHandler(w http.ResponseWriter, r *http.Request) {
	// ตรวจสอบว่าเป็น HTTP Method GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// ตั้งค่า Header ตอบกลับเป็น UTF-8 เพื่อรองรับภาษาไทย
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// ส่งข้อความออกไป
	fmt.Fprint(w, "นี่คือหัวหน้า")
}

// dev-junior-2
func somsriHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "นี่คือสมาชิก: สมศรี")
}

// dev-junior-1
func somchaiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "นี่คือสมาชิก: สมชาย")
}

func main() {
	http.HandleFunc("/head", headHandler)
	http.HandleFunc("/somsri", somsriHandler)
	http.HandleFunc("/somchai", somchaiHandler)
	fmt.Println("Server is running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
