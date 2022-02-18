package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type IPerson interface {
	Hello() string
	Data() string
}

type Person struct {
	Name    string
	Age     int
	Married bool
}

func NewPerson(name string, age int, married bool) IPerson {
	return &Person{
		Name:    name,
		Age:     age,
		Married: married,
	}
}

func (p *Person) Hello() string {
	return "Hello " + p.Name
}

func (p *Person) Data() string {
	return "Name: " + p.Name + ", Age: " + strconv.Itoa(p.Age) + ", Married: " + strconv.FormatBool(p.Married)
}

func TestPerson(t *testing.T) {
	person := NewPerson("Billy", 23, false)
	person2 := NewPerson("Oyen", 21, false)

	data := person.Data()
	data2 := person2.Data()

	fmt.Println(data2)

	assert.Equal(t, "Name: Billy, Age: 23, Married: false", data)
}
