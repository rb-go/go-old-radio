package go_old_radio_processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var processer *Processor

func TestNewProcessor(t *testing.T) {
	processer = NewProcessor()

	assert.Equal(t, processer.EchoLevelMKV(), 0.3)
}

func TestLevelWithFading(t *testing.T) {
	expect := 1241.9879802341036
	res := levelWithFading(1028, 1, 3, 4, 5, 6)
	if res != expect {
		t.Errorf("level_with_fading expected %f returned unexpected %f", expect, res)
	}
}

func TestSignalLevelMKV(t *testing.T) {

	//if res != expect {
	//	t.Errorf("expected %f but returned unexpected %f", expect, res)
	//}
}

func TestSignalFadingFastPeriodRate(t *testing.T) {

}

func TestSignalFadingMediumPeriodRate(t *testing.T) {

}

func TestSignalFadingLongPeriodRate(t *testing.T) {

}

func TestSignalFadingExtralongPeriodRate(t *testing.T) {

}

func TestSignalFadingFastPeriodIncreaser(t *testing.T) {

}

func TestSignalFadingMediumPeriodIncreaser(t *testing.T) {

}

func TestSignalFadingLongPeriodIncreaser(t *testing.T) {

}

func TestSignalFadingExtralongPerioIncreaser(t *testing.T) {

}

func TestEchoLevelMKV(t *testing.T) {

}

func TestEchoFadingFastPeriodRate(t *testing.T) {

}

func TestEchoFadingMediumPeriodRate(t *testing.T) {

}

func TestEchoFadingLongPeriodRate(t *testing.T) {

}

func TestEchoFadingExtralongPeriodRate(t *testing.T) {

}

func TestEchoFadingFastPeriodIncreaser(t *testing.T) {

}

func TestEchoFadingMediumPeriodIncreaser(t *testing.T) {

}

func TestEchoFadingLongPeriodIncreaser(t *testing.T) {

}

func TestEchoFadingExtralongPeriodIncreaser(t *testing.T) {

}

func TestBgSignalLevelMKV(t *testing.T) {

}

func TestBgSignalFadingFastPeriodRate(t *testing.T) {

}

func TestBgSignalFadingMediumPeriodRate(t *testing.T) {

}

func TestBgSignalFadingLongPeriodRate(t *testing.T) {

}

func TestBgSignalFadingExtralongPeriodRate(t *testing.T) {

}

func TestBgSignalFadingFastPeriodIncreaser(t *testing.T) {

}

func TestBgSignalFadingMediumPeriodIncreaser(t *testing.T) {

}

func TestBgSignalFadingLongPeriodIncreaser(t *testing.T) {

}

func TestBgSignalFadingExtralongPeriodIncreaser(t *testing.T) {

}

func TestInterferenceNoiseLevelMKV(t *testing.T) {

}

func TestInterferenceNoiseFadingFastPeriodRate(t *testing.T) {

}

func TestInterferenceNoiseFadingMediumPeriodRate(t *testing.T) {

}

func TestInterferenceNoiseFadingLongPeriodRate(t *testing.T) {

}

func TestInterferenceNoiseFadingExtralongPeriodRate(t *testing.T) {

}

func TestInterferenceNoiseFadingFastPeriodIncreaser(t *testing.T) {

}

func TestInterferenceNoiseFadingMediumPeriodIncreaser(t *testing.T) {

}

func TestInterferenceNoiseFadingLongPeriodIncreaser(t *testing.T) {

}

func TestInterferenceNoiseFadingExtralongPeriodIncreaser(t *testing.T) {

}

func TestInterferenceNoiseFreqFactorRate(t *testing.T) {

}

func TestInterferenceNoiseFreqFactorIncreaser(t *testing.T) {

}

func TestInterferenceNoiseFreqFactorAdditional(t *testing.T) {

}

func TestNoiseLevelMKV(t *testing.T) {

}

