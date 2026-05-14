package ra

import "time"

// PresenceTimeout is how stale a RichPresenceMsgDate can be before
// we consider the user no longer actively playing.
const PresenceTimeout = 130 * time.Second

// UserSummary is the relevant subset of the API_GetUserSummary response.
type UserSummary struct {
	User                string `json:"User"`
	LastGameID          int    `json:"LastGameID"`
	RichPresenceMsg     string `json:"RichPresenceMsg"`
	RichPresenceMsgDate string `json:"RichPresenceMsgDate"`
	Status              string `json:"Status"`
	UserPic             string `json:"UserPic"`
	TotalPoints         int    `json:"TotalPoints"`
	TotalSoftcorePoints int    `json:"TotalSoftcorePoints"`
	TotalTruePoints     int    `json:"TotalTruePoints"`
	Rank                int    `json:"Rank"`
	TotalRanked         int    `json:"TotalRanked"`
}

// AvatarURL returns the full URL for the user's profile picture.
func (u UserSummary) AvatarURL() string {
	if u.UserPic == "" {
		return ""
	}
	return "https://media.retroachievements.org" + u.UserPic
}

// IsActive returns true if the user has fresh rich presence data.
func (u UserSummary) IsActive() bool {
	if u.LastGameID == 0 || u.RichPresenceMsg == "" {
		return false
	}
	t, err := time.Parse("2006-01-02 15:04:05", u.RichPresenceMsgDate)
	if err != nil {
		return false
	}
	return time.Since(t.UTC()) < PresenceTimeout
}

// Game is the relevant subset of the API_GetGame response.
type Game struct {
	Title       string `json:"Title"`
	ConsoleID   int    `json:"ConsoleID"`
	ConsoleName string `json:"ConsoleName"`
	ImageIcon   string `json:"ImageIcon"`
}

// ArtURL returns the full URL for the game's cover art image.
func (g Game) ArtURL() string {
	if g.ImageIcon == "" {
		return ""
	}
	return "https://media.retroachievements.org" + g.ImageIcon
}

// consoleIconURLs maps RetroAchievements console IDs to their system icon URLs.
var consoleIconURLs = map[int]string{
	1:   "https://static.retroachievements.org/assets/images/system/md.png",
	2:   "https://static.retroachievements.org/assets/images/system/n64.png",
	3:   "https://static.retroachievements.org/assets/images/system/snes.png",
	4:   "https://static.retroachievements.org/assets/images/system/gb.png",
	5:   "https://static.retroachievements.org/assets/images/system/gba.png",
	6:   "https://static.retroachievements.org/assets/images/system/gbc.png",
	7:   "https://static.retroachievements.org/assets/images/system/nes.png",
	8:   "https://static.retroachievements.org/assets/images/system/pce.png",
	9:   "https://static.retroachievements.org/assets/images/system/scd.png",
	10:  "https://static.retroachievements.org/assets/images/system/32x.png",
	11:  "https://static.retroachievements.org/assets/images/system/sms.png",
	12:  "https://static.retroachievements.org/assets/images/system/ps1.png",
	13:  "https://static.retroachievements.org/assets/images/system/lynx.png",
	14:  "https://static.retroachievements.org/assets/images/system/ngp.png",
	15:  "https://static.retroachievements.org/assets/images/system/gg.png",
	16:  "https://static.retroachievements.org/assets/images/system/gc.png",
	17:  "https://static.retroachievements.org/assets/images/system/jag.png",
	18:  "https://static.retroachievements.org/assets/images/system/ds.png",
	19:  "https://static.retroachievements.org/assets/images/system/wii.png",
	21:  "https://static.retroachievements.org/assets/images/system/ps2.png",
	23:  "https://static.retroachievements.org/assets/images/system/mo2.png",
	24:  "https://static.retroachievements.org/assets/images/system/mini.png",
	25:  "https://static.retroachievements.org/assets/images/system/2600.png",
	27:  "https://static.retroachievements.org/assets/images/system/arc.png",
	28:  "https://static.retroachievements.org/assets/images/system/vb.png",
	29:  "https://static.retroachievements.org/assets/images/system/msx.png",
	33:  "https://static.retroachievements.org/assets/images/system/sg1k.png",
	37:  "https://static.retroachievements.org/assets/images/system/cpc.png",
	38:  "https://static.retroachievements.org/assets/images/system/a2.png",
	39:  "https://static.retroachievements.org/assets/images/system/sat.png",
	40:  "https://static.retroachievements.org/assets/images/system/dc.png",
	41:  "https://static.retroachievements.org/assets/images/system/psp.png",
	43:  "https://static.retroachievements.org/assets/images/system/3do.png",
	44:  "https://static.retroachievements.org/assets/images/system/cv.png",
	45:  "https://static.retroachievements.org/assets/images/system/intv.png",
	46:  "https://static.retroachievements.org/assets/images/system/vect.png",
	47:  "https://static.retroachievements.org/assets/images/system/8088.png",
	49:  "https://static.retroachievements.org/assets/images/system/pc-fx.png",
	51:  "https://static.retroachievements.org/assets/images/system/7800.png",
	53:  "https://static.retroachievements.org/assets/images/system/ws.png",
	56:  "https://static.retroachievements.org/assets/images/system/ngcd.png",
	57:  "https://static.retroachievements.org/assets/images/system/chf.png",
	63:  "https://static.retroachievements.org/assets/images/system/wsv.png",
	69:  "https://static.retroachievements.org/assets/images/system/duck.png",
	71:  "https://static.retroachievements.org/assets/images/system/ard.png",
	72:  "https://static.retroachievements.org/assets/images/system/wasm4.png",
	73:  "https://static.retroachievements.org/assets/images/system/a2001.png",
	74:  "https://static.retroachievements.org/assets/images/system/vc4000.png",
	75:  "https://static.retroachievements.org/assets/images/system/elek.png",
	76:  "https://static.retroachievements.org/assets/images/system/pccd.png",
	77:  "https://static.retroachievements.org/assets/images/system/jcd.png",
	78:  "https://static.retroachievements.org/assets/images/system/dsi.png",
	80:  "https://static.retroachievements.org/assets/images/system/uze.png",
	102: "https://static.retroachievements.org/assets/images/system/exe.png",
}

// ConsoleIconURL returns the icon URL for the game's console, or "" if unknown.
func (g Game) ConsoleIconURL() string {
	return consoleIconURLs[g.ConsoleID]
}

// UserProgress is the relevant subset of the API_GetUserProgress response.
type UserProgress struct {
	NumPossibleAchievements int `json:"NumPossibleAchievements"`
	NumAchieved             int `json:"NumAchieved"`
	NumAchievedHardcore     int `json:"NumAchievedHardcore"`
}
