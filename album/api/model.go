package api


type Album struct {
	UserID int    `gorm:"column:userId" json:"userId"`
	ID     int    `gorm:"primary_key;column:id" json:"id"`
	Title  string `gorm:"column:title" json:"title"`
}

type Photo struct {
	AlbumID      int    `gorm:"column:albumId" json:"albumId"`
	ID           int    `gorm:"primary_key;column:id" json:"id"`
	Title        string `gorm:"column:title" json:"title"`
	URL          string `gorm:"column:url" json:"url"`
	ThumbnailURL string `gorm:"column:thumbnailUrl" json:"thumbnailUrl"`
}
