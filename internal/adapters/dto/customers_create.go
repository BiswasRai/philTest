package dto

import (
	"errors"
	"net/http"
    "io/ioutil"
	// "strconv"
    "encoding/json"
)

type CreateCustomerRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Location string `json:"location"`
	LifetimeValue float64 `json:"lifetimeValue"`
}

func NewCreateCustomerDTO() *CreateCustomerRequest {
    return &CreateCustomerRequest{
        Name: "",
		Email: "",
		Location: "",
		LifetimeValue: 0.0,
    }
}

func ValidatePayloadCreateCustomerRequest(r *http.Request) (*CreateCustomerRequest, error) {
    var createCustomerRequest CreateCustomerRequest

    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return nil, errors.New("failed to read request body")
    }

    log.WithFields(map[string]interface{}{"body": string(body)}).Info("Request body")

    if err := json.Unmarshal(body, &createCustomerRequest); err != nil {
        log.WithError(err).Error("failed to unmarshal request body")
        return nil, errors.New("failed to verify request body")
    }

    if createCustomerRequest.Name == "" {
        return nil, errors.New("name is required")
    }

    if createCustomerRequest.Email == "" {
        return nil, errors.New("email is required")
    }

    if createCustomerRequest.Location == "" {
        return nil, errors.New("location is required")
    }

    if createCustomerRequest.LifetimeValue == 0 {
        return nil, errors.New("lifetimeValue is required")
    }

    if err := json.Unmarshal(body, &createCustomerRequest); err != nil {
        log.WithError(err).Error("failed to unmarshal request body")
        return nil, errors.New("failed to request body")
    }

    return &createCustomerRequest, nil
}