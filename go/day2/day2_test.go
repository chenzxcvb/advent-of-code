package day2

import "testing"

func TestDay2_Part1(t *testing.T) {
	d := New()
	d.Open()

	answer := d.Part1()

	// 13628
	if answer != "15" {
		t.Fatalf("Expected 15 but got %s", answer)
	}
	d.Close()
}
