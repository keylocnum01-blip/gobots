package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	"strings"
	"os/exec"

	"github.com/kardianos/osext"
	"../botstate"
	"../library/linetcr"
	"../library/hashmap"
	"../utils"
)

func ReloginProgram() error {
	file, botstate.Err := osext.Executable()
	if botstate.Err != nil {
		return botstate.Err
	}
	botstate.Err = syscall.Exec(file, os.Args, os.Environ())
	if botstate.Err != nil {
		return botstate.Err
	}
	return nil
}

func SaveData() {
	file, _ := json.MarshalIndent(botstate.Data, "", "  ")
	_ = ioutil.WriteFile(botstate.DATABASE, file, 0644)
}

func GracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Sutting down application.")
		os.Exit(0)
	}()
}

func Resprem() {
	rngcmd := botstate.GetComs(5, "clone")
	rngcmd = botstate.GetComs(5, "nukeqr")
	rngcmd = botstate.GetComs(5, "joinqr")
	rngcmd = botstate.GetComs(5, "cancelall")
	rngcmd = botstate.GetComs(5, "kickall")
	rngcmd = botstate.GetComs(8, "none")
	rngcmd = botstate.GetComs(8, "max")
	rngcmd = botstate.GetComs(7, "allowall")
	rngcmd = botstate.GetComs(7, "denyall")
	rngcmd = botstate.GetComs(6, "hostage")
	rngcmd = botstate.GetComs(5, "backup")
	rngcmd = botstate.GetComs(6, "upgname")
	rngcmd = botstate.GetComs(6, "welcome")
	rngcmd = botstate.GetComs(5, "sendimage")
	rngcmd = botstate.GetComs(5, "leave")
	rngcmd = botstate.GetComs(6, "announce")
	rngcmd = botstate.GetComs(7, "unban")
	rngcmd = botstate.GetComs(7, "bio")
	rngcmd = botstate.GetComs(7, "tag")
	rngcmd = botstate.GetComs(7, "image")
	rngcmd = botstate.GetComs(7, "contact")
	rngcmd = botstate.GetComs(7, "ban")
	rngcmd = botstate.GetComs(7, "kick")
	rngcmd = botstate.GetComs(7, "vkick")
	rngcmd = botstate.GetComs(7, "invite")
	rngcmd = botstate.GetComs(7, "cancel")
	rngcmd = botstate.GetComs(8, "ungban")
	rngcmd = botstate.GetComs(5, "unbot")
	rngcmd = botstate.GetComs(9, "tagall")
	rngcmd = botstate.GetComs(5, "statusall")
	rngcmd = botstate.GetComs(6, "status")
	rngcmd = botstate.GetComs(6, "whois")
	rngcmd = botstate.GetComs(6, "mute")
	rngcmd = botstate.GetComs(5, "fuck")
	rngcmd = botstate.GetComs(5, "setlimiter")
	rngcmd = botstate.GetComs(5, "setcancel")
	rngcmd = botstate.GetComs(5, "setkick")
	rngcmd = botstate.GetComs(5, "setinvite")
	rngcmd = botstate.GetComs(5, "msgfresh")
	rngcmd = botstate.GetComs(5, "msglimit")
	rngcmd = botstate.GetComs(5, "msgstatus")
	rngcmd = botstate.GetComs(7, "msglurk")
	rngcmd = botstate.GetComs(5, "msgclearban")
	rngcmd = botstate.GetComs(7, "msgleave")
	rngcmd = botstate.GetComs(8, "speed")
	rngcmd = botstate.GetComs(9, "lurk")
	rngcmd = botstate.GetComs(7, "msgwelcome")
	rngcmd = botstate.GetComs(5, "msgrespon")
	rngcmd = botstate.GetComs(5, "setrname")
	rngcmd = botstate.GetComs(5, "setsname")
	rngcmd = botstate.GetComs(5, "notification")
	rngcmd = botstate.GetComs(5, "killmode")
	rngcmd = botstate.GetComs(5, "unowner")
	rngcmd = botstate.GetComs(7, "name")
	rngcmd = botstate.GetComs(5, "Stats")
	rngcmd = botstate.GetComs(5, "buyers")
	rngcmd = botstate.GetComs(5, "upname")
	rngcmd = botstate.GetComs(5, "upstatus")
	rngcmd = botstate.GetComs(5, "acceptall")
	rngcmd = botstate.GetComs(5, "declineall")
	rngcmd = botstate.GetComs(7, "abort")
	rngcmd = botstate.GetComs(5, "accept")
	rngcmd = botstate.GetComs(5, "decline")
	rngcmd = botstate.GetComs(5, "invme")
	rngcmd = botstate.GetComs(5, "gleave")
	rngcmd = botstate.GetComs(5, "Purgeallbans")
	rngcmd = botstate.GetComs(5, "purgeall")
	rngcmd = botstate.GetComs(7, "unsend")
	rngcmd = botstate.GetComs(2, "makers")
	rngcmd = botstate.GetComs(5, "upvcover")
	rngcmd = botstate.GetComs(2, "unseller")
	rngcmd = botstate.GetComs(2, "clearseller")
	rngcmd = botstate.GetComs(3, "sellers")
	rngcmd = botstate.GetComs(2, "seller")
	rngcmd = botstate.GetComs(1, "unmaker")
	rngcmd = botstate.GetComs(1, "clearmaker")
	rngcmd = botstate.GetComs(5, "upvimage")
	rngcmd = botstate.GetComs(5, "upcover")
	rngcmd = botstate.GetComs(5, "upimage")
	rngcmd = botstate.GetComs(3, "clearbuyer")
	rngcmd = botstate.GetComs(3, "unbuyer")
	rngcmd = botstate.GetComs(3, "buyer")
	rngcmd = botstate.GetComs(8, "gaccess")
	rngcmd = botstate.GetComs(5, "allbanlist")
	rngcmd = botstate.GetComs(5, "access")
	rngcmd = botstate.GetComs(7, "expel")
	rngcmd = botstate.GetComs(5, "listcmd")
	rngcmd = botstate.GetComs(4, "owner")
	rngcmd = botstate.GetComs(4, "hide")
	rngcmd = botstate.GetComs(4, "unhide")
	rngcmd = botstate.GetComs(4, "hidelist")
	rngcmd = botstate.GetComs(4, "clearhide")
	rngcmd = botstate.GetComs(7, "mid")
	rngcmd = botstate.GetComs(7, "cleargowner")
	rngcmd = botstate.GetComs(5, "notification")
	rngcmd = botstate.GetComs(4, "clearowner")
	rngcmd = botstate.GetComs(5, "unmaster")
	rngcmd = botstate.GetComs(6, "unmute")
	rngcmd = botstate.GetComs(4, "clearlistcmd")
	rngcmd = botstate.GetComs(5, "setcmd")
	rngcmd = botstate.GetComs(7, "gowner")
	rngcmd = botstate.GetComs(5, "master")
	rngcmd = botstate.GetComs(8, "gojoin")
	rngcmd = botstate.GetComs(7, "ungowner")
	rngcmd = botstate.GetComs(9, "setgroup")
	rngcmd = botstate.GetComs(6, "setbot")
	rngcmd = botstate.GetComs(5, "runtime")
	rngcmd = botstate.GetComs(5, "timenow")
	rngcmd = botstate.GetComs(4, "timeleft")
	rngcmd = botstate.GetComs(9, "say")
	rngcmd = botstate.GetComs(6, "curl")
	rngcmd = botstate.GetComs(6, "ourl")
	rngcmd = botstate.GetComs(9, "here")
	rngcmd = botstate.GetComs(8, "gbanlist")
	rngcmd = botstate.GetComs(5, "clearcache")
	rngcmd = botstate.GetComs(5, "clears")
	rngcmd = botstate.GetComs(5, "cleargban")
	rngcmd = botstate.GetComs(5, "clearbot")
	rngcmd = botstate.GetComs(5, "botlist")
	rngcmd = botstate.GetComs(7, "bans")
	rngcmd = botstate.GetComs(7, "fixed")
	rngcmd = botstate.GetComs(8, "gban")
	rngcmd = botstate.GetComs(5, "bot")
	rngcmd = botstate.GetComs(5, "stay")
	rngcmd = botstate.GetComs(5, "leaveall")
	rngcmd = botstate.GetComs(8, "go")
	rngcmd = botstate.GetComs(7, "stayall")
	rngcmd = botstate.GetComs(5, "bringall")
	rngcmd = botstate.GetComs(6, "listprotect")
	rngcmd = botstate.GetComs(7, "cleargadmin")
	rngcmd = botstate.GetComs(7, "clearban")
	rngcmd = botstate.GetComs(5, "clearadmin")
	rngcmd = botstate.GetComs(5, "upallname")
	rngcmd = botstate.GetComs(5, "upallstatus")
	rngcmd = botstate.GetComs(8, "limitout")
	rngcmd = botstate.GetComs(6, "sayall")
	rngcmd = botstate.GetComs(7, "count")
	rngcmd = botstate.GetComs(9, "ping")
	rngcmd = botstate.GetComs(7, "leave")
	rngcmd = botstate.GetComs(2, "addallsquads")
	rngcmd = botstate.GetComs(3, "addallbots")
	rngcmd = botstate.GetComs(5, "limits")
	rngcmd = botstate.GetComs(4, "adds")
	rngcmd = botstate.GetComs(5, "friends")
	rngcmd = botstate.GetComs(5, "upvallcover")
	rngcmd = botstate.GetComs(5, "upvallimage")
	rngcmd = botstate.GetComs(6, "unsend")
	rngcmd = botstate.GetComs(5, "upallcover")
	rngcmd = botstate.GetComs(5, "upallimage")
	rngcmd = botstate.GetComs(6, "rollcall")
	rngcmd = botstate.GetComs(7, "respon")
	rngcmd = botstate.GetComs(7, "banlist")
	rngcmd = botstate.GetComs(5, "antitag")
	rngcmd = botstate.GetComs(7, "admins")
	rngcmd = botstate.GetComs(8, "gadmin")
	rngcmd = botstate.GetComs(5, "squadmid")
	rngcmd = botstate.GetComs(8, "ungadmin")
	rngcmd = botstate.GetComs(6, "unadmin")
	rngcmd = botstate.GetComs(6, "masters")
	rngcmd = botstate.GetComs(6, "gowners")
	rngcmd = botstate.GetComs(6, "admin")
	rngcmd = botstate.GetComs(5, "unfuck")
	rngcmd = botstate.GetComs(4, "remotegroup")
	rngcmd = botstate.GetComs(5, "groupinfo")
	rngcmd = botstate.GetComs(5, "banpurge")
	rngcmd = botstate.GetComs(5, "autoban")
	rngcmd = botstate.GetComs(5, "autopurge")
	rngcmd = botstate.GetComs(5, "Canceljoin")
	rngcmd = botstate.GetComs(5, "nukejoin")
	rngcmd = botstate.GetComs(5, "groups")
	rngcmd = botstate.GetComs(5, "gourl")
	rngcmd = botstate.GetComs(5, "groupcast")
	rngcmd = botstate.GetComs(5, "fucklist")
	rngcmd = botstate.GetComs(6, "mutelist")
	rngcmd = botstate.GetComs(5, "autojoin")
	rngcmd = botstate.GetComs(5, "ajsjoin")
	rngcmd = botstate.GetComs(4, "perm")
	rngcmd = botstate.GetComs(4, "permlist")
	rngcmd = botstate.GetComs(4, "clearallprotect")
	rngcmd = botstate.GetComs(4, "clearmute")
	rngcmd = botstate.GetComs(4, "clearfuck")
	rngcmd = botstate.GetComs(5, "clearmaster")
	rngcmd = botstate.GetComs(0, "creator")
	rngcmd = botstate.GetComs(0, "clearcreator")
	rngcmd = botstate.GetComs(0, "uncreator")
	rngcmd = botstate.GetComs(0, "creators")
	fmt.Println(rngcmd)
}

