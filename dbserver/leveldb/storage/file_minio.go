package storage

type dirInfo

type file struct {
	pfd        poll.FD
	name       string
	dirinfo    *dirInfo // nil unless directory being read
	appendMode bool     // whether file is opened for appending
}
