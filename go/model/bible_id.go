package model

import "fmt"

type BibleID struct {
	SN            int    `gorm:"primary_key;Column:SN"`
	KindSN        int    `gorm:"Column:KindSN"`
	ChapterNumber int    `gorm:"Column:ChapterNumber"`
	NewOrOld      bool   `gorm:"Column:NewOrOld"`
	PinYin        string `gorm:"Column:PinYin"`
	ShortName     string `gorm:"Column:ShortName"`
	FullName      string `gorm:"Column:FullName"`
}

var NewOrOldMap = map[bool]string{
	true:  "新约",
	false: "旧约",
}

func (bibleID *BibleID) TableName() string {
	return "BibleID"
}

func (bibleID *BibleID) String() string {

	return fmt.Sprintf("%s --------------------------- %d章-%d页-%s", bibleID.FullName, bibleID.KindSN, bibleID.ChapterNumber, NewOrOldMap[bibleID.NewOrOld])
}