var hosts = "https://api.vhtear.com/"
var apikey = "senzu"

func DisableLetterSealing(AuthToken string) {
	Headers := `{
		"AuthToken": "` + AuthToken + `",
		"Msg_Id": "",
		"Device": "ANDROID",
		"Version": "11.5.2",
		"System_Name": "Android OS",
		"System_Ver": "9.1.1",
		"x-lal": "en_US"
	}`
	requestBody := strings.NewReader(Headers)
	res, botstate.Err := http.Post(hosts+"rm_LetterSealing="+apikey, "application/json; charset=UTF-8", requestBody)
	if botstate.Err != nil {
		fmt.Println("Disable Letter Sealing Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Disable Letter Sealing Success")
	}
}

func BackSeave() {
	botstate.DetectCall = botstate.Data.DetectcallBack
	botstate.AutoPro = botstate.Data.AutoproBack
	botstate.AutoPurge = botstate.Data.AutoPurgeBack
	botstate.ProtectMode = botstate.Data.ProtectmodeBack
	botstate.PowerMode = botstate.Data.PowermodeBack
	botstate.KickBanQr = botstate.Data.KickbanqrBack
	botstate.MediaDl = botstate.Data.MediadlBack
	botstate.AutoLike = botstate.Data.AutolikeBack
	botstate.AutoBc = botstate.Data.AutobcBack
	botstate.NukeJoin = botstate.Data.NukejoinBack
	botstate.Canceljoin = botstate.Data.CanceljoinBack
	botstate.AutoJointicket = botstate.Data.AutojointicketBack
	botstate.AutoTranslate = botstate.Data.AutotranslateBack
	botstate.ModeBackup = botstate.Data.ModebackupBack
	botstate.Autojoin = botstate.Data.AutojoinBack
	botstate.Ajsjoin = botstate.Data.AjsjoinBack
	botstate.TypeJoin = botstate.Data.TypejoinBack
	botstate.Typebc = botstate.Data.TypebcBack
	botstate.TypeTrans = botstate.Data.TypetransBack
	botstate.MaxKick = botstate.Data.Maxkick
	botstate.MaxCancel = botstate.Data.Maxcancel
	botstate.MaxInvite = botstate.Data.Maxinvite
	botstate.CancelPend = botstate.Data.Cancelpend
	botstate.TimeBackup = time.Time{}
	botstate.MsSname = botstate.Data.SnameBack
	botstate.MsRname = botstate.Data.RnameBack
	botstate.MsgRespon = botstate.Data.ResponBack
	botstate.Stkid = botstate.Data.KickSticker.Stkid
	botstate.Stkpkgid = botstate.Data.KickSticker.Stkpkgid
	botstate.Stkid2 = botstate.Data.ResponSticker.Stkid2
	botstate.Stkpkgid2 = botstate.Data.ResponSticker.Stkpkgid2
	botstate.Stkid3 = botstate.Data.StayallSticker.Stkid3
	botstate.Stkpkgid3 = botstate.Data.StayallSticker.Stkpkgid3
	botstate.Stkid4 = botstate.Data.LeaveSticker.Stkid4
	botstate.Stkpkgid4 = botstate.Data.LeaveSticker.Stkpkgid4
	botstate.Stkid5 = botstate.Data.KickallSticker.Stkid5
	botstate.Stkpkgid5 = botstate.Data.KickallSticker.Stkpkgid5
	botstate.Stkid6 = botstate.Data.BypassSticker.Stkid6
	botstate.Stkpkgid6 = botstate.Data.BypassSticker.Stkpkgid6
	botstate.Stkid7 = botstate.Data.InviteSticker.Stkid7
	botstate.Stkpkgid7 = botstate.Data.InviteSticker.Stkpkgid7
	botstate.Stkid8 = botstate.Data.ClearbanSticker.Stkid8
	botstate.Stkpkgid8 = botstate.Data.ClearbanSticker.Stkpkgid8
	botstate.Stkid9 = botstate.Data.CancelallSticker.Stkid9
	botstate.Stkpkgid9 = botstate.Data.CancelallSticker.Stkpkgid9
	botstate.MsgBroadcast = botstate.Data.BroadcastBack
	if len(botstate.Data.TimeBanBack) != 0 {
		now := time.Now()
		for a := range botstate.Data.TimeBanBack {
			if utils.InArrayString(botstate.Squadlist, a) {
				tims := botstate.Data.TimeBanBack[a]
				if now.Sub(tims) < 24*time.Hour {
					self := botstate.GetKorban(a)
					if !linetcr.InArrayCl(linetcr.KickBans, self) {
						linetcr.KickBans = append(linetcr.KickBans, self)
						self.TimeBan = tims
					}
					self.Limited = true
					if _, ok := linetcr.GetBlock.Get(self.MID); !ok {
						linetcr.GetBlock.Set(self.MID, tims)
					}
				}

			}
		}
	}
	if len(botstate.Data.CreatorBack) != 0 {
		for _, i := range botstate.Data.CreatorBack {
			botstate.UserBot.AddCreator(i)
		}
	}
	if len(botstate.Data.MakerBack) != 0 {
		for _, i := range botstate.Data.MakerBack {
			botstate.UserBot.AddMaker(i)
		}
	}
	if len(botstate.Data.SellerBack) != 0 {
		for _, i := range botstate.Data.SellerBack {
			botstate.UserBot.AddSeller(i)
		}
	}
	if len(botstate.Data.BuyerBack) != 0 {
		for _, i := range botstate.Data.BuyerBack {
			botstate.UserBot.AddBuyer(i)
		}
	}
	if len(botstate.Data.OwnerBack) != 0 {
		for _, i := range botstate.Data.OwnerBack {
			botstate.UserBot.AddOwner(i)
		}
	}
	if len(botstate.Data.MasterBack) != 0 {
		for _, i := range botstate.Data.MasterBack {
			botstate.UserBot.AddMaster(i)
		}
	}
	if len(botstate.Data.AdminBack) != 0 {
		for _, i := range botstate.Data.AdminBack {
			botstate.UserBot.AddAdmin(i)
		}
	}
	if len(botstate.Data.BotBack) != 0 {
		for _, i := range botstate.Data.BotBack {
			botstate.UserBot.AddBot(i)
		}
	}
	if len(botstate.Data.ProkickBack) != 0 {
		for _, to := range botstate.Data.ProkickBack {
			Room := linetcr.GetRoom(to)
			Room.ProKick = true
		}
	}
	if len(botstate.Data.ProCancelBack) != 0 {
		for _, to := range botstate.Data.ProCancelBack {
			Room := linetcr.GetRoom(to)
			Room.ProCancel = true
		}
	}
	if len(botstate.Data.ProInviteBack) != 0 {
		for _, to := range botstate.Data.ProInviteBack {
			Room := linetcr.GetRoom(to)
			Room.ProInvite = true
		}
	}
	if len(botstate.Data.ProQrBack) != 0 {
		for _, to := range botstate.Data.ProQrBack {
			Room := linetcr.GetRoom(to)
			Room.ProQr = true
		}
	}
	if len(botstate.Data.ProNoteBack) != 0 {
		for _, to := range botstate.Data.ProNoteBack {
			Room := linetcr.GetRoom(to)
			Room.ProNote = true
		}
	}
	if len(botstate.Data.ProNameBack) != 0 {
		for _, to := range botstate.Data.ProNameBack {
			Room := linetcr.GetRoom(to)
			Room.ProName = true
		}
	}
	if len(botstate.Data.ProAlbumBack) != 0 {
		for _, to := range botstate.Data.ProAlbumBack {
			Room := linetcr.GetRoom(to)
			Room.ProAlbum = true
		}
	}
	if len(botstate.Data.ProPictureBack) != 0 {
		for _, to := range botstate.Data.ProPictureBack {
			Room := linetcr.GetRoom(to)
			Room.ProPicture = true
		}
	}
	if len(botstate.Data.ProjoinBack) != 0 {
		for _, to := range botstate.Data.ProjoinBack {
			Room := linetcr.GetRoom(to)
			Room.ProJoin = true
		}
	}
	if len(botstate.Data.AnnunceBack) != 0 {
		for _, to := range botstate.Data.AnnunceBack {
			Room := linetcr.GetRoom(to)
			Room.Announce = true
		}
	}
	if len(botstate.Data.ProLinkBack) != 0 {
		for _, to := range botstate.Data.ProLinkBack {
			Room := linetcr.GetRoom(to)
			Room.ProLink = true
		}
	}
	if len(botstate.Data.ProFlexBack) != 0 {
		for _, to := range botstate.Data.ProFlexBack {
			Room := linetcr.GetRoom(to)
			Room.ProFlex = true
		}
	}
	if len(botstate.Data.ProImageBack) != 0 {
		for _, to := range botstate.Data.ProImageBack {
			Room := linetcr.GetRoom(to)
			Room.ProImage = true
		}
	}
	if len(botstate.Data.ProVideoBack) != 0 {
		for _, to := range botstate.Data.ProVideoBack {
			Room := linetcr.GetRoom(to)
			Room.ProVideo = true
		}
	}
	if len(botstate.Data.ProCallBack) != 0 {
		for _, to := range botstate.Data.ProCallBack {
			Room := linetcr.GetRoom(to)
			Room.ProCall = true
		}
	}
	if len(botstate.Data.ProSpamBack) != 0 {
		for _, to := range botstate.Data.ProSpamBack {
			Room := linetcr.GetRoom(to)
			Room.ProSpam = true
		}
	}
	if len(botstate.Data.ProStickerBack) != 0 {
		for _, to := range botstate.Data.ProStickerBack {
			Room := linetcr.GetRoom(to)
			Room.ProSticker = true
		}
	}
	if len(botstate.Data.ProContactBack) != 0 {
		for _, to := range botstate.Data.ProContactBack {
			Room := linetcr.GetRoom(to)
			Room.ProContact = true
		}
	}
	if len(botstate.Data.ProPostBack) != 0 {
		for _, to := range botstate.Data.ProPostBack {
			Room := linetcr.GetRoom(to)
			Room.ProPost = true
		}
	}
	if len(botstate.Data.ProFileBack) != 0 {
		for _, to := range botstate.Data.ProFileBack {
			Room := linetcr.GetRoom(to)
			Room.ProFile = true
		}
	}
	if len(botstate.Data.GadminBack) != 0 {
		for to := range botstate.Data.GadminBack {
			Room := linetcr.GetRoom(to)
			if len(botstate.Data.GadminBack[to]) != 0 {
				for _, user := range botstate.Data.GadminBack[to] {
					if !utils.InArrayString(Room.Gadmin, user) {
						Room.Gadmin = append(Room.Gadmin, user)
					}
				}
			}
		}
	}
	if len(botstate.Data.GownerBack) != 0 {
		for to := range botstate.Data.GownerBack {
			Room := linetcr.GetRoom(to)
			if len(botstate.Data.GownerBack[to]) != 0 {
				for _, user := range botstate.Data.GownerBack[to] {
					if !utils.InArrayString(Room.Gowner, user) {
						Room.Gowner = append(Room.Gowner, user)
					}
				}
			}
		}
	}
	if len(botstate.Data.GbanBack) != 0 {
		for to := range botstate.Data.GbanBack {
			Room := linetcr.GetRoom(to)
			if len(botstate.Data.GbanBack[to]) != 0 {
				for _, user := range botstate.Data.GbanBack[to] {
					if MemUser(to, user) {
						if !utils.InArrayString(Room.Gban, user) {
							Room.Gban = append(Room.Gban, user)
						}
					}
				}
			}
		}
	}
	if len(botstate.Data.BanBack) != 0 {
		for _, user := range botstate.Data.BanBack {
			botstate.Banned.AddBan(user)
		}
	}
	if len(botstate.Data.FuckBack) != 0 {
		for _, user := range botstate.Data.FuckBack {
			botstate.Banned.AddFuck(user)
		}
	}
	if len(botstate.Data.LockBack) != 0 {
		for _, user := range botstate.Data.LockBack {
			botstate.Banned.AddBan2(user)
		}
	}
	if len(botstate.Data.MuteBack) != 0 {
		for _, user := range botstate.Data.MuteBack {
			botstate.Banned.AddMute(user)
		}
	}
	if len(botstate.Data.WordbanBack) != 0 {
		for _, msg := range botstate.Data.WordbanBack {
			if !utils.InArrayString(botstate.Data.WordbanBack, msg) {
				botstate.Data.WordbanBack = append(botstate.Data.WordbanBack, msg)
			}
		}
	}
	botstate.TimeSave = time.Now()
	fmt.Println("\nSUCCESS RUN ALL BOTS")
}

