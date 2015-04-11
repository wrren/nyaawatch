package main

import "net/smtp"
import "strconv"
import "log"

func Notify( episode Series, config NotificationConfig ) {
	auth := smtp.PlainAuth( "", config.From, config.Password, config.Server )
	to := []string{ config.To }
	msg := []byte( "To: " + config.To + "\r\nSubject: NyaaWatch Downloaded " + episode.Name + " Episode " + strconv.Itoa( episode.Episode ) + "\r\n\r\n" )
	err := smtp.SendMail( config.Server + ":" + strconv.Itoa( config.Port ), auth, config.From, to, msg ) 
   	
   	if err != nil {
      		log.Fatal( err )
   	}
}