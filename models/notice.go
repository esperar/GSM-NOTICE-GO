package models

type Notice struct {
	Id        int    `json:id`
	Title     string `json:title`
	Content   string `json:content`
	CreatedBy int    `json:user_id`
}
