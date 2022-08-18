package types

import "fmt"

type Header struct {
	Add                   string                 `xml:"xmlns:add,attr"`
	MessageID             string                 `xml:"add:MessageID"`
	Action                string                 `xml:"add:Action"`
	To                    string                 `xml:"add:To"`
	Security              *Security              `xml:"oas:Security"`
	AMASecurityHostedUser *AMASecurityHostedUser `xml:"AMA_SecurityHostedUser"`
	Session               *Session
}

type AMASecurityHostedUser struct {
	Xmlns  string `xml:"xmlns,attr"`
	UserID UserID `xml:"UserID"`
}

func NewAMASecurityHostedUser(config AmadeusConfig) *AMASecurityHostedUser {
	return &AMASecurityHostedUser{
		Xmlns:  "http://xml.amadeus.com/2010/06/Security_v1",
		UserID: NewUserID(config),
	}
}

type UserID struct {
	AgentDutyCode  string `xml:"AgentDutyCode,attr"`
	RequestorType  string `xml:"RequestorType,attr"`
	PseudoCityCode string `xml:"PseudoCityCode,attr"`
	POSType        string `xml:"POS_Type,attr"`
}

func NewUserID(config AmadeusConfig) UserID {
	return UserID{
		AgentDutyCode:  config.AgentDutyCode,
		RequestorType:  config.RequestorType,
		PseudoCityCode: config.PseudoCityCode,
		POSType:        config.POSType,
	}
}

func NewHeader(messageID, action string, session *Session, config AmadeusConfig) Header {
	var security *Security
	var hostedUser *AMASecurityHostedUser

	if session == nil || session.TransactionStatusCode == SessionStatusStart {
		security = NewSecurity(config)
		hostedUser = NewAMASecurityHostedUser(config)
	}

	return Header{
		Add:                   "http://www.w3.org/2005/08/addressing",
		MessageID:             messageID,
		Action:                action,
		To:                    fmt.Sprintf("%s%s", config.URL, config.WSAP),
		Security:              security,
		AMASecurityHostedUser: hostedUser,
		Session:               session,
	}
}
