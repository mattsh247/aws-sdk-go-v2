// Code generated by smithy-go-codegen DO NOT EDIT.


package migrationhubconfig

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/smithy-go/encoding/httpbinding"
	"github.com/aws/smithy-go/middleware"
	"path"
	smithy "github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig/types"
)

type awsAwsjson11_serializeOpCreateHomeRegionControl struct {
}

func (*awsAwsjson11_serializeOpCreateHomeRegionControl) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateHomeRegionControl) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}
	
	input, ok := in.Parameters.(*CreateHomeRegionControlInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}
	
	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
	    request.Request.URL.Path = operationPath
	} else {
	    request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
	    if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
	        request.Request.URL.Path += "/"
	    }
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMigrationHubMultiAccountService.CreateHomeRegionControl")
	
	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentCreateHomeRegionControlInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request
	
	return next.HandleSerialize(ctx, in)
}
type awsAwsjson11_serializeOpDescribeHomeRegionControls struct {
}

func (*awsAwsjson11_serializeOpDescribeHomeRegionControls) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeHomeRegionControls) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}
	
	input, ok := in.Parameters.(*DescribeHomeRegionControlsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}
	
	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
	    request.Request.URL.Path = operationPath
	} else {
	    request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
	    if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
	        request.Request.URL.Path += "/"
	    }
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMigrationHubMultiAccountService.DescribeHomeRegionControls")
	
	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeHomeRegionControlsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request
	
	return next.HandleSerialize(ctx, in)
}
type awsAwsjson11_serializeOpGetHomeRegion struct {
}

func (*awsAwsjson11_serializeOpGetHomeRegion) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGetHomeRegion) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}
	
	input, ok := in.Parameters.(*GetHomeRegionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}
	
	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
	    request.Request.URL.Path = operationPath
	} else {
	    request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
	    if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
	        request.Request.URL.Path += "/"
	    }
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMigrationHubMultiAccountService.GetHomeRegion")
	
	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentGetHomeRegionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	
	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request
	
	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentTarget(v *types.Target, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()
	
	if v.Id != nil {
		ok := object.Key("Id")
		ok.String(*v.Id)
	}
	
	if len(v.Type) > 0 {
		ok := object.Key("Type")
		ok.String(string(v.Type))
	}
	
	return nil
}

func awsAwsjson11_serializeOpDocumentCreateHomeRegionControlInput(v *CreateHomeRegionControlInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()
	
	if v.DryRun {
		ok := object.Key("DryRun")
		ok.Boolean(v.DryRun)
	}
	
	if v.HomeRegion != nil {
		ok := object.Key("HomeRegion")
		ok.String(*v.HomeRegion)
	}
	
	if v.Target != nil {
		ok := object.Key("Target")
		if err := awsAwsjson11_serializeDocumentTarget(v.Target, ok); err != nil {
			return err
		}
	}
	
	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeHomeRegionControlsInput(v *DescribeHomeRegionControlsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()
	
	if v.ControlId != nil {
		ok := object.Key("ControlId")
		ok.String(*v.ControlId)
	}
	
	if v.HomeRegion != nil {
		ok := object.Key("HomeRegion")
		ok.String(*v.HomeRegion)
	}
	
	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}
	
	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}
	
	if v.Target != nil {
		ok := object.Key("Target")
		if err := awsAwsjson11_serializeDocumentTarget(v.Target, ok); err != nil {
			return err
		}
	}
	
	return nil
}

func awsAwsjson11_serializeOpDocumentGetHomeRegionInput(v *GetHomeRegionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()
	
	return nil
}
