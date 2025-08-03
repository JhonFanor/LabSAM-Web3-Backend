package services

import (
	"fmt"
	"lamsam-web3-backend/config"

	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendVerificationEmail(email, token string) error
	SendPasswordResetEmail(email, token string) error
	Send(to, subject, body string) error
}

type emailService struct {
	gmail  *config.GmailConfig
	server *config.ServerConfig
}

func NewEmailService(gmail *config.GmailConfig, server *config.ServerConfig) EmailService {
	return &emailService{
		gmail:  gmail,
		server: server,
	}
}

func (s *emailService) SendVerificationEmail(email, token string) error {
	link := fmt.Sprintf(s.server.SERVER+"/verify-email?token=%s", token)
	body := fmt.Sprintf("Bienvenido. Verifica tu correo haciendo clic en este enlace: <a href=\"%s\">Verificar correo</a>", link)
	return s.Send(email, "Verifica tu cuenta", body)
}

func (s *emailService) SendPasswordResetEmail(email, token string) error {
	resetLink := fmt.Sprintf(s.server.FRONTEND+"reset-password?token=%s", token)
	body := fmt.Sprintf("Haz clic en el siguiente enlace para restablecer tu contraseña: <a href=\"%s\">Restablecer contraseña</a>", resetLink)
	return s.Send(email, "Recuperación de contraseña", body)
}

func (s *emailService) Send(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.gmail.GMAIL_USER)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, s.gmail.GMAIL_USER, s.gmail.GMAIL_PASS)
	return d.DialAndSend(m)
}
