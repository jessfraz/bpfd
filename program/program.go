package program

import (
	"fmt"
)

var (
	// All registered programs.
	programs map[string]InitFunc
)

// InitFunc initializes the program.
type InitFunc func() (Program, error)

// Program defines the basic capabilities of a program.
type Program interface {
	// String returns a string representation of this program.
	String() string
	// Load creates the bpf module and starts collecting the data for the program.
	Load() error
	// Unload closes the bpf module and all the probes that all attached to it.
	Unload() error
	// WatchEvents starts the go routine to watch the events for the program.
	WatchEvents() error
}

// Init initialized the program map.
func init() {
	programs = make(map[string]InitFunc)
}

// Register registers an InitFunc for the program.
func Register(name string, initFunc InitFunc) error {
	if _, exists := programs[name]; exists {
		return fmt.Errorf("Name already registered %s", name)
	}
	programs[name] = initFunc

	return nil
}

// Get initializes and returns the registered program.
func Get(name string) (Program, error) {
	if initFunc, exists := programs[name]; exists {
		return initFunc()
	}

	return nil, fmt.Errorf("program %q does not exist as a supported program", name)
}

// List all the registered programs.
func List() []string {
	keys := []string{}
	for k := range programs {
		keys = append(keys, k)
	}
	return keys
}
