package handler

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"../botstate"
	"../library/SyncService"
	"../library/linetcr"
	"../utils"
)

type Mentions struct {
	MENTIONEES []struct {
		M   string `json:"M"`
		S   string `json:"S"`
		E   string `json:"E"`
		Mid string `json:"M"`
	} `json:"MENTIONEES"`
}

func Bot(op *SyncService.Operation, client *linetcr.Account, ch chan int) {
	defer utils.PanicHandle("Bot")
	op.Message = client.DecryptE2EEMessage(op.Message)
	msg := op.Message
	if msg.ToType != 2 {
		return
	}
	if _, ok := botstate.Commandss.Get(op.CreatedTime); ok {
		return
	} else {
		botstate.Commandss.Set(op.CreatedTime, client)
	}
	if time.Now().Sub(botstate.Timeabort) >= 60*time.Second {
		Abort()
	}
	if botstate.AutoBackBot {
		time.AfterFunc(time.Duration(botstate.Timebk)*time.Second, func() {
			all := []string{}
			GetSquad(client, op.Message.To)
			room := linetcr.GetRoom(op.Message.To)
			cuk := room.Client
			alls := []*linetcr.Account{}
			for _, x := range botstate.ClientBot {
				if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
					alls = append(alls, x)
				}
			}
			for _, x := range botstate.ClientBot {
				if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
					all = append(all, x.MID)
				} else {
					break
				}
			}
			ClAct := len(alls)
			if ClAct != 0 {
				sort.Slice(alls, func(i, j int) bool {
					return alls[i].KickPoint < alls[j].KickPoint
				})
				for _, a := range all {
					alls[0].InviteIntoGroupNormal(op.Message.To, []string{a})
					time.Sleep(100 * time.Millisecond)
					GetSquad(client, op.Message.To)
				}
			}
		})
	}
	// Rname and Sname are not used in the snippet but declared in main.go
	// Rname := botstate.MsRname
	// Sname := botstate.MsSname
	sender := op.Message.From_
	text := op.Message.Text
	receiver := op.Message.To
	var pesan = strings.ToLower(text)
	var to string
	mentions := Mentions{}
	if op.Message.ToType == 0 {
		to = sender
	} else {
		to = receiver
	}
	if botstate.DetectCall && msg.ToType == 2 {
		Room := linetcr.GetRoom(to)
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
		pr, _ := client.GetContact(sender)
		if msg.ContentMetadata["GC_MEDIA_TYPE"] == "AUDIO" || msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
			a := "「 DETECT 」"
			a += "\n › Type: Callgroup"
			a += "\n    • Type: AUDIO"
			a += "\n    • Date: " + Date
			a += "\n    • Time: " + Time
			a += fmt.Sprintf("\n    • Group: %s", Room.Name)
			a += "\n    • Host: " + pr.DisplayName
			client.SendMessage(to, botstate.Fancy(a))
		} else {
			if msg.ContentMetadata["GC_MEDIA_TYPE"] == "VIDEO" || msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
				a := "「 DETECT 」"
				a += "\n › Type: Callgroup"
				a += "\n    • Type: VIDEO"
				a += "\n    • Date: " + Date
				a += "\n    • Time: " + Time
				a += fmt.Sprintf("\n    • Group: %s", Room.Name)
				a += "\n    • Host: " + pr.DisplayName
				client.SendMessage(to, botstate.Fancy(a))
			}
		}
	}
	if len(botstate.Sinderremote) != 0 {
		if utils.InArrayString(botstate.Sinderremote, sender) {
			if botstate.Remotegrupid != "" {
				to = botstate.Remotegrupid
			}
		}
	}
	mentionlist := []string{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["MENTION"]), &mentions)
	for _, mention := range mentions.MENTIONEES {
		if !utils.InArrayString(mentionlist, mention.Mid) {
			mentionlist = append(mentionlist, mention.Mid)
		}
	}
	var Rplay = ""
	var room *linetcr.LineRoom
	var bks = []*linetcr.Account{}
	room = linetcr.GetRoom(to)
	bks = room.Client
	if len(bks) == 0 {
		GetSquad(client, to)
		room = linetcr.GetRoom(to)
		bks = room.Client
	}
	sort.Slice(room.Ava, func(i, j int) bool {
		return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
	})
	
	// Continuation of Bot function logic would go here...
	// Since I cannot read the entire function at once, I will just create the file with what I have and then append the rest.
	// But Bot function is huge. I should probably use a script to migrate it entirely.
}
