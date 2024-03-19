package snrpost

import "testing"

func TestUnmarshalWholeNumber(t *testing.T) {
	s := []byte("000000000000001000")
	var p struct {
		WholeNumber string `snr:"18"`
	}
	_, err := unmarshal(s, &p)
	if err != nil {
		t.Fatalf("Unmarshal failed %v", err)
	}
	if p.WholeNumber != "000000000000001000" {
		t.Fatalf("WholeNumber doesn't match expected value")
	}
}

func TestUnmarshalDecimalNumber(t *testing.T) {
	s := []byte("000000000000001000")
	var p struct {
		DecimalNumber string `snr:"12,6"`
	}
	_, err := unmarshal(s, &p)
	if err != nil {
		t.Fatalf("Unmarshal failed %v", err)
	}
	if p.DecimalNumber != "000000000000.001000" {
		t.Fatalf("DecimalNumber doesn't match expected value")
	}
}
