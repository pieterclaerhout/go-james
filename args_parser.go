package james

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	// ErrArgsMissing is what is returned when no arguments could be found
	ErrArgsMissing = errors.New("No arguments were found")

	// ErrArgsDecodeError is what is returned when the arguments could be decoded (empty or invalid JSON)
	ErrArgsDecodeError = errors.New("Failed to decoce arguments")
)

// parseArgsInto checks if the args slice contain arguments as their second parameter (as with os.Args, the first
// argument is the path to the executable).
//
// If so, it will parse the parameter from a JSON string into the variable destination
func parseArgsInto(args []string, destination interface{}) error {

	if len(args) < 2 {
		return ErrArgsMissing
	}

	if err := json.Unmarshal([]byte(args[1]), &destination); err != nil {
		return errors.Wrap(err, ErrArgsDecodeError.Error())
	}

	return nil

}
