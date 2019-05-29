package usecase

import "model"


type Pembeli struct {
	id int
	name string
	countPembelian int
	countItem int
	status int
}

func IsPembelian (pembeli Pembeli) bool{
	if pembeli.status == 1 && pembeli.countPembelian <=5 && pembeli.countItem <=10{
		return true
	}else if pembeli.status == 2 && pembeli.countItem <=10{
		return true
	}else {
		return false
	}
	
}
func BeliBarang (pembeli Pembeli) bool {
	if IsPembelian (pembeli) == true {
		pembeli.countPembelian++
		return true
	}else{
		return false
	}
}

func AmbilBarang (pembeli Pembeli) bool{
	if IsPembelian (pembeli) == true{
		pembeli.countItem++
		return true
	}else{
		return false
	}
}