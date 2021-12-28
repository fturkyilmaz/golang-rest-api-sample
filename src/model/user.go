package model

import "fmt"

func init() {
	fmt.Println("User package initialized")
}

type User struct {
	ID                    int    `json:"id"`
	UID                   string `json:"uid"`
	Password              string `json:"password"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Username              string `json:"username"`
	Email                 string `json:"email"`
	Avatar                string `json:"avatar"`
	Gender                string `json:"gender"`
	PhoneNumber           string `json:"phone_number"`
	SocialInsuranceNumber string `json:"social_insurance_number"`
	DateOfBirth           string `json:"date_of_birth"`
	Employment            struct {
		Title    string `json:"title"`
		KeySkill string `json:"key_skill"`
	} `json:"employment"`
	Address struct {
		City          string `json:"city"`
		StreetName    string `json:"street_name"`
		StreetAddress string `json:"street_address"`
		ZipCode       string `json:"zip_code"`
		State         string `json:"state"`
		Country       string `json:"country"`
		Coordinates   struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"coordinates"`
	} `json:"address"`
	CreditCard struct {
		CcNumber string `json:"cc_number"`
	} `json:"credit_card"`
	Subscription struct {
		Plan          string `json:"plan"`
		Status        string `json:"status"`
		PaymentMethod string `json:"payment_method"`
		Term          string `json:"term"`
	} `json:"subscription"`
}
