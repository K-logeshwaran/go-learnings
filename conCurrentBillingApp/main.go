package main

import (
	"fmt"
)

func main() {
	billChannel := make(chan Bill, 10)
	for {
		newBill := Newbill()
		newBill.createNewBill()
		option := " "
		for option != "s" {
			// getting user's option
			fmt.Print("Choose options (a - to add item,s - to save bill,t - to add tip) :")
			fmt.Scan(&option)

			switch option {
			case "a":
				fmt.Println("Add bill")

				// adding new items to bill
				newBill.Items = append(newBill.Items, addItem())
				fmt.Println("Item added")

			case "t":
				fmt.Println("Tip")
				newBill.addTip()
				fmt.Println("Tip added")

			case "s":
				billChannel <- *newBill
				go saveBill(billChannel)

			default:
				fmt.Println("Please  enter a valid option")

			}
		}
		go saveBill(billChannel)
	}

}

// func to createNewBill

// func to add items
func addItem() Item {
	// local item variable
	var locItm Item

	//getting item name
	fmt.Print("Enter Item Name :")
	fmt.Scan(&locItm.Name)

	//getting item price
	fmt.Print("Enter Item Price:")
	fmt.Scan(&locItm.Price)
	return locItm
}

func saveBill(data chan Bill) {
	//time.Sleep(2 * time.Second)
	for b := range data {
		fmt.Println("Save")
		b.GetTotal()
		b.ToString()
		b.GenerateBill()
		fmt.Print("\n", b)
		fmt.Println("\nBilll generated successfully at " + b.Name + ".txt")
	}

}
