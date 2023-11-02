package interop

import (
	"encoding/json"
	"net/http"

	ttnerrors "go.thethings.network/lorawan-stack/v3/pkg/errors"
)

func writeError(w http.ResponseWriter, _ *http.Request, header MessageHeader, err error) {
	answerHeader, headerErr := header.AnswerHeader()
	if headerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code, desc := ResultOther, err.Error()
	for errDef, protocolCode := range errorResultCodes {
		if ttnerrors.Resemble(errDef, err) {
			if c, ok := protocolCode[header.ProtocolVersion]; ok {
				code = c
				if cause := ttnerrors.Cause(err); cause != nil {
					desc = cause.Error()
				} else {
					desc = ""
				}
			}
			break
		}
	}

	msg := ErrorMessage{
		MessageHeader: answerHeader,
		Result: Result{
			ResultCode:  code,
			Description: desc,
		},
	}
	json.NewEncoder(w).Encode(msg) //nolint:errcheck
}
