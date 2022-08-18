package types

import (
	"encoding/xml"

	"github.com/goccy/go-json"
)

type SessionStatus string

const (
	SessionStatusStart    = "Start"
	SessionStatusInSeries = "InSeries"
	SessionStatusEnd      = "End"
)

type Session struct {
	XMLName xml.Name `xml:"awsse:Session"`

	Awsse                 string `xml:"xmlns:awsse,attr"`
	TransactionStatusCode string `xml:"TransactionStatusCode,attr" json:"status"`
	SessionId             string `xml:"awsse:SessionId,omitempty" json:"id"`
	SequenceNumber        string `xml:"awsse:SequenceNumber,omitempty" json:"number"`
	SecurityToken         string `xml:"awsse:SecurityToken,omitempty" json:"token"`
}

func NewSession() *Session {
	return &Session{
		Awsse:                 "http://xml.amadeus.com/2010/06/Session_v3",
		TransactionStatusCode: SessionStatusStart,
	}
}

func (s Session) Encode() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func DecodeSession(encoded string) *Session {
	session := NewSession()
	json.Unmarshal([]byte(encoded), &session)
	return session
}
