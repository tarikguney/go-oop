package main

import ("fmt")

type Planet struct{
	Name string 
}

// Defining a method on Planet.
func (p *Planet) GetName() string{
	return p.Name
}
// Similar to the extension methods like C#
// these methods can come from other packages too.
func (p *Planet) SetName(name string){
	p.Name = name
}

func main(){
	var planet Planet = Planet{Name:"World"}

	fmt.Println(planet.GetName())
	// Seeing any problem?
	SetName(planet)
	fmt.Println(planet.Name)
}

func SetName(p Planet){
	p.SetName("Jupiter")
}