package band

var (
	ParseChMask   = parseChMask
	ParseChMask16 = parseChMask16
	ParseChMask72 = parseChMask72
	ParseChMask96 = parseChMask96

	GenerateChMask16     = generateChMask16
	GenerateChMask96     = generateChMask96
	MakeGenerateChMask72 = makeGenerateChMask72

	ErrUnsupportedChMaskCntl = errUnsupportedChMaskCntl

	CompareDataRates = compareDataRates

	BoolPtr = boolPtr
)
