package types

type InfoSource string
type EchoToken string
type SearchCacheLevel string

const (
	InfoSourceLeisure      InfoSource = "Leisure"
	InfoSourceDistribution InfoSource = "Distibution"
)

const (
	EchoTokenMultiSingle EchoToken = "MultiSingle"
)

const (
	SearchCacheLevelLive SearchCacheLevel = "Live"
)

type OTAHotelAvailRQ struct {
	SummaryOnly          bool                 `xml:"SummaryOnly,attr"`
	RateRangeOnly        bool                 `xml:"RateRangeOnly,attr"`
	EchoToken            EchoToken            `xml:"EchoToken,attr"`
	RequestedCurrency    string               `xml:"RequestedCurrency,attr"`
	Version              string               `xml:"Version,attr"`
	PrimaryLangID        string               `xml:"PrimaryLangID,attr"`
	SearchCacheLevel     SearchCacheLevel     `xml:"SearchCacheLevel,attr"`
	MaxResponses         int                  `xml:"MaxResponses,attr"`
	AvailRatesOnly       bool                 `xml:"AvailRatesOnly,attr"`
	ExactMatchOnly       bool                 `xml:"ExactMatchOnly,attr"`
	RateDetailsInd       bool                 `xml:"RateDetailsInd,attr"`
	AvailRequestSegments AvailRequestSegments `xml:"AvailRequestSegments"`
}

type AvailRequestSegments struct {
	AvailRequestSegment AvailRequestSegment `xml:"AvailRequestSegment"`
}

type AvailRequestSegment struct {
	InfoSource          InfoSource          `xml:"InfoSource,attr"`
	HotelSearchCriteria HotelSearchCriteria `xml:"HotelSearchCriteria"`
}

type HotelSearchCriteria struct {
	AvailableOnlyIndicator bool        `xml:"AvailableOnlyIndicator,attr"`
	Criterion              []Criterion `xml:"Criterion"`
}

type Criterion struct {
	ExactMatch         bool                `xml:"ExactMatch,attr"`
	StayDateRange      *StayDateRange      `xml:"StayDateRange"`
	RatePlanCandidates *RatePlanCandidates `xml:"RatePlanCandidates"`
	Profiles           *Profiles           `xml:"Profiles"`
	RoomStayCandidates *RoomStayCandidates `xml:"RoomStayCandidates"`
	HotelRef           []HotelRef          `xml:"HotelRef"`
}

type StayDateRange struct {
	Start string `xml:"Start,attr"`
	End   string `xml:"End,attr"`
}

type RatePlanCandidates struct {
	RatePlanCandidate []RatePlanCandidate `xml:"RatePlanCandidate"`
}

type RatePlanCandidate struct {
	RatePlanCode string `xml:"RatePlanCode,attr"`
}

type Profiles struct {
	ProfileInfo ProfileInfo `xml:"ProfileInfo"`
}

type RoomStayCandidates struct {
	RoomStayCandidate RoomStayCandidate `xml:"RoomStayCandidate"`
}

type ProfileInfo struct {
	Profile Profile `xml:"Profile"`
}

type Profile struct {
	ProfileType int      `xml:"ProfileType,attr"`
	Customer    Customer `xml:"Customer"`
}

type Customer struct {
	Address            Address            `xml:"Address"`
	CitizenCountryName CitizenCountryName `xml:"CitizenCountryName"`
}

type Address struct {
	UseType     int         `xml:"UseType,attr"`
	CountryName CountryName `xml:"CountryName"`
}

type CountryName struct {
	Code string `xml:"Code,attr"`
}

type CitizenCountryName struct {
	Code string `xml:"Code,attr"`
}

type HotelRef struct {
	HotelCode        string `xml:"HotelCode,attr"`
	ChainCode        string `xml:"ChainCode,attr"`
	HotelCityCode    string `xml:"HotelCityCode,attr"`
	HotelCodeContext string `xml:"HotelCodeContext,attr"`
}

type RoomStayCandidate struct {
	RoomID      int         `xml:"RoomID,attr"`
	Quantity    int         `xml:"Quantity,attr"`
	GuestCounts GuestCounts `xml:"GuestCounts"`
}

type GuestCounts struct {
	GuestCount []GuestCount `xml:"GuestCount"`
}

func NewGuestCounts(adults int) GuestCounts {
	return GuestCounts{
		GuestCount: []GuestCount{
			{
				AgeQualifyingCode: 10,
				Count:             adults,
			},
		},
	}
}

type GuestCount struct {
	AgeQualifyingCode int `xml:"AgeQualifyingCode,attr"`
	Count             int `xml:"Count,attr"`
}

var (
	DefaultRatePlans = []RatePlanCandidate{
		{RatePlanCode: "ABC"},
		{RatePlanCode: "AOM"},
		{RatePlanCode: "PRO"},
		{RatePlanCode: "COR"},
		{RatePlanCode: "X68"},
		{RatePlanCode: "H81"},
		{RatePlanCode: "H4Y"},
		{RatePlanCode: "YR1"},
		{RatePlanCode: "YR9"},
		{RatePlanCode: "Y45"},
	}
)

const (
	DefaultCustomerCountry = "US"
)

func NewOTAHotelAvailRQ(echoToken EchoToken, currency string, infoSource InfoSource, start, end, country string, adults int, hotelRefs []HotelRef) OTAHotelAvailRQ {
	if country == "" {
		country = DefaultCustomerCountry
	}
	return OTAHotelAvailRQ{
		SummaryOnly:       true,
		RateRangeOnly:     true,
		EchoToken:         echoToken,
		RequestedCurrency: currency,
		Version:           "4.000",
		PrimaryLangID:     "en",
		SearchCacheLevel:  SearchCacheLevelLive,
		MaxResponses:      64,
		AvailRatesOnly:    true,
		ExactMatchOnly:    false,
		RateDetailsInd:    true,
		AvailRequestSegments: AvailRequestSegments{
			AvailRequestSegment: AvailRequestSegment{
				InfoSource: infoSource,
				HotelSearchCriteria: HotelSearchCriteria{
					AvailableOnlyIndicator: true,
					Criterion: []Criterion{
						{
							ExactMatch: true,
							StayDateRange: &StayDateRange{
								Start: start,
								End:   end,
							},
							RatePlanCandidates: &RatePlanCandidates{
								RatePlanCandidate: DefaultRatePlans,
							},
							Profiles: &Profiles{
								ProfileInfo: ProfileInfo{
									Profile: Profile{
										ProfileType: 1,
										Customer: Customer{
											Address: Address{
												UseType: 7,
												CountryName: CountryName{
													Code: country,
												},
											},
											CitizenCountryName: CitizenCountryName{
												Code: country,
											},
										},
									},
								},
							},
							RoomStayCandidates: &RoomStayCandidates{
								RoomStayCandidate: RoomStayCandidate{
									RoomID:      1,
									Quantity:    1,
									GuestCounts: NewGuestCounts(adults),
								},
							},
						},
						{
							ExactMatch: true,
							HotelRef:   hotelRefs,
						},
					},
				},
			},
		},
	}
}
