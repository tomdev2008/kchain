package types

type Image struct {
	Metadata
	Thumb    string `gorm:"size:255"`
	Original string `gorm:"size:255"`
	Ext      string `gorm:"unique" binding:"required" `
	Width    uint   `gorm:"default:0"`
	Height   uint   `gorm:"default:0"`
	Size     uint   `gorm:"default:0"`
}
