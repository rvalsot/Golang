package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type unit struct {
	name    string
	hp      int
	attack  int
	m_armor int
	p_armor int
}

type building struct {
	unit
	creates []string
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	fmt.Println("Methods")
	fmt.Println("--------")
	/*
		guy1 := person{"Juan", "Pijamas", 21}
		guy2 := person{"Mikhail", "Ostrobov", 23}
		fmt.Printf("Method type %T \n", guy1.fullName())
		fmt.Println(guy1.fullName())
		fmt.Println(guy2.fullName())
	*/
	cataphract := unit{
		name:    "cataphract",
		hp:      110,
		attack:  9,
		m_armor: 2,
		p_armor: 1,
	}
	castle_units := []string{"cataphract", "trebuchet", "petard"}
	castle :=
		building{
			unit: unit{
				name:    "castle",
				hp:      4800,
				attack:  11 * 4,
				m_armor: 8,
				p_armor: 11,
			},
			creates: castle_units,
		}
	var castle_pointer *building = &castle
	fmt.Println(castle)
	fmt.Println(cataphract)
	fmt.Println(*castle_pointer)
}
