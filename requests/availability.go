package requests

import (
	"amadeus-xml/types"

	"github.com/google/uuid"
)

const (
	ActionAvailablity       = "http://webservices.amadeus.com/Hotel_MultiSingleAvailability_10.0"
	DefaultHotelCodeContext = "1A"
)

func NewAvailabilityRequest(infoSource types.InfoSource, start string, end string, currency string, country string, adults int, hotels []string, session *types.Session, config types.AmadeusConfig) (envelope types.Envelope, action string) {
	messageID := uuid.NewString()

	hotelRefs := make([]types.HotelRef, len(hotels))
	for i, hotel := range hotels {
		hotelRefs[i] = types.HotelRef{
			HotelCode:        hotel,
			ChainCode:        hotel[:2],
			HotelCityCode:    hotel[2:5],
			HotelCodeContext: "1A",
		}
	}

	action = ActionAvailablity
	envelope = types.NewEnvelope(
		types.NewHeader(messageID, ActionAvailablity, session, config),
		types.Body{
			OTAHotelAvailRQ: types.NewOTAHotelAvailRQ(
				"MultiSingle",
				currency,
				infoSource,
				start,
				end,
				country,
				adults,
				hotelRefs,
			),
		},
	)

	return
}
