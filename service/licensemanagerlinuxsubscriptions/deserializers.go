// Code generated by smithy-go-codegen DO NOT EDIT.


package licensemanagerlinuxsubscriptions

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"encoding/json"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	smithy "github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithyio "github.com/aws/smithy-go/io"
	"strings"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions/types"
)

type awsRestjson1_deserializeOpGetServiceSettings struct {
}

func (*awsRestjson1_deserializeOpGetServiceSettings) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetServiceSettings) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil { return out, metadata, err }
	
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}
	
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetServiceSettings(response, &metadata)
	}
	output := &GetServiceSettingsOutput{}
	out.Result = output
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(response.Body, ringBuffer)
	
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}
	
	err = awsRestjson1_deserializeOpDocumentGetServiceSettingsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}
	
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetServiceSettings(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())
	
	errorCode := "UnknownError"
	errorMessage := errorCode
	
	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 { errorCode = restjson.SanitizeErrorCode(headerCode) }
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 { errorCode = restjson.SanitizeErrorCode(jsonCode) }
	if len(message) != 0 { errorMessage = message }
	
	switch {
		case strings.EqualFold("InternalServerException", errorCode):
			return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)
		
		case strings.EqualFold("ThrottlingException", errorCode):
			return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)
		
		case strings.EqualFold("ValidationException", errorCode):
			return awsRestjson1_deserializeErrorValidationException(response, errorBody)
		
		default:
			genericError := &smithy.GenericAPIError{
				Code: errorCode,
				Message: errorMessage,
			}
			return genericError
		
	}
}

func awsRestjson1_deserializeOpDocumentGetServiceSettingsOutput(v **GetServiceSettingsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *GetServiceSettingsOutput
	if *v == nil {
		sv = &GetServiceSettingsOutput{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "HomeRegions":
				if err := awsRestjson1_deserializeDocumentStringList(&sv.HomeRegions, value); err != nil {
					return err
				}
			
			case "LinuxSubscriptionsDiscovery":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected LinuxSubscriptionsDiscovery to be of type string, got %T instead", value)
					}
					sv.LinuxSubscriptionsDiscovery = types.LinuxSubscriptionsDiscovery(jtv)
				}
			
			case "LinuxSubscriptionsDiscoverySettings":
				if err := awsRestjson1_deserializeDocumentLinuxSubscriptionsDiscoverySettings(&sv.LinuxSubscriptionsDiscoverySettings, value); err != nil {
					return err
				}
			
			case "Status":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected Status to be of type string, got %T instead", value)
					}
					sv.Status = types.Status(jtv)
				}
			
			case "StatusMessage":
				if err := awsRestjson1_deserializeDocumentStringMap(&sv.StatusMessage, value); err != nil {
					return err
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpListLinuxSubscriptionInstances struct {
}

func (*awsRestjson1_deserializeOpListLinuxSubscriptionInstances) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListLinuxSubscriptionInstances) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil { return out, metadata, err }
	
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}
	
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListLinuxSubscriptionInstances(response, &metadata)
	}
	output := &ListLinuxSubscriptionInstancesOutput{}
	out.Result = output
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(response.Body, ringBuffer)
	
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}
	
	err = awsRestjson1_deserializeOpDocumentListLinuxSubscriptionInstancesOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}
	
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListLinuxSubscriptionInstances(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())
	
	errorCode := "UnknownError"
	errorMessage := errorCode
	
	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 { errorCode = restjson.SanitizeErrorCode(headerCode) }
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 { errorCode = restjson.SanitizeErrorCode(jsonCode) }
	if len(message) != 0 { errorMessage = message }
	
	switch {
		case strings.EqualFold("InternalServerException", errorCode):
			return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)
		
		case strings.EqualFold("ThrottlingException", errorCode):
			return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)
		
		case strings.EqualFold("ValidationException", errorCode):
			return awsRestjson1_deserializeErrorValidationException(response, errorBody)
		
		default:
			genericError := &smithy.GenericAPIError{
				Code: errorCode,
				Message: errorMessage,
			}
			return genericError
		
	}
}

