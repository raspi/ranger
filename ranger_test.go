package ranger

import (
	"testing"
)

func TestGetIntegerRanges(t *testing.T) {
	nums := []int8{-1, 0, 1}
	want := Range[int8]{Start: -1, End: 1}

	ranges, err := GetIntegerRanges(nums)
	if err != nil {
		t.Fail()
	}

	got := ranges[0]

	if !(got.Start == want.Start && got.End == want.End) {
		t.Fail()
	}

}

func TestGetIntegerRangesSingle(t *testing.T) {
	nums := []int8{42}
	want := Range[int8]{Start: 42, End: 42}

	ranges, err := GetIntegerRanges(nums)
	if err != nil {
		t.Fail()
	}

	got := ranges[0]

	if !(got.Start == want.Start && got.End == want.End) {
		t.Fail()
	}

}

func TestGetIntegerRangesSingleAndRange(t *testing.T) {
	nums := []int8{42, 10, 11, 12}

	wantSingle := Range[int8]{Start: 42, End: 42}
	want := Range[int8]{Start: 10, End: 12}

	ranges, err := GetIntegerRanges(nums)
	if err != nil {
		t.Fail()
	}

	// 10-12
	got := ranges[0]

	if !(got.Start == want.Start && got.End == want.End) {
		t.Fail()
	}

	// 42
	gotSingle := ranges[1]

	if !(gotSingle.Start == wantSingle.Start && gotSingle.End == wantSingle.End) {
		t.Fail()
	}

}

func TestGetIntegerRangesSingleAndRange2(t *testing.T) {
	nums := []int8{10, 11, 12, 42}

	wantSingle := Range[int8]{Start: 42, End: 42}
	want := Range[int8]{Start: 10, End: 12}

	ranges, err := GetIntegerRanges(nums)
	if err != nil {
		t.Fail()
	}

	// 10-12
	got := ranges[0]

	if !(got.Start == want.Start && got.End == want.End) {
		t.Fail()
	}

	// 42
	gotSingle := ranges[1]

	if !(gotSingle.Start == wantSingle.Start && gotSingle.End == wantSingle.End) {
		t.Fail()
	}

}
