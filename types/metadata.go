package types

type Metadata struct {
	Type        string
	Title       string `gorm:"type:text" binding:"required"`
	Language    string
	Category    string
	Created     float32
	Abstract    string `gorm:"type:text"`
	ContentHash string `gorm:"size:255"`
	DNA         string `gorm:"size:255;unique_index"`
	License     string `gorm:"type:text" binding:"required"`
	Signature   string `sql:"-" binding:"required"`
	Extra       string `gorm:"type:text" binding:"required"`
	BlockHash   string `gorm:"size:255"`
}


