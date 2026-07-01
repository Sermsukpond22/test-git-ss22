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

func main() {
	http.HandleFunc("/head", headHandler)

	fmt.Println("Server is running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
