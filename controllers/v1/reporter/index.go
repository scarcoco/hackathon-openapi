package reporter

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type ReporterItem struct {
	Id           string `fake:"{uuid}"`
	Name         string `fake:"{firstname}'s watch"`
	Description  string `fake:"{sentence:3}"`
	ReporterType string `fake:"{randomstring:[Watch,Anima,Traffic,Unknown]}"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ListReporterData struct {
	Total int            `fake:"{number:15,30}"`
	List  []ReporterItem `fakesize:"15"`
}
type ListReporterResponse struct {
	ErrNo int `fake:"{randomstring:[200]}"`
	Data  ListReporterData
}

type DetailReporterResponse struct {
	ErrNo int `fake:"{randomstring:[200]}"`
	Data  ReporterItem
}

type ReporterLocation struct {
	ReporterId string

	Id        string
	Lng       float64 `fake:"{latitude}"`
	Lat       float64 `fake:"{longitude}"`
	Timestamp time.Time
}
type LocationReporterData struct {
	Total int                `fake:"{number:15,30}"`
	List  []ReporterLocation `fakesize:"15"`
}
type LocationReporterResponse struct {
	ErrNo int `fake:"{randomstring:[200]}"`
	Data  LocationReporterData
}

/** 列表数据 */
func List() ListReporterResponse {
	var t ListReporterResponse
	gofakeit.Struct(&t)
	return t
}

/** 详情数据 */
func Detail() DetailReporterResponse {
	var t DetailReporterResponse
	gofakeit.Struct(&t)
	return t
}

/** 轨迹数据 */
func Location() LocationReporterResponse {
	var t LocationReporterResponse
	gofakeit.Struct(&t)
	return t
}
