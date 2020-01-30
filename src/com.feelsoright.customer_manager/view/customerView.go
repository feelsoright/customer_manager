package view

import (
	"com.feelsoright.customer_manager/model"
	"com.feelsoright.customer_manager/service"
	"com.feelsoright.customer_manager/utils"
	"fmt"
)

type CustomerView struct {
	CustomerService *service.CustomerService
}

func (cv *CustomerView) showAllCustomer() {
	// 获取客户列表
	customers := cv.CustomerService.GetAllCustomer()
	utils.PrintLine("客户信息列表")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.ToString())
	}
	utils.PrintLine("信息列表结束")
}

func (cv *CustomerView) addCustomer() {
	utils.PrintLine("添加客户信息")
	name := utils.GetStringColumn("姓名")
	gender := utils.GetStringColumn("性别")
	age := utils.GetIntColumn("年龄")
	phone := utils.GetStringColumn("电话")
	email := utils.GetStringColumn("邮箱")
	customer := model.NewCustomerNotId(name, gender, age, phone, email)
	flag := cv.CustomerService.AddCustomer(customer)
	if flag {
		utils.PrintLine("添加信息成功")
		return
	}
	utils.PrintLine("添加信息失败")
}

func (cv *CustomerView) delCustomer() {
	utils.PrintLine("删除客户信息")
	id := utils.GetIntColumn("请输入需要删除客户的ID(-1退出)")
	if id == -1 || !utils.YOrN("确认要删除该客户吗？(y/n)") {
		utils.PrintLine("退出删除信息")
		return
	}

	flag := cv.CustomerService.DelCustomer(id)
	if flag {
		utils.PrintLine("删除信息成功")
		return
	}
	utils.PrintLine("删除信息失败")

}

func (cv *CustomerView) updateCustomer() {
	utils.PrintLine("修改客户信息")
	id := utils.GetIntColumn("请输入需要修改客户的ID(-1退出)")
	customer := cv.CustomerService.FindCustomerById(id)
	if customer.Id == 0 && id != -1 {
		fmt.Println("未查询到该客户信息！")
		return
	}
	if id == -1 || !utils.YOrN("确认要修改该客户信息吗？(y/n)") {
		utils.PrintLine("退出修改信息")
		return
	}
	name := utils.GetStringColumn("姓名")
	gender := utils.GetStringColumn("性别")
	age := utils.GetIntColumn("年龄")
	phone := utils.GetStringColumn("电话")
	email := utils.GetStringColumn("邮箱")
	if name != "" {
		customer.Name = name
	}
	if gender != "" {
		customer.Gender = name
	}
	if gender != "" {
		customer.Age = age
	}
	if phone != "" {
		customer.Phone = phone
	}
	if email != "" {
		customer.Email = email
	}
	customer = model.NewCustomer(id, name, gender, age, phone, email)
	flag := cv.CustomerService.UpdateCustomer(customer)
	if flag {
		utils.PrintLine("修改信息成功")
		return
	}
	utils.PrintLine("修改信息失败")
}

func (cv *CustomerView) IndexMenu() {
	loop := true
	for {
		inputKey := ""
		fmt.Println("--------------客户信息管理系统--------------")
		fmt.Println("              1、 添 加 客 户")
		fmt.Println("              2、 修 改 客 户")
		fmt.Println("              3、 删 除 客 户")
		fmt.Println("              4、 客 户 列 表")
		fmt.Println("              5、 退       出")
		fmt.Print("请选择(1-5):")
		_, _ = fmt.Scanln(&inputKey)
		switch inputKey {
		case "1":
			fmt.Println("当前选择添加客户")
			cv.addCustomer()
		case "2":
			fmt.Println("当前选择修改客户")
			cv.updateCustomer()
		case "3":
			fmt.Println("当前选择删除客户")
			cv.delCustomer()
		case "4":
			fmt.Println("当前选择客户列表")
			cv.showAllCustomer()
		case "5":
			loop = !utils.YOrN("确认要退出系统吗？(y/n)")
		default:
			fmt.Println("您的选择有误，请重新选择！")
		}
		if !loop {
			break
		}
	}
	fmt.Println("感谢您的使用，再见！")
}
