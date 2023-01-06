package model

import (
	"fmt"
)

type Bible struct {
	ID         uint    `gorm:"primary_key;Column:ID"`
	VolumeSN   int     `gorm:"primary_key;Column:VolumeSN"`
	ChapterSN  int     `gorm:"Column:ChapterSN"`
	VerseSN    int     `gorm:"Column:VerseSN"`
	Lection    string  `gorm:"Column:Lection"`
	SoundBegin float32 `gorm:"Column:SoundBegin"`
	SoundEnd   float32 `gorm:"Column:SoundEnd"`
}

func (bible *Bible) TableName() string {
	return "Bible"
}

func (bible *Bible) String() string {

	return fmt.Sprintf("%d章%d页%d节:%s", bible.VolumeSN, bible.ChapterSN, bible.VerseSN, bible.Lection)
}
