package main

import "testing"

func TestPasswordChecker(t *testing.T) {
	testData := map[string]bool{
		"111111": false,
		"123456": false,
		"223450": false,
		"223459": true,
		"112233": true,
		"123444": false,
		"111122": true,
	}
	for password, expected := range testData {
		actual := checkPassword(password)
		if actual != expected {
			t.Errorf("Password %s was supposed to check %t but was %t", password, expected, actual)
		}
	}
}
