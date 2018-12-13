package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}

//Fight comment
func Fight(p *Character, q *Character) {
	for p.IsAlive || q.IsAlive {
		p.Hit(q)
		q.Hit(p)
	}
}
