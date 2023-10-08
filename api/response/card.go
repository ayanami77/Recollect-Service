package response

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/lib/pq"
	"time"
)

type CardResponse struct {
	CardID         string
	UserID         string
	Period         string
	Title          string
	Content        string
	AnalysisResult string
	Tags           pq.StringArray
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func ToCardResponse(card *entity.Card) CardResponse {
	cardResponse := CardResponse{
		CardID:         card.CardID,
		UserID:         card.UserID,
		Period:         card.Period,
		Title:          card.Title,
		Content:        card.Content,
		AnalysisResult: card.AnalysisResult,
		Tags:           card.Tags,
		CreatedAt:      card.CreatedAt,
		UpdatedAt:      card.UpdatedAt,
		DeletedAt:      card.DeletedAt,
	}
	return cardResponse
}
