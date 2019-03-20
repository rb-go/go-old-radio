package oldradio

// SIGNAL PART

//SignalLevelMKV ...
func (proc *Processor) SignalLevelMKV() float64 {
	return proc.signalLevelMKV
}

//SetSignalLevelMKV ...
func (proc *Processor) SetSignalLevelMKV(signalLevelMKV float64) {
	proc.Lock()
	proc.signalLevelMKV = signalLevelMKV
	proc.Unlock()
}

//SignalFadingFastPeriodRate ...
func (proc *Processor) SignalFadingFastPeriodRate() float64 {
	return proc.signalFadingFastPeriodRate
}

//SetSignalFadingFastPeriodRate ...
func (proc *Processor) SetSignalFadingFastPeriodRate(signalFadingFastPeriodRate float64) {
	proc.Lock()
	proc.signalFadingFastPeriodRate = signalFadingFastPeriodRate
	proc.Unlock()
}

//SignalFadingMediumPeriodRate ...
func (proc *Processor) SignalFadingMediumPeriodRate() float64 {
	return proc.signalFadingMediumPeriodRate
}

//SetSignalFadingMediumPeriodRate ...
func (proc *Processor) SetSignalFadingMediumPeriodRate(signalFadingMediumPeriodRate float64) {
	proc.Lock()
	proc.signalFadingMediumPeriodRate = signalFadingMediumPeriodRate
	proc.Unlock()
}

//SignalFadingLongPeriodRate ...
func (proc *Processor) SignalFadingLongPeriodRate() float64 {
	return proc.signalFadingLongPeriodRate
}

//SetSignalFadingLongPeriodRate ...
func (proc *Processor) SetSignalFadingLongPeriodRate(signalFadingLongPeriodRate float64) {
	proc.Lock()
	proc.signalFadingLongPeriodRate = signalFadingLongPeriodRate
	proc.Unlock()
}

//SignalFadingExtralongPeriodRate ...
func (proc *Processor) SignalFadingExtralongPeriodRate() float64 {
	return proc.signalFadingExtralongPeriodRate
}

//SetSignalFadingExtralongPeriodRate ...
func (proc *Processor) SetSignalFadingExtralongPeriodRate(signalFadingExtralongPeriodRate float64) {
	proc.Lock()
	proc.signalFadingExtralongPeriodRate = signalFadingExtralongPeriodRate
	proc.Unlock()
}

//SignalFadingFastPeriodIncreaser ...
func (proc *Processor) SignalFadingFastPeriodIncreaser() float64 {
	return proc.signalFadingFastPeriodIncreaser
}

//SetSignalFadingFastPeriodIncreaser ...
func (proc *Processor) SetSignalFadingFastPeriodIncreaser(signalFadingFastPeriodIncreaser float64) {
	proc.Lock()
	proc.signalFadingFastPeriodIncreaser = signalFadingFastPeriodIncreaser
	proc.Unlock()
}

//SignalFadingMediumPeriodIncreaser ...
func (proc *Processor) SignalFadingMediumPeriodIncreaser() float64 {
	return proc.signalFadingMediumPeriodIncreaser
}

//SetSignalFadingMediumPeriodIncreaser ...
func (proc *Processor) SetSignalFadingMediumPeriodIncreaser(signalFadingMediumPeriodIncreaser float64) {
	proc.Lock()
	proc.signalFadingMediumPeriodIncreaser = signalFadingMediumPeriodIncreaser
	proc.Unlock()
}

//SignalFadingLongPeriodIncreaser ...
func (proc *Processor) SignalFadingLongPeriodIncreaser() float64 {
	return proc.signalFadingLongPeriodIncreaser
}

//SetSignalFadingLongPeriodIncreaser ...
func (proc *Processor) SetSignalFadingLongPeriodIncreaser(signalFadingLongPeriodIncreaser float64) {
	proc.Lock()
	proc.signalFadingLongPeriodIncreaser = signalFadingLongPeriodIncreaser
	proc.Unlock()
}

//SignalFadingExtralongPerioIncreaser ...
func (proc *Processor) SignalFadingExtralongPerioIncreaser() float64 {
	return proc.signalFadingExtralongPerioIncreaser
}

