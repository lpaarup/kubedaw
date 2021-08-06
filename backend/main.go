package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"net"

	pb "github.com/lpaarup/kubedaw/backend/api"
	"github.com/sirupsen/logrus"
	grcp "google.golang.org/grpc"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	logrus.Infof("Listening on port %d", *port)
	s := grcp.NewServer()
	pb.RegisterAudioDataServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedAudioDataServer
	t int
}

func (s *server) Request(ctx context.Context, req *pb.DataRequest) (*pb.Data, error) {
	logrus.Debugf("requesting %d frames, %d channels, %d sampleRate, %d maxAmplitude",
		req.NumFrames, req.NumChannels, req.SampleRate, req.MaxAmplitude)
	b := new(bytes.Buffer)
	b.Grow(int(req.NumFrames * req.NumChannels * 2))
	w := (math.Pi * 2.0 * 440.0) / float64(req.SampleRate)
	for frame := 0; frame < int(req.NumFrames); frame++ {
		val := math.Round(float64(req.MaxAmplitude) * math.Sin(w*float64(s.t)))
		s.t++
		for channel := 0; channel < int(req.NumChannels); channel++ {
			binary.Write(b, binary.LittleEndian, int16(val))
		}
	}
	return &pb.Data{Audio: b.Bytes()}, nil
}
