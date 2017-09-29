package main

import ("fmt")

type Planet struct{
	Name string 
}

// Defining a method on Planet.
func (p Planet) GetName() string{
	return p.Name
}

func main(){
	planet := Planet{Name:"World"}
	fmt.Println(planet.GetName())
}