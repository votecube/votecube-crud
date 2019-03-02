package location

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */
func DeserializePollTowns(data []byte, cursor *int64, dataLen int64, err error) (models.PollsTownSlice, error) {
	var pollsTowns models.PollsTownSlice
	var numPollsTowns int64

	if err != nil {
		return pollsTowns, err
	}

	numPollsTowns, err = deserialize.RNum(data, cursor, dataLen, err)

	if numPollsTowns == 0 {
		return pollsTowns, err
	}

	pollsTowns = make(models.PollsTownSlice, numPollsTowns)

	for i := int64(0); i < numPollsTowns; i++ {
		var continentId int64
		continentId, err = deserialize.RNum(data, cursor, dataLen, err)

		if err != nil {
			return pollsTowns, err
		}

		pollsTowns[i] = &models.PollsTown{
			TownID: continentId,
		}
	}

	return pollsTowns, err
}
