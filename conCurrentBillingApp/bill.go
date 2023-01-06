package main

import (
	"fmt"
	"os"
)

type Bill struct {
	Name       string
	Items      []Item
	Tip        float64
	total      float64
	billString string
}

type Item struct {
	Name  string
	Price float64
}

func (this *Bill) GetTotal() {
	this.total = 0 + this.Tip
	for _, val := range this.Items {
		this.total += val.Price
	}
}

func (this *Bill) GenerateBill() {
	fileName := this.Name + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_, err2 := file.WriteString(this.billString)
	if err2 != nil {
		panic(err)
	}
}

func (this *Bill) ToString() {
	str := fmt.Sprintf("\t\tBill Name \t %v\n\n Item\tPrice", this.Name)

	for _, val := range this.Items {
		str += fmt.Sprintf("\n %v\t₹ %v", val.Name, val.Price)
	}
	str += fmt.Sprintf("\n Tip\t₹ %v", this.Tip)
	str += "\n ------------------------------------------"
	str += fmt.Sprintf("\n Total\t₹ %v", this.total)
	str += "\n ------------------------------------------"
	this.billString = str
}

func (bill *Bill) addTip() {
	fmt.Print("Enter Tip :")
	fmt.Scan(&bill.Tip)
}

func (bil *Bill) createNewBill() {
	fmt.Println("Create a new bill name :")
	fmt.Scan(&bil.Name)
}

func Newbill() *Bill {
	return &Bill{}
}
