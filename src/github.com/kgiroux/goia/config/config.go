package config

// Config that will contains configuration;
type Config struct {
	log           bool
	learningRate  float64
	epochNumber   int
	entriesNumber int
	dataSetNumber int
}

// Config return the config
func GetConfig() Config {
	return Config{log: true, learningRate: 0.08, epochNumber: 100000, entriesNumber: 2, dataSetNumber: 6}
}

// setter
func (m *Config) setLog(value bool) {
	m.log = value
}

// Log return the value for the log
func (m *Config) Log() bool {
	return m.log
}

// setter
func (m *Config) setLearningRate(value float64) {
	m.learningRate = value
}

// LearningRate return the value for the learning rate
func (m *Config) LearningRate() float64 {
	return m.learningRate
}

// setter
func (m *Config) setEpochNumber(value int) {
	m.epochNumber = value
}

// EpochNumber return the value for the learning rate
func (m *Config) EpochNumber() int {
	return m.epochNumber
}

// SetEntriesNumber define the entries number
func (m *Config) SetEntriesNumber(value int) {
	m.entriesNumber = value
}

// GetEntriesNumber return the value for the entries Number
func (m *Config) GetEntriesNumber() int {
	return m.entriesNumber
}

// SetDataSetNumber define the DataSet number
func (m *Config) SetDataSetNumber(value int) {
	m.dataSetNumber = value
}

// GetDataSetNumber return the value for the DataSetNumber
func (m *Config) GetDataSetNumber() int {
	return m.dataSetNumber
}
