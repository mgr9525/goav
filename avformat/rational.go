// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#include <libavutil/avutil.h>
import "C"
import "github.com/mgr9525/goav/avcodec"

func newRational(r C.struct_AVRational) avcodec.Rational {
	return avcodec.NewRational(int(r.num), int(r.den))
}
