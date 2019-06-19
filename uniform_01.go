package mt19937

type uniform_01 struct{
	m_eng *MT19937
}

func Dist01(eng *MT19937) *uniform_01{
	dist := &uniform_01{
		m_eng : eng,
	}
	return dist
}

func (dist *uniform_01)Float64() float64{
	return float64(dist.m_eng.Random()) / float64(^uint64(0))
}