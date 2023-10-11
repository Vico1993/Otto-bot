package handles

import "github.com/Vico1993/Otto-bot/internal/service"

var ottoService = service.NewOttoService()
var chats map[string]*service.Chat = make(map[string]*service.Chat)
