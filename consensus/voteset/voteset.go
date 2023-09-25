package voteset

import (
	"fmt"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/types/validator"
	"github.com/pactus-project/pactus/types/vote"
	"github.com/pactus-project/pactus/util/errors"
)

type voteSet struct {
	voteType   vote.Type
	round      int16
	validators map[crypto.Address]*validator.Validator
	totalPower int64
}

func newVoteSet(voteType vote.Type, round int16, totalPower int64,
	validators map[crypto.Address]*validator.Validator,
) *voteSet {
	return &voteSet{
		round:      round,
		voteType:   voteType,
		validators: validators,
		totalPower: totalPower,
	}
}

func (vs *voteSet) Type() vote.Type {
	return vs.voteType
}

// Round returns the round number for the VoteSet.
func (vs *voteSet) Round() int16 {
	return vs.round
}

// verifyVote checks if the given vote is valid.
// It returns the voting power of if valid, or an error if not.
func (vs *voteSet) verifyVote(v *vote.Vote) (int64, error) {
	signer := v.Signer()
	val := vs.validators[signer]
	if val == nil {
		return 0, errors.Errorf(errors.ErrInvalidAddress,
			"cannot find validator %s in committee", signer)
	}

	if err := v.Verify(val.PublicKey()); err != nil {
		return 0, errors.Errorf(errors.ErrInvalidSignature,
			"failed to verify vote")
	}

	return val.Power(), nil
}

// maximumFaultyPower calculates the maximum amount of power that can be faulty.
// Given n is the total power, $$ f_max = (n-1) / 3 $$
func (vs *voteSet) maximumFaultyPower(power int64) int64 {
	// Adding 2 before division to ensure rounding up
	return (vs.totalPower - 1 + 2) / 3
}

// hasThreeFPlusOnePower checks whether the given power is greater than 3f+1,
// where f is the maximum faulty power.
func (vs *voteSet) hasThreeFPlusOnePower(power int64) bool {
	return power > (3*vs.maximumFaultyPower(power) + 1)
}

// hasTwoFPlusOnePower checks whether the given power is greater than 2f+1,
// where f is the maximum faulty power.
func (vs *voteSet) hasTwoFPlusOnePower(power int64) bool {
	return power > (2*vs.maximumFaultyPower(power) + 1)
}

// hasFPlusOnePower checks whether the given power is greater than f+1,
// where f is the maximum faulty power.
func (vs *voteSet) hasFPlusOnePower(power int64) bool {
	return power > (vs.maximumFaultyPower(power) + 1)
}

func (vs *voteSet) String() string {
	return fmt.Sprintf("{%v/%s TOTAL:%v}", vs.round, vs.voteType, vs.totalPower)
}
