package args

import "flag"

func ParseArgs() (int, string) {
	var tLimit int
	var csvLoc string
	flag.IntVar(
		&tLimit,
		"limit",
		30,
		"the time limit for the quiz in seconds ",
	)
	flag.StringVar(
		&csvLoc,
		"csv",
		"problems.csv",
		"a csv file in the format 'question,answer'",
	)
	flag.Parse()

	return tLimit, csvLoc
}
