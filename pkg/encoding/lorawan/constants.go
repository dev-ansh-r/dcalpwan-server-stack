package lorawan

const (
	// JoinAcceptWithoutCFListLength represents the length of a JoinAccept binary message which does
	// not contain a CFList field. It already takes into consideration the presence of the MHDR and the MIC.
	JoinAcceptWithoutCFListLength = 17
	// JoinAcceptWithCFListLength represents the length of a JoinAccept binary message which contains
	// a CFList field. It already takes into consideration the presence of the MHDR and the MIC.
	JoinAcceptWithCFListLength = JoinAcceptWithoutCFListLength + 16
)
