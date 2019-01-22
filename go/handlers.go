package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Value struct {
	Id float64 `json:"id,omitempty"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func sum(w http.ResponseWriter, r *http.Request) { //ผลหาร
	var value Value //struc
	_ = json.NewDecoder(r.Body).Decode(&value)

	t := time.Now()
	y := float64(t.Year())
	mon := t.Month()
	var i int = int(mon)
	var g = float64(i)
	d := float64(t.Day())
	var all = y + g + d
	fmt.Fprintln(w, "Current Date is :", y, mon, d)

	inputdate := value.Id
	// fmt.Fprintln(w, "inputdate is:", inputdate)
	// aa, err := strconv.ParseFloat(inputdate, 64)
	// if err == nil {
	// fmt.Println(aa)
	// }
	if inputdate == 0 {
		fmt.Fprintln(w, "No value please retried ")
		return
	}
	fmt.Fprintln(w, "Sum are ", all, " / ", inputdate)
	var ll = float64(all / inputdate)
	json.NewEncoder(w).Encode(ll)
	fmt.Fprintf(w, "Answer is %.4f:", ll)

}

func date(w http.ResponseWriter, r *http.Request) { //วันที่
	t := time.Now()
	y := float64(t.Year())
	mon := t.Month()
	var i int = int(mon)
	var g = float64(i)
	d := float64(t.Day())
	var all = y + g + d
	fmt.Fprintln(w, "Current Date is :", y, mon, d)
	fmt.Fprintf(w, "Sum date is %.4f:", all)
}

func isNumeric(s string) bool { //เช็คค่า int
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