func Abort() {
	botstate.Remotegrupidto = ""
	botstate.StartChangeImg = false
	botstate.StartChangevImg = false
	botstate.StartChangevImg2 = false
	botstate.Sinderremote = []string{}
	botstate.Remotegrupid = ""
	botstate.RemoteOwner = false
	botstate.RemoteMaster = false
	botstate.RemoteAdmin = false
	botstate.RemoteContact = false
	botstate.RemoteBan = false
	botstate.MidRemote = []string{}
	botstate.Changepic = []*linetcr.Account{}
	botstate.ChangName = false
	botstate.ChangCover = false
	botstate.ChangPict = false
	botstate.ChangeBio = false
	botstate.ChangVpict = false
	botstate.ChangVcover = false
	botstate.AllCheng = false
	botstate.MsgBio = ""
	botstate.MsgName = ""
	botstate.BcImage = false
	botstate.StartBc = false
	botstate.BcVideo = false
	botstate.StartBcV = false
	botstate.GBcImage = false
	botstate.GStartBc = false
	botstate.GBcVideo = false
	botstate.GStartBcV = false
	botstate.FBcImage = false
	botstate.FStartBc = false
	botstate.FBcVideo = false
	botstate.FStartBcV = false
	botstate.SAVEBcImage = false
	botstate.StartSaveBc = false
	botstate.FixedToken = false
	botstate.From_Token = ""
	botstate.Group_Token = ""
	botstate.Timeabort = time.Now()
}

