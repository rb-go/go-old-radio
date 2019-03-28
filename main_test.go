package goorp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var processor *Processor

func TestNewProcessor(t *testing.T) {
	processor = NewProcessor()
	assert.Equal(t, 0.3, processor.EchoLevelMKV())
}

func TestLevelWithFading(t *testing.T) {
	res := levelWithFading(1028, 1, 3, 4, 5, 6)
	expect := 1241.9879802341036
	assert.Equal(t, expect, res)
}

func TestProcessor_ProcessPipes(t *testing.T) {

	mainStream, err := os.Open("./test_audio_files/ringtone.mp3") // For read access.
	defer mainStream.Close()
	assert.NoError(t, err)

	bgStream, err := os.Open("./test_audio_files/ringtone_long.mp3") // For read access.
	defer bgStream.Close()
	assert.NoError(t, err)

	errChan := make(chan error)
	outChan := make(chan byte)

	go processor.ProcessPipes(mainStream, bgStream, 1221, outChan, errChan)

	var cycleBreak bool
	for {
		if cycleBreak {
			break
		}
		select {
		case err := <-errChan:
			assert.Error(t, err, "main pipe ended")
			cycleBreak = true
			break
		case out := <-outChan:
			assert.NotNil(t, out)
			break
		}
	}
}

func TestProcessor_ProcessPipes2(t *testing.T) {

	mainStream, err := os.Open("./test_audio_files/ringtone_long.mp3") // For read access.
	defer mainStream.Close()
	assert.NoError(t, err)

	bgStream, err := os.Open("./test_audio_files/ringtone.mp3") // For read access.
	defer bgStream.Close()
	assert.NoError(t, err)

	errChan := make(chan error)
	outChan := make(chan byte)

	go processor.ProcessPipes(mainStream, bgStream, 1028, outChan, errChan)

	var cycleBreak bool
	for {
		if cycleBreak {
			break
		}
		select {
		case err := <-errChan:
			assert.Error(t, err, ErrBackgroundPipeEnded)
			cycleBreak = true
			break
		case out := <-outChan:
			assert.NotNil(t, out)
			break
		}
	}
}

func TestSignalLevelMKV(t *testing.T) {
	expect := 10.1
	processor.SetSignalLevelMKV(expect)
	res := processor.SignalLevelMKV()
	assert.Equal(t, expect, res)
}

