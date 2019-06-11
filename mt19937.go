package mt19937


//---------------------------------Public------------------------------------------

type MT19937 struct{
	state []uint64
	index int
}

func New() *MT19937{
	mt := &MT19937{
		state: make([]uint64, n),
		index: n,
	}
	mt.Seed(5489)
	return mt
}

func (mt *MT19937) Seed(seed uint64){
	x := mt.state
	x[0] = seed
	for i := 1; i < n; i++{
		x[i] = f * (x[i-1] ^ (x[i-1] >> (w-2))) + uint64(i)
	}
}

func (mt *MT19937)  Random() uint64 {
	x:=mt.state
	if (mt.index == n){
		mt.twist()
	}
	var z uint64 = x[mt.index]

	mt.index++

 	z ^= ((z >> u) & d)
 	z ^= ((z << s) & b)
 	z ^= ((z << t) & c)
    z ^= (z >> l)
    return z
}

//---------------------------------Private------------------------------------------

const(
	//算法中用到的变量如下所示：
    //w：长度（以bit为单位）
    //n：递归长度
    //m：周期参数，用作第三阶段的偏移量
    //r：低位掩码/低位要提取的位数
    //a：旋转矩阵的参数
    //b,c：TGFSR的掩码
    //s,t：TGFSR的位移量
    //u,d,l：额外梅森旋转所需的掩码和位移量
    //f：初始化梅森旋转链所需参数
    
	w = 64
	n = 312
	m = 156
	r = 31

	a = 0xb5026f5aa96619e9

	u = 29
	d = 0x5555555555555555

	s = 17
	b = 0x71d67fffeda60000

	t = 37
	c = 0xfff7eee000000000

	l = 43

	f = 6364136223846793005
)

func (mt *MT19937) twist(){
	x := mt.state
	const lower_mask uint64 = 1 << r -1
	const upper_mask uint64 = ^lower_mask

	for j := 0; j < n - m; j ++ {
		var y uint64 =(x[j] & upper_mask) | (x[j + 1] & lower_mask)
		x[j] = x[j + m] ^ (y >> 1) ^ ((x[j + 1] & 1) * a)
	}

	for j := n - m; j < n - 1; j ++ {
		var y uint64 =(x[j] & upper_mask) | (x[j + 1] & lower_mask)
		x[j] = x[j - (n - m)] ^ (y >> 1) ^ ((x[j + 1] & 1) * a)
	}


	var y uint64 = (x[n-1] & upper_mask) | (x[0] & lower_mask)
	x[n - 1] = x[m - 1] ^ (y >> 1) ^ ((x[0] & 1) * a)
	mt.index = 0
}