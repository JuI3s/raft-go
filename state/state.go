package state

// Persistent state on all servers
type StatePersistent struct {
	// latest term server has seen (initialized to 0 on first boot, increases monotonically)
	currentTerm int

	// candidateId that received vote in current term (or null if none)
	votedFor int

	// log entries; each entry contains command for state machine, and term when entry was received by leader (first index is 1)
	log []int
}

// Volatile state on all servers
type StateVolatileServer struct {
	commitIndex int
	lastApplied int
}
