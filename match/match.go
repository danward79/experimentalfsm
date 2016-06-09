package match

import "github.com/danward79/fsm/machine"

//match looks at an input stream and matches it to a state
//state is stored as a map.
//map[event]state

type matchThing struct {
	InputStream chan []string
	Events      map[string]machine.State
}

//Matcher interface
type Matcher interface {
	Match() machine.State
}

//Match returns a state, which matches an event
func Match() machine.State {
	return ""
}
