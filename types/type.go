package types

import (
	"github.com/Iffigues/mywebsite/linuxtool"
	"github.com/Iffigues/mywebsite/config"
	//"polaroid/pk"
	"github.com/gorilla/sessions"
)

type Data struct {
	Store *sessions.CookieStore
	Conf  config.Config
	Commande  *linuxtool.Commande
	//Db    *pk.Pk
}
