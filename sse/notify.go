package sse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

const format = "event:%s\ndata:%s\n\n"

type eventMessage struct {
	EventName string
	Data      any
}

func NewEventMessage(eventName string, data any) eventMessage {
	return eventMessage{
		EventName: eventName,
		Data:      data,
	}
}

type Notifier interface {
	BroadcastMessage(msg eventMessage)
	BoadcastMessageToAccount(msg eventMessage, accountId int64, userId *int64)
	SendMessageToUser(msg eventMessage, userId int64)
}

type HandlerEvent struct {
	m                sync.Mutex
	clients          map[int64]*client
	clientsByAccount map[int64][]int64
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clients:          make(map[int64]*client),
		clientsByAccount: make(map[int64][]int64),
	}
}

func (h *HandlerEvent) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	accountIdParam := r.URL.Query().Get("accountId")
	userIdParam := r.URL.Query().Get("userId")

	accountId, err := strconv.ParseInt(accountIdParam, 10, 64)
	if err != nil {
		LogError(err, "Invalid accountId parameter")
		return
	}

	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		LogError(err, "Invalid userId parameter")
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c := newClient(userId, accountId)

	h.registerClient(c)
	LogInfo(fmt.Sprintf("Connected client %d of account %d\n", userId, accountId))
	c.onLine(r.Context(), w, flusher)
	LogInfo(fmt.Sprintf("Disconnected client %d of account %d\n", userId, accountId))
	h.unregisterClient(userId)
}

func (h *HandlerEvent) registerClient(c *client) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clients[c.UserID] = c
	clientsOfAccount := h.clientsByAccount[c.AccountID]
	clientsOfAccount = append(clientsOfAccount, c.UserID)
	h.clientsByAccount[c.AccountID] = clientsOfAccount
}

func (h *HandlerEvent) unregisterClient(id int64) {
	h.m.Lock()
	defer h.m.Unlock()
	delete(h.clients, id)
}

func (c *client) onLine(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case m := <-c.sendMessage:
			data, err := json.Marshal(m.Data)
			if err != nil {
				LogError(err, "")
			}

			fmt.Fprintf(w, format, m.EventName, string(data))
			flusher.Flush()
		case <-ctx.Done():
			return
		}
	}
}

func (h *HandlerEvent) BroadcastMessage(msg eventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, client := range h.clients {
		client.sendMessage <- msg
	}
}

func (h *HandlerEvent) BoadcastMessageToAccount(msg eventMessage, accountId int64, userId *int64) {
	h.m.Lock()
	defer h.m.Unlock()
	clientsOfAccount := h.clientsByAccount[accountId]
	for _, userIdIt := range clientsOfAccount {
		client := h.clients[userIdIt]
		if client != nil {
			if userId == nil || *userId != userIdIt {
				client.sendMessage <- msg
			}
		}
	}
}

func (h *HandlerEvent) SendMessageToUser(msg eventMessage, userId int64) {
	h.m.Lock()
	defer h.m.Unlock()
	client := h.clients[userId]
	if client != nil {
		client.sendMessage <- msg
	}
}
