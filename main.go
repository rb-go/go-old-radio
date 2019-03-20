package oldradio

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"sync"
)

const echoSizeSamples float64 = 1000

func levelWithFading(level, t, fastPeriod, mediumPeriod, longPeriod, extralongPeriod float64) float64 {
	return level + 0.4*level*math.Sin(t/extralongPeriod) + 0.3*level*math.Sin(t/longPeriod) + 0.2*level*math.Sin(t/mediumPeriod) + 0.1*level*math.Sin(t/fastPeriod)
}

//Processor ...
type Processor struct {
	sync.Mutex

	signalLevelMKV                      float64
	signalFadingFastPeriodRate          float64
	signalFadingMediumPeriodRate        float64
	signalFadingLongPeriodRate          float64
	signalFadingExtralongPeriodRate     float64
	signalFadingFastPeriodIncreaser     float64
	signalFadingMediumPeriodIncreaser   float64
	signalFadingLongPeriodIncreaser     float64
	signalFadingExtralongPerioIncreaser float64

	echoLevelMKV                       float64
	echoFadingFastPeriodRate           float64
	echoFadingMediumPeriodRate         float64
	echoFadingLongPeriodRate           float64
	echoFadingExtralongPeriodRate      float64
	echoFadingFastPeriodIncreaser      float64
	echoFadingMediumPeriodIncreaser    float64
	echoFadingLongPeriodIncreaser      float64
	echoFadingExtralongPeriodIncreaser float64

	bgSignalLevelMKV                       float64
	bgSignalFadingFastPeriodRate           float64
	bgSignalFadingMediumPeriodRate         float64
	bgSignalFadingLongPeriodRate           float64
	bgSignalFadingExtralongPeriodRate      float64
	bgSignalFadingFastPeriodIncreaser      float64
	bgSignalFadingMediumPeriodIncreaser    float64
	bgSignalFadingLongPeriodIncreaser      float64
	bgSignalFadingExtralongPeriodIncreaser float64

	interferenceNoiseLevelMKV                       float64
	interferenceNoiseFadingFastPeriodRate           float64
	interferenceNoiseFadingMediumPeriodRate         float64
	interferenceNoiseFadingLongPeriodRate           float64
	interferenceNoiseFadingExtralongPeriodRate      float64
	interferenceNoiseFadingFastPeriodIncreaser      float64
	interferenceNoiseFadingMediumPeriodIncreaser    float64
	interferenceNoiseFadingLongPeriodIncreaser      float64
	interferenceNoiseFadingExtralongPeriodIncreaser float64

	interferenceNoiseFreqFactorRate       float64
	interferenceNoiseFreqFactorIncreaser  float64
	interferenceNoiseFreqFactorAdditional float64

	// White (or brown) noise
	noiseLevelMKV float64
}

//NewProcessor ...
func NewProcessor() *Processor {
	proc := &Processor{
		signalLevelMKV:                      10.0,
		signalFadingFastPeriodRate:          1000,
		signalFadingMediumPeriodRate:        15000,
		signalFadingLongPeriodRate:          40000,
		signalFadingExtralongPeriodRate:     200000,
		signalFadingFastPeriodIncreaser:     300,
		signalFadingMediumPeriodIncreaser:   17500,
		signalFadingLongPeriodIncreaser:     60000,
		signalFadingExtralongPerioIncreaser: 150000,

		echoLevelMKV:                       0.3,
		echoFadingFastPeriodRate:           2000,
		echoFadingMediumPeriodRate:         7000,
		echoFadingLongPeriodRate:           50000,
		echoFadingExtralongPeriodRate:      200000,
		echoFadingFastPeriodIncreaser:      1500,
		echoFadingMediumPeriodIncreaser:    4000,
		echoFadingLongPeriodIncreaser:      25000,
		echoFadingExtralongPeriodIncreaser: 150000,

		bgSignalLevelMKV:                       0.2,
		bgSignalFadingFastPeriodRate:           2000,
		bgSignalFadingMediumPeriodRate:         7000,
		bgSignalFadingLongPeriodRate:           50000,
		bgSignalFadingExtralongPeriodRate:      200000,
		bgSignalFadingFastPeriodIncreaser:      1500,
		bgSignalFadingMediumPeriodIncreaser:    4000,
		bgSignalFadingLongPeriodIncreaser:      25000,
		bgSignalFadingExtralongPeriodIncreaser: 150000,

		interferenceNoiseLevelMKV:                       0.02,
		interferenceNoiseFadingFastPeriodRate:           2000,
		interferenceNoiseFadingMediumPeriodRate:         7000,
		interferenceNoiseFadingLongPeriodRate:           50000,
		interferenceNoiseFadingExtralongPeriodRate:      200000,
		interferenceNoiseFadingFastPeriodIncreaser:      1500,
		interferenceNoiseFadingMediumPeriodIncreaser:    4000,
		interferenceNoiseFadingLongPeriodIncreaser:      25000,
		interferenceNoiseFadingExtralongPeriodIncreaser: 150000,

		interferenceNoiseFreqFactorRate:       1000,
		interferenceNoiseFreqFactorIncreaser:  1.0,
		interferenceNoiseFreqFactorAdditional: 3000.0,

		noiseLevelMKV: 0.1,
	}
	return proc
}

