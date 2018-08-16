package main

type Event struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Sub  string `json:"sub"`
}

type Person struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Pictures []string `json:"pictures"`
}

type Place struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	People []*Person `json:"people"`
}

type Claims struct {
	Aud             string `json:"aud"`
	AuthTime        string `json:"auth_time"`
	CognitoUsername string `json:"cognito:username"`
	Email           string `json:"email"`
	EmailVerified   string `json:"email_verified"`
	EventID         string `json:"event_id"`
	Exp             string `json:"exp"`
	FamilyName      string `json:"family_name"`
	Gender          string `json:"gender"`
	GivenName       string `json:"given_name"`
	Iat             string `json:"iat"`
	Iss             string `json:"iss"`
	Picture         string `json:"picture"`
	Sub             string `json:"sub"`
	TokenUse        string `json:"token_use"`
}
