// Code generated by "string2enum -linecomment -type Linkage /home/u/Desktop/go/src/github.com/llir/llvm/ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

const _Linkage_name = "noneappendingavailable_externallycommoninternallinkoncelinkonce_odrprivateweakweak_odrexternalextern_weak"

var _Linkage_index = [...]uint8{0, 4, 13, 33, 39, 47, 55, 67, 74, 78, 86, 94, 105}

func LinkageFromString(s string) enum.Linkage {
	if len(s) == 0 {
		return 0
	}
	for i := range _Linkage_index[:len(_Linkage_index)-1] {
		if s == _Linkage_name[_Linkage_index[i]:_Linkage_index[i+1]] {
			return enum.Linkage(i)
		}
	}
	panic(fmt.Errorf("unable to locate Linkage enum corresponding to %q", s))
}
