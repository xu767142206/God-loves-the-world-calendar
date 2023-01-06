package utils

import (
	"encoding/json"
	"fmt"
	"github.com/liujiawm/gocalendar"
	"sync"
	"time"
)

const TIME_FORMAT = "2006年01月02日15点04分05秒"

type SolarTermItem gocalendar.SolarTermItem

func (st *SolarTermItem) String() string {
	return fmt.Sprintf("%s:%s", st.Name, st.Time.Format(TIME_FORMAT))
}

type Calendar struct {
	Gocalendar *gocalendar.Calendar //日历
	TimeFormat string
}

// Result 结果
type Result struct {
	Date              string   `json:"date"`
	SolarTermItem     string   `json:"st"`                 // 节气
	IsAccidental      int      `json:"isam"`               // 0为本月日期,-1为上一个月的日期,1为下一个月的日期,
	Festival          []string `json:"festival"`           // 公历节日
	FestivalSecondary []string `json:"festival_secondary"` // 公历节日
	GZ                string   `json:"gz"`                 // 干支
	LunarDate         string   `json:"ld"`                 // 农历
	StarSign          string   `json:"ss"`                 // 星座
}

var calendar *Calendar
var once sync.Once

func GetCalendar() *Calendar {

	once.Do(func() {
		c := gocalendar.NewCalendar(gocalendar.CalendarConfig{
			Grid:            gocalendar.GridDay,
			FirstWeek:       0,
			SolarTerms:      true,
			Lunar:           true,
			HeavenlyEarthly: true,
			NightZiHour:     true,
			StarSign:        true,
			TimeZoneName:    "Asia/Shanghai",
		})
		calendar = &Calendar{
			Gocalendar: c,
			TimeFormat: TIME_FORMAT,
		}
	})

	return calendar
}

// Get 获取指定时间的日历数据
func (calendar Calendar) Get(nowTime time.Time) (*Result, error) {

	result := GetCalendar().Gocalendar.
		GenerateWithDate(nowTime.Year(), int(nowTime.Month()), nowTime.Day())
	if len(result) == 0 {
		return nil, fmt.Errorf("获取日历失败")
	}

	toDay := result[0]

	solarTermItem := (*SolarTermItem)(toDay.SolarTerm)

	return &Result{
		Date:              toDay.Time.Format("2006年01月02日"),
		StarSign:          toDay.StarSign.Name,
		Festival:          toDay.Festival.Show,
		FestivalSecondary: toDay.Festival.Secondary,
		SolarTermItem:     fmt.Sprint(solarTermItem),
		GZ:                fmt.Sprint(toDay.GZ),
		LunarDate:         fmt.Sprint(toDay.LunarDate),
		IsAccidental:      toDay.IsAccidental,
	}, nil
}

// GetTimeNow 获取当天的日历数据
func (calendar Calendar) GetTimeNow() (*Result, error) {
	return calendar.Get(time.Now())
}

func (result *Result) Json() []byte {
	marshal, _ := json.Marshal(result)
	return marshal
}

func (result *Result) String() string {

	return fmt.Sprintf("[\n \t日期:%s \n \t星座:%s\n \t公假:%s\n \t其他假:%s\n \t节气:%s\n \t干支:%s\n \t农历:%s \n]",
		result.Date,
		result.StarSign,
		result.Festival,
		result.FestivalSecondary,
		result.SolarTermItem,
		result.GZ,
		result.LunarDate,
	)
}
