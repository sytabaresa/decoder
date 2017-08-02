package decoder

import "fmt"

func ToNibles(x byte) (int, int) {
	y := int(x)
	return y >> 4, 0x0F & y
}

func ToBCD(x byte) (int, error) {
	u, l := ToNibles(x)
	if u > 9 || l > 9 {
		return 0, fmt.Errorf("No esta bien codificado en BCD: %x", x)
	}
	y := int(x)
	return (y>>4)*10 + (0x0F & y), nil
}
