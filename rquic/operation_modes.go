package rquic


var BTOOnly = false // bT
var BTOtoRTT float64 = 0 // 10/3 // * sRTT, will use fixed BTO if <= 0
var BTOMargin = 1   // bM?


var PauseEncodingWith = PauseEncodingNever
const(
	PauseEncodingNever            = iota // pN
	PauseEncodingTillFirstLoss           // pS
	PauseEncodingWithResidualLoss        // pL
	// What else?
)
var PauseEncodingStr = []string{
	"NoPause",
	"PauseTillFirstLoss",
	"ResidualLoss",
}
func PauseEncodingExplained() string {
	return PauseEncodingStr[PauseEncodingWith]
}

// if PauseEncodingWith == PauseEncodingWithResidualLoss
var ResLossFactor = 0.0 // residualOff = residualTarget * ResLossFactor // pL{?*10}



var LimRateToDecBuffer = false	// rL
