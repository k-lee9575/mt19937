package mt19937

import(
	"log"
)

type bernoulli_distribution struct{
	m_prob float64
	m_eng *MT19937
}

func DistBernolli(eng *MT19937, prob float64) *bernoulli_distribution{
	if(prob < 0 || prob > 1){
		log.Println("ERROR! prob must be between 0 and 1")
		return nil
	}
	dist := &bernoulli_distribution{
		m_prob : prob,
		m_eng : eng,
	}
	return dist
}

func (dist *bernoulli_distribution) Bool() bool{
	if (dist.m_prob == 0){
		return false
	}else{
		return float64(dist.m_eng.Random()) <= dist.m_prob * float64(^uint64(0))
	}
}