func SaveProHistory() {
	botstate.AllowDoOnce = 0
	for i := range botstate.ClientBot {
		botstate.Data.Kikhistory = botstate.Data.Kikhistory + botstate.ClientBot[i].Ckick
		botstate.Data.Invhistory = botstate.Data.Invhistory + botstate.ClientBot[i].Cinvite
		botstate.Data.Canclhistory = botstate.Data.Canclhistory + botstate.ClientBot[i].Ccancel
		botstate.ClientBot[i].Ckick = 0
		botstate.ClientBot[i].Cinvite = 0
		botstate.ClientBot[i].Ccancel = 0
	}
}

func CheckChatBan() {
	defer utils.PanicHandle("CheckChatBan")
	if botstate.AllowDoOnce == 0 {
		for _, cl := range botstate.ClientBot {
			if !linetcr.InArrayCl(linetcr.KickBanChat, cl) && !cl.Frez {
				r, _ := cl.GetHomeProfile(cl.MID)
				if linetcr.GetBannedChat(r) == 1 {
					linetcr.BanChatAdd(cl)
				}
			}
		}
		botstate.AllowDoOnce++
	}
}


func Addpermcmd(client *linetcr.Account, to string, test1 string, test2 string) {
	x := 0
	numr := 0
	list := ""
	if test1 == "dev" {
		x = 0
	} else if test1 == "creator" {
		x = 1
	} else if test1 == "maker" {
		x = 2
	} else if test1 == "seller" {
		x = 3
	} else if test1 == "buyer" {
		x = 4
	} else if test1 == "owner" {
		x = 5
	} else if test1 == "master" {
		x = 6
	} else if test1 == "admin" {
		x = 7
	} else if test1 == "gowner" {
		x = 8
	} else if test1 == "gadmin" {
		x = 9
	} else {
		list += "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin."
	}
	if list != "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin." {
		cmd2 := test2
		_, value := botstate.SetHelper.Rngcmd[cmd2]
		if value == true {
			if botstate.SetHelper.Rngcmd[cmd2] != x {
				botstate.SetHelper.Rngcmd[cmd2] = x
				numr = 5
			}
		} else {
			list += "botstate.Command not found.\nUse ths botstate.Command First."
		}
	}
	if list != "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin." {
		if list != "botstate.Command not found.\nUse ths botstate.Command First." {
			if numr != 5 {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("%v is already a %v command.\n", cmd2, cmd1)
			} else {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("Changed permission to %v for: %v \n", cmd1, cmd2)
			}
			client.SendMessage(to, botstate.Fancy(list))
		} else {
			client.SendMessage(to, botstate.Fancy(list+"\n"))
		}
	} else {
		client.SendMessage(to, botstate.Fancy(list+"\n"))
	}

}