//SetSignalFadingExtralongPerioIncreaser ...
func (proc *Processor) SetSignalFadingExtralongPerioIncreaser(signalFadingExtralongPerioIncreaser float64) {
	proc.Lock()
	proc.signalFadingExtralongPerioIncreaser = signalFadingExtralongPerioIncreaser
	proc.Unlock()
}

// ECHO PART

//EchoLevelMKV ...
func (proc *Processor) EchoLevelMKV() float64 {
	return proc.echoLevelMKV
}

//SetEchoLevelMKV ...
func (proc *Processor) SetEchoLevelMKV(echoLevelMKV float64) {
	proc.Lock()
	proc.echoLevelMKV = echoLevelMKV
	proc.Unlock()
}

//EchoFadingFastPeriodRate  ...
func (proc *Processor) EchoFadingFastPeriodRate() float64 {
	return proc.echoFadingFastPeriodRate
}

//SetEchoFadingFastPeriodRate  ...
func (proc *Processor) SetEchoFadingFastPeriodRate(echoFadingFastPeriodRate float64) {
	proc.Lock()
	proc.echoFadingFastPeriodRate = echoFadingFastPeriodRate
	proc.Unlock()
}

//EchoFadingMediumPeriodRate  ...
func (proc *Processor) EchoFadingMediumPeriodRate() float64 {
	return proc.echoFadingMediumPeriodRate
}

//SetEchoFadingMediumPeriodRate  ...
func (proc *Processor) SetEchoFadingMediumPeriodRate(echoFadingMediumPeriodRate float64) {
	proc.Lock()
	proc.echoFadingMediumPeriodRate = echoFadingMediumPeriodRate
	proc.Unlock()
}

//EchoFadingLongPeriodRate  ...
func (proc *Processor) EchoFadingLongPeriodRate() float64 {
	return proc.echoFadingLongPeriodRate
}

//SetEchoFadingLongPeriodRate  ...
func (proc *Processor) SetEchoFadingLongPeriodRate(echoFadingLongPeriodRate float64) {
	proc.Lock()
	proc.echoFadingLongPeriodRate = echoFadingLongPeriodRate
	proc.Unlock()
}

//EchoFadingExtralongPeriodRate  ...
func (proc *Processor) EchoFadingExtralongPeriodRate() float64 {
	return proc.echoFadingExtralongPeriodRate
}

//SetEchoFadingExtralongPeriodRate  ...
func (proc *Processor) SetEchoFadingExtralongPeriodRate(echoFadingExtralongPeriodRate float64) {
	proc.Lock()
	proc.echoFadingExtralongPeriodRate = echoFadingExtralongPeriodRate
	proc.Unlock()
}

//EchoFadingFastPeriodIncreaser  ...
func (proc *Processor) EchoFadingFastPeriodIncreaser() float64 {
	return proc.echoFadingFastPeriodIncreaser
}

//SetEchoFadingFastPeriodIncreaser  ...
func (proc *Processor) SetEchoFadingFastPeriodIncreaser(echoFadingFastPeriodIncreaser float64) {
	proc.Lock()
	proc.echoFadingFastPeriodIncreaser = echoFadingFastPeriodIncreaser
	proc.Unlock()
}

//EchoFadingMediumPeriodIncreaser  ...
func (proc *Processor) EchoFadingMediumPeriodIncreaser() float64 {
	return proc.echoFadingMediumPeriodIncreaser
}

//SetEchoFadingMediumPeriodIncreaser  ...
func (proc *Processor) SetEchoFadingMediumPeriodIncreaser(echoFadingMediumPeriodIncreaser float64) {
	proc.Lock()
	proc.echoFadingMediumPeriodIncreaser = echoFadingMediumPeriodIncreaser
	proc.Unlock()
}

//EchoFadingLongPeriodIncreaser  ...
func (proc *Processor) EchoFadingLongPeriodIncreaser() float64 {
	return proc.echoFadingLongPeriodIncreaser
}

//SetEchoFadingLongPeriodIncreaser  ...
func (proc *Processor) SetEchoFadingLongPeriodIncreaser(echoFadingLongPeriodIncreaser float64) {
	proc.Lock()
	proc.echoFadingLongPeriodIncreaser = echoFadingLongPeriodIncreaser
	proc.Unlock()
}

//EchoFadingExtralongPeriodIncreaser  ...
func (proc *Processor) EchoFadingExtralongPeriodIncreaser() float64 {
	return proc.echoFadingExtralongPeriodIncreaser
}