func awsRestjson1_deserializeOpDocumentListLinuxSubscriptionInstancesOutput(v **ListLinuxSubscriptionInstancesOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *ListLinuxSubscriptionInstancesOutput
	if *v == nil {
		sv = &ListLinuxSubscriptionInstancesOutput{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "Instances":
				if err := awsRestjson1_deserializeDocumentInstanceList(&sv.Instances, value); err != nil {
					return err
				}
			
			case "NextToken":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.NextToken = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpListLinuxSubscriptions struct {
}

func (*awsRestjson1_deserializeOpListLinuxSubscriptions) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListLinuxSubscriptions) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil { return out, metadata, err }
	
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}
	
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListLinuxSubscriptions(response, &metadata)
	}
	output := &ListLinuxSubscriptionsOutput{}
	out.Result = output
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(response.Body, ringBuffer)
	
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}
	
	err = awsRestjson1_deserializeOpDocumentListLinuxSubscriptionsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}
	
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListLinuxSubscriptions(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())
	
	errorCode := "UnknownError"
	errorMessage := errorCode
	
	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 { errorCode = restjson.SanitizeErrorCode(headerCode) }
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 { errorCode = restjson.SanitizeErrorCode(jsonCode) }
	if len(message) != 0 { errorMessage = message }
	
	switch {
		case strings.EqualFold("InternalServerException", errorCode):
			return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)
		
		case strings.EqualFold("ThrottlingException", errorCode):
			return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)
		
		case strings.EqualFold("ValidationException", errorCode):
			return awsRestjson1_deserializeErrorValidationException(response, errorBody)
		
		default:
			genericError := &smithy.GenericAPIError{
				Code: errorCode,
				Message: errorMessage,
			}
			return genericError
		
	}
}

func awsRestjson1_deserializeOpDocumentListLinuxSubscriptionsOutput(v **ListLinuxSubscriptionsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *ListLinuxSubscriptionsOutput
	if *v == nil {
		sv = &ListLinuxSubscriptionsOutput{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "NextToken":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.NextToken = ptr.String(jtv)
				}
			
			case "Subscriptions":
				if err := awsRestjson1_deserializeDocumentSubscriptionList(&sv.Subscriptions, value); err != nil {
					return err
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

type awsRestjson1_deserializeOpUpdateServiceSettings struct {
}

func (*awsRestjson1_deserializeOpUpdateServiceSettings) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpUpdateServiceSettings) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil { return out, metadata, err }
	
	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}
	
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorUpdateServiceSettings(response, &metadata)
	}
	output := &UpdateServiceSettingsOutput{}
	out.Result = output
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(response.Body, ringBuffer)
	
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}
	
	err = awsRestjson1_deserializeOpDocumentUpdateServiceSettingsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}
	
	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorUpdateServiceSettings(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())
	
	errorCode := "UnknownError"
	errorMessage := errorCode
	
	headerCode := response.Header.Get("X-Amzn-ErrorType")
	if len(headerCode) != 0 { errorCode = restjson.SanitizeErrorCode(headerCode) }
	
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	jsonCode, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	if len(headerCode) == 0 && len(jsonCode) != 0 { errorCode = restjson.SanitizeErrorCode(jsonCode) }
	if len(message) != 0 { errorMessage = message }
	
	switch {
		case strings.EqualFold("InternalServerException", errorCode):
			return awsRestjson1_deserializeErrorInternalServerException(response, errorBody)
		
		case strings.EqualFold("ThrottlingException", errorCode):
			return awsRestjson1_deserializeErrorThrottlingException(response, errorBody)
		
		case strings.EqualFold("ValidationException", errorCode):
			return awsRestjson1_deserializeErrorValidationException(response, errorBody)
		
		default:
			genericError := &smithy.GenericAPIError{
				Code: errorCode,
				Message: errorMessage,
			}
			return genericError
		
	}
}

func awsRestjson1_deserializeOpDocumentUpdateServiceSettingsOutput(v **UpdateServiceSettingsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *UpdateServiceSettingsOutput
	if *v == nil {
		sv = &UpdateServiceSettingsOutput{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "HomeRegions":
				if err := awsRestjson1_deserializeDocumentStringList(&sv.HomeRegions, value); err != nil {
					return err
				}
			
			case "LinuxSubscriptionsDiscovery":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected LinuxSubscriptionsDiscovery to be of type string, got %T instead", value)
					}
					sv.LinuxSubscriptionsDiscovery = types.LinuxSubscriptionsDiscovery(jtv)
				}
			
			case "LinuxSubscriptionsDiscoverySettings":
				if err := awsRestjson1_deserializeDocumentLinuxSubscriptionsDiscoverySettings(&sv.LinuxSubscriptionsDiscoverySettings, value); err != nil {
					return err
				}
			
			case "Status":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected Status to be of type string, got %T instead", value)
					}
					sv.Status = types.Status(jtv)
				}
			
			case "StatusMessage":
				if err := awsRestjson1_deserializeDocumentStringMap(&sv.StatusMessage, value); err != nil {
					return err
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeErrorInternalServerException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InternalServerException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	err := awsRestjson1_deserializeDocumentInternalServerException(&output, shape)
	
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	
	return output
}

