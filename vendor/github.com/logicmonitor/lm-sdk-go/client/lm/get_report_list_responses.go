// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// GetReportListReader is a Reader for the GetReportList structure.
type GetReportListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetReportListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportListOK creates a GetReportListOK with default headers values
func NewGetReportListOK() *GetReportListOK {
	return &GetReportListOK{}
}

/*GetReportListOK handles this case with default header values.

successful operation
*/
type GetReportListOK struct {
	Payload *models.ReportPaginationResponse
}

func (o *GetReportListOK) Error() string {
	return fmt.Sprintf("[GET /report/reports][%d] getReportListOK  %+v", 200, o.Payload)
}

func (o *GetReportListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportListDefault creates a GetReportListDefault with default headers values
func NewGetReportListDefault(code int) *GetReportListDefault {
	return &GetReportListDefault{
		_statusCode: code,
	}
}

/*GetReportListDefault handles this case with default header values.

Error
*/
type GetReportListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get report list default response
func (o *GetReportListDefault) Code() int {
	return o._statusCode
}

func (o *GetReportListDefault) Error() string {
	return fmt.Sprintf("[GET /report/reports][%d] getReportList default  %+v", o._statusCode, o.Payload)
}

func (o *GetReportListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
