# Realtime P2P Chat  (WORK IN PROGRESS)

Um sistema de chat ponto a ponto (P2P) em tempo real, que permite a comunicação direta entre usuários sem a necessidade de um servidor centralizado.  

Disclaimer: O frontend ainda não está desenvolvido e o back-end se encontra em estágio inicial, pois estou usando o projeto para tirar dúvidas práticas sobre a linguágem

## Funcionalidades  

- 📡 **Comunicação em tempo real**: Envio e recebimento instantâneo de mensagens.  
- 🔒 **Conexão P2P segura**: Mensagens são transmitidas diretamente entre os usuários.  
- 💬 **Interface simples**: UI intuitiva para uma experiência de uso otimizada.  
- 📱 **Compatibilidade multiplataforma**: Funciona em diferentes dispositivos e navegadores.  

## Tecnologias Utilizadas  

- **Frontend**: (WIP) Será usado Vue3 ou React 
- **Backend**: Golang
- **Comunicação em tempo real**: WebRTC ou WebSockets  
- **Outras bibliotecas ou frameworks**: (WIP) Socket.io 

## Instalação e Execução  

### Pré-requisitos  

- [Instalar Go](https://golang.org/dl/) (versão 1.XX ou superior)  

### Passos  

1. Clone o repositório:  
   ```bash
   git clone https://github.com/NestorBrasileiro/realtime-p2p-chat.git
   cd realtime-p2p-chat
   
2. Instale as dependências:
   ```bash
   go mod tidy

3. Execute o projeto:
   ```bash
   go run cmd/main.go


