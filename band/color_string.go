// generated by stringer -type=Color ./types.go; DO NOT EDIT

package band

import "fmt"

const _Color_name = "RedGreenYellow"

var _Color_index = [...]uint8{0, 3, 8, 14}

func (i Color) String() string {
	if i < 0 || i+1 >= Color(len(_Color_index)) {
		return fmt.Sprintf("Color(%d)", i)
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}
