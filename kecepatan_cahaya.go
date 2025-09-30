// Seberapa lapa kita sampai di Mars?
package main

import "fmt"

func main() {
	const kecepatanCahaya = 100800 // km/s
	var jarak = 96300000           // km
	const satuhariBumi = 24        // jam
	fmt.Println(jarak/kecepatanCahaya/satuhariBumi, "hari")
}