func SaveBackup() {
	fmt.Println("start Save botstate.Data *__*")
	botstate.Data.DetectcallBack = botstate.DetectCall
	botstate.Data.AutoproBack = botstate.AutoPro
	botstate.Data.AutoPurgeBack = botstate.AutoPurge
	botstate.Data.ProtectmodeBack = botstate.ProtectMode
	botstate.Data.PowermodeBack = botstate.PowerMode
	botstate.Data.KickbanqrBack = botstate.KickBanQr
	botstate.Data.MediadlBack = botstate.MediaDl
	botstate.Data.AutolikeBack = botstate.AutoLike
	botstate.Data.AutobcBack = botstate.AutoBc
	botstate.Data.NukejoinBack = botstate.NukeJoin
	botstate.Data.CanceljoinBack = botstate.Canceljoin
	botstate.Data.AutojointicketBack = botstate.AutoJointicket
	botstate.Data.AutotranslateBack = botstate.AutoTranslate
	botstate.Data.ModebackupBack = botstate.ModeBackup
	botstate.Data.AutojoinBack = botstate.Autojoin
	botstate.Data.AjsjoinBack = botstate.Ajsjoin
	botstate.Data.TypejoinBack = botstate.TypeJoin
	botstate.Data.TypebcBack = botstate.Typebc
	botstate.Data.TypetransBack = botstate.TypeTrans
	botstate.Data.Maxkick = botstate.MaxKick
	botstate.Data.Maxcancel = botstate.MaxCancel
	botstate.Data.Maxinvite = botstate.MaxInvite
	botstate.Data.Cancelpend = botstate.CancelPend
	SaveProHistory()
	botstate.Data.GbanBack = map[string][]string{}
	botstate.Data.GownerBack = map[string][]string{}
	botstate.Data.GadminBack = map[string][]string{}
	botstate.Data.BanBack = []string{}
	botstate.Data.LockBack = []string{}
	botstate.Data.SnameBack = botstate.MsSname
	botstate.Data.RnameBack = botstate.MsRname
	botstate.Data.ResponBack = botstate.MsgRespon
	botstate.Data.KickSticker.Stkid = botstate.Stkid
	botstate.Data.KickSticker.Stkpkgid = botstate.Stkpkgid
	botstate.Data.ResponSticker.Stkid2 = botstate.Stkid2
	botstate.Data.ResponSticker.Stkpkgid2 = botstate.Stkpkgid2
	botstate.Data.StayallSticker.Stkid3 = botstate.Stkid3
	botstate.Data.StayallSticker.Stkpkgid3 = botstate.Stkpkgid3
	botstate.Data.LeaveSticker.Stkid4 = botstate.Stkid4
	botstate.Data.LeaveSticker.Stkpkgid4 = botstate.Stkpkgid4
	botstate.Data.KickallSticker.Stkid5 = botstate.Stkid5
	botstate.Data.KickallSticker.Stkpkgid5 = botstate.Stkpkgid5
	botstate.Data.BypassSticker.Stkid6 = botstate.Stkid6
	botstate.Data.BypassSticker.Stkpkgid6 = botstate.Stkpkgid6
	botstate.Data.InviteSticker.Stkid7 = botstate.Stkid7
	botstate.Data.InviteSticker.Stkpkgid7 = botstate.Stkpkgid7
	botstate.Data.ClearbanSticker.Stkid8 = botstate.Stkid8
	botstate.Data.ClearbanSticker.Stkpkgid8 = botstate.Stkpkgid8
	botstate.Data.CancelallSticker.Stkid9 = botstate.Stkid9
	botstate.Data.CancelallSticker.Stkpkgid9 = botstate.Stkpkgid9
	botstate.Data.BroadcastBack = botstate.MsgBroadcast
	botstate.Data.FuckBack = []string{}
	botstate.Data.MuteBack = []string{}
	botstate.Data.AnnunceBack = []string{}
	botstate.Data.ProQrBack = []string{}
	botstate.Data.ProNameBack = []string{}
	botstate.Data.ProPictureBack = []string{}
	botstate.Data.ProNoteBack = []string{}
	botstate.Data.ProAlbumBack = []string{}
	botstate.Data.ProjoinBack = []string{}
	botstate.Data.ProInviteBack = []string{}
	botstate.Data.ProCancelBack = []string{}
	botstate.Data.ProkickBack = []string{}
	botstate.Data.ProLinkBack = []string{}
	botstate.Data.ProFlexBack = []string{}
	botstate.Data.ProImageBack = []string{}
	botstate.Data.ProVideoBack = []string{}
	botstate.Data.ProCallBack = []string{}
	botstate.Data.ProSpamBack = []string{}
	botstate.Data.ProStickerBack = []string{}
	botstate.Data.ProContactBack = []string{}
	botstate.Data.ProPostBack = []string{}
	botstate.Data.ProFileBack = []string{}
	botstate.Data.CreatorBack = []string{}
	botstate.Data.MakerBack = []string{}
	botstate.Data.SellerBack = []string{}
	botstate.Data.BuyerBack = []string{}
	botstate.Data.OwnerBack = []string{}
	botstate.Data.MasterBack = []string{}
	botstate.Data.AdminBack = []string{}
	botstate.Data.BotBack = []string{}
	botstate.Data.TimeBanBack = map[string]time.Time{}
	if len(linetcr.KickBans) != 0 {
		for _, cl := range linetcr.KickBans {
			if _, ok := linetcr.GetBlock.Get(cl.MID); ok {
				botstate.Data.TimeBanBack[cl.MID] = cl.TimeBan
			}
		}
	}
	for _, room := range linetcr.SquadRoom {
		botstate.Data.GbanBack[room.Id] = []string{}
		botstate.Data.GownerBack[room.Id] = []string{}
		botstate.Data.GadminBack[room.Id] = []string{}
		if room.ProKick {
			botstate.Data.ProkickBack = append(botstate.Data.ProkickBack, room.Id)
		}
		if room.ProCancel {
			botstate.Data.ProCancelBack = append(botstate.Data.ProCancelBack, room.Id)
		}
		if room.ProInvite {
			botstate.Data.ProInviteBack = append(botstate.Data.ProInviteBack, room.Id)
		}
		if room.ProQr {
			botstate.Data.ProQrBack = append(botstate.Data.ProQrBack, room.Id)
		}
		if room.ProName {
			botstate.Data.ProNameBack = append(botstate.Data.ProNameBack, room.Id)
		}
		if room.ProPicture {
			botstate.Data.ProPictureBack = append(botstate.Data.ProPictureBack, room.Id)
		}
		if room.ProNote {
			botstate.Data.ProNoteBack = append(botstate.Data.ProNoteBack, room.Id)
		}
		if room.ProAlbum {
			botstate.Data.ProAlbumBack = append(botstate.Data.ProAlbumBack, room.Id)
		}
		if room.ProJoin {
			botstate.Data.ProjoinBack = append(botstate.Data.ProjoinBack, room.Id)
		}
		if room.ProLink {
			botstate.Data.ProLinkBack = append(botstate.Data.ProLinkBack, room.Id)
		}
		if room.ProFlex {
			botstate.Data.ProFlexBack = append(botstate.Data.ProFlexBack, room.Id)
		}
		if room.ProImage {
			botstate.Data.ProImageBack = append(botstate.Data.ProImageBack, room.Id)
		}
		if room.ProVideo {
			botstate.Data.ProVideoBack = append(botstate.Data.ProVideoBack, room.Id)
		}
		if room.ProCall {
			botstate.Data.ProCallBack = append(botstate.Data.ProCallBack, room.Id)
		}
		if room.ProSpam {
			botstate.Data.ProSpamBack = append(botstate.Data.ProSpamBack, room.Id)
		}
		if room.ProSticker {
			botstate.Data.ProStickerBack = append(botstate.Data.ProStickerBack, room.Id)
		}
		if room.ProContact {
			botstate.Data.ProContactBack = append(botstate.Data.ProContactBack, room.Id)
		}
		if room.ProPost {
			botstate.Data.ProPostBack = append(botstate.Data.ProPostBack, room.Id)
		}
		if room.ProFile {
			botstate.Data.ProFileBack = append(botstate.Data.ProFileBack, room.Id)
		}
		if room.Announce {
			botstate.Data.AnnunceBack = append(botstate.Data.AnnunceBack, room.Id)
		}
	}
	if len(botstate.UserBot.Creator) != 0 {
		for _, i := range botstate.UserBot.Creator {
			if !utils.InArrayString(botstate.Data.CreatorBack, i) {
				botstate.Data.CreatorBack = append(botstate.Data.CreatorBack, i)
			}
		}
	}
	if len(botstate.UserBot.Maker) != 0 {
		for _, i := range botstate.UserBot.Maker {
			if !utils.InArrayString(botstate.Data.MakerBack, i) {
				botstate.Data.MakerBack = append(botstate.Data.MakerBack, i)
			}
		}
	}
	if len(botstate.UserBot.Seller) != 0 {
		for _, i := range botstate.UserBot.Seller {
			if !utils.InArrayString(botstate.Data.SellerBack, i) {
				botstate.Data.SellerBack = append(botstate.Data.SellerBack, i)
			}
		}
	}
	if len(botstate.UserBot.Buyer) != 0 {
		for _, i := range botstate.UserBot.Buyer {
			if !utils.InArrayString(botstate.Data.BuyerBack, i) {
				botstate.Data.BuyerBack = append(botstate.Data.BuyerBack, i)
			}
		}
	}
	if len(botstate.UserBot.Owner) != 0 {
		for _, i := range botstate.UserBot.Owner {
			if !utils.InArrayString(botstate.Data.OwnerBack, i) {
				botstate.Data.OwnerBack = append(botstate.Data.OwnerBack, i)
			}
		}
	}
	if len(botstate.UserBot.Master) != 0 {
		for _, i := range botstate.UserBot.Master {
			if !utils.InArrayString(botstate.Data.MasterBack, i) {
				botstate.Data.MasterBack = append(botstate.Data.MasterBack, i)
			}
		}
	}
	if len(botstate.UserBot.Admin) != 0 {
		for _, i := range botstate.UserBot.Admin {
			if !utils.InArrayString(botstate.Data.AdminBack, i) {
				botstate.Data.AdminBack = append(botstate.Data.AdminBack, i)
			}
		}
	}
	if len(botstate.UserBot.Bot) != 0 {
		for _, i := range botstate.UserBot.Bot {
			if !utils.InArrayString(botstate.Data.BotBack, i) {
				botstate.Data.BotBack = append(botstate.Data.BotBack, i)
			}
		}
	}
	if len(botstate.Data.GbanBack) != 0 {
		for to := range botstate.Data.GbanBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gban) != 0 {
				for _, i := range Room.Gban {
					if MemUser(to, i) {
						if !utils.InArrayString(botstate.Data.GbanBack[to], i) {
							botstate.Data.GbanBack[to] = append(botstate.Data.GbanBack[to], i)
						}
					}
				}
			}
		}
	}
	if len(botstate.Data.GownerBack) != 0 {
		for to := range botstate.Data.GownerBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gowner) != 0 {
				for _, i := range Room.Gowner {
					if !utils.InArrayString(botstate.Data.GownerBack[to], i) {
						botstate.Data.GownerBack[to] = append(botstate.Data.GownerBack[to], i)
					}
				}
			}
		}
	}
	if len(botstate.Data.GadminBack) != 0 {
		for to := range botstate.Data.GadminBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gadmin) != 0 {
				for _, i := range Room.Gadmin {
					if !utils.InArrayString(botstate.Data.GadminBack[to], i) {
						botstate.Data.GadminBack[to] = append(botstate.Data.GadminBack[to], i)
					}
				}
			}
		}
	}
	if len(botstate.Banned.Banlist) != 0 {
		for _, i := range botstate.Banned.Banlist {
			if MemAccsess(i) {
				if !utils.InArrayString(botstate.Data.BanBack, i) {
					botstate.Data.BanBack = append(botstate.Data.BanBack, i)
				}
			}
		}
	}
	if len(botstate.Banned.Fucklist) != 0 {
		for _, i := range botstate.Banned.Fucklist {
			if MemAccsess(i) {
				if !utils.InArrayString(botstate.Data.FuckBack, i) {
					botstate.Data.FuckBack = append(botstate.Data.FuckBack, i)
				}
			}
		}
	}
	if len(botstate.Banned.Locklist) != 0 {
		for _, i := range botstate.Banned.Locklist {
			if MemAccsess(i) {
				if !utils.InArrayString(botstate.Data.LockBack, i) {
					botstate.Data.LockBack = append(botstate.Data.LockBack, i)
				}
			}
		}
	}
	if len(botstate.Banned.Mutelist) != 0 {
		for _, i := range botstate.Banned.Mutelist {
			if MemAccsess(i) {
				if !utils.InArrayString(botstate.Data.MuteBack, i) {
					botstate.Data.MuteBack = append(botstate.Data.MuteBack, i)
				}
			}
		}
	}
	if len(botstate.Data.WordbanBack) != 0 {
		for _, msg := range botstate.Data.WordbanBack {
			if !utils.InArrayString(botstate.Data.WordbanBack, msg) {
				botstate.Data.WordbanBack = append(botstate.Data.WordbanBack, msg)
			}
		}
	}
	fmt.Println("done save botstate.Data *__*")
	SaveData()
}

func LeaveallGroups(client *linetcr.Account, to string) []string {
	allg := []string{}
	for i := range botstate.ClientBot {
		groups, _ := botstate.ClientBot[i].GetGroupIdsJoined()
		grup, _ := botstate.ClientBot[i].GetGroups(groups)
		for _, gi := range grup {
			if gi.ChatMid != to {
				botstate.ClientBot[i].LeaveGroup(gi.ChatMid)
				time.Sleep(1 * time.Second)
				if !utils.InArrayString(allg, gi.ChatMid) {
					allg = append(allg, gi.ChatMid)
				}
			}
		}
	}
	return allg
}