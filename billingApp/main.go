package main

import "fmt"

var NewBill Bill

func main() {
	createNewBill(&NewBill)
	option := " "
	for option != "s" {
		// getting user's option
		fmt.Print("Choose options (a - to add item,s - to save bill,t - to add tip) :")
		fmt.Scan(&option)

		switch option {
		case "a":
			fmt.Println("Add bill")

			// adding new items to bill
			NewBill.Items = append(NewBill.Items, addItem())
			fmt.Println("Item added")

		case "t":
			fmt.Println("Tip")
			addTip(&NewBill)
			fmt.Println("Tip added")

		case "s":
			fmt.Println("Save")
			NewBill.GetTotal()
			NewBill.ToString()
			NewBill.GenerateBill()
			fmt.Print("\n", NewBill)
			fmt.Println("\nBilll generated successfully at " + NewBill.Name + ".txt")

		default:
			fmt.Println("Please  enter a valid option")

		}
	}

}

// func to createNewBill
func createNewBill(bil *Bill) {
	fmt.Print("Create a new bill name :")
	fmt.Scan(&bil.Name)
}

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

func addTip(bill *Bill) {
	fmt.Print("Enter Tip :")
	fmt.Scan(&bill.Tip)
}
