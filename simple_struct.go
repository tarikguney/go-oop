package main

import ("fmt")

type Planet struct{
	Name string 
}

// Defining a method on Planet.
func (p Planet) GetName() string{
	return p.Name
}

func (p Planet) SetName(name string){
	p.Name = name
}

func main(){
	var planet Planet = Planet{Name:"World"}
	new_planet := planet
	new_planet.SetName("Mars")

	fmt.Println(new_planet.Name)
	fmt.Println(planet.GetName())
}