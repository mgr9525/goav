package main

import (
	"log"

	"github.com/mgr9525/goav/avcodec"
	"github.com/mgr9525/goav/avdevice"
	"github.com/mgr9525/goav/avfilter"
	"github.com/mgr9525/goav/avformat"
	"github.com/mgr9525/goav/avutil"
	"github.com/mgr9525/goav/swresample"
	"github.com/mgr9525/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
