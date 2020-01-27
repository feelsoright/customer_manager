package main

import (
	"com.feelsoright.capita_manager/service"
	"com.feelsoright.capita_manager/view"
)

func main() {
	customerView := view.CustomerView{
		CustomerService: service.NewCustomerService(),
	}
	// 初始化customerService
	customerView.IndexMenu()
}
