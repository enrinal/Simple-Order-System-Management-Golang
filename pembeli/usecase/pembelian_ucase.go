package usecase

import (
	"fmt"
	"simple-order-go/models"
)

type model models.Pembeli

func IsPembelian(pembeli model) bool {
	if pembeli.Status == 1 && pembeli.CountPembelian <= 5 && pembeli.CountItem <= 10 {
		return true
	} else if pembeli.Status == 2 && pembeli.CountItem <= 10 {
		return true
	} else {
		return false
	}
}


func (pembeli *model) BeliBarang() bool{
	if IsPembelian(*pembeli) == true {
		pembeli.CountPembelian++
		return true
	} else {
		return false
	}
}


func (pembeli *model) AmbilBarang() bool{
	if IsPembelian(*pembeli) == true {
		pembeli.CountItem++
		return true
	} else {
		return false
	}
}

func Print(pembeli model){
	fmt.Printf("%d %s %d %d %d",pembeli.ID, pembeli.Name, pembeli.CountItem, pembeli.CountPembelian, pembeli.Status)
}

func main (){
	pembeli := model{1,"nama",1,1,1}
	pembeli.BeliBarang()
	pembeli.AmbilBarang()
	Print(pembeli)
	//Print(pembeli)
}
