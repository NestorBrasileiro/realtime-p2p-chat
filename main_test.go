package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

// Testa se o servidor WebSocket responde corretamente.
func TestWebSocketConnection(t *testing.T) {
	// Cria um servidor HTTP de teste
	server := httptest.NewServer(http.HandlerFunc(handleConnections))
	defer server.Close()

	// Constrói a URL do WebSocket
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	// Define o dialer para o cliente WebSocket
	dialer := websocket.DefaultDialer

	// Tenta se conectar ao WebSocket
	ws, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("WebSocket connection failed: %v", err)
	}
	defer ws.Close()

	// Teste de envio de mensagem
	testMessage := "Hello WebSocket!"
	err = ws.WriteMessage(websocket.TextMessage, []byte(testMessage))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// Recebe a mensagem do servidor
	_, receivedMessage, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	// Compara a mensagem enviada e recebida
	if string(receivedMessage) != testMessage {
		t.Fatalf("Expected message %s but got %s", testMessage, string(receivedMessage))
	}
}
