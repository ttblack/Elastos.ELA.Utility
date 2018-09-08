package p2p

import "os"

type logWriter struct{}

func (logWriter) Write(p []byte)(n int, err error) {
	os.Stdout.Write(p)

}
