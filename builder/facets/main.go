package main

import "fmt"

/*
Autor: Francisco Nogales
Descripcion: El Builder es un patron de diseño que permite crear un objeto
	a partir de una serie de pasos que se deben seguir para obtener el objeto final.

Observaciones Personales: En el facets es importante que el builder principal, sea una estructura,
	y esta sea implementada en un los otros builders.
*/

type Person struct {
	StreetAddress string
	PostalCode    string
	City          string
	CompanyName   string
	Position      string
	AnualIncome   int
}

// PersonBuilder es la estructura principal del builder.
type PersonBuilder struct {
	person *Person
}

// Lives es la inicializacion del builder de la vivienda de la persona. Estos Setters retornan el builder para poder ser llamados en cadena.
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{
		PersonBuilder: *b,
	}
}

// Works es la inicializacion del builder del trabajo de la persona
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{
		PersonBuilder: *b,
	}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	it.person.StreetAddress = street
	return it
}

func (it *PersonAddressBuilder) WithPostalCode(code string) *PersonAddressBuilder {
	it.person.PostalCode = code
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonJobBuilder) WorksAsA(position string) *PersonJobBuilder {
	it.person.Position = position
	return it
}

func (it *PersonJobBuilder) At(company string) *PersonJobBuilder {
	it.person.CompanyName = company
	return it
}

func (it *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	it.person.AnualIncome = annualIncome
	return it
}

func (it *PersonBuilder) Build() *Person {
	return it.person
}

func main() {
	var pb = NewPersonBuilder()
	pb.Lives().At("campeche 12").In("Caborca").WithPostalCode("83000")
	pb.Works().At("Yofio").WorksAsA("Software Developer").Earning(20000)
	person := pb.Build()
	fmt.Println(*person)
}
