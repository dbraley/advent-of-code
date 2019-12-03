package main

import "fmt"

// TODO: read this from the file or something less Q&D
var moduleMasses = []int {
73910,
57084,
102852,
134452,
108006,
134228,
102765,
60642,
64819,
54335,
82480,
135119,
73027,
107087,
108254,
111944,
83790,
128585,
52889,
53870,
145120,
96863,
57105,
97702,
75324,
70566,
120914,
95808,
86568,
143498,
125093,
71370,
122889,
67808,
133643,
52806,
103532,
126487,
54807,
121402,
57580,
75759,
84225,
102232,
112367,
95635,
132871,
102903,
51997,
74565,
63674,
97410,
96965,
55711,
53547,
117482,
107957,
108175,
136622,
144235,
80407,
78670,
114870,
145967,
148646,
75955,
84293,
129590,
144067,
142192,
79117,
123861,
68546,
148675,
88932,
91493,
127808,
96517,
130687,
137822,
77726,
110502,
130509,
98370,
136008,
142920,
81358,
112950,
101057,
86547,
128714,
135401,
55903,
66606,
105404,
55276,
57427,
101556,
91111,
79585,
}

func main() {
	moduleFuel := TransformAndSum(moduleMasses, Fuel)
	moduleAndFuelFuel := TransformAndSum(moduleMasses, Fuel2)
	fmt.Printf("Total Fuel needed for %v modules: %v\n", len(moduleMasses), moduleFuel)
	fmt.Printf("Total Fuel needed for %v modules & Fuel: %v\n", len(moduleMasses), moduleAndFuelFuel)
}

func TransformAndSum(numbers []int, t func(int)int)  int {
	sum := 0
	for _, val := range numbers {
		sum += t(val)
	}
	return sum
}

func Fuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func Fuel2(mass int) int {
	newFuel := Fuel(mass)
	if newFuel > 0 {
		return newFuel + Fuel2(newFuel)
	}
	return newFuel
}
