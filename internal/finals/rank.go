package finals

type League string

const (
	Bronze   League = "Bronze"
	Silver   League = "Silver"
	Gold     League = "Gold"
	Platinum League = "Platinum"
	Diamond  League = "Diamond"
	Ruby     League = "Ruby"
)

type Rank struct {
	Bracket int    `json:"bracket"`
	League  League `json:"league"`
}
