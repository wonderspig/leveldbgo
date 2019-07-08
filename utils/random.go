package utils

 

const (
	M = 2147483647 //2^31-1
	A = 16807   // bits 14, 8, 7, 5, 2, 1, 0

)

type Random struct{
	seed uint32
}


func NewRandom(sed uint32)*Random{
	seedLocal := sed & M 
	if seedLocal ==0 || seedLocal == M {
		seedLocal=1
	}
	return &Random{
		seed: seedLocal,
	}
}

func (random* Random)Next()uint32{
	var product uint64 =(uint64)(random.seed * A)
	random.seed=(uint32)((product>>31)+(product&M))
	if random.seed > M{
		random.seed-=M
	}
	return random.seed
}

func (random*Random)Uniform(n uint32)uint32{
	return (uint32)(random.Next() % n)
}

func (random*Random)OneIn(n uint32)bool{
	return (random.Next() %n ==0 )
}


func (random*Random)Skewed(maxLog int)uint32{
	return random.Uniform(1<<random.Uniform(uint32(maxLog+1)))
}