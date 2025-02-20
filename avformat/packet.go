// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/mgr9525/goav/avcodec"
)

func toCPacket(pkt *avcodec.Packet) *C.struct_AVPacket {
	return (*C.struct_AVPacket)(unsafe.Pointer(pkt))
}

func fromCPacket(pkt *C.struct_AVPacket) *avcodec.Packet {
	return (*avcodec.Packet)(unsafe.Pointer(pkt))
}
