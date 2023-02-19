package raftik

// NodeState represents the state of a raft node
type NodeState int

const (
	// raft node is in follower state
	StateFollower NodeState = iota
	// raft node is in candidate state
	StateCandidate
	// raft node is in leader state
	StateLeader
)

type Config struct {
	// Raft node ID
	ID string

	// Raft node address
	Address string

	// Raft node state
	State NodeState
}

func NewConfig() *Config {
	return &Config{}
}
