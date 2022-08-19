package types

import "encoding/xml"

type HotelInfo struct {
	SendData    bool `xml:"SendData,attr"`
	HotelStatus bool `xml:"HotelStatus,attr"`
}

type FacilityInfo struct {
	SendMeetingRooms bool `xml:"SendMeetingRooms,attr"`
	SendGuestRooms   bool `xml:"SendGuestRooms,attr"`
	SendRestaurants  bool `xml:"SendRestaurants,attr"`
}

type Policies struct {
	SendPolicies bool `xml:"SendPolicies,attr"`
}

type AreaInfo struct {
	SendRefPoints   bool `xml:"SendRefPoints,attr"`
	SendAttractions bool `xml:"SendAttractions,attr"`
	SendRecreations bool `xml:"SendRecreations,attr"`
}

type AffiliationInfo struct {
	SendLoyalPrograms bool `xml:"SendLoyalPrograms,attr"`
	SendAwards        bool `xml:"SendAwards,attr"`
}

type ContactInfo struct {
	SendData bool `xml:"SendData,attr"`
}

type MultimediaObjects struct {
	SendData bool `xml:"SendData,attr"`
}

type ContentInfo struct {
	XMLName xml.Name `xml:"ContentInfo"`

	Name string `xml:"Name,attr"`
}

type ContentInfos struct {
	ContentInfo []ContentInfo
}

type HotelDescriptiveInfo struct {
	HotelCode         string            `xml:"HotelCode,attr"`
	HotelInfo         HotelInfo         `xml:"HotelInfo"`
	FacilityInfo      FacilityInfo      `xml:"FacilityInfo"`
	Policies          Policies          `xml:"Policies"`
	AreaInfo          AreaInfo          `xml:"AreaInfo"`
	AffiliationInfo   AffiliationInfo   `xml:"AffiliationInfo"`
	ContactInfo       ContactInfo       `xml:"ContactInfo"`
	MultimediaObjects MultimediaObjects `xml:"MultimediaObjects"`
	ContentInfos      ContentInfos      `xml:"ContentInfos"`
}

type HotelDescriptiveInfos struct {
	HotelDescriptiveInfo []HotelDescriptiveInfo `xml:"HotelDescriptiveInfo"`
}

type OTAHotelDescriptiveInfoRQ struct {
	XMLName xml.Name `xml:"OTA_HotelDescriptiveInfoRQ"`

	EchoToken             EchoToken             `xml:"EchoToken,attr"`
	Version               string                `xml:"Version,attr"`
	PrimaryLangID         string                `xml:"PrimaryLangID,attr"`
	HotelDescriptiveInfos HotelDescriptiveInfos `xml:"HotelDescriptiveInfos"`
}

func NewHotelDescriptiveInfo(hotelCode string) HotelDescriptiveInfo {
	return HotelDescriptiveInfo{
		HotelCode:         hotelCode,
		HotelInfo:         HotelInfo{SendData: true, HotelStatus: true},
		FacilityInfo:      FacilityInfo{SendMeetingRooms: true, SendGuestRooms: true, SendRestaurants: true},
		Policies:          Policies{SendPolicies: true},
		AreaInfo:          AreaInfo{SendRefPoints: true, SendAttractions: true, SendRecreations: true},
		AffiliationInfo:   AffiliationInfo{SendLoyalPrograms: true, SendAwards: true},
		ContactInfo:       ContactInfo{SendData: true},
		MultimediaObjects: MultimediaObjects{SendData: true},
		ContentInfos: ContentInfos{
			ContentInfo: []ContentInfo{
				{
					Name: "SecureMultimediaURLs",
				},
			},
		},
	}
}

func NewOTAHotelDescriptiveInfoRQ(descriptiveInfos []HotelDescriptiveInfo) OTAHotelDescriptiveInfoRQ {
	return OTAHotelDescriptiveInfoRQ{
		EchoToken:     EchoTokenWithParsing,
		Version:       "6.001",
		PrimaryLangID: "en",
		HotelDescriptiveInfos: HotelDescriptiveInfos{
			HotelDescriptiveInfo: descriptiveInfos,
		},
	}
}
