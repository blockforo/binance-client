package binance_connector

import (
	"context"
	"encoding/json"
	"strconv"
)

type OrderPlacementService struct {
	websocketAPI            *WebsocketAPIClient
	symbol                  string
	side                    string
	orderType               string
	timeInForce             *string
	price                   *float64
	quantity                *float64
	quoteOrderQty           *float64
	newClientOrderID        *string
	newOrderRespType        *string
	stopPrice               *float64
	trailingDelta           *int64
	icebergQty              *float64
	strategyID              *int
	strategyType            *int
	selfTradePreventionMode *string
	recvWindow              *int64
}

func (s *OrderPlacementService) Symbol(symbol string) *OrderPlacementService {
	s.symbol = symbol
	return s
}

func (s *OrderPlacementService) Side(side string) *OrderPlacementService {
	s.side = side
	return s
}

func (s *OrderPlacementService) OrderType(orderType string) *OrderPlacementService {
	s.orderType = orderType
	return s
}

func (s *OrderPlacementService) TimeInForce(timeInForce string) *OrderPlacementService {
	s.timeInForce = &timeInForce
	return s
}

func (s *OrderPlacementService) Price(price float64) *OrderPlacementService {
	s.price = &price
	return s
}

func (s *OrderPlacementService) Quantity(quantity float64) *OrderPlacementService {
	s.quantity = &quantity
	return s
}

func (s *OrderPlacementService) QuoteOrderQty(quoteOrderQty float64) *OrderPlacementService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *OrderPlacementService) NewClientOrderID(newClientOrderID string) *OrderPlacementService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *OrderPlacementService) NewOrderRespType(newOrderRespType string) *OrderPlacementService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *OrderPlacementService) StopPrice(stopPrice float64) *OrderPlacementService {
	s.stopPrice = &stopPrice
	return s
}

func (s *OrderPlacementService) TrailingDelta(trailingDelta int64) *OrderPlacementService {
	s.trailingDelta = &trailingDelta
	return s
}

func (s *OrderPlacementService) IcebergQty(icebergQty float64) *OrderPlacementService {
	s.icebergQty = &icebergQty
	return s
}

func (s *OrderPlacementService) StrategyID(strategyID int) *OrderPlacementService {
	s.strategyID = &strategyID
	return s
}

func (s *OrderPlacementService) StrategyType(strategyType int) *OrderPlacementService {
	s.strategyType = &strategyType
	return s
}

