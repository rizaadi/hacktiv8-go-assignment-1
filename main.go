package main

import (
	"fmt"
	"os"
	"strings"
)

type People struct {
	Id      int
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {
	peoples := []People{
		{Id: 0, Name: "Achmad Yusuf", Address: "Madura", Job: "Programmer", Reason: "Nintendo Switch"},
		{Id: 1, Name: "Saeful Kurniawan", Address: "Jogja", Job: "Project Manager", Reason: "Work Life Balance"},
		{Id: 2, Name: "Reza Wira", Address: "Jekartwah", Job: "Petani", Reason: "Slow living"},
		{Id: 3, Name: "Thomas", Address: "Gresik", Job: "Guru", Reason: "Iseng aja 🔥🔥"},
	}
	var args = os.Args[1:]
	if len(args) == 0 {
		fmt.Println("List Nama kosong")
		return
	}
	var newPeoples []People

	for _, arg := range args {
		found := false
		for _, person := range peoples {
			if strings.EqualFold(person.Name, arg) {
				newPeoples = append(newPeoples, person)
				found = true
			}
		}
		if !found {
			fmt.Printf("Orang dengan nama '%s' tidak ada\n", arg)
		}
	}

	for _, newPeople := range newPeoples {
		fmt.Println("ID :", newPeople.Id)
		fmt.Println("Nama :", newPeople.Name)
		fmt.Println("Alamat :", newPeople.Address)
		fmt.Println("Pekerjaan :", newPeople.Job)
		fmt.Println("Alasan :", newPeople.Reason)
	}
}
