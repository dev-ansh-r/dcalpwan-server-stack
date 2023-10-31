package redis

import (
	"github.com/redis/go-redis/v9"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

var (
	errDecode              = errors.Define("decode", "decode value")
	errEncode              = errors.Define("encode", "encode value")
	errInvalidKeyValueType = errors.DefineInvalidArgument("value_type", "invalid value type for key `{key}`")
	errNoArguments         = errors.DefineInvalidArgument("no_arguments", "no arguments")
	errNotFound            = errors.DefineNotFound("not_found", "entity not found")
	errStore               = errors.Define("store", "store error")
	errTransactionFailed   = errors.DefineAborted("transaction_failed", "transaction failed")
)

// ConvertError converts Redis error into errors.Error.
func ConvertError(err error) error {
	ttnErr, ok := errors.From(err)
	if ok {
		return ttnErr
	}
	switch err {
	case nil:
		return nil
	case redis.Nil:
		return errNotFound.New()
	case redis.TxFailedErr:
		return errTransactionFailed.New()
	}
	return errStore.WithCause(err)
}
