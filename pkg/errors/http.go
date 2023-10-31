package errors

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
)

// httpStatusCodes maps status codes to HTTP codes.
// See package google.golang.org/genproto/googleapis/rpc/code and google.golang.org/grpc/codes for details.
var httpStatusCodes = map[uint32]int{
	uint32(codes.OK):                 http.StatusOK,
	uint32(codes.Canceled):           499, // Client Closed Request
	uint32(codes.Unknown):            http.StatusInternalServerError,
	uint32(codes.InvalidArgument):    http.StatusBadRequest,
	uint32(codes.DeadlineExceeded):   http.StatusGatewayTimeout,
	uint32(codes.NotFound):           http.StatusNotFound,
	uint32(codes.AlreadyExists):      http.StatusConflict,
	uint32(codes.PermissionDenied):   http.StatusForbidden,
	uint32(codes.Unauthenticated):    http.StatusUnauthorized,
	uint32(codes.ResourceExhausted):  http.StatusTooManyRequests,
	uint32(codes.FailedPrecondition): http.StatusBadRequest,
	uint32(codes.Aborted):            http.StatusConflict,
	uint32(codes.OutOfRange):         http.StatusBadRequest,
	uint32(codes.Unimplemented):      http.StatusNotImplemented,
	uint32(codes.Internal):           http.StatusInternalServerError,
	uint32(codes.Unavailable):        http.StatusServiceUnavailable,
	uint32(codes.DataLoss):           http.StatusInternalServerError,
}

// httpErrorCodes maps HTTP codes to status codes.
// See package google.golang.org/genproto/googleapis/rpc/code and google.golang.org/grpc/codes for details.
var httpErrorCodes = map[int]uint32{
	http.StatusOK:                  uint32(codes.OK),
	499:                            uint32(codes.Canceled), // Client Closed Request
	http.StatusInternalServerError: uint32(codes.Unknown),
	http.StatusBadRequest:          uint32(codes.InvalidArgument),
	http.StatusGatewayTimeout:      uint32(codes.DeadlineExceeded),
	http.StatusNotFound:            uint32(codes.NotFound),
	http.StatusConflict:            uint32(codes.AlreadyExists),
	http.StatusForbidden:           uint32(codes.PermissionDenied),
	http.StatusUnauthorized:        uint32(codes.Unauthenticated),
	http.StatusTooManyRequests:     uint32(codes.ResourceExhausted),
	http.StatusNotImplemented:      uint32(codes.Unimplemented),
	http.StatusServiceUnavailable:  uint32(codes.Unavailable),
}

// ToHTTPStatusCode maps an error to HTTP response codes.
func ToHTTPStatusCode(err error) int {
	for _, err := range Stack(err) {
		code := Code(err)
		if code == uint32(codes.Unknown) {
			continue
		}
		if status, ok := httpStatusCodes[code]; ok {
			return status
		}
	}
	return http.StatusInternalServerError
}

// FromHTTPStatusCode maps an HTTP response code to an error.
func FromHTTPStatusCode(status int, publicAttributes ...string) *Error {
	code := uint32(codes.Unknown)
	if c, ok := httpErrorCodes[status]; ok {
		code = c
	}
	return build(&Definition{
		namespace:        namespace(2),
		messageFormat:    http.StatusText(status),
		publicAttributes: publicAttributes,
		code:             code,
	}, 4)
}

// ToHTTP writes the error to the HTTP response.
func ToHTTP(in error, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if ttnErr, ok := From(in); ok {
		w.WriteHeader(ToHTTPStatusCode(ttnErr))
		return json.NewEncoder(w).Encode(ttnErr)
	}
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(in)
}

// FromHTTP reads an error from the HTTP response.
func FromHTTP(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	defer resp.Body.Close()
	var err Error
	if decErr := json.NewDecoder(resp.Body).Decode(&err); decErr != nil {
		return decErr
	}
	return &err
}