/*
SignalLevelMKV()
SetSignalLevelMKV(signalLevelMKV float64)
SignalFadingFastPeriodRate()
SetSignalFadingFastPeriodRate(signalFadingFastPeriodRate float64)
SignalFadingMediumPeriodRate()
SetSignalFadingMediumPeriodRate(signalFadingMediumPeriodRate float64)
SignalFadingLongPeriodRate()
SetSignalFadingLongPeriodRate(signalFadingLongPeriodRate float64)
SignalFadingExtralongPeriodRate()
SetSignalFadingExtralongPeriodRate(signalFadingExtralongPeriodRate float64)
SignalFadingFastPeriodIncreaser()
SetSignalFadingFastPeriodIncreaser(signalFadingFastPeriodIncreaser float64)
SignalFadingMediumPeriodIncreaser()
SetSignalFadingMediumPeriodIncreaser(signalFadingMediumPeriodIncreaser float64)
SignalFadingLongPeriodIncreaser()
SetSignalFadingLongPeriodIncreaser(signalFadingLongPeriodIncreaser float64)
SignalFadingExtralongPerioIncreaser()
SetSignalFadingExtralongPerioIncreaser(signalFadingExtralongPerioIncreaser float64)
EchoLevelMKV()
SetEchoLevelMKV(echoLevelMKV float64)
EchoFadingFastPeriodRate()
SetEchoFadingFastPeriodRate(echoFadingFastPeriodRate float64)
EchoFadingMediumPeriodRate()
SetEchoFadingMediumPeriodRate(echoFadingMediumPeriodRate float64)
EchoFadingLongPeriodRate()
SetEchoFadingLongPeriodRate(echoFadingLongPeriodRate float64)
EchoFadingExtralongPeriodRate()
SetEchoFadingExtralongPeriodRate(echoFadingExtralongPeriodRate float64)
EchoFadingFastPeriodIncreaser()
SetEchoFadingFastPeriodIncreaser(echoFadingFastPeriodIncreaser float64)
EchoFadingMediumPeriodIncreaser()
SetEchoFadingMediumPeriodIncreaser(echoFadingMediumPeriodIncreaser float64)
EchoFadingLongPeriodIncreaser()
SetEchoFadingLongPeriodIncreaser(echoFadingLongPeriodIncreaser float64)
EchoFadingExtralongPeriodIncreaser()
SetEchoFadingExtralongPeriodIncreaser(echoFadingExtralongPeriodIncreaser float64)
BgSignalLevelMKV()
SetBgSignalLevelMKV(bgSignalLevelMKV float64)
BgSignalFadingFastPeriodRate()
SetBgSignalFadingFastPeriodRate(bgSignalFadingFastPeriodRate float64)
BgSignalFadingMediumPeriodRate()
SetBgSignalFadingMediumPeriodRate(bgSignalFadingMediumPeriodRate float64)
BgSignalFadingLongPeriodRate()
SetBgSignalFadingLongPeriodRate(bgSignalFadingLongPeriodRate float64)
BgSignalFadingExtralongPeriodRate()
SetBgSignalFadingExtralongPeriodRate(bgSignalFadingExtralongPeriodRate float64)
BgSignalFadingFastPeriodIncreaser()
SetBgSignalFadingFastPeriodIncreaser(bgSignalFadingFastPeriodIncreaser float64)
BgSignalFadingMediumPeriodIncreaser()
SetBgSignalFadingMediumPeriodIncreaser(bgSignalFadingMediumPeriodIncreaser float64)
BgSignalFadingLongPeriodIncreaser()
SetBgSignalFadingLongPeriodIncreaser(bgSignalFadingLongPeriodIncreaser float64)
BgSignalFadingExtralongPeriodIncreaser()
SetBgSignalFadingExtralongPeriodIncreaser(bgSignalFadingExtralongPeriodIncreaser float64)
InterferenceNoiseLevelMKV()
SetInterferenceNoiseLevelMKV(interferenceNoiseLevelMKV float64)
InterferenceNoiseFadingFastPeriodRate()
SetInterferenceNoiseFadingFastPeriodRate(interferenceNoiseFadingFastPeriodRate float64)
InterferenceNoiseFadingMediumPeriodRate()
SetInterferenceNoiseFadingMediumPeriodRate(interferenceNoiseFadingMediumPeriodRate float64)
InterferenceNoiseFadingLongPeriodRate()
SetInterferenceNoiseFadingLongPeriodRate(interferenceNoiseFadingLongPeriodRate float64)
InterferenceNoiseFadingExtralongPeriodRate()
SetInterferenceNoiseFadingExtralongPeriodRate(interferenceNoiseFadingExtralongPeriodRate float64)
InterferenceNoiseFadingFastPeriodIncreaser()
SetInterferenceNoiseFadingFastPeriodIncreaser(interferenceNoiseFadingFastPeriodIncreaser float64)
InterferenceNoiseFadingMediumPeriodIncreaser()
SetInterferenceNoiseFadingMediumPeriodIncreaser(interferenceNoiseFadingMediumPeriodIncreaser float64)
InterferenceNoiseFadingLongPeriodIncreaser()
SetInterferenceNoiseFadingLongPeriodIncreaser(interferenceNoiseFadingLongPeriodIncreaser float64)
InterferenceNoiseFadingExtralongPeriodIncreaser()
SetInterferenceNoiseFadingExtralongPeriodIncreaser(interferenceNoiseFadingExtralongPeriodIncreaser float64)
InterferenceNoiseFreqFactorRate()
SetInterferenceNoiseFreqFactorRate(interferenceNoiseFreqFactorRate float64)
InterferenceNoiseFreqFactorIncreaser()
SetInterferenceNoiseFreqFactorIncreaser(interferenceNoiseFreqFactorIncreaser float64)
InterferenceNoiseFreqFactorAdditional()
SetInterferenceNoiseFreqFactorAdditional(interferenceNoiseFreqFactorAdditional float64)
NoiseLevelMKV()
SetNoiseLevelMKV
*/
