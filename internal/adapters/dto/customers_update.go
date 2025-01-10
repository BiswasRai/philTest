package dto

import (
    "errors"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type UpdateCustomerRequest struct {
    Name     string `json:"name"`
    Location string `json:"location"`
}

func NewUpdateCustomerDTO() *UpdateCustomerRequest {
    return &UpdateCustomerRequest{
        Name:     "",
        Location: "",
    }
}

func ValidatePayloadUpdateCustomerRequest(r *http.Request) (*UpdateCustomerRequest, error) {
    var updateCustomerRequest UpdateCustomerRequest
     body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return nil, errors.New("failed to read request body")
    }

     log.WithFields(map[string]interface{}{"body": string(body)}).Info("Request body")

    if err := json.Unmarshal(body, &updateCustomerRequest); err != nil {
        log.WithError(err).Error("failed to unmarshal request body")
        return nil, errors.New("failed to verify request body")
    }

    if updateCustomerRequest.Name == "" {
        return nil, errors.New("name is required")
    }

    if updateCustomerRequest.Location == "" {
        return nil, errors.New("location is required")
    }



    return &updateCustomerRequest, nil
}