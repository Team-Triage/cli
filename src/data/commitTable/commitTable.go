package commitTable

import (
	"fmt"

	"github.com/team-triage/triage/types"
)

var hash map[int]types.CommitStore = make(map[int]types.CommitStore) // map where keys are offsets and 'ack/nack' is boolean
var CommitHash *types.SafeCommitHash = types.MakeSafeCommitHash(hash)

func Delete(offset int) {
	currentOffset := offset
	for {
		if _, ok := CommitHash.Read(currentOffset); ok {
			CommitHash.Delete(currentOffset)
			// delete(CommitHash, currentOffset)
			fmt.Printf("COMMIT TABLE: Deleting entry at offset %v\n", currentOffset)
			currentOffset--
		} else {
			break
		}
	}
}
