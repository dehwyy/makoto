package utils

import (
	"testing"
)

var hasher = NewHasher()

func TestHash(t *testing.T) {
	test_cases := []string{"HashTest1", "", "3", " ", "hashtest5?"}

	for _, test_case := range test_cases {
		hashed_s, err := hasher.Hash(test_case)

		// I guess that `hashed_s` should always be kinda large (~at least 20 symbols)
		//  for example `12232ceb-f5a5-458f-ab19-ba6ffc938256`
		if len(hashed_s) == 0 {
			t.Errorf("Error: hash should not be empty")
		}

		if err != nil {
			t.Error(err)
		}
	}
}

// `Compare` fn would always return falsy value in this test_fn
func TestCompare(t *testing.T) {
	test_cases := []string{"CompareTest1", "", "3", " ", "comparetest5?"}

	for _, test_case := range test_cases {

		// is_equal should always be `false` as the compared string is always wrong
		is_equal := hasher.Compare(test_case, "fake_hash")
		if is_equal {
			t.Errorf("Error: wrong behaviour of Compare method")
		}
	}
}

func TestHasher(t *testing.T) {
	test_cases := []string{"test_case", "", "123", "CASE", " "}

	for _, test_case := range test_cases {
		hashed_s, _ := hasher.Hash(test_case)
		is_equal := hasher.Compare(test_case, hashed_s)

		// should be equal
		if !is_equal {
			t.Errorf("Error: passwords are supposed to be equal. %s is not hash of given string: %s", hashed_s, test_case)
		}

		mutated_s := test_case + "mut"
		is_equal = hasher.Compare(mutated_s, hashed_s)

		// should NOT be equal as the strings are different
		if is_equal {
			t.Errorf("Error: passwords are not supposed to be equal. %s is somehow hash of mutated string: %s", hashed_s, mutated_s)
		}
	}
}