//SetEchoFadingExtralongPeriodIncreaser  ...
func (proc *Processor) SetEchoFadingExtralongPeriodIncreaser(echoFadingExtralongPeriodIncreaser float64) {
	proc.Lock()
	proc.echoFadingExtralongPeriodIncreaser = echoFadingExtralongPeriodIncreaser
	proc.Unlock()
}

// BG SIGNAL PART

//BgSignalLevelMKV  ...
func (proc *Processor) BgSignalLevelMKV() float64 {
	return proc.bgSignalLevelMKV
}

//SetBgSignalLevelMKV  ...
func (proc *Processor) SetBgSignalLevelMKV(bgSignalLevelMKV float64) {
	proc.Lock()
	proc.bgSignalLevelMKV = bgSignalLevelMKV
	proc.Unlock()
}

//BgSignalFadingFastPeriodRate  ...
func (proc *Processor) BgSignalFadingFastPeriodRate() float64 {
	return proc.bgSignalFadingFastPeriodRate
}

//SetBgSignalFadingFastPeriodRate  ...
func (proc *Processor) SetBgSignalFadingFastPeriodRate(bgSignalFadingFastPeriodRate float64) {
	proc.Lock()
	proc.bgSignalFadingFastPeriodRate = bgSignalFadingFastPeriodRate
	proc.Unlock()
}

//BgSignalFadingMediumPeriodRate  ...
func (proc *Processor) BgSignalFadingMediumPeriodRate() float64 {
	return proc.bgSignalFadingMediumPeriodRate
}

//SetBgSignalFadingMediumPeriodRate  ...
func (proc *Processor) SetBgSignalFadingMediumPeriodRate(bgSignalFadingMediumPeriodRate float64) {
	proc.Lock()
	proc.bgSignalFadingMediumPeriodRate = bgSignalFadingMediumPeriodRate
	proc.Unlock()
}

//BgSignalFadingLongPeriodRate  ...
func (proc *Processor) BgSignalFadingLongPeriodRate() float64 {
	return proc.bgSignalFadingLongPeriodRate
}

//SetBgSignalFadingLongPeriodRate  ...
func (proc *Processor) SetBgSignalFadingLongPeriodRate(bgSignalFadingLongPeriodRate float64) {
	proc.Lock()
	proc.bgSignalFadingLongPeriodRate = bgSignalFadingLongPeriodRate
	proc.Unlock()
}

//BgSignalFadingExtralongPeriodRate  ...
func (proc *Processor) BgSignalFadingExtralongPeriodRate() float64 {
	return proc.bgSignalFadingExtralongPeriodRate
}

//SetBgSignalFadingExtralongPeriodRate  ...
func (proc *Processor) SetBgSignalFadingExtralongPeriodRate(bgSignalFadingExtralongPeriodRate float64) {
	proc.Lock()
	proc.bgSignalFadingExtralongPeriodRate = bgSignalFadingExtralongPeriodRate
	proc.Unlock()
}

//BgSignalFadingFastPeriodIncreaser  ...
func (proc *Processor) BgSignalFadingFastPeriodIncreaser() float64 {
	return proc.bgSignalFadingFastPeriodIncreaser
}

//SetBgSignalFadingFastPeriodIncreaser  ...
func (proc *Processor) SetBgSignalFadingFastPeriodIncreaser(bgSignalFadingFastPeriodIncreaser float64) {
	proc.Lock()
	proc.bgSignalFadingFastPeriodIncreaser = bgSignalFadingFastPeriodIncreaser
	proc.Unlock()
}

//BgSignalFadingMediumPeriodIncreaser  ...
func (proc *Processor) BgSignalFadingMediumPeriodIncreaser() float64 {
	return proc.bgSignalFadingMediumPeriodIncreaser
}

//SetBgSignalFadingMediumPeriodIncreaser  ...
func (proc *Processor) SetBgSignalFadingMediumPeriodIncreaser(bgSignalFadingMediumPeriodIncreaser float64) {
	proc.Lock()
	proc.bgSignalFadingMediumPeriodIncreaser = bgSignalFadingMediumPeriodIncreaser
	proc.Unlock()
}

//BgSignalFadingLongPeriodIncreaser  ...
func (proc *Processor) BgSignalFadingLongPeriodIncreaser() float64 {
	return proc.bgSignalFadingLongPeriodIncreaser
}

