package data

import "time"
"""
	startsAt := strings.Replace(strings.Split(v[i].(map[string]interface{})["startsAt"].(string), ".")[0], "T", " ", -1)

	duration := strconv.FormatInt(getHourDiffer(startsAt, time.Now().Format("2006-01-02 15:04:05")), 10)
"""


func getHourDiffer(start_time, end_time string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 60
		//minute = diff / 60
		return hour
	} else {
		return hour
	}

}



// func TimeDiff(oldTime, newTime string) int {
// 	//传入的时间，必须是规定的格式 "2006-01-02 15:04:05"
// 	//用于计算2个时间之间有多少秒的差别

// 	old_time, _ := time.Parse("2006-01-02 15:04:05", oldTime)
// 	new_time, _ := time.Parse("2006-01-02 15:04:05", newTime)

// 	if old_time.Unix() >= new_time.Unix() {
// 		subTimes := old_time.Sub(new_time).Seconds()
// 		return int(subTimes)
// 	} else {
// 		subTimes := new_time.Sub(old_time).Seconds()
// 		return int(subTimes)
// 	}
// }