// generated by stringer -type=Region ./types.go; DO NOT EDIT

package band

import "fmt"

const _Region_name = "NorteNordesteSudesteCentroOesteSul"

var _Region_index = [...]uint8{0, 5, 13, 20, 31, 34}

func (i Region) String() string {
	if i < 0 || i+1 >= Region(len(_Region_index)) {
		return fmt.Sprintf("Region(%d)", i)
	}
	return _Region_name[_Region_index[i]:_Region_index[i+1]]
}