//SetBgSignalFadingLongPeriodIncreaser  ...
func (proc *Processor) SetBgSignalFadingLongPeriodIncreaser(bgSignalFadingLongPeriodIncreaser float64) {
	proc.Lock()
	proc.bgSignalFadingLongPeriodIncreaser = bgSignalFadingLongPeriodIncreaser
	proc.Unlock()
}

//BgSignalFadingExtralongPeriodIncreaser  ...
func (proc *Processor) BgSignalFadingExtralongPeriodIncreaser() float64 {
	return proc.bgSignalFadingExtralongPeriodIncreaser
}

//SetBgSignalFadingExtralongPeriodIncreaser  ...
func (proc *Processor) SetBgSignalFadingExtralongPeriodIncreaser(bgSignalFadingExtralongPeriodIncreaser float64) {
	proc.Lock()
	proc.bgSignalFadingExtralongPeriodIncreaser = bgSignalFadingExtralongPeriodIncreaser
	proc.Unlock()
}

// INTERFERENCE NOISE PART

//InterferenceNoiseLevelMKV  ...
func (proc *Processor) InterferenceNoiseLevelMKV() float64 {
	return proc.interferenceNoiseLevelMKV
}

//SetInterferenceNoiseLevelMKV  ...
func (proc *Processor) SetInterferenceNoiseLevelMKV(interferenceNoiseLevelMKV float64) {
	proc.Lock()
	proc.interferenceNoiseLevelMKV = interferenceNoiseLevelMKV
	proc.Unlock()
}

//InterferenceNoiseFadingFastPeriodRate  ...
func (proc *Processor) InterferenceNoiseFadingFastPeriodRate() float64 {
	return proc.interferenceNoiseFadingFastPeriodRate
}

//SetInterferenceNoiseFadingFastPeriodRate  ...
func (proc *Processor) SetInterferenceNoiseFadingFastPeriodRate(interferenceNoiseFadingFastPeriodRate float64) {
	proc.Lock()
	proc.interferenceNoiseFadingFastPeriodRate = interferenceNoiseFadingFastPeriodRate
	proc.Unlock()
}

//InterferenceNoiseFadingMediumPeriodRate  ...
func (proc *Processor) InterferenceNoiseFadingMediumPeriodRate() float64 {
	return proc.interferenceNoiseFadingMediumPeriodRate
}

//SetInterferenceNoiseFadingMediumPeriodRate  ...
func (proc *Processor) SetInterferenceNoiseFadingMediumPeriodRate(interferenceNoiseFadingMediumPeriodRate float64) {
	proc.Lock()
	proc.interferenceNoiseFadingMediumPeriodRate = interferenceNoiseFadingMediumPeriodRate
	proc.Unlock()
}

//InterferenceNoiseFadingLongPeriodRate  ...
func (proc *Processor) InterferenceNoiseFadingLongPeriodRate() float64 {
	return proc.interferenceNoiseFadingLongPeriodRate
}

//SetInterferenceNoiseFadingLongPeriodRate  ...
func (proc *Processor) SetInterferenceNoiseFadingLongPeriodRate(interferenceNoiseFadingLongPeriodRate float64) {
	proc.Lock()
	proc.interferenceNoiseFadingLongPeriodRate = interferenceNoiseFadingLongPeriodRate
	proc.Unlock()
}

//InterferenceNoiseFadingExtralongPeriodRate  ...
func (proc *Processor) InterferenceNoiseFadingExtralongPeriodRate() float64 {
	return proc.interferenceNoiseFadingExtralongPeriodRate
}

//SetInterferenceNoiseFadingExtralongPeriodRate  ...
func (proc *Processor) SetInterferenceNoiseFadingExtralongPeriodRate(interferenceNoiseFadingExtralongPeriodRate float64) {
	proc.Lock()
	proc.interferenceNoiseFadingExtralongPeriodRate = interferenceNoiseFadingExtralongPeriodRate
	proc.Unlock()
}

//InterferenceNoiseFadingFastPeriodIncreaser  ...
func (proc *Processor) InterferenceNoiseFadingFastPeriodIncreaser() float64 {
	return proc.interferenceNoiseFadingFastPeriodIncreaser
}

