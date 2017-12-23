package trade

import (
	"github.com/go-kit/kit/endpoint"
	"strings"
	"errors"
	"context"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

//request
type TradeRequest struct {
	RequestType string
	Min int
	Max int
}

//response
type TradeResponse struct {
	Message string `json:"message"`
	Err     error `json:"err,omitempty"`
}

// endpoints wrapper
type Endpoints struct {
	TradeEndpoint endpoint.Endpoint
}

// creating Lorem Ipsum Endpoint
func MakeTradeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TradeRequest)

		var (
			txt string
			min, max int
		)

		min = req.Min
		max = req.Max

		if strings.EqualFold(req.RequestType, "Word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence"){
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "Paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}

		return TradeResponse{Message: txt}, nil
	}

}