package main

import "fmt"

// Create new stuct
type User struct {
	name     string
	age      int64
	password string
	score    []int
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name1 string) {
	u.name = name1
}

func (u User) getHighScore() int {
	high := 0
	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

func main() {
	user1 := User{"John", 23, "pass", []int{1, 10, 25, 345, 1234, 9, 29}}
	fmt.Println(user1.getHighScore())

}
