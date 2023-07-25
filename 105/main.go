package main

import (
	"fmt"
	"sync"
)

type VoteBox struct {
	votes map[string]int
	lock  sync.RWMutex
}

func NewVoteBox() *VoteBox {
	return &VoteBox{
		votes: make(map[string]int),
	}
}

// 投票
func (vb *VoteBox) Vote(candidate string) {
	vb.lock.Lock()
	defer vb.lock.Unlock()

	vb.votes[candidate]++
}

// 获取候选人的得票数
func (vb *VoteBox) GetVotes(candidate string) int {
	vb.lock.RLock()
	defer vb.lock.RUnlock()

	return vb.votes[candidate]
}

func main() {
	vb := NewVoteBox()

	// 并发投票
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		vb.Vote("A")
	}()

	go func() {
		defer wg.Done()
		vb.Vote("B")
	}()

	go func() {
		defer wg.Done()
		vb.Vote("C")
	}()

	wg.Wait()

	// 输出每个候选人的得票数
	fmt.Println("Candidate A Votes:", vb.GetVotes("A"))
	fmt.Println("Candidate B Votes:", vb.GetVotes("B"))
	fmt.Println("Candidate C Votes:", vb.GetVotes("C"))
}
