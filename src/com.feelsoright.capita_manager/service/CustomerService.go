package service

import (
	"com.feelsoright.capita_manager/model"
)

// 客户增删改查service
type CustomerService struct {
	customers []model.Customer
	// 数量
	CustomerNum int
}

// 工厂模式
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customer := model.NewCustomer(1, "测试001", "男", 18, "13800138088", "708838862@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 获取所有客户
func (cs *CustomerService) GetAllCustomer() []model.Customer {
	return cs.customers
}

// 添加客户
func (cs *CustomerService) AddCustomer(customer model.Customer) bool {
	if customer.Id == 0 {
		cs.CustomerNum++
		customer.Id = cs.CustomerNum
	}
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据ID删除客户
func (cs *CustomerService) DelCustomer(id int) bool {
	for index, val := range cs.customers {
		if val.Id == id {
			cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
			cs.CustomerNum--
			return true
		}
	}
	return false
}

// 根据ID查询客户
func (cs *CustomerService) FindCustomerById(id int) model.Customer {
	for _, val := range cs.customers {
		if val.Id == id {
			return val
		}
	}
	return model.Customer{}
}

// 更新客户信息
func (cs *CustomerService) UpdateCustomer(customer model.Customer) bool {
	for index, val := range cs.customers {
		if val.Id == customer.Id {
			cs.customers[index] = customer
			return true
		}
	}
	return false
}
