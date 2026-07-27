package electionday

import "fmt"

func NewVoteCounter(init int) *int{
	return &init
}

func VoteCount(cnt *int)int{
	if cnt == nil{
		return 0
	}
	return *cnt
}

func IncrementVoteCount(cnt *int, votes int){
	*cnt += votes
}

func NewElectionResult(name string, votes int) *ElectionResult {
	var res *ElectionResult
	res = &ElectionResult{Name: name, Votes: votes}
	return res
}

func DisplayResult(res *ElectionResult) string{
	return fmt.Sprintf("%s (%v)", res.Name, res.Votes)
}

func DecrementVotesOfCandidate(res map[string]int, name string) map[string]int{
	res[name]--
	return res
}
