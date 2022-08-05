// Code generated by gen.go. DO NOT EDIT.

package logpb

import (
	"strconv"

	"github.com/cockroachdb/cockroach/pkg/util/jsonbytes"
	"github.com/cockroachdb/redact"
)

// AppendJSONFields implements the EventPayload interface.
func (m *CommonEventDetails) AppendJSONFields(printComma bool, b redact.RedactableBytes) (bool, redact.RedactableBytes) {

	if m.Timestamp != 0 {
		if printComma {
			b = append(b, ',')
		}
		printComma = true
		b = append(b, "\"Timestamp\":"...)
		b = strconv.AppendInt(b, int64(m.Timestamp), 10)
	}

	if m.EventType != "" {
		if printComma {
			b = append(b, ',')
		}
		printComma = true
		b = append(b, "\"EventType\":\""...)
		b = redact.RedactableBytes(jsonbytes.EncodeString([]byte(b), string(m.EventType)))
		b = append(b, '"')
	}

	return printComma, b
}
