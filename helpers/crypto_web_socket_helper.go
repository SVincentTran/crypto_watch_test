package helpers

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Subscription struct {
	StreamSubscription `json:"streamSubscription"`
}

type StreamSubscription struct {
	Resource string `json:"resource"`
}

type SubscribeRequest struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Update struct {
	MarketUpdate struct {
		Market struct {
			MarketId int `json:"marketId,string"`
		} `json:"market"`

		TradesUpdate struct {
			Trades []Trade `json:"trades"`
		} `json:"tradesUpdate"`
	} `json:"marketUpdate"`
}

type Trade struct {
	Timestamp     int `json:"timestamp,string"`
	TimestampNano int `json:"timestampNano,string"`

	Price  string `json:"priceStr"`
	Amount string `json:"amountStr"`
}

const APIKEY = "W537CZ928IQU4A0HAM57"

var update Update

func InitWebSocketClient() (bool, error) {
	// First initialization of the websocket connection
	var err error
	client, _, err := websocket.DefaultDialer.Dial("wss://stream.cryptowat.ch/connect?apikey="+APIKEY, nil)
	if err != nil {
		log.Printf("Error while connection to websocket %s", err)
		return false, err
	}
	defer client.Close()

	_, message, err := client.ReadMessage()
	if err != nil {
		log.Printf("Error while reading auth message %s", err)
		return false, err
	}
	var authResult struct {
		AuthenticationResult struct {
			Status string `json:"status"`
		} `json:"authenticationResult"`
	}

	err = json.Unmarshal(message, &authResult)
	if err != nil {
		log.Printf("Error while unmarshalling auth result %s", err)
		return false, err
	}

	// Subscribe to cryptowatch resources for the first time
	resources := []string{
		"markets:4:trades",
	}

	subMessage := struct {
		Subcribe SubscribeRequest `json:"subscribe"`
	}{}

	for _, resource := range resources {
		subMessage.Subcribe.Subscriptions = append(subMessage.Subcribe.Subscriptions,
			Subscription{
				StreamSubscription: StreamSubscription{
					Resource: resource,
				},
			},
		)
	}

	msg, _ := json.Marshal(subMessage)
	err = client.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Printf("Error while subscribing to resources %s", err)
		return false, err
	}

	for {
		_, message, err := client.ReadMessage()
		if err != nil {
			log.Printf("Error while reading from connection %s", err)
		}

		err = json.Unmarshal(message, &update)
		if err != nil {
			log.Printf("Error while unmarshalling message %s", err)
		} else {
			if len(update.MarketUpdate.TradesUpdate.Trades) != 0 {
				for _, trade := range update.MarketUpdate.TradesUpdate.Trades {
					log.Printf("ETH/USD trade on market %d: %s %s", update.MarketUpdate.Market.MarketId, trade.Price, trade.Amount)
				}
			}
		}
	}
}

func GetUpdates() Update {
	return update
}
