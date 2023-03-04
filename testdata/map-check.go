package sth

func foo() {

	type info struct {
		A int
		B string
	}

	tmpMap := make(map[int]*info)
	tmpMap[1] = &info{
		A: 11,
		B: "11",
	}
	tmpMap[2] = &info{
		A: 22,
		B: "2",
	}
	tmpMap[3] = &info{
		A: 33,
		B: "33",
	}

	t := tmpMap[4].B
	println(t)
}
