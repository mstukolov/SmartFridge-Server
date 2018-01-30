package main

import "reflect"

type Object interface {
	print()
}

type person struct{
	id int
	name string
}
func (p *person) print(){
	println(p.name)
}

type sportsmen struct{
	id int
	name string
}
func (s *sportsmen) print(){
	println(s.name)
}

func main(){
	p1 := person{id:1, name:"max"}
	s1 := sportsmen{id:2, name:"sveta"}
	toString(&p1)
	toString(&s1)
}

func toString(a Object){
	val := reflect.ValueOf(a).Elem()
	println(val.String())
	a.print()
}
