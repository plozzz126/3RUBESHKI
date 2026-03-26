package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

func newStd(name string, age int) Student{
	return Student{
		Name: name,
		Age: age,
	}
}

func ShowSt(s []Student){
	for _, v := range s {
		fmt.Println(v.Name, v.Age)
	}
}

func Search(students []Student, name string) *Student{
	for i := range students{
		if students[i].Name == name{
			return &students[i]
		}
	}
	return nil
}

func del(students []Student, name string) []Student{
	for i := range students {
		if students[i].Name == name {
			return append(students[:i], students[i+1:]...)
		}
	}
	return students
}


func main(){
	slm := []Student{
		{Name: "Nurzhan", Age: 40},
		{Name: "Oleg", Age: 22},
		{Name: "Iliya", Age: 4},
	}

	var choice int

	for {
		fmt.Println("1. Показать всех")
		fmt.Println("2. Найти")
		fmt.Println("3. Удалить")
		fmt.Println("4. Добавить")
		fmt.Println("0. Выйти")
		
		switch choice {
			case 1:
				ShowSt(slm)
			case 2:
				var name string 
				fmt.Println("Введите имя: ")
				fmt.Scan(&name)

				s := Search(slm, name)
				if s != nil{
					fmt.Println("пользователь найден")
				}else {
					fmt.Println("не найден")
				}
			case 3:
				var name string 
				fmt.Println("Введите имя: ")
				fmt.Scan(&name)

				slm = del(slm, name)
			case 4:
				var name string 
				var age int

				fmt.Println("Имя: ")
				fmt.Scan(&name)
				fmt.Println("возраст: ")
				fmt.Scan(&age)

				s1 := newStd(name, age)
				slm = append(slm, s1)

			case 0:
				return
		}
	}
}


// type cit struct{
// 	City string
// }
// type st struct{
// 	name string
// 	age int
//  cit City 
// } 
// func (s *st) SetName(name string){
// 	s.name = name
// } вот типо пример с приватным просто эт было написано в конце что бы применить там менять надо
// и с композицией 






