package tools

import (
	"testing"
)

var hasher = NewHasher()

func TestHasher_Hash(t *testing.T) {
	values := []string{
		"HashTest1",
		"",
		"3",
		" ",
		"hashtest5?",
	}

	for _, test_case := range values {
		_, err := hasher.Hash(test_case)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestHasher_Compare(t *testing.T) {
	values := []string{
		"CompareTest1",
		"",
		"3",
		" ",
		"comparetest5?",
	}

	for _, test_case := range values {
		err := hasher.Compare(test_case, "fake_hash")
		if err == true {
			t.Errorf("Error: wrong behaviour of Compare method")
		}
	}
}

func TestHasher(t *testing.T) {
	// TC stands for TestCases
	type TC struct {
		str string
	}

	values := []TC{
		{"testCase1"},
		{""},
		{"12345678"},
		{"testcase4"},
		{" "}}

	for _, test_case := range values {
		hashed_s, _ := hasher.Hash(test_case.str)
		is_equal := hasher.Compare(test_case.str, hashed_s)

		if !is_equal {
			t.Errorf("Error: passwords are supposed to be equal. %s is not hash of given string: %s", hashed_s, test_case.str)
		}

		mutated_str := test_case.str + "mutation"
		is_equal = hasher.Compare(mutated_str, hashed_s)
		
		if is_equal {
			t.Errorf("Error: passwords are not supposed to be equal. %s is somehow hash of mutated string: %s", hashed_s, mutated_str)
		}
	}
}
