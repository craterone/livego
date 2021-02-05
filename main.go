package main

import (
	"fmt"
	"github.com/gwuhaolin/livego/protocol/rtmp"
	"net"
	"path"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

var VERSION = "master"

var rtmpAddr string

func startRtmp(stream *rtmp.RtmpStream) {
	rtmpAddr = ":1935"

	rtmpListen, err := net.Listen("tcp", rtmpAddr)
	if err != nil {
		log.Fatal(err)
	}

	var rtmpServer *rtmp.Server

	rtmpServer = rtmp.NewRtmpServer(stream, nil)

	defer func() {
		if r := recover(); r != nil {
			log.Error("RTMP server panic: ", r)
		}
	}()
	log.Info("RTMP Listen On ", rtmpAddr)
	rtmpServer.Serve(rtmpListen)
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("livego panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()

	log.Infof(`
     _     _            ____       
    | |   (_)_   _____ / ___| ___  
    | |   | \ \ / / _ \ |  _ / _ \ 
    | |___| |\ V /  __/ |_| | (_) |
    |_____|_| \_/ \___|\____|\___/ 
        version: %s
	`, VERSION)

	stream := rtmp.NewRtmpStream()

	startRtmp(stream)
}
