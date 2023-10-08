package response

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/lib/pq"
	"time"
)

type CardResponse struct {
	CardID         string         `json:"card_id"`
	Period         string         `json:"period"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	AnalysisResult string         `json:"analysis_result"`
	Tags           pq.StringArray `json:"tags"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
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
