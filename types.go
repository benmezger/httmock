package main

type HTTPSpecMethodResponse struct {
	Status  int
	Payload string
}

type HTTPSpecMethodRequest struct {
	Args string
	Body string
}

type HTTPSpecMethod struct {
	Request  HTTPSpecMethodRequest
	Response HTTPSpecMethodResponse
}

type HTTPSpec struct {
	Paths map[string]map[string]HTTPSpecMethod
}
