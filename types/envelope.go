package types

import (
	"encoding/xml"
)

// type LiteralBooleanValue string

// type LiteralBoolean struct {
// 	value LiteralBooleanValue
// }

// func (lb LiteralBoolean) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
// 	return xml.Attr{Name: name, Value: string(lb.value)}, nil
// }

// var (
// 	LiteralFalse = LiteralBoolean{value: "false"}
// 	LiteralTrue  = LiteralBoolean{value: "true"}
// )

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Link    string   `xml:"xmlns:link,attr"`
	Sec     string   `xml:"xmlns:sec,attr"`
	Typ     string   `xml:"xmlns:typ,attr"`
	Iat     string   `xml:"xmlns:iat,attr"`
	App     string   `xml:"xmlns:app,attr"`
	Ses     string   `xml:"xmlns:ses,attr"`
	Header  Header   `xml:"soapenv:Header"`
	Body    Body     `xml:"soapenv:Body"`
}

type Body struct {
	OTAHotelAvailRQ           *OTAHotelAvailRQ
	OTAHotelDescriptiveInfoRQ *OTAHotelDescriptiveInfoRQ
}

func NewEnvelope(header Header, body Body) Envelope {
	return Envelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Link:    "http://wsdl.amadeus.com/2010/06/ws/Link_v1",
		Sec:     "http://xml.amadeus.com/2010/06/Security_v1",
		Typ:     "http://xml.amadeus.com/2010/06/Types_v1",
		Iat:     "http://www.iata.org/IATA/2007/00/IATA2010.1",
		App:     "http://xml.amadeus.com/2010/06/AppMdw_CommonTypes_v3",
		Ses:     "http://xml.amadeus.com/2010/06/Session_v3",
		Header:  header,
		Body:    body,
	}
}
