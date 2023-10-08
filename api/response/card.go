package response

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/lib/pq"
	"time"
)

type CardResponse struct {
	CardID         string
	Period         string
	Title          string
	Content        string
	AnalysisResult string
	Tags           pq.StringArray
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func ToCardResponse(card *entity.Card) CardResponse {
	cardResponse := CardResponse{
		CardID:         card.CardID,
		Period:         card.Period,
		Title:          card.Title,
		Content:        card.Content,
		AnalysisResult: card.AnalysisResult,
		Tags:           card.Tags,
		CreatedAt:      card.CreatedAt,
		UpdatedAt:      card.UpdatedAt,
	}
	return cardResponse
}

func ToCardsResponse(cards *[]entity.Card) []CardResponse {
	cardsResponse := []CardResponse{}
	for _, card := range *cards {
		cardResponse := ToCardResponse(&card)
		cardsResponse = append(cardsResponse, cardResponse)
	}
	return cardsResponse
}
