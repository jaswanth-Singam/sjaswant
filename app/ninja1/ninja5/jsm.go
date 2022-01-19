package main

import "fmt"

type Hospital struct {
	name    string
	reason  string
	receipt int
}

func (P Hospital) Recipt() int {
	return P.receipt
}
func main() {
	p1 := Hospital{
		name:    "ammar_hospital",
		reason:  "fever",
		receipt: 125,
	}
	fmt.Printf("i am going %v for %v my receipt %d\n", p1.name, p1.reason, p1.receipt)
	x := p1.Recipt()
	fmt.Println(x)
}
