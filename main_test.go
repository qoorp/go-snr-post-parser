package snrpost

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestValidUnmarshal(t *testing.T) {
	f, err := os.Open("./test/test_data.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := bufio.NewReader(f)
	// Skip post 100
	rd.ReadLine()
	line, _, err := rd.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	var av AviserPost
	UnmarshalAviserPost(line, &av)
	fmt.Println(av)
	var p800 Post800
	UnmarshalData(line, &p800)
	fmt.Println(p800.Firma)
}

func TestNilUnmarshal(t *testing.T) {
	_, err := unmarshal(nil, nil)
	if err == nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestNonPtrUnmarshal(t *testing.T) {
	var p Post800
	unmarshal(nil, p)
}
