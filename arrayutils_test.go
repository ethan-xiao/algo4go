package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLinerArray2SymmetricTree(t *testing.T) {
	b, err := json.Marshal(linerArray2SymmetricTree([]int{1, 2, 3}))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(string(b), "{\"val\":1,\"left\":{\"val\":2},\"right\":{\"val\":3}}") {
		t.Fatal()
	}

	b, err = json.Marshal(linerArray2SymmetricTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(string(b), "{\"val\":1,\"left\":{\"val\":2,\"left\":{\"val\":4},\"right\":{\"val\":5,\"left\":{\"val\":8,\"left\":{\"val\":14},\"right\":{\"val\":15}},\"right\":{\"val\":9}}},\"right\":{\"val\":3,\"left\":{\"val\":6,\"left\":{\"val\":10},\"right\":{\"val\":11}},\"right\":{\"val\":7,\"left\":{\"val\":12},\"right\":{\"val\":13}}}}") {
		t.Fatal()
	}

	b, err = json.Marshal(linerArray2SymmetricTree([]int{1, 2, 2, 3, 4, 4, 3}))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(string(b), "{\"val\":1,\"left\":{\"val\":2,\"left\":{\"val\":3},\"right\":{\"val\":4}},\"right\":{\"val\":2,\"left\":{\"val\":4},\"right\":{\"val\":3}}}") {
		t.Fatal()
	}
}
