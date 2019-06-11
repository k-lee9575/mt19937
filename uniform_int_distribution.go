package mt19937

import(
	"log"
)

type uniform_int_distribution struct{
	m_min int64
	m_max int64
	m_eng *MT19937
}

func DistInt64(eng *MT19937, begin int64, end int64) *uniform_int_distribution{
	if (begin > end){
		log.Println("begin is not allowed to be greater than end!")
		return nil
	}

	dist := &uniform_int_distribution{
		m_min : begin,
		m_max : end,
		m_eng : eng,
	}
	return dist
}

func (dist *uniform_int_distribution) Int64() int64{
	var rng uint64

	if (dist.m_min >= 0){
		rng = uint64(dist.m_max) - uint64(dist.m_min)
	} else if (dist.m_max >= 0){
		rng = uint64(dist.m_max) + uint64(-(dist.m_min + 1)) + 1
	} else {
		rng = uint64(dist.m_max - dist.m_min)
	}

	if (rng == 0){
		return dist.m_min
	} else if(rng == ^uint64(0)){
		return int64(dist.m_eng.Random())
	}
	bucket_size := ^uint64(0) / (rng + 1)
	if (^uint64(0) % (rng + 1) == rng){
		bucket_size ++
	}
	result := dist.m_eng.Random() / bucket_size
	return int64(result) + dist.m_min


}