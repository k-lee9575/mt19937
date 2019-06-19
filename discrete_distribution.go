package mt19937

import(
	"log"
)

type aliasTableUnit struct{
	weight float64
	index int
}

type aliasTable struct{
	m_aliasTable []aliasTableUnit
	m_eng *MT19937
}

func DistDiscrete(eng *MT19937, list []int) *aliasTable{
	if(len(list) == 0){
		log.Println("Error! list is not allowed empty!")
		return nil
	}
	var weightSum int
	for _, weight := range list {
		weightSum += weight
	}
	average := float64(weightSum) / float64(len(list))

	table := &aliasTable{
		m_aliasTable : make([]aliasTableUnit, len(list), len(list)),
		m_eng : eng,
	}

	below_average := make([]aliasTableUnit, 0, len(list))
	above_average := make([]aliasTableUnit, 0, len(list))

	for i, weight := range list {
		val := float64(weight) / average
		unit := aliasTableUnit{
			weight : val,
			index : i,
		}
		if (val < 1){
			below_average = append(below_average, unit)
		}else{
			above_average = append(above_average, unit)
		}
	}

	posA := 0
	posB := 0
	for posB < len(below_average) && posA < len(above_average) {
		table.m_aliasTable[below_average[posB].index] = aliasTableUnit{
			weight : below_average[posB].weight,
			index : above_average[posA].index,
		}
		above_average[posA].weight -= (1 - below_average[posB].weight)
		if (above_average[posA].weight < 1){
			below_average[posB] = above_average[posA]
			posA ++
		} else{
			posB ++
		}
	}

	for ;posB < len(below_average); posB ++{
		table.m_aliasTable[below_average[posB].index].weight = float64(1)
	}
	for ;posA < len(above_average); posA ++{
		table.m_aliasTable[above_average[posA].index].weight = float64(1)
	}

	return table
}

func (dist *aliasTable)Discrete() int{
	result := int(DistInt64(dist.m_eng, 0, int64(len(dist.m_aliasTable) - 1)).Int64())
	test := Dist01(dist.m_eng).Float64()
	if(test < dist.m_aliasTable[result].weight){
		return result
	} else {
		return dist.m_aliasTable[result].index
	}
}