func (s *OrderPlacementService) SelfTradePreventionMode(selfTradePreventionMode string) *OrderPlacementService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *OrderPlacementService) RecvWindow(recvWindow int64) *OrderPlacementService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderPlacementService) Do(ctx context.Context) (*OrderPlacementResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}

	if s.timeInForce != nil {
		parameters["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		parameters["price"] = strconv.FormatFloat(*s.price, 'f', -1, 64)
	}
	if s.quantity != nil {
		parameters["quantity"] = strconv.FormatFloat(*s.quantity, 'f', -1, 64)
	}
	if s.quoteOrderQty != nil {
		parameters["quoteOrderQty"] = strconv.FormatFloat(*s.quoteOrderQty, 'f', -1, 64)
	}
	if s.newClientOrderID != nil {
		parameters["newClientOrderId"] = *s.newClientOrderID
	}
	if s.newOrderRespType != nil {
		parameters["newOrderRespType"] = *s.newOrderRespType
	}
	if s.stopPrice != nil {
		parameters["stopPrice"] = strconv.FormatFloat(*s.stopPrice, 'f', -1, 64)
	}
	if s.trailingDelta != nil {
		parameters["trailingDelta"] = strconv.FormatInt(*s.trailingDelta, 10)
	}
	if s.icebergQty != nil {
		parameters["icebergQty"] = strconv.FormatFloat(*s.icebergQty, 'f', -1, 64)
	}
	if s.strategyID != nil {
		parameters["strategyId"] = strconv.Itoa(*s.strategyID)
	}
	if s.strategyType != nil {
		parameters["strategyType"] = strconv.Itoa(*s.strategyType)
	}
	if s.selfTradePreventionMode != nil {
		parameters["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "order.place",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderPlacementResponse OrderPlacementResponse
		err = json.Unmarshal(response, &orderPlacementResponse)
		if err != nil {
			return nil, err
		}
		return &orderPlacementResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderPlacementResponse struct {
	ID         string                `json:"id"`
	Status     int                   `json:"status"`
	Error      *WsAPIErrorResponse   `json:"error,omitempty"`
	Result     *OrderPlacementResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit     `json:"rateLimits"`
}

type OrderPlacementResult struct {
	Symbol              string      `json:"symbol,omitempty"`
	OrderID             int64       `json:"orderId,omitempty"`
	OrderListID         int64       `json:"orderListId,omitempty"`
	ClientOrderID       string      `json:"clientOrderId,omitempty"`
	TransactTime        int64       `json:"transactTime,omitempty"`
	Price               string      `json:"price,omitempty"`
	OrigQty             string      `json:"origQty,omitempty"`
	ExecutedQty         string      `json:"executedQty,omitempty"`
	CummulativeQuoteQty string      `json:"cummulativeQuoteQty,omitempty"`
	Status              string      `json:"status,omitempty"`
	TimeInForce         string      `json:"timeInForce,omitempty"`
	Type                string      `json:"type,omitempty"`
	Side                string      `json:"side,omitempty"`
	WorkingTime         int64       `json:"workingTime,omitempty"`
	IcebergQty          string      `json:"icebergQty,omitempty"`
	PreventedMatchID    int64       `json:"preventedMatchId,omitempty"`
	PreventedQuantity   string      `json:"preventedQuantity,omitempty"`
	StopPrice           string      `json:"stopPrice,omitempty"`
	StrategyID          int64       `json:"strategyId,omitempty"`
	StrategyType        int64       `json:"strategyType,omitempty"`
	TrailingDelta       string      `json:"trailingDelta,omitempty"`
	TrailingTime        int64       `json:"trailingTime,omitempty"`
	Fills               []OrderFill `json:"fills,omitempty"`
	SelfTradePrevention string      `json:"selfTradePreventionMode,omitempty"`
}

type OrderFill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeID         int64  `json:"tradeId"`
}

type TestOrderPlacementService struct {
	websocketAPI            *WebsocketAPIClient
	symbol                  string
	side                    string
	orderType               string
	timeInForce             *string
	price                   *float64
	quantity                *float64
	quoteOrderQty           *float64
	newClientOrderID        *string
	newOrderRespType        *string
	stopPrice               *float64
	trailingDelta           *int64
	icebergQty              *float64
	strategyID              *int
	strategyType            *int
	selfTradePreventionMode *string
	recvWindow              *int64
}

func (s *TestOrderPlacementService) Symbol(symbol string) *TestOrderPlacementService {
	s.symbol = symbol
	return s
}

func (s *TestOrderPlacementService) Side(side string) *TestOrderPlacementService {
	s.side = side
	return s
}

func (s *TestOrderPlacementService) OrderType(orderType string) *TestOrderPlacementService {
	s.orderType = orderType
	return s
}

func (s *TestOrderPlacementService) TimeInForce(timeInForce string) *TestOrderPlacementService {
	s.timeInForce = &timeInForce
	return s
}

func (s *TestOrderPlacementService) Price(price float64) *TestOrderPlacementService {
	s.price = &price
	return s
}

func (s *TestOrderPlacementService) Quantity(quantity float64) *TestOrderPlacementService {
	s.quantity = &quantity
	return s
}

func (s *TestOrderPlacementService) QuoteOrderQty(quoteOrderQty float64) *TestOrderPlacementService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *TestOrderPlacementService) NewClientOrderID(newClientOrderID string) *TestOrderPlacementService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *TestOrderPlacementService) NewOrderRespType(newOrderRespType string) *TestOrderPlacementService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *TestOrderPlacementService) StopPrice(stopPrice float64) *TestOrderPlacementService {
	s.stopPrice = &stopPrice
	return s
}

func (s *TestOrderPlacementService) TrailingDelta(trailingDelta int64) *TestOrderPlacementService {
	s.trailingDelta = &trailingDelta
	return s
}

func (s *TestOrderPlacementService) IcebergQty(icebergQty float64) *TestOrderPlacementService {
	s.icebergQty = &icebergQty
	return s
}

func (s *TestOrderPlacementService) StrategyID(strategyID int) *TestOrderPlacementService {
	s.strategyID = &strategyID
	return s
}

func (s *TestOrderPlacementService) StrategyType(strategyType int) *TestOrderPlacementService {
	s.strategyType = &strategyType
	return s
}

func (s *TestOrderPlacementService) SelfTradePreventionMode(selfTradePreventionMode string) *TestOrderPlacementService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *TestOrderPlacementService) RecvWindow(recvWindow int64) *TestOrderPlacementService {
	s.recvWindow = &recvWindow
	return s
}

