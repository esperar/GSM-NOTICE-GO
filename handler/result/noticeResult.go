package result

type NoticeResult struct {
	Id        int    `json:id`
	Title     string `json:title`
	CreatedBy int    `json:user_id`
}
