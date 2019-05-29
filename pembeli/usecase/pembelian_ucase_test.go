package usecase

import "testing"


func TestBeliBarang(t *testing.T) {

	areaTests := []struct{
		pembeli Pembeli
		want bool
	}{
		{Pembeli{1,"Enrinal",1,5,2},true},
		{Pembeli{1,"Enrinal",1,20,1},false},
		{Pembeli{1,"Enrinal",1,10,2},true},
	}

	for _, tt := range areaTests {
        got := BeliBarang(tt.pembeli)
        if got != tt.want {
            t.Errorf("not Passed")
        }
    }
}

func TestAmbilBarang(t *testing.T) {

	areaTests := []struct{
		pembeli Pembeli
		want bool
	}{
		{Pembeli{1,"Enrinal",1,5,2},true},
		{Pembeli{1,"Enrinal",1,20,1},false},
		{Pembeli{1,"Enrinal",1,10,2},true},
	}

	for _, tt := range areaTests {
        got := AmbilBarang(tt.pembeli)
        if got != tt.want {
            t.Errorf("not Passed")
        }
    }
}


// func TestAmbilBarang(t *testing.T) {
// 	pembeli := Pembeli{1,"Enrinal",1,1,1}
//     got := AmbilBarang(pembeli)
//     want := true

//     if got != want {
//         t.Errorf("not Passed")
//     }
// }




