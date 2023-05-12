package linuxCommand

import (
	"bytes"
)

type Server struct {
	UnimplementedRigServiceServer
	UnimplementedFortuneServiceServer
	UnimplementedCowServiceServer
	UnimplementedFigletServiceServer
	UnimplementedToiletServiceServer
	UnimplementedXkcdpassServiceServer
}

type Command struct {
	command string
	args	[]string
	out bytes.Buffer
	err bytes.Buffer
	path string
}
