package data

// [username], password
var Users = map[string]string {
	"yash": "pass",
	"tush": "password",
}

var Books = map[int]string {
	1: "Game of thrones",
	2: "Animal farm",
	3: "Crime and punishment",
}

var Orders = make(map[string][]int)

var NumberOfRequests int
