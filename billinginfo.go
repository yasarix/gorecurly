package gorecurly

import (
	"encoding/xml"
)


//Billing Info struct
type BillingInfo struct {
	XMLName xml.Name `xml:"billing_info"`
	endpoint string
	r *Recurly
	Account *AccountStub `xml:"account,omitempty"`
	AccountCode string `xml:"account_code,omitempty"`
	FirstName string `xml:"first_name,omitempty"`
	LastName string `xml:"last_name,omitempty"`
	Address1 string `xml:"address1,omitempty"`
	Address2 string `xml:"address2,omitempty"`
	City string `xml:"city,omitempty"`
	State string `xml:"state,omitempty"`
	Zip string `xml:"zip,omitempty"`
	Country string `xml:"country,omitempty"`
	Phone string `xml:"phone,omitempty"`
	VatNumber string `xml:"vat_number,omitempty"`
	IPAddress string `xml:"ip_address,omitempty"`
	IPAddressCountry string `xml:"ip_address_country,omitempty"`
	Number string `xml:"number,omitempty"`
	FirstSix string `xml:"first_six,omitempty"`
	LastFour string `xml:"last_four,omitempty"`
	VerificationValue string `xml:"verification_value,omitempty"`
	CardType string `xml:"card_type,omitempty"`
	Month int `xml:"month,omitempty"`
	Year int `xml:"year,omitempty"`
	BillingAgreementID string `xml:"billing_agreement_id,omitempty"`
}

//Update an billing info 
func (b *BillingInfo) Update() (error) {
	newbilling := new(BillingInfo)
	*newbilling = *b
	newbilling.AccountCode = ""
	newbilling.Account = nil
	newbilling.FirstSix = ""
	newbilling.LastFour = ""
	newbilling.CardType = ""
	return b.r.doUpdate(newbilling,ACCOUNTS + "/" + b.Account.GetCode() + "/" + BILLINGINFO)
}

//Delete billing info for an account
func (b *BillingInfo) Delete() (error) {
	return b.r.doDelete(ACCOUNTS + "/" + b.AccountCode + "/" + BILLINGINFO)
}

func (b BillingInfo) GetAccount() (Account,error) {
	return b.r.GetAccount(b.Account.GetCode())
}
