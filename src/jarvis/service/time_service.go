package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"jarvis/data"
	"jarvis/log"
	"jarvis/util"
	"math"
	"strings"
	"time"
)

type Time struct{}

func (t Time) DurationToString(d time.Duration) string {
	s := ""
	hours := math.Floor(d.Hours())
	if hours > 0 {
		s += fmt.Sprintf("%v hour", hours)
	}
	if hours > 1 {
		s += "s"
	}
	minutes := int(d.Minutes()) % 60
	if minutes > 0 {
		if hours > 0 {
			s += " "
		}
		s += fmt.Sprintf("%v minute", minutes)
		if minutes > 1 {
			s += "s"
		}
	}
	seconds := int(d.Seconds()) % 60
	if seconds > 0 {
		if hours > 0 || minutes > 0 {
			s += " "
		}
		s += fmt.Sprintf("%v second", seconds)
		if seconds > 1 {
			s += "s"
		}
	}
	return s
}

func (t Time) StringToDuration(durStr string) (time.Duration, error) {
	// Run the duration string through a bunch of processing to get it into a time.Duration format that can be parsed by go
	durStr = strings.Replace(durStr, " seconds", "s", -1)
	durStr = strings.Replace(durStr, " second", "s", -1)
	durStr = strings.Replace(durStr, " minutes", "m", -1)
	durStr = strings.Replace(durStr, " minute", "m", -1)
	durStr = strings.Replace(durStr, " hours", "h", -1)
	durStr = strings.Replace(durStr, " hour", "h", -1)
	durStr = strings.Replace(durStr, " ", "", -1)
	d, err := time.ParseDuration(durStr)
	if err != nil {
		return d, errors.New("Apologies, but I can't seem to parse your duration string.")
	}
	if d.Hours() < 0 || d.Minutes() < 0 || d.Seconds() < 0 {
		return d, errors.New("Apologies, but my functionality does not include the recognition of negative time.")
	}
	return d, nil
}

// Converts a time object into a printable string
// Ignores date properties on the time and just looks at the hour and minute
func (tt Time) TimeToStringWithoutDate(t time.Time) string {
	result := ""
	hour := t.Hour() % 12
	result += fmt.Sprintf("%v:%2v", hour, t.Minute())
	if t.Hour() > 12 {
		result += "pm"
	} else {
		result += "am"
	}
	return result
}

func (tt Time) StringToTime(ts string, user string) (time.Time, error) {

	// Get the user's prefered timezone
	in, datum := data.GetDatum("my timezone", user)
	if !in {
		return time.Now(), errors.New("I don't seem to have your prefered timezone stored.")
	}
	userTz, err := time.LoadLocation(datum)
	if err != nil {
		return time.Now(), errors.New("The timezone I have on file for you appears to be invalid.")
	}

	// Get the current time and then convert it into the user's timezone
	currentTime := time.Now()
	currentTime = currentTime.In(userTz)

	// Convert the user's time into the requested time by dep-injecting the
	// current time down into this computational process
	t, err := tt.inTz(ts, currentTime)
	if err != nil {
		return t, err
	}

	return t, nil

}

func (tt Time) inTz(ts string, currentTime time.Time) (time.Time, error) {
	defaultErr := errors.New("Apologies, but I can't seem to read the time you gave me.")

	// Case 1: "Remind me at 8 to do X"
	// NOW will parse this to always mean "8AM of the current day", but I want this to actually mean
	//  - "8PM Today if it is after 8am today but before 8pm today"
	//  - "8AM Tomorrow if it is after 8AM today and after 8pm today"
	works, t, err := tt.parseAbsTimeLoneNumber(ts, currentTime, defaultErr)
	if err != nil {
		return t, err
	}
	if works {
		return t, nil
	}

	// Case 2: "Remind me at 8pm to do X"
	// NOW cannot parse this, so we help it out a little bit
	works, t, err = tt.parseAbsTimeWithCycle(ts, currentTime, defaultErr)
	if err != nil {
		return t, err
	}
	if works {
		return t, nil
	}

	// Case 3: "Remind me at 8am tomorrow to do x"
	// NOW cannot parse this
	works, t, err = tt.parseAbsTimeWithTomorrow(ts, currentTime, defaultErr)
	if err != nil {
		return t, err
	}
	if works {
		return t, nil
	}

	// At this point we pass the timestamp over to NOW to parse
	t, err = now.Parse(ts)
	if err != nil {
		return t, defaultErr
	}
	return t, nil

}

