// Package crypto implements LoRaWAN crypto.
package crypto

func reverse(in []byte) []byte {
	l := len(in)
	out := make([]byte, l)
	for i := 0; i < l; i++ {
		out[l-i-1] = in[i]
	}
	return out
}
