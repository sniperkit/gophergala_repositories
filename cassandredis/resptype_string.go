// generated by stringer -type=respType; DO NOT EDIT

package cassandredis

import "fmt"

const _respType_name = "respSimpleStringrespErrorrespIntegerrespBulkStringrespArrayrespUnknown"

var _respType_index = [...]uint8{16, 25, 36, 50, 59, 70}

func (i respType) String() string {
	i -= 1
	if i >= respType(len(_respType_index)) {
		return fmt.Sprintf("respType(%d)", i+1)
	}
	hi := _respType_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _respType_index[i-1]
	}
	return _respType_name[lo:hi]
}