func awsRestjson1_deserializeErrorThrottlingException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ThrottlingException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	err := awsRestjson1_deserializeDocumentThrottlingException(&output, shape)
	
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	
	return output
}

func awsRestjson1_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ValidationException{}
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])
	
	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	err := awsRestjson1_deserializeDocumentValidationException(&output, shape)
	
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError {
			Err: fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}
	
	errorBody.Seek(0, io.SeekStart)
	
	return output
}

func awsRestjson1_deserializeDocumentInstance(v **types.Instance, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.Instance
	if *v == nil {
		sv = &types.Instance{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "AccountID":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.AccountID = ptr.String(jtv)
				}
			
			case "AmiId":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.AmiId = ptr.String(jtv)
				}
			
			case "InstanceID":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.InstanceID = ptr.String(jtv)
				}
			
			case "InstanceType":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.InstanceType = ptr.String(jtv)
				}
			
			case "LastUpdatedTime":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.LastUpdatedTime = ptr.String(jtv)
				}
			
			case "ProductCode":
				if err := awsRestjson1_deserializeDocumentProductCodeList(&sv.ProductCode, value); err != nil {
					return err
				}
			
			case "Region":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Region = ptr.String(jtv)
				}
			
			case "Status":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Status = ptr.String(jtv)
				}
			
			case "SubscriptionName":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.SubscriptionName = ptr.String(jtv)
				}
			
			case "UsageOperation":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.UsageOperation = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentInstanceList(v *[]types.Instance, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var cv []types.Instance
	if *v == nil {
		cv = []types.Instance{}
		} else {
			cv = *v
		}
	
	for _, value := range shape {
		var col types.Instance
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentInstance(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)
		
	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentInternalServerException(v **types.InternalServerException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.InternalServerException
	if *v == nil {
		sv = &types.InternalServerException{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "message":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Message = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentLinuxSubscriptionsDiscoverySettings(v **types.LinuxSubscriptionsDiscoverySettings, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.LinuxSubscriptionsDiscoverySettings
	if *v == nil {
		sv = &types.LinuxSubscriptionsDiscoverySettings{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "OrganizationIntegration":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected OrganizationIntegration to be of type string, got %T instead", value)
					}
					sv.OrganizationIntegration = types.OrganizationIntegration(jtv)
				}
			
			case "SourceRegions":
				if err := awsRestjson1_deserializeDocumentStringList(&sv.SourceRegions, value); err != nil {
					return err
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentProductCodeList(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var cv []string
	if *v == nil {
		cv = []string{}
		} else {
			cv = *v
		}
	
	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)
		
	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentStringList(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var cv []string
	if *v == nil {
		cv = []string{}
		} else {
			cv = *v
		}
	
	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)
		
	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentStringMap(v *map[string]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
		} else {
			mv = *v
		}
	
	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal
		
	}
	*v = mv
	return nil
}

func awsRestjson1_deserializeDocumentSubscription(v **types.Subscription, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.Subscription
	if *v == nil {
		sv = &types.Subscription{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "InstanceCount":
				if value != nil {
					jtv, ok := value.(json.Number)
					if !ok {
						return fmt.Errorf("expected BoxLong to be json.Number, got %T instead", value)
					}
					i64, err := jtv.Int64()
					if err != nil { return err }
					sv.InstanceCount = ptr.Int64(i64)
				}
			
			case "Name":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Name = ptr.String(jtv)
				}
			
			case "Type":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Type = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentSubscriptionList(v *[]types.Subscription, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var cv []types.Subscription
	if *v == nil {
		cv = []types.Subscription{}
		} else {
			cv = *v
		}
	
	for _, value := range shape {
		var col types.Subscription
		destAddr := &col
		if err := awsRestjson1_deserializeDocumentSubscription(&destAddr, value); err != nil {
			return err
		}
		col = *destAddr
		cv = append(cv, col)
		
	}
	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentThrottlingException(v **types.ThrottlingException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.ThrottlingException
	if *v == nil {
		sv = &types.ThrottlingException{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "message":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Message = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentValidationException(v **types.ValidationException, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}
	
	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}
	
	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
		} else {
			sv = *v
		}
	
	for key, value := range shape {
		switch key {
			case "message":
				if value != nil {
					jtv, ok := value.(string)
					if !ok {
						return fmt.Errorf("expected String to be of type string, got %T instead", value)
					}
					sv.Message = ptr.String(jtv)
				}
			
			default:
				_, _ = key, value
			
		}
	}
	*v = sv
	return nil
}
