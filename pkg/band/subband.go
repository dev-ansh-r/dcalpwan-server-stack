package band

import "time"

// SubBandParameters contains the sub-band frequency range, duty cycle and Tx power.
type SubBandParameters struct {
	MinFrequency uint64
	MaxFrequency uint64
	DutyCycle    float32
	MaxEIRP      float32
}

// Comprises returns whether the duty cycle applies to the given frequency.
func (d SubBandParameters) Comprises(frequency uint64) bool {
	return frequency >= d.MinFrequency && frequency <= d.MaxFrequency
}

// MaxEmissionDuring the period passed as parameter, that is allowed by that duty cycle.
func (d SubBandParameters) MaxEmissionDuring(period time.Duration) time.Duration {
	return time.Duration(d.DutyCycle * float32(period))
}
