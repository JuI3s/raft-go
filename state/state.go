package state

// Persistent state on all servers
type StatePersistent struct {
	// latest term server has seen (initialized to 0 on first boot, increases monotonically)
	currentTerm uint

	// candidateId that received vote in current term (or null if none)
	votedFor uint

	// log entries; each entry contains command for state machine, and term when entry was received by leader (first index is 1)
	log []string
}

// Volatile state on all servers
type StateVolatileServer struct {
	// index of highest log entry known to be committed (initialized to 0, increases monotonically)
	commitIndex uint

	// index of highest log entry applied to state machine (initialized to 0, increases monotonically)
	lastApplied uint
}

type StateVolatileLeader struct {
	// for each server, index of the next log entry to send to that server (initialized to leader last log index + 1)
	nextIndex []uint

	// for each server, index of highest log entry known to be replicated on server (initialized to 0, increases monotonically)
	matchIndex []uint
}
