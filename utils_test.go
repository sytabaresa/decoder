package decoder

import "testing"

func TestToNibles(t *testing.T) {
	test := map[byte]([]int){
		0x12: []int{1, 2},
		0x1F: []int{1, 15},
		0xD2: []int{13, 2},
		0xAF: []int{10, 15},
	}

	for k, v := range test {
		a, b := ToNibles(k)
		if a != v[0] || b != v[1] {
			t.Errorf("%v != %v,%v", k, a, b)
		}
	}
}

func TestToBCD(t *testing.T) {
	a, err := ToBCD(0x31)
	if a != 31 || err != nil {
		t.Errorf("%v is not 31", a)
	}
	b, err := ToBCD(0x9F)
	if b != 0 || err == nil {
		t.Errorf("mal validado valor %v, error %s", b, err)
	}
}
