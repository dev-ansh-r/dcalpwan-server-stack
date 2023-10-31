package fetch

import "go.thethings.network/lorawan-stack/v3/pkg/errors"

var (
	errCouldNotFetchFile    = errors.Define("fetch_file", "fetch file `{filename}`")
	errCouldNotReadFile     = errors.DefineCorruption("read_file", "read file `{filename}`")
	errFilenameNotSpecified = errors.DefineInvalidArgument("filename_not_specified", "filename not specified")
	errFileNotFound         = errors.DefineNotFound("file_not_found", "file `{filename}` not found")
	errSchemeNotSpecified   = errors.DefineInvalidArgument("scheme_not_specified", "URI scheme not specified")
	errSchemeSpecified      = errors.DefineInvalidArgument("scheme_specified", "URI scheme should not be specified")
	errVolumeSpecified      = errors.DefineInvalidArgument("volume_specified", "volume should not be specified")
)
