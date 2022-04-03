package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	category int
	quantity int
	itemCost float64
}

var category []string

var shoppingList = map[string]Item{}

func printList() {
	for k, v := range shoppingList {
		fmt.Println("Category:", category[v.category], "- Item:", k, "Quantity:", v.quantity, "Unit Cost:", v.itemCost)
	}
}

func printCurrent() {
	fmt.Println("Print Current Data.")
	if len(shoppingList) > 0 {
		for i, j := range shoppingList {
			fmt.Printf("%s - {%d %d %.2f}\n", i, j.category, j.quantity, j.itemCost)
		}
	} else {
		fmt.Println("No data found!")
	}

}

func generatReport() {
	var input int
	fmt.Println("Generate Report")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")
	fmt.Println("\nChoose your report: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		fmt.Println("Total cost by Category.")
		for idx, name := range category {
			//	fmt.Println(idx, name)
			var categoryTotal float64 = 0
			for _, itemStruct := range shoppingList {
				if itemStruct.category == idx {
					//	fmt.Println(itemName, itemStruct)
					//	fmt.Println(itemStruct.quantity, itemStruct.itemCost)
					totalCost := float64(itemStruct.quantity) * itemStruct.itemCost
					//	fmt.Println(totalCost)
					categoryTotal = categoryTotal + totalCost
				}
			}
			fmt.Println(name, ":", categoryTotal)
		}
	case 2:
		printList()
	case 3:
		mainMenu()

	}
}
func mainMenu() {
	var input int
	fmt.Println("Shopping List Application")
	fmt.Println("=========================")
	fmt.Println("1. View entire shopping list.")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items.")
	fmt.Println("4. Modify Items.")
	fmt.Println("5. Delete Item")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7. Add New Category Name")
	fmt.Println("Select your choice: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		printList()
	case 2:
		generatReport()
	case 3:
		addItems()
	case 4:
		modExistItem()
	case 5:
		deleteItem()
	case 6:
		printCurrent()
	case 7:
		addNewCat()

	}
	mainMenu()
}

func addItems() {
	var input string
	var cat string
	var uni int
	var unitCost float64
	var item Item
	var cat1 int

	fmt.Println("What is the name of your item?")
	fmt.Scanln(&input)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&cat)
	fmt.Println("How many units are there?")
	fmt.Scanln(&uni)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&unitCost)

	for i, v := range category {
		//fmt.Println(i, v)
		if v == cat {
			cat1 = i
		}
	}

	item.category = cat1
	item.quantity = uni
	item.itemCost = unitCost

	//fmt.Println(input)
	//fmt.Println(item)
	shoppingList[input] = item
}
func deleteItem() {
	var input string
	fmt.Println("Enter item name to delete: ")
	fmt.Scanln(&input)
	if _, exist := shoppingList[input]; exist {
		delete(shoppingList, input)
	} else {
		fmt.Println("Item not found. Nothing to delete!")
	}
}
func addNewCat() {
	var input string
	var exist bool = false
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&input)
	if len(input) > 0 {
		for i, v := range category {
			//fmt.Println(i = index , v = category name)
			if v == input {
				fmt.Println(input, "already exist at index", i)
				exist = true
			}
		}
		if exist == false {
			category = append(category, input)
			for idx, value := range category {
				if value == input {
					fmt.Println("New category:", input, "added at index", idx, "!")
				}
			}
		}
	} else {
		fmt.Println("No Input Found!")
	}
}
func modExistItem() {
	var input string
	var input1 string
	var input2 string
	var input3 string
	var input4 string
	var newCategory int
	var newQuantity int
	var newCost float64
	fmt.Println("Which item would you wish to modify?")
	fmt.Scanln(&input)
	item1 := shoppingList[input]
	fmt.Println("Current item name is ", input, "- Category is", category[item1.category], "- Quantity is ", item1.quantity, "- Unit Cost", item1.itemCost)
	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&input1)
	if len(input1) == 0 {
		input1 = input
	}
	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&input2)
	if len(input2) == 0 {
		newCategory = item1.category
	} else {
		value, _ := strconv.Atoi(input2)
		newCategory = value
	}
	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&input3)
	if len(input3) == 0 {
		newQuantity = item1.quantity
	} else {
		value, _ := strconv.Atoi(input3)
		newQuantity = value
	}
	fmt.Println("Enter new Unit cost. Enter for no change.")
	fmt.Scanln(&input4)
	if len(input4) == 0 {
		newCost = item1.itemCost
	} else {
		value, _ := strconv.ParseFloat(input4, 64)
		newCost = value
	}
	fmt.Println(input1, input2, input3, input4)
	if input == input1 {
		fmt.Println("No changes to item name made.")
	} else {
		delete(shoppingList, input)
		shoppingList[input1] = item1
	}
	if newCategory == item1.category {
		fmt.Println("No changes to category made.")
	} else {
		item1.category = newCategory
		if input == input1 {
			shoppingList[input] = item1
		} else {
			shoppingList[input1] = item1
		}
	}
	if newQuantity == item1.quantity {
		fmt.Println("No change to quantity made.")
	} else {
		item1.quantity = newQuantity
		if input == input1 {
			shoppingList[input] = item1
		} else {
			shoppingList[input1] = item1
		}
	}
	if newCost == item1.itemCost {
		fmt.Println("No changes to unit cost made.")
	} else {
		item1.itemCost = newCost
		if input == input1 {
			shoppingList[input] = item1
		} else {
			shoppingList[input1] = item1
		}
	}
}
func main() {
	category = append(category, "Household")
	category = append(category, "Food")
	category = append(category, "Drinks")

	shoppingList["Cups"] = Item{0, 5, 3.00}
	shoppingList["Cake"] = Item{1, 3, 1.00}
	shoppingList["Sprite"] = Item{2, 5, 2.00}
	shoppingList["Fork"] = Item{0, 4, 3.00}
	shoppingList["Bread"] = Item{1, 2, 2.00}
	shoppingList["Plates"] = Item{0, 4, 3.00}
	shoppingList["Coke"] = Item{2, 5, 2.00}

	mainMenu()
}