//ProcessPipes
func (proc *Processor) ProcessPipes(mainPipe io.Reader, bgPipe io.Reader, freq int64, outChan chan byte, errChan chan error) {

	// Randomize from frequency value
	rand.Seed(freq)

	// Signal component levels
	// Payload signal

	signalFadingFastPeriod := math.Mod(rand.Float64(), proc.SignalFadingFastPeriodRate()) + proc.SignalFadingFastPeriodIncreaser()
	signalFadingMediumPeriod := math.Mod(rand.Float64(), proc.SignalFadingMediumPeriodRate()) + proc.SignalFadingMediumPeriodIncreaser()
	signalFadingLongPeriod := math.Mod(rand.Float64(), proc.SignalFadingLongPeriodRate()) + proc.SignalFadingLongPeriodIncreaser()
	signalFadingExtralongPeriod := math.Mod(rand.Float64(), proc.SignalFadingExtralongPeriodRate()) + proc.SignalFadingExtralongPerioIncreaser()

	// Payload signal echo
	echoFadingFastPeriod := math.Mod(rand.Float64(), proc.EchoFadingFastPeriodRate()) + proc.EchoFadingFastPeriodIncreaser()
	echoFadingMediumPeriod := math.Mod(rand.Float64(), proc.EchoFadingMediumPeriodRate()) + proc.EchoFadingMediumPeriodIncreaser()
	echoFadingLongPeriod := math.Mod(rand.Float64(), proc.EchoFadingLongPeriodRate()) + proc.EchoFadingLongPeriodIncreaser()
	echoFadingExtralongPeriod := math.Mod(rand.Float64(), proc.EchoFadingExtralongPeriodRate()) + proc.EchoFadingExtralongPeriodIncreaser()

	// Background signal
	bgSignalFadingFastPeriod := math.Mod(rand.Float64(), proc.BgSignalFadingFastPeriodRate()) + proc.BgSignalFadingFastPeriodIncreaser()
	bgSignalFadingMediumPeriod := math.Mod(rand.Float64(), proc.BgSignalFadingMediumPeriodRate()) + proc.BgSignalFadingMediumPeriodIncreaser()
	bgSignalFadingLongPeriod := math.Mod(rand.Float64(), proc.BgSignalFadingLongPeriodRate()) + proc.BgSignalFadingLongPeriodIncreaser()
	bgSignalFadingExtralongPeriod := math.Mod(rand.Float64(), proc.BgSignalFadingExtralongPeriodRate()) + proc.BgSignalFadingExtralongPeriodIncreaser()

	// Interference noise
	interferenceNoiseFreqFactor := math.Mod(rand.Float64(), proc.InterferenceNoiseFreqFactorRate())/proc.InterferenceNoiseFreqFactorAdditional() + proc.InterferenceNoiseFreqFactorIncreaser()
	interferenceNoiseFadingFastPeriod := math.Mod(rand.Float64(), proc.InterferenceNoiseFadingFastPeriodRate()) + proc.InterferenceNoiseFadingFastPeriodIncreaser()
	interferenceNoiseFadingMediumPeriod := math.Mod(rand.Float64(), proc.InterferenceNoiseFadingMediumPeriodRate()) + proc.InterferenceNoiseFadingMediumPeriodIncreaser()
	interferenceNoiseFadingLongPeriod := math.Mod(rand.Float64(), proc.InterferenceNoiseFadingLongPeriodRate()) + proc.InterferenceNoiseFadingLongPeriodIncreaser()
	interferenceNoiseFadingExtralongPeriod := math.Mod(rand.Float64(), proc.InterferenceNoiseFadingExtralongPeriodRate()) + proc.InterferenceNoiseFadingExtralongPeriodIncreaser()

	// Signal component values
	var signal, echo, bgSignal, noise, noiseSrc, interference float64

	// Current levels
	var signalCurLevel, echoCurLevel, bgSignalCurLevel, noiseCurLevel, interferenceCurLevel, commonCurLevel, commonSignalSrc float64

	// Other vars
	var output, prevOutput float64

	// Time counter
	var timeCounter float64

	mainPP := bufio.NewReader(mainPipe)
	bgPP := bufio.NewReader(bgPipe)

	var echoBuffer [int(echoSizeSamples)]float64

	for {
		mainSignalRune, err := mainPP.ReadByte()
		if err != nil {
			errChan <- fmt.Errorf("main pipe ended")
			break
		}
		signal = float64(mainSignalRune)

		bgSignalByte, err := bgPP.ReadByte()
		if err != nil {
			errChan <- fmt.Errorf("background pipe ended")
			break
		}
		bgSignal = float64(bgSignalByte)

		echo = echoBuffer[int(math.Mod(timeCounter, echoSizeSamples))]

		interference = 100 + 100*math.Sin(timeCounter/interferenceNoiseFreqFactor)

		noiseSrc = math.Mod(rand.Float64(), 256)
		noise = 0.5*noise + 0.5*noiseSrc

		// Current levels
		signalCurLevel = levelWithFading(proc.SignalLevelMKV(), timeCounter, signalFadingFastPeriod, signalFadingMediumPeriod, signalFadingLongPeriod, signalFadingExtralongPeriod)
		echoCurLevel = levelWithFading(proc.EchoLevelMKV(), timeCounter, echoFadingFastPeriod, echoFadingMediumPeriod, echoFadingLongPeriod, echoFadingExtralongPeriod)
		bgSignalCurLevel = levelWithFading(proc.BgSignalLevelMKV(), timeCounter, bgSignalFadingFastPeriod, bgSignalFadingMediumPeriod, bgSignalFadingLongPeriod, bgSignalFadingExtralongPeriod)
		noiseCurLevel = proc.NoiseLevelMKV()
		interferenceCurLevel = levelWithFading(proc.InterferenceNoiseLevelMKV(), timeCounter, interferenceNoiseFadingFastPeriod, interferenceNoiseFadingMediumPeriod, interferenceNoiseFadingLongPeriod, interferenceNoiseFadingExtralongPeriod)

		// Common level
		commonCurLevel = signalCurLevel + echoCurLevel + bgSignalCurLevel + noiseCurLevel + interferenceCurLevel

		// Common signal
		commonSignalSrc = signal*signalCurLevel + echo*echoCurLevel + bgSignal*bgSignalCurLevel + noise*noiseCurLevel + interference*interferenceCurLevel

		// Output with AGC
		output = commonSignalSrc / commonCurLevel

		outputPrepared := 0.5*output + 0.5*prevOutput

		outChan <- byte(outputPrepared)

		echoBuffer[int(math.Mod(timeCounter, echoSizeSamples))] = signal
		prevOutput = output

		timeCounter++
	}
}

//| oggenc -Q -r -B 8 -R 8000 -C 1 - -o - | ezstream -c /radio/config/play-to-stream-vorbis-${RADIO_FREQ}am.xml &