func (tt Time) parseAbsTimeLoneNumber(ts string, currentTime time.Time, defaultErr error) (bool, time.Time, error) {
	if util.NewRegex("^[0-9]{1,2}$").Matches(ts) || util.NewRegex("^[0-9]{1,2}:[0-9]{2}$").Matches(ts) {
		log.Trace("Parsing absolute time assuming lone number")
		t, err := now.Parse(ts)
		if err != nil {
			log.Trace("Error: %v\n", err.Error())
			return false, t, defaultErr
		}
		if t.Hour() > 12 {
			if t.Before(currentTime) {
				t = t.Add(24 * time.Hour)
			}
		} else {
			if t.Before(currentTime) {
				t = t.Add(12 * time.Hour)
			}
			if t.Before(currentTime) {
				t = t.Add(12 * time.Hour)
			}
		}
		return true, t, nil
	} else {
		return false, currentTime, nil
	}
}

func (tt Time) parseAbsTimeWithCycle(ts string, currentTime time.Time, defaultErr error) (bool, time.Time, error) {
	r := util.NewRegex("^[0-9]{1,2}:?([0-9]{2})?(am|pm|AM|PM)$")
	if r.Matches(ts) {
		log.Trace("Parsing absolute time assuming cyclic number")
		withoutCycle := strings.Replace(ts, "am", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "AM", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "pm", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "PM", "", -1)
		t, err := now.Parse(withoutCycle)
		if err != nil {
			log.Trace("Error: %v\n", err.Error())
			return false, t, defaultErr
		}
		if t.Hour() > 12 {
			log.Trace("Time provided is military with cycle; invalid")
			return false, t, defaultErr
		}
		if strings.Contains(ts, "AM") || strings.Contains(ts, "am") {
			if t.Before(currentTime) {
				t = t.Add(24 * time.Hour)
			}
		}
		if strings.Contains(ts, "PM") || strings.Contains(ts, "pm") {
			t = t.Add(12 * time.Hour)
			if t.Before(currentTime) {
				t = t.Add(24 * time.Hour)
			}
		}
		return true, t, nil
	} else {
		return false, currentTime, nil
	}
}

func (tt Time) parseAbsTimeWithTomorrow(ts string, currentTime time.Time, defaultErr error) (bool, time.Time, error) {
	r := util.NewRegex("^[0-9]{1,2}:?([0-9]{2})?(am|pm|AM|PM) [Tt]omorrow$")
	if r.Matches(ts) {
		log.Trace("Parsing absolute time assuming cycling time with tomorrow")
		withoutCycle := strings.Replace(ts, " Tomorrow", "", -1)
		withoutCycle = strings.Replace(withoutCycle, " tomorrow", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "am", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "AM", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "pm", "", -1)
		withoutCycle = strings.Replace(withoutCycle, "PM", "", -1)
		t, err := now.Parse(withoutCycle)
		if err != nil {
			log.Trace("Error: %v\n", err.Error())
			return false, t, defaultErr
		}
		if t.Hour() > 12 {
			log.Trace("Time provided is military with cycle; invalid")
			return false, t, defaultErr
		}
		if strings.Contains(ts, "AM") || strings.Contains(ts, "am") {
			t = t.Add(24 * time.Hour)
		}
		if strings.Contains(ts, "PM") || strings.Contains(ts, "pm") {
			t = t.Add(36 * time.Hour)
		}
		return true, t, nil
	} else {
		return false, currentTime, nil
	}
}
