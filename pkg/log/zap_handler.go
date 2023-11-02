package log

import (
	"os"
	"sort"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZap returns a new log handler based on the zap library.
func NewZap(encoding string) (Handler, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // levels are currently still filtered by the Logger.
		Encoding:         encoding,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	config.DisableCaller = true       // omit caller from log message.
	config.EncoderConfig.TimeKey = "" // omit timestamp from log message.
	if encoding == "console" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		colorterm, _ := strconv.ParseBool(os.Getenv("COLORTERM"))
		if !colorterm && os.Getenv("COLORTERM") == "truecolor" {
			colorterm = true
		}
		if colorterm {
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
	}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &zapHandler{Logger: logger}, nil
}

type zapHandler struct {
	*zap.Logger
}

func (h *zapHandler) HandleLog(e Entry) error {
	var logFunc func(msg string, fields ...zap.Field)
	switch e.Level() {
	case DebugLevel:
		logFunc = h.Logger.Debug
	case InfoLevel:
		logFunc = h.Logger.Info
	case WarnLevel:
		logFunc = h.Logger.Warn
	case ErrorLevel:
		logFunc = h.Logger.Error
	case FatalLevel:
		logFunc = h.Logger.Fatal
	default:
		return nil
	}
	logFields := e.Fields().Fields()
	var fieldNames []string
	for k := range logFields {
		fieldNames = append(fieldNames, k)
	}
	sort.Strings(fieldNames)
	var zapFields []zap.Field
	for _, k := range fieldNames {
		zapFields = append(zapFields, zap.Any(k, logFields[k]))
	}
	logFunc(e.Message(), zapFields...)
	return nil
}