//SetInterferenceNoiseFadingFastPeriodIncreaser  ...
func (proc *Processor) SetInterferenceNoiseFadingFastPeriodIncreaser(interferenceNoiseFadingFastPeriodIncreaser float64) {
	proc.Lock()
	proc.interferenceNoiseFadingFastPeriodIncreaser = interferenceNoiseFadingFastPeriodIncreaser
	proc.Unlock()
}

//InterferenceNoiseFadingMediumPeriodIncreaser  ...
func (proc *Processor) InterferenceNoiseFadingMediumPeriodIncreaser() float64 {
	return proc.interferenceNoiseFadingMediumPeriodIncreaser
}

//SetInterferenceNoiseFadingMediumPeriodIncreaser  ...
func (proc *Processor) SetInterferenceNoiseFadingMediumPeriodIncreaser(interferenceNoiseFadingMediumPeriodIncreaser float64) {
	proc.Lock()
	proc.interferenceNoiseFadingMediumPeriodIncreaser = interferenceNoiseFadingMediumPeriodIncreaser
	proc.Unlock()
}

//InterferenceNoiseFadingLongPeriodIncreaser  ...
func (proc *Processor) InterferenceNoiseFadingLongPeriodIncreaser() float64 {
	return proc.interferenceNoiseFadingLongPeriodIncreaser
}

//SetInterferenceNoiseFadingLongPeriodIncreaser  ...
func (proc *Processor) SetInterferenceNoiseFadingLongPeriodIncreaser(interferenceNoiseFadingLongPeriodIncreaser float64) {
	proc.Lock()
	proc.interferenceNoiseFadingLongPeriodIncreaser = interferenceNoiseFadingLongPeriodIncreaser
	proc.Unlock()
}

//InterferenceNoiseFadingExtralongPeriodIncreaser  ...
func (proc *Processor) InterferenceNoiseFadingExtralongPeriodIncreaser() float64 {
	return proc.interferenceNoiseFadingExtralongPeriodIncreaser
}

//SetInterferenceNoiseFadingExtralongPeriodIncreaser  ...
func (proc *Processor) SetInterferenceNoiseFadingExtralongPeriodIncreaser(interferenceNoiseFadingExtralongPeriodIncreaser float64) {
	proc.Lock()
	proc.interferenceNoiseFadingExtralongPeriodIncreaser = interferenceNoiseFadingExtralongPeriodIncreaser
	proc.Unlock()
}

//InterferenceNoiseFreqFactorRate  ...
func (proc *Processor) InterferenceNoiseFreqFactorRate() float64 {
	return proc.interferenceNoiseFreqFactorRate
}

//SetInterferenceNoiseFreqFactorRate  ...
func (proc *Processor) SetInterferenceNoiseFreqFactorRate(interferenceNoiseFreqFactorRate float64) {
	proc.Lock()
	proc.interferenceNoiseFreqFactorRate = interferenceNoiseFreqFactorRate
	proc.Unlock()
}

//InterferenceNoiseFreqFactorIncreaser  ...
func (proc *Processor) InterferenceNoiseFreqFactorIncreaser() float64 {
	return proc.interferenceNoiseFreqFactorIncreaser
}

//SetInterferenceNoiseFreqFactorIncreaser  ...
func (proc *Processor) SetInterferenceNoiseFreqFactorIncreaser(interferenceNoiseFreqFactorIncreaser float64) {
	proc.Lock()
	proc.interferenceNoiseFreqFactorIncreaser = interferenceNoiseFreqFactorIncreaser
	proc.Unlock()
}

//InterferenceNoiseFreqFactorAdditional  ...
func (proc *Processor) InterferenceNoiseFreqFactorAdditional() float64 {
	return proc.interferenceNoiseFreqFactorAdditional
}

//SetInterferenceNoiseFreqFactorAdditional  ...
func (proc *Processor) SetInterferenceNoiseFreqFactorAdditional(interferenceNoiseFreqFactorAdditional float64) {
	proc.Lock()
	proc.interferenceNoiseFreqFactorAdditional = interferenceNoiseFreqFactorAdditional
	proc.Unlock()
}

// NOISE PART

//NoiseLevelMKV  ...
func (proc *Processor) NoiseLevelMKV() float64 {
	return proc.noiseLevelMKV
}

//SetNoiseLevelMKV  ...
func (proc *Processor) SetNoiseLevelMKV(noiseLevelMKV float64) {
	proc.Lock()
	proc.noiseLevelMKV = noiseLevelMKV
	proc.Unlock()
}
