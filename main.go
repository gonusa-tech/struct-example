package main

import "fmt"

type Student struct {
	ID   int
	Name string
}

type Classroom struct {
	Name     string
	Students []Student
}

type Course struct {
	Name       string
	Prequisite *Course
}

func main() {

	// disini s menyimpan value dari data student.
	s := Student{ID: 1, Name: "Budi"}

	fmt.Printf("%s memiliki ID %d\n", s.Name, s.ID)

	// sedangkan p menyimpan amalat ke memory tempat value dari student disimpan.

	p := &Student{ID: 1, Name: "Budi"}
	fmt.Printf("%s memiliki ID %d\n", p.Name, p.ID)

	// Course
	class := Classroom{
		Name: "Kelas Kimia",
		Students: []Student{
			{
				ID:   2,
				Name: "Tono",
			},
			s,
		},
	}

	fmt.Printf("%+v\n", class)

	// Contoh Penggunaan Course
	course := Course{
		Name: "Struktur Data",
		Prequisite: &Course{
			Name:       "Dasar Pemograman",
			Prequisite: nil,
		},
	}

	fmt.Printf("%+v", course)

}
