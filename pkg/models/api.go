package models

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	StatusOK      = StatusModel{Status: http.StatusOK, Message: "Ok"}
	StatusCreated = StatusModel{Status: http.StatusCreated, Message: "Created"}
)

type SortOrder string

const (
	SortOrderName    = SortOrder("name")
	SortOrderCreated = SortOrder("created")
)

type StatusModel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AuthLoginResponse struct {
	AccessToken string    `json:"access_token"`
	Deadline    time.Time `json:"deadline"`
}

type CreateSoundRequest struct {
	Sound

	UploadId   string `json:"upload_id"`
	YouTubeURL string `json:"youtube_url"`

	Normalize bool `json:"normalize"`
	Overdrive bool `json:"overdrive"`
}

type UpdateSoundRequest struct {
	Sound
}

type SoundUploadResponse struct {
	UploadId string    `json:"upload_id"`
	Deadline time.Time `json:"deadline"`
}

type SetVolumeRequest struct {
	Volume int `json:"volume"`
}

type FastTrigger struct {
	FastTrigger string `json:"fast_trigger"`
}

type PlaybackStats struct {
	Ident string `json:"ident"`
	Count int    `json:"count"`
}

type StateStats struct {
	NSoudns int `json:"n_sounds"`
	NPlays  int `json:"n_plays"`
}

type OTAResponse struct {
	Deadline   time.Time `json:"deadline"`
	Token      string    `json:"token"`
	QRCodeData string    `json:"qrcode_data"`
}

type User struct {
	discordgo.User

	AvatarURL string `json:"avatar_url"`
	IsOwner   bool   `json:"is_owner"`
}

func UserFromUser(u discordgo.User) User {
	return User{
		User:      u,
		AvatarURL: u.AvatarURL(""),
	}
}

type ApiKey struct {
	ApiKey string `json:"api_key"`
}
