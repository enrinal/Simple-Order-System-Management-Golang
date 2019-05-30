package usecase

import (
	"testing"
)


func TestBeliBarang(t *testing.T) {

	areaTests := []struct{
		pembeli model
		want bool
	}{
		{model{1,"Enrinal",1,5,2},true},
		{model{1,"Enrinal",1,20,1},false},
		{model{1,"Enrinal",1,10,2},true},
	}

	for _, tt := range areaTests {
        got := tt.pembeli.BeliBarang()
        if got != tt.want {
            t.Errorf("not Passed")
        }
    }
}

func TestAmbilBarang(t *testing.T) {

	areaTests := []struct{
		pembeli model
		want bool
	}{
		{model{1,"Enrinal",1,5,2},true},
		{model{1,"Enrinal",1,20,1},false},
		{model{1,"Enrinal",1,10,2},true},
	}

	for _, tt := range areaTests {
        got := tt.pembeli.AmbilBarang()
        if got != tt.want {
            t.Errorf("not Passed")
        }
    }
}

// func TestProsesPembelian(t *testing.T){
// 	areaTests :=[]struct{
// 		pembeli model
// 		want bool
// 	}{
// 	{model{1,"Enrinal",1,5,2},true},
// 	{model{1,"Enrinal",1,20,1},false},
// 	{model{1,"Enrinal",1,10,2},true},
// }
// for _, tt := range areaTests {
//         got := AmbilBarang(tt.pembeli)
//         if got != tt.want {
//             t.Errorf("not Passed")
//         }
//     }
// }





