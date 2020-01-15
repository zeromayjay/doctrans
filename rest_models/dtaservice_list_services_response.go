// Code generated by go-swagger; DO NOT EDIT.

package rest_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DtaserviceListServicesResponse dtaservice list services response
// swagger:model dtaserviceListServicesResponse
type DtaserviceListServicesResponse struct {

	// services
	Services []string `json:"services"`
}

// Validate validates this dtaservice list services response
func (m *DtaserviceListServicesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtaserviceListServicesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtaserviceListServicesResponse) UnmarshalBinary(b []byte) error {
	var res DtaserviceListServicesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
