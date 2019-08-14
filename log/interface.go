package log

type Adapter interface {
	//  Writing method writes the log bytes to the
	// destination, for example: file, terminal, network,
	// system log component.
	Writing(p []byte) error

	// Flush method flushes the contents of the buffer to the destination.
	Flush()

	// Close method is used to close all open resources.
	Close()
}
