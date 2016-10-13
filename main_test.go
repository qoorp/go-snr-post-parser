package snrpost

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestValidUnmarshal(t *testing.T) {
	f, err := os.Open("/home/dlq/qoorp/baslyft-800,840/1line")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := bufio.NewReader(f)
	line, _, err := rd.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	var postID PostID
	UnmarshalPostID(line, &postID)
	var p800 Post800
	UnmarshalData(line, &p800)
	fmt.Println(p800)
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
