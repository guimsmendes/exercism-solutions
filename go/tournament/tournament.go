package tournament

type Team struct {
	name    string
	matches int
	wins    int
	draws   int
	losses  int
	points  int
}

const head = "Team                           | MP |  W |  D |  L |  P"