func (s *TestOrderPlacementService) Do(ctx context.Context) (*OrderPlacementResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}

	if s.timeInForce != nil {
		parameters["timeInForce"] = *s.timeInForce
	}
	if s.price != nil {
		parameters["price"] = strconv.FormatFloat(*s.price, 'f', -1, 64)
	}
	if s.quantity != nil {
		parameters["quantity"] = strconv.FormatFloat(*s.quantity, 'f', -1, 64)
	}
	if s.quoteOrderQty != nil {
		parameters["quoteOrderQty"] = strconv.FormatFloat(*s.quoteOrderQty, 'f', -1, 64)
	}
	if s.newClientOrderID != nil {
		parameters["newClientOrderId"] = *s.newClientOrderID
	}
	if s.newOrderRespType != nil {
		parameters["newOrderRespType"] = *s.newOrderRespType
	}
	if s.stopPrice != nil {
		parameters["stopPrice"] = strconv.FormatFloat(*s.stopPrice, 'f', -1, 64)
	}
	if s.trailingDelta != nil {
		parameters["trailingDelta"] = strconv.FormatInt(*s.trailingDelta, 10)
	}
	if s.icebergQty != nil {
		parameters["icebergQty"] = strconv.FormatFloat(*s.icebergQty, 'f', -1, 64)
	}
	if s.strategyID != nil {
		parameters["strategyId"] = strconv.Itoa(*s.strategyID)
	}
	if s.strategyType != nil {
		parameters["strategyType"] = strconv.Itoa(*s.strategyType)
	}
	if s.selfTradePreventionMode != nil {
		parameters["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "order.test",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderPlacementResponse OrderPlacementResponse
		err = json.Unmarshal(response, &orderPlacementResponse)
		if err != nil {
			return nil, err
		}
		return &orderPlacementResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderStatusService struct {
	websocketAPI      *WebsocketAPIClient
	symbol            string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (s *OrderStatusService) Symbol(symbol string) *OrderStatusService {
	s.symbol = symbol
	return s
}

func (s *OrderStatusService) OrderId(orderId int64) *OrderStatusService {
	s.orderId = &orderId
	return s
}

func (s *OrderStatusService) ClientOrderID(origClientOrderId string) *OrderStatusService {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *OrderStatusService) RecvWindow(recvWindow int64) *OrderStatusService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderStatusService) Do(ctx context.Context) (*OrderStatusResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.orderId != nil {
		parameters["orderId"] = strconv.FormatInt(*s.orderId, 10)
	}

	if s.origClientOrderId != nil {
		parameters["origClientOrderId"] = *s.origClientOrderId
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "order.status",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderStatusResponse OrderStatusResponse
		err = json.Unmarshal(response, &orderStatusResponse)
		if err != nil {
			return nil, err
		}
		return &orderStatusResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderStatusResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *OrderStatusResult  `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type OrderStatusResult struct {
	Symbol                  string `json:"symbol"`
	OrderID                 int64  `json:"orderId"`
	OrderListID             int    `json:"orderListId"`
	ClientOrderID           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice,omitempty"`
	IcebergQty              string `json:"icebergQty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StrategyID              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

type OrderCancelService struct {
	websocketAPI       *WebsocketAPIClient
	symbol             string
	orderID            *int64
	origClientOrderID  *string
	newClientOrderID   *string
	cancelRestrictions *string
	recvWindow         *int64
}

func (s *OrderCancelService) Symbol(symbol string) *OrderCancelService {
	s.symbol = symbol
	return s
}

func (s *OrderCancelService) OrderID(orderID int64) *OrderCancelService {
	s.orderID = &orderID
	return s
}

func (s *OrderCancelService) OrigClientOrderID(origClientOrderID string) *OrderCancelService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *OrderCancelService) NewClientOrderId(newClientOrderID string) *OrderCancelService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *OrderCancelService) CancelRestrictions(cancelRestrictions string) *OrderCancelService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

func (s *OrderCancelService) RecvWindow(recvWindow int64) *OrderCancelService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderCancelService) Do(ctx context.Context) (*OrderCancelResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.orderID != nil {
		parameters["orderId"] = strconv.FormatInt(*s.orderID, 10)
	}

	if s.origClientOrderID != nil {
		parameters["origClientOrderId"] = *s.origClientOrderID
	}

	if s.newClientOrderID != nil {
		parameters["newClientOrderId"] = *s.newClientOrderID
	}

	if s.cancelRestrictions != nil {
		parameters["cancelRestrictions"] = *s.cancelRestrictions
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "order.cancel",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderCancelResponse OrderCancelResponse
		err = json.Unmarshal(response, &orderCancelResponse)
		if err != nil {
			return nil, err
		}
		return &orderCancelResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderCancelResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     *OrderCancelResult  `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type OrderCancelResult struct {
	Symbol                  string `json:"symbol"`
	OrigClientOrderId       string `json:"origClientOrderId"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	IcebergQty              string `json:"icebergQty,omitempty"`
	PreventedMatchId        int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity       string `json:"preventedQuantity,omitempty"`
	StopPrice               string `json:"stopPrice,omitempty"`
	StrategyID              int64  `json:"strategyId,omitempty"`
	StrategyType            int64  `json:"strategyType,omitempty"`
	TrailingDelta           string `json:"trailingDelta,omitempty"`
	TrailingTime            int64  `json:"trailingTime,omitempty"`
}

type OrderCancelReplaceService struct {
	websocketAPI            *WebsocketAPIClient
	symbol                  string
	cancelReplaceMode       string
	cancelOrderID           *int64
	cancelOrigClientOrderID *string
	cancelNewClientOrderID  *string
	side                    string
	orderType               string
	timeInForce             *string
	price                   *float64
	quantity                *float64
	quoteOrderQty           *float64
	newClientOrderID        *string
	newOrderRespType        *string
	stopPrice               *float64
	trailingDelta           *float64
	icebergQty              *float64
	strategyID              *int
	strategyType            *int
	selfTradePreventionMode *string
	cancelRestrictions      *string
	recvWindow              *int64
}

func (s *OrderCancelReplaceService) Symbol(symbol string) *OrderCancelReplaceService {
	s.symbol = symbol
	return s
}

func (s *OrderCancelReplaceService) CancelReplaceMode(cancelReplaceMode string) *OrderCancelReplaceService {
	s.cancelReplaceMode = cancelReplaceMode
	return s
}

func (s *OrderCancelReplaceService) CancelOrderID(cancelOrderID int64) *OrderCancelReplaceService {
	s.cancelOrderID = &cancelOrderID
	return s
}

func (s *OrderCancelReplaceService) CancelOrigClientOrderId(cancelOrigClientOrderID string) *OrderCancelReplaceService {
	s.cancelOrigClientOrderID = &cancelOrigClientOrderID
	return s
}

func (s *OrderCancelReplaceService) CancelNewClientOrderId(cancelNewClientOrderID string) *OrderCancelReplaceService {
	s.cancelNewClientOrderID = &cancelNewClientOrderID
	return s
}

func (s *OrderCancelReplaceService) Side(side string) *OrderCancelReplaceService {
	s.side = side
	return s
}

func (s *OrderCancelReplaceService) OrderType(orderType string) *OrderCancelReplaceService {
	s.orderType = orderType
	return s
}

func (s *OrderCancelReplaceService) TimeInForce(timeInForce string) *OrderCancelReplaceService {
	s.timeInForce = &timeInForce
	return s
}

func (s *OrderCancelReplaceService) Price(price float64) *OrderCancelReplaceService {
	s.price = &price
	return s
}

func (s *OrderCancelReplaceService) Quantity(quantity float64) *OrderCancelReplaceService {
	s.quantity = &quantity
	return s
}

func (s *OrderCancelReplaceService) QuoteOrderQty(quoteOrderQty float64) *OrderCancelReplaceService {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *OrderCancelReplaceService) NewClientOrderId(newClientOrderID string) *OrderCancelReplaceService {
	s.newClientOrderID = &newClientOrderID
	return s
}

func (s *OrderCancelReplaceService) NewOrderRespType(newOrderRespType string) *OrderCancelReplaceService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *OrderCancelReplaceService) StopPrice(stopPrice float64) *OrderCancelReplaceService {
	s.stopPrice = &stopPrice
	return s
}

func (s *OrderCancelReplaceService) TrailingDelta(trailingDelta float64) *OrderCancelReplaceService {
	s.trailingDelta = &trailingDelta
	return s
}

func (s *OrderCancelReplaceService) IcebergQty(icebergQty float64) *OrderCancelReplaceService {
	s.icebergQty = &icebergQty
	return s
}

func (s *OrderCancelReplaceService) StrategyID(strategyID int) *OrderCancelReplaceService {
	s.strategyID = &strategyID
	return s
}

func (s *OrderCancelReplaceService) StrategyType(strategyType int) *OrderCancelReplaceService {
	s.strategyType = &strategyType
	return s
}

func (s *OrderCancelReplaceService) SelfTradePreventionMode(selfTradePreventionMode string) *OrderCancelReplaceService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *OrderCancelReplaceService) CancelRestrictions(cancelRestrictions string) *OrderCancelReplaceService {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

func (s *OrderCancelReplaceService) RecvWindow(recvWindow int64) *OrderCancelReplaceService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderCancelReplaceService) Do(ctx context.Context) (*OrderCancelReplaceResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol":            s.symbol,
		"cancelReplaceMode": s.cancelReplaceMode,
		"side":              s.side,
		"type":              s.orderType,
	}

	if s.cancelOrderID != nil {
		parameters["cancelOrderId"] = strconv.FormatInt(*s.cancelOrderID, 10)
	}

	if s.cancelOrigClientOrderID != nil {
		parameters["cancelOrigClientOrderId"] = *s.cancelOrigClientOrderID
	}

	if s.cancelNewClientOrderID != nil {
		parameters["cancelNewClientOrderId"] = *s.cancelNewClientOrderID
	}

	if s.timeInForce != nil {
		parameters["timeInForce"] = *s.timeInForce
	}

	if s.price != nil {
		parameters["price"] = strconv.FormatFloat(*s.price, 'f', -1, 64)
	}

	if s.quantity != nil {
		parameters["quantity"] = strconv.FormatFloat(*s.quantity, 'f', -1, 64)
	}

	if s.quoteOrderQty != nil {
		parameters["quoteOrderQty"] = strconv.FormatFloat(*s.quoteOrderQty, 'f', -1, 64)
	}

	if s.newClientOrderID != nil {
		parameters["newClientOrderId"] = *s.newClientOrderID
	}

	if s.newOrderRespType != nil {
		parameters["newOrderRespType"] = *s.newOrderRespType
	}

	if s.stopPrice != nil {
		parameters["stopPrice"] = strconv.FormatFloat(*s.stopPrice, 'f', -1, 64)
	}

	if s.trailingDelta != nil {
		parameters["trailingDelta"] = strconv.FormatFloat(*s.trailingDelta, 'f', -1, 64)
	}

	if s.icebergQty != nil {
		parameters["icebergQty"] = strconv.FormatFloat(*s.icebergQty, 'f', -1, 64)
	}

	if s.strategyID != nil {
		parameters["strategyId"] = strconv.Itoa(*s.strategyID)
	}

	if s.strategyType != nil {
		parameters["strategyType"] = strconv.Itoa(*s.strategyType)
	}

	if s.selfTradePreventionMode != nil {
		parameters["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}

	if s.cancelRestrictions != nil {
		parameters["cancelRestrictions"] = *s.cancelRestrictions
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "order.cancelReplace",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderCancelReplaceResponse OrderCancelReplaceResponse
		err = json.Unmarshal(response, &orderCancelReplaceResponse)
		if err != nil {
			return nil, err
		}
		return &orderCancelReplaceResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderCancelReplaceResponse struct {
	ID         string                    `json:"id"`
	Status     int                       `json:"status"`
	Error      *WsAPIErrorResponse       `json:"error,omitempty"`
	Result     *OrderCancelReplaceResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit         `json:"rateLimits"`
}

type OrderCancelReplaceResult struct {
	CancelResult     string                `json:"cancelResult,omitempty"`
	NewOrderResult   string                `json:"newOrderResult,omitempty"`
	CancelResponse   *OrderCancelledResult `json:"cancelResponse,omitempty"`
	NewOrderResponse *OrderPlacedResult    `json:"newOrderResponse,omitempty"`
}

type OrderPlacedResult struct {
	Symbol                  string  `json:"symbol,omitempty"`
	OrderId                 int64   `json:"orderId,omitempty"`
	OrderListId             int     `json:"orderListId,omitempty"`
	ClientOrderId           string  `json:"clientOrderId,omitempty"`
	TransactTime            int64   `json:"transactTime,omitempty"`
	Price                   string  `json:"price,omitempty"`
	OrigQty                 string  `json:"origQty,omitempty"`
	ExecutedQty             string  `json:"executedQty,omitempty"`
	CummulativeQuoteQty     string  `json:"cummulativeQuoteQty,omitempty"`
	Status                  string  `json:"status,omitempty"`
	TimeInForce             string  `json:"timeInForce,omitempty"`
	Type                    string  `json:"type,omitempty"`
	Side                    string  `json:"side,omitempty"`
	WorkingTime             uint64  `json:"workingTime,omitempty"`
	Fills                   []*Fill `json:"fills,omitempty"`
	SelfTradePreventionMode string  `json:"selfTradePreventionMode,omitempty"`
}

type Fill struct {
	Price           string `json:"price,omitempty"`
	Qty             string `json:"qty,omitempty"`
	Commission      string `json:"commission,omitempty"`
	CommissionAsset string `json:"commissionAsset,omitempty"`
	TradeID         int64  `json:"tradeId,omitempty"`
}

type OrderCancelledResult struct {
	Symbol                  string `json:"symbol,omitempty"`
	OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
	OrderId                 int64  `json:"orderId,omitempty"`
	OrderListId             int    `json:"orderListId,omitempty"`
	ClientOrderId           string `json:"clientOrderId,omitempty"`
	Price                   string `json:"price,omitempty"`
	OrigQty                 string `json:"origQty,omitempty"`
	ExecutedQty             string `json:"executedQty,omitempty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty,omitempty"`
	Status                  string `json:"status,omitempty"`
	TimeInForce             string `json:"timeInForce,omitempty"`
	Type                    string `json:"type,omitempty"`
	Side                    string `json:"side,omitempty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
}

type OpenOrdersStatusService struct {
	websocketAPI *WebsocketAPIClient
	symbol       *string
	recvWindow   *int64
}

func (s *OpenOrdersStatusService) Symbol(symbol string) *OpenOrdersStatusService {
	s.symbol = &symbol
	return s
}

func (s *OpenOrdersStatusService) RecvWindow(recvWindow int64) *OpenOrdersStatusService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OpenOrdersStatusService) Do(ctx context.Context) (*OpenOrdersStatusResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{}

	if s.symbol != nil {
		parameters["symbol"] = *s.symbol
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "openOrders.status",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var openOrdersStatusResponse OpenOrdersStatusResponse
		err = json.Unmarshal(response, &openOrdersStatusResponse)
		if err != nil {
			return nil, err
		}
		return &openOrdersStatusResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OpenOrdersStatusResponse struct {
	ID         string              `json:"id"`
	Status     int                 `json:"status"`
	Error      *WsAPIErrorResponse `json:"error,omitempty"`
	Result     []*OpenOrdersResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit   `json:"rateLimits"`
}

type OpenOrdersResult struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty"`
	Time                    uint64 `json:"time"`
	UpdateTime              uint64 `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             uint64 `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type OpenOrdersCancelAllService struct {
	websocketAPI *WebsocketAPIClient
	symbol       string
	recvWindow   *int64
}

func (s *OpenOrdersCancelAllService) Symbol(symbol string) *OpenOrdersCancelAllService {
	s.symbol = symbol
	return s
}

func (s *OpenOrdersCancelAllService) RecvWindow(recvWindow int64) *OpenOrdersCancelAllService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OpenOrdersCancelAllService) Do(ctx context.Context) (*OpenOrdersCancelAllResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "openOrders.cancelAll",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var openOrdersCancelAllResponse OpenOrdersCancelAllResponse
		err = json.Unmarshal(response, &openOrdersCancelAllResponse)
		if err != nil {
			return nil, err
		}
		return &openOrdersCancelAllResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OpenOrdersCancelAllResponse struct {
	ID         string                    `json:"id"`
	Status     int                       `json:"status"`
	Error      *WsAPIErrorResponse       `json:"error,omitempty"`
	Result     []*OpenOrdersCancelResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit         `json:"rateLimits"`
}

type OpenOrdersCancelResult struct {
	Symbol                  string `json:"symbol"`
	OrigClientOrderID       string `json:"origClientOrderId"`
	OrderId                 int64  `json:"orderId"`
	OrderListID             int    `json:"orderListId"`
	ClientOrderID           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty"`
	StrategyID              int64  `json:"strategyId"`
	StrategyType            int64  `json:"strategyType"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type OrderListPlaceService struct {
	websocketAPI            *WebsocketAPIClient
	symbol                  string
	side                    string
	price                   float64
	quantity                float64
	listClientOrderID       *string
	limitClientOrderID      *string
	limitIcebergQty         *float64
	limitStrategyID         *int
	limitStrategyType       *int
	stopPrice               *float64
	trailingDelta           *float64
	stopClientOrderID       *string
	stopLimitPrice          *float64
	stopLimitTimeInForce    *string
	stopIcebergQty          *float64
	stopStrategyID          *int
	stopStrategyType        *int
	newOrderRespType        *string
	selfTradePreventionMode *string
	recvWindow              *int64
}

func (s *OrderListPlaceService) Symbol(symbol string) *OrderListPlaceService {
	s.symbol = symbol
	return s
}

func (s *OrderListPlaceService) Side(side string) *OrderListPlaceService {
	s.side = side
	return s
}

func (s *OrderListPlaceService) Price(price float64) *OrderListPlaceService {
	s.price = price
	return s
}

func (s *OrderListPlaceService) Quantity(quantity float64) *OrderListPlaceService {
	s.quantity = quantity
	return s
}

func (s *OrderListPlaceService) ListClientOrderId(listClientOrderID string) *OrderListPlaceService {
	s.listClientOrderID = &listClientOrderID
	return s
}

func (s *OrderListPlaceService) LimitClientOrderID(limitClientOrderID string) *OrderListPlaceService {
	s.limitClientOrderID = &limitClientOrderID
	return s
}

func (s *OrderListPlaceService) LimitIcebergQty(limitIcebergQty float64) *OrderListPlaceService {
	s.limitIcebergQty = &limitIcebergQty
	return s
}

func (s *OrderListPlaceService) LimitStrategyID(limitStrategyID int) *OrderListPlaceService {
	s.limitStrategyID = &limitStrategyID
	return s
}

func (s *OrderListPlaceService) LimitStrategyType(limitStrategyType int) *OrderListPlaceService {
	s.limitStrategyType = &limitStrategyType
	return s
}

func (s *OrderListPlaceService) StopPrice(stopPrice float64) *OrderListPlaceService {
	s.stopPrice = &stopPrice
	return s
}

func (s *OrderListPlaceService) StopClientOrderID(stopClientOrderID string) *OrderListPlaceService {
	s.stopClientOrderID = &stopClientOrderID
	return s
}

func (s *OrderListPlaceService) StopLimitPrice(stopLimitPrice float64) *OrderListPlaceService {
	s.stopLimitPrice = &stopLimitPrice
	return s
}

func (s *OrderListPlaceService) StopLimitTimeInForce(stopLimitTimeInForce string) *OrderListPlaceService {
	s.stopLimitTimeInForce = &stopLimitTimeInForce
	return s
}

func (s *OrderListPlaceService) StopIcebergQty(stopIcebergQty float64) *OrderListPlaceService {
	s.stopIcebergQty = &stopIcebergQty
	return s
}

func (s *OrderListPlaceService) StopStrategyID(stopStrategyID int) *OrderListPlaceService {
	s.stopStrategyID = &stopStrategyID
	return s
}

func (s *OrderListPlaceService) StopStrategyType(stopStrategyType int) *OrderListPlaceService {
	s.stopStrategyType = &stopStrategyType
	return s
}

func (s *OrderListPlaceService) NewOrderRespType(newOrderRespType string) *OrderListPlaceService {
	s.newOrderRespType = &newOrderRespType
	return s
}

func (s *OrderListPlaceService) SelfTradePreventionMode(selfTradePreventionMode string) *OrderListPlaceService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *OrderListPlaceService) RecvWindow(recvWindow int64) *OrderListPlaceService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderListPlaceService) Do(ctx context.Context) (*OrderListPlaceResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol":   s.symbol,
		"side":     s.side,
		"price":    strconv.FormatFloat(s.price, 'f', -1, 64),
		"quantity": strconv.FormatFloat(s.quantity, 'f', -1, 64),
	}

	if s.listClientOrderID != nil {
		parameters["listClientOrderId"] = *s.listClientOrderID
	}

	if s.limitClientOrderID != nil {
		parameters["limitClientOrderId"] = *s.limitClientOrderID
	}

	if s.limitIcebergQty != nil {
		parameters["limitIcebergQty"] = strconv.FormatFloat(*s.limitIcebergQty, 'f', -1, 64)
	}

	if s.limitStrategyID != nil {
		strconv.Itoa(*s.limitStrategyID)
	}

	if s.limitStrategyType != nil {
		strconv.Itoa(*s.limitStrategyType)
	}

	if s.stopPrice != nil {
		parameters["stopPrice"] = strconv.FormatFloat(*s.stopPrice, 'f', -1, 64)
	}

	if s.trailingDelta != nil {
		parameters["trailingDelta"] = strconv.FormatFloat(*s.trailingDelta, 'f', -1, 64)
	}

	if s.stopClientOrderID != nil {
		parameters["stopClientOrderId"] = *s.stopClientOrderID
	}

	if s.stopLimitPrice != nil {
		parameters["stopLimitPrice"] = strconv.FormatFloat(*s.stopLimitPrice, 'f', -1, 64)
	}

	if s.stopLimitTimeInForce != nil {
		parameters["stopLimitTimeInForce"] = *s.stopLimitTimeInForce
	}

	if s.stopIcebergQty != nil {
		parameters["stopIcebergQty"] = strconv.FormatFloat(*s.stopIcebergQty, 'f', -1, 64)
	}

	if s.stopStrategyID != nil {
		parameters["stopStrategyId"] = strconv.Itoa(*s.stopStrategyID)
	}

	if s.stopStrategyType != nil {
		parameters["stopStrategyType"] = strconv.Itoa(*s.stopStrategyType)
	}

	if s.newOrderRespType != nil {
		parameters["newOrderRespType"] = *s.newOrderRespType
	}

	if s.selfTradePreventionMode != nil {
		parameters["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "orderList.place",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderListPlaceResponse OrderListPlaceResponse
		err = json.Unmarshal(response, &orderListPlaceResponse)
		if err != nil {
			return nil, err
		}
		return &orderListPlaceResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderListPlaceResponse struct {
	ID         string                `json:"id"`
	Status     int                   `json:"status"`
	Error      *WsAPIErrorResponse   `json:"error,omitempty"`
	Result     *OrderListPlaceResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit     `json:"rateLimits"`
}

type OrderListPlaceResult struct {
	OrderListId       int            `json:"orderListId"`
	ContingencyType   string         `json:"contingencyType"`
	ListStatusType    string         `json:"listStatusType"`
	ListOrderStatus   string         `json:"listOrderStatus"`
	ListClientOrderId string         `json:"listClientOrderId"`
	TransactionTime   int64          `json:"transactionTime"`
	Symbol            string         `json:"symbol"`
	Orders            []*OrderInfo   `json:"orders"`
	OrderReports      []*OrderReport `json:"orderReports"`
}

type OrderInfo struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OrderReport struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            uint64 `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             uint64 `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type OrderListStatusService struct {
	websocketAPI      *WebsocketAPIClient
	origClientOrderId *string
	orderListId       *int
	recvWindow        *int64
}

func (s *OrderListStatusService) OrigClientOrderId(origClientOrderId string) *OrderListStatusService {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *OrderListStatusService) OrderListId(orderListId int) *OrderListStatusService {
	s.orderListId = &orderListId
	return s
}

func (s *OrderListStatusService) RecvWindow(recvWindow int64) *OrderListStatusService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderListStatusService) Do(ctx context.Context) (*OrderListStatusResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{}

	if s.origClientOrderId != nil {
		parameters["origClientOrderId"] = *s.origClientOrderId
	}

	if s.orderListId != nil {
		parameters["orderListId"] = strconv.Itoa(*s.orderListId)
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "orderList.status",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderListStatusResponse OrderListStatusResponse
		err = json.Unmarshal(response, &orderListStatusResponse)
		if err != nil {
			return nil, err
		}
		return &orderListStatusResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderListStatusResponse struct {
	ID         string                 `json:"id"`
	Status     int                    `json:"status"`
	Error      *WsAPIErrorResponse    `json:"error,omitempty"`
	Result     *OrderListStatusResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit      `json:"rateLimits"`
}

type OrderListStatusResult struct {
	OrderListId       int                     `json:"orderListId"`
	ContingencyType   string                  `json:"contingencyType"`
	ListStatusType    string                  `json:"listStatusType"`
	ListOrderStatus   string                  `json:"listOrderStatus"`
	ListClientOrderId string                  `json:"listClientOrderId"`
	TransactionTime   uint64                  `json:"transactionTime"`
	Symbol            string                  `json:"symbol"`
	Orders            []*OrderListStatusOrder `json:"orders"`
}

type OrderListStatusOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OrderListCancelService struct {
	websocketAPI      *WebsocketAPIClient
	symbol            string
	orderListId       *int
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (s *OrderListCancelService) Symbol(symbol string) *OrderListCancelService {
	s.symbol = symbol
	return s
}

func (s *OrderListCancelService) OrderListId(orderListId int) *OrderListCancelService {
	s.orderListId = &orderListId
	return s
}

func (s *OrderListCancelService) ListClientOrderId(listClientOrderId string) *OrderListCancelService {
	s.listClientOrderId = &listClientOrderId
	return s
}

func (s *OrderListCancelService) NewClientOrderId(newClientOrderId string) *OrderListCancelService {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *OrderListCancelService) RecvWindow(recvWindow int64) *OrderListCancelService {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderListCancelService) Do(ctx context.Context) (*OrderListCancelResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{
		"symbol": s.symbol,
	}

	if s.orderListId != nil {
		parameters["orderListId"] = strconv.Itoa(*s.orderListId)
	}

	if s.listClientOrderId != nil {
		parameters["listClientOrderId"] = *s.listClientOrderId
	}

	if s.newClientOrderId != nil {
		parameters["newClientOrderId"] = *s.newClientOrderId
	}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "orderList.cancel",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var orderListCancelResponse OrderListCancelResponse
		err = json.Unmarshal(response, &orderListCancelResponse)
		if err != nil {
			return nil, err
		}
		return &orderListCancelResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OrderListCancelResponse struct {
	ID         string                 `json:"id"`
	Status     int                    `json:"status"`
	Error      *WsAPIErrorResponse    `json:"error,omitempty"`
	Result     *OrderListCancelResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit      `json:"rateLimits"`
}

type OrderListCancelResult struct {
	OrderListId       int64                         `json:"orderListId"`
	ContingencyType   string                        `json:"contingencyType"`
	ListStatusType    string                        `json:"listStatusType"`
	ListOrderStatus   string                        `json:"listOrderStatus"`
	ListClientOrderId string                        `json:"listClientOrderId"`
	TransactionTime   uint64                        `json:"transactionTime"`
	Symbol            string                        `json:"symbol"`
	Orders            []*OrderListCancelOrder       `json:"orders"`
	OrderReports      []*OrderListCancelOrderReport `json:"orderReports"`
}

type OrderListCancelOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OrderListCancelOrderReport struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int64  `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactionTime         uint64 `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type OpenOrderListsStatusService struct {
	websocketAPI *WebsocketAPIClient
	recvWindow   *int64
}

func (s *OpenOrderListsStatusService) Do(ctx context.Context) (*OpenOrderListsStatusResponse, error) {
	// Generate a random UUID
	id, err := getUUID()

	// Make sure it was generated correctly
	if err != nil {
		return nil, err
	}

	parameters := map[string]string{}

	if s.recvWindow != nil {
		parameters["recvWindow"] = strconv.FormatInt(*s.recvWindow, 10)
	}

	signedParams, err := websocketAPISignature(s.websocketAPI.APIKey, s.websocketAPI.APISecret, parameters)
	if err != nil {
		panic(err)
	}

	payload := map[string]interface{}{
		"id":     id,
		"method": "openOrderLists.status",
		"params": signedParams,
	}

	messageCh := make(chan []byte)
	s.websocketAPI.ReqResponseMap[id] = messageCh

	err2 := s.websocketAPI.SendMessage(payload)
	if err2 != nil {
		return nil, err2
	}

	defer delete(s.websocketAPI.ReqResponseMap, id)

	select {
	case response := <-messageCh:
		var openOrderListsStatusResponse OpenOrderListsStatusResponse
		err = json.Unmarshal(response, &openOrderListsStatusResponse)
		if err != nil {
			return nil, err
		}
		return &openOrderListsStatusResponse, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type OpenOrderListsStatusResponse struct {
	ID         string                  `json:"id"`
	Status     int                     `json:"status"`
	Error      *WsAPIErrorResponse     `json:"error,omitempty"`
	Result     []*OpenOrderListsResult `json:"result,omitempty"`
	RateLimits []*WsAPIRateLimit       `json:"rateLimits"`
}

type OpenOrderListsResult struct {
	OrderListId       int                  `json:"orderListId"`
	ContingencyType   string               `json:"contingencyType"`
	ListStatusType    string               `json:"listStatusType"`
	ListOrderStatus   string               `json:"listOrderStatus"`
	ListClientOrderId string               `json:"listClientOrderId"`
	TransactionTime   uint64               `json:"transactionTime"`
	Symbol            string               `json:"symbol"`
	Orders            []*OpenOrderListItem `json:"orders"`
}

type OpenOrderListItem struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}
