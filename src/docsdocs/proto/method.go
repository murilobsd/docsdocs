package proto

// Commom Protocol methods.

// Method type
type Method uint64

const (
	// MethodGet ...
	MethodGet Method = 1 << iota
	// MethodPost ...
	MethodPost
	// MethodDelete ...
	MethodDelete
)
