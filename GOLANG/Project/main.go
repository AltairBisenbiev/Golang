package main

import (
	"fmt"
	"sort"
)

var id int = 1

type User struct {
	id       int
	Name     string
	Age      int
	password string
	email    string
}
type Data struct {
	users []User
	items []Item
}
type Item struct {
	name  string
	owner User
	price int
	rate  int
}

func (d *Data) signup() {
	fmt.Println("What is your name?")
	var name string
	var age int
	var email string
	var password string
	fmt.Scan(&name)
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("Write your Email:")
	fmt.Scan(&email)
	fmt.Println("Write a password:")
	fmt.Scan(&password)
	newUser := User{id, name, age, password, email}
	id = id + 1
	d.users = append(d.users, newUser)

}
func (d *Data) addItem(a User) {
	fmt.Println("Write a name of item: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Write a price of item: ")
	var price int
	fmt.Scan(&price)
	newItem := Item{name, a, price, 10}
	d.items = append(d.items, newItem)

}
func (d *Data) signin() User {
	var email1 string
	var password1 string
	fmt.Println("Write your Email:")
	fmt.Scan(&email1)
	fmt.Println("Write a password:")
	fmt.Scan(&password1)

	for i := 0; i < len(d.users); i++ {
		if d.users[i].email == email1 && d.users[i].password == password1 {
			return d.users[i]
		}
	}

	ni := User{-1, "", 0, "", ""}
	return ni

}
func (d *Data) search(name string, price int, rate int) []Item {
	result := []Item{}
	for i := 0; i < len(d.items); i++ {
		if name == "none" || name == d.items[i].name {
			if price <= d.items[i].price {
				if rate <= d.items[i].rate {
					result = append(result, d.items[i])
				}
			}
		}

	}
	return result
}
func equals(a Item, b Item) bool {
	if a.name == b.name && a.owner.Name == b.owner.Name && a.price == b.price {
		return true
	} else {
		return false
	}

}
func (d *Data) giveRate(a Item, rate int) {

	for i := 0; i < len(d.items); i++ {
		if equals(a, d.items[i]) {
			d.items[i].rate = (d.items[i].rate + rate) / 2
		}
	}

}
func main() {
	dat := &Data{}
	var command string
	var cur User
	user := false
	for command != "Exit" {
		if user == false {
			fmt.Println("1. Login ->")
			fmt.Println("2. SignUp ->")
			fmt.Println("3. Exit <-")
			fmt.Scan(&command)

			if command == "Login" {
				cur = dat.signin()
				if cur.id != -1 {
					user = true
				}

			} else if command == "SignUp" {
				dat.signup()
			}

		} else {
			fmt.Println("---- Welcome! ----")
			fmt.Println("1. Add (item) ")
			fmt.Println("2. Search (item) ")
			fmt.Println("3. Exit")
			if command == "Add" {
				dat.addItem(cur)
			} else if command == "Search" {
				fmt.Println("---- Welcome! ----")
				var name string
				var price int
				var rate int
				fmt.Println("Write a name of item that u want to find")
				fmt.Println("(If u don't want write 'none')")
				fmt.Scan(&name)
				fmt.Println("Write a 'x< price' of item that u want to find")
				fmt.Println("(If u don't want, write '0')")
				fmt.Scan(&price)
				fmt.Println("Write a rating of item that u want to find")
				fmt.Println("(If u don't want, write '0')")
				fmt.Scan(&rate)

				result := dat.search(name, price, rate)
				sort.SliceStable(result, func(i, j int) bool {
					return result[i].name < result[j].name
				})
				sort.SliceStable(result, func(i, j int) bool {
					return result[i].price < result[j].price
				})
				sort.SliceStable(result, func(i, j int) bool {
					return result[i].rate < result[j].rate
				})
				for i := 0; i < len(result); i++ {
					fmt.Println(i, result[i].name, result[i].price, result[i].rate)
				}
				fmt.Println("Choose id of item:")
				var idd int
				fmt.Scan(idd)
				fmt.Println("Choose rate to give to item:")
				var rating int
				fmt.Scan(rating)
				dat.giveRate(result[id], rating)
			}
			fmt.Scan(&command)
		}

	}

}
