package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂构造
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 工厂构造
func NewCustomerNotId(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// toString
func (c Customer) ToString() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}
