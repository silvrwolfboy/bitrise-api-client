// Code generated by go-swagger; DO NOT EDIT.

package build_artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/slapec93/bitrise-api-client/models"
)

// ArtifactListReader is a Reader for the ArtifactList structure.
type ArtifactListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtifactListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArtifactListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewArtifactListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewArtifactListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewArtifactListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewArtifactListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewArtifactListOK creates a ArtifactListOK with default headers values
func NewArtifactListOK() *ArtifactListOK {
	return &ArtifactListOK{}
}

/*ArtifactListOK handles this case with default header values.

OK
*/
type ArtifactListOK struct {
	Payload *models.V0ArtifactListResponseModel
}

func (o *ArtifactListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}][%d] artifactListOK  %+v", 200, o.Payload)
}

func (o *ArtifactListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ArtifactListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListBadRequest creates a ArtifactListBadRequest with default headers values
func NewArtifactListBadRequest() *ArtifactListBadRequest {
	return &ArtifactListBadRequest{}
}

/*ArtifactListBadRequest handles this case with default header values.

Bad Request
*/
type ArtifactListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ArtifactListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}][%d] artifactListBadRequest  %+v", 400, o.Payload)
}

func (o *ArtifactListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListUnauthorized creates a ArtifactListUnauthorized with default headers values
func NewArtifactListUnauthorized() *ArtifactListUnauthorized {
	return &ArtifactListUnauthorized{}
}

/*ArtifactListUnauthorized handles this case with default header values.

Unauthorized
*/
type ArtifactListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ArtifactListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}][%d] artifactListUnauthorized  %+v", 401, o.Payload)
}

func (o *ArtifactListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListNotFound creates a ArtifactListNotFound with default headers values
func NewArtifactListNotFound() *ArtifactListNotFound {
	return &ArtifactListNotFound{}
}

/*ArtifactListNotFound handles this case with default header values.

Not Found
*/
type ArtifactListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ArtifactListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}][%d] artifactListNotFound  %+v", 404, o.Payload)
}

func (o *ArtifactListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArtifactListInternalServerError creates a ArtifactListInternalServerError with default headers values
func NewArtifactListInternalServerError() *ArtifactListInternalServerError {
	return &ArtifactListInternalServerError{}
}

/*ArtifactListInternalServerError handles this case with default header values.

Internal Server Error
*/
type ArtifactListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ArtifactListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/builds/{build-slug}][%d] artifactListInternalServerError  %+v", 500, o.Payload)
}

func (o *ArtifactListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}