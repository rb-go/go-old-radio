package go_old_radio_processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var processer *Processor

func TestNewProcessor(t *testing.T) {
	processer = NewProcessor()
	assert.Equal(t, 0.3, processer.EchoLevelMKV())
}

func TestLevelWithFading(t *testing.T) {
	res := levelWithFading(1028, 1, 3, 4, 5, 6)
	expect := 1241.9879802341036
	assert.Equal(t, expect, res)
}

func TestSignalLevelMKV(t *testing.T) {
	expect := 10.1
	processer.SetSignalLevelMKV(expect)
	res := processer.SignalLevelMKV()
	assert.Equal(t, expect, res)
}

func TestSignalFadingFastPeriodRate(t *testing.T) {
	expect := 1000.1
	processer.SetSignalFadingFastPeriodRate(expect)
	res := processer.SignalFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingMediumPeriodRate(t *testing.T) {
	expect := 15000.1
	processer.SetSignalFadingMediumPeriodRate(expect)
	res := processer.SignalFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingLongPeriodRate(t *testing.T) {
	expect := 40000.1
	processer.SetSignalFadingLongPeriodRate(expect)
	res := processer.SignalFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processer.SetSignalFadingExtralongPeriodRate(expect)
	res := processer.SignalFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingFastPeriodIncreaser(t *testing.T) {
	expect := 300.1
	processer.SetSignalFadingFastPeriodIncreaser(expect)
	res := processer.SignalFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 17500.1
	processer.SetSignalFadingMediumPeriodIncreaser(expect)
	res := processer.SignalFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingLongPeriodIncreaser(t *testing.T) {
	expect := 60000.1
	processer.SetSignalFadingLongPeriodIncreaser(expect)
	res := processer.SignalFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingExtralongPerioIncreaser(t *testing.T) {
	expect := 150000.1
	processer.SetSignalFadingExtralongPerioIncreaser(expect)
	res := processer.SignalFadingExtralongPerioIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoLevelMKV(t *testing.T) {
	expect := 0.31
	processer.SetEchoLevelMKV(expect)
	res := processer.EchoLevelMKV()
	assert.Equal(t, expect, res)
}

func TestEchoFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processer.SetEchoFadingFastPeriodRate(expect)
	res := processer.EchoFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processer.SetEchoFadingMediumPeriodRate(expect)
	res := processer.EchoFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processer.SetEchoFadingLongPeriodRate(expect)
	res := processer.EchoFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processer.SetEchoFadingExtralongPeriodRate(expect)
	res := processer.EchoFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processer.SetEchoFadingFastPeriodIncreaser(expect)
	res := processer.EchoFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processer.SetEchoFadingMediumPeriodIncreaser(expect)
	res := processer.EchoFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processer.SetEchoFadingLongPeriodIncreaser(expect)
	res := processer.EchoFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processer.SetEchoFadingExtralongPeriodIncreaser(expect)
	res := processer.EchoFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalLevelMKV(t *testing.T) {
	expect := 0.21
	processer.SetBgSignalLevelMKV(expect)
	res := processer.BgSignalLevelMKV()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processer.SetBgSignalFadingFastPeriodRate(expect)
	res := processer.BgSignalFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processer.SetBgSignalFadingMediumPeriodRate(expect)
	res := processer.BgSignalFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processer.SetBgSignalFadingLongPeriodRate(expect)
	res := processer.BgSignalFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processer.SetBgSignalFadingExtralongPeriodRate(expect)
	res := processer.BgSignalFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processer.SetBgSignalFadingFastPeriodIncreaser(expect)
	res := processer.BgSignalFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processer.SetBgSignalFadingMediumPeriodIncreaser(expect)
	res := processer.BgSignalFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processer.SetBgSignalFadingLongPeriodIncreaser(expect)
	res := processer.BgSignalFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processer.SetBgSignalFadingExtralongPeriodIncreaser(expect)
	res := processer.BgSignalFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseLevelMKV(t *testing.T) {
	expect := 0.021
	processer.SetInterferenceNoiseLevelMKV(expect)
	res := processer.InterferenceNoiseLevelMKV()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processer.SetInterferenceNoiseFadingFastPeriodRate(expect)
	res := processer.InterferenceNoiseFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processer.SetInterferenceNoiseFadingMediumPeriodRate(expect)
	res := processer.InterferenceNoiseFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processer.SetInterferenceNoiseFadingLongPeriodRate(expect)
	res := processer.InterferenceNoiseFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processer.SetInterferenceNoiseFadingExtralongPeriodRate(expect)
	res := processer.InterferenceNoiseFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processer.SetInterferenceNoiseFadingFastPeriodIncreaser(expect)
	res := processer.InterferenceNoiseFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processer.SetInterferenceNoiseFadingMediumPeriodIncreaser(expect)
	res := processer.InterferenceNoiseFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processer.SetInterferenceNoiseFadingLongPeriodIncreaser(expect)
	res := processer.InterferenceNoiseFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processer.SetInterferenceNoiseFadingExtralongPeriodIncreaser(expect)
	res := processer.InterferenceNoiseFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorRate(t *testing.T) {
	expect := 1000.1
	processer.SetInterferenceNoiseFreqFactorRate(expect)
	res := processer.InterferenceNoiseFreqFactorRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorIncreaser(t *testing.T) {
	expect := 1.01
	processer.SetInterferenceNoiseFreqFactorIncreaser(expect)
	res := processer.InterferenceNoiseFreqFactorIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorAdditional(t *testing.T) {
	expect := 3000.1
	processer.SetInterferenceNoiseFreqFactorAdditional(expect)
	res := processer.InterferenceNoiseFreqFactorAdditional()
	assert.Equal(t, expect, res)
}

func TestNoiseLevelMKV(t *testing.T) {
	expect := 0.11
	processer.SetNoiseLevelMKV(expect)
	res := processer.NoiseLevelMKV()
	assert.Equal(t, expect, res)
}
