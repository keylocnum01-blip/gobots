package handler

import (
    "fmt"
    "strconv"
    "time"
	"../botstate"
	"../config"
	"../library/linetcr"
	"../utils"
    "../library/linetcr/SyncService"
)

func LogAccess(client *linetcr.Account, group, from, tipe string, targets []string, tempat int32) {
	defer utils.PanicHandle("logAccess")
	if !botstate.LogMode || SendMyseller(from) {
		return
	}
	nm, _, _ := client.GetChatList(group)
	var ts = ""
	if tipe == "ban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! banned %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! banned %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Unbanned %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Unbanned %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "kick" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! kicked %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! kicked %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "mkick" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! kicked %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! kicked %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "inv" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! invited %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! invited %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "cancel" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! canceled %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! canceled %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "owner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Owner:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Owner from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "master" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Master:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Master from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "admin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Admin:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Admin from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "bot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Bot:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Bot from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "gowner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Gowner:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Gowner from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "gadmin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Promoted %v user's as Gadmin:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Promoted %v user's as Gadmin from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "fuck" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Added %v user's to FuckList:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Added %v user's to FuckList from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "gban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Added %v user's to GBanList:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Added %v user's to GBanList from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "mute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! Added %v user's to MuteList:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! Added %v user's to MuteList from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	}
	if botstate.LogGroup != "" {
		if len(targets) == 1 {
			client.SendMention(botstate.LogGroup, ts, targets)
		} else {
			client.SendMessage(botstate.LogGroup, ts)
		}
	}
}

func LogLast(op *SyncService.Operation, midds string) {
	defer linetcr.PanicOnly()
	if op.Type == 26 {
		if op.Message.ContentType == 18 {
			return
		}
	}
	botstate.LastActive.Set(midds, op)
}

func LogFight(room *linetcr.LineRoom) {
	defer utils.PanicHandle("logfight")
	if botstate.LogMode {
		var tx = ""
		for i := 0; i < len(botstate.ClientBot); i++ {
			exe := botstate.ClientBot[i]
			if !exe.Frez {
				g, botstate.Err := exe.GetGroupMember(room.Id)
				if botstate.Err != nil {
					continue
				} else {
					room.Name = g
					break
				}
			}
		}

		tx += fmt.Sprintf("Squad action's in Group:\n%s\n", room.Name)
		if room.Kick != 0 {
			tx += fmt.Sprintf("\nKick's: %v", room.Kick)
		}
		if room.Invite != 0 {
			tx += fmt.Sprintf("\nInvite's: %v", room.Invite)
		}
		if room.Cancel != 0 {
			tx += fmt.Sprintf("\nCancel's: %v", room.Cancel)
		}
		if room.Kick == 0 && room.Invite == 0 && room.Cancel == 0 {
			room.Kick = 0
			room.Invite = 0
			room.Cancel = 0
			return
		}
		roomLog := linetcr.GetRoom(botstate.LogGroup)
		if len(roomLog.Client) != 0 {
			exe, botstate.Err := botstate.SelectBot(roomLog.Client[0], botstate.LogGroup)
			if botstate.Err == nil {
				if exe != nil {
					exe.SendMessage(botstate.LogGroup, botstate.Fancy(tx))
				}
			} else {
				botstate.LogMode = false
				botstate.LogGroup = ""
			}
		}
	}
	room.Kick = 0
	room.Invite = 0
	room.Cancel = 0
}

func GenerateTimeLog(client *linetcr.Account, to string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a := time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0" + strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0" + strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0" + strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := "Date : " + dd + "-" + MM + "-" + yyyy + "\nTime : " + hhconv + ":" + mmconv + ":" + ssconv
	client.SendMessage(to, botstate.Fancy(times))
}

