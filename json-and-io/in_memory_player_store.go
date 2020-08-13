package main

// NewInMemoryPlayerStore stores wins by player and the count
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore stores wins by player and the count
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin will record win for the player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore will get the player score by name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// GetLeague returns the list of player in the league
func (i *InMemoryPlayerStore) GetLeague() []Player {
	league := []Player{}
	for name, wins := range i.store {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
