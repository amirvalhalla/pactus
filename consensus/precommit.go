package consensus

import (
	"github.com/pactus-project/pactus/types/proposal"
	"github.com/pactus-project/pactus/types/vote"
)

type precommitState struct {
	*consensus
	hasVoted bool
}

func (s *precommitState) enter() {
	s.hasVoted = false

	s.decide()
}

func (s *precommitState) decide() {
	s.vote()

	precommits := s.log.PrecommitVoteSet(s.round)
	precommitQH := precommits.QuorumHash()
	if precommitQH != nil {
		s.logger.Debug("pre-commit has quorum", "hash", precommitQH.ShortString())

		roundProposal := s.log.RoundProposal(s.round)
		if roundProposal == nil {
			// There is a consensus about a proposal that we don't have yet.
			// Ask peers for this proposal.
			s.logger.Info("query for a decided proposal", "hash", precommitQH.ShortString())
			s.queryProposal()
		} else {
			// To ensure we have voted and won't be absent from the certificate
			if s.hasVoted {
				s.enterNewState(s.commitState)
			}
		}
	}
}

func (s *precommitState) vote() {
	if s.hasVoted {
		return
	}

	roundProposal := s.log.RoundProposal(s.round)
	if roundProposal == nil {
		s.queryProposal()
		s.logger.Debug("no proposal yet")
		return
	}

	if !roundProposal.IsForBlock(*s.preparedHash) {
		s.log.SetRoundProposal(s.round, nil)
		s.queryProposal()
		s.logger.Warn("double proposal detected",
			"roundProposal", roundProposal,
			"prepared", s.preparedHash.ShortString())
		return
	}

	// Everything is good
	s.signAddPrecommitVote(*s.preparedHash)
	s.hasVoted = true
}

func (s *precommitState) onAddVote(v *vote.Vote) {
	if v.Type() == vote.VoteTypePrecommit ||
		v.Type() == vote.VoteTypeCPPreVote {
		s.decide()
	}
}

func (s *precommitState) onSetProposal(_ *proposal.Proposal) {
	s.decide()
}

func (s *precommitState) onTimeout(_ *ticker) {
	// Ignore timeouts
}

func (s *precommitState) name() string {
	return "precommit"
}
