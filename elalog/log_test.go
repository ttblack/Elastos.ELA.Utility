package elalog

import "testing"

func TestStdout(t *testing.T) {
	Stdout.Info("Stdout print Info log")
	Stdout.Infof("Stdout print Infof %v", Stdout)

	Stdout.Warn("Stdout print Warn log")
	Stdout.Warnf("Stdout print Warnf %v", Stdout)

	Stdout.Error("Stdout print Error log")
	Stdout.Errorf("Stdout print Errorf %v", Stdout)

	Stdout.Fatal("Stdout print Fatal log")
	Stdout.Fatalf("Stdout print Fatalf %v", Stdout)

	Stdout.Trace("Stdout print Trace log")
	Stdout.Tracef("Stdout print Tracef %v", Stdout)

	Stdout.Debug("Stdout print Debug log")
	Stdout.Debugf("Stdout print Debugf %v", Stdout)
}