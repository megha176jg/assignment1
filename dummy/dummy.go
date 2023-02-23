package dummy

import (
	"video_api/models"
)

var Videos []models.Video

func DummyData() {
	Videos = []models.Video{
		{ID: "abc123", Title: "Stay:Intersteller", View: 2},
		{ID: "xyz123", Title: "Pastlives:Sapiens", View: 5},
		{ID: "pqr123", Title: "Save Your Tears:Weeknd", View: 6},
	}
}

func GetVideos() []models.Video {
	return Videos
}
