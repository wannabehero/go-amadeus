package types

import "amadeus/utils"

type Security struct {
	Oas           string        `xml:"xmlns:oas,attr"`
	Oas1          string        `xml:"xmlns:oas1,attr"`
	UsernameToken UsernameToken `xml:"oas:UsernameToken"`
}

type UsernameToken struct {
	ID       string   `xml:"oas1:Id,attr"`
	Username string   `xml:"oas:Username"`
	Nonce    Nonce    `xml:"oas:Nonce"`
	Password Password `xml:"oas:Password"`
	Created  string   `xml:"oas1:Created"`
}

type Nonce struct {
	Text         string `xml:",chardata"`
	EncodingType string `xml:"EncodingType,attr"`
}

type Password struct {
	Text string `xml:",chardata"`
	Type string `xml:"Type,attr"`
}

func NewSecurity(config AmadeusConfig) *Security {
	timestamp, nonce, signature := utils.CreateSignature(config.Password)

	return &Security{
		Oas:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		Oas1: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		UsernameToken: UsernameToken{
			ID:       "UsernameToken-1",
			Username: config.Username,
			Nonce: Nonce{
				EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
				Text:         nonce,
			},
			Password: Password{
				Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest",
				Text: signature,
			},
			Created: timestamp,
		},
	}
}