func AppendLastSticker(s []*config.Stickers, e *config.Stickers) []*config.Stickers {
	defer linetcr.PanicOnly()
	s = config.RemoveSticker(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func CheckMessage(waktu int64, typ int8) bool {
	if typ == 1 {
		for _, wkt := range botstate.TimeSend {
			if wkt == waktu {
				return false
				break
			}
		}
		botstate.TimeSend = append(botstate.TimeSend, waktu)
		return true
	}
	return false
}

func AppendLastD(s [][]string, e []string) [][]string {
	defer linetcr.PanicOnly()
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func AppendLast(s []string, e string) []string {
	defer linetcr.PanicOnly()
	s = utils.RemoveString(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func LlistCheck(client *linetcr.Account, to string, typec string, nCount int, sender string, rplay string, mentionlist []string) (ss []string) {
	saodd := []string{}
	pendlast := []string{}
	if len(mentionlist) != 0 {
		for a := range mentionlist {
			if !utils.InArrayString(saodd, mentionlist[a]) && !utils.InArrayString(botstate.Squadlist, mentionlist[a]) {
				saodd = append(saodd, mentionlist[a])
			}

		}
		return saodd
	} else if rplay != "" {
		if !utils.InArrayString(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "rplay" {
		if !utils.InArrayString(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "lmid" {
		g, ok := botstate.Lastmid.Get(to)
		if !ok {
			g = [][]string{}
			botstate.Lastmid.Set(to, g)
		} else {
			num := nCount
			c := g.([][]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i]...)
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lmessage" {
		g, ok := botstate.Lastmessage.Get(to)
		if !ok {
			g = [][]string{}
			botstate.Lastmessage.Set(to, g)
		} else {
			num := nCount
			c := g.([][]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i]...)
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "linvite" {
		g, ok := botstate.Lastinvite.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastinvite.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lkick" {
		g, ok := botstate.Lastkick.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastkick.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcancel" {
		g, ok := botstate.Lastcancel.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastcancel.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lqr" {
		g, ok := botstate.Lastupdate.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastupdate.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ljoin" {
		g, ok := botstate.Lastjoin.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastjoin.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ltag" {
		g, ok := botstate.Lasttag.Get(to)
		if !ok {
			g = []string{}
			botstate.Lasttag.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcon" {
		g, ok := botstate.Lastcon.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastcon.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lleave" {
		g, ok := botstate.Lastleave.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastleave.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "@me" {
		if !utils.InArrayString(saodd, sender) {
			saodd = append(saodd, sender)
		}
		return saodd
	} else if typec == "@all" {
		Member := client.GetChatListMem(to)
		for i := 0; i < len(Member); i++ {
			if !utils.InArrayString(saodd, Member[i]) {
				saodd = append(saodd, Member[i])
			}
		}
		return saodd
	} else if typec == "lbanlist" {
		for i := 0; i < len(botstate.Banned.Banlist); i++ {
			if !utils.InArrayString(saodd, botstate.Banned.Banlist[i]) {
				saodd = append(saodd, botstate.Banned.Banlist[i])
			}
		}
		return saodd
	} else if typec == "pend" {
		tcr := strings.Replace(typec, "pend ", "", 1)
		numb, _ := strconv.Atoi(tcr)
		_, _, pind := client.GetChatList(to)
		if numb == 0 {
			for n, i := range pind {
				if n < botstate.CancelPend {
					if !utils.InArrayString(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	} else if typec == "numpend" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !utils.InArrayString(saodd, i) {
				pendlast = append(pendlast, i)
			}
			num := nCount
			lk := len(pendlast)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, pendlast[i])
					}
					no++
				}
				return saodd
			}
		}
		return saodd
	} else if typec == "pendingall" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !utils.InArrayString(saodd, i) {
				saodd = append(saodd, i)
			}
		}
		return saodd
	} else if typec == "@oa" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
		      contact, _ := client.GetContact(i)
			if contact != nil {
				if contact.CapableBuddy {
					if !utils.InArrayString(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	}
	return saodd
}

func Savejoin(Pelaku string, Optime int64) {
	defer linetcr.PanicOnly()
	ix := config.IndexOf(botstate.Detectjoin.User, Pelaku)
	if ix == -1 {
		botstate.Detectjoin.User = append(botstate.Detectjoin.User, Pelaku)
		botstate.Detectjoin.Time = append(botstate.Detectjoin.Time, Optime)
	} else {
		botstate.Detectjoin.Time[ix] = Optime
	}
}


func LogOp(op *SyncService.Operation, client *linetcr.Account) {
	defer linetcr.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	if tipe == 124 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 133 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 130 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 122 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 55 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 128 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 26 {
		msg := op.Message
		if utils.InArrayString(botstate.Squadlist, msg.From_) {
			return
		}
		LogLast(op, msg.From_)
	}
}

func NotifBot(client *linetcr.Account, to string, tipe string, targets string) {
       names1, _ := client.GetGroupMember(to)
       loc, _ := time.LoadLocation("Asia/Jakarta")
	a := time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0" + strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0" + strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0" + strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	Date := dd + "-" + MM + "-" + yyyy
	Time := hhconv + ":" + mmconv + ":" + ssconv
	tcr, _ := client.GetContact(targets)
	if tipe == "invite" {
	       list := fmt.Sprintf("NOTIFI INVITE:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "kick" {
	       list := fmt.Sprintf("NOTIFI KICK:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "cancel" {
	       list := fmt.Sprintf("NOTIFI CANCEL:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "delete" {
	       list := fmt.Sprintf("NOTIFI DELETE ACCOUNT:\n  . Mid : %s", tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "call" {
	       list := fmt.Sprintf("NOTIFI CALL ME:\n  . Mid : %s", tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "join" {
	       list := fmt.Sprintf("NOTIFI MEMBER JOIN:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	} else if tipe == "leave" {
	       list := fmt.Sprintf("NOTIFI MEMBER LEAVE:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(botstate.ClientBot) != 0 {
		      botstate.ClientBot[0].SendNewText(botstate.LogGroup, list)
	     }
	}
}