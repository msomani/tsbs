package main

import "time"

var (
	RandByteString      = []byte("rand")       // heap optimization
	RandTotalByteString = []byte("rand-total") // heap optimization
)

var (
	// Field keys for 'cpu' points.
	RandFieldKeys = [][]byte{
		[]byte("usage_user"),
		[]byte("usage_system"),
		[]byte("usage_idle"),
		[]byte("usage_nice"),
		[]byte("usage_iowait"),
		[]byte("usage_irq"),
		[]byte("usage_softirq"),
		[]byte("usage_steal"),
		[]byte("usage_guest"),
		[]byte("usage_guest_nice"),
	}
)

type RandMeasurement struct {
	timestamp     time.Time
	distributions []Distribution
}

func NewRandMeasurement(start time.Time) *RandMeasurement {
	distributions := make([]Distribution, len(RandFieldKeys))
	for i := range distributions {
		distributions[i] = &UniformDistribution{
			Low:  0.0,
			High: 100.0,
		}
	}
	return &RandMeasurement{
		timestamp:     start,
		distributions: distributions,
	}
}

func (m *RandMeasurement) Tick(d time.Duration) {
	m.timestamp = m.timestamp.Add(d)
	for i := range m.distributions {
		m.distributions[i].Advance()
	}
}

func (m *RandMeasurement) ToPoint(p *Point) {
	p.SetMeasurementName(RandByteString)
	p.SetTimestamp(&m.timestamp)

	for i := range m.distributions {
		p.AppendField(RandFieldKeys[i], m.distributions[i].Get())
	}
}