func TestSignalFadingFastPeriodRate(t *testing.T) {
	expect := 1000.1
	processor.SetSignalFadingFastPeriodRate(expect)
	res := processor.SignalFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingMediumPeriodRate(t *testing.T) {
	expect := 15000.1
	processor.SetSignalFadingMediumPeriodRate(expect)
	res := processor.SignalFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingLongPeriodRate(t *testing.T) {
	expect := 40000.1
	processor.SetSignalFadingLongPeriodRate(expect)
	res := processor.SignalFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processor.SetSignalFadingExtralongPeriodRate(expect)
	res := processor.SignalFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestSignalFadingFastPeriodIncreaser(t *testing.T) {
	expect := 300.1
	processor.SetSignalFadingFastPeriodIncreaser(expect)
	res := processor.SignalFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 17500.1
	processor.SetSignalFadingMediumPeriodIncreaser(expect)
	res := processor.SignalFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingLongPeriodIncreaser(t *testing.T) {
	expect := 60000.1
	processor.SetSignalFadingLongPeriodIncreaser(expect)
	res := processor.SignalFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestSignalFadingExtralongPerioIncreaser(t *testing.T) {
	expect := 150000.1
	processor.SetSignalFadingExtralongPerioIncreaser(expect)
	res := processor.SignalFadingExtralongPerioIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoLevelMKV(t *testing.T) {
	expect := 0.31
	processor.SetEchoLevelMKV(expect)
	res := processor.EchoLevelMKV()
	assert.Equal(t, expect, res)
}

func TestEchoFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processor.SetEchoFadingFastPeriodRate(expect)
	res := processor.EchoFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processor.SetEchoFadingMediumPeriodRate(expect)
	res := processor.EchoFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processor.SetEchoFadingLongPeriodRate(expect)
	res := processor.EchoFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processor.SetEchoFadingExtralongPeriodRate(expect)
	res := processor.EchoFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestEchoFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processor.SetEchoFadingFastPeriodIncreaser(expect)
	res := processor.EchoFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processor.SetEchoFadingMediumPeriodIncreaser(expect)
	res := processor.EchoFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processor.SetEchoFadingLongPeriodIncreaser(expect)
	res := processor.EchoFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestEchoFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processor.SetEchoFadingExtralongPeriodIncreaser(expect)
	res := processor.EchoFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalLevelMKV(t *testing.T) {
	expect := 0.21
	processor.SetBgSignalLevelMKV(expect)
	res := processor.BgSignalLevelMKV()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processor.SetBgSignalFadingFastPeriodRate(expect)
	res := processor.BgSignalFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processor.SetBgSignalFadingMediumPeriodRate(expect)
	res := processor.BgSignalFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processor.SetBgSignalFadingLongPeriodRate(expect)
	res := processor.BgSignalFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processor.SetBgSignalFadingExtralongPeriodRate(expect)
	res := processor.BgSignalFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processor.SetBgSignalFadingFastPeriodIncreaser(expect)
	res := processor.BgSignalFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processor.SetBgSignalFadingMediumPeriodIncreaser(expect)
	res := processor.BgSignalFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processor.SetBgSignalFadingLongPeriodIncreaser(expect)
	res := processor.BgSignalFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestBgSignalFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processor.SetBgSignalFadingExtralongPeriodIncreaser(expect)
	res := processor.BgSignalFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseLevelMKV(t *testing.T) {
	expect := 0.021
	processor.SetInterferenceNoiseLevelMKV(expect)
	res := processor.InterferenceNoiseLevelMKV()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingFastPeriodRate(t *testing.T) {
	expect := 2000.1
	processor.SetInterferenceNoiseFadingFastPeriodRate(expect)
	res := processor.InterferenceNoiseFadingFastPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingMediumPeriodRate(t *testing.T) {
	expect := 7000.1
	processor.SetInterferenceNoiseFadingMediumPeriodRate(expect)
	res := processor.InterferenceNoiseFadingMediumPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingLongPeriodRate(t *testing.T) {
	expect := 50000.1
	processor.SetInterferenceNoiseFadingLongPeriodRate(expect)
	res := processor.InterferenceNoiseFadingLongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingExtralongPeriodRate(t *testing.T) {
	expect := 200000.1
	processor.SetInterferenceNoiseFadingExtralongPeriodRate(expect)
	res := processor.InterferenceNoiseFadingExtralongPeriodRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingFastPeriodIncreaser(t *testing.T) {
	expect := 1500.1
	processor.SetInterferenceNoiseFadingFastPeriodIncreaser(expect)
	res := processor.InterferenceNoiseFadingFastPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingMediumPeriodIncreaser(t *testing.T) {
	expect := 4000.1
	processor.SetInterferenceNoiseFadingMediumPeriodIncreaser(expect)
	res := processor.InterferenceNoiseFadingMediumPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingLongPeriodIncreaser(t *testing.T) {
	expect := 25000.1
	processor.SetInterferenceNoiseFadingLongPeriodIncreaser(expect)
	res := processor.InterferenceNoiseFadingLongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFadingExtralongPeriodIncreaser(t *testing.T) {
	expect := 150000.1
	processor.SetInterferenceNoiseFadingExtralongPeriodIncreaser(expect)
	res := processor.InterferenceNoiseFadingExtralongPeriodIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorRate(t *testing.T) {
	expect := 1000.1
	processor.SetInterferenceNoiseFreqFactorRate(expect)
	res := processor.InterferenceNoiseFreqFactorRate()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorIncreaser(t *testing.T) {
	expect := 1.01
	processor.SetInterferenceNoiseFreqFactorIncreaser(expect)
	res := processor.InterferenceNoiseFreqFactorIncreaser()
	assert.Equal(t, expect, res)
}

func TestInterferenceNoiseFreqFactorAdditional(t *testing.T) {
	expect := 3000.1
	processor.SetInterferenceNoiseFreqFactorAdditional(expect)
	res := processor.InterferenceNoiseFreqFactorAdditional()
	assert.Equal(t, expect, res)
}

func TestNoiseLevelMKV(t *testing.T) {
	expect := 0.11
	processor.SetNoiseLevelMKV(expect)
	res := processor.NoiseLevelMKV()
	assert.Equal(t, expect, res)
}
