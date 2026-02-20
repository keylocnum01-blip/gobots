﻿﻿﻿﻿﻿﻿﻿﻿﻿﻿﻿package handler

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"../botstate"
	"../library/SyncService"
	"../library/linetcr"
	"../utils"
)

func Fancy(text string) string {
	return text
}

func MentionList(op *SyncService.Operation) []string {
	msg := op.Message
	str := fmt.Sprintf("%v", msg.ContentMetadata["MENTION"])
	taglist := utils.GetMentionData(str)

	return taglist
}

func GetIP() net.IP {
	conn, botstate.Err := net.Dial("udp", "0.0.0.0:80")
	//conn, botstate.Err := net.Dial("udp", "8.8.8.8:80")
	if botstate.Err != nil {
		log.Fatal(botstate.Err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func Checkserver(ip string) bool {
	for _, v := range botstate.ListIp {
		if v == ip {
			return true
		}
	}
	return false
}

func QrKick(client *linetcr.Account, to string) {
	if len(botstate.Banned.Banlist) != 0 {
		for _, v := range botstate.Banned.Banlist {
			if linetcr.IsMembers(client, to, v) == true {
				go func(v string) { client.DeleteOtherFromChats(to, []string{v}) }(v)
			}
			if linetcr.IsPending(client, to, v) == true {
				go func(v string) { client.CancelChatInvitations(to, []string{v}) }(v)
			}
		}
	}
}

func Ungban(group string, asu string) {
	room := linetcr.GetRoom(group)
	if utils.InArrayString(room.Gban, asu) {
		room.Gban = utils.RemoveString(room.Gban, asu)
	}
}

func Addgban(asu string, group string) {
	room := linetcr.GetRoom(group)
	if !utils.InArrayString(room.Gban, asu) && asu != "" {
		room.Gban = append(room.Gban, asu)
	}
}

func AddbanOp3(mid []string) {
	for _, m := range mid {
		botstate.Banned.AddBan(m)
	}
}

func Joinsave(Pelaku string, Optime int64) {
	defer linetcr.PanicOnly()
	ix := utils.IndexOf(botstate.Detectjoin.User, Pelaku)
	if ix == -1 {
		botstate.Detectjoin.User = append(botstate.Detectjoin.User, Pelaku)
		botstate.Detectjoin.Time = append(botstate.Detectjoin.Time, Optime)
	} else {
		botstate.Detectjoin.Time[ix] = Optime
	}
}

func RandomToString(count int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, count)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MaxRevision(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func SendBigImage(AuthToken string, grp string, imglink string) {
	cmd, _ := exec.Command("python3", "libliff/liff.py", AuthToken, grp, imglink).Output()
	fmt.Println("\033[33m" + string(cmd) + "\033[39m")
}

func SendMycreator(sender string) bool {
	if utils.InArrayString(botstate.DEVELOPER, sender) {
		return true
	}
	if botstate.UserBot.GetCreator(sender) {
		return true
	}
	return false
}

func SendMymaker(sender string) bool {
	if SendMycreator(sender) {
		return true
	}
	if botstate.UserBot.GetMaker(sender) {
		return true
	}
	return false
}

func SendMyseller(sender string) bool {
	if SendMymaker(sender) {
		return true
	}
	if botstate.UserBot.GetSeller(sender) {
		return true
	}
	return false
}

func SendMybuyer(sender string) bool {
	if SendMyseller(sender) {
		return true
	}
	if botstate.UserBot.GetBuyer(sender) {
		return true
	}
	return false
}

func SendMyowner(sender string) bool {
	if SendMybuyer(sender) {
		return true
	}
	if botstate.UserBot.GetOwner(sender) {
		return true
	}
	return false
}

func SendMymaster(sender string) bool {
	if SendMyowner(sender) {
		return true
	}
	if botstate.UserBot.GetMaster(sender) {
		return true
	}
	return false
}

func SendMyadmin(sender string) bool {
	if SendMymaster(sender) {
		return true
	}
	if botstate.UserBot.GetAdmin(sender) {
		return true
	}
	return false
}

func SendMygowner(to, sender string) bool {
	if SendMyadmin(sender) {
		return true
	}
	room := linetcr.GetRoom(to)
	if utils.InArrayString(room.Gowner, sender) {
		return true
	}
	return false
}

func SendMygadmin(to, sender string) bool {
	if SendMygowner(to, sender) {
		return true
	}
	room := linetcr.GetRoom(to)
	if utils.InArrayString(room.Gadmin, sender) {
		return true
	}
	return false
}

func GetCodeprem(sender string) string {
	if utils.InArrayString(botstate.DEVELOPER, sender) {
		return "0"
	} else if botstate.UserBot.GetCreator(sender) {
		return "1"
	} else if botstate.UserBot.GetMaker(sender) {
		return "2"
	} else if botstate.UserBot.GetSeller(sender) {
		return "3"
	} else if botstate.UserBot.GetBuyer(sender) {
		return "4"
	} else if botstate.UserBot.GetOwner(sender) {
		return "5"
	} else if botstate.UserBot.GetMaster(sender) {
		return "6"
	} else if botstate.UserBot.GetAdmin(sender) {
		return "7"
	} else if botstate.UserBot.GetBot(sender) {
		return "100"
	}
	return "8"
}

func GetSquad(tok *linetcr.Account, to string) []*linetcr.Account {
	defer utils.PanicHandle("GetSquad")
	nm, memlist, invitee := tok.GetChatList(to)
	Bots := []*linetcr.Account{}
	MIdbot := []string{}
	GoClint := []*linetcr.Account{}
	Gomid := []string{}
	for _, ym := range memlist {
		if utils.InArrayString(botstate.Squadlist, ym) {
			idx := botstate.GetKorban(ym)
			MIdbot = append(MIdbot, ym)
			Bots = append(Bots, idx)
		}
	}
	room := linetcr.GetRoom(to)
	room.Name = nm
	for _, ym := range invitee {
		if utils.InArrayString(botstate.Squadlist, ym) {
			Gomid = append(Gomid, ym)
			idx := botstate.GetKorban(ym)
			GoClint = append(GoClint, idx)
		}
	}
	room.AddSquad(MIdbot, Bots, GoClint, Gomid)
	return Bots
}

func MemUser(group, user string) bool {
	room := linetcr.GetRoom(group)
	if utils.InArrayString(room.Gowner, user) || utils.InArrayString(room.Gadmin, user) {
		return true
	}
	if SendMyadmin(user) {
		return true
	}
	return false
}

func MemBan(group, user string) bool {
	room := linetcr.GetRoom(group)
	if utils.InArrayString(room.Gban, user) {
		return true
	}
	if botstate.Banned.GetBan(user) {
		return true
	}
	return false
}

func MemBan2(group, user string) bool {
	return MemBan(group, user)
}

func IsBlacklist(client *linetcr.Account, user string) bool {
	if botstate.Banned.GetBan(user) {
		return true
	}
	return false
}

func FmtDurations(d time.Duration) string {
	d = d.Round(time.Second)
	x := d
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	if x < 60*time.Second {
		return fmt.Sprintf("%v", x)
	} else if x < 3600*time.Second {
		return fmt.Sprintf("%02dMin", m)
	} else if x < 86400*time.Second {
		return fmt.Sprintf("%02dH %02dMin", h%24, m)
	} else {
		return fmt.Sprintf("%02dD %02dH %02dMin", h/24, h%24, m)
	}
}

func GetContactName(client *linetcr.Account, id string) string {
	contact, botstate.Err := client.GetContact(id)
	if botstate.Err != nil {
		return "Closed Account"
	}
	return contact.DisplayName
}




func Contains(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}
	return false
}

func IsBlacklist(client *linetcr.Account, from string) bool {
	if Contains(botstate.Banned.Banlist, from) == true {
		return true
	}
	return false
}

func IsBlacklist2(client *linetcr.Account, from string) bool {
	if Contains(botstate.Banned.Locklist, from) == true {
		return true
	}
	return false
}

func IsMember(client *linetcr.Account, from string, groups string) bool {
	res, _ := client.GetGroup(groups)
	memlist := res.Members
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}

func Hstg(to, u string) {
	room := linetcr.GetRoom(to)
	if !utils.InArrayString(room.LeaveBack, u) {
		room.LeaveBack = append(room.LeaveBack, u)
	}
}



func DataMention(to string, mtxt string, targetlist []string) {
	defer utils.PanicHandle("DataMention")
	Room := linetcr.GetRoom(to)
	if Room.Id != "" {
		Cans := Room.Client
		if len(Cans) != 0 {
			memlist := []string{}
			for _, mid := range targetlist {
				if !utils.InArrayString(botstate.Squadlist, mid) {
					memlist = append(memlist, mid)
				}
			}
			if len(memlist) != 0 {
				if len(memlist) <= 20 || len(Cans) == 1 {
					cl := Cans[0]
					cl.SendPollMention(to, mtxt, memlist)
					time.Sleep(1 * time.Second)
				} else {
					tx := ""
					nob := 0
					ta := false
					tag := []string{}
					z := len(memlist) / 20
					y := z + 1
					for i := 0; i < y; i++ {
						if !ta {
							tx += fmt.Sprintf("%s\n", mtxt)
							ta = true
						}
						if i == z {
							tag = memlist[i*20:]
							no := i * 20
							no += 1
							for i := 0; i < len(tag); i++ {
								iki := no + i
								if iki < 10 {
									tx += fmt.Sprintf("0%v. @!\n", iki)
								} else {
									tx += fmt.Sprintf("%v. @!\n", iki)
								}
							} 
						} else {
							tag = memlist[i*20 : (i+1)*20]
							no := i * 20
							no += 1
							for i := 0; i < len(tag); i++ {
								iki := no + i
								if iki < 10 {
									tx += fmt.Sprintf("0%v. @!\n", iki)
								} else {
									tx += fmt.Sprintf("%v. @!\n", iki)
								}
							}
						}
						if len(tag) != 0 {
							if nob >= len(Cans) {
								nob = 0
							}
							Cans[nob].SendMention(to, tx, tag)
							nob++
						}
						tx = ""
					}
				}
			}
		}
	}
}

func SelectallBot(client *linetcr.Account, to string) ([]*linetcr.Account, error) {
	botstate.Err, _, memlist := client.GetGroupMembers(to)
	if botstate.Err != nil {
		return nil, botstate.Err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := botstate.GetKorban(mid)
			exe = append(exe, cl)
		}
	}
	if len(exe) != 0 {
		return exe, botstate.Err
	}
	return nil, botstate.Err
}

func BotDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	return fmt.Sprintf("%2d Days, %2d Hours, %2d Mins.", h/24, h%24, m)
}

func AppendLastD(s [][]string, e []string) [][]string {
	defer utils.PanicHandle("Helper")
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func AppendLast(s []string, e string) []string {
	defer utils.PanicHandle("Helper")
	s = utils.RemoveString(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func CheckPermission(num int, sinder string, group string) bool {
	if num == 0 {
		if utils.InArrayString(botstate.DEVELOPER, sinder) {
			return true
		}
	} else if num == 1 {
		if SendMycreator(sinder) {
			return true
		}
	} else if num == 2 {
		if SendMymaker(sinder) {
			return true
		}
	} else if num == 3 {
		if SendMyseller(sinder) {
			return true
		}
	} else if num == 4 {
		if SendMybuyer(sinder) {
			return true
		}
	} else if num == 5 {
		if SendMyowner(sinder) {
			return true
		}
	} else if num == 6 {
		if SendMymaster(sinder) {
			return true
		}
		return false
	} else if num == 7 {
		if SendMyadmin(sinder) {
			return true
		}
	} else if num == 8 {
		if SendMygowner(group, sinder) {
			return true
		}
	} else if num == 9 {
		if SendMygadmin(group, sinder) {
			return true
		}
	}
	return false
}


func InArrayChat(arr []*talkservice.Chat, str *talkservice.Chat) bool {
	for _, tar := range arr {
		if tar.ChatMid == str.ChatMid {
			return true
		}
	}
	return false
}

func InArrayInt64(arr []int64, str int64) bool {
	for _, tar := range arr {
		if tar == str {
			return true
		}
	}
	return false
}

func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func TimeDown(Fucking int) bool {
	switch Fucking {
	case 0:
		time.Sleep(200 * time.Millisecond)
		return true
	case 1:
		time.Sleep(400 * time.Millisecond)
		return true
	case 2:
		time.Sleep(600 * time.Millisecond)
		return true
	case 3:
		time.Sleep(800 * time.Millisecond)
		return true
	case 4:
		time.Sleep(1000 * time.Millisecond)
		return true
	case 5:
		time.Sleep(1200 * time.Millisecond)
		return true
	case 6:
		time.Sleep(1400 * time.Millisecond)
		return true
	case 7:
		time.Sleep(1600 * time.Millisecond)
		return true
	case 8:
		time.Sleep(1800 * time.Millisecond)
		return true
	case 9:
		time.Sleep(2000 * time.Millisecond)
		return true
	case 10:
		time.Sleep(2200 * time.Millisecond)
		return true
	case 11:
		time.Sleep(2400 * time.Millisecond)
		return true
	case 12:
		time.Sleep(2600 * time.Millisecond)
		return true
	case 13:
		time.Sleep(2800 * time.Millisecond)
		return true
	case 14:
		time.Sleep(3000 * time.Millisecond)
		return true
	case 15:
		time.Sleep(3200 * time.Millisecond)
		return true
	case 16:
		time.Sleep(3400 * time.Millisecond)
		return true
	case 17:
		time.Sleep(3600 * time.Millisecond)
		return true
	case 18:
		time.Sleep(3800 * time.Millisecond)
		return true
	case 19:
		time.Sleep(4000 * time.Millisecond)
		return true
	case 20:
		time.Sleep(4200 * time.Millisecond)
		return true
	case 21:
		time.Sleep(4400 * time.Millisecond)
		return true
	case 22:
		time.Sleep(4600 * time.Millisecond)
		return true
	case 23:
		time.Sleep(4800 * time.Millisecond)
		return true
	default:
		return false
	}
}

func StripOut(kata string) string {
	kata = strings.TrimSpace(kata)
	return kata
}

func GetTxt(from string, client *linetcr.Account, pesan string, rname string, sname string, Mid string, MentionMsg []string, group string) string {
	var txt string
	ca, ok := squadMention(MentionMsg)
	if ok {
		pr, _ := ca.GetContact(ca.MID)
		name := pr.DisplayName
		Vs := fmt.Sprintf("@%v", name)
		Vs = strings.ToLower(Vs)
		Vs = strings.TrimSuffix(Vs, " ")
		txt = strings.Replace(pesan, Vs, "", 1)
		txt = strings.TrimPrefix(txt, " ")
		for _, men := range MentionMsg {
			prs, _ := ca.GetContact(men)
			names := prs.DisplayName
			jj := fmt.Sprintf("@%v", names)
			jj = strings.ToLower(jj)
			jj = strings.TrimSuffix(jj, " ")
			txt = strings.Replace(txt, jj, "", 1)
			txt = StripOut(txt)
		}
		botstate.Used = rname
	}
	if strings.HasPrefix(pesan, rname) {
		txt = strings.Replace(pesan, rname, "", 1)
		botstate.Used = rname
	} else if strings.HasPrefix(pesan, sname) {
		txt = strings.Replace(pesan, sname, "", 1)
		botstate.Used = sname
	}
	txt = StripOut(txt)
	return txt
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

func GetGrade(num int) string {
	if num == 0 {
		return "Dev"
	} else if num == 1 {
		return "Creator"
	} else if num == 2 {
		return "Maker"
	} else if num == 3 {
		return "seller"
	} else if num == 4 {
		return "Buyer"
	} else if num == 5 {
		return "Owner"
	} else if num == 6 {
		return "Master"
	} else if num == 7 {
		return "Admin"
	} else if num == 8 {
		return "Gowner"
	} else if num == 9 {
		return "Gadmin"
	}
	return "None"
}

func RemoveSticker(items []*Stickers, item *Stickers) []*Stickers {
	defer linetcr.PanicOnly()
	newitems := []*Stickers{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

func AppendLastSticker(s []*Stickers, e *Stickers) []*Stickers {
	defer linetcr.PanicOnly()
	s = RemoveSticker(s, e)
	s = append(s, e)
	if len(s) >= 1000 {
		s = s[100:]
		return s
	}
	return s
}

func GetArg() string {
	args := os.Args
	if len(os.Args) <= 1 {
		fmt.Println("\033[0;31m not enoght args")
		fmt.Println("\033[37m try :\n\t  \033[33m <app-name> <arg>")
		fmt.Println("\033[37m for example:\n\t \033[33m  ./botline 123")
		fmt.Println("\033[37m or\n\t \033[33m go run *go 123")
		os.Exit(0)
	}
	return args[1]
}

func GetKey(cmd string) string {
	mp := linetcr.HashToMap(botstate.CmdHelper)
	for k, v := range mp {
		if v.(string) == cmd {
			return k
		}
	}
	return cmd
}

func MemAccsess(from string) bool {
	if utils.InArrayString(botstate.Squadlist, from) {
		return false
	} else if botstate.UserBot.GetBot(from) {
		return false
	} else if utils.InArrayString(botstate.DEVELOPER, from) {
		return false
	} else if botstate.UserBot.GetCreator(from) {
		return false
	} else if botstate.UserBot.GetMaker(from) {
		return false
	} else if botstate.UserBot.GetSeller(from) {
		return false
	} else if botstate.UserBot.GetBuyer(from) {
		return false
	} else if botstate.UserBot.GetOwner(from) {
		return false
	} else if botstate.UserBot.GetMaster(from) {
		return false
	} else if botstate.UserBot.GetAdmin(from) {
		return false
	}
	return true
}

func MemUserN(group string, from string) bool {
	Room := linetcr.GetRoom(group)
	if botstate.UserBot.GetBot(from) {
		return false
	} else if utils.InArrayString(botstate.DEVELOPER, from) {
		return false
	} else if botstate.UserBot.GetCreator(from) {
		return false
	} else if botstate.UserBot.GetMaker(from) {
		return false
	} else if botstate.UserBot.GetSeller(from) {
		return false
	} else if botstate.UserBot.GetBuyer(from) {
		return false
	} else if botstate.UserBot.GetOwner(from) {
		return false
	} else if botstate.UserBot.GetMaster(from) {
		return false
	} else if botstate.UserBot.GetAdmin(from) {
		return false
	} else if utils.InArrayString(Room.Gowner, from) {
		return false
	} else if utils.InArrayString(Room.Gadmin, from) {
		return false
	}
	return true
}

func MemEx(to, user string) bool {
	defer linetcr.PanicOnly()
	if botstate.Banned.GetEx(user) {
		return true
	}
	return false
}

func CheckKickUser(group string, user string, invited string) bool {
	Room := linetcr.GetRoom(group)
	if utils.InArrayString(botstate.DEVELOPER, invited) {
		if !utils.InArrayString(botstate.DEVELOPER, user) {
			return true
		}
	} else if botstate.UserBot.GetCreator(invited) {
		if !SendMycreator(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetMaker(invited) {
		if !SendMymaker(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetSeller(invited) {
		if !SendMyseller(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetBuyer(invited) {
		if !SendMybuyer(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetOwner(invited) {
		if !SendMyowner(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetMaster(invited) {
		if !SendMymaster(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetAdmin(invited) {
		if !SendMyadmin(user) && !Allbotlist(user) {
			return true
		}
	} else if utils.InArrayString(Room.Gowner, invited) {
		if !SendMygowner(group, user) && !Allbotlist(user) {
			return true
		}
	} else if utils.InArrayString(Room.Gadmin, invited) {
		if MemUser(group, user) {
			return true
		}
	} else if botstate.UserBot.GetBot(invited) {
		if MemUser(group, user) {
			return true
		}
	}
	return false
}

func CheckBan(client *linetcr.Account, group string) []string {
	list := []string{}
	botstate.Err, _, memlist := client.GetGroupMembers(group)
	if botstate.Err != nil {
		return list
	}
	for mid, _ := range memlist {
		if MemUser(group, mid) {
			if MemBan(group, mid) {
				list = append(list, mid)
			}
		}
	}
	return list
}