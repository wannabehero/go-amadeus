package requests

import (
	"amadeus/types"

	"github.com/google/uuid"
)

const (
	ActionDescriptiveInfo = "http://webservices.amadeus.com/OTA_HotelDescriptiveInfoRQ_07.1_1A2007A"
)

func NewDescriptiveInfoRequest(hotels []string, session *types.Session, config types.AmadeusConfig) (envelope types.Envelope, action string) {
	descriptiveInfos := make([]types.HotelDescriptiveInfo, len(hotels))
	for i, hotel := range hotels {
		descriptiveInfos[i] = types.NewHotelDescriptiveInfo(hotel)
	}

	descriptiveInfoRQ := types.NewOTAHotelDescriptiveInfoRQ(descriptiveInfos)

	action = ActionDescriptiveInfo
	envelope = types.NewEnvelope(
		types.NewHeader(uuid.NewString(), ActionDescriptiveInfo, session, config),
		types.Body{
			OTAHotelDescriptiveInfoRQ: &descriptiveInfoRQ,
		},
	)

	return
}
