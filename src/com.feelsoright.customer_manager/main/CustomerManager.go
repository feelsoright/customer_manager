package main

import (
	"com.feelsoright.customer_manager/service"
	"com.feelsoright.customer_manager/view"
)

func main() {
	customerView := view.CustomerView{
		CustomerService: service.NewCustomerService(),
	}
	// 初始化customerService
	customerView.IndexMenu()
}
