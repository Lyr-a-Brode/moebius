// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package client

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Error Object with error type and description
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// XTraceID defines model for xTraceID.
type XTraceID = openapi_types.UUID

// ErrorBadRequest Object with error type and description
type ErrorBadRequest = Error

// ErrorInternal Object with error type and description
type ErrorInternal = Error

// UploadBlobSuccessResponse defines model for UploadBlobSuccessResponse.
type UploadBlobSuccessResponse struct {
	BlobId string `json:"blob_id"`
}

// UploadBlobMultipartBody defines parameters for UploadBlob.
type UploadBlobMultipartBody struct {
	File   openapi_types.File `json:"file"`
	Format string             `json:"format"`
}

// UploadBlobParams defines parameters for UploadBlob.
type UploadBlobParams struct {
	// XTraceID Globally unique identifier of the request in UUID v4 format
	XTraceID XTraceID `json:"X-Trace-ID"`
}

// UploadBlobMultipartRequestBody defines body for UploadBlob for multipart/form-data ContentType.
type UploadBlobMultipartRequestBody UploadBlobMultipartBody