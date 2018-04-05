package neurone

// DataSet that will contains data;
type DataSet struct {
	entries []float64
	target  float64
}

// GetDataSet : Constructor a DataSet Object
func GetDataSet() DataSet {
	return DataSet{}
}

// GetDataSetArray : Constructor a DataSet Object
func GetDataSetArray(size int) []DataSet {
	return make([]DataSet, size)
}

// GetEntries : return entries
func (m *DataSet) GetEntries() []float64 {
	return m.entries
}

// GetTarget : return target of the dataset
func (m *DataSet) GetTarget() float64 {
	return m.target
}

// SetEntries : Definition des entries
func (m *DataSet) SetEntries(value []float64) {
	m.entries = value
}

// SetTarget : Definition des targets
func (m *DataSet) SetTarget(value float64) {
	m.target = value
}
