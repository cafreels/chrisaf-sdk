package lotr_sdk

type Movie struct {
	ID                         string  `json:"_id"`
	Name                       string  `json:"name"`
	RuntimeInMinutes           int     `json:"runtimeInMinutes"`
	BudgetInMillions           float64 `json:"budgetInMillions"`
	BoxOfficeRevenueInMillions float64 `json:"boxOfficeRevenueInMillions"`
	AcademyAwardNominations    int     `json:"academyAwardNominations"`
	AcademyAwardWins           int     `json:"academyAwardWins"`
	RottenTomatoesScore        float64 `json:"rottenTomatoesScore"`
}

type Quote struct {
	ID        string `json:"_id"`
	Dialog    string `json:"dialog"`
	Movie     string `json:"movie"`
	Character string `json:"character"`
}

func (c *LOTRClient) GetMovies() ([]Movie, PaginationStats, error) {
	result := &genericResponse[Movie]{}
	_, err := c.R().SetResult(result).Get("/movie")
	if err != nil {
		return nil, PaginationStats{}, err
	}
	return result.Docs, result.PaginationStats, err
}
