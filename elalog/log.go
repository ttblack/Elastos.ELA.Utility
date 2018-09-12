package elalog

func init() {
	Disabled = &disableLog{}
	Stdout = newStdout()
}
