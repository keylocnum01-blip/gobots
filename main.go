package main
//BOT_GO_MULTY_WAR_PROTECTION
//RECODE_BY: SELFTCR™
//ID_LINE: code-bot
//NEW__All_UPDATE: 21-04-2024

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"github.com/shirou/gopsutil/mem"
	"unicode/utf8"
	"regexp"
	"log"
	"github.com/kardianos/osext"
	"github.com/tidwall/gjson"
	"net"
	"os/signal"
	"syscall"
	"strconv"
	"runtime"
	"net/http"
	"runtime/debug"
	"sort"
	"strings"
	"time"
	"./library/linetcr"
       "./library/unistyle"
	"github.com/panjf2000/ants"
	"./library/hashmap"
	"./library/SyncService"
	mod "./library/modcompact"
	call "./library/libcall/call"
	namegenerator "./library/RandomName"
	talkservice "./library/linethrift"
	valid "github.com/asaskevich/govalidator"
	"github.com/shirou/gopsutil/host"
	"github.com/opalmer/check-go-version/api"
)
var fancy = unistyle.Normal
//kick_sticker
var getStickerKick int = 0
var stkid string = ""
var stkpkgid string = ""
//respon_sticker
var getStickerRespon int = 0
var stkid2 string = ""
var stkpkgid2 string = ""
//stayall_sticker
var getStickerStayall int = 0
var stkid3 string = ""
var stkpkgid3 string = ""
//leave_sticker
var getStickerLeave int = 0
var stkid4 string = ""
var stkpkgid4 string = ""
//kickall_sticker
var getStickerKickall int = 0
var stkid5 string = ""
var stkpkgid5 string = ""
//bypass_sticker
var getStickerBypass int = 0
var stkid6 string = ""
var stkpkgid6 string = ""
//invite_sticker
var getStickerInvite int = 0
var stkid7 string = ""
var stkpkgid7 string = ""
//clearban_sticker
var getStickerClearban int = 0
var stkid8 string = ""
var stkpkgid8 string = ""
//cancelall_sticker
var getStickerCancelall int = 0
var stkid9 string = ""
var stkpkgid9 string = ""
var (
	GO               = getArg()
	Whitelist        = []string{}
	SetHelper        = &linetcr.Helper{Rngcmd: make(map[string]int)}
	DB               *DATA
	ClientBot        []*linetcr.Account
	Midlist          []string
	aclear           = time.Now()
	Grupas           []*talkservice.Group
	Poll             *linetcr.Account
	Self             *linetcr.Account
	cpu              int
	err 			 error
	botleave         = &hashmap.HashMap{}
	changepic        []*linetcr.Account
	timeabort        = time.Now()
	TimeSave         = time.Now()
	TimeClear         = time.Now()
	Laststicker      = &hashmap.HashMap{}
	ChangCover       = false
	MsgRespon        = "Ready !!!"
	TimeBc         = time.Now()
	MsgBroadcast        = "tester"
	TimeBroadcast        = 1
	AutoBc        = false
	AutoBackBot        = false
	Timebk    = 6
	MsgLock           = "Success Clear %v locklist."
	MsgBan           = "Success Clear %v blacklist."
	MsFresh          = "✓"
	MsLimit          = "✘"
	MsSname          = "x"
	MsRname          = "a"
	AllCheng         = false
	Lastleave        = &hashmap.HashMap{}
	ChangPict        = false
	ChangName        = false
	ProtectMode      = false
	AutokickBan      = true
	ChangVpict       = false
	ChangVcover      = false
	ChangeBio        = false
	CmdHelper        = &hashmap.HashMap{}
	cewel            = &hashmap.HashMap{}
	cleave           = &hashmap.HashMap{}
	ScanTarget         = false
	PowerMode         = false
	PublicMode         = true
	LockMode         = false
	NukeJoin         = false
	AutoBan          = true
	KickBanQr         = false
	canceljoin       = false
	Autojoin         = "off"
	Ajsjoin         = "qr"
	backlist         = &hashmap.HashMap{}
	cekoptime        = []int64{}
	Ceknuke          = &hashmap.HashMap{}
	Cekstaybot       = &hashmap.HashMap{}
	Commands         = &linetcr.Command{}
	Waitlistin       = map[string][]string{}
	AutoproN         = false
	LogMode          = false
       DetectCall        = false
       AutoLike        = false
       BomLike        = false
       MediaDl        = false
	LogGroup         = ""
	delayed          = 10 * time.Second
	MsgBio           = ""
	MsgName          = ""
	FixedToken       = false
	From_Token       = ""
	Group_Token      = ""
	StartChangeImg   = false
	StartChangevImg  = false
	StartChangevImg2 = false
	AutoPro          = false
	Command          = &hashmap.HashMap{}
	tempginv         = []string{}
	remotegrupidto   = ""
	ModeBackup       = "inv"
	checkHaid        = []string{}
	botStart         = time.Now()
	TimeBackup       = time.Now()
	oplist           = []int64{}
	oplistinvite     = []int64{}
	PurgeOP          = []int64{}
	oplistjoin       = []int64{}
	AutoPurge        = true
	BcImage        = false
	StartBc        = false
	BcVideo        = false
	StartBcV        = false
	GBcImage        = false
	GStartBc        = false
	GBcVideo        = false
	GStartBcV        = false
	FBcImage        = false
	FStartBc        = false
	FBcVideo        = false
	FStartBcV        = false
	SAVEBcImage        = false
	startSaveBc        = false
	ClientMid        = map[string]*linetcr.Account{}
	Squadlist        = []string{}
	argsRaw          = os.Args
	Sinderremote     = []string{}
	StartChangeVideo = false
	tempgroup        = []string{}
	Lastinvite       = &hashmap.HashMap{}
	Lastkick         = &hashmap.HashMap{}
	Lastjoin         = &hashmap.HashMap{}
	Lastcancel       = &hashmap.HashMap{}
	Nkick            = &hashmap.HashMap{}
	Lastupdate       = &hashmap.HashMap{}
	Lastmid          = &hashmap.HashMap{}
	filterop         = &hashmap.HashMap{}
	Lasttag          = &hashmap.HashMap{}
	Lastcon          = &hashmap.HashMap{}
	Lastmessage       = &hashmap.HashMap{}
	Commandss        = &hashmap.HashMap{}
	Detectjoin       = &linetcr.SaveJoin{}
	Banned           = &linetcr.BanUser{Banlist: []string{}, Fucklist: []string{}, Mutelist: []string{}, Exlist: []string{}, Locklist: []string{}}
	UserBot          = &linetcr.Access{Creator: []string{}, Maker: []string{}, Seller: []string{}, Buyer: []string{}, Owner: []string{}, Master: []string{}, Admin: []string{}, Bot: []string{}}
	timeSend         = []int64{}
	opkick 			 = []int64{}
	opjoin           = []string{}
	Cekpurge         = []int64{}
	MaxCancel        = 4
	MaxKick          = 3
	MaxInvite          = 5
	CancelPend          = 5
	CountSpam         = "5"
	CountAjs         = "2"
	AllowDoOnce       = 0
	LockAjs      = false
	cekGo            = []int64{}
	UpdatePicture    = map[string]bool{}
	UpdateCover      = map[string]bool{}
	UpdateVProfile   = map[string]bool{}
	UpdateVCover     = map[string]bool{}
	Qrwar            = false
	filterWar        = &kickop{Kick: []string{}, Inv: []string{}, Opinv: []int64{}}
	ColorCyan        = "\033[36m"
	ColorReset       = "\033[0m"
	Data             DATA
	remotegrupid     = ""
	LastActive       = &hashmap.HashMap{}
	used             = ""
	IPServer         string
	Killmode         = "none"
	Typebc         = "none"
	AutoJointicket = false
	TypeJoin         = "none"
	AutoTranslate        = false
	TypeTrans         = "EN"
	filtermsg   = &hashmap.HashMap{}
	Opinvite    = []int64{}
	stringToInt = []rune("01")
	DATABASE    = "database/" + getArg() + ".json"
	CREATOR      = []string{"u0950ac5584daf10c380a53085378775d"}
	DEVELOPER      = []string{"u0950ac5584daf10c380a53085378775d","u0fbc4b49c6194469781badf7d3194284"}
	TeamNotif      = []string{""}
	AntiJs      = []string{}
	MidBc      = []string{}
	MidRemote      = []string{}
	RemoteOwner = false
	RemoteMaster = false
	RemoteAdmin = false
	RemoteContact = false
	RemoteBan = false
	HostName    = []string{
		"legy-jp-addr-long",
	}
	carierMap = map[string]string{
		"IOSIPAD":     "51089, 1-0",
		"IOS":         "51089, 1-0",
		"ANDROIDLITE": "51000",
		"ANDROID":     "51010",
		"CHROMEOS":    "",
		"DESKTOPMAC":  "",
		"DESKTOPWIN":  "",
		"CHANNELCP":   "51010",
	}
	helppublic = []string{
		"ɢᴇᴛᴘɪᴄᴛ [@]",
		"ɢᴇᴛɴᴀᴍᴇ [@]",
		"ɢᴇᴛʙɪᴏ [@]",
		"ɢᴇᴛᴄᴏᴠᴇʀ [@]",
		"ɢᴇxᴛʀᴀᴄᴏᴠᴇʀ [@]",
		"ɢᴇᴛꜱᴛᴏʀʏ [@]",
		"ɢᴇᴛᴍɪᴅ [@]",
		"ɢᴇᴛᴄᴏɴᴛᴀᴄᴛ [@]",
		"ɢᴇᴛᴄᴀʟʟ",
		"ɢʀᴏᴜᴘɪɴꜰᴏ",
		"ᴍᴇɴᴛɪᴏɴᴀʟʟ",
		"ʟᴜʀᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴡᴇʟᴄᴏᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴅᴇᴛᴇᴄᴛᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
	}
	helpreply = []string{
		"ɢᴇᴛᴘɪᴄᴛ",
		"ɢᴇᴛɴᴀᴍᴇ",
		"ɢᴇᴛʙɪᴏ",
		"ɢᴇᴛᴄᴏᴠᴇʀ",
		"ɢᴇxᴛʀᴀᴄᴏᴠᴇʀ",
		"ɢᴇᴛꜱᴛᴏʀʏ",
		"ɢᴇᴛᴍɪᴅ",
		"ɢᴇᴛᴄᴏɴᴛᴀᴄᴛ",
	}
	helppro = []string{
		"ᴀʟʟᴘʀᴏᴛᴇᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪɴᴠɪᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴋɪᴄᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀɴᴄᴇʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏQʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢɴᴀᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢᴘɪᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɴᴏᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴀʟʙᴜᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏʟɪɴᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰʟᴇx [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪᴍᴀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴠɪᴅᴇᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴘᴀᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴛɪᴄᴋᴇʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴏɴᴛᴀᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴘᴏꜱᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰɪʟᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴛᴇᴄᴛᴍᴀx [ᴏɴ/ᴏꜰꜰ]",
	}
	ListIp = []string{}
	helpdeveloper = []string{
		"ᴄʀᴇᴀᴛᴏʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴄʀᴇᴀᴛᴏʀ [@/ʟᴀꜱᴛ]",
		"ᴄʀᴇᴀᴛᴏʀꜱ",
		"ᴄʟᴇᴀʀᴄʀᴇᴀᴛᴏʀ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ʀᴇʙᴏᴏᴛᴠᴘꜱ",
	}
	helpcreator = []string{
		"ᴀᴘᴘɴᴀᴍᴇ",
		"ᴜꜱᴇʀᴀɢᴇɴᴛ",
		"ʜᴏꜱᴛɴᴀᴍᴇ",
		"ᴄᴇᴋᴛᴏᴋᴇɴ",
		"ᴍᴀᴋᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴍᴀᴋᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴍᴀᴋᴇʀꜱ",
		"ᴄʟᴇᴀʀᴍᴀᴋᴇʀ",
		"ᴇɴᴀʙʟᴇᴇ2ᴇᴇ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ᴀᴅᴅꜱ",
	}
	helpmaker = []string{
		"ᴀᴅᴅᴀʟʟꜱQᴜᴀᴅꜱ",
		"ᴄᴏꜱQᴜᴀᴅ",
		"ᴀᴅᴅᴛᴏᴋᴇɴ [ᴛᴇxᴛ]",
		"ᴜɴᴛᴏᴋᴇɴ [ɴᴜᴍ]",
		"ʟɪꜱᴛ ᴛᴏᴋᴇɴ",
		"ꜱᴛᴀᴛᴜꜱ ᴛᴏᴋᴇɴ",
		"ʀᴇᴍᴏᴠᴇʟɪᴍɪᴛꜱ",
		"ʀᴇᴍᴛᴏᴋᴇɴʙᴀɴꜱ",
		"ꜱᴇʟʟᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴꜱᴇʟʟᴇʀ [@/ʟᴀꜱᴛ]",
		"ꜱᴇʟʟᴇʀꜱ",
		"ᴄʟᴇᴀʀꜱᴇʟʟᴇʀ",
		"ᴍᴀᴋᴇʀꜱ",
		"ᴀᴅᴅꜰʀɪᴇɴᴅꜱ [@]",
		"ꜰʀɪᴇɴᴅꜱ [@]",
		"ᴄʟᴇᴀʀꜰʀɪᴇɴᴅꜱ",
		"ᴜɴꜰʀɪᴇɴᴅʙᴀɴꜱ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ʀᴇʙᴏᴏᴛ",
		"ʀᴜɴᴀʟʟ",
	}
	helpseller = []string{
		"ᴀᴅᴅᴅᴀʏ [ɴᴜᴍ]",
		"ᴀᴅᴅᴡᴇᴇᴋ [ɴᴜᴍ]",
		"ᴀᴅᴅᴍᴏɴᴛʜ [ɴᴜᴍ]",
		"ꜱᴇᴛᴅᴀᴛᴇ [ᴛᴇxᴛ]",
		"ꜱᴇʟʟᴇʀꜱ",
		"ʙᴜʏᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴʙᴜʏᴇʀ [@/ʟᴀꜱᴛ]",
		"ʙᴜʏᴇʀꜱ",
		"ᴄʟᴇᴀʀʙᴜʏᴇʀ",
		"ʙᴏᴛʟɪꜱᴛ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ᴜᴘᴀʟʟɴᴀᴍᴇ [ᴛᴇxᴛ]",
		"ᴜᴘᴀʟʟꜱᴛᴀᴛᴜꜱ [ᴛᴇxᴛ]",
		"ᴜᴘᴀʟʟᴄᴏᴠᴇʀ",
		"ᴜᴘᴀʟʟɪᴍᴀɢᴇ",
		"ᴜᴘᴠᴀʟʟᴄᴏᴠᴇʀ",
		"ᴜᴘᴠᴀʟʟɪᴍᴀɢᴇ",
		"ᴄʟᴇᴀʀʙᴏᴛ",
		"ᴄʟᴇᴀʀʙᴀɴ",
		"ᴄʟᴇᴀʀᴏᴡɴᴇʀ",
		"ᴄʟᴇᴀʀᴍᴀꜱᴛᴇʀ",
		"ᴄʟᴇᴀʀᴀᴅᴍɪɴ",
		"ᴄʟᴇᴀʀꜰᴜᴄᴋ",
		"ᴄʟᴇᴀʀʜɪᴅᴇ",
		"ᴄʟᴇᴀʀʟɪꜱᴛᴄᴍᴅ",
		"ᴄʟᴇᴀʀᴍᴜᴛᴇ",
		"ᴄʟᴇᴀʀᴀʟʟᴘʀᴏᴛᴇᴄᴛ",
	}
	helpbuyer = []string{
		"ᴀʙᴏᴜᴛ",
		"ᴀᴄᴄᴇꜱꜱ",
		"ᴀᴄᴄᴇᴘᴛ [ɴᴜᴍ]",
		"ᴀᴄᴄᴇᴘᴛᴀʟʟ",
		"ᴀᴜᴛᴏɴᴀᴍᴇ",
		"ᴀᴜᴛᴏɪᴍᴀɢᴇ",
		"ᴀᴜᴛᴏᴄᴏᴠᴇʀ",
		"ʀᴀɴᴅᴏᴍᴘʀᴏꜰɪʟᴇ",
		"ʙᴜʏᴇʀꜱ",
		"ᴄʟᴇᴀʀᴀʟʟᴘʀᴏᴛᴇᴄᴛ",
		"ᴄʟᴇᴀʀʙᴀɴ",
		"ᴄʟᴇᴀʀᴍᴀꜱᴛᴇʀ",
		"ᴄʟᴇᴀʀᴀᴅᴍɪɴ",
		"ᴄʟᴇᴀʀꜰᴜᴄᴋ",
		"ᴄʟᴇᴀʀʜɪᴅᴇ",
		"ᴄʟᴇᴀʀʟɪꜱᴛᴄᴍᴅ",
		"ᴄʟᴇᴀʀᴍᴜᴛᴇ",
		"ᴄʟᴏɴᴇ [@/ʟᴀꜱᴛ]",
		"ᴅᴇᴄʟɪɴᴇ [ɴᴜᴍ]",
		"ᴅᴇᴄʟɪɴᴇᴀʟʟ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ɢʟᴇᴀᴠᴇ [ɴᴜᴍ]",
		"ɢᴏᴜʀʟ [ɴᴜᴍ]",
		"ɢʀᴏᴜᴘꜱ",
		"ɢʀᴏᴜᴘɪɴꜰᴏ [ɴᴜᴍ]",
		"ɢʀᴏᴜᴘᴄᴀꜱᴛ [ᴛᴇxᴛ]",
		"ɢʀᴏᴜᴘʙᴄ ɪᴍᴀɢᴇ",
		"ɢʀᴏᴜᴘʙᴄ ᴠɪᴅᴇᴏ",
		"ꜰʀɪᴇɴᴅᴄᴀꜱᴛ [ᴛᴇxᴛ]",
		"ꜰʀɪᴇɴᴅʙᴄ ɪᴍᴀɢᴇ",
		"ꜰʀɪᴇɴᴅʙᴄ ᴠɪᴅᴇᴏ",
		"ʙᴄᴀꜱᴛ [ᴍɪᴅ]:[ᴍꜱɢ]",
		"ᴜɴɪᴄᴏᴅᴇ [ᴍɪᴅ]:[ɴᴜᴍ]",
		"ꜱᴛᴀʀᴛʙᴄᴍɪᴅ ɪᴍᴀɢᴇ",
		"ꜱᴛᴀʀᴛʙᴄᴍɪᴅ ᴠɪᴅᴇᴏ",
		"ʟɪꜱᴛ ʙʀᴏᴀᴅᴄᴀꜱᴛ",
		"ʜɪᴅᴇ [@/ʟᴀꜱᴛ]",
		"ᴜɴʜɪᴅᴇ [@/ʟᴀꜱᴛ]",
		"ʜɪᴅᴇʟɪꜱᴛ",
		"ɪɴᴠᴍᴇ [ɴᴜᴍ]",
		"ʟᴇᴀᴠᴇᴀʟʟ",
		"ʟɪꜱᴛ ᴘʀᴏᴛᴇᴄᴛ",
		"ᴍꜱɢᴄʟᴇᴀʀʙᴀɴ [ᴛᴇxᴛ]",
		"ᴍꜱɢꜰʀᴇꜱʜ [ᴛᴇxᴛ]",
		"ᴍꜱɢʟɪᴍɪᴛ [ᴛᴇxᴛ]",
		"ᴍꜱɢʀᴇꜱᴘᴏɴ [ᴛᴇxᴛ]",
		"ᴍꜱɢꜱᴛᴀᴛᴜꜱ [ᴛᴇxᴛ]",
		"ɴᴜᴋᴇᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴏᴡɴᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴏᴡɴᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴏᴡɴᴇʀꜱ",
		"ᴄʟᴇᴀʀᴏᴡɴᴇʀ",
		"ᴘᴇʀᴍ [ʀᴀɴᴋ - ᴄᴍᴅ]",
		"ᴘᴇʀᴍʟɪꜱᴛ:",
		"ʀᴇᴍᴏᴛᴇɢʀᴏᴜᴘ [ɴᴜᴍ]",
		"ʀᴇᴍᴏᴛᴇᴏᴡɴᴇʀ [ɴᴜᴍ]",
		"ʀᴇᴍᴏᴛᴇᴍᴀꜱᴛᴇʀ [ɴᴜᴍ]",
		"ʀᴇᴍᴏᴛᴇᴀᴅᴍɪɴ [ɴᴜᴍ]",
		"ʀᴇᴍᴏᴛᴇᴄᴏɴᴛᴀᴄᴛ [ɴᴜᴍ]",
		"ʀᴇᴍᴏᴛᴇʙᴀɴ [ɴᴜᴍ]",
		"ᴜᴘᴄᴏᴠᴇʀ",
		"ᴜᴘɪᴍᴀɢᴇ",
		"ᴜᴘᴠᴄᴏᴠᴇʀ",
		"ᴜᴘᴠɪᴍᴀɢᴇ",
		"ᴜᴘɴᴀᴍᴇ [ᴛᴇxᴛ]",
		"ᴜᴘꜱᴛᴀᴛᴜꜱ [ᴛᴇxᴛ]",
		"ᴜᴘɢɴᴀᴍᴇ [ᴛᴇxᴛ]",
		"ꜱᴇᴛᴄᴀɴᴄᴇʟ [ɴᴜᴍ]",
		"ꜱᴇᴛʟᴏɢᴏ [ᴛᴇxᴛ]",
		"ꜱᴇᴛᴋɪᴄᴋ [ɴᴜᴍ]",
		"ꜱᴇᴛɪɴᴠɪᴛᴇ [ɴᴜᴍ]",
		"ꜱᴇᴛʟɪᴍɪᴛᴇʀ [ɴᴜᴍ]",
		"ꜱᴇᴛᴘᴇɴᴅ [ɴᴜᴍ]",
		"ꜱᴇᴛʀɴᴀᴍᴇ [ᴛᴇxᴛ]",
		"ꜱᴇᴛꜱɴᴀᴍᴇ [ᴛᴇxᴛ]",
		"ꜱᴛᴀᴛᴜꜱᴀʟʟ",
		"ꜱᴛᴀᴛꜱ",
		"ᴛɪᴍᴇʟᴇꜰᴛ",
	}
	helpowner = []string{
		"ᴀʟʟʙᴀɴʟɪꜱᴛ",
		"ᴀʟʟɢᴀᴄᴄᴇꜱꜱ",
		"ᴀɴᴛɪᴛᴀɢ [ᴏɴ/ᴏꜰꜰ]",
		"ʙʀɪɴɢᴀʟʟ",
		"ᴄʟᴇᴀʀᴄᴀᴄʜᴇ",
		"ᴄʟᴇᴀʀꜱ",
		"ᴄʟᴇᴀʀᴀʟʟ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ꜰᴜᴄᴋ [@/ʟᴀꜱᴛ]",
		"ᴜɴꜰᴜᴄᴋ [@/ʟᴀꜱᴛ]",
		"ꜰᴜᴄᴋʟɪꜱᴛ",
		"ᴊᴏɪɴQʀ [ʟɪɴᴋ]",
		"ɴᴜᴋᴇQʀ [ʟɪɴᴋ]",
		"ᴀᴊꜱᴊᴏɪɴ [ɪɴᴠ/Qʀ]",
		"ᴀᴜᴛᴏᴊᴏɪɴ [ɪɴᴠ/Qʀ]",
		"ᴀᴜᴛᴏᴘʀᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ᴀᴜᴛᴏᴘᴜʀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴀᴜᴛᴏᴛɪᴄᴋᴇᴛ [ɴᴏʀᴍᴀʟ]",
		"ᴀᴜᴛᴏᴛɪᴄᴋᴇᴛ [ɴᴜᴋᴇ/ᴏꜰꜰ]",
		"ᴀᴜᴛᴏʙᴄ [ᴍꜱɢ/ɪᴍɢ/ᴏꜰꜰ]",
		"ᴀᴜᴛᴏᴛʀᴀɴꜱ [ᴄᴏᴅᴇ/ᴏꜰꜰ]",
		"ᴀᴜᴛᴏʟɪᴋᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʙᴀᴄᴋᴜᴘ [ɪɴᴠ/Qʀ]",
		"ᴋɪᴄᴋʙᴀɴQʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴍᴏᴅᴇᴡᴀʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴍᴏᴅᴇᴘʀᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴏᴄᴋᴀᴊꜱ [ᴏɴ/ᴏꜰꜰ]",
		"ʙᴀᴄᴋᴜꜱᴇʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴄᴀɴᴄᴇʟᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴅᴇᴛᴇᴄᴛᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴍᴇᴅɪᴀᴅʟ [ᴏɴ/ᴏꜰꜰ]",
		"ʙᴏᴍʟɪᴋᴇ [ɴᴜᴍ/ᴏꜰꜰ]",
		"ꜰʟᴇx [1-2/ꜰᴏᴏᴛᴇʀ/ᴏꜰꜰ]",
		"ʟɪꜰꜰ [1-5]",
		"ʟɪꜱᴛᴄᴍᴅ",
		//"ʀᴇꜱᴇɴᴅ [ɴᴜᴍ]",
		"ᴛᴇxᴛᴍᴏᴅᴇ [1-10]",
		"ᴡᴏʀᴅʙᴀɴᴀᴅᴅ [ᴛᴇxᴛ]",
		"ᴡᴏʀᴅʙᴀɴᴅᴇʟ [ᴛᴇxᴛ]",
		"ᴡᴏʀᴅʙᴀɴʟɪꜱᴛ",
		"ᴡᴏʀᴅʙᴀɴᴄʟᴇᴀʀ",
		"ᴄᴏᴜɴᴛꜱᴘᴀᴍ [ɴᴜᴍ]",
		"ᴄᴏᴜɴᴛᴀᴊꜱ [ɴᴜᴍ]",
		"ɢʀᴏᴜᴘꜱ",
		"ɢʀᴏᴜᴘɪɴꜰᴏ [ɴᴜᴍ]",
		"ɢʀᴏᴜᴘᴄᴀꜱᴛ [ɴᴜᴍ]",
		"ᴍꜱɢʙᴄ [ᴛᴇxᴛ]",
		"ᴋɪᴄᴋᴀʟʟ",
		"ᴄᴀɴᴄᴇʟᴀʟʟ",
		"ᴄʜᴇᴄᴋ ᴀʟʟ",
		"ᴄʜᴇᴄᴋʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ᴍᴀꜱᴛᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴍᴀꜱᴛᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴍᴀꜱᴛᴇʀꜱ",
		"ᴄʟᴇᴀʀᴍᴀꜱᴛᴇʀ",
		"ᴏᴡɴᴇʀꜱ",
		"ʙᴏᴛ [@/ʟᴀꜱᴛ]",
		"ᴜɴʙᴏᴛ [@/ʟᴀꜱᴛ]",
		"ʙᴏᴛʟɪꜱᴛ",
		"ᴄʟᴇᴀʀʙᴏᴛ",
		"ᴘᴜʀɢᴇᴀʟʟ",
		"ᴘᴜʀɢᴇᴀʟʟʙᴀɴꜱ [ɴᴜᴍ]",
		"ɴᴜᴋᴇʙᴀɴ",
		"ɴᴜᴋᴇʙᴏᴛ",
		"ʀᴜɴᴛɪᴍᴇ",
		"ᴛɪᴍᴇɴᴏᴡ",
		"ꜱQᴜᴀᴅᴍɪᴅ",
		"ꜱᴇᴛʙᴏᴛ",
		"ꜱᴇᴛᴄᴍᴅ [ᴄᴍᴅ]",
		"ꜱᴛɪᴄᴋᴇʀᴄᴍᴅ",
		"ꜱᴛɪᴄᴋᴇʀᴄʟᴇᴀʀ",
		"ꜱᴛɪᴄᴋᴇʀᴋɪᴄᴋ",
		"ꜱᴛɪᴄᴋᴇʀɪɴᴠɪᴛᴇ",
		"ꜱᴛɪᴄᴋᴇʀᴋɪᴄᴋᴀʟʟ",
		"ꜱᴛɪᴄᴋᴇʀᴄᴀɴᴄᴇʟ",
		"ꜱᴛɪᴄᴋᴇʀʙʏᴘᴀꜱꜱ",
		"ꜱᴛɪᴄᴋᴇʀꜱᴛᴀʏᴀʟʟ",
		"ꜱᴛɪᴄᴋᴇʀʟᴇᴀᴠᴇ",
		"ꜱᴛɪᴄᴋᴇʀʀᴇꜱᴘᴏɴ",
		"ꜱᴛɪᴄᴋᴇʀᴄʟᴇᴀʀʙᴀɴ",
	}
	helpmaster = []string{
		"ᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ᴀᴅᴍɪɴꜱ",
		"ᴄʟᴇᴀʀᴀᴅᴍɪɴ",
		"ᴀᴜᴛᴏᴍᴜᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴀɴɴᴏᴜɴᴄᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴄʟᴇᴀʀʙᴀɴ",
		"ʙᴀɴꜱ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ꜰɪxᴇᴅ",
		"ʙʀɪɴɢ [ɴᴜᴍ]",
		"ᴄᴏɴᴛᴀᴄᴛ [@/ʟᴀꜱᴛ]",
		"ᴄᴏᴜɴᴛ",
		"ᴄᴜʀʟ",
		"ʜᴏꜱᴛᴀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴍᴀꜱᴛᴇʀꜱ",
		"ᴍᴜᴛᴇ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴍᴜᴛᴇ [@/ʟᴀꜱᴛ]",
		"ᴍᴜᴛᴇʟɪꜱᴛ",
		"ᴏᴜʀʟ",
		"ʀᴏʟʟᴄᴀʟʟ",
		"ꜱᴀʏᴀʟʟ",
		"ꜱᴇᴛɢʀᴏᴜᴘ",
		"ꜱᴛᴀʏᴀʟʟ",
		"ꜱᴛᴀʏ [ɴᴜᴍ]",
		"ʟᴇᴀᴠᴇ",
		"ꜱᴘᴇᴇᴅ",
		"ꜱᴛᴀᴛᴜꜱ",
		"ꜱᴛᴀᴛᴜꜱ ᴀᴅᴅ",
		"ʟɪᴍɪᴛᴏᴜᴛ",
		"ᴄʟᴇᴀʀɢᴏᴡɴᴇʀ",
		"ᴄʟᴇᴀʀɢᴀᴅᴍɪɴ",
		"ᴜɴꜱᴇɴᴅ",
		"ᴡʜᴏɪꜱ [@/ʟᴀꜱᴛ]",
	}
	helpadmin = []string{
		"ᴀᴅᴍɪɴꜱ",
		"ᴀʟʟᴘʀᴏᴛᴇᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪɴᴠɪᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴋɪᴄᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀɴᴄᴇʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏQʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢɴᴀᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢᴘɪᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɴᴏᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴀʟʙᴜᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏʟɪɴᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰʟᴇx [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪᴍᴀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴠɪᴅᴇᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴘᴀᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴛɪᴄᴋᴇʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴏɴᴛᴀᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴘᴏꜱᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰɪʟᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴛᴇᴄᴛᴍᴀx [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋ ɴᴀᴍᴇ",
		"ʟᴜʀᴋ ᴍᴇɴᴛɪᴏɴ",
		"ʟᴜʀᴋ ʜɪᴅᴇ",
		"ʟᴜʀᴋ ɪᴍᴀɢᴇ",
		"ʟᴜʀᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋꜱ",
		"ᴡᴇʟᴄᴏᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴇᴀᴠᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴍꜱɢʟᴜʀᴋ [ᴛᴇxᴛ]",
		"ᴍꜱɢʟᴇᴀᴠᴇ [ᴛᴇxᴛ]",
		"ᴍꜱɢᴡᴇʟᴄᴏᴍᴇ [ᴛᴇxᴛ]",
		"ꜱᴘᴀᴍᴄᴀʟʟᴛᴏ [@/ʟᴀꜱᴛ]",
		"ꜱᴘᴀᴍᴄᴀʟʟ [ɴᴜᴍ]",
		"ᴛɪᴋᴛᴏᴋ [ᴜꜱᴇʀ]",
		"ɪɴꜱᴛᴀɢʀᴀᴍ [ᴜꜱᴇʀ]",
		"ꜱᴍᴜʟᴇ [ᴜꜱᴇʀ]",
		"ʏᴏᴜᴛᴜʙᴇ [ᴛᴇxᴛ]",
		"ʏᴏᴜᴛᴜʙᴇᴅʟ [ᴜʀʟ]",
		"ᴊᴏᴏx [ᴛᴇxᴛ]",
		"ᴛᴇxᴛɪᴍᴀɢᴇ [ᴛᴇxᴛ]",
		"ᴄᴀʟᴄᴜʟᴀᴛᴏʀ [ᴛᴇxᴛ]",
		"ᴀʀᴛɪɴᴀᴍᴀ [ᴛᴇxᴛ]",
		"ꜱɪᴍɪ [ᴛᴇxᴛ]",
		"ᴄᴜᴀᴄᴀ [ᴛᴇxᴛ]",
		"ᴘɪɴᴛᴇʀᴇꜱᴛ [ᴛᴇxᴛ]",
		"ᴠɪᴅᴇᴏᴘᴏʀɴ [ᴛᴇxᴛ]",
		"ɢɪᴍᴀɢᴇ [ᴛᴇxᴛ]",
		"ʜᴇɴᴛᴀɪ",
		"ᴘᴏʀɴꜱᴛᴀʀᴛ",
		"ꜱᴇᴛɢʀᴏᴜᴘ",
		"ᴜɴꜱᴇɴᴅ",
		"ɢᴇᴛᴄᴀʟʟ",
		"ʙᴀɴꜱ",
		"ꜰɪxᴇᴅ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ʙᴀɴʟɪꜱᴛ",
		"ᴄʟᴇᴀʀʙᴀɴ",
		"ɢᴀᴄᴄᴇꜱꜱ",
		"ɢᴏᴡɴᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴɢᴏᴡɴᴇʀ [@/ʟᴀꜱᴛ]",
		"ɢᴏᴡɴᴇʀꜱ",
		"ᴄʟᴇᴀʀɢᴏᴡɴᴇʀ",
		"ɢᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴɢᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ɢᴀᴅᴍɪɴꜱ",
		"ᴄʟᴇᴀʀɢᴀᴅᴍɪɴ",
		"ɢʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴɢʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ɢʙᴀɴʟɪꜱᴛ",
		"ᴄʟᴇᴀʀɢʙᴀɴ",
		"ʜᴇʀᴇ",
		"ꜱᴛᴀʏ [ɴᴜᴍ]",
		"ꜱᴛᴀʏᴀʟʟ",
		"ʟᴇᴀᴠᴇ",
		"ꜱᴛᴀᴛᴜꜱ",
		"ꜱᴛᴀᴛᴜꜱ ᴀᴅᴅ",
		"ʟɪᴍɪᴛᴏᴜᴛ",
		"ɢᴏ [ɴᴜᴍ]",
		"ɢᴏᴊᴏɪɴ",
		"ᴋɪᴄᴋ [@/ʟᴀꜱᴛ]",
		"ᴠᴋɪᴄᴋ [@/ʟᴀꜱᴛ]",
		"ɪɴᴠɪᴛᴇ [@/ʟᴀꜱᴛ]",
		"ᴄᴀɴᴄᴇʟ [@/ʟᴀꜱᴛ]",
		"ɪᴍᴀɢᴇ [@/ʟᴀꜱᴛ]",
		"ʙɪᴏ [@/ʟᴀꜱᴛ]",
		"ᴄᴏᴠᴇʀ [@/ʟᴀꜱᴛ]",
		"ᴇxᴛʀᴀᴄᴏᴠᴇʀ [@/ʟᴀꜱᴛ]",
		"ꜱᴛᴏʀʏ [@/ʟᴀꜱᴛ]",
		"ᴍɪᴅ [@/ʟᴀꜱᴛ]",
		"ɴᴀᴍᴇ [@/ʟᴀꜱᴛ]",
		"ᴘɪɴɢ",
		"ᴘʀᴇꜰɪx",
		"ʀᴇꜱᴘᴏɴ",
		"ʀɴᴀᴍᴇ",
		"ꜱɴᴀᴍᴇ",
		"ʟᴀꜱᴛꜱᴇᴛ",
		"ꜱᴀʏ [ᴛᴇxᴛ]",
		"ᴛᴀɢ",
		"ᴛᴀɢᴀʟʟ",
	}
	helpgowner = []string{
		"ɢᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴɢᴀᴅᴍɪɴ [@/ʟᴀꜱᴛ]",
		"ɢᴀᴅᴍɪɴꜱ",
		"ᴄʟᴇᴀʀɢᴀᴅᴍɪɴ",
		"ɢʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ᴜɴɢʙᴀɴ [@/ʟᴀꜱᴛ]",
		"ɢʙᴀɴʟɪꜱᴛ",
		"ᴄʟᴇᴀʀɢʙᴀɴ",
		"ᴡᴇʟᴄᴏᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴇᴀᴠᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋ ɴᴀᴍᴇ",
		"ʟᴜʀᴋ ᴍᴇɴᴛɪᴏɴ",
		"ᴀʟʟᴘʀᴏᴛᴇᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪɴᴠɪᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴋɪᴄᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀɴᴄᴇʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏQʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢɴᴀᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢᴘɪᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɴᴏᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴀʟʙᴜᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏʟɪɴᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰʟᴇx [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪᴍᴀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴠɪᴅᴇᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴘᴀᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴛɪᴄᴋᴇʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴏɴᴛᴀᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴘᴏꜱᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰɪʟᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴛᴇᴄᴛᴍᴀx [ᴏɴ/ᴏꜰꜰ]",
		"ꜱᴇᴛɢʀᴏᴜᴘ",
		"ɢᴀᴄᴄᴇꜱꜱ",
		"ʟɪᴍɪᴛᴏᴜᴛ",
		"ꜱᴛᴀᴛᴜꜱ",
		"ꜱᴛᴀᴛᴜꜱ ᴀᴅᴅ",
		"ꜱᴘᴇᴇᴅ",
		"ɢᴏ [ɴᴜᴍ]",
		"ɢᴏᴊᴏɪɴ",
		"ᴛᴀɢᴀʟʟ",
	}
	helpgadmin = []string{
		"ᴡᴇʟᴄᴏᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴇᴀᴠᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ʟᴜʀᴋ ɴᴀᴍᴇ",
		"ʟᴜʀᴋ ᴍᴇɴᴛɪᴏɴ",
		"ᴀʟʟᴘʀᴏᴛᴇᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪɴᴠɪᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴋɪᴄᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀɴᴄᴇʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴊᴏɪɴ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏQʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢɴᴀᴍᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɢᴘɪᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɴᴏᴛᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴀʟʙᴜᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏʟɪɴᴋ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰʟᴇx [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏɪᴍᴀɢᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴠɪᴅᴇᴏ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴀʟʟ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴘᴀᴍ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜱᴛɪᴄᴋᴇʀ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴄᴏɴᴛᴀᴄᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴘᴏꜱᴛ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏꜰɪʟᴇ [ᴏɴ/ᴏꜰꜰ]",
		"ᴘʀᴏᴛᴇᴄᴛᴍᴀx [ᴏɴ/ᴏꜰꜰ]",
		"ꜱᴇᴛɢʀᴏᴜᴘ",
		"ɢᴀᴄᴄᴇꜱꜱ",
		"ᴛᴀɢᴀʟʟ",
	}
	details = map[string]string{
		"shutdown":     "'%s%s'\n\nShutting down the bot's.",
		"perm":         "'%s%s .<grade>.<command>'\n\nAvailable grade buyer/owner/master/admin",
		"nukejoin":     "'%s%s' on/off\nkickall member's while bot has invited..",
		"announce":     "'%s%s on/off'\n\nEnable detect announce.",
		"hostage":      "'%s%s on/off'\n\nEnable auto invite leave member.",
		"accept":       "'%s%s <number>'\n\nAccept group invitation by number.",
		"reject":       "'%s%s <number>'\n\nReject group invitation by number.",
		"welcome":      "'%s%s on/off'\n\nEnable welcome message.",
		"leave":        "'%s%s on/off'\n\nEnable leave message.",
		"setcmd":       "'%s%s <state> <command>'\n\nUsed to enabling/disabling command\nAvailable state lock/unlock/disable/enable.",
		"fixed":        "'%s%s'\n\nIf bot's error, please use this command to autofix.",
		"notification":      "'%s%s <state>'\n\nUsed to see bot's activity.\nAvailable state on/off",
		"go":           "'%s%s <number>'\n\nSet bot to stay on group invitation.\nDefault is 2 bot.",
		"unseller":     "'%s%s <range/lcon/lkick/etc>'Used to expel seller.\nAvailable range '<', '>', '-', ',' with number.",
		"unbuyer":      "'%s%s <range/lcon/lkick/etc>'Used to expel buyer.\nAvailable range '<', '>', '-', ',' with number.",
		"unowner":      "'%s%s <range/lcon/lkick/etc>'Used to expel owner.\nAvailable range '<', '>', '-', ',' with number.",
		"unadmin":      "'%s%s <range/lcon/lkick/etc>'Used to expel admin.\nAvailable range '<', '>', '-', ',' with number.",
		"unmaster":     "'%s%s <range/lcon/lkick/etc>'Used to expel master.\nAvailable range '<', '>', '-', ',' with number.",
		"ungowner":     "'%s%s <range/lcon/lkick/etc>'Used to expel gowner.\nAvailable range '<', '>', '-', ',' with number.",
		"ungadmin":     "'%s%s <range/lcon/lkick/etc>'Used to expel gadmin.\nAvailable range '<', '>', '-', ',' with number.",
		"clearseller":  "'%s%s'\n\nClearing all sellers.",
		"clearbuyer":   "'%s%s'\n\nClearing all buyer list.",
		"clearowner":   "'%s%s'\n\nClearing all owner list.",
		"clearmaster":  "'%s%s'\n\nClearing all master list.",
		"clearadmin":   "'%s%s'\n\nClearing all admin list.",
		"cleargadmin":  "'%s%s'\n\nClearing all gadmin list.",
		"cleargowner":  "'%s%s'\n\nClearing all gowner list.",
		"clearbot":     "'%s%s'\n\nClearing all bot list.",
		"clearban":     "'%s%s'\n\nClearing all ban list.",
		"clearfuck":    "'%s%s'\n\nClearing all fuck list.",
		"clearmute":    "'%s%s'\n\nClearing all mute list.",
		"cleargban":    "'%s%s'\n\nClearing all gban list.",
		"clears":    "'%s%s'\n\nClearing all squad messages.",
		"upvallimage":  "'%s%s'\n\nUpdating all bot's video profile.",
		"upvimage":     "'%s%s'\n\nUpdating all bot's video profile.",
		"upallimage":   "'%s%s'\n\nUpdating all bot's picture profile.",
		"upimage":      "'%s%s'\n\nUpdating bot's profile picture.",
		"upvallcover":  "'%s%s'\n\nUpdating all bot's video cover.",
		"upvcover":     "'%s%s @tag bot'\n\nUpdating bot's video cover.",
		"upcover":      "'%s%s' @tag\n\nUpdating bot's cover picture.",
		"upallcover":   "'%s%s'\n\nUpdating all bot's cover picture.",
		"upname":       "'%s%s newname'\n\nUpdating bot's displayname.",
		"upallname":    "'%s%s newname'\n\nUpdating all bot's displayname.",
		"leaveall":     "'%s%s'\n\nleave all bot's from all group's.",
		"groups":       "'%s%s'\n\nsee bot group's.",
		"stayall":      "'%s%s'\n\naccepting all group invitation.",
		"setcom":       "'%s%s .key .value'\n\nChange command.",
		"upstatus":     "'%s%s <status message>'\n\nUpdating bot's profile bio.",
		"upallstatus":  "'%s%s <status message>'\n\nUpdating all bot's profile bio.",
		"kick":         "'%s%s @tag/lcon/lkick/etc'\n\nKick member's.",
		"prefix":       "'%s%s on/off'\n\nEnable/disable prefix.",
		"list protect": "'%s%s'\n\nShow all protection group's.",
		"invme":        "'%s%s gnumber'\n\nInvite user to the destination group.",
		"autojoin":     "'%s%s qr/inv/off'\n\nForcing bot's to joinall while invited.",
		"ajsjoin":     "'%s%s qr/inv/off'\n\nForcing bot's to joinall while invited.",
		"autoban":      "'%s%s on/off'\n\nAuto banned user.",
		"sellers":      "'%s%s'\n\nShow seller list.",
		"buyers":       "'%s%s'\n\nShow buyer list.",
		"owners":       "'%s%s'\n\nShow owner list.",
		"masters":      "'%s%s'\n\nShow master list.",
		"admins":       "'%s%s'\n\nShow admin list.",
		"gowners":      "'%s%s'\n\nShow gowner list.",
		"gadmins":      "'%s%s'\n\nShow gadmin list.",
		"botlist":      "'%s%s'\n\nShow bot list.",
		"banlist":      "'%s%s'\n\nShow ban list.",
		"fucklist":     "'%s%s'\n\nShow fuck list.",
		"mutelist":     "'%s%s'\n\nShow mutelist list.",
		"gbanlist":     "'%s%s'\n\nShow gban list.",
		"hides":        "'%s%s'\n\nShow Invisible user.",
		"hide":         "'%s%s @tag/lcon/lkick/etc'\n\nAdded user to invisible list.",
		"kickall":      "'%s%s'\n\nKick all group member's.",
		"group info":   "'%s%s'\n\nShow all group member's./pendings/access",
		"autopurge":    "'%s%s on/off'\n\nEnable autopurge.",
		"lurk":         "'%s%s on/off'\n\nEnable lurking mode.",
		"lurkmsg":      "'%s%s <message>'\n\nSet lurk message.\nUse @! for placing user tagging.",
		"antitag":      "'%s%s on/off'\n\nEnable antitag.",
		"killmode":     "'%s%s kill/purge/on/off/range'\n\nKiller mode to kick all banlist/squad.",
		"autopro":      "'%s%s on/off'\n\nAuto protect max while bot's join.",
		"setlimit":     "'%s%s number'\n\nSet max kick in killmode /bot.",
		"stay":         "'%s%s number'\n\nSet amount of bot's in group invite via link invitation.",
		"bringall":     "'%s%s'\n\nBring all bot's by invitation.",
		"bring":        "'%s%s number'\n\nSet amount of bot's in group via invitation.",
		"here":         "'%s%s'\n\nShow amount of bot's in group.",
		"friends":      "'%s%s'\n\nShow all bot's friends.",
		"msgrespon":    "'%s%s respon'\n\nSet bot's response.",
		"msgwelcome":   "'%s%s <message>'\n\nSet welcome message each group.\nParameter for changing need to adding @user for replacing username and @group for replacing groupname.",
		"setrname":     "'%s%s newrname'\n\nChange the rname prefix.",
		"setsname":     "'%s%s newsname'\n\nChange the sname prefix.",
		"invite":       "'%s%s @tag/lcon/lkick/etc'\n\nInvite target to the group's.",
		"clone":        "'%s%s @tag/lcon/lkick/etc @tagbot'\n\nCloning targte profile.",
		"gaccess":      "'%s%s'\n\nSee all group access list.",
		"limitout":     "'%s%s'\n\nLeave the kicbanned bot's.",
		"say":          "'%s%s word'\n\nThe bot's would said the word.",
		"sayall":       "'%s%s word'\n\nAll bot's would said the word.",
		"expel":        "'%s%s @tag/lcon/lkick/etc'\n\nUsed to expel user access.",
		"respon":       "'%s%s'\n\nBot response.",
		"ping":         "'%s%s'\n\nBot response.",
		"permlist":     "'%s%s key'\n\nGet the command value.",
		"setgroup":     "'%s%s'\n\nShow the group preset status in group.",
		"setbot":          "'%s%s'\n\nShow the bot's set.",
		"help":         "'%s%s'\n\nShow the help command.",
		"deny":         "'%s%s invite/kick/qr/join/cancel/off/all/max'\n\nEnable the protection.",
		"allow":        "'%s%s invite/kick/qr/join/cancel/all'\n\nDisable the protection.",
		"ourl":         "'%s%s'\n\nOpen group links.",
		"curl":         "'%s%s'\n\nClose group links.",
		"mysquad":      "'%s%s'\n\nSend squad contact's",
		"count":        "'%s%s'\n\nShow bot's number.",
		"speed":        "'%s%s'\n\nShow bot response speed.",
		"unsend":       "'%s%s count'\n\nUnsend recent bot's message.\nIf count not definde, it would unsend all recent message.",
		"tagall":       "'%s%s'\n\nTagging all member's.",
		"ftagall":      "'%s%s'\n\nTagging all member's with sticker.",
		"access":       "'%s%s'\n\nShow all bot access.",
		"bans":         "'%s%s'\n\nShow the bot's status.",
		"runtime":      "'%s%s'\n\nShiw the bot's time alive.",
		"timeleft":     "'%s%s'\n\nShow the bot's timeleft.",
		"linvite":      "'%s%s'\n\nShow the last invited in group.",
		"lkick":        "'%s%s'\n\nShow the last kicked in group.",
		"lmid":         "'%s%s'\n\nShow the last mid in group.",
		"lcon":         "'%s%s'\n\nShow the last contact in group.",
		"ltag":         "'%s%s'\n\nShow the last tag in group.",
		"lban":         "'%s%s'\n\nShow the last banned in group.",
		"lcancel":      "'%s%s'\n\nShow the last cancel in group.",
		"lqr":          "'%s%s'\n\nShow the last upded qr in group.",
		"ljoin":        "'%s%s'\n\nShow the last join in group.",
		"lleave":       "'%s%s'\n\nShow the last leave in group.",
		"abort":        "'%s%s'\n\nAborting command.",
		"groupcast":    "'%s%s <your word>'\n\nBroadcasting message to all groups.",
		"contact":      "'%s%s @tag/lcon/lkick/etc'\n\nUsed to get contact's.",
		"rollcall":     "'%s%s'\n\nShow bot's name.",
		"gojoin":       "'%s%s'\n\nJoining bot's from invitation list.",
		"mid":          "'%s%s @tag/lcon/lkick/etc'\n\nGet midlist.",
		"name":         "'%s%s @tag/lcon/lkick/etc'\n\nGet namelist.",
		"purgeall":     "'%s%s'\n\nPurge all banlist in all group.",
		"squadmid":     "'%s%s'\n\nShow all bots mid.",
		"whois":        "'%s%s @tag/lcon/lkick/etc'\n\nSee user info.",
		"cancel":       "'%s%s @tag/lcon/lkick/etc'\n\nCancel group invitation.",
		"remotegroup":       "'%s%s:'\n\nthe right number\nSee group number with command groups.\nExample:\n  remote: 2 gmember.\nund send command.",
	}
)
type (
	mentions struct {
		MENTIONEES []struct {
			Start string `json:"S"`
			End   string `json:"E"`
			Mid   string `json:"M"`
		} `json:"MENTIONEES"`
	}
	kickop struct {
		Kick  []string
		Inv   []string
		Opinv []int64
	}
	emots struct {
		STICON struct {
			RESOURCES [] struct {
				PRODUCTID string `json:"productId"`
				STICONID string `json:"sticonId"`
			}`json:"resources"`
		}`json:"sticon"`
	}
	Stickers struct {
		Id  string
		Pid string
	}
	clustering struct {
		mem string
		tm  int64
		fr  []string
	}
	DATA struct {
		Authoken      []string             `json:"Authoken"`
		CreatorBack   []string             `json:"CreatorBack"`
		MakerBack   []string             `json:"MakerBack"`
		BuyerBack     []string             `json:"BuyerBack"`
		OwnerBack     []string             `json:"OwnerBack"`
		MasterBack    []string             `json:"MasterBack"`
		AdminBack     []string             `json:"AdminBack"`
		ResponBack    string               `json:"ResponBack"`
		BroadcastBack    string               `json:"BroadcastBack"`
		RnameBack     string               `json:"RnameBack"`
		SnameBack     string               `json:"SnameBack"`
		BotBack       []string             `json:"BotBack"`
		Dalltime      string               `json:"Dalltime"`
		Logobot       string               `json:"Logobot"`
		SellerBack    []string             `json:"SellerBack"`
		BanBack       []string             `json:"BanBack"`
		FuckBack      []string             `json:"FuckBack"`
		LockBack       []string             `json:"LockBack"`
		WordbanBack      []string             `json:"WordbanBack"`
		Limit         string               `json:"MLimit"`
		Fresh         string               `json:"MFfresh"`
		MuteBack      []string             `json:"MuteBack"`
		AnnunceBack   []string             `json:"AnnunceBack"`
		ProNameBack   []string             `json:"ProNameBack"`
		ProPictureBack[]string             `json:"ProPictureBack"`
		ProNoteBack	  []string             `json:"ProNoteBack"`
		ProAlbumBack  []string             `json:"ProAlbumBack"`
		ProQrBack     []string             `json:"ProQrBack"`
		ProjoinBack   []string             `json:"ProjoinBack"`
		ProInviteBack []string             `json:"ProInviteBack"`
		ProCancelBack []string             `json:"ProCancelBack"`
		ProkickBack   []string             `json:"ProkickBack"`
		GbanBack      map[string][]string  `json:"GbanBack"`
		GadminBack    map[string][]string  `json:"GadminBack"`
		GownerBack    map[string][]string  `json:"GownerBack"`
		TimeBanBack   map[string]time.Time `json:"TimeBanBack"`
		ProLinkBack  []string             `json:"ProLinkBack"`
		ProFlexBack  []string             `json:"ProFlexBack"`
		ProImageBack  []string             `json:"ProImageBack"`
		ProVideoBack  []string             `json:"ProVideoBack"`
		ProCallBack  []string             `json:"ProCallBack"`
		ProSpamBack  []string             `json:"ProSpamBack"`
		ProStickerBack  []string             `json:"ProStickerBack"`
		ProContactBack  []string             `json:"ProContactBack"`
		ProPostBack  []string             `json:"ProPostBack"`
		ProFileBack  []string             `json:"ProFileBack"`
		Kikhistory       int             `json:"Kikhistory"`
		Invhistory       int             `json:"Invhistory"`
		Canclhistory       int             `json:"Canclhistory"`
		Maxkick       int             `json:"Maxkick"`
		Maxcancel       int             `json:"Maxcancel"`
		Maxinvite       int             `json:"Maxinvite"`
		Cancelpend       int             `json:"Cancelpend"`
		AutoproBack    bool                 `json:"AutoproBack"`
		RestartBack     string          `json:"RestartBack"`
		AutoPurgeBack  bool                 `json:"AutoPurgeBack"`
		ProtectmodeBack   bool               `json:"ProtectmodeBack"`
		PowermodeBack  bool                 `json:"PowermodeBack"`
		KickbanqrBack  bool                 `json:"KickbanqrBack"`
		MediadlBack  bool                 `json:"MediadlBack"`
		AutolikeBack  bool                 `json:"AutolikeBack"`
		AutobcBack  bool                 `json:"AutobcBack"`
		AutojointicketBack  bool                 `json:"AutojointicketBack"`
		AutotranslateBack  bool                 `json:"AutotranslateBack"`
		NukejoinBack  bool                 `json:"NukejoinBack"`
		CanceljoinBack  bool                 `json:"CanceljoinBack"`
		DetectcallBack  bool                 `json:"DetectcallBack"`
		ModebackupBack    string               `json:"ModebackupBack"`
		AutojoinBack    string               `json:"AutojoinBack"`
		AjsjoinBack    string               `json:"AjsjoinBack"`
		TypejoinBack    string               `json:"TypejoinBack"`
		TypebcBack    string               `json:"TypebcBack"`
		TypetransBack    string               `json:"TypetransBack"`

		KickSticker struct{
			Stkid string `json:"stkid"`
			Stkpkgid string `json:"stkpkgid"`
		}`json:"Command-sticker-kick"`
		ResponSticker struct{
			Stkid2 string `json:"stkid2"`
			Stkpkgid2 string `json:"stkpkgid2"`
		}`json:"Command-sticker-respon"`
		StayallSticker struct{
			Stkid3 string `json:"stkid3"`
			Stkpkgid3 string `json:"stkpkgid3"`
		}`json:"Command-sticker-stayall"`
		LeaveSticker struct{
			Stkid4 string `json:"stkid4"`
			Stkpkgid4 string `json:"stkpkgid4"`
		}`json:"Command-sticker-leave"`
		KickallSticker struct{
			Stkid5 string `json:"stkid5"`
			Stkpkgid5 string `json:"stkpkgid5"`
		}`json:"Command-sticker-kickall"`
		BypassSticker struct{
			Stkid6 string `json:"stkid6"`
			Stkpkgid6 string `json:"stkpkgid6"`
		}`json:"Command-sticker-bypass"`
		InviteSticker struct{
			Stkid7 string `json:"stkid7"`
			Stkpkgid7 string `json:"stkpkgid7"`
		}`json:"Command-sticker-invite"`
		ClearbanSticker struct{
			Stkid8 string `json:"stkid8"`
			Stkpkgid8 string `json:"stkpkgid8"`
		}`json:"Command-sticker-clearban"`
		CancelallSticker struct{
			Stkid9 string `json:"stkid9"`
			Stkpkgid9 string `json:"stkpkgid9"`
		}`json:"Command-sticker-cancelall"`
	}
)

func Resprem() {
	rngcmd := GetComs(5, "clone")
	rngcmd = GetComs(5, "nukeqr")
	rngcmd = GetComs(5, "joinqr")
	rngcmd = GetComs(5, "cancelall")
	rngcmd = GetComs(5, "kickall")
	rngcmd = GetComs(8, "none")
	rngcmd = GetComs(8, "max")
	rngcmd = GetComs(7, "allowall")
	rngcmd = GetComs(7, "denyall")
	rngcmd = GetComs(6, "hostage")
	rngcmd = GetComs(5, "backup")
	rngcmd = GetComs(6, "upgname")
	rngcmd = GetComs(6, "welcome")
	rngcmd = GetComs(5, "sendimage")
	rngcmd = GetComs(5, "leave")
	rngcmd = GetComs(6, "announce")
	rngcmd = GetComs(7, "unban")
	rngcmd = GetComs(7, "bio")
	rngcmd = GetComs(7, "tag")
	rngcmd = GetComs(7, "image")
	rngcmd = GetComs(7, "contact")
	rngcmd = GetComs(7, "ban")
	rngcmd = GetComs(7, "kick")
	rngcmd = GetComs(7, "vkick")
	rngcmd = GetComs(7, "invite")
	rngcmd = GetComs(7, "cancel")
	rngcmd = GetComs(8, "ungban")
	rngcmd = GetComs(5, "unbot")
	rngcmd = GetComs(9, "tagall")
	rngcmd = GetComs(5, "statusall")
	rngcmd = GetComs(6, "status")
	rngcmd = GetComs(6, "whois")
	rngcmd = GetComs(6, "mute")
	rngcmd = GetComs(5, "fuck")
	rngcmd = GetComs(5, "setlimiter")
	rngcmd = GetComs(5, "setcancel")
	rngcmd = GetComs(5, "setkick")
	rngcmd = GetComs(5, "setinvite")
	rngcmd = GetComs(5, "msgfresh")
	rngcmd = GetComs(5, "msglimit")
	rngcmd = GetComs(5, "msgstatus")
	rngcmd = GetComs(7, "msglurk")
	rngcmd = GetComs(5, "msgclearban")
	rngcmd = GetComs(7, "msgleave")
	rngcmd = GetComs(8, "speed")
	rngcmd = GetComs(9, "lurk")
	rngcmd = GetComs(7, "msgwelcome")
	rngcmd = GetComs(5, "msgrespon")
	rngcmd = GetComs(5, "setrname")
	rngcmd = GetComs(5, "setsname")
	rngcmd = GetComs(5, "notification")
	rngcmd = GetComs(5, "killmode")
	rngcmd = GetComs(5, "unowner")
	rngcmd = GetComs(7, "name")
	rngcmd = GetComs(5, "Stats")
	rngcmd = GetComs(5, "buyers")
	rngcmd = GetComs(5, "upname")
	rngcmd = GetComs(5, "upstatus")
	rngcmd = GetComs(5, "acceptall")
	rngcmd = GetComs(5, "declineall")
	rngcmd = GetComs(7, "abort")
	rngcmd = GetComs(5, "accept")
	rngcmd = GetComs(5, "decline")
	rngcmd = GetComs(5, "invme")
	rngcmd = GetComs(5, "gleave")
	rngcmd = GetComs(5, "Purgeallbans")
	rngcmd = GetComs(5, "purgeall")
	rngcmd = GetComs(7, "unsend")
	rngcmd = GetComs(2, "makers")
	rngcmd = GetComs(5, "upvcover")
	rngcmd = GetComs(2, "unseller")
	rngcmd = GetComs(2, "clearseller")
	rngcmd = GetComs(3, "sellers")
	rngcmd = GetComs(2, "seller")
	rngcmd = GetComs(1, "unmaker")
	rngcmd = GetComs(1, "clearmaker")
	rngcmd = GetComs(5, "upvimage")
	rngcmd = GetComs(5, "upcover")
	rngcmd = GetComs(5, "upimage")
	rngcmd = GetComs(3, "clearbuyer")
	rngcmd = GetComs(3, "unbuyer")
	rngcmd = GetComs(3, "buyer")
	rngcmd = GetComs(8, "gaccess")
	rngcmd = GetComs(5, "allbanlist")
	rngcmd = GetComs(5, "access")
	rngcmd = GetComs(7, "expel")
	rngcmd = GetComs(5, "listcmd")
	rngcmd = GetComs(4, "owner")
	rngcmd = GetComs(4, "hide")
	rngcmd = GetComs(4, "unhide")
	rngcmd = GetComs(4, "hidelist")
	rngcmd = GetComs(4, "clearhide")
	rngcmd = GetComs(7, "mid")
	rngcmd = GetComs(7, "cleargowner")
	rngcmd = GetComs(5, "notification")
	rngcmd = GetComs(4, "clearowner")
	rngcmd = GetComs(5, "unmaster")
	rngcmd = GetComs(6, "unmute")
	rngcmd = GetComs(4, "clearlistcmd")
	rngcmd = GetComs(5, "setcmd")
	rngcmd = GetComs(7, "gowner")
	rngcmd = GetComs(5, "master")
	rngcmd = GetComs(8, "gojoin")
	rngcmd = GetComs(7, "ungowner")
	rngcmd = GetComs(9, "setgroup")
	rngcmd = GetComs(6, "setbot")
	rngcmd = GetComs(5, "runtime")
	rngcmd = GetComs(5, "timenow")
	rngcmd = GetComs(4, "timeleft")
	rngcmd = GetComs(9, "say")
	rngcmd = GetComs(6, "curl")
	rngcmd = GetComs(6, "ourl")
	rngcmd = GetComs(9, "here")
	rngcmd = GetComs(8, "gbanlist")
	rngcmd = GetComs(5, "clearcache")
	rngcmd = GetComs(5, "clears")
	rngcmd = GetComs(5, "cleargban")
	rngcmd = GetComs(5, "clearbot")
	rngcmd = GetComs(5, "botlist")
	rngcmd = GetComs(7, "bans")
	rngcmd = GetComs(7, "fixed")
	rngcmd = GetComs(8, "gban")
	rngcmd = GetComs(5, "bot")
	rngcmd = GetComs(5, "stay")
	rngcmd = GetComs(5, "leaveall")
	rngcmd = GetComs(8, "go")
	rngcmd = GetComs(7, "stayall")
	rngcmd = GetComs(5, "bringall")
	rngcmd = GetComs(6, "listprotect")
	rngcmd = GetComs(7, "cleargadmin")
	rngcmd = GetComs(7, "clearban")
	rngcmd = GetComs(5, "clearadmin")
	rngcmd = GetComs(5, "upallname")
	rngcmd = GetComs(5, "upallstatus")
	rngcmd = GetComs(8, "limitout")
	rngcmd = GetComs(6, "sayall")
	rngcmd = GetComs(7, "count")
	rngcmd = GetComs(9, "ping")
	rngcmd = GetComs(7, "leave")
	rngcmd = GetComs(2, "addallsquads")
	rngcmd = GetComs(3, "addallbots")
	rngcmd = GetComs(5, "limits")
	rngcmd = GetComs(4, "adds")
	rngcmd = GetComs(5, "friends")
	rngcmd = GetComs(5, "upvallcover")
	rngcmd = GetComs(5, "upvallimage")
	rngcmd = GetComs(6, "unsend")
	rngcmd = GetComs(5, "upallcover")
	rngcmd = GetComs(5, "upallimage")
	rngcmd = GetComs(6, "rollcall")
	rngcmd = GetComs(7, "respon")
	rngcmd = GetComs(7, "banlist")
	rngcmd = GetComs(5, "antitag")
	rngcmd = GetComs(7, "admins")
	rngcmd = GetComs(8, "gadmin")
	rngcmd = GetComs(5, "squadmid")
	rngcmd = GetComs(8, "ungadmin")
	rngcmd = GetComs(6, "unadmin")
	rngcmd = GetComs(6, "masters")
	rngcmd = GetComs(6, "gowners")
	rngcmd = GetComs(6, "admin")
	rngcmd = GetComs(5, "unfuck")
	rngcmd = GetComs(4, "remotegroup")
	rngcmd = GetComs(5, "groupinfo")
	rngcmd = GetComs(5, "banpurge")
	rngcmd = GetComs(5, "autoban")
	rngcmd = GetComs(5, "autopurge")
	rngcmd = GetComs(5, "canceljoin")
	rngcmd = GetComs(5, "nukejoin")
	rngcmd = GetComs(5, "groups")
	rngcmd = GetComs(5, "gourl")
	rngcmd = GetComs(5, "groupcast")
	rngcmd = GetComs(5, "fucklist")
	rngcmd = GetComs(6, "mutelist")
	rngcmd = GetComs(5, "autojoin")
	rngcmd = GetComs(5, "ajsjoin")
	rngcmd = GetComs(4, "perm")
	rngcmd = GetComs(4, "permlist")
	rngcmd = GetComs(4, "clearallprotect")
	rngcmd = GetComs(4, "clearmute")
	rngcmd = GetComs(4, "clearfuck")
	rngcmd = GetComs(5, "clearmaster")
	rngcmd = GetComs(0, "creator")
	rngcmd = GetComs(0, "clearcreator")
	rngcmd = GetComs(0, "uncreator")
	rngcmd = GetComs(0, "creators")
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
	res, err := http.Post(hosts+"rm_LetterSealing="+apikey, "application/json; charset=UTF-8", requestBody)
	if err != nil {
		fmt.Println("Disable Letter Sealing Gagal")
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("Disable Letter Sealing Success")
	}
}
func sendBigImage(AuthToken string, grp string, imglink string) {
	cmd, _ := exec.Command("python3","libliff/liff.py",AuthToken,grp,imglink).Output()
	fmt.Println("\033[33m"+string(cmd)+"\033[39m")
}
func detectSquad(client *linetcr.Account, to, pelaku string) {
	chat := client.GetChat([]string{to}, true, true)
	if chat == nil { return }
	memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
	var createdTime int64
	for mid, tt := range memberMids {
		if pelaku == mid {
			createdTime = tt
			break
		}
	}
	for mid, tt := range memberMids {
		ct := float64(createdTime/1000 - tt/1000)
		if valid.Abs(ct) <= 1000 {
			if MemUser(to, mid) {
				Banned.AddEx(mid)
			}
		}
	}
}
func main() {
	if len(os.Args) < 1 {
		fmt.Println("do not forget arguement")
		os.Exit(1)
	}
	defer ants.Release()
	defer linetcr.PanicOnly()
	debug.SetGCPercent(500)
	cpu = runtime.NumCPU()
	jsonFile, err := os.Open(DATABASE)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Data)
	IPServer = fmt.Sprintf("%v", GetIP())
	fmt.Println("\n[ START RUN GOLANG ]")
	mod.LoginBase()
	mod.TMoreCompact()
	go gracefulShutdown()
	for no, tok := range Data.Authoken {
		time.Sleep(250 * time.Millisecond)
		sort := rand.Intn(9999-1000) + 1000
		app := fmt.Sprintf("ANDROID\t14.10.0\tAndroid OS\t13.0.%v", sort)
		mids := strings.Split(tok, ":")
		mid := mids[0]
		var ua = fmt.Sprintf("Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)")
		cl, err := linetcr.CreateNewLogin(tok, no, mid, app, ua, HostName[0])
		if err == nil {
			fmt.Println("\n\n  ❏ Name : " + cl.Namebot + "\n  ❏ Mid : " + cl.MID + "\n  ❏ Location : " + cl.Locale + "\n  ❏ Bots No: " + fmt.Sprintf("%v", no+1))
			if linetcr.IsFriends(cl, DEVELOPER[0]) == false {cl.FindAndAddContactsByMidV2(DEVELOPER[0]);time.Sleep(1 * time.Second)}
			cl.LoadPrimaryE2EEKeys()
			ClientBot = append(ClientBot, cl)
			ClientMid[cl.MID] = cl
			Squadlist = append(Squadlist, cl.MID)
			err2 := cl.DisableE2ee()
			if err2 != nil {
				fmt.Println(err2)
			}
			aaa, err3 := cl.GetE2EEPublicKeys()
			if err3 == nil {
				for _, aa := range aaa {
					cl.RemoveE2EEPublicKey(aa)
				}
			} else {
				fmt.Println(err3)
			}
			r, _ := cl.GetHomeProfile(cl.MID)
			if linetcr.GetBannedChat(r) == 1 {
				linetcr.BanChatAdd(cl)
				fmt.Println("  ❏ Status : Banned")
			} else {
				fmt.Println("  ❏ Status : Normal")
			}
		} else {
			rs := err.Error()
			if strings.Contains(rs, "INTERNAL_ERROR") || strings.Contains(rs, "AUTHENTICATION_FAILED") {
				fmt.Println("\n  ❏ Status : Limited" + "\n  ❏ Mid : " + mid + "\n  ❏ Bots No: " + fmt.Sprintf("%v", no+1))
				cl.MID = mid
				cl.Limited = true
			} else {
				logs := fmt.Sprintf("\n\n  ❏ No: %v ERROR: %s", no+1, err)
				fmt.Println(logs)
			}
		}
	}
	for m := range linetcr.HashToMap(linetcr.GetBlock) {
		if !InArray2(Squadlist, m) {
			linetcr.GetBlock.Del(m)
		}
	}
	ch := make(chan int, len(ClientBot))
	if len(ClientBot) != 0 {
		acl := len(ClientBot)
		for x := 0; x < acl; x++ {
			cc := x
			cla := ClientBot[cc]
			runtime.Gosched()
			go RunBot(cla, ch)
		}
		list := append([]*linetcr.Account{}, ClientBot...)
		sort.Slice(list, func(i, j int) bool {
			return list[i].KickCount < list[j].KickCount
		})
		for i, cl := range list {
			kk := i * 30
			cl.KickPoint = kk
			ko := i * 10
			cl.CustomPoint = ko
		}
		Resprem()
		for i := range ClientBot {
			for _, x := range Squadlist {
				if !InArray2(ClientBot[i].Squads, x) && x != ClientBot[i].MID {
					ClientBot[i].Squads = append(ClientBot[i].Squads, x)
				}
			}
		}
		ClientBot[0].SendMessage(DEVELOPER[0], ClientBot[0].Namebot)
		if Data.RestartBack != "" {
			_, memlist, _ := ClientBot[0].GetChatList(Data.RestartBack)
			if len(memlist) != 0 {
				for _, mid := range memlist {
					if InArray2(Squadlist, mid) {
						cl := GetKorban(mid)
						if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
							cl.SendMessage(Data.RestartBack, "Done Reboot ♪♪\n")
							Data.RestartBack = ""
							break
						}
					}
				}
			}
			Data.RestartBack = ""
		}
		go func() {
			for {
				autoset()
				time.Sleep(3 * time.Second)
			}
		}()
		for v := range ch {
			if v == 69 {
				break
			}
		}
		fmt.Println("HAVE FUN")
	}
}


//DONE
func BackSeave() {
	DetectCall = Data.DetectcallBack
	AutoPro = Data.AutoproBack
	AutoPurge = Data.AutoPurgeBack
	ProtectMode = Data.ProtectmodeBack
	PowerMode = Data.PowermodeBack
	KickBanQr = Data.KickbanqrBack
	MediaDl = Data.MediadlBack
	AutoLike = Data.AutolikeBack
	AutoBc = Data.AutobcBack
	NukeJoin = Data.NukejoinBack
	canceljoin = Data.CanceljoinBack
	AutoJointicket = Data.AutojointicketBack
	AutoTranslate = Data.AutotranslateBack
	ModeBackup = Data.ModebackupBack
	Autojoin = Data.AutojoinBack
	Ajsjoin = Data.AjsjoinBack
	TypeJoin = Data.TypejoinBack
	Typebc = Data.TypebcBack
	TypeTrans = Data.TypetransBack
	MaxKick = Data.Maxkick
	MaxCancel = Data.Maxcancel
	MaxInvite = Data.Maxinvite
	CancelPend = Data.Cancelpend
	TimeBackup = time.Time{}
	MsSname = Data.SnameBack
	MsRname = Data.RnameBack
	MsgRespon = Data.ResponBack
	stkid = Data.KickSticker.Stkid
	stkpkgid = Data.KickSticker.Stkpkgid
	stkid2 = Data.ResponSticker.Stkid2
	stkpkgid2 = Data.ResponSticker.Stkpkgid2
	stkid3 = Data.StayallSticker.Stkid3
	stkpkgid3 = Data.StayallSticker.Stkpkgid3
	stkid4 = Data.LeaveSticker.Stkid4
	stkpkgid4 = Data.LeaveSticker.Stkpkgid4
	stkid5 = Data.KickallSticker.Stkid5
	stkpkgid5 = Data.KickallSticker.Stkpkgid5
	stkid6 = Data.BypassSticker.Stkid6
	stkpkgid6 = Data.BypassSticker.Stkpkgid6
	stkid7 = Data.InviteSticker.Stkid7
	stkpkgid7 = Data.InviteSticker.Stkpkgid7
	stkid8 = Data.ClearbanSticker.Stkid8
	stkpkgid8 = Data.ClearbanSticker.Stkpkgid8
	stkid9 = Data.CancelallSticker.Stkid9
	stkpkgid9 = Data.CancelallSticker.Stkpkgid9
	MsgBroadcast = Data.BroadcastBack
	if len(Data.TimeBanBack) != 0 {
		now := time.Now()
		for a := range Data.TimeBanBack {
			if InArray2(Squadlist, a) {
				tims := Data.TimeBanBack[a]
				if now.Sub(tims) < 24*time.Hour {
					self := GetKorban(a)
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
	if len(Data.CreatorBack) != 0 {
		for _, i := range Data.CreatorBack {
			UserBot.AddCreator(i)
		}
	}
	if len(Data.MakerBack) != 0 {
		for _, i := range Data.MakerBack {
			UserBot.AddMaker(i)
		}
	}
	if len(Data.SellerBack) != 0 {
		for _, i := range Data.SellerBack {
			UserBot.AddSeller(i)
		}
	}
	if len(Data.BuyerBack) != 0 {
		for _, i := range Data.BuyerBack {
			UserBot.AddBuyer(i)
		}
	}
	if len(Data.OwnerBack) != 0 {
		for _, i := range Data.OwnerBack {
			UserBot.AddOwner(i)
		}
	}
	if len(Data.MasterBack) != 0 {
		for _, i := range Data.MasterBack {
			UserBot.AddMaster(i)
		}
	}
	if len(Data.AdminBack) != 0 {
		for _, i := range Data.AdminBack {
			UserBot.AddAdmin(i)
		}
	}
	if len(Data.BotBack) != 0 {
		for _, i := range Data.BotBack {
			UserBot.AddBot(i)
		}
	}
	if len(Data.ProkickBack) != 0 {
		for _, to := range Data.ProkickBack {
			Room := linetcr.GetRoom(to)
			Room.ProKick = true
		}
	}
	if len(Data.ProCancelBack) != 0 {
		for _, to := range Data.ProCancelBack {
			Room := linetcr.GetRoom(to)
			Room.ProCancel = true
		}
	}
	if len(Data.ProInviteBack) != 0 {
		for _, to := range Data.ProInviteBack {
			Room := linetcr.GetRoom(to)
			Room.ProInvite = true
		}
	}
	if len(Data.ProQrBack) != 0 {
		for _, to := range Data.ProQrBack {
			Room := linetcr.GetRoom(to)
			Room.ProQr = true
		}
	}
	if len(Data.ProNoteBack) != 0 {
		for _, to := range Data.ProNoteBack {
			Room := linetcr.GetRoom(to)
			Room.ProNote = true
		}
	}
	if len(Data.ProNameBack) != 0 {
		for _, to := range Data.ProNameBack {
			Room := linetcr.GetRoom(to)
			Room.ProName = true
		}
	}
	if len(Data.ProAlbumBack) != 0 {
		for _, to := range Data.ProAlbumBack {
			Room := linetcr.GetRoom(to)
			Room.ProAlbum = true
		}
	}
	if len(Data.ProPictureBack) != 0 {
		for _, to := range Data.ProPictureBack {
			Room := linetcr.GetRoom(to)
			Room.ProPicture = true
		}
	}
	if len(Data.ProjoinBack) != 0 {
		for _, to := range Data.ProjoinBack {
			Room := linetcr.GetRoom(to)
			Room.ProJoin = true
		}
	}
	if len(Data.AnnunceBack) != 0 {
		for _, to := range Data.AnnunceBack {
			Room := linetcr.GetRoom(to)
			Room.Announce = true
		}
	}
	if len(Data.ProLinkBack) != 0 {
		for _, to := range Data.ProLinkBack {
			Room := linetcr.GetRoom(to)
			Room.ProLink = true
		}
	}
	if len(Data.ProFlexBack) != 0 {
		for _, to := range Data.ProFlexBack {
			Room := linetcr.GetRoom(to)
			Room.ProFlex = true
		}
	}
	if len(Data.ProImageBack) != 0 {
		for _, to := range Data.ProImageBack {
			Room := linetcr.GetRoom(to)
			Room.ProImage = true
		}
	}
	if len(Data.ProVideoBack) != 0 {
		for _, to := range Data.ProVideoBack {
			Room := linetcr.GetRoom(to)
			Room.ProVideo = true
		}
	}
	if len(Data.ProCallBack) != 0 {
		for _, to := range Data.ProCallBack {
			Room := linetcr.GetRoom(to)
			Room.ProCall = true
		}
	}
	if len(Data.ProSpamBack) != 0 {
		for _, to := range Data.ProSpamBack {
			Room := linetcr.GetRoom(to)
			Room.ProSpam = true
		}
	}
	if len(Data.ProStickerBack) != 0 {
		for _, to := range Data.ProStickerBack {
			Room := linetcr.GetRoom(to)
			Room.ProSticker = true
		}
	}
	if len(Data.ProContactBack) != 0 {
		for _, to := range Data.ProContactBack {
			Room := linetcr.GetRoom(to)
			Room.ProContact = true
		}
	}
	if len(Data.ProPostBack) != 0 {
		for _, to := range Data.ProPostBack {
			Room := linetcr.GetRoom(to)
			Room.ProPost = true
		}
	}
	if len(Data.ProFileBack) != 0 {
		for _, to := range Data.ProFileBack {
			Room := linetcr.GetRoom(to)
			Room.ProFile = true
		}
	}
	if len(Data.GadminBack) != 0 {
		for to := range Data.GadminBack {
			Room := linetcr.GetRoom(to)
			if len(Data.GadminBack[to]) != 0 {
				for _, user := range Data.GadminBack[to] {
					if !InArray2(Room.Gadmin, user) {
						Room.Gadmin = append(Room.Gadmin, user)
					}
				}
			}
		}
	}
	if len(Data.GownerBack) != 0 {
		for to := range Data.GownerBack {
			Room := linetcr.GetRoom(to)
			if len(Data.GownerBack[to]) != 0 {
				for _, user := range Data.GownerBack[to] {
					if !InArray2(Room.Gowner, user) {
						Room.Gowner = append(Room.Gowner, user)
					}
				}
			}
		}
	}
	if len(Data.GbanBack) != 0 {
		for to := range Data.GbanBack {
			Room := linetcr.GetRoom(to)
			if len(Data.GbanBack[to]) != 0 {
				for _, user := range Data.GbanBack[to] {
					if MemUser(to, user) {
						if !InArray2(Room.Gban, user) {
							Room.Gban = append(Room.Gban, user)
						}
					}
				}
			}
		}
	}
	if len(Data.BanBack) != 0 {
		for _, user := range Data.BanBack {
			Banned.AddBan(user)
		}
	}
	if len(Data.FuckBack) != 0 {
		for _, user := range Data.FuckBack {
			Banned.AddFuck(user)
		}
	}
	if len(Data.LockBack) != 0 {
		for _, user := range Data.LockBack {
			Banned.AddBan2(user)
		}
	}
	if len(Data.MuteBack) != 0 {
		for _, user := range Data.MuteBack {
			Banned.AddMute(user)
		}
	}
	if len(Data.WordbanBack) != 0 {
		for _, msg := range Data.WordbanBack {
			if !InArray2(Data.WordbanBack, msg) {
				Data.WordbanBack = append(Data.WordbanBack, msg)
			}
		}
	}
	TimeSave = time.Now()
	fmt.Println("\nSUCCESS RUN ALL BOTS")
}
func SaveProHistory() {
	AllowDoOnce = 0
	for i := range ClientBot {
		Data.Kikhistory = Data.Kikhistory + ClientBot[i].Ckick
		Data.Invhistory = Data.Invhistory + ClientBot[i].Cinvite
		Data.Canclhistory = Data.Canclhistory + ClientBot[i].Ccancel
		ClientBot[i].Ckick = 0
		ClientBot[i].Cinvite = 0
		ClientBot[i].Ccancel = 0
	}
}

func CheckChatBan() {
	defer panicHandle("CheckChatBan")
	if AllowDoOnce == 0 {
		for _, cl := range ClientBot {
			if !linetcr.InArrayCl(linetcr.KickBanChat, cl) && !cl.Frez {
				r, _ := cl.GetHomeProfile(cl.MID)
				if linetcr.GetBannedChat(r) == 1 {
					linetcr.BanChatAdd(cl)
				}
			}
		}
		AllowDoOnce++
	}
}


//DONE
func autoset() {
	defer panicHandle("autoset")
	now := time.Now()
	for _, room := range linetcr.SquadRoom {
		if !room.Fight.IsZero() {
			if now.Sub(room.Fight) >= 3*time.Second {
				if AutoPro {
					room.AutoBro()
				}
				room.Fight = time.Time{}
				var cll *linetcr.Account
				if len(room.Client) != 0 {
					cll = room.Client[0]
					name, mem, pending := cll.GetChatList(room.Id)
					room.Name = name
					room.Reset()
					sort.Slice(room.Ava, func(i, j int) bool {
						return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
					})
					sort.Slice(room.HaveClient, func(i, j int) bool {
						return room.HaveClient[i].KickPoint < room.HaveClient[j].KickPoint
					})
					exe := []*linetcr.Account{}
					for _, cls := range room.Client {
						if InArray2(mem, cls.MID) && !cls.Frez && !cls.Limited {
							exe = append(exe, cls)
						}
					}
					room.HaveClient = exe
					if len(exe) != 0 {
						Backup := []string{}
						li, ok := backlist.Get(room.Id)
						if ok {
							mems := li.([]string)
							for _, l := range mems {
								if !InArray2(mem, l) && !InArray2(Backup, l) && !InArray2(pending, l) {
									Backup = append(Backup, l)
								}
							}
						}
						backlist.Set(room.Id, []string{})
						if len(Backup) != 0 {
							celek := len(Backup)
							no := 0
							bat := 5
							ClAct := len(exe)
							if ClAct != 0 {
								if celek < bat {
									for _, cl := range exe {
										cl.GetRecommendationIds()
										for _, mid := range Backup {
											linetcr.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range Backup {
											if InArray2(fl, mid) {
												bb = append(bb, mid)
												Backup = Remove(Backup, mid)
											}
										}
										if len(bb) != 0 {
											cl.InviteIntoGroupNormal(room.Id, bb)
										}
										if len(Backup) == 0 {
											break
										}
									}
								} else {
									hajar := []string{}
									z := celek / bat
									y := z + 1
									for i := 0; i < y; i++ {
										if no >= ClAct {
											no = 0
										}
										client := exe[no]
										if i == z {
											hajar = Backup[i*bat:]
										} else {
											hajar = Backup[i*bat : (i+1)*bat]
										}
										if len(hajar) != 0 {
											client.GetRecommendationIds()
											for _, mid := range hajar {
												linetcr.AddContact3(client, mid)
											}
											fl, _ := client.GetAllContactIds()
											bb := []string{}
											for _, mid := range hajar {
												if InArray2(fl, mid) {
													bb = append(bb, mid)
												}
											}
											if len(bb) != 0 {
												client.InviteIntoGroupNormal(room.Id, bb)
											}
										}
										no += 1
									}
								}
							}
						}
					}
				} else {
					linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				}
				filterWar.clear()
				Nkick = &hashmap.HashMap{}
				filterop = &hashmap.HashMap{}
				oplist = []int64{}
				Ceknuke = &hashmap.HashMap{}
				cekoptime = []int64{}
				PurgeOP = []int64{}
				filtermsg = &hashmap.HashMap{}
				opjoin = []string{}
				room.ListInvited = []string{}
				Cekpurge = []int64{}
				AutoproN = false
				cekGo = []int64{}
			}
		}
		if !room.Leave.IsZero() {
			if now.Sub(room.Leave) >= 3*time.Second {
				room.Leave = time.Time{}
				if len(room.LeaveBack) != 0 {
					var cll *linetcr.Account
					if len(room.Client) != 0 {
						cll = room.Client[0]
						botleave = &hashmap.HashMap{}
						name, mem, invs := cll.GetChatList(room.Id)
						room.Name = name
						exe := []*linetcr.Account{}
						for _, cls := range room.Client {
							if InArray2(mem, cls.MID) && !InArray2(room.GoMid, cls.MID) {
								exe = append(exe, cls)
							}
						}
						inv := []string{}
						asu := room.LeaveBack
						room.LeaveBack = []string{}
						if len(exe) != 0 {
							for _, l := range asu {
								if !MemBan(room.Id, l) && !InArray2(inv, l) && !InArray2(mem, l) && !InArray2(invs, l) {
									inv = append(inv, l)
								}
							}
							if len(inv) != 0 {
								cls := exe
								for _, cl := range cls {
									if !cl.Limited {
										cl.GetRecommendationIds()
										for _, mid := range inv {
											linetcr.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range inv {
											if InArray2(fl, mid) {
												bb = append(bb, mid)
											}
										}
										cl.InviteIntoGroupNormal(room.Id, bb)
										for _, mid := range bb {
											if MemUser(room.Id, mid) {
												cl.UnFriend(mid)
											}
										}
										break
									}
								}
							}
						}
					} else {
						linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
					}
				}
			}
		}
	}
	for _, cl := range linetcr.Waitadd {
		if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
			v, ok := linetcr.BlockAdd.Get(cl.MID)
			if !ok {
				if now.Sub(cl.TimeBan) >= 1*time.Hour  {
					cl.Limitadd = false
					cl.Add = 0
					cl.Lastadd = now
					linetcr.Waitadd = linetcr.RemoveCl(linetcr.Waitadd, cl)
					linetcr.BlockAdd.Del(cl.MID)
				}
			} else {
				if now.Sub(v.(time.Time)) >= 24*time.Hour {
					linetcr.BlockAdd.Del(cl.MID)
					cl.Limitadd = false
					cl.Add = 0
					cl.Lastadd = now
					linetcr.Waitadd = linetcr.RemoveCl(linetcr.Waitadd, cl)
					linetcr.BlockAdd.Del(cl.MID)
				}
			}
		}
	}
	for _, cl := range ClientBot {
		if now.Sub(cl.Lastadd) >= 1*time.Hour  {
			cl.Add = 0
			cl.Lastadd = now
		}
		if now.Sub(cl.Lastkick) >= 1*time.Hour  {
			cl.TempKick = 0
			cl.TempInv = 0
		}
	}
	for _, cl := range linetcr.KickBans {
		if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
			v, ok := linetcr.GetBlock.Get(cl.MID)
			if !ok {
				if now.Sub(cl.TimeBan) >= 1*time.Hour  {
					linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
					cl.Limited = false
					cl.TempKick = 0
					cl.TempInv = 0
					cl.Frez = false
					linetcr.GetBlock.Del(cl.MID)
				}
			} else {
				if now.Sub(v.(time.Time)) >= 24*time.Hour {
					linetcr.GetBlock.Del(cl.MID)
					linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
					cl.Limited = false
					cl.Frez = false
					cl.TempKick = 0
					cl.TempInv = 0
					cl.KickCount = 0
					cl.KickPoint = 0
					cl.InvCount = 0
					cl.CountDay = 0
				}
			}
		}
	}
	for m, v := range linetcr.HashToMap(linetcr.GetBlockAdd) {
		cl := GetKorban(m)
		if cl.Limited {
			if now.Sub(v.(time.Time)) >= 1*time.Hour  {
				cl.Limitadd = false
				linetcr.GetBlockAdd.Del(cl.MID)
			}
		}
	}
	if now.Sub(aclear) >= 30*time.Second {
		filterop = &hashmap.HashMap{}
		Nkick = &hashmap.HashMap{}
		filterWar.clear()
		oplist = []int64{}
		timeSend = []int64{}
		Ceknuke = &hashmap.HashMap{}
		cekoptime = []int64{}
		filtermsg = &hashmap.HashMap{}
		aclear = now
		PurgeOP = []int64{}
		Cekpurge = []int64{}
		opjoin = []string{}
		cekGo = []int64{}
		AutoproN = false
	}
	if now.Sub(TimeSave) >= 1*time.Hour {
		SaveBackup()
		TimeBackup = now
	}
	if !TimeBackup.IsZero() {
		BackSeave()
	}
	if now.Sub(TimeClear) >= 3*time.Hour {
		time.Sleep(3*time.Hour)
		exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
		exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
		fmt.Println("\n clear vps cache automatically every 6 hours")
	}
	if AutoBc {
		if now.Sub(TimeBc) >= time.Duration(TimeBroadcast)*time.Minute {
			exe := []*linetcr.Account{}
			for _, cls := range ClientBot {
				exe = append(exe, cls)
			}
			if len(exe) != 0 {
				cls := exe
				for _, cl := range cls {
					if cl.Limited == false {
						gr, _ := cl.GetGroupIdsJoined()
						time.Sleep(250 * time.Second)
						for _, gi := range gr {
							if Typebc == "msg" {
								cl.SendMessage(gi, MsgBroadcast)
							} else if Typebc == "image" {
								cl.SendFoto(gi, linetcr.Imagebc)
							}
						}
					}
				}
			}
		}
	}
}

//DONE
func back(to, u string) {
	li, ok := backlist.Get(to)
	if ok {
		list := li.([]string)
		if !InArray2(list, u) {
			list = append(list, u)
		}
		backlist.Set(to, list)
	} else {
		list := []string{u}
		backlist.Set(to, list)
	}
}

//DONE


func InArrayChat(arr []*talkservice.Chat, str *talkservice.Chat) bool {
	for _, tar := range arr {
		if tar.ChatMid == str.ChatMid {
			return true
		}
	}
	return false
}

//NEED FIX
func clone(p *linetcr.Account, pp string, vp string, co string, cv string, name string, status string) {
	if pp != "" && vp != "" {
		err := p.UpdateVideoProfile(vp)
		if err == nil {
			err := p.UpdatePictureProfile(pp, "v")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
		os.Remove(vp)
		os.Remove(pp)
	} else if pp != "" {
		err := p.UpdatePictureProfile(pp, "p")
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(pp)
	}
	if co != "" && cv == "" {
		err := p.UpdateCover(co)
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(co)
	} else if co != "" && cv != "" {
		p.UpdateCoverVideo(cv)
		err := p.UpdateCoverWithVideo(co)
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(cv)
		os.Remove(co)
	}
	p.UpdateProfileName(name)
	p.UpdateProfileBio(status)
	p.Namebot = name
}

//DONE
func Gtotal(client *linetcr.Account, gid string) string {
	list := ""
	GetSquad(client, gid)
	Room := linetcr.GetRoom(gid)
	_, mem, pending := client.GetChatList(gid)
	Glist := []string{}
	mGlist := []string{}
	for _, from := range mem {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(Glist, from) {
				Glist = append(Glist, from)
			}
		}
	}
	for _, from := range pending {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(mGlist, from) {
				mGlist = append(mGlist, from)
			}
		}
	}
	list += fmt.Sprintf("Gtotal: %s", Room.Name)
	pp := len(mGlist)
	list += "\n Member: 38"
	list += fmt.Sprintf("\n Pending: %v", pp)
	list += "\n Total: 38"
	list += "\n Freeinvite: 462"
	return list
}
func InfoCreator(client *linetcr.Account) string {
       list := ""
	if len(CREATOR) != 0 {
		cuh, _ := client.GetContacts(CREATOR)
		for _, prs := range cuh {
			name := prs.DisplayName
			list += fmt.Sprintf("\n ⚙️ Developer : %v", name)
		}
	}
	return list
}
func InfoGroup(client *linetcr.Account, gid string) string {
	list := ""
	GetSquad(client, gid)
	Room := linetcr.GetRoom(gid)
	_, mem, pending := client.GetChatList(gid)
	Developer := []string{}
	creator := []string{}
	buyer := []string{}
	owner := []string{}
	master := []string{}
	admin := []string{}
	gowner := []string{}
	gadmin := []string{}
	squad := []string{}
	bot := []string{}
	ban := []string{}
	fuck := []string{}
	mute := []string{}
	Gban := []string{}
	Glist := []string{}
	Maker := []string{}
	Seller := []string{}
	mGlist := []string{}
	for _, from := range mem {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(Glist, from) {
				Glist = append(Glist, from)
			}
		} else if UserBot.GetCreator(from) {
			creator = append(creator, from)
		} else if UserBot.GetMaker(from) {
			Maker = append(Maker, from)
		} else if UserBot.GetSeller(from) {
			Seller = append(Seller, from)
		} else if InArray2(DEVELOPER, from) {
			Developer = append(Developer, from)
		} else if UserBot.GetBuyer(from) {
			buyer = append(buyer, from)
		} else if UserBot.GetOwner(from) {
			owner = append(owner, from)
		} else if UserBot.GetMaster(from) {
			master = append(master, from)
		} else if UserBot.GetAdmin(from) {
			admin = append(admin, from)
		} else if InArray2(Room.Gowner, from) {
			gowner = append(gowner, from)
		} else if InArray2(Room.Gadmin, from) {
			gadmin = append(gadmin, from)
		} else if UserBot.GetBot(from) {
			bot = append(bot, from)
		} else if Banned.GetFuck(from) {
			fuck = append(fuck, from)
		} else if Banned.GetBan(from) {
			ban = append(ban, from)
		} else if Banned.GetMute(from) {
			mute = append(mute, from)
		} else if InArray2(Room.Gban, from) {
			Gban = append(Gban, from)
		} else if InArray2(Squadlist, from) {
			squad = append(squad, from)
		}
	}
	for _, from := range pending {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !InArray2(mGlist, from) {
				mGlist = append(mGlist, from)
			}
		} else if UserBot.GetCreator(from) {
			creator = append(creator, from)
		} else if UserBot.GetMaker(from) {
			Maker = append(Maker, from)
		} else if UserBot.GetSeller(from) {
			Seller = append(Seller, from)
		} else if InArray2(DEVELOPER, from) {
			Developer = append(Developer, from)
		} else if UserBot.GetBuyer(from) {
			buyer = append(buyer, from)
		} else if UserBot.GetOwner(from) {
			owner = append(owner, from)
		} else if UserBot.GetMaster(from) {
			master = append(master, from)
		} else if UserBot.GetAdmin(from) {
			admin = append(admin, from)
		} else if InArray2(Room.Gowner, from) {
			gowner = append(gowner, from)
		} else if InArray2(Room.Gadmin, from) {
			gadmin = append(gadmin, from)
		} else if UserBot.GetBot(from) {
			bot = append(bot, from)
		} else if Banned.GetFuck(from) {
			fuck = append(fuck, from)
		} else if Banned.GetBan(from) {
			ban = append(ban, from)
		} else if Banned.GetMute(from) {
			mute = append(mute, from)
		} else if InArray2(Room.Gban, from) {
			Gban = append(Gban, from)
		} else if InArray2(Squadlist, from) {
			squad = append(squad, from)
		}
	}
	list += fmt.Sprintf("Group Info: %s", Room.Name)
	if len(Glist) != 0 {
		list += "\n\nMember: \n"
		cuh, _ := client.GetContacts(Glist)
		for _, prs := range cuh {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}
	if len(mGlist) != 0 {
		chp, _ := client.GetContacts(mGlist)
		list += "\n\n Pending: \n"
		for _, prs := range chp {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}
	if len(Glist)+len(mGlist) != len(pending)+len(mem) {
		list += "\n\nUsers have access:\n"
		if len(Developer) != 0 {
			list += "\nExist in Developer:\n"
			for n, xx := range Developer {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(creator) != 0 {
			list += "\nExist in Creator:\n"
			for n, xx := range creator {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Maker) != 0 {
			list += "\nExist in Maker:\n"
			for n, xx := range Maker {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Seller) != 0 {
			list += "\nExist in Seller:\n"
			for n, xx := range Seller {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(buyer) != 0 {
			list += "\nExist in Buyers:\n"
			for n, xx := range buyer {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(owner) != 0 {
			list += "\nExist in Owners:\n"
			for n, xx := range owner {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(master) != 0 {
			list += "\nExist in Masters:\n"
			for n, xx := range master {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(admin) != 0 {
			list += "\nExist in Admins:\n"
			for n, xx := range admin {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(gowner) != 0 {
			list += "\nExist in Gowners:\n"
			for n, xx := range gowner {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(gadmin) != 0 {
			list += "\nExist in Gadmins:\n"
			for n, xx := range gadmin {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(bot) != 0 {
			list += "\nExist in Botlist\n"
			for n, xx := range bot {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(squad) != 0 {
			list += "\nExist in Squad:\n"
			for n, xx := range squad {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(ban) != 0 {
			list += "Exist in Banlist:\n"
			for n, xx := range ban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(fuck) != 0 {
			list += "\nExist in Fucklist:\n"
			for n, xx := range fuck {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Gban) != 0 {
			list += "\nExist in Gbanlist:\n\n"
			for n, xx := range Gban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(mute) != 0 {
			list += "\nExist in Mutelist:\n\n"
			for n, xx := range mute {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
	}
	return list
}

//DONE
func nukeAll(Client *linetcr.Account, Group string) {
	defer linetcr.PanicOnly()
	memlist := []string{}
	_, memlists := Client.GetGroupMember(Group)
	act := []*linetcr.Account{}
	for mid, _ := range memlists {
		if MemUser(Group, mid) {
			memlist = append(memlist, mid)
		} else if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				act = append(act, cl)
			}
		}
	}
	lact := len(act)
	if lact == 0 {
		return
	} else {
		sort.Slice(act, func(i, j int) bool {
			return act[i].KickPoint < act[j].KickPoint
		})
		celek := len(memlist)
		if celek < MaxKick || lact == 1 {
			cl := act[0]
			for _, mem := range memlist {
				go cl.DeleteOtherFromChat(Group, mem)
			}
		} else {
			hajar := []string{}
			z := celek / MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= lact {
					no = 0
				}
				go func(Group string, no int, i int, z int, memlist []string, act []*linetcr.Account) {
					Client = act[no]
					if i == z {
						hajar = memlist[i*MaxKick:]
					} else {
						hajar = memlist[i*MaxKick : (i+1)*MaxKick]
					}
					if len(hajar) != 0 {
						for _, target := range hajar {
							go Client.DeleteOtherFromChat(Group, target)
						}
					}
				}(Group, no, i, z, memlist, act)
				no += 1
			}
		}
		linetcr.GetRoom(Group).HaveClient = act
	}
}

//GechatLlinetcr
func KickCancelBan129(client *linetcr.Account, to string) {
	go runtime.Gosched()
	x,_ := client.NewGetChat(to)
	grup := x.Chats[0].Extra.GroupExtra.MemberMids
	g,_ := client.NewGetChat(to)
	for _, v := range client.Backup {if _, cek := grup[v]; cek {if v == client.MID {var Batas = 0;if g != nil {targets := g.Chats[0].Extra.GroupExtra.MemberMids;targets2 := g.Chats[0].Extra.GroupExtra.InviteeMids;listMid := []string{};listMid2 := []string{};for cok := range targets {listMid = append(listMid, cok)};for cok := range targets2 {listMid2 = append(listMid2, cok)};for v := range Banned.Banlist {go runtime.Gosched();if contains(listMid, Banned.Banlist[v]) {go runtime.Gosched();go func(v string) {client.NewkickGroup(to, v)}(Banned.Banlist[v]);Batas = Batas + 1;if int64(Batas) >= int64(MaxKick) {Batas = 0;break}} else if contains(listMid2, Banned.Banlist[v]) {go runtime.Gosched();go func(v string) {client.NewcancelGroup(to, v)}(Banned.Banlist[v]);Batas = Batas + 1;if int64(Batas) >= int64(MaxCancel) {Batas = 0;break}}}}};break} else {continue}
	}
}
func KickBan132V2(client *linetcr.Account, to string) {
	var Batas = 0
	chat, _ := client.NewGetChat(to)
	if chat != nil {
		memb := chat.Chats[0].Extra.GroupExtra.MemberMids
		for x := range client.Backup {if _, blog := memb[client.Backup[x]]; blog {if client.MID == client.Backup[x] {go func(){for mid := range memb {if IsBlacklist(client, mid) == true {go func(mid string) {go client.NewkickGroup(to, mid)}(mid);Batas = Batas + 1;if int64(Batas) >= int64(MaxKick) {Batas = 0;break}}}}()};break} else {continue}
		}
	}
	runtime.GOMAXPROCS(cpu)
}
func CancelBan125V2(client *linetcr.Account, to string) {
	var Batas = 0
	g, _ := client.NewGetChat(to)
	if g != nil {
		memb := g.Chats[0].Extra.GroupExtra.MemberMids
		memb1 := g.Chats[0].Extra.GroupExtra.InviteeMids
		for x := range client.Backup {if _, blog := memb[client.Backup[x]]; blog {if client.MID == client.Backup[x] {var wg sync.WaitGroup;wg.Add(len(memb1));go func(){for mid := range memb1 {if IsBlacklist(client, mid) == true {go func(mid string) {defer wg.Done();go client.NewcancelGroup(to, mid)}(mid);Batas = Batas + 1;if int64(Batas) >= int64(MaxCancel) {Batas = 0;break}}}}();wg.Wait()};break} else {continue}
		}
	}
	runtime.GOMAXPROCS(cpu)
}
func KickBan132(client *linetcr.Account, to string, user string) {
	if len(user) > 10 {go func() {FastKick(client, to)}()
	}
	runtime.GOMAXPROCS(cpu)
}
func CancelBan125(client *linetcr.Account, to string, korban []string) {
	if len(korban) > 10 {go func() {FastCancel(client, to)}()
	}
	runtime.GOMAXPROCS(cpu)
}
func FastKick(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChats([]string{to})
	mem := c.Chats[0].Extra.GroupExtra.MemberMids
	for k, _ := range mem {
		if IsBlacklist(client, k) == true {go func(k string) {go client.DeleteOtherFromChats(to, []string{k})
			}(k)
			Batas = Batas + 1
			if Batas >= MaxKick-1 {
				Batas = 0
				break
			}
		}
	}
}
func FastCancel(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChats([]string{to})
	pend := c.Chats[0].Extra.GroupExtra.InviteeMids
	for k, _ := range pend {
		if IsBlacklist(client, k) == true {go func(k string) {go client.CancelChatInvitations(to, []string{k})
			}(k)
			Batas = Batas + 1
			if Batas >= MaxCancel-1 {
				Batas = 0
				break
			}
		}
	}
}
func NodeBans(client *linetcr.Account, to string, korban []string) {
	if len(korban) > 3 {go func() {FastKick(client, to);FastCancel(client, to)}()
	}
	runtime.GOMAXPROCS(cpu)
}
func PurgeFaster(client *linetcr.Account, to string) {
	for x := range Squadlist {if linetcr.IsMembers(client, to, Squadlist[x]) == true {if client.MID == Squadlist[x] {go func() {FastKick(client, to);FastCancel(client, to)}()};break} else {continue}
	}
	runtime.GOMAXPROCS(cpu)
}

func KickMemBan(client *linetcr.Account, to string) {
	var Batas = 0;c, _ := client.GetChatsV2(to);mem := c.Extra.GroupExtra.MemberMids;for k, _ := range mem {if MemBan(to, k) == true {go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k);Batas = Batas + 1;if Batas >= MaxKick-1 {Batas = 0;break}}
	}
}
func CanMemBan(client *linetcr.Account, to string) {
	var Batas = 0;c, _ := client.GetChatsV2(to);pend := c.Extra.GroupExtra.InviteeMids;for k, _ := range pend {if MemBan(to, k) == true {go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k);Batas = Batas + 1;if Batas >= MaxKick-1 {Batas = 0;break}}
	}
}
func KickCancel(client *linetcr.Account, to string) {
	for _, v := range Banned.Banlist {
		if Banned.GetBan(v) {
			if linetcr.IsMembers(client, to, v) == true {go func(v string) {client.DeleteOtherFromChats(to, []string{v})}(v)
			}
			if linetcr.IsPending(client, to, v) == true {if Banned.GetBan(v) {go func(v string) {client.CancelChatInvitations(to, []string{v})}(v)}
			}
		}
	}
}

func QrKick(client *linetcr.Account, to string) {
	if len(Banned.Banlist) != 0 {
		for _, v := range Banned.Banlist {
			if linetcr.IsMembers(client, to, v) == true {go func() {client.DeleteOtherFromChats(to, []string{v})}()}
			if linetcr.IsPending(client, to, v) == true {go func() {client.CancelChatInvitations(to, []string{v})}()}
		}
	}
}

func AcceptJoin(client *linetcr.Account, Group string) {
	defer panicHandle("AcceptJoin")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(Group)
	if AutoPro {
		Room.AutoBro()
	}
	if canceljoin {
		CanceljoinBot(client, Group)
	} else if NukeJoin {
		nukeAll(client, Group)
	}
	if Autojoin == "qr" {
		AutojoinQr(client, Group)
	} else {
		if Autojoin == "inv" {
			Setinviteto(client, Group, client.Squads)
		}
	}
}
func AcceptJoinV2(client *linetcr.Account, Group string) {
	defer panicHandle("AcceptJoinV2")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(Group)
	if AutoPro {
		Room.AutoBro()
	}
	_, memlist := client.GetGroupMember(Group)
	oke := []string{}
	ban := []string{}
	exe := []*linetcr.Account{}
	Botss := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			oke = append(oke, mid)
			cl := GetKorban(mid)
			Botss = append(Botss, cl)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemBan(Group, mid) {
			ban = append(ban, mid)
		}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		Room.HaveClient = exe
		Room.Client = Botss
		Room.Bot = oke
		linetcr.SetAva(Group, oke)
		if canceljoin {
			CanceljoinBot(client, Group)
		} else if NukeJoin {
			nukeAll(client, Group)
		}
		if AutoPurge {
			if len(ban) != 0 {
				no := 0
				ah := 0
				for _, target := range ban {
					go func(target string, no int) {exe[no].DeleteOtherFromChats(Group, []string{target})
					}(target, no)
					if ah >= MaxKick {
						no++
						if no >= len(exe) {
							no = 0
						}
						ah = 0
					}
					ah++
				}
				for _, enemy := range ban {
					go func(enemy string, no int) {exe[no].CancelChatInvitations(Group, []string{enemy})
					}(enemy, no)
					if ah >= MaxCancel {
						no++
						if no >= len(exe) {
							no = 0
						}
						ah = 0
					}
					ah++
				}
			}
		}
		if Autojoin == "qr" {
			AutojoinQr(exe[0], Group)
		} else {
			if Autojoin == "inv" {
				Setinviteto(exe[0], Group, exe[0].Squads)
			}
		}
	}
}
func getArg() string {
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

func abort() {
	remotegrupidto = ""
	StartChangeImg = false
	StartChangevImg = false
	StartChangevImg2 = false
	Sinderremote = []string{}
	remotegrupid = ""
	RemoteOwner = false
	RemoteMaster = false
	RemoteAdmin = false
	RemoteContact = false
	RemoteBan = false
	MidRemote = []string{}
	changepic = []*linetcr.Account{}
	ChangName = false
	ChangCover = false
	ChangPict = false
	ChangeBio = false
	ChangVpict = false
	ChangVcover = false
	AllCheng = false
	MsgBio = ""
	MsgName = ""
	BcImage        = false
	StartBc        = false
	BcVideo        = false
	StartBcV        = false
	GBcImage        = false
	GStartBc        = false
	GBcVideo        = false
	GStartBcV        = false
	FBcImage        = false
	FStartBc        = false
	FBcVideo        = false
	FStartBcV        = false
	SAVEBcImage        = false
	startSaveBc        = false
	FixedToken = false
	From_Token = ""
	Group_Token = ""
	timeabort = time.Now()
}

func AllBanList(self *linetcr.Account) string {
	listadm := "𝗔𝗹𝗹 𝗯𝗮𝗻𝗹𝗶𝘀𝘁𝘀:\n"
	if len(Banned.Banlist) != 0 {
		//listadm += "\n\n ☠️ 𝗕𝗮𝗻𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Banlist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(Banned.Fucklist) != 0 {
		listadm += "\n\n ☠️ 𝗙𝘂𝗰𝗸𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Fucklist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(Banned.Mutelist) != 0 {
		listadm += "\n\n ☠️ 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range Banned.Mutelist {
			num++
			rengs := strconv.Itoa(num)
			new := self.Getcontactuser(xd)
			if new != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				x, _ := self.GetContact(xd)
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	return listadm
}

func Checkserver(ip string) bool {
	if InArray2(ListIp, ip) {
		return true
	}
	return false
}

func MemBan(to, user string) bool {
	defer linetcr.PanicOnly()
	if Banned.GetBan(user) {
		return true
	} else if Banned.GetFuck(user) {
		return true
	} else {
		Room := linetcr.GetRoom(to)
		if InArray2(Room.Gban, user) {
			return true
		}
		return false
	}
	//return false
}

func MemBan2(to, user string) bool {
	defer linetcr.PanicOnly()
	if Banned.GetBan(user) {
		return true
	} else if Banned.GetFuck(user) {
		return true
	} else if Banned.GetMute(user) {
		return true
	} else {
		Room := linetcr.GetRoom(to)
		if InArray2(Room.Gban, user) {
			return true
		}
		return false
	}
	//return false
}

func Cekbanwhois(client *linetcr.Account, to string, targets []string) {
	room := linetcr.GetRoom(to)
	list := ""
	if len(targets) > 1 {
		ban := []string{}
		fuck := []string{}
		mute := []string{}
		Gban := []string{}
		for _, from := range targets {
			if Banned.GetFuck(from) {
				fuck = append(fuck, from)
			} else if Banned.GetBan(from) {
				ban = append(ban, from)
			} else if Banned.GetMute(from) {
				mute = append(mute, from)
			} else if InArray2(room.Gban, from) {
				Gban = append(Gban, from)
			}
		}
		if len(ban) != 0 {
			list += "𝗘𝘅𝗶𝘀𝘁.𝗶𝗻 𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range ban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(fuck) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗳𝘂𝗰𝗸𝗹𝗶𝘀𝘁:\n"
			for n, xx := range fuck {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Gban) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range Gban {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(mute) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁:\n"
			for n, xx := range mute {
				rengs := strconv.Itoa(n + 1)
				new := client.Getcontactuser(xx)
				if new != nil {
					list += rengs + ". Closed Account \n"
				} else {
					x, _ := client.GetContact(xx)
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
	} else {
		for _, from := range targets {
			if Banned.GetFuck(from) {
				list += "User have access exist in fuck list."
			} else if Banned.GetBan(from) {
				list += "User have access exist in ban list."
			} else if InArray2(room.Gban, from) {
				list += "User have access exist in gban list."
			} else if Banned.GetMute(from) {
				list += "User have access exist in mute list."
			}

		}
	}
	if list != "" {
		client.SendMessage(to, fancy(list))
	}
}

func autokickban(client *linetcr.Account, to string, target string) {
	if AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, []string{target})
			go client.CancelChatInvitations(aa, []string{target})
		}
	}
}

func Purgemode(Client *linetcr.Account, Group string) {
	defer panicHandle("Purgemode")
	_, memlists := Client.GetGroupMember(Group)
	for target, _ := range memlists {
		if MemBan(Group, target) {
			go Client.DeleteOtherFromChat(Group, target)
		}
	}
}

func KIckbansPurges(client *linetcr.Account, group string) {
	defer linetcr.PanicOnly()
	gr, _ := client.GetGroupIdsJoined()
	for _, aa := range gr {
		c, _ := client.GetChats([]string{aa})
		zxc := c.Chats[0].Extra.GroupExtra.MemberMids
		for k, _ := range zxc {
			if IsBlacklist(client, k) == true {
				go func(k string) {
					go client.DeleteOtherFromChats(aa, []string{k})
				}(k)
			}
		}
		_, _, pind := client.GetChatList(aa)
		for _, i := range pind {
			if IsBlacklist(client, i) == true {
				go func(i string) {
					go client.CancelChatInvitations(aa, []string{i})
				}(i)
			}
		}
	}
	client.SendMessage(group, fancy("Success nukebanlist"))
}
//Llinetcr_Allbanlist
func KickAllBan(client *linetcr.Account, to string) {
	_, mem, pend := client.GetChatList(to)
	for _, k := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k)
		}
	}
	for _, i := range pend {
		if IsBlacklist(client, i) == true {
			go func(i string) {go client.CancelChatInvitations(to, []string{i})}(i)
		}
	}
	runtime.GOMAXPROCS(cpu)
}

func KIckbansPurges1(client *linetcr.Account, group string) {
	defer linetcr.PanicOnly()
	gr, _ := client.GetGroupIdsJoined()
	nus := []string{}
	list := ""
	list += fmt.Sprintf("Purged %v groups: \n", len(gr))
	for num, aa := range gr {
		num++
		//list += fmt.Sprintf("%v- %v :", num, name)
		for _, v := range Banned.Banlist {
			if linetcr.IsMembers(client, aa, v) == true {
				if Banned.GetBan(v) {
					go func(v string) { client.DeleteOtherFromChats(aa, []string{v}) }(v)
					if linetcr.IsPending(client, aa, v) == true {
						client.CancelChatInvitations(aa, []string{v})
					}
					new := client.Getcontactuser(v)
					rengs := strconv.Itoa(num)
					if new != nil {
						list += "\n " + rengs + ". Closed Account"
					} else {
						x, _ := client.GetContact(v)
						nus = append(nus, v)
						list += "\n " + rengs + ". " + x.DisplayName
					}
				}
			}
		}
	}
	list += fmt.Sprintf("\n\nTotal kicks: %v.", len(nus))
	client.SendMessage(group, fancy(list))
}

func Ungban(group string, asu string) {
	room := linetcr.GetRoom(group)
	if InArray2(room.Gban, asu) {
		room.Gban = Remove(room.Gban, asu)
	}
}

func Addgban(asu string, group string) {
	room := linetcr.GetRoom(group)
	if !InArray2(room.Gban, asu) && asu != "" {
		room.Gban = append(room.Gban, asu)
	}
}

func IndexOf(data []string, element string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func Joinsave(Pelaku string, Optime int64) {
	defer linetcr.PanicOnly()
	ix := IndexOf(Detectjoin.User, Pelaku)
	if ix == -1 {
		Detectjoin.User = append(Detectjoin.User, Pelaku)
		Detectjoin.Time = append(Detectjoin.Time, Optime)
	} else {
		Detectjoin.Time[ix] = Optime
	}
}

func Checkmulti(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if InArray2(list2, v) {
			return true
		}
	}
	return false
}

func AddbanOp3(mid []string) {
	for _, m := range mid {
		Banned.AddBan(m)
	}
}

func SaveData() {
	defer linetcr.PanicOnly()
	file, _ := json.MarshalIndent(Data, "", "  ")
	_ = ioutil.WriteFile(DATABASE, file, 0644)
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Sutting down application.")
		os.Exit(0)
	}()
}
var letters = []rune("123456789")
func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomToString(count int) string {
	numb := make([]rune, count)
	for i := range numb {
		numb[i] = stringToInt[rand.Intn(len(stringToInt))]
	}
	return string(numb)
}

func panicHandle(s string) {
	if r := recover(); r != nil {
		Ides := fmt.Sprintf("\nEror: %s \nFunc: %v", r, s)
		println(Ides)
	}
}

func GetMentionData(data string) []string {
	var midmen []string
	var midbefore []string
	res := mentions{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArray2(midbefore, v.Mid) == false {
			midbefore = append(midbefore, v.Mid)
			midmen = append(midmen, v.Mid)
		}
	}

	return midmen
}

func MentionList(op *SyncService.Operation) []string {
	msg := op.Message
	str := fmt.Sprintf("%v", msg.ContentMetadata["MENTION"])
	taglist := GetMentionData(str)

	return taglist
}

func GetIP() net.IP {
	conn, err := net.Dial("udp", "0.0.0.0:80")
	//conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func SendMycreator(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	}
	return false
}
func SendMymaker(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	}
	return false
}
func SendMyseller(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	}
	return false
}
func SendMybuyer(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	}
	return false
}

func SendMyowner(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	}
	return false
}
func SendMymaster(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	}
	return false
}
func SendMyadmin(from string) bool {
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	}
	return false
}

func SendMygowner(group string, from string) bool {
	Room := linetcr.GetRoom(group)
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	} else if InArray2(Room.Gowner, from) == true {
		return true
	}
	return false
}
func SendMygadmin(group string, from string) bool {
	Room := linetcr.GetRoom(group)
	if InArray2(DEVELOPER, from) {
		return true
	} else if UserBot.GetCreator(from) {
		return true
	} else if UserBot.GetMaker(from) {
		return true
	} else if UserBot.GetSeller(from) {
		return true
	} else if UserBot.GetBuyer(from) {
		return true
	} else if UserBot.GetOwner(from) {
		return true
	} else if UserBot.GetMaster(from) {
		return true
	} else if UserBot.GetAdmin(from) {
		return true
	} else if InArray2(Room.Gowner, from) {
		return true
	} else if InArray2(Room.Gadmin, from) {
		return true
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
func GetCodeprem(num int, sinder string, group string) bool {
	if num == 0 {
		if InArray2(DEVELOPER, sinder) {
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

func PerCheckList() string {
	list := ""
	var test1 string
	if SetHelper.Rngcmd != nil {
		list += "✠ 𝗟𝗶𝘀𝘁 𝗽𝗲𝗿𝗺 :\n\n"
		for i := range SetHelper.Rngcmd {
			if SetHelper.Rngcmd[i] == 0 {
				test1 = "Dev"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 1 {
				test1 = "Creator"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 2 {
				test1 = "Maker"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 3 {
				test1 = "Seller"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 4 {
				test1 = "Buyer"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 5 {
				test1 = "Owner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 6 {
				test1 = "Master"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 7 {
				test1 = "Admin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 8 {
				test1 = "Gowner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if SetHelper.Rngcmd[i] == 9 {
				test1 = "Gadmin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			}
		}
	}
	return list
}
func MaxRevision(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
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
		_, value := SetHelper.Rngcmd[cmd2]
		if value == true {
			if SetHelper.Rngcmd[cmd2] != x {
				SetHelper.Rngcmd[cmd2] = x
				numr = 5
			}
		} else {
			list += "Command not found.\nUse ths Command First."
		}
	}
	if list != "Rank not found.\nUse .perm <rank> <command>\nAvailable ranks: \nbuyer/owner/master/admin/gowner/gadmin." {
		if list != "Command not found.\nUse ths Command First." {
			if numr != 5 {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("%v is already a %v command.\n", cmd2, cmd1)
			} else {
				cmd1 := test1
				cmd2 := test2
				list += fmt.Sprintf("Changed permission to %v for: %v \n", cmd1, cmd2)
			}
			client.SendMessage(to, fancy(list))
		} else {
			client.SendMessage(to, fancy(list+"\n"))
		}
	} else {
		client.SendMessage(to, fancy(list+"\n"))
	}

}
func GetComs(gr int, data string) int {
	defer linetcr.PanicOnly()
	_, value := SetHelper.Rngcmd[data]
	if value == false {
		SetHelper.Rngcmd[data] = gr
	}
	xx := SetHelper.Rngcmd[data]
	return xx
}
func CheckAccount(user string) *linetcr.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}
func Checkuser(client *linetcr.Account, group string) ([]*linetcr.Account, []string) {
	list := []string{}
	err, _, memlist := client.GetGroupMembers(group)
	if err != nil {
		return nil, list
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := CheckAccount(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemUser(group, mid) {
			list = append(list, mid)
		}
	}
	if len(exe) != 0 && len(list) != 0 {
		return exe, list
	}
	return nil, list
}
func CheckBan(client *linetcr.Account, group string) []string {
	list := []string{}
	err, _, memlist := client.GetGroupMembers(group)
	if err != nil {
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
func GETgrade(num int) string {
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
func CanceljoinBot(Client *linetcr.Account, Group string) {
	defer panicHandle("CanceljoinBot")
	_, _, pind := Client.GetChatList(Group)
	for _, i := range pind {
		if linetcr.IsPending(Client, Group, i) == true {
			Client.CancelChatInvitations(Group, []string{i})
		}
	}
}
func Nukjoin(Client *linetcr.Account, Optime int64, Group string) {
	defer panicHandle("Nukejoin")
	_, ok := Ceknuke.Get(Optime)
	if !ok {
		Ceknuke.Set(Optime, 1)
	} else {
		return
	}
	exe, list := Checkuser(Client, Group)
	if exe != nil {
		no := 0
		i := 0
		lm := len(list)
		acts := []*linetcr.Account{}
		var cl *linetcr.Account
		for ; i < lm; i++ {
			if no >= len(exe) {
				no = 0
			}
			acts = append(acts, exe[no])
			no += 1
		}
		for n, target := range list {
			go func(n int, target string) {
				cl = acts[n]
				cl.DeleteOtherFromChats(Group, []string{target})
			}(n, target)
		}
		_, _, pind := Client.GetChatList(Group)
		for _, p := range pind {
			if MemUser(Group, p) {
				if linetcr.IsPending(Client, Group, p) == true {
					Client.CancelChatInvitations(Group, []string{p})
				}
			}
		}
	}
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
func AutopurgeEnd(client *linetcr.Account, Group string, mem []string) {
	defer panicHandle("AutopurgeEnd")
	for _, target := range mem {
		client.DeleteOtherFromChats(Group, []string{target})
	}
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
func CheckMessage(waktu int64, typ int8) bool {
	if typ == 1 {
		for _, wkt := range timeSend {
			if wkt == waktu {
				return false
				break
			}
		}
		timeSend = append(timeSend, waktu)
		return true
	}
	return false
}
func Remove(s []string, r string) []string {
	new := make([]string, len(s))
	copy(new, s)
	for i, v := range new {
		if v == r {
			return append(new[:i], new[i+1:]...)
		}
	}
	return s
}
func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	x := d
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	if x < 60*time.Second {
		return fmt.Sprintf("%v", x)
	} else if x < 3600*time.Second {
		return fmt.Sprintf("%02dMinutes's", m)
	} else if x < 86400*time.Second {
		return fmt.Sprintf("%02dHour's %02dMinute's", h%24, m)
	} else {
		return fmt.Sprintf("%02dDay's %02dHour's %02dMinute's", h/24, h%24, m)
	}
}

func CancelEnd(client *linetcr.Account, Group string, mem []string) {
	defer panicHandle("CancelEnd")
	for _, target := range mem {
		client.CancelChatInvitations(Group, []string{target})
	}
}

func Setpurgealln(client *linetcr.Account, to string, invits []string) {
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		} else if linetcr.IsPending(client, to, cc) == true {
			client.CancelChatInvitations(to, []string{cc})
		}
	}

}
func SelectBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], err
	}
	return nil, err
}
func CheckBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], err
	}
	return nil, err
}
func GetKorban(user string) *linetcr.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}
func squadMention(mlist []string) (m *linetcr.Account, b bool) {
	for _, l := range mlist {
		if InArray2(Squadlist, l) {
			cl := GetKorban(l)
			return cl, true
		}
	}
	return nil, false
}
func StripOut(kata string) string {
	kata = strings.TrimSpace(kata)
	return kata
}

func Setinvitetomsg(client *linetcr.Account, to string, invits []string) []string {
	defer panicHandle("Setinvitetomsg")
	bans := []string{}
	news := []string{}
	room := linetcr.GetRoom(to)
	exe := room.HaveClient
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == false && linetcr.IsPending(client, to, cc) == false {
			if !MemBan(to, cc) {
				if linetcr.IsFriends(client, cc) == false {
					client.FindAndAddContactsByMidV2(cc)
					time.Sleep(250 * time.Millisecond)
				}
				news = append(news, cc)
			} else {
				bans = append(bans, cc)

			}
		}
	}
	if len(news) != 0 && len(exe) != 0 {
		celek := len(news)
		no := 0
		bat := 5
		ClAct := len(exe)
		if ClAct != 0 {
			if celek < bat {
				for _, cl := range exe {
					cl.GetRecommendationIds()
					for _, mid := range news {
						linetcr.AddContact3(cl, mid)
					}
					fl, _ := cl.GetAllContactIds()
					bb := []string{}
					for _, mid := range news {
						if InArray2(fl, mid) {
							bb = append(bb, mid)
							news = Remove(news, mid)
						}
					}
					if len(bb) != 0 {
						cl.InviteIntoGroupNormal(to, bb)
					}
					if len(news) == 0 {
						break
					}
				}
			} else {
				hajar := []string{}
				z := celek / bat
				y := z + 1
				for i := 0; i < y; i++ {
					if no >= ClAct {
						no = 0
					}
					client := exe[no]
					if i == z {
						hajar = news[i*bat:]
					} else {
						hajar = news[i*bat : (i+1)*bat]
					}
					if len(hajar) != 0 {
						client.GetRecommendationIds()
						for _, mid := range hajar {
							linetcr.AddContact3(client, mid)
						}
						fl, _ := client.GetAllContactIds()
						bb := []string{}
						for _, mid := range hajar {
							if InArray2(fl, mid) {
								bb = append(bb, mid)
							}
						}
						if len(bb) != 0 {
							client.InviteIntoGroupNormal(to, bb)
						}
					}
					no += 1
				}
			}
		}
	}
	return bans
}
func gettxt(from string, client *linetcr.Account, pesan string, rname string, sname string, Mid string, MentionMsg []string, group string) string {
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
		used = rname
	}
	if strings.HasPrefix(pesan, rname) {
		txt = strings.Replace(pesan, rname, "", 1)
		used = rname
	} else if strings.HasPrefix(pesan, sname) {
		txt = strings.Replace(pesan, sname, "", 1)
		used = sname
	}
	txt = StripOut(txt)
	return txt
}
func LogLast(op *SyncService.Operation, midds string) {
	defer linetcr.PanicOnly()
	if op.Type == 26 {
		if op.Message.ContentType == 18 {
			return
		}
	}
	LastActive.Set(midds, op)
}
func LogOp(op *SyncService.Operation, client *linetcr.Account) {
	defer linetcr.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	if tipe == 124 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 133 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 130 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 122 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 55 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 128 {
		if InArray2(Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 26 {
		msg := op.Message
		if InArray2(Squadlist, msg.From_) {
			return
		}
		LogLast(op, msg.From_)
	}
}
func Setkickto(client *linetcr.Account, to string, invits []string) {
	defer panicHandle("Setkickto")
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		}
	}

}
func addCon(cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !InArray2(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {cl.FindAndAddContactsByMidV2(con);time.Sleep(3 * time.Second)}
			}
		}
		n += 1
	}
}

func addConSq(to string, cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !InArray2(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {time.Sleep(5 * time.Second);cl.FindAndAddContactsByMidV2(con);time.Sleep(5 * time.Second)}
			}
			cl.SendMessage(to, fancy("Added friends"))
		}
		n += 1
	}
}


func addConSqV2(cons []string) {
	for _, cl := range ClientBot {
		for _, con := range cons {
			if linetcr.IsFriends(cl, con) == false && con != cl.MID {
				time.Sleep(5 * time.Second)
				cl.FindAndAddContactsByMidV5(con)
				time.Sleep(250 * time.Millisecond)
			}
		}
	}
}

func addConSingle(cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !InArray2(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {cl.FindAndAddContactsByMidV2(con);time.Sleep(3 * time.Second)}
			}
		}
		n += 1
	}
}

func clearCon() {
	n := 0
	for _, cl := range ClientBot {
		fl, _ := cl.GetAllContactIds()
		for _, x := range fl {
			if !InArray2(Squadlist, x) && !InArray2(DEVELOPER, x){
				cl.UnFriend(x)
				time.Sleep(2 * time.Second)
			}	
		}
		n += 1
	}
}
func RemBanFriends(client *linetcr.Account, to string) {
	defer panicHandle("RemBanFriends")
	if AllowDoOnce == 0 {
		donedel := []string{}
		for _, cl := range ClientBot {
			friendz, _ := cl.GetAllContactIds()
			for _, con := range friendz {
				if !linetcr.InArrayCl(linetcr.KickBanChat, cl) && !cl.Frez {
					r, _ := cl.GetHomeProfile(con)
					if linetcr.GetBannedChat(r) == 1 {
						cl.UnFriend(con)
						if !InArray2(donedel, con) {
							donedel = append(donedel, con)
						}
					}
				} else {
					r, _ := client.GetHomeProfile(con)
					if linetcr.GetBannedChat(r) == 1 {
						cl.UnFriend(con)
						if !InArray2(donedel, con) {
							donedel = append(donedel, con)
						}
					}
				}
			}
		}
		if len(donedel) == 0 {
			client.SendMessage(to, fancy("Nothing Deleted || No C_Ban Friends\n"))
		} else {
			DataMention(to, "✭ Unfriend Banz ✭\n", donedel)
		}
		AllowDoOnce++
	} else {
		client.SendMessage(to, fancy("No Data For C_Ban Friends"))
	}
}
func DataMention(to string, mtxt string, targetlist []string) {
	defer panicHandle("DataMention")
	Room := linetcr.GetRoom(to)
	if Room.Id != "" {
		Cans := Room.Client
		if len(Cans) != 0 {
			memlist := []string{}
			for _, mid := range targetlist {
				if !InArray2(Squadlist, mid) {
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

func ReloginProgram() error {
	file, err := osext.Executable()
	if err != nil {
		return err
	}
	err = syscall.Exec(file, os.Args, os.Environ())
	if err != nil {
		return err
	}
	return nil
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
	client.SendMessage(to, fancy(times))
}
func SelectallBot(client *linetcr.Account, to string) ([]*linetcr.Account, error) {
	err, _, memlist := client.GetGroupMembers(to)
	if err != nil {
		return nil, err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			exe = append(exe, cl)
		}
	}
	if len(exe) != 0 {
		return exe, err
	}
	return nil, err
}

func Setcancelto(client *linetcr.Account, to string, invits []string) {
	defer panicHandle("Setcancelto")
	for _, x := range invits {
		if linetcr.IsPending(client, to, x) == true {
			client.CancelChatInvitations(to, []string{x})
		}
	}
}

func cekOp(optime int64) bool {
	for _, tar := range oplist {
		if tar == optime {
			return false
		}
	}
	oplist = append(oplist, optime)
	return true
}

func cekOpinvite(optime int64) bool {
	for _, tar := range oplistinvite {
		if tar == optime {
			return false
		}
	}
	oplistinvite = append(oplistinvite, optime)
	return true
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
	s = Remove(s, e)
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
			if !InArray2(saodd, mentionlist[a]) && !InArray2(Squadlist, mentionlist[a]) {
				saodd = append(saodd, mentionlist[a])
			}

		}
		return saodd
	} else if rplay != "" {
		if !InArray2(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "rplay" {
		if !InArray2(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "lmid" {
		g, ok := Lastmid.Get(to)
		if !ok {
			g = [][]string{}
			Lastmid.Set(to, g)
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
		g, ok := Lastmessage.Get(to)
		if !ok {
			g = [][]string{}
			Lastmessage.Set(to, g)
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
		g, ok := Lastinvite.Get(to)
		if !ok {
			g = []string{}
			Lastinvite.Set(to, g)
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
		g, ok := Lastkick.Get(to)
		if !ok {
			g = []string{}
			Lastkick.Set(to, g)
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
		g, ok := Lastcancel.Get(to)
		if !ok {
			g = []string{}
			Lastcancel.Set(to, g)
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
		g, ok := Lastupdate.Get(to)
		if !ok {
			g = []string{}
			Lastupdate.Set(to, g)
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
		g, ok := Lastjoin.Get(to)
		if !ok {
			g = []string{}
			Lastjoin.Set(to, g)
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
		g, ok := Lasttag.Get(to)
		if !ok {
			g = []string{}
			Lasttag.Set(to, g)
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
		g, ok := Lastcon.Get(to)
		if !ok {
			g = []string{}
			Lastcon.Set(to, g)
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
		g, ok := Lastleave.Get(to)
		if !ok {
			g = []string{}
			Lastleave.Set(to, g)
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
		if !InArray2(saodd, sender) {
			saodd = append(saodd, sender)
		}
		return saodd
	} else if typec == "@all" {
		Member := client.GetChatListMem(to)
		for i := 0; i < len(Member); i++ {
			if !InArray2(saodd, Member[i]) {
				saodd = append(saodd, Member[i])
			}
		}
		return saodd
	} else if typec == "lbanlist" {
		for i := 0; i < len(Banned.Banlist); i++ {
			if !InArray2(saodd, Banned.Banlist[i]) {
				saodd = append(saodd, Banned.Banlist[i])
			}
		}
		return saodd
	} else if typec == "pend" {
		tcr := strings.Replace(typec, "pend ", "", 1)
		numb, _ := strconv.Atoi(tcr)
		_, _, pind := client.GetChatList(to)
		if numb == 0 {
			for n, i := range pind {
				if n < CancelPend {
					if !InArray2(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	} else if typec == "numpend" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !InArray2(saodd, i) {
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
			if !InArray2(saodd, i) {
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
					if !InArray2(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	}
	return saodd
}
func savejoin(Pelaku string, Optime int64) {
	defer linetcr.PanicOnly()
	ix := IndexOf(Detectjoin.User, Pelaku)
	if ix == -1 {
		Detectjoin.User = append(Detectjoin.User, Pelaku)
		Detectjoin.Time = append(Detectjoin.Time, Optime)
	} else {
		Detectjoin.Time[ix] = Optime
	}
}
func LogGet(op *SyncService.Operation) {
	defer linetcr.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	korban := op.Param3
	if tipe == 124 || tipe == 123 {
		var invites []string
		if tipe == 124 {
			invites = strings.Split(korban, "\x1e")
		} else {
			invites = strings.Split(pelaku, "\x1e")
		}
		ll := len(invites)
		if ll != 0 {
			g, ok := Lastinvite.Get(op.Param1)
			if !ok {
				Lastinvite.Set(op.Param1, invites)
			} else {
				c := g.([]string)
				for _, can := range invites {
					c = AppendLast(c, can)
				}
				Lastinvite.Set(op.Param1, c)
			}
		}

	} else if tipe == 133 {
		g, ok := Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 132 {
		g, ok := Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 130 {
		g, ok := Lastjoin.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastjoin.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastjoin.Set(op.Param1, c)
		}
	} else if tipe == 125 {
		g, ok := Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 126 {
		g, ok := Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 122 {
		g, ok := Lastupdate.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastupdate.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastupdate.Set(op.Param1, c)
		}

	} else if tipe == 128 {
		g, ok := Lastleave.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			Lastleave.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			Lastleave.Set(op.Param1, c)
		}

	} else if tipe == 26 {
		var MentionMsg = MentionList(op)
		msg := op.Message
		if InArray2(Squadlist, msg.From_) {
			return
		}
		if len(MentionMsg) != 0 {
			g, ok := Lasttag.Get(msg.To)
			if !ok {
				g = MentionMsg
				Lasttag.Set(msg.To, g)
			} else {
				c := g.([]string)
				for _, men := range MentionMsg {
					c = AppendLast(c, men)
				}
				Lasttag.Set(msg.To, c)
			}
		} else if msg.ContentType == 13 {
			mids := msg.ContentMetadata["mid"]
			g, ok := Lastcon.Get(msg.To)
			if !ok {
				g = []string{mids}
				Lastcon.Set(msg.To, g)
			} else {
				c := g.([]string)
				c = AppendLast(c, mids)
				Lastcon.Set(msg.To, c)
			}

		} else if msg.ContentType == 7 {
			var ids []string
			var pids []string
			zx := msg.ContentMetadata
			vok, cook := zx["REPLACE"]
			if cook {
				ress := gjson.Get(vok, "sticon")
				mp := ress.Map()
				yo := mp["resources"]
				vls := yo.Array()
				for _, vl := range vls {
					mm := vl.Map()
					pids = append(pids, mm["productId"].String())
					ids = append(ids, mm["sticonId"].String())
				}
			} else {
				ids = []string{zx["STKID"]}
				pids = []string{zx["STKPKGID"]}
			}

			g, ok := Laststicker.Get(msg.To)
			if !ok {
				g = []*Stickers{&Stickers{Id: ids[0], Pid: pids[0]}}
				Laststicker.Set(msg.To, g)
			} else {
				c := g.([]*Stickers)
				c = AppendLastSticker(c, &Stickers{Id: ids[0], Pid: pids[0]})
				Laststicker.Set(msg.To, c)
			}

		} else if msg.ContentType == 0 {
			if strings.Contains(msg.Text, "u") {
				regex, _ := regexp.Compile(`u\w{32}`)
				links := regex.FindAllString(msg.Text, -1)
				mmd := []string{}
				for _, a := range links {
					if len(a) == 33 {
						mmd = append(mmd, a)
					}
				}
				if len(mmd) != 0 {
					g, ok := Lastmid.Get(msg.To)
					if !ok {
						g = [][]string{mmd}
						Lastmid.Set(msg.To, g)
					} else {
						c := g.([][]string)
						c = AppendLastD(c, mmd)
						Lastmid.Set(msg.To, c)
					}
				}
			}
			sender := op.Message.From_
			if MemUser(op.Param1, sender) && msg.ToType == 2 {
				mmu := []string{}
				if !InArray2(mmu, sender) {
					mmu = append(mmu, sender)
				}
				if len(mmu) != 0 {
					g, ok := Lastmessage.Get(msg.To)
					if !ok {
						g = [][]string{mmu}
						Lastmessage.Set(msg.To, g)
					} else {
						c := g.([][]string)
						c = AppendLastD(c, mmu)
						Lastmessage.Set(msg.To, c)
					}
				}
			}
		}
	}
}
func BanAll(memlist []string, Group string) {
	ilen := len(memlist)
	for i := 0; i < ilen; i++ {
             if MemUser(Group, memlist[i]) {
	            Banned.AddBan(memlist[i])
	      }
       }
}
func botDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	return fmt.Sprintf("%2d Days, %2d Hours, %2d Mins.", h/24, h%24, m)
}
func (self *kickop) ceko(pelaku int64) bool {
	if !InArrayInt64(self.Opinv, pelaku) {
		self.Opinv = append(self.Opinv, pelaku)
		return true
	}
	return false
}

func (self *kickop) cek(pelaku string) bool {
	if !InArray2(self.Kick, pelaku) {
		self.Kick = append(self.Kick, pelaku)
		return true
	}
	return false
}

func (self *kickop) del(pelaku string) {
	self.Kick = Remove(self.Kick, pelaku)
}

func (self *kickop) ceki(pelaku string) bool {
	defer linetcr.PanicOnly()
	if !InArray2(self.Inv, pelaku) {
		self.Inv = append(self.Inv, pelaku)
		return true
	}
	return false
}

func (self *kickop) deli(pelaku string) {
	self.Inv = Remove(self.Inv, pelaku)
}

func (self *kickop) clear() {
	self.Inv = []string{}
	self.Kick = []string{}
	self.Opinv = []int64{}
}

func Checklistexpel(client *linetcr.Account, to string, targets []string, pl int, sinder string) {
	Room := linetcr.GetRoom(to)
	if len(targets) > 1 {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if InArray2(DEVELOPER, from) {
				if !InArray2(DEVELOPER, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetCreator(from) {
				if InArray2(DEVELOPER, sinder) {
					target = append(target, from)
					UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaker(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					UserBot.DelMaker(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetSeller(from) {
				if SendMymaker(sinder) {
					target = append(target, from)
					UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = Remove(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = Remove(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Expeled from Buyer\n"
			} else if pl == 2 {
				list += "Expeled from Owner\n"
			} else if pl == 3 {
				list += "Expeled from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Maker\n"
			} else if pl == 10 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				logAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				logAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				logAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				logAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				logAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				logAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				logAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, fancy(list))
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, fancy(list))
		}
	} else {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if InArray2(DEVELOPER, from) {
				if !InArray2(DEVELOPER, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetCreator(from) {
				if InArray2(DEVELOPER, sinder) {
					target = append(target, from)
					UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaker(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					UserBot.DelMaker(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetSeller(from) {
				if SendMymaker(sinder) {
					target = append(target, from)
					UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = Remove(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if InArray2(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = Remove(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Removed from Buyer\n"
			} else if pl == 2 {
				list += "Removed from Owner\n"
			} else if pl == 3 {
				list += "Removed from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Maker\n"
			} else if pl == 10 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				logAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				logAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				logAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				logAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				logAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				logAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				logAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, fancy(list))
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, fancy(list))
		}
	}
}
func Checkqr() {
	Qrwar = true
	time.Sleep(1 * time.Second)
	Qrwar = false
}
func Cmdlistcheck() string {
	list2 := "𝗟𝗶𝘀𝘁 𝗖𝗺𝗱:\n\n"
	list := ""
	if Commands.Botname != "" {
		list += fmt.Sprintf(" - Botname: %s\n", Commands.Botname)
	}
	if Commands.Upallimage != "" {
		list += fmt.Sprintf(" - Upallimage: %s\n", Commands.Upallimage)
	}
	if Commands.Upallcover != "" {
		list += fmt.Sprintf(" - Upallcover: %s\n", Commands.Upallcover)
	}
	if Commands.Unsend != "" {
		list += fmt.Sprintf(" - Unsend: %s\n", Commands.Unsend)
	}
	if Commands.Upvallimage != "" {
		list += fmt.Sprintf(" - Upvallimage: %s\n", Commands.Upvallimage)
	}
	if Commands.Upvallcover != "" {
		list += fmt.Sprintf(" - Upvallcover: %s\n", Commands.Upvallcover)
	}
	if Commands.Appname != "" {
		list += fmt.Sprintf(" - Appname: %s\n", Commands.Appname)
	}
	if Commands.Useragent != "" {
		list += fmt.Sprintf(" - Useragent: %s\n", Commands.Useragent)
	}
	if Commands.Hostname != "" {
		list += fmt.Sprintf(" - Hostname: %s\n", Commands.Hostname)
	}
	if Commands.Friends != "" {
		list += fmt.Sprintf(" - Friends: %s\n", Commands.Friends)
	}
	if Commands.Adds != "" {
		list += fmt.Sprintf(" - Adds: %s\n", Commands.Adds)
	}
	if Commands.Limits != "" {
		list += fmt.Sprintf(" - Limits: %s\n", Commands.Limits)
	}
	if Commands.Addallbots != "" {
		list += fmt.Sprintf(" - Addallbots: %s\n", Commands.Addallbots)
	}
	if Commands.Addallsquads != "" {
		list += fmt.Sprintf(" - Addallsquads: %s\n", Commands.Addallsquads)
	}
	if Commands.Leave != "" {
		list += fmt.Sprintf(" - Leave: %s\n", Commands.Leave)
	}
	if Commands.Respon != "" {
		list += fmt.Sprintf(" - Respon: %s\n", Commands.Respon)
	}
	if Commands.Ping != "" {
		list += fmt.Sprintf(" - Ping: %s\n", Commands.Ping)
	}
	if Commands.Count != "" {
		list += fmt.Sprintf(" - Count: %s\n", Commands.Count)
	}
	if Commands.Limitout != "" {
		list += fmt.Sprintf(" - 1111111: %s\n", Commands.Limitout)
	}
	if Commands.Access != "" {
		list += fmt.Sprintf(" - Access: %s\n", Commands.Access)
	}
	if Commands.Allbanlist != "" {
		list += fmt.Sprintf(" - Allbanlist: %s\n", Commands.Allbanlist)
	}
	if Commands.Allgaccess != "" {
		list += fmt.Sprintf(" - Allgaccess: %s\n", Commands.Allgaccess)
	}
	if Commands.Gaccess != "" {
		list += fmt.Sprintf(" - Gaccess: %s\n", Commands.Gaccess)
	}
	if Commands.Checkram != "" {
		list += fmt.Sprintf(" - Checkram: %s\n", Commands.Checkram)
	}
	if Commands.Backups != "" {
		list += fmt.Sprintf(" - Backups: %s\n", Commands.Backups)
	}
	if Commands.Upimage != "" {
		list += fmt.Sprintf(" - Upimage: %s\n", Commands.Upimage)
	}
	if Commands.Upcover != "" {
		list += fmt.Sprintf(" - Upcover: %s\n", Commands.Upcover)
	}
	if Commands.Upvimage != "" {
		list += fmt.Sprintf(" - Upvimage: %s\n", Commands.Upvimage)
	}
	if Commands.Upvcover != "" {
		list += fmt.Sprintf(" - Upvcover: %s\n", Commands.Upvcover)
	}
	if Commands.Bringall != "" {
		list += fmt.Sprintf(" - Bringall: %s\n", Commands.Bringall)
	}
	if Commands.Purgeall != "" {
		list += fmt.Sprintf(" - Purgeall: %s\n", Commands.Purgeall)
	}
	if Commands.Banlist != "" {
		list += fmt.Sprintf(" - Banlist: %s\n", Commands.Banlist)
	}
	if Commands.Clearban != "" {
		list += fmt.Sprintf(" - Clearban: %s\n", Commands.Clearban)
	}
	if Commands.Stayall != "" {
		list += fmt.Sprintf(" - Stayall: %s\n", Commands.Stayall)
	}
	if Commands.Clearchat != "" {
		list += fmt.Sprintf(" - Clearchat: %s\n", Commands.Clearchat)
	}
	if Commands.Here != "" {
		list += fmt.Sprintf(" - Here: %s\n", Commands.Here)
	}
	if Commands.Speed != "" {
		list += fmt.Sprintf(" - Speed: %s\n", Commands.Speed)
	}
	if Commands.Status != "" {
		list += fmt.Sprintf(" - Status: %s\n", Commands.Status)
	}
	if Commands.Tagall != "" {
		list += fmt.Sprintf(" - Tagall: %s\n", Commands.Tagall)
	}
	if Commands.Kick != "" {
		list += fmt.Sprintf(" - Kick: %s\n", Commands.Kick)
	}
	if Commands.Max != "" {
		list += fmt.Sprintf(" - Protect Max: %s\n", Commands.Max)
	}
	if Commands.None != "" {
		list += fmt.Sprintf(" - Protect None: %s\n", Commands.None)
	}
	if Commands.Kickall != "" {
		list += fmt.Sprintf(" - Kickall: %s\n", Commands.Kickall)
	}
	if Commands.Cancelall != "" {
		list += fmt.Sprintf(" - Cancelall: %s\n", Commands.Cancelall)
	}
	if list != "" {
		return list2 + list

	} else {
		return "Not found set Cmd.\n"
	}
}
func bToMb(b uint64) uint64 {
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
func LeaveallGroups(client *linetcr.Account, to string) []string {
	allg := []string{}
	for i := range ClientBot {
		groups, _ := ClientBot[i].GetGroupIdsJoined()
		grup, _ := ClientBot[i].GetGroups(groups)
		for _, gi := range grup {
			if gi.ChatMid != to {
				ClientBot[i].LeaveGroup(gi.ChatMid)
				time.Sleep(1 * time.Second)
				if !InArray2(allg, gi.ChatMid) {
					allg = append(allg, gi.ChatMid)
				}
			}
		}
	}
	return allg
}
func logAccess(client *linetcr.Account, group, from, tipe string, targets []string, tempat int32) {
	defer panicHandle("logAccess")
	if !LogMode || SendMyseller(from) {
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
	} else if tipe == "unban" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! unbaned %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! unbaned %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@! promoted owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! promoted owner %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unowner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! expeled owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! expeled owner %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@! promoted bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! promoted bot %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unbot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! expeled bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! expeled bot %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@! muted %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! muted %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unmute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@! unmuted %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@! unmuted %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@!  fuck %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  fuck %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@!  promoted Master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  promoted Master %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unmaster" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled Master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled Master %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@!  promoted admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  promoted admin %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "unadmin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled admin %v user's from \n%s\n\nTarget:", len(targets), nm)
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
		ts += fmt.Sprintf("@!  promoted gowner %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "ungowner" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  expeled gowner %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "gadmin" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  promoted gadmin %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "ungadmin" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  expeled gadmin %v user's from \n%s\n\nTarget:", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "expel" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  expeled access %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  expeled access %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearowner" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all owner %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all owner %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearmaster" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all master %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all master %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearadmin" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all admin %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all admin %v user's from \n%s\n\nTarget:", len(targets), nm)
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
			ts += fmt.Sprintf("@!  cleared all banlist %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all banlist %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearlock" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all locklist %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all locklist %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearbot" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all bot %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all bot %v user's from \n%s\n\nTarget:", len(targets), nm)
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		}
	} else if tipe == "clearmute" {
		if len(targets) == 0 {
			return
		}
		if tempat == 1 {
			ts += fmt.Sprintf("@!  cleared all mutelist %v user's:\n", len(targets))
			cuh, _ := client.GetContacts(targets)
			for _, prs := range cuh {
				name := prs.DisplayName
				ts += fmt.Sprintf("\n   %s", name)
			}
		} else {
			ts += fmt.Sprintf("@!  cleared all mutelist %v user's from \n%s\n\nTarget:", len(targets), nm)
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
		ts += fmt.Sprintf("@!  kick %v user's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "cancel" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  cancel %v invitation's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "invite" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  invite %v user's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "addfriends" {
		if len(targets) == 0 {
			return
		}
		ts += fmt.Sprintf("@!  addfriends %v user's from\n%s\n\n", len(targets), nm)
		cuh, _ := client.GetContacts(targets)
		for _, prs := range cuh {
			name := prs.DisplayName
			ts += fmt.Sprintf("\n   %s", name)
		}
	} else if tipe == "Kickall" {
		ts += fmt.Sprintf("@!  Nukeall %v user's from\n%s", len(targets), nm)
	} else if tipe == "purgeall" {
		ts += fmt.Sprintf("@!  purgeall %v user's from\n%s", len(targets), nm)
	} else if tipe == "purge" {
		ts += fmt.Sprintf("@!  purge %v user's in:\n%s", len(targets), nm)
	} else if tipe == "cancelall" {
		ts += fmt.Sprintf("@!  cancelall %v invitation's from\n%s", len(targets), nm)
	} else if tipe == "leave" {
		ts += fmt.Sprintf("@!  bot's leave from\n%s", nm)
	} else if tipe == "bringbot" {
		ts += fmt.Sprintf("@!  Invite bot's\n%s", nm)
	} else if tipe == "addfrind" {
		ts += fmt.Sprintf("@! he added %v as friend\nMid : \n%s", group, from)
	} else if tipe == "callme" {
		ts += fmt.Sprintf("@! he call me %v as friend\nMid : \n%s", group, from)
	} else if tipe == "deleteaccount" {
		ts += fmt.Sprintf("@! he call me %v as friend\nMid : \n%s", group, from)
	}
	room := linetcr.GetRoom(LogGroup)
	if len(room.Client) != 0 {
		exe, err := SelectBot(room.Client[0], LogGroup)
		if err == nil {
			if exe != nil {
				exe.SendMention(LogGroup, ts, []string{from})
			}
		} else {
			LogMode = false
			LogGroup = ""
		}
	}
}

func checkunbanbots(client *linetcr.Account, to string, targets []string, pl int, sinder string) {
	room := linetcr.GetRoom(to)
	target := []string{}
	for _, from := range targets {
		if Banned.GetFuck(from) {
			target = append(target, from)
			Banned.DelFuck(from)
		} else if Banned.GetBan(from) {
			target = append(target, from)
			Banned.DelBan(from)
		} else if InArray2(room.Gban, from) {
			target = append(target, from)
			Ungban(to, from)
		} else if Banned.GetMute(from) {
			target = append(target, from)
			Banned.DelMute(from)
		}
	}
	if len(target) != 0 {
		list := ""
		if pl == 1 {
			list += "Removed from banlist:\n"
		} else if pl == 2 {
			list += "Removed from fucklist:\n"
		} else if pl == 3 {
			list += "Removed from gbanlist:\n"
		} else if pl == 4 {
			list += "Removed from mutelist:\n"
		}
		for i := range target {
			list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
		}
		client.SendPollMention(to, list, target)
		if pl == 1 {
			logAccess(client, to, sinder, "unban", target, 2)
		} else if pl == 2 {
			logAccess(client, to, sinder, "unfuck", target, 2)
		} else if pl == 3 {
			logAccess(client, to, sinder, "ungban", target, 2)
		} else if pl == 4 {
			logAccess(client, to, sinder, "unmute", target, 2)
		}
	} else {
		list := ""
		if pl == 1 {
			list += "User(s) not in banlist.\n"
		} else if pl == 2 {
			list += "User(s) not in fucklist.\n"
		} else if pl == 3 {
			list += "User(s) not in gbanlist.\n"
		} else if pl == 4 {
			list += "User(s) not in mutelist.\n"
		}
		client.SendMessage(to, fancy(list))
	}
}
func CheckExprd(s *linetcr.Account, to string, sender string) bool {
	base := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
	d := fmt.Sprintf("%v", Data.Dalltime)
	has := strings.Split(d, "-")
	has2 := strings.Split(has[2], "T")
	yy, _ := strconv.Atoi(has[0])
	mm, _ := strconv.Atoi(has[1])
	timeup, _ := strconv.Atoi(has2[0])
	batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
	if batas.Before(base) {
		if !SendMycreator(sender) {
			s.SendMessage(to, fancy("Sorry your bots is expired, Please Contact with our Creator to renew your squad. ;-)"))
			return false
		}
		return true
	}
	return true
}
func fmtDurations(d time.Duration) string {
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
func CekDuedate() time.Time {
	bod := string(Data.Dalltime)
	date, _ := time.Parse(time.RFC3339, bod)
	return date
}
func CheckLastActive(client *linetcr.Account, targets string) string {
	list := ""
	mek, tu := LastActive.Get(targets)
	if tu {
		asu := mek.(*SyncService.Operation)
		if asu.Type == 55 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			list += "- LastActive: " + sp + "\n- Type: Read Message\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 124 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			invites := strings.Split(asu.Param3, "\x1e")
			nos := 0
			her := ""
			for _, ampemng := range invites {
				nos += 1
				pr, _ := client.GetContact(ampemng)
				her += fmt.Sprintf("\n  %v. %v", nos, pr.DisplayName)
			}
			list += "- LastActive: " + sp + "\n- Type: Invited member\n- Group: " + names1 + "\n- Target: " + her + "\n\n"
		} else if asu.Type == 133 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- LastActive: " + sp + "\n- Type : Kick member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 126 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- LastActive: " + sp + "\n- Type: Cancel member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 26 {
			msg := asu.Message
			if msg.ToType == 2 {
				names1, _ := client.GetGroupMember(msg.To)
				cok := asu.CreatedTime / 1000
				i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
				tm := time.Unix(i, 0)
				ss := time.Since(tm)
				sp := fmtDuration(ss)
				tx := ""
				if msg.ContentType == 0 {
					tx = msg.Text
				} else {
					tx = "Non Text Message"
				}
				list += "- LastActive: " + sp + "\n- Type: Send Message\n- Group: " + names1 + "\n- Message: " + tx + "\n\n"
			}
		} else if asu.Type == 130 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			list += "- LastActive: " + sp + "\n- Type: Join Group\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 122 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := fmtDuration(ss)
			var ti string
			if asu.Param3 == "4" {
				g, _ := client.GetGroup3(asu.Param1)
				if g.Extra.GroupExtra.PreventedJoinByTicket == false {
					ti = "Open qr"
				} else {
					ti = "Close qr"
				}
			} else if asu.Param3 == "1" {
				ti = "Change Group Name"
			}
			list += "- LastActive: " + sp + "\n- Type: Update Group\n- Group: " + names1 + "\n- Type: " + ti + "\n\n"
		}
	}
	return list
}
func InArray2(ArrList []string, rstr string) bool {
	for _, x := range ArrList {
		if x == rstr {
			return true
		}
	}
	return false
}

func Checklistaccess(client *linetcr.Account, group string, targets []string, pl int, sinder string) {
	Room := linetcr.GetRoom(group)
	if pl == 12 {
		countr := 0
		countr1 := 0
		list := "Account Info: \n\n"
		for n, xx := range targets {
			new := client.Getcontactuser(xx)
			if new != nil {
				list += "Name: Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("Name: %v \n", x.DisplayName)
				status := "status: None\n\n"
				if InArray2(DEVELOPER, targets[n]) {
					status = "status: Developer\n\n"
				} else if UserBot.GetCreator(targets[n]) {
					status = "status: Creators\n\n"
				} else if UserBot.GetMaker(targets[n]) {
					status = "status: Makers\n\n"
				} else if UserBot.GetBuyer(targets[n]) {
					status = "status: Buyer\n\n"
				} else if UserBot.GetOwner(targets[n]) {
					status = "status: Owner\n\n"
				} else if UserBot.GetMaster(targets[n]) {
					status = "status: Master\n\n"
				} else if UserBot.GetAdmin(targets[n]) {
					status = "status: Admin\n\n"
				} else if InArray2(Room.Gowner, targets[n]) {
					status = "status: GroupOwnar\n\n"
				} else if InArray2(Room.Gadmin, targets[n]) {
					status = "status: GroupAdmin\n\n"
				} else if UserBot.GetBot(targets[n]) {
					status = "status: Bot\n\n"
				} else if Banned.GetFuck(targets[n]) {
					status = "status: Fuck\n\n"
				} else if Banned.GetBan(targets[n]) {
					status = "status: Ban\n\n"
				} else if Banned.GetMute(targets[n]) {
					status = "status: Mute\n\n"
				} else if InArray2(Room.Gban, targets[n]) {
					status = "status: Groupban\n\n"
				} else if InArray2(Squadlist, targets[n]) {
					status = "status: My team\n\n"
				} else if UserBot.GetSeller(targets[n]) {
					status = "status: My Seller\n\n"
				}
				list += status
				if !InArray2(checkHaid, targets[n]) {
					new := CheckLastActive(client, targets[n])
					list += new
				}
				listGroup := "\nMember of:\n"
				listPinde := "\nPending of:\n"
				grs, _ := client.GetGroupIdsJoined()
				groups, _ := client.GetGroups(grs)
				for _, x := range groups {
					if linetcr.IsMembers(client, x.ChatMid, targets[n]) == true {
						countr = countr + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listGroup += nm + "\n"
					}
					if linetcr.IsPending(client, x.ChatMid, targets[n]) == true {
						countr1 = countr1 + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listPinde += nm + "\n"
					}
				}
				if countr != 0 {
					list += fmt.Sprintf("Groups: %v\n", countr)

				} else {
					list += "Groups: 0\n"
				}
				if countr1 != 0 {
					list += fmt.Sprintf("Pendings: %v\n", countr1)
				} else {
					list += "Pendings: 0\n"
				}
				if countr != 0 {
					if !InArray2(checkHaid, targets[n]) {
						list += listGroup
					}
				}
				if countr1 != 0 {
					if !InArray2(checkHaid, targets[n]) {
						list += listPinde
					}
				}

			}
		}
		client.SendMessage(group, fancy(list))
	} else if pl == 16 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			new := client.Getcontactuser(xx)
			if new != nil {
				list += rengs + ". Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

			}
		}
		client.SendMessage(group, fancy(list))
	} else if pl == 14 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			new := client.Getcontactuser(xx)
			if new != nil {
				list += rengs + ". Closed Account \n"
			} else {
				x, _ := client.GetContact(xx)
				list += fmt.Sprintf("%v. %v\n_%v\n", n+1, x.DisplayName, targets[n])

			}
		}
		client.SendMessage(group, fancy(list))
	} else {
		if len(targets) > 1 {
			Developer := []string{}
			creator := []string{}
			buyer := []string{}
			owner := []string{}
			master := []string{}
			admin := []string{}
			gowner := []string{}
			gadmin := []string{}
			squad := []string{}
			bot := []string{}
			ban := []string{}
			fuck := []string{}
			mute := []string{}
			Gban := []string{}
			Glist := []string{}
			Maker := []string{}
			Seller := []string{}
			for _, from := range targets {
				if MemUser(group, from) && !MemBan2(group, from) {
					if !InArray2(Glist, from) {
						Glist = append(Glist, from)
					}
				} else if UserBot.GetCreator(from) {
					creator = append(creator, from)
				} else if UserBot.GetMaker(from) {
					Maker = append(Maker, from)
				} else if UserBot.GetSeller(from) {
					Seller = append(Seller, from)
				} else if InArray2(DEVELOPER, from) {
					Developer = append(Developer, from)
				} else if UserBot.GetBuyer(from) {
					buyer = append(buyer, from)
				} else if UserBot.GetOwner(from) {
					owner = append(owner, from)
				} else if UserBot.GetMaster(from) {
					master = append(master, from)
				} else if UserBot.GetAdmin(from) {
					admin = append(admin, from)
				} else if InArray2(Room.Gowner, from) {
					gowner = append(gowner, from)
				} else if InArray2(Room.Gadmin, from) {
					gadmin = append(gadmin, from)
				} else if UserBot.GetBot(from) {
					bot = append(bot, from)
				} else if Banned.GetFuck(from) {
					fuck = append(fuck, from)
				} else if Banned.GetBan(from) {
					ban = append(ban, from)
				} else if Banned.GetMute(from) {
					mute = append(mute, from)
				} else if InArray2(Room.Gban, from) {
					Gban = append(Gban, from)
				} else if InArray2(Squadlist, from) {
					squad = append(squad, from)
				}
			}
			list2 := ""
			if len(Glist) != 0 {
				if pl == 1 {
					list2 += "Promoted as Buyer:\n\n"
				} else if pl == 2 {
					list2 += "Promoted as Owner:\n\n"
				} else if pl == 3 {
					list2 += "Promoted as Master:\n\n"
				} else if pl == 4 {
					list2 += "Promoted as Admin:\n\n"
				} else if pl == 5 {
					list2 += "Promoted as Bot:\n\n"
				} else if pl == 6 {
					list2 += "Promoted as Gowner:\n\n"
				} else if pl == 7 {
					list2 += "Promoted as Gadmin\n\n"
				} else if pl == 8 {
					list2 += "Added to banlist:\n\n"
				} else if pl == 9 {
					list2 += "Added to fucklist:\n\n"
				} else if pl == 10 {
					list2 += "Added to gbanlist:\n\n"
				} else if pl == 11 {
					list2 += "Added to mutelist:\n\n"
				} else if pl == 13 {
					list2 += "Added to Makerlist:\n\n"
				} else if pl == 15 {
					list2 += "Added to Creatorlist:\n\n"
				} else if pl == 17 {
					list2 += "Added to Sellerlist:\n\n"
				} else if pl == 18 {
					list2 += "Added to Friendlist:\n\n"
				}
				for n, xx := range Glist {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list2 += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list2 += rengs + ". " + x.DisplayName + "\n"
						if pl == 1 {
							UserBot.AddBuyer(xx)
						} else if pl == 2 {
							UserBot.AddOwner(xx)
						} else if pl == 13 {
							UserBot.AddMaker(xx)
						} else if pl == 15 {
							UserBot.AddCreator(xx)
						} else if pl == 3 {
							UserBot.AddMaster(xx)
						} else if pl == 4 {
							UserBot.AddAdmin(xx)
						} else if pl == 5 {
							UserBot.AddBot(xx)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, xx)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, xx)
						} else if pl == 8 {
							Banned.AddBan(xx)
						} else if pl == 9 {
							Banned.AddBan(xx)
						} else if pl == 10 {
							Addgban(xx, group)
						} else if pl == 11 {
							Banned.AddBan(xx)
						} else if pl == 17 {
							UserBot.AddSeller(xx)
						}
					}
				}
				if pl == 2 {
					logAccess(client, group, sinder, "owner", Glist, 2)
				} else if pl == 3 {
					logAccess(client, group, sinder, "master", Glist, 2)
				} else if pl == 4 {
					logAccess(client, group, sinder, "admin", Glist, 2)
				} else if pl == 5 {
					logAccess(client, group, sinder, "bot", Glist, 2)
				} else if pl == 6 {
					logAccess(client, group, sinder, "gowner", Glist, 2)
				} else if pl == 7 {
					logAccess(client, group, sinder, "gadmin", Glist, 2)
				} else if pl == 8 {
					logAccess(client, group, sinder, "ban", Glist, 2)
				} else if pl == 9 {
					logAccess(client, group, sinder, "fuck", Glist, 2)
				} else if pl == 10 {
					logAccess(client, group, sinder, "gban", Glist, 2)
				} else if pl == 11 {
					logAccess(client, group, sinder, "mute", Glist, 2)
				}
			}
			list := "Users have access:\n"
			if len(Developer) != 0 {
				list += "\nExist in Developers:\n"
				for n, xx := range Developer {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(creator) != 0 {
				list += "\nExist in Creators:\n"
				for n, xx := range creator {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Maker) != 0 {
				list += "\nExist in Makers:\n"
				for n, xx := range Maker {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Seller) != 0 {
				list += "\nExist in Sellers:\n"
				for n, xx := range Seller {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(buyer) != 0 {
				list += "\nExist in Buyers:\n"
				for n, xx := range buyer {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(owner) != 0 {
				list += "\nExist in Owners:\n"
				for n, xx := range owner {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(master) != 0 {
				list += "\nExist in Masters:\n"
				for n, xx := range master {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(admin) != 0 {
				list += "\nExist in Admins:\n"
				for n, xx := range admin {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gowner) != 0 {
				list += "\nExist in Gowners:\n"
				for n, xx := range gowner {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gadmin) != 0 {
				list += "\nExist in Gadmins:\n"
				for n, xx := range gadmin {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(bot) != 0 {
				list += "\nExist in Botlist\n"
				for n, xx := range bot {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(squad) != 0 {
				list += "\nExist in Squads:\n"
				for n, xx := range squad {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(ban) != 0 {
				list += "Exist in Banlist:\n"
				for n, xx := range ban {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(fuck) != 0 {
				list += "\nExist in Fucklist:\n"
				for n, xx := range fuck {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Gban) != 0 {
				list += "\nExist in Gbanlist:\n\n"
				for n, xx := range Gban {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(mute) != 0 {
				list += "\nExist in Mutelist:\n\n"
				for n, xx := range mute {
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(xx)
					if new != nil {
						list += rengs + ". Closed Account \n"
					} else {
						x, _ := client.GetContact(xx)
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if list != "Users have access:\n" {
				if list2 != "" {
					list2 += "\n"
				}
				client.SendMessage(group, fancy(list2+list))
			} else {
				client.SendMessage(group, fancy(list2))
			}
		} else {
			list := ""
			for n, from := range targets {
				if InArray2(DEVELOPER, from) {
					list += "User have access exist in Developer."
				} else if UserBot.GetCreator(from) {
					list += "User have access exist in Creator list."
				} else if UserBot.GetMaker(from) {
					list += "User have access exist in Maker list."
				} else if UserBot.GetSeller(from) {
					list += "User have access exist in seller list."
				} else if UserBot.GetBuyer(from) {
					list += "User have access exist in buyer list."
				} else if UserBot.GetOwner(from) {
					list += "User have access exist in owner list."
				} else if UserBot.GetMaster(from) {
					list += "User have access exist in master list."
				} else if UserBot.GetAdmin(from) {
					list += "User have access exist in admin list."
				} else if InArray2(Room.Gowner, from) {
					list += "User have access exist in gowner list."
				} else if InArray2(Room.Gadmin, from) {
					list += "User have access exist in gadmin list."
				} else if UserBot.GetBot(from) {
					list += "User have access exist in bot list."
				} else if Banned.GetFuck(from) {
					list += "User have access exist in fuck list."
				} else if Banned.GetBan(from) {
					list += "User have access exist in ban list."
				} else if InArray2(Room.Gban, from) {
					list += "User have access exist in gban list."
				} else if InArray2(Squadlist, from) {
					list += "User have access exist in squad list."
				} else if Banned.GetMute(from) {
					list += "User have access exist in mute list."
				} else if MemUser(group, from) && !MemBan2(group, from) {
					if pl == 1 {
						list += "Promoted as Buyer:\n"
					} else if pl == 2 {
						list += "Promoted as Owner:\n"
					} else if pl == 3 {
						list += "Promoted as Master:\n"
					} else if pl == 4 {
						list += "Promoted as Admin:\n"
					} else if pl == 5 {
						list += "Promoted as Bot:\n"
					} else if pl == 6 {
						list += "Promoted as Gowner:\n"
					} else if pl == 7 {
						list += "Promoted as Gadmin:\n"
					} else if pl == 8 {
						list += "Added to banlist:\n"
					} else if pl == 9 {
						list += "Added to fucklist:\n"
					} else if pl == 10 {
						list += "Added to gbanlist:\n"
					} else if pl == 11 {
						list += "Added to mutelist:\n"
					} else if pl == 13 {
						list += "Added to Makerlist:\n"
					} else if pl == 15 {
						list += "Added to Creatorlist:\n"
					} else if pl == 17 {
						list += "Added to Sellerlist:\n"
					} else if pl == 18 {
						list += "Added to Friendlist:\n"
					}
					rengs := strconv.Itoa(n + 1)
					new := client.Getcontactuser(from)
					if new != nil {
						list += "\n   " + rengs + ". Closed Account"
					} else {
						x, _ := client.GetContact(from)
						list += "\n   " + rengs + ". " + x.DisplayName
						if pl == 1 {
							UserBot.AddBuyer(from)
						} else if pl == 2 {
							UserBot.AddOwner(from)
						} else if pl == 3 {
							UserBot.AddMaster(from)
						} else if pl == 4 {
							UserBot.AddAdmin(from)
						} else if pl == 5 {
							UserBot.AddBot(from)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, from)
						} else if pl == 13 {
							UserBot.AddMaker(from)
						} else if pl == 15 {
							UserBot.AddCreator(from)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, from)
						} else if pl == 8 {
							autokickban(client, group, from)
							Banned.AddBan(from)
						} else if pl == 9 {
							Banned.AddFuck(from)
						} else if pl == 10 {
							Addgban(from, group)
						} else if pl == 11 {
							Banned.AddMute(from)
						} else if pl == 17 {
							UserBot.AddSeller(from)
						}
					}
					if pl == 2 {
						logAccess(client, group, sinder, "owner", []string{from}, 2)
					} else if pl == 3 {
						logAccess(client, group, sinder, "master", []string{from}, 2)
					} else if pl == 4 {
						logAccess(client, group, sinder, "admin", []string{from}, 2)
					} else if pl == 5 {
						logAccess(client, group, sinder, "bot", []string{from}, 2)
					} else if pl == 6 {
						logAccess(client, group, sinder, "gowner", []string{from}, 2)
					} else if pl == 7 {
						logAccess(client, group, sinder, "gadmin", []string{from}, 2)
					} else if pl == 8 {
						logAccess(client, group, sinder, "ban", []string{from}, 2)
					} else if pl == 9 {
						logAccess(client, group, sinder, "fuck", []string{from}, 2)
					} else if pl == 10 {
						logAccess(client, group, sinder, "gban", []string{from}, 2)
					} else if pl == 11 {
						logAccess(client, group, sinder, "mute", []string{from}, 2)
					}
				}

			}
			client.SendMessage(group, fancy(list))
		}
	}
}

func Purgesip(Group string, cl *linetcr.Account) {
	defer panicHandle("purgesip")
	mem := make(chan []string)
	go func(m chan []string) {
		memlistss := []string{}
		_, memlists := cl.GetGroupMember(Group)
		for target := range memlists {
			if MemBan(Group, target) {
				memlistss = append(memlistss, target)
			}
		}
		m <- memlistss
	}(mem)
	Cans := linetcr.Actor(Group)
	ClAct := len(Cans)
	hajar := []string{}
	var client *linetcr.Account
	memlist := <-mem
	celek := len(memlist)
	if celek > MaxKick {
		if ClAct != 0 {
			z := celek / MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= ClAct {
					no = 0
				}
				if i != 0 {
					client = Cans[no]
				} else {
					client = cl
				}
				if i == z {
					hajar = memlist[i*MaxKick:]
				} else {
					hajar = memlist[i*MaxKick : (i+1)*MaxKick]
				}
				if len(hajar) != 0 {
					for _, target := range hajar {
						go client.DeleteOtherFromChats(Group, []string{target})
					}
				}
				no += 1
			}
		} else if !cl.Limited {
			for _, target := range memlist {
				go cl.DeleteOtherFromChats(Group, []string{target})
			}
		}
	} else if !cl.Limited {
		for _, target := range memlist {
			go cl.DeleteOtherFromChats(Group, []string{target})
		}
	}
}

func kickBl(client *linetcr.Account, to string) {
	defer panicHandle("detectBl")
	memlist := []string{}
	_, memlists := client.GetGroupMember(to)
	act := []*linetcr.Account{}
	for mid, _ := range memlists {
		if MemBan(to, mid) {
			memlist = append(memlist, mid)
		} else if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited {
				act = append(act, cl)
			}
		}
	}
	lact := len(act)
	if lact == 0 {
		return
	} else {
		sort.Slice(act, func(i, j int) bool {
			return act[i].KickPoint < act[j].KickPoint
		})
		celek := len(memlist)
		if celek < MaxKick || lact == 1 {
			cl := act[0]
			for _, mem := range memlist {
				go cl.DeleteOtherFromChat(to, mem)
			}
		} else {
			hajar := []string{}
			z := celek / MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= lact {
					no = 0
				}
				go func(to string, no int, i int, z int, memlist []string, act []*linetcr.Account) {
					client = act[no]
					if i == z {
						hajar = memlist[i*MaxKick:]
					} else {
						hajar = memlist[i*MaxKick : (i+1)*MaxKick]
					}
					if len(hajar) != 0 {
						for _, target := range hajar {
							go client.DeleteOtherFromChat(to, target)
						}
					}
				}(to, no, i, z, memlist, act)
				no += 1
			}
		}
	}
}

func GetSquad(tok *linetcr.Account, to string) []*linetcr.Account {
	defer panicHandle("GetSquad")
	nm, memlist, invitee := tok.GetChatList(to)
	Bots := []*linetcr.Account{}
	MIdbot := []string{}
	GoClint := []*linetcr.Account{}
	Gomid := []string{}
	for _, ym := range memlist {
		if InArray2(Squadlist, ym) {
			idx := GetKorban(ym)
			MIdbot = append(MIdbot, ym)
			Bots = append(Bots, idx)
		}
	}
	room := linetcr.GetRoom(to)
	room.Name = nm
	for _, ym := range invitee {
		if InArray2(Squadlist, ym) {
			Gomid = append(Gomid, ym)
			idx := GetKorban(ym)
			GoClint = append(GoClint, idx)
		}
	}
	room.AddSquad(MIdbot, Bots, GoClint, Gomid)
	return Bots
}
func qrGo22(client *linetcr.Account, cans []*linetcr.Account, to string) {
	numb := len(ClientBot)
	if numb > 0 && numb <= len(ClientBot) {
		GetSquad(client, to)
		room := linetcr.GetRoom(to)
		aa := len(room.Client)
		if aa > numb {
			c := aa - numb
			ca := 0
			list := append([]*linetcr.Account{}, room.Client...)
			sort.Slice(list, func(i, j int) bool {
				return list[i].KickPoint > list[j].KickPoint
			})
			for _, o := range list {
				o.LeaveGroup(to)
				ca = ca + 1
				if ca == c {
					break
				}
			}
			GetSquad(client, to)
		} else if aa < numb {
			ti, err := client.ReissueChatTicket(to)
			if err == nil {
				go client.UpdateChatQrV2(to, false)
				all := []*linetcr.Account{}
				room := linetcr.GetRoom(to)
				cuk := room.Client
				for _, x := range ClientBot {
					if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
						all = append(all, x)
					}
				}
				sort.Slice(all, func(i, j int) bool {
				return all[i].KickPoint < all[j].KickPoint
				})
				g := numb - aa
				var wg sync.WaitGroup
				wi := GetSquad(client, to)
				for i := 0; i < len(all); i++ {
					if i == g {
						break
					}
					l := all[i]
					if l != client && !linetcr.InArrayCl(wi, l) {
						wg.Add(1)
						go func() {
							l.AcceptTicket(to, ti)
							QrKick(client, to)
							wg.Done()
						}()
					}
				}
				wg.Wait()
				client.UpdateChatQrV2(to, true)
				GetSquad(client, to)
			}
		}
	}
}
func AutojoinQr22(client *linetcr.Account, to string) {
	numb := len(ClientBot)
	if numb > 0 && numb <= len(ClientBot) {
		GetSquad(client, to)
		room := linetcr.GetRoom(to)
		aa := len(room.Client)
		if aa > numb {
			c := aa - numb
			ca := 0
			list := append([]*linetcr.Account{}, room.Client...)
			sort.Slice(list, func(i, j int) bool {
				return list[i].KickPoint > list[j].KickPoint
			})
			for _, o := range list {
				o.LeaveGroup(to)
				ca = ca + 1
				if ca == c {
					break
				}
			}
			GetSquad(client, to)
		} else if aa < numb {
			ti, err := client.ReissueChatTicket(to)
			if err == nil {
				go client.UpdateChatQrV2(to, false)
				all := []*linetcr.Account{}
				room := linetcr.GetRoom(to)
				cuk := room.Client
				for _, x := range ClientBot {
					if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
						all = append(all, x)
					}
				}
				sort.Slice(all, func(i, j int) bool {
				return all[i].KickPoint < all[j].KickPoint
				})
				g := numb - aa
				var wg sync.WaitGroup
				wi := GetSquad(client, to)
				for i := 0; i < len(all); i++ {
					if i == g {
						break
					}
					l := all[i]
					if l != client && !linetcr.InArrayCl(wi, l) {
						wg.Add(1)
						go func() {
							l.AcceptTicket(to, ti)
							QrKick(client, to)
							wg.Done()
						}()
					}
				}
				wg.Wait()
				client.UpdateChatQrV2(to, true)
				GetSquad(client, to)
			}
		}
	}
}
func AutojoinQr(client *linetcr.Account, to string) {
	defer panicHandle("AutojoinQr")
	ti, err := client.ReissueChatTicket(to)
	if err == nil {
		go client.UpdateChatQrV2(to, false)
		all := []*linetcr.Account{}
		room := linetcr.GetRoom(to)
		cuk := room.Client
		for _, x := range ClientBot {
			if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
				all = append(all, x)
			}
		}
		sort.Slice(all, func(i, j int) bool {
			return all[i].KickPoint < all[j].KickPoint
		})
		var wg sync.WaitGroup
		wi := GetSquad(client, to)
		for i := 0; i < len(all); i++ {
			l := all[i]
			if l != client && !linetcr.InArrayCl(wi, l) {
				wg.Add(1)
				go func() {
					l.AcceptTicket(to, ti)
					wg.Done()
				}()
			}
		}
		wg.Wait()
		client.UpdateChatQrV2(to, true)
		GetSquad(client, to)
	}
}
func qrGo(cl *linetcr.Account, cans []*linetcr.Account, to string) {
	defer panicHandle("QR_go")
	Room := linetcr.GetRoom(to)
	mes := make(chan bool)
	go func() {
		err := cl.UpdateChatQrV2(to, false)
		if err != nil {
			mes <- false
		} else {
			mes <- true
		}
	}()
	Room.Qr = false
	var ticket string
	link, err := cl.ReissueChatTicket(to)
	if err == nil {
		ticket = link
	} else {
		ticket = "error"
	}
	var wg sync.WaitGroup
	if ticket != "error" && ticket != "" {
		ok := <-mes
		if !ok {
			return
		}
		for _, cc := range cans {
			wg.Add(1)
			go func(c *linetcr.Account) {
				err := c.AcceptTicket(to, ticket)
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}(cc)
		}
		wg.Wait()
		Room.Qr = true
	}
	if Room.Qr {
		go func() {
			err := cl.UpdateChatQrV2(to, true)
			if err != nil {
				mes <- true
			} else {
				mes <- false
			}
		}()
	}
}
func hstg(to, u string) {
	room := linetcr.GetRoom(to)
	if !InArray2(room.LeaveBack, u) {
		room.LeaveBack = append(room.LeaveBack, u)
	}
}
func contains(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
	}
	return false
}
func IsBlacklist(client *linetcr.Account, from string) bool {
	if contains(Banned.Banlist, from) == true {
		return true
	}
	return false
}
func IsBlacklist2(client *linetcr.Account, from string) bool {
	if contains(Banned.Locklist, from) == true {
		return true
	}
	return false
}
func IsMember(client *linetcr.Account, from string, groups string) bool {
	res := client.GetGroup(groups)
	memlist := res.Members
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}
//Backup_Qr
func AcceptTicketSimple(client *linetcr.Account, to string, ticketId string){
    runtime.GOMAXPROCS(cpu)
    go func(){
        client.AcceptTicket(to, ticketId)
    }()
    go func(){
        var wg sync.WaitGroup
        wg.Add(len(Banned.Banlist))
        for i:=0; i<len(Banned.Banlist); i++ {
            go func(i int){
                defer wg.Done()
                if linetcr.IsMembers(client, to, Banned.Banlist[i]) {
                    client.DeleteOtherFromChats(to, []string{Banned.Banlist[i]})
                }
            }(i)
        }
        wg.Wait()
    }()
    time.Sleep(500 * time.Nanosecond)            
}
func WarQr(client *linetcr.Account, to string){
	runtime.GOMAXPROCS(cpu)
	GetSquad(client, to)
    chat := client.GetChat([]string{to}, true, false)
    if chat != nil {
        cek := chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket
        if cek == true{go client.UpdateChatQrV2(to, false);time.Sleep(100 * time.Millisecond)}	
        ticket, err := client.ReissueChatTickets(to)
        if err == nil {
            link := ticket
            all := []*linetcr.Account{}
            room := linetcr.GetRoom(to)
            cuk := room.Client
            for _, x := range ClientBot {
                if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
                    all = append(all, x)}}
            sort.Slice(all, func(i, j int) bool {return all[i].KickPoint < all[j].KickPoint})
            var wg sync.WaitGroup
            wi := GetSquad(client, to)
            for i := 0; i < len(all); i++ {
                l := all[i]
                if l != client && !linetcr.InArrayCl(wi, l) {
                    wg.Add(1)
                    go func() {
                        SpamAcceptQR(l, to, link)
                        wg.Done()
                    }()
                }
            }
            wg.Wait()  
        }
        runtime.GOMAXPROCS(1)
    }
}
func SpamAcceptQR(client *linetcr.Account, to string, ticketId string) {
	runtime.GOMAXPROCS(100)
	go func(){AcceptTicketSimple(client, to, ticketId)}()
	go func(){WarQr(client, to)}()
	time.Sleep(200 * time.Nanosecond)
}

//Backup_124
func JKickFuck(client *linetcr.Account, to string) {
	c, _ := client.GetGroup3(to)
	zxc := c.Extra.GroupExtra.MemberMids
	for k, _ := range zxc {
		if IsBlacklist(client, k) == true {
			go func(to string, k string) {
				go client.DeleteOtherFromChats(to, []string{k})
			}(to,k)
		}
	}
}
func JCancelFuck(client *linetcr.Account, to string) {
	c, _ := client.GetGroup3(to)
	zxc := c.Extra.GroupExtra.InviteeMids
	for k, _ := range zxc {
		if IsBlacklist(client, k) == true {
			go func(to string, k string) {
				go client.CancelChatInvitations(to, []string{k})
			}(to,k)
		}
	}
}
func KickCansWar(client *linetcr.Account, to string) {
	go func() {
		JKickFuck(client, to)
	}()
	go func() {
		JCancelFuck(client, to)
	}()
}
func AcceptWar(client *linetcr.Account, to string, korban []string) {
	go func() {
		AccGroup(client, to);JoinLlinetcrBan(client, to);for _, i := range korban {go func(i string) {go client.InviteIntoGroupNormal(to, []string{i})}(i)}
	}()
}
func AccKickBan(client *linetcr.Account, to string) {
	go func() {
		AccGroup(client, to);JoinLlinetcrBan(client, to)
	}()
}
func AccGroup(client *linetcr.Account, to string) {
	go func() {client.NewacceptGroup(to)
	}()
	//go func() {client.AcceptGroupInvitationNormal(to)
	//}()
}
//Llinetcr_Banlist
func JoinLlinetcrBan(client *linetcr.Account, to string) {
	_, mem, pend := client.GetChatList(to)
	for _, k := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k)
		}
	}
	for _, i := range pend {
		if IsBlacklist(client, i) == true {
			go func(i string) {go client.CancelChatInvitations(to, []string{i})}(i)
		}
	}
	runtime.GOMAXPROCS(cpu)
}
func KickBan129(client *linetcr.Account, to string) {
	_, mem, pend := client.GetChatList(to)
	for _, k := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k)
		}
	}
	for _, i := range pend {
		if IsBlacklist(client, i) == true {
			go func(i string) {go client.CancelChatInvitations(to, []string{i})}(i)
		}
	}
	runtime.GOMAXPROCS(cpu)
}
//Backup_133
func groupBackupWar(client *linetcr.Account, to string) {
	for x := range Squadlist {if linetcr.IsMembers(client, to, Squadlist[x]) == true {if client.MID == Squadlist[x] {go func() {KickCansWar(client, to)}()};break} else {continue}
	}
	runtime.GOMAXPROCS(cpu)
}

func RunBot(client *linetcr.Account, ch chan int) {
	defer panicHandle("RunBot")
	runtime.GOMAXPROCS(cpu)
	client.Revision = -1
	for {
		multiFunc, err := client.SyncLoad(100)
		if err != nil || len(multiFunc) == 0 {
			continue
		}
		go func(fetch []*SyncService.Operation) {
			for _, op := range multiFunc {
				if op.Type == 124 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					rngcmd := GetComs(7, "invitebot")
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := linetcr.GetRoom(Group)
					if InArray2(invited, client.MID) {
						if linetcr.IoGOBot(Group, client) {
							if InArray2(client.Squads, user) {
								go func(client *linetcr.Account, Group string){
									go AccKickBan(client, Group)
								}(client, Group)
							} else if UserBot.GetBot(user) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if AutoPurge {
									go JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !InArrayInt64(cekGo, Optime) {cekGo = append(cekGo, Optime)
									AcceptJoin(client, Group)}
								}
							} else if GetCodeprem(rngcmd, user, Group) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if AutoPurge {
									go JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !InArrayInt64(cekGo, Optime) {cekGo = append(cekGo, Optime)
									AcceptJoin(client, Group)}
								}
							} else {
								grs, _ := client.GetGroupIdsJoined()
								if InArray2(grs, Group) {
									client.LeaveGroup(Group)
									fl, _ := client.GetAllContactIds()
									if InArray2(fl, user) {
										client.UnFriend(user)
									}
								}
							}
						}
					} else {
						Optime := op.CreatedTime
						if Room.ProInvite {
							if MemUser(Group, user) {
								go func(client *linetcr.Account, Group string, user string) {
									if filterWar.cek(user) {
										KickCancelProtect(client, user, invited[0], Group)
										Banned.AddBan(user)
										filterWar.del(user)
									}
								}(client, Group, user)
								if AutoPurge {
									go func(client *linetcr.Account, Group string, user string) {
										if filterWar.ceki(user) {
											Banned.AddBan(user)
											kickProtect(client, Group, user)
											filterWar.deli(user)
										}
									}(client, Group, user)
									if filterWar.ceko(Optime) {
										Room.ListInvited = invited
										BanAll(invited, Group)
										go cancelBanInv(client, invited, Group)
									}
								}
							} else {
								if filterWar.ceko(Optime) {
									go cancelallcek(client, invited, Group)
								}
							}
						} else {
							if MemBan(Group, user) {
								go func() {
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
									}
								}()
								if filterWar.ceko(Optime) {
									BanAll(invited, Group)
									go cancelBanInv(client, invited, Group)
									go func() {NodeBans(client, Group, invited)}()
								}
							} else {
								if MemUser(Group, user) {
									go func() {
										if filterWar.ceki(user) {
											for _, vo := range invited {
												if MemBan(Group, vo) {
													Banned.AddBan(user)
													kickPelaku(client, Group, user)
													break
												}
											}
											filterWar.deli(user)
										}
									}()
									if filterWar.ceko(Optime) {
										go cancelallcek(client, invited, Group)
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "invite", user)}
				}
				if op.Type == 133 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if MemUser(Group, user) {
							Banned.AddBan(user)
						}
					} else if !InArray2(Room.GoMid, client.MID) {
						if InArray2(client.Squads, Invited) {
							if MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									Banned.AddBan(user)
									go func() {
										if filterWar.cek(user) {
											groupBackupKick(client, Group, user, true)
											filterWar.del(user)
										}
									}()
									if filterWar.cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										filterWar.del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										if MemUser(Group, user) {
											Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProKick {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := Nkick.Get(user); !ok {
											Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {kickPelaku(client, Group, user)
											Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								Banned.AddBan(user)
							}
						}
						if MemUser(Group, user) {
							if filterWar.ceki(user) {
								GhostEnd(client, Group, Optime, user, true)
								filterWar.deli(user)
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "kick", user)}
				}
				if op.Type == 129 {
					runtime.GOMAXPROCS(cpu)
					Group, user := op.Param1, op.Param2
					if PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if filterWar.ceki(user) {
								AccKickBan(client, Group)
								filterWar.deli(user)
							}
						}()
					}}
				}
				if op.Type == 130 {
					runtime.GOMAXPROCS(cpu)
					Group, user := op.Param1, op.Param2
					Room := linetcr.GetRoom(Group)
					if linetcr.IoGOBot(Group, client) {
						if Room.ProJoin || AutoproN == true {
							if MemUser(Group, user) {
								if filterWar.ceki(user) {
										Banned.AddBan(user)
									kickPelaku(client, Group, user)
									filterWar.deli(user)
								}
							}
						} else {
							if MemBan(Group, user) {
								if MemUser(Group, user) {
									if filterWar.ceki(user) {
										Banned.AddBan(user)
										kickPelaku(client, Group, user)
										filterWar.deli(user)
									}
								}
							} else {
								if InArray2(Room.ListInvited, user) {
									if MemUser(Group, user) {
										if cekjoin(user) {
											kickPelaku(client, Group, user)
											deljoin(user)
											Room.ListInvited = Remove(Room.ListInvited, user)
										}
									} else {
										Room.ListInvited = Remove(Room.ListInvited, user)
									}
								} else {
									if Room.Welcome {
										if _, ok := cewel.Get(user); !ok {
											cewel.Set(user, 1)
											if cekjoin(user) {
												if !InArray2(Squadlist, user) {
													Room.WelsomeSet(client, Group, user)
												}
											}
										}
									} else {
										if LockMode == true {
											if MemUser(Group, user) {
												if filterWar.ceki(user) {
													Banned.AddBan2(user)
												}
											}
										}
									}
								}
							}
						}
					}
					Optime := op.CreatedTime
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				}
				if op.Type == 122 {
					runtime.GOMAXPROCS(cpu)
					Group, user, invited := op.Param1, op.Param2, op.Param3
					Optime := op.CreatedTime
					Room := linetcr.GetRoom(Group)
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if MemUser(Group, user) {
							if Room.ProQr || AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												err := cl.UpdateChatQrV2(Group, true)
												if err == nil {
													break
												}
											}
										}()
										if filterWar.ceki(user) {
											kickPelaku(client, Group, user)
											filterWar.deli(user)
											Banned.AddBan(user)
										}
									}
								}
							} else if KickBanQr == true {
								if invited == "4" {
									if cekOp2(Optime) {go func() {cans := linetcr.Actor(Group);for _, cl := range cans {err := cl.UpdateChatQrV2(Group, true);if err == nil {break}}}()
										if filterWar.ceki(user) {kickPelaku(client, Group, user);filterWar.deli(user);Banned.AddBan(user)}
									}
								}
							} else if Room.ProPicture || AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										if filterWar.ceki(user) {
											kickPelaku(client, Group, user)
											filterWar.deli(user)
										}
									}
								}
							} else if Room.ProName || AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												err := cl.UpdateChatName(Group, Room.Name)
												if err == nil {
													break
												}
											}
										}()
										if filterWar.ceki(user) {
											kickPelaku(client, Group, user)
											filterWar.deli(user)
										}
									}
								}
							} else {
								if MemBan(Group, user) {
									if invited == "4" {
										if cekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													err := cl.UpdateChatQrV2(Group, true)
													if err == nil {
														break
													}
												}
											}()
											if filterWar.ceki(user) {
												kickPelaku(client, Group, user)
												filterWar.deli(user)
												Banned.AddBan(user)
											}
										}
									} else if invited == "4" {
										if cekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													err := cl.UpdateChatName(Group, Room.Name)
													if err == nil {
														break
													}
												}
											}()
											if filterWar.ceki(user) {
												kickPelaku(client, Group, user)
												filterWar.deli(user)
											}
										}
									}
								}
							}

						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				}
				if op.Type == 126 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if MemUser(Group, user) {
							Banned.AddBan(user)
						}
					} else if !InArray2(Room.GoMid, client.MID) {
						if InArray2(client.Squads, Invited) {
							if MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									Banned.AddBan(user)
									go func() {
										if filterWar.cek(user) {
											groupBackupCans(client, Group, user, true)
											filterWar.del(user)
										}
									}()
									if filterWar.cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										filterWar.del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if filterWar.ceki(user) {
										kickPelaku(client, Group, user)
										filterWar.deli(user)
										if MemUser(Group, user) {
											Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProCancel {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := Nkick.Get(user); !ok {
											Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {kickPelaku(client, Group, user)
											Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								Banned.AddBan(user)
							}
						}
						if MemUser(Group, user) {
							if filterWar.ceki(user) {
								GhostEnd(client, Group, Optime, user, true)
								filterWar.deli(user)
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "cancel", user)}
				} else if op.Type == 33 {
					Group := op.Param1
					if LogMode && !InArray2(client.Squads, Group) {logAccess(client, client.Namebot, Group, "deleteaccount", []string{}, 2)}
				} else if op.Type == 5 {
					Group := op.Param1
					if LogMode && !InArray2(client.Squads, Group) {logAccess(client, client.Namebot, Group, "addfrind", []string{}, 2)}
				} else if op.Type == 50 {
					Group := op.Param1
					if LogMode && !InArray2(client.Squads, Group) {logAccess(client, client.Namebot, Group, "callme", []string{}, 2)}
				} else if op.Type == 55 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if cekOp(Optime) {
							if MemBan(Group, user) {
								kickPelaku(client, Group, user)
							} else {
								Room := linetcr.GetRoom(Group)
								if Room.Lurk && !InArray2(checkHaid, user) {
									Room.CheckLurk(client, Group, user)
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 26 {
					msg := op.Message
					Optime := op.CreatedTime
					if msg.ContentType != 18 {
						if _, ok := Command.Get(Optime); !ok {
							Command.Set(Optime, client)
							if _, ok := filterop.Get(Optime); !ok {
								filterop.Set(Optime, 1)
								Bot(op, client, ch)
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 128 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if !InArray2(Squadlist, user) {
						Room := linetcr.GetRoom(Group)
						if Room.Backleave {
							jangan := true
							tm, ok := botleave.Get(user)
							if ok {
								if time.Now().Sub(tm.(time.Time)) < 5*time.Second {
									jangan = false
								}
							}
							if jangan {
								if filterWar.ceki(user) {
									if !MemBan(Group, user) && !InArray2(Squadlist, user) && !UserBot.GetBot(user) && !InArray2(Room.GoMid, user) {
										hstg(Group, user)
										Room.Leave = time.Now()
									}
								}
							}
						} else {
							if Room.Leavebool {
								if _, ok := cleave.Get(user); !ok {
									cleave.Set(user, 1)
									if !MemBan(Group, user) && !InArray2(Squadlist, user) && !UserBot.GetBot(user) && !InArray2(Room.GoMid, user) {
										Room.LeaveSet(client, Group, user)
									}
								}
							}
						}
					}
					if _, ok := filtermsg.Get(Optime); !ok {
						filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 30 {
					Group := op.Param1
					Room := linetcr.GetRoom(Group)
					if Room.Announce && linetcr.IoGOBot(Group, client) {
						Optime := op.CreatedTime
						if cekOp(Optime) {
							Room.CheckAnnounce(client, Group)
						}
					}
				} else if op.Type == 123 {
					client.CInvite()
				} else if op.Type == 132 {
					runtime.GOMAXPROCS(cpu)
					Group, user := op.Param1, op.Param2
					if PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if filterWar.ceki(user) {
								KickBan132(client, Group, user)
								filterWar.deli(user)
							}
						}()
					}}
					client.CountKick()
				} else if op.Type == 125 {
					runtime.GOMAXPROCS(cpu)
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := linetcr.GetRoom(Group)
					if PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if filterWar.ceko(Optime) {
								BanAll(invited, Group)
								go CancelBan125(client, Group, invited)
							}
						}()
					}}
					if Room.ProInvite {
						if MemUser(Group, user) {
							go func(client *linetcr.Account, Group string, user string) {
								if filterWar.cek(user) {
									KickCancelProtect(client, user, invited[0], Group)
									Banned.AddBan(user)
									filterWar.del(user)
								}
							}(client, Group, user)
						}
					}
					client.CCancel()
				}
			}
		}(multiFunc)
		for _, ops := range multiFunc {
			client.SetSyncRevision(ops.Revision)
		}
	}
}

func getKey(cmd string) string {
	mp := linetcr.HashToMap(CmdHelper)
	for k, v := range mp {
		if v.(string) == cmd {
			return k
		}
	}
	return cmd
}

func SaveBackup() {
	fmt.Println("start Save Data *__*")
	Data.DetectcallBack = DetectCall
	Data.AutoproBack = AutoPro
	Data.AutoPurgeBack = AutoPurge
	Data.ProtectmodeBack = ProtectMode
	Data.PowermodeBack = PowerMode
	Data.KickbanqrBack = KickBanQr
	Data.MediadlBack = MediaDl
	Data.AutolikeBack = AutoLike
	Data.AutobcBack = AutoBc
	Data.NukejoinBack = NukeJoin
	Data.CanceljoinBack = canceljoin
	Data.AutojointicketBack = AutoJointicket
	Data.AutotranslateBack = AutoTranslate
	Data.ModebackupBack = ModeBackup
	Data.AutojoinBack = Autojoin
	Data.AjsjoinBack = Ajsjoin
	Data.TypejoinBack = TypeJoin
	Data.TypebcBack = Typebc
	Data.TypetransBack = TypeTrans
	Data.Maxkick = MaxKick
	Data.Maxcancel = MaxCancel
	Data.Maxinvite = MaxInvite
	Data.Cancelpend = CancelPend
	SaveProHistory()
	Data.GbanBack = map[string][]string{}
	Data.GownerBack = map[string][]string{}
	Data.GadminBack = map[string][]string{}
	Data.BanBack = []string{}
	Data.LockBack = []string{}
	Data.SnameBack = MsSname
	Data.RnameBack = MsRname
	Data.ResponBack = MsgRespon
	Data.KickSticker.Stkid = stkid
	Data.KickSticker.Stkpkgid = stkpkgid
	Data.ResponSticker.Stkid2 = stkid2
	Data.ResponSticker.Stkpkgid2 = stkpkgid2
	Data.StayallSticker.Stkid3 = stkid3
	Data.StayallSticker.Stkpkgid3 = stkpkgid3
	Data.LeaveSticker.Stkid4 = stkid4
	Data.LeaveSticker.Stkpkgid4 = stkpkgid4
	Data.KickallSticker.Stkid5 = stkid5
	Data.KickallSticker.Stkpkgid5 = stkpkgid5
	Data.BypassSticker.Stkid6 = stkid6
	Data.BypassSticker.Stkpkgid6 = stkpkgid6
	Data.InviteSticker.Stkid7 = stkid7
	Data.InviteSticker.Stkpkgid7 = stkpkgid7
	Data.ClearbanSticker.Stkid8 = stkid8
	Data.ClearbanSticker.Stkpkgid8 = stkpkgid8
	Data.CancelallSticker.Stkid9 = stkid9
	Data.CancelallSticker.Stkpkgid9 = stkpkgid9
	Data.BroadcastBack = MsgBroadcast
	Data.FuckBack = []string{}
	Data.MuteBack = []string{}
	Data.AnnunceBack = []string{}
	Data.ProQrBack = []string{}
	Data.ProNameBack = []string{}
	Data.ProPictureBack = []string{}
	Data.ProNoteBack = []string{}
	Data.ProAlbumBack = []string{}
	Data.ProjoinBack = []string{}
	Data.ProInviteBack = []string{}
	Data.ProCancelBack = []string{}
	Data.ProkickBack = []string{}
	Data.ProLinkBack = []string{}
	Data.ProFlexBack = []string{}
	Data.ProImageBack = []string{}
	Data.ProVideoBack = []string{}
	Data.ProCallBack = []string{}
	Data.ProSpamBack = []string{}
	Data.ProStickerBack = []string{}
	Data.ProContactBack = []string{}
	Data.ProPostBack = []string{}
	Data.ProFileBack = []string{}
	Data.CreatorBack = []string{}
	Data.MakerBack = []string{}
	Data.SellerBack = []string{}
	Data.BuyerBack = []string{}
	Data.OwnerBack = []string{}
	Data.MasterBack = []string{}
	Data.AdminBack = []string{}
	Data.BotBack = []string{}
	Data.TimeBanBack = map[string]time.Time{}
	if len(linetcr.KickBans) != 0 {
		for _, cl := range linetcr.KickBans {
			if _, ok := linetcr.GetBlock.Get(cl.MID); ok {
				Data.TimeBanBack[cl.MID] = cl.TimeBan
			}
		}
	}
	for _, room := range linetcr.SquadRoom {
		Data.GbanBack[room.Id] = []string{}
		Data.GownerBack[room.Id] = []string{}
		Data.GadminBack[room.Id] = []string{}
		if room.ProKick {
			Data.ProkickBack = append(Data.ProkickBack, room.Id)
		}
		if room.ProCancel {
			Data.ProCancelBack = append(Data.ProCancelBack, room.Id)
		}
		if room.ProInvite {
			Data.ProInviteBack = append(Data.ProInviteBack, room.Id)
		}
		if room.ProQr {
			Data.ProQrBack = append(Data.ProQrBack, room.Id)
		}
		if room.ProName {
			Data.ProNameBack = append(Data.ProNameBack, room.Id)
		}
		if room.ProPicture {
			Data.ProPictureBack = append(Data.ProPictureBack, room.Id)
		}
		if room.ProNote {
			Data.ProNoteBack = append(Data.ProNoteBack, room.Id)
		}
		if room.ProAlbum {
			Data.ProAlbumBack = append(Data.ProAlbumBack, room.Id)
		}
		if room.ProJoin {
			Data.ProjoinBack = append(Data.ProjoinBack, room.Id)
		}
		if room.ProLink {
			Data.ProLinkBack = append(Data.ProLinkBack, room.Id)
		}
		if room.ProFlex {
			Data.ProFlexBack = append(Data.ProFlexBack, room.Id)
		}
		if room.ProImage {
			Data.ProImageBack = append(Data.ProImageBack, room.Id)
		}
		if room.ProVideo {
			Data.ProVideoBack = append(Data.ProVideoBack, room.Id)
		}
		if room.ProCall {
			Data.ProCallBack = append(Data.ProCallBack, room.Id)
		}
		if room.ProSpam {
			Data.ProSpamBack = append(Data.ProSpamBack, room.Id)
		}
		if room.ProSticker {
			Data.ProStickerBack = append(Data.ProStickerBack, room.Id)
		}
		if room.ProContact {
			Data.ProContactBack = append(Data.ProContactBack, room.Id)
		}
		if room.ProPost {
			Data.ProPostBack = append(Data.ProPostBack, room.Id)
		}
		if room.ProFile {
			Data.ProFileBack = append(Data.ProFileBack, room.Id)
		}
		if room.Announce {
			Data.AnnunceBack = append(Data.AnnunceBack, room.Id)
		}
	}
	if len(UserBot.Creator) != 0 {
		for _, i := range UserBot.Creator {
			if !InArray2(Data.CreatorBack, i) {
				Data.CreatorBack = append(Data.CreatorBack, i)
			}
		}
	}
	if len(UserBot.Maker) != 0 {
		for _, i := range UserBot.Maker {
			if !InArray2(Data.MakerBack, i) {
				Data.MakerBack = append(Data.MakerBack, i)
			}
		}
	}
	if len(UserBot.Seller) != 0 {
		for _, i := range UserBot.Seller {
			if !InArray2(Data.SellerBack, i) {
				Data.SellerBack = append(Data.SellerBack, i)
			}
		}
	}
	if len(UserBot.Buyer) != 0 {
		for _, i := range UserBot.Buyer {
			if !InArray2(Data.BuyerBack, i) {
				Data.BuyerBack = append(Data.BuyerBack, i)
			}
		}
	}
	if len(UserBot.Owner) != 0 {
		for _, i := range UserBot.Owner {
			if !InArray2(Data.OwnerBack, i) {
				Data.OwnerBack = append(Data.OwnerBack, i)
			}
		}
	}
	if len(UserBot.Master) != 0 {
		for _, i := range UserBot.Master {
			if !InArray2(Data.MasterBack, i) {
				Data.MasterBack = append(Data.MasterBack, i)
			}
		}
	}
	if len(UserBot.Admin) != 0 {
		for _, i := range UserBot.Admin {
			if !InArray2(Data.AdminBack, i) {
				Data.AdminBack = append(Data.AdminBack, i)
			}
		}
	}
	if len(UserBot.Bot) != 0 {
		for _, i := range UserBot.Bot {
			if !InArray2(Data.BotBack, i) {
				Data.BotBack = append(Data.BotBack, i)
			}
		}
	}
	if len(Data.GbanBack) != 0 {
		for to := range Data.GbanBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gban) != 0 {
				for _, i := range Room.Gban {
					if MemUser(to, i) {
						if !InArray2(Data.GbanBack[to], i) {
							Data.GbanBack[to] = append(Data.GbanBack[to], i)
						}
					}
				}
			}
		}
	}
	if len(Data.GownerBack) != 0 {
		for to := range Data.GownerBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gowner) != 0 {
				for _, i := range Room.Gowner {
					if !InArray2(Data.GownerBack[to], i) {
						Data.GownerBack[to] = append(Data.GownerBack[to], i)
					}
				}
			}
		}
	}
	if len(Data.GadminBack) != 0 {
		for to := range Data.GadminBack {
			Room := linetcr.GetRoom(to)
			if len(Room.Gadmin) != 0 {
				for _, i := range Room.Gadmin {
					if !InArray2(Data.GadminBack[to], i) {
						Data.GadminBack[to] = append(Data.GadminBack[to], i)
					}
				}
			}
		}
	}
	if len(Banned.Banlist) != 0 {
		for _, i := range Banned.Banlist {
			if MemAccsess(i) {
				if !InArray2(Data.BanBack, i) {
					Data.BanBack = append(Data.BanBack, i)
				}
			}
		}
	}
	if len(Banned.Fucklist) != 0 {
		for _, i := range Banned.Fucklist {
			if MemAccsess(i) {
				if !InArray2(Data.FuckBack, i) {
					Data.FuckBack = append(Data.FuckBack, i)
				}
			}
		}
	}
	if len(Banned.Locklist) != 0 {
		for _, i := range Banned.Locklist {
			if MemAccsess(i) {
				if !InArray2(Data.LockBack, i) {
					Data.LockBack = append(Data.LockBack, i)
				}
			}
		}
	}
	if len(Banned.Mutelist) != 0 {
		for _, i := range Banned.Mutelist {
			if MemAccsess(i) {
				if !InArray2(Data.MuteBack, i) {
					Data.MuteBack = append(Data.MuteBack, i)
				}
			}
		}
	}
	if len(Data.WordbanBack) != 0 {
		for _, msg := range Data.WordbanBack {
			if !InArray2(Data.WordbanBack, msg) {
				Data.WordbanBack = append(Data.WordbanBack, msg)
			}
		}
	}
	fmt.Println("done save Data *__*")
	SaveData()
}
func Checkkickuser(group string, user string, invited string) bool {
	Room := linetcr.GetRoom(group)
	if InArray2(DEVELOPER, invited) {
		if !InArray2(DEVELOPER, user) {
			return true
		}
	} else if UserBot.GetCreator(invited) {
		if !SendMycreator(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetMaker(invited) {
		if !SendMymaker(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetSeller(invited) {
		if !SendMyseller(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetBuyer(invited) {
		if !SendMybuyer(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetOwner(invited) {
		if !SendMyowner(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetMaster(invited) {
		if !SendMymaster(user) && !Allbotlist(user) {
			return true
		}
	} else if UserBot.GetAdmin(invited) {
		if !SendMyadmin(user) && !Allbotlist(user) {
			return true
		}
	} else if InArray2(Room.Gowner, invited) {
		if !SendMygowner(group, user) && !Allbotlist(user) {
			return true
		}
	} else if InArray2(Room.Gadmin, invited) {
		if MemUser(group, user) {
			return true
		}
	} else if UserBot.GetBot(invited) {
		if MemUser(group, user) {
			return true
		}
	}
	return false
}
func MemUser(group string, from string) bool {
	Room := linetcr.GetRoom(group)
	if InArray2(Squadlist, from) {
		return false
	} else if UserBot.GetBot(from) {
		return false
	} else if InArray2(DEVELOPER, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetMaker(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	} else if InArray2(Room.Gowner, from) {
		return false
	} else if InArray2(Room.Gadmin, from) {
		return false
	} else if InArray2(TeamNotif, from) {
		return false
	}
	return true
}
func MemAccsess(from string) bool {
	if InArray2(Squadlist, from) {
		return false
	} else if UserBot.GetBot(from) {
		return false
	} else if InArray2(DEVELOPER, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetMaker(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	}
	return true
}

func MemUserN(group string, from string) bool {
	Room := linetcr.GetRoom(group)
	if UserBot.GetBot(from) {
		return false
	} else if InArray2(DEVELOPER, from) {
		return false
	} else if UserBot.GetCreator(from) {
		return false
	} else if UserBot.GetMaker(from) {
		return false
	} else if UserBot.GetSeller(from) {
		return false
	} else if UserBot.GetBuyer(from) {
		return false
	} else if UserBot.GetOwner(from) {
		return false
	} else if UserBot.GetMaster(from) {
		return false
	} else if UserBot.GetAdmin(from) {
		return false
	} else if InArray2(Room.Gowner, from) {
		return false
	} else if InArray2(Room.Gadmin, from) {
		return false
	}
	return true
}

func Allbotlist(user string) bool {
	if InArray2(Squadlist, user) {
		return true
	} else if UserBot.GetBot(user) {
		return true
	}
	return false
}

func Upsetcmd(text string, text2 string) string {
	count := 0
	if text == "rollcall" {
		Commands.Botname = text2
		count = count + 1
	} else if text == "upallimage" {
		Commands.Upallimage = text2
		count = count + 1
	} else if text == "upallcover" {
		Commands.Upallcover = text2
		count = count + 1
	} else if text == "unsend" {
		Commands.Unsend = text2
		count = count + 1
	} else if text == "upvallimage" {
		Commands.Upvallimage = text2
		count = count + 1
	} else if text == "upvallcover" {
		Commands.Upvallcover = text2
		count = count + 1
	} else if text == "appname" {
		Commands.Appname = text2
		count = count + 1
	} else if text == "useragent" {
		Commands.Useragent = text2
		count = count + 1
	} else if text == "hostname" {
		Commands.Hostname = text2
		count = count + 1
	} else if text == "friends" {
		Commands.Friends = text2
		count = count + 1
	} else if text == "adds" {
		Commands.Adds = text2
		count = count + 1
	} else if text == "limits" {
		Commands.Limits = text2
		count = count + 1
	} else if text == "addallbots" {
		Commands.Addallbots = text2
		count = count + 1
	} else if text == "addallsquads" {
		Commands.Addallsquads = text2
		count = count + 1
	} else if text == "leave" {
		Commands.Leave = text2
		count = count + 1
	} else if text == "respon" {
		Commands.Respon = text2
		count = count + 1
	} else if text == "ping" {
		Commands.Ping = text2
		count = count + 1
	} else if text == "count" {
		Commands.Count = text2
		count = count + 1
	} else if text == "limitout" {
		Commands.Limitout = text2
		count = count + 1
	} else if text == "access" {
		Commands.Access = text2
		count = count + 1
	} else if text == "allbanlist" {
		Commands.Allbanlist = text2
		count = count + 1
	} else if text == "allgaccess" {
		Commands.Allgaccess = text2
		count = count + 1
	} else if text == "gaccess" {
		Commands.Gaccess = text2
		count = count + 1
	} else if text == "checkram" {
		Commands.Checkram = text2
		count = count + 1
	} else if text == "upimage" {
		Commands.Upimage = text2
		count = count + 1
	} else if text == "upcover" {
		Commands.Upcover = text2
		count = count + 1
	} else if text == "upvimage" {
		Commands.Upvimage = text2
		count = count + 1
	} else if text == "upvcover" {
		Commands.Upvcover = text2
		count = count + 1
	} else if text == "Purgeall" {
		Commands.Purgeall = text2
		count = count + 1
	} else if text == "banlist" {
		Commands.Banlist = text2
		count = count + 1
	} else if text == "clearban" {
		Commands.Clearban = text2
		count = count + 1
	} else if text == "bringall" {
		Commands.Bringall = text2
		count = count + 1
	} else if text == "stayall" {
		Commands.Stayall = text2
		count = count + 1
	} else if text == "clears" {
		Commands.Clearchat = text2
		count = count + 1
	} else if text == "here" {
		Commands.Here = text2
		count = count + 1
	} else if text == "speed" {
		Commands.Speed = text2
		count = count + 1
	} else if text == "status" {
		Commands.Status = text2
		count = count + 1
	} else if text == "tagall" {
		Commands.Tagall = text2
		count = count + 1
	} else if text == "kick" {
		Commands.Kick = text2
		count = count + 1
	} else if text == "max" {
		Commands.Max = text2
		count = count + 1
	} else if text == "none" {
		Commands.None = text2
		count = count + 1
	} else if text == "kickall" {
		Commands.Kickall = text2
		count = count + 1
	} else if text == "cancelall" {
		Commands.Cancelall = text2
		count = count + 1
	}
	if count != 0 {
		kowe := text
		jancuk := text2
		newsend := "Changed cmd: " + kowe + " to " + jancuk + "\n"
		return newsend
	}
	return ""
}

////NEW
func addwl(g string, w []string) {
	for _, mid := range w {
		if !MemBan(g, mid) {
			if !InArray2(Whitelist, mid) && MemUser(g, mid) {
				Whitelist = append(Whitelist, mid)
			}
		}
	}
}
func CancelEnemy(client *linetcr.Account, mem []string, to string) {
	defer panicHandle("CancelEnemy")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, []string{target})
		}(target, no)
		if ah >= MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}
func KickCancelProtect(client *linetcr.Account, ktrg string, ctrg string, to string) {
    defer panicHandle("KickCancelProtect")
    runtime.GOMAXPROCS(cpu)
    go func() {
        client.DeleteOtherFromChats(to, []string{ktrg})
        client.CancelChatInvitations(to, []string{ctrg})
    }()
    time.Sleep(25 * time.Millisecond)
}
func kickPelaku(client *linetcr.Account, to, pelaku string) {
	defer panicHandle("kickPelaku")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if InArray2(Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {go func(pelaku string, no int) {exe[no].DeleteOtherFromChats(to, []string{pelaku})}(pelaku, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	linetcr.SetAva(to, oke)
}
func kickProtect(client *linetcr.Account, to, pelaku string) {
	defer panicHandle("kickProtect")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if InArray2(Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if _, ok := memlist[pelaku]; ok {exe[0].DeleteOtherFromChats(to, []string{pelaku})}
	}
	linetcr.SetAva(to, oke)
}
func GhostEnd(client *linetcr.Account, Group string, Optime int64, pelaku string, cek bool) {
	Room := linetcr.GetRoom(Group)
	_, memlist := client.GetGroupMember(Group)
	oke := []string{}
	ban := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			oke = append(oke, mid)
		} else if MemBan(Group, mid) {
			ban = append(ban, mid)
		}
	}
	if len(oke) == 0 {
		if !InArrayInt64(cekGo, Optime) {
			cekGo = append(cekGo, Optime)
			cls := []*linetcr.Account{}
			Bot2 := Room.Bot
			bots := Room.HaveClient
			for n, cl := range Room.GoClient {
				if n < 2 {
					go cl.AcceptGroupInvitationNormal(Group)
					cls = append(cls, cl)
				}
			}
			cc := len(cls)
			if cc != 0 {
				if Ajsjoin == "qr" {
					qrGo22(cls[0], bots, Group)
				} else if Ajsjoin == "inv" {
					cls[0].InviteIntoChatPollVer(Group, Bot2)
				}
				if ProtectMode {
					chat := cls[0].GetChat([]string{Group}, true, true)
					if chat == nil { return }
					memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
					var createdTime int64
					for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
					}
					for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(Group, mid) {Banned.AddBan(mid);ban = append(ban, mid)}}
					}
					for _, target := range ban {go func(target string) {cls[0].DeleteOtherFromChats(Group, []string{target})}(target)
					}
				}
				if len(ban) != 0 {
					for _, target := range ban {
						go func(target string) {cls[0].DeleteOtherFromChats(Group, []string{target})}(target)
					}
				}
				for _, cl := range cls {time.Sleep(1 * time.Second);cl.LeaveGroup(Group);linetcr.GetRoom(Group).RevertGo(cl)
				}
				time.Sleep(1 * time.Second)
				clbot := Room.Client
				for _, cbot := range clbot {if !cbot.Limited {for i := range Room.GoMid {cbot.InviteIntoGroupNormal(Group, []string{Room.GoMid[i]})};break}
				}
			}
		}
	}
}

func cancelBanInv(client *linetcr.Account, mem []string, to string) {
	defer panicHandle("cancelBanInv")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	Cans := Room.HaveClient
	clen := len(Cans)
	if clen != 0 {no := 0;ah := 0;if len(mem) > 50 {mem = mem[:50]};for _, target := range mem {var wg sync.WaitGroup;wg.Add(len(mem));go func(target string) {defer wg.Done();Cans[no].CancelChatInvitations(to, []string{target})}(target);if ah >= MaxCancel {no++;if no >= clen {no = 0};ah = 0};ah++}
	}
}

func cancelall(client *linetcr.Account, mem []string, to string) {
	defer panicHandle("cancelall")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, []string{target})
		}(target, no)
		if ah >= MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}

func getfuckV1(client *linetcr.Account, vo string, to string) {
	defer panicHandle("getfuck")
	runtime.GOMAXPROCS(cpu)
	for _, v := range Banned.Banlist {
		if Banned.GetBan(v) {
			if linetcr.IsPending(client, to, v) == true {
				go func(v string) { 
					client.CancelChatInvitations(to, []string{v})
				}(v)
			}
			if linetcr.IsMembers(client, to, v) == true {
				if Banned.GetBan(v) {
					go func(v string) { 
						client.DeleteOtherFromChats(to, []string{v}) 
					}(v)
				}
			}
		}
	}
}

func getfuck(cl *linetcr.Account, vo string, Group string) {
	defer panicHandle("getfuck")
	runtime.GOMAXPROCS(cpu)
	if MemBan(Group, vo) {
		cl.CancelChatInvitations(Group, []string{vo})
	}
}

func MemEx(to, user string) bool {
	defer linetcr.PanicOnly()
	if Banned.GetEx(user) {
		return true
	}
	return false
}

func cancelallcek(client *linetcr.Account, mem []string, to string) {
	defer panicHandle("cancelallcek")
	runtime.GOMAXPROCS(cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	Cans := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				Cans = append(Cans, cl)
			}
			oke = append(oke, mid)
		}
	}
	if len(Cans) != 0 {
		sort.Slice(Cans, func(i, j int) bool {
			return Cans[i].KickPoint < Cans[j].KickPoint
		})
		Room.HaveClient = Cans
		no := 0
		ah := 0
		for _, target := range mem {
			go getfuck(Cans[no], target, to)
			if ah >= MaxCancel {
				no++
				if no >= len(Cans) {
					no = 0
				}
				ah = 0
			}
			ah++
		}
	}
}

func Purgeact(Group string, cl *linetcr.Account) {
	defer panicHandle("purgeact")
	mem := make(chan []string)
	go func(m chan []string) {
		memlistss := []string{}
		_, memlists := cl.GetGroupMember(Group)
		for target, _ := range memlists {
			if MemBan(Group, target) {
				memlistss = append(memlistss, target)
			}
		}
		m <- memlistss
	}(mem)
	Cans := []*linetcr.Account{}
	for _, ym := range linetcr.Actor(Group) {
		if ym.Limited {
			Cans = append(Cans, ym)
		}
	}
	ClAct := len(Cans)
	if ClAct != 0 {
		no := 0
		memlist := <-mem
		for _, target := range memlist {
			if no >= ClAct {
				no = 0
			}
			cl = Cans[no]
			go cl.DeleteOtherFromChat(Group, target)
			no += 1
		}
	} else if cl.Limited {
		memlist := <-mem
		for _, target := range memlist {
			go cl.DeleteOtherFromChat(Group, target)
		}
	}
}
func groupBackupKick(client *linetcr.Account, to, pelaku string, cek bool) {
	defer panicHandle("groupBackupKick")
	Room := linetcr.GetRoom(to)
	memlist, pending := client.GetChatListMap(to)
	ban := []string{}
	ban2 := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if InArray2(Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	for mid2, _ := range pending {if MemBan(to, mid2) {ban2 = append(ban2, mid2)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban2) != 0 {
			no := 0
			ah := 0
			for _, target2 := range ban2 {go func(target2 string, no int) {exe[no].CancelChatInvitations(to, []string{target2})}(target2, no);if ah >= MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	//} else {
		//no := 0
		//ah := 0
		//if _, ok := memlist[pelaku]; ok {
			//exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		//}
	}
	runtime.GOMAXPROCS(cpu)
}
func groupBackupKickV2(client *linetcr.Account, to, pelaku string, cek bool) {
	defer panicHandle("groupBackupKickV2")
	Room := linetcr.GetRoom(to)
	memlist, pending := client.GetChatListMap(to)
	ban := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if InArray2(Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	for mid2, _ := range pending {if MemBan(to, mid2) {ban = append(ban, mid2)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
			for _, target2 := range ban {go func(target2 string, no int) {exe[no].CancelChatInvitations(to, []string{target2})}(target2, no);if ah >= MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	} else {
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {
			exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	runtime.GOMAXPROCS(cpu)
}
func groupBackupCans(client *linetcr.Account, to, pelaku string, cek bool) {
	defer panicHandle("groupBackupCans")
	Room := linetcr.GetRoom(to)
	memlist, _ := client.GetChatListMap(to)
	ban := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if InArray2(Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	} else {
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {
			exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	runtime.GOMAXPROCS(cpu)
}

func cekKick(optime int64) bool {
	for _, tar := range opkick {
		if tar == optime {
			return false
		}
	}
	opkick = append(opkick, optime)
	return true
}

func deljoin(user string) {
	for _, us := range opjoin {
		if us == user {
			opjoin = Remove(opjoin, user)
		}
	}
}
func Setinviteto(client *linetcr.Account, to string, invits []string) {
	news := []string{}
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == false && linetcr.IsPending(client, to, cc) == false {
			news = append(news, cc)
		}
	}
	if len(news) != 0 {
		client.InviteIntoChatPollVer(to, news)
	}
}
func invBackup(exe *linetcr.Account, to string, oke []string, korban string) {
	exe.InviteIntoGroupNormal(to, []string{korban})
}
func openqr(exe []*linetcr.Account, to string, mes chan bool) {
	defer panicHandle("QR_backupupdate")
	Room := linetcr.GetRoom(to)
	Room.Qr = false
	for _, cl := range exe {
		err = cl.UpdateChatQrV2(to, false)
		if err == nil {
			mes <- true
			return
		}
	}
	mes <- false
}

func getTicket(exe []*linetcr.Account, to string, lnk chan string) {
	defer panicHandle("gettiket")
	ClAct := len(exe)
	if ClAct > 1 {
		for i := ClAct - 1; i >= 0; i-- {
			cls := exe[i]
			link, err := cls.ReissueChatTicket(to)
			if err == nil {
				lnk <- link
				return
			}
		}
		lnk <- "error"
		return

	} else {
		link, err := exe[0].ReissueChatTicket(to)
		if err == nil {
			lnk <- link
		} else {
			lnk <- "error"
		}
		return
	}
}

func qrBackup(exe []*linetcr.Account, to string, oke []string) {
	defer panicHandle("qrBackup")
	lnk := make(chan string)
	Room := linetcr.GetRoom(to)
	mes := make(chan bool)
	go openqr(exe, to, mes)
	go getTicket(exe, to, lnk)
	bot := linetcr.GetRoom(to).Bot
	cans := []*linetcr.Account{}
	for _, mid := range bot {
		if !InArray2(oke, mid) {
			cl := GetKorban(mid)
			if cl.Limited {
				cans = append(cans, cl)
			}
		}
	}
	var wg sync.WaitGroup
	linetcr.GetRoom(to).Purge = false
	ticket := <-lnk
	if len(ticket) > 5 {
		ok := <-mes
		if ok {
			for _, cc := range cans {
				wg.Add(1)
				go func(c *linetcr.Account) {
					c.AcceptTicket(to, ticket)
					wg.Done()
				}(cc)
			}
			wg.Wait()
			Room.Qr = true
		} else {
			Room.Qr = true
		}

	} else {
		Room.Qr = true
	}
}
func groupBackupInv2(client *linetcr.Account, to string) {
	defer panicHandle("groupBackupInv2")
	all := []string{}
	GetSquad(client, to)
	room := linetcr.GetRoom(to)
	cuk := room.Client
	exe := []*linetcr.Account{}
	for _, x := range ClientBot {
		if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
			exe = append(exe, x)
		}
	}
	for _, x := range ClientBot {
		if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
			all = append(all, x.MID)
		} else {
			break
		}
	}
	ClAct := len(exe)
	if ClAct != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		for _, a := range all {
			exe[0].InviteIntoGroupNormal(to, []string{a})
	            	time.Sleep(250 * time.Millisecond)
			GetSquad(client, to)
		}
	}
	runtime.GOMAXPROCS(cpu)
}
func groupBackupInv(client *linetcr.Account, to string, optime int64, korban string) {
	defer panicHandle("groupBackupInv")
	memlist, _ := client.GetChatListMap(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
			oke = append(oke, mid)
		}
	}
	ClAct := len(exe)
	if ClAct != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		if ModeBackup == "inv" {
			invBackup(exe[0], to, oke, korban)
		} else if ModeBackup == "qr" {
			qrBackup(exe, to, oke)
		}
		linetcr.SetAva(to, oke)
	}
	runtime.GOMAXPROCS(cpu)
}
func getBot(client *linetcr.Account, to string) []*linetcr.Account {
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if InArray2(Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
		}
	}
	sort.Slice(exe, func(i, j int) bool {
		return exe[i].KickPoint < exe[j].KickPoint
	})
	linetcr.GetRoom(to).HaveClient = exe
	return exe
}

func AddContact2(cl *linetcr.Account, con string) int {
	fl, _ := cl.GetAllContactIds()
	if !InArray2(fl, con) {
		if con != cl.MID && !cl.Limitadd {
			_, err := cl.FindAndAddContactsByMidV2(con)
			if err != nil {
				println(fmt.Sprintf("%v", err.Error()))
				return 0
			}
			return 1
		} else {
			return 0
		}
	}
	return 1
}

func kickDirt(client *linetcr.Account, to, pelaku string) {
	runtime.GOMAXPROCS(cpu)
	cans := linetcr.Actor(to)
	for _, cl := range cans {
		if linetcr.GetRoom(to).Act(cl) {
			in := cl.DeleteOtherFromChat(to, pelaku)
			if in == 35 || in == 10 {
				continue
			} else {
				break
			}
		}
	}
}

func CekPurge(optime int64) bool {
	defer linetcr.PanicOnly()
	for _, tar := range PurgeOP {
		if tar == optime {
			return false
		}
	}
	PurgeOP = append(PurgeOP, optime)
	return true
}

func cekjoin(optime string) bool {
	defer linetcr.PanicOnly()
	for _, tar := range opjoin {
		if tar == optime {
			return false
		}
	}
	opjoin = append(opjoin, optime)
	return true
}

func cekOp2(optime int64) bool {
	for _, tar := range cekoptime {
		if tar == optime {
			return false
		}
	}
	cekoptime = append(cekoptime, optime)
	return true
}

// here func notif SELFTCR™

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
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "kick" {
	       list := fmt.Sprintf("NOTIFI KICK:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "cancel" {
	       list := fmt.Sprintf("NOTIFI CANCEL:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "delete" {
	       list := fmt.Sprintf("NOTIFI DELETE ACCOUNT:\n  . Mid : %s", tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "call" {
	       list := fmt.Sprintf("NOTIFI CALL ME:\n  . Mid : %s", tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "join" {
	       list := fmt.Sprintf("NOTIFI MEMBER JOIN:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	} else if tipe == "leave" {
	       list := fmt.Sprintf("NOTIFI MEMBER LEAVE:\n  . Group: %s\n  . Mid : %s", names1, tcr.Mid)
	       list += "\n  . Name: "+tcr.DisplayName+" \n  . Date: "+Date+"\n  . Time: "+Time
	      if len(ClientBot) != 0 {
		      ClientBot[0].SendNewText(LogGroup, list)
	     }
	}
}

func LogFight(room *linetcr.LineRoom) {
	defer panicHandle("logfight")
	if LogMode {
		var tx = ""
		for i := 0; i < len(ClientBot); i++ {
			exe := ClientBot[i]
			if !exe.Frez {
				g, err := exe.GetGroupMember(room.Id)
				if err != nil {
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
		room := linetcr.GetRoom(LogGroup)
		if len(room.Client) != 0 {
			exe, err := SelectBot(room.Client[0], LogGroup)
			if err == nil {
				if exe != nil {
					exe.SendMessage(LogGroup, fancy(tx))
				}
			} else {
				LogMode = false
				LogGroup = ""
			}
		}
	}
	room.Kick = 0
	room.Invite = 0
	room.Cancel = 0
}

func Bot(op *SyncService.Operation, client *linetcr.Account, ch chan int) {
	defer panicHandle("Bot")
	op.Message = client.DecryptE2EEMessage(op.Message)
	msg := op.Message
	if msg.ToType != 2 {
		return
	}
	if _, ok := Commandss.Get(op.CreatedTime); ok {
		return
	} else {
		Commandss.Set(op.CreatedTime, client)
	}
	if time.Now().Sub(timeabort) >= 60*time.Second {
		abort()
	}
       if AutoBackBot {
		time.AfterFunc(time.Duration(Timebk)*time.Second, func() {  
			//groupBackupInv2(client, op.Message.To)
			all := []string{}
			GetSquad(client, op.Message.To)
			room := linetcr.GetRoom(op.Message.To)
			cuk := room.Client
			alls := []*linetcr.Account{}
			for _, x := range ClientBot {
				if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
					alls = append(alls, x)
				}
			}
			for _, x := range ClientBot {
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
	Rname := MsRname
	Sname := MsSname
	sender := op.Message.From_
	text := op.Message.Text
	receiver := op.Message.To
	var pesan = strings.ToLower(text)
	var to string
	mentions := mentions{}
	if op.Message.ToType == 0 {
		to = sender
	} else {
		to = receiver
	}
       if DetectCall && msg.ToType == 2 {
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
                   a += "\n    • Date: "+Date
                   a += "\n    • Time: "+Time
                   a += fmt.Sprintf("\n    • Group: %s", Room.Name)
		      a += "\n    • Host: "+pr.DisplayName
                   client.SendMessage(to, fancy(a))
		} else {
                   if msg.ContentMetadata["GC_MEDIA_TYPE"] == "VIDEO" || msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
                        a := "「 DETECT 」"
                        a += "\n › Type: Callgroup"
                        a += "\n    • Type: VIDEO"
                        a += "\n    • Date: "+Date
                        a += "\n    • Time: "+Time
                        a += fmt.Sprintf("\n    • Group: %s", Room.Name)
		           a += "\n    • Host: "+pr.DisplayName
                        client.SendMessage(to, fancy(a))
                  }
             }
       }
	if len(Sinderremote) != 0 {
		if InArray2(Sinderremote, sender) {
			if remotegrupid != "" {
				remotegrupidto = to
				to = remotegrupid
			}
		}
	}
	mentionlist := []string{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["MENTION"]), &mentions)
	for _, mention := range mentions.MENTIONEES {
		if !InArray2(mentionlist, mention.Mid) {
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
	bk := []*linetcr.Account{}
	bk2 := []*linetcr.Account{}
	for _, n := range bks {
		bk = append(bk, n)
		if !n.Limited {
			bk2 = append(bk2, n)
		}
	}
	clen := len(bk2)
	if clen != 0 {
		client = bk2[0]
		room.Exe = bk2[0]
		room.Limit = false
	} else {
		room.Limit = true
	}
	if room.AntiTag && MemUser(to, msg.From_) && len(mentionlist) != 0 && !room.Automute {
		if room.Limit {
			client.SendMessage(to, fancy("All bot in here banned, please try invite another bot"))
			return
		}
		if client.Limited == false {
			client.DeleteOtherFromChats(to, []string{msg.From_})
		} else {
			for _, bot := range bk {
				if bot.Limited == false {
					bot.DeleteOtherFromChats(to, []string{msg.From_})
					break
				}
			}
		}
	}
	if op.Message.RelatedMessageId != "" && len(mentionlist) == 0 {
		asu, _ := client.GetRecentMessagesV2(op.Message.To)
		for _, xx := range asu {
			if xx.ID == op.Message.RelatedMessageId {
				Rplay = xx.From_
				break
			}
		}
	}
	emots := emots{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["REPLACE"]), &emots)
	for _, stiker := range emots.STICON.RESOURCES {
		if GetCodeprem(5, sender, to) {
		if !MemUser(to, msg.From_) {
			if getStickerKick == 1 {
				if !MemUser(to, msg.From_) {
					stkid = stiker.PRODUCTID
					stkpkgid = stiker.STICONID
					SaveBackup()
					getStickerKick = 0
					client.SendMessage(to, fancy("kick by emote updated"))
				}
			} else if getStickerRespon == 1 {
				if !MemUser(to, msg.From_) {
					stkid2 = stiker.PRODUCTID
					stkpkgid2 = stiker.STICONID
					SaveBackup()
					getStickerRespon = 0
					client.SendMessage(to, fancy("respon by emote updated"))
				}
			} else if getStickerStayall == 1 {
				if !MemUser(to, msg.From_) {
					stkid3 = stiker.PRODUCTID
					stkpkgid3 = stiker.STICONID
					SaveBackup()
					getStickerStayall = 0
					client.SendMessage(to, fancy("stayall by emote updated"))
				}
			} else if getStickerLeave == 1 {
				if !MemUser(to, msg.From_) {
					stkid4 = stiker.PRODUCTID
					stkpkgid4 = stiker.STICONID
					SaveBackup()
					getStickerLeave = 0
					client.SendMessage(to, fancy("leave by emote updated"))
				}
			} else if getStickerKickall == 1 {
				if !MemUser(to, msg.From_) {
					stkid5 = stiker.PRODUCTID
					stkpkgid5 = stiker.STICONID
					SaveBackup()
					getStickerKickall = 0
					client.SendMessage(to, fancy("kickall by emote updated"))
				}
			} else if getStickerBypass == 1 {
				if !MemUser(to, msg.From_) {
					stkid6 = stiker.PRODUCTID
					stkpkgid6 = stiker.STICONID
					SaveBackup()
					getStickerBypass = 0
					client.SendMessage(to, fancy("bypass by emote updated"))
				}
			} else if getStickerInvite == 1 {
				if !MemUser(to, msg.From_) {
					stkid7 = stiker.PRODUCTID
					stkpkgid7 = stiker.STICONID
					SaveBackup()
					getStickerInvite = 0
					client.SendMessage(to, fancy("invite by emote updated"))
				}
			} else if getStickerClearban == 1 {
				if !MemUser(to, msg.From_) {
					stkid8 = stiker.PRODUCTID
					stkpkgid8 = stiker.STICONID
					SaveBackup()
					getStickerClearban = 0
					client.SendMessage(to, fancy("clearban by emote updated"))
				}
			} else if getStickerCancelall == 1 {
				if !MemUser(to, msg.From_) {
					stkid9 = stiker.PRODUCTID
					stkpkgid9 = stiker.STICONID
					SaveBackup()
					getStickerCancelall = 0
					client.SendMessage(to, fancy("cancelall by emote updated"))
				}
			} else if stiker.PRODUCTID == stkid && stiker.STICONID == stkpkgid {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
							if AutoBan {
										Banned.AddBan(lists[i])
							}
							listuser = append(listuser, lists[i])
						}
					}
				}
				fmt.Println(listuser)
				if len(listuser) != 0 {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, listuser)
						logAccess(client, to, sender, "kick", listuser, msg.ToType)
					}
				}
			} else if stiker.PRODUCTID == stkid2 && stiker.STICONID == stkpkgid2 {
				for _, p := range bk {
					go p.SendMessage(to, fancy(MsgRespon))
				}
			} else if stiker.PRODUCTID == stkid3 && stiker.STICONID == stkpkgid3 {
				if room.Limit {
					client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
					return
				}
				numb := len(ClientBot)
				if numb > 0 && numb <= len(ClientBot) {
					GetSquad(client, to)
					room := linetcr.GetRoom(to)
					aa := len(room.Client)
					if aa > numb {
						c := aa - numb
						ca := 0
						list := append([]*linetcr.Account{}, room.Client...)
						sort.Slice(list, func(i, j int) bool {
							return list[i].KickPoint > list[j].KickPoint
						})
						for _, o := range list {
							o.LeaveGroup(to)
							ca = ca + 1
							if ca == c {
								break
							}
						}
						GetSquad(client, to)
					} else if aa < numb {
						ti, err := client.ReissueChatTicket(to)
						if err == nil {
							go client.UpdateChatQrV2(to, false)
							all := []*linetcr.Account{}
							room := linetcr.GetRoom(to)
							cuk := room.Client
							for _, x := range ClientBot {
								if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
									all = append(all, x)
								}
							}
							sort.Slice(all, func(i, j int) bool {
								return all[i].KickPoint < all[j].KickPoint
							})
							g := numb - aa
							var wg sync.WaitGroup
							wi := GetSquad(client, to)
							for i := 0; i < len(all); i++ {
								if i == g {
									break
								}
								l := all[i]
								if l != client && !linetcr.InArrayCl(wi, l) {
									if !l.Limited {
										wg.Add(1)
										go func() {
											l.AcceptTicket(to, ti)
											wg.Done()
										}()
									}
								}
							}
							wg.Wait()
							client.UpdateChatQrV2(to, true)
							GetSquad(client, to)
							logAccess(client, to, sender, "bringbot", []string{}, 2)
							SaveBackup()
							aa := len(room.Client)
							var name string
							name = fmt.Sprintf("Ready %v bots here", aa)
							client.SendMessage(to, fancy(name))
							//newsend += name + "\n"
						}
					}
				}
			} else if stiker.PRODUCTID == stkid4 && stiker.STICONID == stkpkgid4 {
				_, mem := client.GetGroupInvitation(to)
				anu := []string{}
				for m := range mem {
					if InArray2(Squadlist, m) {
						anu = append(anu, m)
					}
				}
				if len(anu) != 0 {
					for _, mid := range anu {
						cl := GetKorban(mid)
						cl.AcceptGroupInvitationNormal(to)
					}
				}
				GetSquad(client, to)
				room := linetcr.GetRoom(to)
				bk = room.Client
				for _, cl := range bk {
					go cl.LeaveGroup(to)
				}
				if LogGroup == to {
					LogMode = false
					LogGroup = ""
				}
				linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				SaveBackup()
				logAccess(client, to, sender, "leave", []string{}, msg.ToType)
			} else if stiker.PRODUCTID == stkid5 && stiker.STICONID == stkpkgid5 {
				_, memlist, _ := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist {
					if InArray2(Squadlist, mid) {
						cl := GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				max := len(exe) * 100
				lkick := []string{}
				for n, v := range memlist {
					if MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0
				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					go func(to string, i int) {
						target := lkick[i]
						cl := nom[i]
						cl.DeleteOtherFromChats(to, []string{target})
					}(to, i)
				}
				logAccess(client, to, sender, "kickall", lkick, msg.ToType)
			} else if stiker.PRODUCTID == stkid7 && stiker.STICONID == stkpkgid7 {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if !InArray2(listuser, lists[i]) {
							listuser = append(listuser, lists[i])
						}
					}
				}
				if len(listuser) != 0 {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						lists := Setinvitetomsg(exe, to, listuser)
						if len(lists) != 0 {
									Cekbanwhois(client, to, lists)
						}
						logAccess(client, to, sender, "invite", listuser, msg.ToType)
					} else {
						client.SendMessage(to, fancy("Please add another bot that has a ban invite."))
					}
				}
			} else if stiker.PRODUCTID == stkid8 && stiker.STICONID == stkpkgid8 {
				if len(Banned.Banlist) != 0 {
					msgcbn := fmt.Sprintf(MsgBan, len(Banned.Banlist))
					logAccess(client, to, sender, "clearban", Banned.Banlist, msg.ToType)
					client.SendMessage(to, fancy(msgcbn))
					Banned.Banlist = []string{}
					Banned.Exlist = []string{}
				} else {
					client.SendMessage(to, fancy("Ban list is empty."))
				}
			} else if stiker.PRODUCTID == stkid9 && stiker.STICONID == stkpkgid9 {
				_, memlist2, memlist := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist2 {
					if InArray2(Squadlist, mid) {
						cl := GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				lkick := []string{}
				max := len(exe) * 10
				for n, v := range memlist {
					if MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0

				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					target := lkick[i]
					cl := nom[i]
					ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
				}
				logAccess(client, to, sender, "cancelall", lkick, msg.ToType)
			}
		}
	}}
	if AutoTranslate {
		if MemUser(to, sender) {
			if strings.Contains(pesan, pesan) {
				client.TranslateMe("ID", pesan)
				filepath := fmt.Sprintf("trMe.txt")
				b, err := ioutil.ReadFile(filepath)
				if err != nil {
					fmt.Print(err)
				}
				code := string(b)
				list := code
				client.SendMessage(to, fancy(list))
			}
		}
	}
	if AutoTranslate {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				Type := TypeTrans
				client.TranslateYou(Type, pesan)
				filepath := fmt.Sprintf("trYou.txt")
				b, err := ioutil.ReadFile(filepath)
				if err != nil {
					fmt.Print(err)
				}
				code := string(b)
				list := code
				client.SendMessage(to, fancy(list))
			}
		}
	}
	if RemoteOwner {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := StripOut(pesan)
                           var err error
				if err != nil {
					client.SendMessage(to, fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, MidRemote)
					listuser := []string{}
					x := 2
					if len(simpan) != 0 {
						for i := range simpan {
							if !InArray2(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						Checklistaccess(client, to, listuser, x, sender)
					}
				}
			}
			RemoteOwner = false
			MidRemote = []string{}
		}
	}
	if RemoteMaster {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := StripOut(pesan)
                           var err error
				if err != nil {
					client.SendMessage(to, fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, MidRemote)
					listuser := []string{}
					x := 3
					if len(simpan) != 0 {
						for i := range simpan {
							if !InArray2(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						Checklistaccess(client, to, listuser, x, sender)
					}
				}
			}
			RemoteMaster = false
			MidRemote = []string{}
		}
	}
	if RemoteAdmin {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := StripOut(pesan)
                           var err error
				if err != nil {
					client.SendMessage(to, fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, MidRemote)
					listuser := []string{}
					x := 4
					if len(simpan) != 0 {
						for i := range simpan {
							if !InArray2(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						Checklistaccess(client, to, listuser, x, sender)
					}
				}
			}
			RemoteAdmin = false
			MidRemote = []string{}
		}
	}
	if RemoteContact {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := StripOut(pesan)
                           var err error
				if err != nil {
					client.SendMessage(to, fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, MidRemote)
					listuser := []string{}
					if len(simpan) != 0 {
						for i := range simpan {
							if !InArray2(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
					}
					if len(listuser) != 0 {
						for _, mek := range listuser {
							client.FindAndAddContactsByMidV2(mek)
							go client.SendContact(to, mek)
						}
					}
				}
			}
			RemoteContact = false
			MidRemote = []string{}
		}
	}
	if RemoteBan {
		if !MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := StripOut(pesan)
                           var err error
				if err != nil {
					client.SendMessage(to, fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, MidRemote)
					listuser := []string{}
					if len(simpan) != 0 {
						for i := range simpan {
							if !InArray2(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
					}
					client.SendPollMention(to, "Added banlist member:", listuser)
					if len(listuser) != 0 {
						for _, mek := range listuser {
							Banned.AddBan(mek)
						}
					}
				}
			}
			RemoteBan = false
			MidRemote = []string{}
		}
	}
//AutoJointicket
	if AutoJointicket {
		if strings.Contains(pesan, "/ti/g") {
			regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
			links := regex.FindAllString(msg.Text, -1)
			tickets := []string{}
			for _, link := range links {
				if !InArray2(tickets, link) {
					tickets = append(tickets, link)
				}
			}
			for _, tick := range tickets {
				tuk := strings.Split(tick, "/")
				ntk := len(tuk) - 1
				ti := tuk[ntk]
				fmt.Println(ti)
				tkt := client.FindChatByTicket(ti)
				client.AcceptTicket(tkt.Chat.ChatMid, ti)
				exe := []*linetcr.Account{}
				for _, p := range bk {
					if p.Limited == false {
						err := p.AcceptTicket(tkt.Chat.ChatMid, ti)
						if err == nil {
							exe = append(exe, p)
						}
					}
				}
				newsend := ""
				if len(exe) != 0 {
					if TypeJoin == "normal" {
						newsend += "Succes Accept Group Ticket"
					} else if TypeJoin == "nuke" {
						go Nukjoin(exe[0], op.CreatedTime, tkt.Chat.ChatMid)
						newsend += "Succes Accept Ticket Nuke Group"
					}
				}
			}
		}
	}
//BOMLIKE
       if BomLike && !MemUser(to, sender) && op.Message.ContentType == 16 {
       	posturl := msg.ContentMetadata["postEndUrl"]
       	if msg.ContentMetadata["serviceType"] == "MH" {
             		mids := strings.Replace(posturl, "https://line.me/R/home/post?userMid=", "", 1)
                    post := strings.Split(mids, "&postId=")
                    client.Bomlike(to, mids, post[1])
		}
	}
//TIMELINE
       if AutoLike && op.Message.ContentType == 16 {
       	posturl := msg.ContentMetadata["postEndUrl"]
       	if msg.ContentMetadata["serviceType"] == "MH" {
             		mids := strings.Replace(posturl, "https://line.me/R/home/post?userMid=", "", 1)
                    post := strings.Split(mids, "&postId=")
                    client.Timeline(to, mids, post[1],"likepost",)
		}
	}
//MEDIA_DL
	if MediaDl {
		if strings.Contains(pesan, "tiktok.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"tiktok")
			}
		}
		if strings.Contains(pesan, "instagram.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"instagram")
			}
		}
		if strings.Contains(pesan, "smule.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"smule")
			}
		}
		if strings.Contains(pesan, "sck.io/p/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"snackvideo")
			}
		}
		if strings.Contains(pesan, "facebook.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"facebook")
			}
		}
	}
//BATAS
	if room.ProLink && MemUser(to, msg.From_) && strings.Contains(pesan, "https://") {
		client.DeleteOtherFromChats(to, []string{msg.From_})
		Banned.AddBan(msg.From_)
	}
	if room.AntiTag && MemUser(to, msg.From_) && strings.Contains(pesan, "@All") {
		client.DeleteOtherFromChats(to, []string{msg.From_})
		Banned.AddBan(msg.From_)
	}
	if op.Message.ContentType == 22 {
	      if room.ProFlex && MemUser(to, msg.From_) {
		       client.DeleteOtherFromChats(to, []string{msg.From_})
		       Banned.AddBan(msg.From_)
                    if msg.ContentMetadata["FLEX_VER"] == "u1" || msg.ContentMetadata["ORGCONTP"] == "FLEX" {
		              client.DeleteOtherFromChats(to, []string{msg.From_})
		              Banned.AddBan(msg.From_)
	             }
	       }
	}
	if op.Message.ContentType == 1 {
		if room.ProImage && MemUser(to, sender) {
		       client.DeleteOtherFromChats(to, []string{sender})
		       Banned.AddBan(sender)
	       }
	}
	if op.Message.ContentType == 2 {
		if room.ProVideo && MemUser(to, sender) {
		       client.DeleteOtherFromChats(to, []string{sender})
		       Banned.AddBan(sender)
	       }
	}
//Broadcast_VIDEO_IMAGE
	if BcImage {
		StartBc = true
		timeabort = time.Now()
	}
	if GBcImage {
		GStartBc = true
		timeabort = time.Now()
	}
	if FBcImage {
		FStartBc = true
		timeabort = time.Now()
	}
	if SAVEBcImage {
		startSaveBc = true
		timeabort = time.Now()
	}
	if op.Message.ContentType == 1 {
		if StartBc {
			if !MemUser(to, sender) {
				if BcImage {
					nCount := 0
					typec := "lmid"
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					addConSingle(lists)
					if len(lists) != 0 {
						for _, i := range lists {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendFoto(i, path)
						}
					}
					StartBc = false
					BcImage = false
					client.SendMessage(to, fancy("Success broadcast image to mid"))
				}
				timeabort = time.Now()
			}
		}
		if GStartBc {
			if !MemUserN(to, sender) {
				if GBcImage {
					gr, _ := client.GetGroupIdsJoined()
					for _, gi := range gr {
						path, _ := client.DownloadObjectMsg(msg.ID)
						client.SendFoto(gi, path)
					}
					GStartBc = false
					GBcImage = false
					client.SendMessage(to, fancy("Success all groupcast image"))
				}
				timeabort = time.Now()
			}
		}
		if FStartBc {
			if !MemUserN(to, sender) {
				if FBcImage {
					for _, cl := range ClientBot {
						fl, _ := cl.GetAllContactIds()
						for _, x := range fl {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendFoto(x, path)
							time.Sleep(3 * time.Second)
						}	
					}
					FStartBc = false
					FBcImage = false
					client.SendMessage(to, fancy("Success all friendcast image"))
				}
				timeabort = time.Now()
			}
		}
		if startSaveBc {
			if !MemUserN(to, sender) {
				if SAVEBcImage {
					client.Downloadbc(msg.ID)
					startSaveBc = false
					SAVEBcImage = false
					client.SendMessage(to, fancy("Success save image broadcast"))
                           	time.Sleep(3 * time.Second)
					client.SendMessage(to, fancy("Broadcast all group runs every 5 minutes"))
				}
				timeabort = time.Now()
			}
		}
	}
	if BcVideo {
		StartBcV = true
		timeabort = time.Now()
	}
	if GBcVideo {
		GStartBcV = true
		timeabort = time.Now()
	}
	if FBcVideo {
		FStartBcV = true
		timeabort = time.Now()
	}
	if op.Message.ContentType == 2 {
		if StartBcV {
			if !MemUser(to, sender) {
				if BcVideo {
					nCount := 0
					typec := "lmid"
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					addConSingle(lists)
					if len(lists) != 0 {
						for _, i := range lists {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendVid(i, path)
						}
					}
					StartBcV = false
					BcVideo = false
					client.SendMessage(to, fancy("Success broadcast video to mid"))
				}
				timeabort = time.Now()
			}
		}
		if GStartBcV {
			if !MemUserN(to, sender) {
				if GBcVideo {
					gr, _ := client.GetGroupIdsJoined()
					for _, gi := range gr {
						path, _ := client.DownloadObjectMsg(msg.ID)
						client.SendVid(gi, path)
					}
					GStartBcV = false
					GBcVideo = false
					client.SendMessage(to, fancy("Success all groupcast video"))
				}
				timeabort = time.Now()
			}
		}
		if FStartBcV {
			if !MemUserN(to, sender) {
				if FBcVideo {
					for _, cl := range ClientBot {
						fl, _ := cl.GetAllContactIds()
						for _, x := range fl {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendVid(x, path)
							time.Sleep(3 * time.Second)
						}	
					}
					FStartBcV = false
					FBcVideo = false
					client.SendMessage(to, fancy("Success all friendcast video"))
				}
				timeabort = time.Now()
			}
		}
	}
//Changepicture
	if ChangPict && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, fancy("Please Send Your Image !!!"))
				StartChangeImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangCover && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, fancy("Please Send Your Image !!!"))
				StartChangeImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangVpict && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, fancy("Please Send Your Video !!!"))
				StartChangevImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangVcover && !AllCheng && !StartChangeImg {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				client.SendMessage(to, fancy("Please Send Your Video !!!"))
				StartChangevImg = true
			}
			timeabort = time.Now()
		}
	} else if ChangName {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				if MsgName != "" {
					for i := range changepic {
						if TimeDown(i) {
							star := MsgName
							changepic[i].UpdateProfileName(star)
							changepic[i].SendMessage(to, fancy("Profile name success updated."))
						}
					}
				} else {
					client.SendMessage(to, fancy("Add name first."))
				}
				ChangName = false
				MsgName = ""
			}
			timeabort = time.Now()
		}
	} else if ChangeBio {
		if len(mentionlist) != 0 {
			changepic = []*linetcr.Account{}
			for _, ym := range mentionlist {
				if InArray2(Squadlist, ym) {
					cl := GetKorban(ym)
					if !linetcr.Checkarri(changepic, cl) {
						changepic = append(changepic, cl)
					}
				}
			}
			if len(changepic) != 0 {
				if MsgBio != "" {
					for i := range changepic {
						if TimeDown(i) {
							star := MsgBio
							changepic[i].UpdateProfileBio(star)
							changepic[i].SendMessage(to, fancy("Profile status success updated."))
						}
					}
				} else {
					client.SendMessage(to, fancy("Add Status first."))
				}
				ChangeBio = false
				MsgBio = ""
			}
			timeabort = time.Now()
		}
	}
	if op.Message.ContentType == 1 {
		if StartChangeImg && len(changepic) != 0 {
			if !MemUser(to, sender) {
				if ChangPict {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for n, p := range changepic {
							if TimeDown(n) {
								go func(p *linetcr.Account) {
									if StartChangevImg2 {
										err := p.UpdatePictureProfile(path, "v")
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, fancy("Update dual profile failure."))
										} else {
											p.SendMessage(to, fancy("Update video picture done."))
										}
									} else {
										err := p.UpdatePictureProfile(path, "p")
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, fancy("Update picture profile failure."))
										} else {
											p.SendMessage(to, fancy("Update Image picture done."))
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.Remove(path)
					} else {
						fmt.Println(err)
						if StartChangevImg2 {
							client.SendMessage(to, fancy("Download video picture Failure."))
						} else {
							client.SendMessage(to, fancy("Download Image picture Failure."))
						}
					}
					StartChangevImg2 = false
					StartChangeImg = false
					ChangPict = false
				} else if ChangCover {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for n, p := range changepic {
							if TimeDown(n) {
								go func(p *linetcr.Account) {
									if StartChangevImg2 {
										err := p.UpdateCoverWithVideo(path)
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, fancy("Update video cover failure."))
										} else {
											p.SendMessage(to, fancy("Update video cover done."))
											time.Sleep(2 * time.Second)
										}
									} else {
										err := p.UpdateCover(path)
										if err != nil {
											fmt.Println(err)
											p.SendMessage(to, fancy("Update picture cover failure."))
										} else {
											p.SendMessage(to, fancy("Update Image cover done."))
											time.Sleep(2 * time.Second)
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.Remove(path)
					} else {
						fmt.Println(err)
						if StartChangevImg2 {
							client.SendMessage(to, fancy("Download video cover Failure."))
						} else {
							client.SendMessage(to, fancy("Download Image cover Failure."))
						}
					}
					StartChangevImg2 = false
					StartChangeImg = false
					ChangCover = false
				}
				timeabort = time.Now()
			}
		}
	} else if op.Message.ContentType == 16 {
		if msg.ContentMetadata["serviceType"] == "GB" && MemUser(msg.To, msg.From_) {
			if room.ProNote {
				exe, _ := SelectBot(client, to)
				if exe != nil {
					Setkickto(exe, to, []string{msg.From_})
					Banned.AddBan(msg.From_)
				}
			}
		}
		if msg.ContentMetadata["serviceType"] == "MH" {
			if MemUser(msg.To, msg.From_) {
				if room.ProPost {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		}
		if msg.ContentMetadata["locKey"] == "BA" || msg.ContentMetadata["locKey"] == "BT" {
			if MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		}
	} else if op.Message.ContentType == 18 {
		if msg.ContentMetadata["LOC_KEY"] == "BD" {
			if MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		} else if msg.ContentMetadata["LOC_KEY"] == "BB" {
			if MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		} else if msg.ContentMetadata["LOC_KEY"] == "BO" {
			if MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		}
	} else if op.Message.ContentType == 18 { //NEW PROCJECT
    	runtime.GOMAXPROCS(cpu)
	    if msg.ContentMetadata["serviceType"] == "AB"  {
	   	   if room.ProAlbum && MemUser(to, msg.From_){
	   	   	  Banned.AddBan(msg.From_)
	   	  	  client.DeleteOtherFromChats(to, []string{msg.From_})
	   	  }
	   }
	} else if op.Message.ContentType == 6 {
		if msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
			if MemUser(msg.To, msg.From_) {
				if room.ProCall {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, []string{msg.From_})
						Banned.AddBan(msg.From_)
					}
				}
			}
		}
              if msg.ToType == 0 {
                    if op.Type != 50 && room.ProSpam {
		             if msg.ContentMetadata["GC_EVT_TYPE"] == "I" {
			            if MemUser(msg.To, msg.From_) {
					     exe, _ := SelectBot(client, to)
					     if exe != nil {
						     Setkickto(exe, to, []string{msg.From_})
						     Banned.AddBan(msg.From_)
					      }
					}
				}
			}
		}
	} else if op.Message.ContentType == 14 {
		if MemUser(msg.To, msg.From_) {
			if room.ProFile {
				exe, _ := SelectBot(client, to)
				if exe != nil {
					Setkickto(exe, to, []string{msg.From_})
					Banned.AddBan(msg.From_)
				}
			}
		}
	} else if op.Message.ContentType == 7 {
		if MemUser(msg.To, msg.From_) {
			if room.ProSticker {
				exe, _ := SelectBot(client, to)
				if exe != nil {
					Setkickto(exe, to, []string{msg.From_})
					Banned.AddBan(msg.From_)
				}
			}
		}//COMMAND_STICKER
		if GetCodeprem(5, sender, to) {
		if !MemUser(msg.To, msg.From_) {
			if getStickerKick == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid = op.Message.ContentMetadata["STKID"]
					stkpkgid = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerKick = 0
					client.SendMessage(to, fancy("kick by sticker updated"))
				}
			} else if getStickerRespon == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid2 = op.Message.ContentMetadata["STKID"]
					stkpkgid2 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerRespon = 0
					client.SendMessage(to, fancy("respon by sticker updated"))
				}
			} else if getStickerStayall == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid3 = op.Message.ContentMetadata["STKID"]
					stkpkgid3 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerStayall = 0
					client.SendMessage(to, fancy("stayall by sticker updated"))
				}
			} else if getStickerLeave == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid4 = op.Message.ContentMetadata["STKID"]
					stkpkgid4 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerLeave = 0
					client.SendMessage(to, fancy("leave by sticker updated"))
				}
			} else if getStickerKickall == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid5 = op.Message.ContentMetadata["STKID"]
					stkpkgid5 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerKickall = 0
					client.SendMessage(to, fancy("kickall by sticker updated"))
				}
			} else if getStickerBypass == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid6 = op.Message.ContentMetadata["STKID"]
					stkpkgid6 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerBypass = 0
					client.SendMessage(to, fancy("bypass by sticker updated"))
				}
			} else if getStickerInvite == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid7 = op.Message.ContentMetadata["STKID"]
					stkpkgid7 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerInvite = 0
					client.SendMessage(to, fancy("invite by sticker updated"))
				}
			} else if getStickerClearban == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid8 = op.Message.ContentMetadata["STKID"]
					stkpkgid8 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerClearban = 0
					client.SendMessage(to, fancy("clearban by sticker updated"))
				}
			} else if getStickerCancelall == 1 {
				if !MemUser(msg.To, msg.From_) {
					stkid9 = op.Message.ContentMetadata["STKID"]
					stkpkgid9 = op.Message.ContentMetadata["STKPKGID"]
					SaveBackup()
					getStickerCancelall = 0
					client.SendMessage(to, fancy("cancelall by sticker updated"))
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid && op.Message.ContentMetadata["STKPKGID"] == stkpkgid {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
							if AutoBan {
										Banned.AddBan(lists[i])
							}
							listuser = append(listuser, lists[i])
						}
					}
				}
				fmt.Println(listuser)
				if len(listuser) != 0 {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						Setkickto(exe, to, listuser)
						logAccess(client, to, sender, "kick", listuser, msg.ToType)
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid2 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid2 {
				for _, p := range bk {
					go p.SendMessage(to, fancy(MsgRespon))
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid3 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid3 {
				if room.Limit {
					client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
					return
				}
				numb := len(ClientBot)
				if numb > 0 && numb <= len(ClientBot) {
					GetSquad(client, to)
					room := linetcr.GetRoom(to)
					aa := len(room.Client)
					if aa > numb {
						c := aa - numb
						ca := 0
						list := append([]*linetcr.Account{}, room.Client...)
						sort.Slice(list, func(i, j int) bool {
							return list[i].KickPoint > list[j].KickPoint
						})
						for _, o := range list {
							o.LeaveGroup(to)
							ca = ca + 1
							if ca == c {
								break
							}
						}
						GetSquad(client, to)
					} else if aa < numb {
						ti, err := client.ReissueChatTicket(to)
						if err == nil {
							go client.UpdateChatQrV2(to, false)
							all := []*linetcr.Account{}
							room := linetcr.GetRoom(to)
							cuk := room.Client
							for _, x := range ClientBot {
								if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
									all = append(all, x)
								}
							}
							sort.Slice(all, func(i, j int) bool {
								return all[i].KickPoint < all[j].KickPoint
							})
							g := numb - aa
							var wg sync.WaitGroup
							wi := GetSquad(client, to)
							for i := 0; i < len(all); i++ {
								if i == g {
									break
								}
								l := all[i]
								if l != client && !linetcr.InArrayCl(wi, l) {
									if !l.Limited {
										wg.Add(1)
										go func() {
											l.AcceptTicket(to, ti)
											wg.Done()
										}()
									}
								}
							}
							wg.Wait()
							client.UpdateChatQrV2(to, true)
							GetSquad(client, to)
							logAccess(client, to, sender, "bringbot", []string{}, 2)
							SaveBackup()
							aa := len(room.Client)
							var name string
							name = fmt.Sprintf("Ready %v bots here", aa)
							client.SendMessage(to, fancy(name))
							//newsend += name + "\n"
						}
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid4 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid4 {
				_, mem := client.GetGroupInvitation(to)
				anu := []string{}
				for m := range mem {
					if InArray2(Squadlist, m) {
						anu = append(anu, m)
					}
				}
				if len(anu) != 0 {
					for _, mid := range anu {
						cl := GetKorban(mid)
						cl.AcceptGroupInvitationNormal(to)
					}
				}
				GetSquad(client, to)
				room := linetcr.GetRoom(to)
				bk = room.Client
				for _, cl := range bk {
					go cl.LeaveGroup(to)
				}
				if LogGroup == to {
					LogMode = false
					LogGroup = ""
				}
				linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				SaveBackup()
				logAccess(client, to, sender, "leave", []string{}, msg.ToType)
			} else if op.Message.ContentMetadata["STKID"] == stkid5 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid5 {
				_, memlist, _ := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist {
					if InArray2(Squadlist, mid) {
						cl := GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				max := len(exe) * 100
				lkick := []string{}
				for n, v := range memlist {
					if MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0
				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					go func(to string, i int) {
						target := lkick[i]
						cl := nom[i]
						cl.DeleteOtherFromChats(to, []string{target})
					}(to, i)
				}
				logAccess(client, to, sender, "kickall", lkick, msg.ToType)
			} else if op.Message.ContentMetadata["STKID"] == stkid7 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid7 {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if !InArray2(listuser, lists[i]) {
							listuser = append(listuser, lists[i])
						}
					}
				}
				if len(listuser) != 0 {
					exe, _ := SelectBot(client, to)
					if exe != nil {
						lists := Setinvitetomsg(exe, to, listuser)
						if len(lists) != 0 {
									Cekbanwhois(client, to, lists)
						}
						logAccess(client, to, sender, "invite", listuser, msg.ToType)
					} else {
						client.SendMessage(to, fancy("Please add another bot that has a ban invite."))
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid8 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid8 {
				if len(Banned.Banlist) != 0 {
					msgcbn := fmt.Sprintf(MsgBan, len(Banned.Banlist))
					logAccess(client, to, sender, "clearban", Banned.Banlist, msg.ToType)
					client.SendMessage(to, fancy(msgcbn))
					Banned.Banlist = []string{}
					Banned.Exlist = []string{}
				} else {
					client.SendMessage(to, fancy("Ban list is empty."))
				}
			} else if op.Message.ContentMetadata["STKID"] == stkid9 && op.Message.ContentMetadata["STKPKGID"] == stkpkgid9 {
				_, memlist2, memlist := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist2 {
					if InArray2(Squadlist, mid) {
						cl := GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				lkick := []string{}
				max := len(exe) * 10
				for n, v := range memlist {
					if MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0

				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					target := lkick[i]
					cl := nom[i]
					ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
				}
				logAccess(client, to, sender, "cancelall", lkick, msg.ToType)
			}
		}}
	} else if op.Message.ContentType == 13 {
		if MemUser(msg.To, msg.From_) {
			if room.ProContact {
				exe, _ := SelectBot(client, to)
				if exe != nil {
					Setkickto(exe, to, []string{msg.From_})
					Banned.AddBan(msg.From_)
				}
			}
		}
	} else if op.Message.ContentType == 2 {
		if StartChangevImg && len(changepic) != 0 {
			if !MemUser(to, sender) {
				if ChangVpict {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for _, p := range changepic {
							go func(p *linetcr.Account) {
								err := p.UpdateVideoProfile(path)
								if err != nil {
									fmt.Println(err)
									p.SendMessage(to, fancy("Update video profile failure."))
								}
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, fancy("Upload video done, now Please Send Your Image !!!"))
						os.Remove(path)
						StartChangevImg2 = true
						ChangPict = true
						StartChangeImg = true
						ChangVpict = false
						StartChangevImg = false
					} else {
						fmt.Println(err)
						client.SendMessage(to, fancy("Download Image Failure."))
					}
				} else if ChangVcover {
					path, err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(changepic))
						for _, p := range changepic {
							go func(p *linetcr.Account) {
								p.UpdateCoverVideo(path)
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, fancy("Upload video done, now Please Send Your Image !!!"))
						os.Remove(path)
						StartChangevImg2 = true
						StartChangeImg = true
						ChangCover = true
						StartChangevImg = false
						ChangVcover = false
					} else {
						fmt.Println(err)
						client.SendMessage(to, fancy("Download Image Failure."))
					}
				}
				timeabort = time.Now()
			}
		}
	} else if msg.ContentType == 0 && msg.Text != "" {
		if FixedToken && From_Token == msg.From_ && Group_Token == to {
			anu := StripOut(pesan)
			if anu == "done reboot" {
				timeabort = time.Now()
				FixedToken = false
				From_Token = ""
				Group_Token = ""
				Data.RestartBack = fmt.Sprintf("%v", to)
				SaveBackup()
				client.SendMessage(to, fmt.Sprintf("Done Editing, You Can Use After %v Seconds.", len(ClientBot)))
				ReloginProgram()
			}
			timeabort = time.Now()
		}
		if room.Automute && MemUser(to, msg.From_) {
			if client.Limited == false {
				client.DeleteOtherFromChats(to, []string{msg.From_})
			} else {
				for _, bot := range bk {
					if bot.Limited == false {
						bot.DeleteOtherFromChats(to, []string{msg.From_})
						break
					}
				}
			}
		} else {
			if MemBan2(to, msg.From_) && MemUser(to, msg.From_) {
				if client.Limited == false {
					client.DeleteOtherFromChats(to, []string{msg.From_})
				} else {
					for _, bot := range bk {
						if bot.Limited == false {
							bot.DeleteOtherFromChats(to, []string{msg.From_})
							break
						}
					}
				}
			}
		}
		cmds := gettxt(sender, client, pesan, Rname, Sname, client.MID, mentionlist, to)
		text := op.Message.Text
		newsend := ""
		var pesan = strings.ToLower(text)
		for _, cmd := range strings.Split(cmds, ",") {
			if strings.HasPrefix(cmd, "creator") && cmd != "creator" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 15
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "maker") && cmd != "makers" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 13
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "creators" {
				rngcmd := GetComs(0, "creators")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Creator) != 0 {
							list := "Creator List:\n"
							for num, xd := range UserBot.Creator {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if cmd == "makers" {
				rngcmd := GetComs(1, "makers")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Maker) != 0 {
							list := "Maker List:\n"
							for num, xd := range UserBot.Maker {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Maker list is empty.\n"
						}
					}
				}
			} else if cmd == "clearcreator" {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Creator) != 0 {
							newsend += fmt.Sprintf("Cleared %v Makerlist\n", len(UserBot.Creator))
							UserBot.ClearCreator()
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmaker" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Maker) != 0 {
							newsend += fmt.Sprintf("Cleared %v Makerlist\n", len(UserBot.Maker))
							UserBot.ClearMaker()
						} else {
							newsend += "Maker list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "uncreator") {
				if GetCodeprem(0, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 10
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "uncreator"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Creator)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmaker") {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmaker"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Maker)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "seller") && cmd != "seller" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "sellers" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Seller) != 0 {
							list := " ✠ 𝗦𝗲𝗹𝗹𝗲𝗿𝘀 ✠ \n"
							for num, xd := range UserBot.Seller {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if cmd == "clearseller" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Seller) != 0 {
							newsend += fmt.Sprintf("Cleared %v sellerlist\n", len(UserBot.Seller))
							UserBot.ClearSeller()
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unseller") {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unseller"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Seller)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "listcmd" {
				rngcmd := GetComs(5, "listcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						list := Cmdlistcheck()
						client.SendMessage(to, fancy(list))
					}
				}
			} else if strings.HasPrefix(cmd, "expel") {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "access" || cmd == Commands.Access && Commands.Access != "" {
				rngcmd := GetComs(7, "ess")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "𝗔𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁:"
						if len(DEVELOPER) != 0 {
							listadm += "\n\n >Developer:"
							for num, xd := range DEVELOPER {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Creator) != 0 {
							listadm += "\n\n >Creator:"
							for num, xd := range UserBot.Creator {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Maker) != 0 {
							listadm += "\n\n >Maker:"
							for num, xd := range UserBot.Maker {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Seller) != 0 {
							listadm += "\n\n >Seller:"
							for num, xd := range UserBot.Seller {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Buyer) != 0 {
							listadm += "\n\n >Buyer:"
							for num, xd := range UserBot.Buyer {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Owner) != 0 {
							listadm += "\n\n >Owner:"
							for num, xd := range UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Master) != 0 {
							listadm += "\n\n >Master:"
							for num, xd := range UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(UserBot.Admin) != 0 {
							listadm += "\n\n >Admin:"
							for num, xd := range UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "𝗔ccess is empty.\n"
						}
					}
				}
			} else if cmd == "allbanlist" || cmd == Commands.Allbanlist && Commands.Allbanlist != "" {
				rngcmd := GetComs(5, "allbanlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listadm := AllBanList(client)
						if listadm != "All Banlist:" {
							newsend += listadm + "\n"
						} else {
							newsend += "Banlist is empty.\n"
						}
					}
				}
			} else if cmd == "gaccess" || cmd == Commands.Gaccess && Commands.Gaccess != "" {
				rngcmd := GetComs(9, "access")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "𝗚𝗮𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁:"
						if len(room.Gowner) != 0 {
							listadm += "\n\n >Gowner:"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(room.Gadmin) != 0 {
							listadm += "\n\n >Gadmin:"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								new := client.Getcontactuser(xd)
								if new != nil {
									listadm += "\n " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									listadm += "\n " + rengs + ". " + x.DisplayName
								}
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "Gaccess is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "buyer") && cmd != "buyers" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setdate ") {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						ha := strings.Split((cmd), "setdate ")
						haj := ha[1]
						haj = StripOut(haj)
						has := strings.Split(haj, "-")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(has[2])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							times := time2.Format(time.RFC3339)
							Data.Dalltime = times
							str := fmt.Sprintf("⚙️ Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\n⚙️ Remaining:\n  %v", botDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "addweek" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 7 * mont
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addday" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addmonth" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 30 * mont
						t := batas.Add(mont)
						Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", botDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "reboot" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						client.SendMessage(to, fancy("Waiting....."))
						ReloginProgram()
						client.SendMessage(to, fancy("Done"))
					}
				}
			} else if cmd == "runall" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						client.SendMessage(to, fancy("Waiting....."))
						ReloginProgram()
						client.SendMessage(to, fancy("Done"))
					}
				}
			} else if strings.HasPrefix(cmd, "unbuyer") {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unbuyer"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Buyer)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "checkram" || cmd == Commands.Checkram && Commands.Checkram != "" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						v, _ := mem.VirtualMemory()
						r := fmt.Sprintf("  ≻ Cpu : %v core\n  ≻ Ram : %v mb\n  ≻ Free : %v mb\n  ≻ Cache : %v mb\n  ≻ UsedPercent : %f %%", cpu, bToMb(v.Used+v.Free+v.Buffers+v.Cached), bToMb(v.Free), bToMb(v.Buffers+v.Cached), v.UsedPercent)
						newsend += r + "\n"
					}
				}
			} else if cmd == "memory" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						v, _ := mem.VirtualMemory()
						r := fmt.Sprintf("  ≻ Cpu : %v core\n  ≻ Ram : %v mb\n  ≻ Free : %v mb\n  ≻ Cache : %v mb\n  ≻ UsedPercent : %f %%", cpu, bToMb(v.Used+v.Free+v.Buffers+v.Cached), bToMb(v.Free), bToMb(v.Buffers+v.Cached), v.UsedPercent)
						newsend += r + "\n"
					}
				}
			} else if cmd == "clearbuyer" {
				if GetCodeprem(3, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Buyer) != 0 {
							newsend += fmt.Sprintf("Cleared %v buyerlist\n", len(UserBot.Buyer))
							UserBot.ClearBuyer()
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "upimage" || cmd == Commands.Upimage && Commands.Upimage != "" {
				rngcmd := GetComs(4, "upimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangPict = true
						newsend += "Which bot's you want to update Pict.\n"
					}
				}
			} else if cmd == "upcover" || cmd == Commands.Upcover && Commands.Upcover != "" {
				rngcmd := GetComs(4, "upcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangCover = true
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if cmd == "upvimage" || cmd == Commands.Upvimage && Commands.Upvimage != "" {
				rngcmd := GetComs(4, "upvimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangVpict = true
						newsend += "Which bot's you want to update Pict ?.\n"
					}
				}
			} else if cmd == "upvcover" || cmd == Commands.Upvcover && Commands.Upvcover != "" {
				rngcmd := GetComs(4, "upvcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ChangVcover = true
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "unsend ") {
				rngcmd := GetComs(5, "unsend")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									Nganu, _ := client.GetRecentMessagesV2(op.Message.To)
									Mid := []string{}
									unsed := []string{}
									for _, chat := range Nganu {
										if InArray2(Squadlist, chat.From_) {
											Mid = append(Mid, chat.ID)
										}
									}
									for i := 0; i < len(Mid); i++ {
										if i < result2 {
											unsed = append(unsed, Mid[i])
										}
									}
									if len(unsed) != 0 {
										exess, _ := SelectallBot(client, to)
										if exess != nil {
											for i := range exess {
												Nganu2, _ := exess[i].GetRecentMessagesV2(op.Message.To)
												for _, chat := range Nganu2 {
													if chat.From_ == exess[i].MID {
														if InArray2(unsed, chat.ID) {
															exess[i].UnsendChatnume(to, chat.ID)
														}
													}
												}
											}
										}
									}
								} else {
									client.SendMessage(to, fancy("out of range."))
								}
							}
						} else {
							client.SendMessage(to, fancy("Msg not fund number"))
						}
					}
				}
			} else if cmd == "purgeall" || cmd == Commands.Purgeall && Commands.Purgeall != "" {
				rngcmd := GetComs(5, "purgeall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						gr, _ := client.GetGroupIdsJoined()
						all := []string{}
						for _, aa := range gr {
							_, memlist, _ := client.GetChatList(aa)
							lkicks := []string{}
							for _, v := range memlist {
								if MemUser(aa, v) {
									lkicks = append(lkicks, v)
								}
							}
	                                        c, _ := client.GetChats([]string{aa})
	                                        zxc := c.Chats[0].Extra.GroupExtra.InviteeMids
	                                        for k, _ := range zxc {
		                                        if IsBlacklist(client, k) == true {
			                                        go func(k string) {
				                                        go client.CancelChatInvitations(aa, []string{k})
			                                        }(k)
		                                        }
	                                        }
							lkick := []string{}
							for _, ban := range lkicks {
								if MemBan(aa, ban) {
									lkick = append(lkick, ban)
									all = append(all, ban)
								}
							}
							nom := []*linetcr.Account{}
							ilen := len(lkick)
							xx := 0
							exe := []*linetcr.Account{}
							for _, c := range linetcr.GetRoom(aa).Client {
								if !c.Limited {
									exe = append(exe, c)
								}
							}
							if len(exe) != 0 {
								for i := 0; i < ilen; i++ {
									if xx < len(exe) {
										nom = append(nom, exe[xx])
										xx += 1
									} else {
										xx = 0
										nom = append(nom, exe[xx])
									}
								}
								for i := 0; i < ilen; i++ {
									target := lkick[i]
									cl := nom[i]
									go cl.DeleteOtherFromChats(aa, []string{target})
								}
								time.Sleep(1 * time.Second)
                                               }
						}
						newsend += fmt.Sprintf("Success purgeall blacklist")
						logAccess(client, to, sender, "purgeall", all, msg.ToType)
					}
				}
			} else if cmd == "2purgeall"  {
				rngcmd := GetComs(5, "2purgeall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						gr, _ := client.GetGroupIdsJoined()
						for _, aa := range gr {
							go KickAllBan(client, aa)
							time.Sleep(1 * time.Second)
						}
						newsend += fmt.Sprintf("Success purgeall blacklist")
					}
				}
			} else if strings.HasPrefix(cmd, "gleave") {
				rngcmd := GetComs(5, "gleave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									su := "gleave"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := StripOut(str)
									hapuss := linetcr.Archimed(st, tempgroup)
									if len(hapuss) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									names := []string{}
									for _, gid := range hapuss {
										name, mem := client.GetGroupInvitation(gid)
										names = append(names, name)
										anu := []string{}
										for m := range mem {
											if InArray2(Squadlist, m) {
												anu = append(anu, m)
											}
										}
										if len(anu) != 0 {
											for _, mid := range anu {
												cl := GetKorban(mid)
												cl.AcceptGroupInvitationNormal(gid)
												linetcr.GetRoom(gid).ConvertGo(cl)
											}
										}
										GetSquad(client, gid)
										room := linetcr.GetRoom(gid)
										bk = room.Client
										for _, cl := range bk {
											go cl.LeaveGroup(gid)
										}
										if LogGroup == gid {
											LogMode = false
											LogGroup = ""
										}
										linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
									}
									strs := strings.Join(names, ", ")
									client.SendMessage(to, fancy("Bot's leave from group: \n\n"+strs))
									SaveBackup()
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invme ") {
				rngcmd := GetComs(5, "invme")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										GetSquad(client, gid)
										room := linetcr.GetRoom(gid)
										bk := room.Client
										name, mem, inv := client.GetChatList(gid)
										if InArray2(mem, msg.From_) {
											client.SendMessage(to, fancy("You was on group "+name))
											return
										} else {
											if InArray2(inv, msg.From_) {
												bk[0].CancelChatInvitations(gid, []string{msg.From_})
											}
											for _, cl := range bk {
												if !cl.Limited && !cl.Limitadd {
													AddContact2(cl, msg.From_)
													fl, _ := cl.GetAllContactIds()
													if InArray2(fl, msg.From_) {
														err := cl.InviteIntoGroupNormal(gid, []string{msg.From_})
														if err != nil {
															code := linetcr.GetCode(err)
															if code != 35 && code != 10 {
																client.SendMessage(to, fancy("You has invited to group "+name))
																return
															}
														} else {
															client.SendMessage(to, fancy("You has invited to group "+name))
															return
														}
													}
												}
											}
											newsend += "Sorry, all bot has invite banned"
										}
									} else {
										newsend += "out of range."
									}
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "decline ") {
				rngcmd := GetComs(2, "decline")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									num, _ := strconv.Atoi(result[1])
									gr := []string{}
									for i := range ClientBot {
										grs, _ := ClientBot[i].GetGroupsInvited()
										if len(grs) != 0 {
											for _, a := range grs {
												if !InArray2(gr, a) {
													gr = append(gr, a)
												}
											}
										}
									}
									grup, _ := client.GetGroups(gr)
									tempgroup := []string{}
									for _, gi := range grup {
										if !InArray2(tempgroup, gi.ChatMid) {
											tempgroup = append(tempgroup, gi.ChatMid)
										}
									}
									if num > 0 && num <= len(tempgroup) {
										exe := []*linetcr.Account{}
										gen := tempgroup[num-1]
										names, _, _ := client.GetChatList(tempgroup[num-1])
										for i := range ClientBot {
											if ClientMid[ClientBot[i].MID].Limited == false {
												grs, _ := ClientBot[i].GetGroupsInvited()
												if InArray2(grs, gen) {
													exe = append(exe, ClientBot[i])
												}
											}
										}
										if len(exe) != 0 {
											for i := range exe {
												exe[i].RejectChatInvitation(gen)
											}
											newsend += fmt.Sprintf("Successfully declined invitation for: %v\n", names)
										}
									} else {
										newsend += "out of range pendinglist.\n"
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "accept") && cmd != "acceptall" {
				rngcmd := GetComs(5, "accept")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									su := "accept"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := StripOut(str)
									hapuss := linetcr.Archimed(st, tempginv)
									if len(hapuss) == 0 {
										newsend += "Please input the right number\nSee group number with command groups"
									} else {
										names := []string{}
										for _, gid := range hapuss {
											name, mem := client.GetGroupInvitation(gid)
											names = append(names, name)
											anu := []string{}
											for m := range mem {
												if InArray2(Squadlist, m) {
													anu = append(anu, m)
												}
											}
											if len(anu) != 0 {
												for _, mid := range anu {
													cl := GetKorban(mid)
													cl.AcceptGroupInvitationNormal(gid)
													linetcr.GetRoom(gid).ConvertGo(cl)
												}
											}
										}
										str := strings.Join(names, ", ")
										newsend += "Bot's join to group \n\n" + str
									}
								}
							}
						}
					}
				}
			} else if cmd == "abort" {
				rngcmd := GetComs(5, "abort")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if remotegrupidto != "" {
							client.SendMessage(remotegrupidto, fancy("Done Have abort."))
						} else {
							newsend += "Done Have abort." + "\n"
						}
						abort()
					}
				}
			} else if cmd == "declineall" {
				rngcmd := GetComs(2, "declineall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						tempgroup := []string{}
						for i := range ClientBot {
							grs, _ := ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !InArray2(tempgroup, gi.ChatMid) {
										tempgroup = append(tempgroup, gi.ChatMid)
									}
									ClientBot[i].RejectChatInvitation(gi.ChatMid)
								}
								time.Sleep(1 * time.Second)
							}

						}
						if len(tempgroup) != 0 {
							newsend += fmt.Sprintf("Successfully declined invitations: (%v)\n", len(tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if cmd == "acceptall" {
				rngcmd := GetComs(4, "acceptall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						tempgroup := []string{}
						for i := range ClientBot {
							grs, _ := ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !InArray2(tempgroup, gi.ChatMid) {
										tempgroup = append(tempgroup, gi.ChatMid)
									}
									ClientBot[i].AcceptGroupInvitationNormal(gi.ChatMid)
									linetcr.GetRoom(gi.ChatMid).ConvertGo(ClientBot[i])
									time.Sleep(1 * time.Second)
								}
							}
						}
						if len(tempgroup) != 0 {
							newsend += fmt.Sprintf("Success accept bot %v Group\n", len(tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "upstatus") {
				rngcmd := GetComs(4, "upstatus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upstatus"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						MsgBio = str
						ChangeBio = true
						timeabort = time.Now()
						newsend += fmt.Sprintf("Which bot's should be Status %v", str)
					}
				}
			} else if strings.HasPrefix(cmd, "upname") {
				rngcmd := GetComs(4, "upname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						aa := utf8.RuneCountInString(str)
						if aa != 0 && aa <= 20 {
							MsgName = str
							ChangName = true
							timeabort = time.Now()
							newsend += fmt.Sprintf("Which bot's should be Name %v", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "tiktokdl") {
				rngcmd := GetComs(4, "tiktokdl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "tiktokdl"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := "https://api.minzteam.xyz/tiktokdl?url=" + str + "&apikey="+apikey
						client.SendVideoWithURL(to, data)
					}
				}
			} else if strings.HasPrefix(cmd, "youtubedl") {
				rngcmd := GetComs(4, "youtubedl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "youtubedl"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						client.Sendmediadl(to, str,"youtube")
					}
				}
			} else if cmd == "buyers" {
				rngcmd := GetComs(4, "buyers")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Buyer) != 0 {
							list := " ✠ 𝗯𝘂𝘆𝗲𝗿𝘀 ✠ \n"
							for num, xd := range UserBot.Buyer {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "history" {
				rngcmd := GetComs(5, "history")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveProHistory()
						list := fmt.Sprintf("History Bot:\n\n Kick: %v \n Cancel: %v \n Invited: %v", Data.Kikhistory, Data.Canclhistory, Data.Invhistory)
						client.SendMessage(to, fancy(list))
					}
				}
			} else if cmd == "2history" {
				rngcmd := GetComs(5, "2history")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						countK := 0
						countinv := 0
						countcancel := 0
						for i := range ClientBot {
							countK = countK + ClientBot[i].Ckick
							countinv = countinv + ClientBot[i].Cinvite
							countcancel = countcancel + ClientBot[i].Ccancel
						}
						list := fmt.Sprintf("History: \n\n Kick: %v \n Cancel: %v \n Invited: %v", countK, countcancel, countinv)
						client.SendMessage(to, fancy(list))
					}
				}
			} else if cmd == "clearhide" {
				rngcmd := GetComs(5, "clearhide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(checkHaid) != 0 {
							logAccess(client, to, sender, "clearhid", checkHaid, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v Hidelist\n", len(checkHaid))
							checkHaid = []string{}
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if cmd == "hidelist" {
				rngcmd := GetComs(5, "hidelist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(checkHaid) != 0 {
							list := " ✠ Hide List ✠ \n"
							for num, xd := range checkHaid {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unhide") || strings.HasPrefix(cmd, "delhide") {
				rngcmd := GetComs(5, "unhide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						list := ""
						listuser := []string{}
						nCount1 := 0
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User removed from hidelist:\n\n"
							test2 := "User not exist in hidelist:\n\n"
							for n, xx := range listuser {
								if InArray2(checkHaid, xx) {
									checkHaid = Remove(checkHaid, xx)
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								new := client.Getcontactuser(xx)
								if new != nil {
									list += rengs + ". Closed Account \n"
								} else {
									x, _ := client.GetContact(xx)
									list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

								}
							}
							if nCount1 != 0 {
								client.SendMessage(to, fancy(test1+list))
							} else {
								client.SendMessage(to, fancy(test2+list))
							}
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unhide"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, checkHaid)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											for _, i := range hapuss {
												if InArray2(checkHaid, i) {
													checkHaid = Remove(checkHaid, i)
													listuser = append(listuser, i)
												}
											}
											if len(listuser) != 0 {
												list += "User removed from hidelist:\n\n"
												for n, xx := range listuser {
													checkHaid = Remove(checkHaid, xx)
													rengs := strconv.Itoa(n + 1)
													new := client.Getcontactuser(xx)
													if new != nil {
														list += rengs + ". Closed Account \n"
													} else {
														x, _ := client.GetContact(xx)
														list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

													}
												}
												client.SendMessage(to, fancy(list))
											}
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "hide") && cmd != "hidelist" {
				rngcmd := GetComs(5, "hide")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						nCount1 := 0
						list := ""
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User added to hidelist:\n\n"
							test2 := "User already exist in hidelist:\n\n"
							for n, xx := range listuser {
								if !InArray2(checkHaid, xx) {
									checkHaid = append(checkHaid, xx)
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								new := client.Getcontactuser(xx)
								if new != nil {
									list += rengs + ". Closed Account \n"
								} else {
									x, _ := client.GetContact(xx)
									list += fmt.Sprintf("%v. %v\n", n+1, x.DisplayName)

								}
							}
							if nCount1 != 0 {
								client.SendMessage(to, fancy(test1+list))
							} else {
								client.SendMessage(to, fancy(test2+list))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "owner") && cmd != "owners" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unowner") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Owner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "clearowner" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Owner) != 0 {
							logAccess(client, to, sender, "clearowner", UserBot.Owner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v ownerlist\n", len(UserBot.Owner))
							UserBot.ClearOwner()
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if cmd == "notification on" {
				rngcmd := GetComs(3, "notification")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if LogGroup == to {
							LogMode = true
							newsend += "Already enabled.\n"
						} else {
							LogMode = true
							LogGroup = to
							newsend += "Notification is enabled.\n"
						}
					}
				}
			} else if cmd == "notification off" {
				rngcmd := GetComs(3, "notification")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if LogGroup == to {
							LogMode = false
							LogGroup = ""
							newsend += "Notification is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setsname ") {
				rngcmd := GetComs(4, "setsname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setsname ", "", 1)
						if result == "," || result == "" {
							MsSname = ","
						} else {
							MsSname = result
						}
						newsend += "Sname set to: " + Sname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setrname ") {
				rngcmd := GetComs(4, "setrname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setrname ", "", 1)
						if result == "," || result == "" {
							MsRname = ","
						} else {
							MsRname = result
						}
						newsend += "Succes update Rname to " + Rname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgrespon") {
				rngcmd := GetComs(4, "msgrespon")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgrespon"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						MsgRespon = str
						newsend += "Message respon set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "upgname") {
				rngcmd := GetComs(4, "upgname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "upgname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						client.UpdateChatName(to, str)
						newsend += "group name has been changed to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlogo") {
				rngcmd := GetComs(4, "setlogo")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "setlogo"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						Data.Logobot = str
						newsend += "Menu logo set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgwelcome") {
				rngcmd := GetComs(7, "msgwelcome")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgwelcome"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.WelcomeMsg = str
						newsend += "Message Welcome set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgleave") {
				rngcmd := GetComs(7, "msgleave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgleave"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLeave = str
						newsend += "Message Leave set to: " + str + "\n"
					}
				}

			} else if strings.HasPrefix(cmd, "msgclearban ") {
				rngcmd := GetComs(4, "msgclearban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgclearban ", "", 1)
						MsgBan = result
						newsend += "Message unban set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglurk") {
				rngcmd := GetComs(7, "msglurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msglurk"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLurk = str
						newsend += "Message sider set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "2addtoken") {
				rngcmd := GetComs(2, "2addtoken")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "2addtoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						Data.Authoken = append(Data.Authoken,str)
						newsend += "Add token: "+str+"\n\nCmd to runall"
						time.Sleep(2 * time.Millisecond)
						SaveBackup()
					}
				}
			} else if strings.HasPrefix(cmd, "addtoken") {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "addtoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						tokenlist := strings.Split(str, "\n")
						for _, token := range tokenlist {
							token_mid := token[:33]
							new := client.Getcontactuser(token_mid)
							if new != nil {
								client.SendMessage(to, fancy("Token Error : ")+token)
							} else {
								ct, _ := client.GetContact(token_mid)
								if !InArray2(Data.Authoken, token) {
									Data.Authoken = append(Data.Authoken, token)
									client.SendMessage(to, ct.DisplayName + fancy(" >> Token Login"))
								} else {
									client.SendMessage(to, ct.DisplayName + fancy(" >> Token Alerdy Login"))
								}
							}

						}

						SaveBackup()
						ReloginProgram()
					}
				}
			} else if strings.HasPrefix(cmd, "list token") {
				rngcmd := GetComs(2, "list token")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len (Data.Authoken) != 0 {
							list :="List Token: ↓"
							for num, xd := range Data.Authoken {
								num++
								rengs := strconv.Itoa(num)
								mids := strings.Split(xd, ":")
								var prof *talkservice.Contact
								prof, _ = client.GetContact(mids[0])
								name := prof.DisplayName
								list += "\n\n  "+rengs+". "+xd+"\n Name: "+name
							}
							list += "\n\nToken limited: ↓"
							for n, cl := range linetcr.KickBans {
								m := cl.MID
								no := n + 1
								pr, _ := client.GetContact(m)
								cl.Namebot = pr.DisplayName
								list += fmt.Sprintf("\n\n%v. %s\n Name: %v", no, cl.AuthToken, cl.Namebot)
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "untoken") {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "untoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						st := StripOut(str)
						hapuss := linetcr.Archimed(st, Data.Authoken)

						if len(hapuss) == 0 {
							newsend += "empty list\n"
						} else {
							can_reboot := false
							for _, token := range hapuss {
								if InArray2(Data.Authoken, token) {
									can_reboot = true
									Data.Authoken = Remove(Data.Authoken, token)
									token_mid := token[:33]
									name := "Closed Account"
									new := client.Getcontactuser(token_mid)
									if new == nil {
										x, _ := client.GetContact(token_mid)
										name = x.DisplayName
									}

									client.SendMessage(to, name + fancy(" >> Token Delete !!!"))
								}
							}
							if can_reboot {
								SaveBackup()
								ReloginProgram()
							} else {
								client.SendMessage(to, fancy("No user deleted"))
							}

						}
					}
				}
			} else if cmd == "status token" {
				rngcmd := GetComs(2, "status token")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						kb := "Status Token:"
						for no, tok := range Data.Authoken {
							mids := strings.Split(tok, ":")
							mid := mids[0]
							if InArray2(Squadlist, mid) {
								xxx := GetKorban(mid)
								if linetcr.InArrayCl(linetcr.KickBanChat, xxx) {
									if xxx.Namebot == "" {
										pr, _ := client.GetContact(mid)
										xxx.Namebot = pr.DisplayName
									}
									kb += fmt.Sprintf("\n\n%v. %s\nMid: %s\nStatus: Banned", no+1, xxx.Namebot, xxx.MID)
								} else {
									kb += fmt.Sprintf("\n\n%v. %s\nmid: %s\nstatus: Normal", no+1, xxx.Namebot, xxx.MID)
								}
							}
						}
						newsend += kb
					}
				}
			} else if cmd == "remtokenbans" {
				rngcmd := GetComs(2, "remtokenbans")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						timeabort = time.Now()
						FixedToken = true
						From_Token = sender
						Group_Token = to
						if FixedToken && From_Token == sender {
							kb := "Deleted Banned:"
							for no, tok := range Data.Authoken {
								mids := strings.Split(tok, ":")
								mid := mids[0]
								Bots := []*linetcr.Account{}
								if InArray2(Squadlist, mid) {
									xxx := GetKorban(mid)
									if linetcr.InArrayCl(linetcr.KickBanChat, xxx) || xxx.Frez {
										if xxx.Namebot == "" {
											pr, _ := client.GetContact(mid)
											xxx.Namebot = pr.DisplayName
										}
										Data.Authoken = Remove(Data.Authoken, tok)
										Squadlist = Remove(Squadlist, mid)
										Bots = append(Bots, xxx)
										if linetcr.InArrayCl(ClientBot, xxx) {
											ClientBot = linetcr.RemoveCl(ClientBot, xxx)
										}
										kb += fmt.Sprintf("\n\n%v. %s\nMid: %s", no+1, xxx.Namebot, xxx.MID)
									}
								}
								if len(Bots) != 0 {
									linetcr.RemoveBot(Bots)
								}
								for i := range ClientBot {
									for _, x := range ClientBot[i].Squads {
										if !InArray2(Squadlist, x) {
											ClientBot[i].Squads = Remove(ClientBot[i].Squads, x)
										}
									}
								}
							}
							SaveData()
							newsend += kb
						}
					}
				}
			} else if cmd == "removelimits" {
				rngcmd := GetComs(2, "removelimits")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						timeabort = time.Now()
						FixedToken = true
						From_Token = sender
						Group_Token = to
						if FixedToken && From_Token == sender {
							kb := "Deleted Limit:"
							for no, tok := range Data.Authoken {
								mids := strings.Split(tok, ":")
								mid := mids[0]
								Bots := []*linetcr.Account{}
								if InArray2(Squadlist, mid) {
									xxx := GetKorban(mid)
									if linetcr.InArrayCl(linetcr.KickBans, xxx) || xxx.Frez {
										if xxx.Namebot == "" {
											pr, _ := client.GetContact(mid)
											xxx.Namebot = pr.DisplayName
										}
										Data.Authoken = Remove(Data.Authoken, tok)
										Squadlist = Remove(Squadlist, mid)
										Bots = append(Bots, xxx)
										if linetcr.InArrayCl(ClientBot, xxx) {
											ClientBot = linetcr.RemoveCl(ClientBot, xxx)
										}
										kb += fmt.Sprintf("\n\n%v. %s\nMid: %s", no+1, xxx.Namebot, xxx.MID)
									}
								}
								if len(Bots) != 0 {
									linetcr.RemoveBot(Bots)
								}
								for i := range ClientBot {
									for _, x := range ClientBot[i].Squads {
										if !InArray2(Squadlist, x) {
											ClientBot[i].Squads = Remove(ClientBot[i].Squads, x)
										}
									}
								}
							}
							SaveData()
							newsend += kb
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cosquad") {
				rngcmd := GetComs(2, "cosquad")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memlist := client.GetGroupMember(to)
						yus := []string{}
						for mek, _ := range memlist {
							if InArray2(Squadlist, mek) {
								yus = append(yus, mek)
							}
						}
						if len(yus) != 0 {
							for _, mek := range yus {
								go client.SendContact(to, mek)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "msgfresh ") {
				rngcmd := GetComs(4, "msgfresh")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgfresh ", "", 1)
						MsFresh = result
						newsend += "Message fresh set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglimit ") {
				rngcmd := GetComs(4, "msglimit")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msglimit ", "", 1)
						MsLimit = result
						newsend += "Message limit set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setkick ") {
				rngcmd := GetComs(4, "setkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						anjay := strings.Split((cmd), " ")
						num, err := strconv.Atoi(anjay[1])
						if err != nil {
							newsend += "Please use number!\n"
						} else {
							MaxKick = num
							newsend += "Limiter kick set to " + anjay[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setinvite ") {
				rngcmd := GetComs(4, "setinvite")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						anjay := strings.Split((cmd), " ")
						num, err := strconv.Atoi(anjay[1])
						if err != nil {
							newsend += "Please use number!\n"
						} else {
							MaxInvite = num
							newsend += "Limiter invite set to " + anjay[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcancel ") {
				rngcmd := GetComs(4, "setcancel")
				if GetCodeprem(rngcmd, sender, to) {
					anjay := strings.Split((cmd), " ")
					num, err := strconv.Atoi(anjay[1])
					if err != nil {
						newsend += "Please use number!\n"
					} else {
						MaxCancel = num
						newsend += "Limiter cancel set to " + anjay[1] + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlimiter ") {
				rngcmd := GetComs(4, "setlimiter")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						no, err := strconv.Atoi(result[1])
						if err != nil {
							newsend += "Please use number!\n"
						} else {
							MaxKick = no
							MaxCancel = no
							MaxInvite = no
							newsend += "Limiter successs set to " + result[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setpend ") {
				rngcmd := GetComs(4, "setpend")
				if GetCodeprem(rngcmd, sender, to) {
					anjay := strings.Split((cmd), " ")
					num, err := strconv.Atoi(anjay[1])
					if err != nil {
						newsend += "Please use number!\n"
					} else {
						CancelPend = num
						newsend += "Cancel pending set to " + anjay[1] + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "fuck") && cmd != "fucklist" {
				rngcmd := GetComs(5, "fuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mute") && cmd != "mutelist" {
				rngcmd := GetComs(5, "mute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 11
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "whois") {
				rngcmd := GetComs(5, "whois")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 12
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "name") {
				rngcmd := GetComs(5, "name")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 16
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mid") {
				rngcmd := GetComs(5, "mid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 14
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmute") {
				rngcmd := GetComs(5, "unmute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmute"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "owners" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Owner) != 0 {
							list := "𝗼𝘄𝗻𝗲𝗿𝘀:\n"
							for num, xd := range UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmaster") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmaster"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Master)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungowner") {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gowner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "gojoin" {
				rngcmd := GetComs(8, "join")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m := range mem {
							if InArray2(Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "master") && cmd != "masters" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gowner") && cmd != "gowners" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcmd ") {
				rngcmd := GetComs(5, "setcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "setcmd ", "")
						texts := strings.Split(txt, " ")
						if len(texts) > 1 {
							new := Upsetcmd(texts[0], texts[1])
							if new != "" {
								newsend += new
							} else {
								newsend += "Cmd not found.\n"
							}
						} else {
							newsend += "Cmd not found.\n"
						}
					}
				}
			} else if cmd == "restartcmd" {
				rngcmd := GetComs(5, "restartcmd")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						Commands.Botname = ""
						Commands.Upallimage = ""
						Commands.Upallcover = ""
						Commands.Unsend = ""
						Commands.Upvallimage = ""
						Commands.Upvallcover = ""
						Commands.Appname = ""
						Commands.Useragent = ""
						Commands.Hostname = ""
						Commands.Friends = ""
						Commands.Adds = ""
						Commands.Limits = ""
						Commands.Addallbots = ""
						Commands.Addallsquads = ""
						Commands.Leave = ""
						Commands.Respon = ""
						Commands.Ping = ""
						Commands.Count = ""
						Commands.Limitout = ""
						Commands.Access = ""
						Commands.Allbanlist = ""
						Commands.Allgaccess = ""
						Commands.Gaccess = ""
						Commands.Checkram = ""
						Commands.Backups = ""
						Commands.Upimage = ""
						Commands.Upcover = ""
						Commands.Upvimage = ""
						Commands.Upvcover = ""
						Commands.Bringall = ""
						Commands.Purgeall = ""
						Commands.Banlist = ""
						Commands.Clearban = ""
						Commands.Stayall = ""
						Commands.Clearchat = ""
						Commands.Here = ""
						Commands.Speed = ""
						Commands.Status = ""
						Commands.Tagall = ""
						Commands.Kick = ""
						Commands.Max = ""
						Commands.None = ""
						Commands.Kickall = ""
						Commands.Cancelall = ""
						newsend += "Done restart all Cmd.\n"
					}
				}
			} else if cmd == "cleargowner" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							logAccess(client, to, sender, "cleargowner", room.Gowner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gownerlist\n", len(room.Gowner))
							room.Gowner = []string{}
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmaster" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Master) != 0 {
							newsend += fmt.Sprintf("Cleared %v masterlist\n", len(UserBot.Master))
							logAccess(client, to, sender, "clearmaster", UserBot.Master, msg.ToType)
							UserBot.ClearMaster()
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if cmd == "clearfuck" {
				rngcmd := GetComs(5, "clearfuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Fucklist) != 0 {
							logAccess(client, to, sender, "clearfuck", Banned.Fucklist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(Banned.Fucklist))
							Banned.Fucklist = []string{}
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmute" {
				rngcmd := GetComs(5, "clearmute")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Mutelist) != 0 {
							logAccess(client, to, sender, "clearmute", Banned.Mutelist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(Banned.Mutelist))
							Banned.Mutelist = []string{}
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "clearallprotect" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						linetcr.ClearProtect()
						newsend += "Cleared allprotected.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "perm ") {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						ditha := strings.ReplaceAll(cmd, "perm ", "")
						cmdLil := strings.Split(ditha, " ")
						Addpermcmd(client, to, cmdLil[0], cmdLil[1])
					}
				}
			} else if cmd == "permlist" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						list := PerCheckList()
						if list != "" {
							newsend += list
						} else {
							newsend += "Not have perm in list.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ajsjoin") {
				rngcmd := GetComs(5, "ajsjoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "ajsjoin"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							if Ajsjoin != "inv" {
								Ajsjoin = "inv"
								newsend += "Ajsjoin Invite enabled.\n"
							} else {
								newsend += "Ajsjoin Already Invite.\n"
							}
						} else if str == "qr" {
							if Ajsjoin != "qr" {
								Ajsjoin = "qr"
								newsend += "Ajsjoin qr enabled.\n"
							} else {
								newsend += "Ajsjoin Already qr.\n"
							}
						} else if str == "off" {
							if Ajsjoin != "off" {
								Ajsjoin = "off"
								newsend += fmt.Sprintf("Ajsjoin %s disabled.\n", Ajsjoin)
							} else {
								newsend += "Ajsjoin Already disabled.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autojoin") {
				rngcmd := GetComs(5, "autojoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "autojoin"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							if Autojoin != "inv" {
								Autojoin = "inv"
								newsend += "Autojoin Invite enabled.\n"
							} else {
								newsend += "Autojoin Already Invite.\n"
							}
						} else if str == "qr" {
							if Autojoin != "qr" {
								Autojoin = "qr"
								newsend += "Autojoin qr enabled.\n"
							} else {
								newsend += "Autojoin Already qr.\n"
							}
						} else if str == "off" {
							if Autojoin != "off" {
								Autojoin = "off"
								newsend += fmt.Sprintf("Autojoin %s disabled.\n", Autojoin)
							} else {
								newsend += "Autojoin Already disabled.\n"
							}
						}
					}
				}
			} else if cmd == "mutelist" {
				rngcmd := GetComs(5, "mutelist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Mutelist) != 0 {
							list := "Mutelist:"
							client.SendPollMention(to, list, Banned.Mutelist)
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "fucklist" {
				rngcmd := GetComs(5, "fucklist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Fucklist) != 0 {
							list := "Fucklist:"
							client.SendPollMention(to, list, Banned.Fucklist)
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "groupcast") {
				rngcmd := GetComs(5, "upcast")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "groupcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							gr, _ := client.GetGroupIdsJoined()
							for _, gi := range gr {
								client.SendMessage(gi, fancy(str))
								go client.SendContact(gi, sender)
							}
							newsend += "Success broadcast to " + strconv.Itoa(len(gr)) + " group\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "friendcast") {
				rngcmd := GetComs(5, "ndcast")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "friendcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							n := 0
							for _, cl := range ClientBot {
								fl, _ := cl.GetAllContactIds()
								for _, x := range fl {
									client.SendMessage(x, fancy(str))
									go client.SendContact(x, sender)
									time.Sleep(3 * time.Second)
								}
							}
							n += 1
							newsend += "Success broadcast to all friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bcimage") {
				rngcmd := GetComs(5, "bcimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "bcimage"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							addConSingle(result)
							for _, gi := range result {
								list := " "+result[2]
								client.SendMessage(gi, fancy(list))
								image := "https://"+result[1]
								
								client.SendFoto(gi, image)
								if !InArray2(MidBc, gi) {
									MidBc = append(MidBc, gi)
								}
							}
							newsend += "Success broadcast image to friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bcast") {
				rngcmd := GetComs(5, "bcast")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "bcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							addConSingle(result)
							for _, gi := range result {
								list := " "+result[1]
								client.SendMessage(gi, fancy(list))
								go client.SendContact(gi, sender)
								if !InArray2(MidBc, gi) {
									MidBc = append(MidBc, gi)
								}
							}
							newsend += "Success broadcast to friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "list broadcast") {
				rngcmd := GetComs(5, "list broadcast")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(MidBc) != 0 {
							list :="List Broadcast:"
							cuh, _ := client.GetContacts(MidBc)
							for num, prs := range cuh {
								num++
								rengs := strconv.Itoa(num)
								name := prs.DisplayName
								list += fmt.Sprintf("\n ."+rengs+" %v", name)
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "accgroupid") {
				rngcmd := GetComs(5, "accgroupid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "accgroupid"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							client.AcceptGroupInvitationNormal(str)
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									err := p.AcceptGroupInvitationNormal(str)
									if err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								newsend += "Succes Accept Group ID"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unicode") {
				rngcmd := GetComs(5, "unicode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "unicode"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							addConSingle(result)
							no, _ := strconv.Atoi(result[1])
							for _, gi := range result {
								for i := 1; i <= no; i++ {
									time.Sleep(time.Second * 3)
									for _, p := range bk {
										filepath := fmt.Sprintf("unicode.txt")
										b, err := ioutil.ReadFile(filepath)
										if err != nil {
											fmt.Print(err)
										}
										code := string(b)
										list := code
										p.SendMessage(gi, fancy(list))
									}
								}
							}
							newsend += "Success send unicode\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gcreate") {
				rngcmd := GetComs(5, "gcreate")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "gcreate"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							addConSingle(result)
							no, _ := strconv.Atoi(result[1])
							for i := 1; i <= no; i++ {
								for _, p := range bk {
									call.CreateGroup("test", result, string(p.AuthToken))
								}
							}
							newsend += "Success create new group\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gourl ") {
				rngcmd := GetComs(5, "gourl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						num, err := strconv.Atoi(result[1])
						if err != nil {
							newsend += "invalid number.\n"
						} else {
							gr := []string{}
							for i := range ClientBot {
								grs, _ := ClientBot[i].GetGroupIdsJoined()
								for _, a := range grs {
									if !InArray2(gr, a) {
										gr = append(gr, a)
									}
								}
							}
							groups, _ := client.GetGroups(gr)
							tempgroup := []string{}
							for _, gi := range groups {
								tempgroup = append(tempgroup, gi.ChatMid)
							}
							if num > 0 && num <= len(tempgroup) {
								gid := tempgroup[num-1]
								tick, err := client.ReissueChatTicket(gid)
								if err == nil {
									var err error
									mes := make(chan bool)
									go func() {
										err = client.UpdateChatQrV2(gid, false)
										if err != nil {
											mes <- false
										} else {
											mes <- true
										}
									}()

									newsend += "https://line.me/R/ti/g/" + tick + "\n"
								}
							} else {
								newsend += "out of range.\n"
							}
						}
					}
				}
			} else if cmd == "allgroups" {
				rngcmd := GetComs(4, "allgroups")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							nm := []string{}
							gr, _ := p.GetGroupIdsJoined()
							for c, a := range gr {
								name, _ := p.GetGroupMember(a)
								c += 1
								name = fmt.Sprintf("%v. %s", c, name)
								nm = append(nm, name)
								GetSquad(p, a)
							}
							
							nm1 := []string{}
							gr1, _ := p.GetGroupIdsInvited()
							for c1, a1 := range gr1 {
								name1, _ := p.GetGroupMember(a1)
								c1 += 1
								name1 = fmt.Sprintf("%v. %s", c1, name1)
								nm1 = append(nm1, name1)
								GetSquad(p, a1)
							}
							stf := "Group list:\n\n"
							str := strings.Join(nm, "\n\n")
							stf1 := "\n\nPending list:\n\n"
							str1 := strings.Join(nm1, "\n")
							p.SendText(to, stf+str+stf1+str1)
						}
					}
				}
			} else if cmd == "groups" {
				rngcmd := GetComs(5, "groups")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						gr := []string{}
						for _, p := range ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsJoined()
								for _, a := range grs {
									if !InArray2(gr, a) {
										gr = append(gr, a)
									}
								}
							}
						}
						nm := []string{}
						grup, _ := client.GetGroups(gr)
						ci := []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups := []*talkservice.Chat{}
						tempgroup = []string{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !InArrayChat(groups, gi) {
										groups = append(groups, gi)
										tempgroup = append(tempgroup, gi.ChatMid)
									}
								}
							}
						}
						for c, a := range groups {
							name, mem := a.ChatName, a.Extra.GroupExtra.MemberMids
							c += 1
							jm := 0
							for mid := range mem {
								if InArray2(Squadlist, mid) {
									jm++
								}
							}
							name = fmt.Sprintf("%v. %s (%v/%v)", c, name, jm, len(mem))
							nm = append(nm, name)
							GetSquad(client, a.ChatMid)
						}
						stf := "All Group List:\n\n"
						str := strings.Join(nm, "\n")
						anu := []string{}
						for _, p := range ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsInvited()
								for _, a := range grs {
									if !InArray2(gr, a) && !InArray2(anu, a) {
										anu = append(anu, a)
									}
								}
							}
						}
						grup, _ = client.GetGroups(anu)
						ci = []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups = []*talkservice.Chat{}
						tempginv = []string{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !InArrayChat(groups, gi) {
										groups = append(groups, gi)
										tempginv = append(tempginv, gi.ChatMid)
									}
								}
							}
						}
						nm = []string{}
						nn := 1
						for _, a := range groups {
							name, mem, inv := a.ChatName, a.Extra.GroupExtra.MemberMids, a.Extra.GroupExtra.InviteeMids
							if name != "" {
								jm := 0
								for mid := range inv {
									if InArray2(Squadlist, mid) {
										jm++
									}
								}
								if jm != 0 {
									name = fmt.Sprintf("%v. %s (invited) (%v/%v)", nn, name, jm, len(mem))
									nm = append(nm, name)
									GetSquad(client, a.ChatMid)
									nn++
								} else {
									tempginv = Remove(tempginv, a.ChatMid)
								}
							} else {
								tempginv = Remove(tempginv, a.ChatMid)
							}
						}
						var strs, strsa = "", ""
						if len(nm) != 0 {
							strs = "\n\nAll Group Invitation:\n\n"
							strsa = strings.Join(nm, "\n")
						}
						newsend += stf + str + strs + strsa
					}
				}
			} else if strings.HasPrefix(cmd, "nukejoin ") {
				rngcmd := GetComs(5, "nukejoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "nukejoin ", "", 1)
						if spl == "on" {
							NukeJoin = true
							newsend += "Nukejoin is enabled.\n"
						} else if spl == "off" {
							NukeJoin = false
							newsend += "Nukejoin is disabled.\n"
						}
					}
				}

			} else if strings.HasPrefix(cmd, "canceljoin ") {
				rngcmd := GetComs(5, "canceljoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "canceljoin ", "", 1)
						if spl == "on" {
							canceljoin = true
							newsend += "canceljoin is enabled.\n"
						} else if spl == "off" {
							canceljoin = false
							newsend += "canceljoin is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopro ") {
				rngcmd := GetComs(5, "autopro")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopro ", "", 1)
						if spl == "on" {
							AutoPro = true
							newsend += "Autopro is enabled.\n"
						} else if spl == "off" {
							AutoPro = false
							newsend += "Autopro is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "detectcall ") {
				rngcmd := GetComs(5, "detectcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "detectcall ", "", 1)
						if spl == "on" {
							DetectCall = true
							newsend += "Detectgroupcall is enabled.\n"
						} else if spl == "off" {
							DetectCall = false
							newsend += "Detectgroupcall is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kickbanqr ") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "kickbanqr ", "", 1)
						if spl == "on" {
							KickBanQr = true
							newsend += "KickbanQr is enabled.\n"
						} else if spl == "off" {
							KickBanQr = false
							newsend += "KickbanQr is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autobackbot ") {
				rngcmd := GetComs(5, "autobackbot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autobackbot ", "", 1)
						if spl == "on" {
							AutoBackBot = true
							newsend += "Autobackbot is enabled.\n"
						} else if spl == "off" {
							AutoBackBot = false
							newsend += "Autobackbot is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopurge ") {
				rngcmd := GetComs(5, "autopurge")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopurge ", "", 1)
						if spl == "on" {
							AutoPurge = true
							newsend += "Autopurge is enabled.\n"
						} else if spl == "off" {
							AutoPurge = false
							newsend += "Autopurge is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autoticket") {
				rngcmd := GetComs(5, "autoticket")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autoticket"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "normal" {
							AutoJointicket = true
							TypeJoin = str
							count = count + 1
							client.SendMessage(to, fancy("AutoJointicket type normal"))
						} else if str == "nuke" {
							AutoJointicket = true
							TypeJoin = str
							count = count + 1
							client.SendMessage(to, fancy("AutoJointicket type nuke"))
						} else if str == "off" {
							AutoJointicket = false
							TypeJoin = "none"
							count = count + 1
							client.SendMessage(to, fancy("AutoJointicket is disabled"))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("Autobroadcast type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autotrans") {
				rngcmd := GetComs(5, "autotrans")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autotrans"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "off" {
							AutoTranslate = false
							TypeTrans = "none"
							count = count + 1
							client.SendMessage(to, fancy("AutoTranslate is disabled"))
						} else if len(str) != 0 {
							AutoTranslate = true
							TypeTrans = str
							count = count + 1
							client.SendMessage(to, fancy("AutoTranslate type "+str))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("AutoTranslate type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "flex ") {
				rngcmd := GetComs(5, "flex")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "flex ", "", 1)
						if spl == "1" {
							linetcr.FlexMode = true
							linetcr.FooterMode = false
							linetcr.FlexMode2 = false
							newsend += "Flexmode 1 is enabled.\n"
						} else if spl == "2" {
							linetcr.FlexMode2 = true
							linetcr.FooterMode = false
							linetcr.FlexMode = false
							newsend += "Flexmode 2 is enabled.\n"
						} else if spl == "footer" {
							linetcr.FooterMode = true
							linetcr.FlexMode2 = false
							linetcr.FlexMode = false
							newsend += "Footermode is enabled.\n"
						} else if spl == "off" {
							linetcr.FooterMode = false
							linetcr.FlexMode = false
							linetcr.FlexMode2 = false
							newsend += "Flexmode is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "liff ") {
				rngcmd := GetComs(5, "liff")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "liff ", "", 1)
						if spl == "1" {
							linetcr.Liffid = "1655425084-3OQ8Mn9J"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "2" {
							linetcr.Liffid = "1656652460-LGKR2XXJ"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "3" {
							linetcr.Liffid = "1655623470-81eDd9kM"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "4" {
							linetcr.Liffid = "1653779160-yw2l2v9d"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "5" {
							linetcr.Liffid = "1657707255-WVxqmM35"
							newsend += "Liff update:\n  "+linetcr.Liffid
						}
					}
				}
			} else if strings.HasPrefix(cmd, "lockajs ") {
				rngcmd := GetComs(5, "lockajs")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "lockajs ", "", 1)
						if spl == "on" {
							LockAjs = true
							newsend += "Lockajs is enabled.\n"
						} else if spl == "off" {
							LockAjs = false
							newsend += "Lockajs is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "modepro ") {
				rngcmd := GetComs(5, "modepro")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "modepro ", "", 1)
						if spl == "on" {
							ProtectMode = true
							newsend += "Mode protect is enabled.\n"
						} else if spl == "off" {
							ProtectMode = false
							newsend += "Mode protect is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autoban ") {
				rngcmd := GetComs(5, "autoban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autoban ", "", 1)
						if spl == "on" {
							AutoBan = true
							newsend += "Autoban is enabled.\n"
						} else if spl == "off" {
							AutoBan = false
							newsend += "Autoban is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "lockmode ") {
				rngcmd := GetComs(5, "lockmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "lockmode ", "", 1)
						if spl == "on" {
							LockMode = true
							newsend += "Lockmode is enabled.\n"
						} else if spl == "off" {
							LockMode = false
							newsend += "Lockmode is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "modewar ") {
				rngcmd := GetComs(5, "modewar")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "modewar ", "", 1)
						if spl == "on" {
							PowerMode = true
							newsend += "Mode war is enabled.\n"
						} else if spl == "off" {
							PowerMode = false
							newsend += "Mode war is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "banpurge ") {
				rngcmd := GetComs(5, "banpurge")
				if GetCodeprem(rngcmd, sender, to) {
					spl := strings.Replace(cmd, "banpurge ", "", 1)
					if spl == "on" {
						AutokickBan = true
						newsend += "Banpurge is enabled.\n"
					} else if spl == "off" {
						AutokickBan = false
						newsend += "Banpurge is disabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "groupinfo ") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										list := InfoGroup(client, gid)
										client.SendMessage(to, fancy(list))
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotegroup ") {
				rngcmd := GetComs(4, "remotegroup")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										remotegrupidto = to
										if !InArray2(Sinderremote, sender) {
											Sinderremote = append(Sinderremote, sender)
										}
										names, _, _ := client.GetChatList(gid)
										remotegrupid = tempgroup[nim]
										GetSquad(client, gid)
										ret := fmt.Sprintf("Group: %v\n\n Send your command.\n", names)
										newsend += ret
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteowner ") {
				rngcmd := GetComs(4, "remoteowner")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !InArray2(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !InArray2(MidRemote, mid) {
													MidRemote = append(MidRemote, mid)
												}
											}
											client.SendMessage(to, fancy(list))
										}
										RemoteOwner = true
										client.SendMessage(to, fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotemaster ") {
				rngcmd := GetComs(4, "remotemaster")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !InArray2(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !InArray2(MidRemote, mid) {
													MidRemote = append(MidRemote, mid)
												}
											}
											client.SendMessage(to, fancy(list))
										}
										RemoteMaster = true
										client.SendMessage(to, fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteadmin ") {
				rngcmd := GetComs(4, "remoteadmin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !InArray2(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !InArray2(MidRemote, mid) {
													MidRemote = append(MidRemote, mid)
												}
											}
											client.SendMessage(to, fancy(list))
										}
										RemoteAdmin = true
										client.SendMessage(to, fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotecontact ") {
				rngcmd := GetComs(4, "remotecontact")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !InArray2(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !InArray2(MidRemote, mid) {
													MidRemote = append(MidRemote, mid)
												}
											}
											client.SendMessage(to, fancy(list))
										}
										RemoteContact = true
										client.SendMessage(to, fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteban ") {
				rngcmd := GetComs(4, "remoteban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, err := strconv.Atoi(result[1])
							if err != nil {
								client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(tempgroup) == 0 {
										client.SendMessage(to, fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(tempgroup)+1 {
										gid := tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !InArray2(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !InArray2(MidRemote, mid) {
													MidRemote = append(MidRemote, mid)
												}
											}
											client.SendMessage(to, fancy(list))
										}
										RemoteBan = true
										client.SendMessage(to, fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unfuck") {
				rngcmd := GetComs(4, "unfuck")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unfuck"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "admin") && cmd != "admins" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "gowners" {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							list := "𝗴𝗼𝘄𝗻𝗲𝗿𝘀:\n"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "masters" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Master) != 0 {
							list := " 𝗺𝗮𝘀𝘁𝗲𝗿𝘀:\n"
							for num, xd := range UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unadmin") {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Admin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungadmin") {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("??𝗹𝗲????𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gadmin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "squadmid" {
				rngcmd := GetComs(5, "squadmid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						e, _ := client.GetProfile()
						list := "All bot's mid\n\n"
						list += "1." + e.DisplayName + "\n\n"
						list += client.MID
						for b, a := range client.Squads {
							b++
							x, _ := client.GetContact(a)
							list += fmt.Sprintf("\n\n%v. %s ", b+1, x.DisplayName)
							list += "\n\n" + a
						}
						newsend += list + "\n"
					}
				}
			} else if cmd == "check all" {
				rngcmd := GetComs(5, "check all")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						e, _ := client.GetProfile()
						list := "All bot's: ↓\n\n"
						list += "1." + e.DisplayName + "\nMid "
						list += client.MID
						for b, a := range client.Squads {
							b++
							x, _ := client.GetContact(a)
							list += fmt.Sprintf("\n\n%v. %s ", b+1, x.DisplayName)
							list += "\nMid " + a
						}
						list += "\n\nBot's limited: ↓"
						for n, cl := range linetcr.KickBans {
							m := cl.MID
							no := n + 1
							pr, _ := client.GetContact(m)
							cl.Namebot = pr.DisplayName
							list += fmt.Sprintf("\n\n%v. %s\nMid %v", no, cl.Namebot, cl.MID)
						}
						newsend += list + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "checkban") && cmd != "checkbans" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							newlist := []string{}
							for _, user := range lists {
								if !InArray2(newlist, user) {
									newlist = append(newlist, user)
								}
							}
							numr := 0
							tex := "Status Account:\n"
							for _, user := range newlist {
								name := client.GetName(user)
								if name == "" {
									name = "Deleted Account"
								}
								numr ++
								tex += fmt.Sprintf("\n%v. %v", numr, name)
								time.Sleep(100 * time.Millisecond)
								r, _ := client.GetHomeProfile(user)
								if linetcr.GetBannedChat(r) == 1 {
									tex += fmt.Sprintf("\n   Status: Banned\n")
								} else {
									tex += fmt.Sprintf("\n   Status: Normal\n")
								}
							}
							newsend += tex
						}
					}
				}
			} else if strings.HasPrefix(cmd, "wordbanadd") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "wordbanadd"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if !InArray2(Data.WordbanBack, str) {
							Data.WordbanBack = append(Data.WordbanBack, str)
							SaveBackup()
							newsend += "Wordban added : " + str + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "wordbandel") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "wordbandel"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if InArray2(Data.WordbanBack, str) {
							Data.WordbanBack = Remove(Data.WordbanBack, str)

							client.SendMessage(to, fancy("Wordban delete : "+str))
						}
						SaveBackup()
					} else {
						client.SendMessage(to, fancy("No wordban deleted"))
					}
				}
			} else if strings.HasPrefix(cmd, "wordbanlist") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len (Data.WordbanBack) != 0 {
							list :="Wordban List: ↓"
							for num, xd := range Data.WordbanBack {
								num++
								rengs := strconv.Itoa(num)
								list += "\n\n  "+rengs+". "+xd
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if cmd == "wordbanclear" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Data.WordbanBack) != 0 {
							newsend += fmt.Sprintf("Cleared %v wordbanlist", len(Data.WordbanBack)) + "\n"
							Data.WordbanBack = []string{}
							SaveBackup()
						} else {
							newsend += "Wordban is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gadmin") && cmd != "gadmins" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "admins" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Admin) != 0 {
							list := "𝗮𝗱𝗺𝗶𝗻𝘀:\n"
							for num, xd := range UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "antitag ") {
				rngcmd := GetComs(5, "antitag")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "antitag ", "", 1)
						if spl == "on" {
							room.AntiTag = true
							newsend += "antitag enabled.\n"
						} else if spl == "off" {
							room.AntiTag = false
							newsend += "antitag disabled.\n"
						}
					}
				}
			} else if cmd == "banlist" || cmd == Commands.Banlist && Commands.Banlist != "" {
				rngcmd := GetComs(7, "banlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Banlist) != 0 {
							listbl := "banlist:"
							client.SendPollMention(to, listbl, Banned.Banlist)
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "locklist" {
				rngcmd := GetComs(5, "locklist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Locklist) != 0 {
							listbl := "locklist:"
							client.SendPollMention(to, listbl, Banned.Locklist)
						} else {
							newsend += "Lock list is empty.\n"
						}
					}
				}
			} else if cmd == "respon" || cmd == Commands.Respon && Commands.Respon != "" {
				rngcmd := GetComs(5, "respon")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							go p.SendMessage(to, fancy(MsgRespon))
						}
					}
				}
			} else if cmd == "rollcall" || cmd == Commands.Botname && Commands.Botname != "" {
				rngcmd := GetComs(5, "rollcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							asss := fmt.Sprintf("%v", p.Namebot)
							go p.SendMessage(to, fancy(asss))
						}
					}
				}
			} else if cmd == "upallimage" || cmd == Commands.Upallimage && Commands.Upallimage != "" {
				rngcmd := GetComs(4, "upallimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*linetcr.Account{}
						for _, p := range bk {
							if !linetcr.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangPict = true
						StartChangeImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, fancy("Send image."))
					}
				}
			} else if cmd == "upallcover" || cmd == Commands.Upallcover && Commands.Upallcover != "" {
				rngcmd := GetComs(4, "upallcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*linetcr.Account{}
						for _, p := range bk {
							if !linetcr.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangCover = true
						StartChangeImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, fancy("Send image."))
					}
				}
			} else if cmd == "unsend" || cmd == Commands.Unsend && Commands.Unsend != "" {
				rngcmd := GetComs(7, "unsend")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.UnsendChat(to)
						}
					}
				}
			} else if cmd == "upvallimage" || cmd == Commands.Upvallimage && Commands.Upvallimage != "" {
				rngcmd := GetComs(4, "upvallimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !linetcr.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangVpict = true
						StartChangevImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, fancy("Send video."))
					}
				}
			} else if cmd == "upvallcover" || cmd == Commands.Upvallcover && Commands.Upvallcover != "" {
				rngcmd := GetComs(4, "upvallcover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						changepic = []*linetcr.Account{}
						for _, p := range bk {
							if !linetcr.Checkarri(changepic, p) {
								changepic = append(changepic, p)
							}
						}
						ChangVcover = true
						StartChangevImg = true
						AllCheng = true
						timeabort = time.Now()
						client.SendMessage(to, fancy("Send video."))
					}
				}
			} else if cmd == "appname" || cmd == Commands.Appname && Commands.Appname != "" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy(string(p.AppName)))
						}
					}
				}
			} else if cmd == "cektoken" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy(string(p.AuthToken)))
						}
					}
				}
			} else if cmd == "useragent" || cmd == Commands.Useragent && Commands.Useragent != "" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy(string(p.UserAgent)))
						}
					}
				}
			} else if cmd == "hostname" || cmd == Commands.Hostname && Commands.Hostname != "" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy(string(p.Host)))
						}
					}
				}
			} else if cmd == "location" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy(string(p.Locale)))
						}
					}
				}
			} else if cmd == "friends" || cmd == Commands.Friends && Commands.Friends != "" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						exe2 := []*linetcr.Account{}
						for _, mid := range mentionlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								exe2 = append(exe2, cl)
							}
						}
						if len(exe2) != 0 {
							for _, p := range exe2 {
								friends, _ := p.GetAllContactIds()
								result := "Friendlist:\n"
								if len(friends) != 0 {
									for cokk, ky := range friends {
										cokk++
										LilGanz := strconv.Itoa(cokk)
										haniku, _ := p.GetContact(ky)
										result += "\n" + LilGanz + ". " + haniku.DisplayName
									}
									client.SendMessage(to, fancy(result))
								} else {
									client.SendMessage(to, fancy("Friend is empty."))
								}
							}
						} else {
							client.SendMessage(to, fancy("Mention Bot First."))
						}
					}
				}
			} else if cmd == "adds" || cmd == Commands.Adds && Commands.Adds != "" {
				if GetCodeprem(1, sender, to) {
					if CheckExprd(client, to, sender) {
						addb := len(linetcr.Waitadd)
						kb := ""
						if addb != 0 {
							kb += fmt.Sprintf("%v/%v bot's got add/friend banned.", addb, len(Squadlist))
							for n, cl := range linetcr.Waitadd {
								m := cl.MID
								no := n + 1
								go client.SendContact(to, m)
								var ta time.Duration
								if _, ok := linetcr.BlockAdd.Get(cl.MID); ok {
									t := cl.Timeadd.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.Timeadd.Add(1*time.Hour )
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, fmtDurations(ta))
							}
						}
						if addb == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}
				}
			} else if cmd == "cek" || cmd == Commands.Limits && Commands.Limits != "" {
				rngcmd := GetComs(5, "limits")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							var asss string
							if p.Limited == true {
								asss += MsLimit
							} else {
								asss += MsFresh
							}
							p.SendMessage(to, fancy(asss))
						}
					}
				}
			} else if cmd == "addallsquads" || cmd == Commands.Addallsquads && Commands.Addallsquads != "" {
				rngcmd := GetComs(2, "addallsquads")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, "Waiting for Add Squad")
						addConSqV2(Squadlist)
						friends, _ := client.GetAllContactIds()
						result := "Success added friends:\n"
						if len(friends) != 0 {
							for cokk, ky := range friends {
								cokk++
								LilGanz := strconv.Itoa(cokk)
								haniku, _ := client.GetContact(ky)
								result += "\n" + LilGanz + ". " + haniku.DisplayName
							}
							client.SendMessage(to, fancy(result))
						}
					}
				}
			} else if cmd == "unfriendbans" {
				if GetCodeprem(2, sender, to) {
					if CheckExprd(client, to, sender) {
						RemBanFriends(client, to)
					}
				}
			} else if cmd == "clearfriends" {
				rngcmd := GetComs(2, "clear friends")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, "Waiting for Clear Friend")
						clearCon()
						asss := "Success clear allfriends."
						client.SendMessage(to, fancy(asss))
					}
				}
			} else if cmd == "leave" || cmd == Commands.Leave && Commands.Leave != "" {
				rngcmd := GetComs(7, "leave")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m := range mem {
							if InArray2(Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						GetSquad(client, to)
						room := linetcr.GetRoom(to)
						bk = room.Client
						for _, cl := range bk {
							go cl.LeaveGroup(to)
						}
						if LogGroup == to {
							LogMode = false
							LogGroup = ""
						}
						linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
						SaveBackup()
						logAccess(client, to, sender, "leave", []string{}, msg.ToType)
					}
				}
			} else if cmd == "ping" || cmd == Commands.Ping && Commands.Ping != "" {
				rngcmd := GetComs(5, "ping")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, fancy("pong"))
						}
					}
				}
			} else if cmd == "count" || cmd == Commands.Count && Commands.Count != "" {
				rngcmd := GetComs(5, "count")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for n, p := range bk {
							p.SendMessage(to, fancy(fmt.Sprintf("%v", n+1)))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "sayall") {
				rngcmd := GetComs(5, "sayall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := ""
						var su = "sayall"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						for _, p := range bk {
							p.SendMessage(to, fancy(str))
						}
					}
				}
			} else if cmd == "limitout" || cmd == Commands.Limitout && Commands.Limitout != "" {
				rngcmd := GetComs(8, "out")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							if p.Limited == true {
								p.LeaveGroup(to)
							}
						}
						GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "upallstatus") {
				rngcmd := GetComs(4, "upallstatus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							str := ""
							var su = "upallstatus"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							for num, p := range bk {
								num++
								rengs := strconv.Itoa(num)
								name := str+" "+rengs
								if TimeDown(num) {
									p.UpdateProfileBio(name)
									p.SendMessage(to, fancy("Profile Bio updated."))
								}
							}
						} else {
							client.SendMessage(to, fancy("Add Bio first."))
						}
						timeabort = time.Now()
					}
				}
			} else if strings.HasPrefix(cmd, "upallname") {
				rngcmd := GetComs(4, "upallname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							var str string
							var su = "upallname"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							aa := utf8.RuneCountInString(str)
							if aa != 0 && aa <= 20 {
								for num, p := range bk {
									num++
									rengs := strconv.Itoa(num)
									name := str+" "+rengs
									if TimeDown(num) {
										p.UpdateProfileName(name)
										p.SendMessage(to, fancy("Profile name success updated."))
									}
								}
							}
						} else {
							client.SendMessage(to, fancy("Add name first."))
						}
					}
				}
			} else if cmd == "autoname" {
				rngcmd := GetComs(4, "autoname")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						seed := time.Now().UTC().UnixNano()
						nameGenerator := namegenerator.NewNameGenerator(seed)
						for i,x := range Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								err := ClientBot[i].UpdateProfileName(nameGenerator.Generate())
								time.Sleep(1 * time.Second)
								if err != nil {
									fmt.Println(err)
									ClientBot[i].SendMessage(to, fancy("Update profile name failure."))
								} else {
									ClientBot[i].SendMessage(to, fancy("Update profile name done."))
								}
							}
						}
					}
				}
			} else if cmd == "randomprofile" {
				rngcmd := GetComs(4, "randomprofile")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						seed := time.Now().UTC().UnixNano()
						nameGenerator := namegenerator.NewNameGenerator(seed)
						for i,x := range Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								ClientBot[i].UpdateProfileName(nameGenerator.Generate())
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								link1 := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								ClientBot[i].UpdateCoverWithURL(link1)
								err := ClientBot[i].UpdateProfilePictureWithURL(link, "p")
								if err != nil {
									fmt.Println(err)
									ClientBot[i].SendMessage(to, fancy("Update random profile failure."))
								} else {
									ClientBot[i].SendMessage(to, fancy("Update random profile done."))
								}
							}
						}
					}
				}
			} else if cmd == "1autoimage" {
				rngcmd := GetComs(4, "1autoimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for i,x := range Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								err := ClientBot[i].UpdateProfilePictureWithURL(link, "p")
								if err != nil {
									fmt.Println(err)
									ClientBot[i].SendMessage(to, fancy("Update picture failure."))
								} else {
									ClientBot[i].SendMessage(to, fancy("Update picture done."))
								}
							}
						}
					}
				}
			} else if cmd == "autocover" {
				rngcmd := GetComs(4, "autocover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for i,x := range Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								err := ClientBot[i].UpdateCoverWithURL(link)
								if err != nil {
									fmt.Println(err)
									ClientBot[i].SendMessage(to, fancy("Update cover failure."))
								} else {
									ClientBot[i].SendMessage(to, fancy("Update cover done."))
								}
							}
						}
					}
				}
			} else if cmd == "1autoname" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						for n, p := range bk {
							if TimeDown(n) {
								p.AutoupdateName(p.AuthToken)
								p.SendMessage(to, fancy("Success update name"))
							}
						}
					}
				}
			} else if cmd == "autoimage" {
				if GetCodeprem(4, sender, to) {
					if CheckExprd(client, to, sender) {
						for n, p := range bk {
							if TimeDown(n) {
								p.AutoupdatePict(p.AuthToken)
								p.SendMessage(to, fancy("Success update pict"))
							}
						}
					}
				}
			} else if cmd == "clearadmin" {
				if GetCodeprem(6, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Admin) != 0 {
							logAccess(client, to, sender, "clearadmin", UserBot.Admin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v adminlist\n", len(UserBot.Admin))
							UserBot.ClearAdmin()
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if cmd == "clearban" || cmd == Commands.Clearban && Commands.Clearban != "" {
				rngcmd := GetComs(7, "rban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Banlist) != 0 {
							msgcbn := fmt.Sprintf(MsgBan, len(Banned.Banlist))
							logAccess(client, to, sender, "clearban", Banned.Banlist, msg.ToType)
							newsend += msgcbn + "\n"
							Banned.Banlist = []string{}
							Banned.Exlist = []string{}
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "clearlock" {
				rngcmd := GetComs(5, "clearlock")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(Banned.Locklist) != 0 {
							msgcbn := fmt.Sprintf(MsgLock, len(Banned.Locklist))
							logAccess(client, to, sender, "clearlock", Banned.Locklist, msg.ToType)
							newsend += msgcbn + "\n"
							Banned.Locklist = []string{}
							Banned.Exlist = []string{}
						} else {
							newsend += "Lock list is empty.\n"
						}
					}
				}
			} else if cmd == "cleargadmin" {
				if GetCodeprem(8, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							logAccess(client, to, sender, "cleargadmin", room.Gadmin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gadminlist\n", len(room.Gadmin))
							room.Gadmin = []string{}
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if cmd == "/list protect" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						res := linetcr.ListProtect()
						client.SendHelp(to, fancy(res + "\n"))
					}
				}
			} else if cmd == "list protect" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "✠ List Protect:"
						fukk, _ := client.GetGroupIdsJoined()
						for num, group := range fukk {
							num++
							rengs := strconv.Itoa(num)
							Room := linetcr.GetRoom(group)
							ret += fmt.Sprintf("\n\n%v. %s\n", rengs, Room.Name)
							if Room.ProQr {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							} else {
								ret += "\n ⚪ » Protect QR"
							}
							if Room.AntiTag {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							}
							if Room.ProKick {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							}
							if Room.ProInvite {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							}
							if Room.ProCancel {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							}
							if Room.ProJoin {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							}
							if Room.ProName {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							}
							if Room.ProPicture {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							}
							if Room.ProNote {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							}
							if Room.ProAlbum {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							}
							if Room.ProLink {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							}
							if Room.ProFlex {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							}
							if Room.ProImage {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							}
							if Room.ProVideo {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							}
							if Room.ProCall {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							}
							if Room.ProSpam {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							}
							if Room.ProSticker {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							}
							if Room.ProContact {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							}
							if Room.ProPost {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							}
							if Room.ProFile {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							}
							if len(Room.GoMid) > 0 {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							}
						}
						client.SendHelp(to, fancy(ret + "\n"))
					}
				}
			} else if cmd == "bringall" || cmd == Commands.Bringall && Commands.Bringall != "" {
				rngcmd := GetComs(5, "bringall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if op.Message.ToType != 2 {
							return
						}
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						exe, _ := SelectBot(client, to)
						if exe != nil {
							Setinviteto(exe, to, exe.Squads)
							logAccess(client, to, sender, "invite", exe.Squads, msg.ToType)
							time.Sleep(1 * time.Second)
							GetSquad(exe, to)
						} else {
							newsend += "Invite banned try with another bot.\n"
						}
					}
				}
			} else if cmd == "stayall" || cmd == Commands.Stayall && Commands.Stayall != "" {
				rngcmd := GetComs(7, "stayall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						//_, mem := client.GetGroupInvitation(to)
						//anu := []string{}
						//for m := range mem {if InArray2(Squadlist, m) {anu = append(anu, m)}
						//}
						//if len(anu) != 0 {for _, mid := range anu {cl := GetKorban(mid);cl.AcceptGroupInvitationNormal(to)}
						//}
						//GetSquad(client, to)
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						numb := len(ClientBot)
						if numb > 0 && numb <= len(ClientBot) {
							GetSquad(client, to)
							room := linetcr.GetRoom(to)
							aa := len(room.Client)
							if aa > numb {
								c := aa - numb
								ca := 0
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for _, o := range list {
									o.LeaveGroup(to)
									ca = ca + 1
									if ca == c {
										break
									}
								}
								GetSquad(client, to)
							} else if aa < numb {
								ti, err := client.ReissueChatTicket(to)
								if err == nil {
									go client.UpdateChatQrV2(to, false)
									all := []*linetcr.Account{}
									room := linetcr.GetRoom(to)
									cuk := room.Client
									for _, x := range ClientBot {
										if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
											all = append(all, x)
										}
									}
									sort.Slice(all, func(i, j int) bool {
										return all[i].KickPoint < all[j].KickPoint
									})
									g := numb - aa
									var wg sync.WaitGroup
									wi := GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
										l := all[i]
										if l != client && !linetcr.InArrayCl(wi, l) {
											if !l.Limited {
												wg.Add(1)
												go func() {
													l.AcceptTicket(to, ti)
													wg.Done()
												}()
											}
										}
									}
									wg.Wait()
									client.UpdateChatQrV2(to, true)
									GetSquad(client, to)
									logAccess(client, to, sender, "bringbot", []string{}, 2)
									SaveBackup()
									aa := len(room.Client)
									var name string
									name = fmt.Sprintf("Ready %v bots here", aa)
									newsend += name + "\n"
								}
							}
						}
						if LockAjs {
							str := CountAjs
							numb, _ := strconv.Atoi(str)
							if numb == 0 {
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for n, o := range list {
									if n < 2 {
										o.LeaveGroup(to)
										linetcr.GetRoom(to).RevertGo(o)

									} else {
										break
									}
								}
								room := linetcr.GetRoom(to)
								cls := room.Client
								for _, cl := range cls {
									if !cl.Limited {
										for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
										}
										break
									}
								}
							} else {
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for n, o := range list {
									if n < numb {
										o.LeaveGroup(to)
										linetcr.GetRoom(to).RevertGo(o)
									} else {
										break
									}
								}
								room := linetcr.GetRoom(to)
								cls := room.Cans()
								for _, cl := range cls {
									if !cl.Limited {
										for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
										}
										break
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "go") && cmd != "gojoin" {
				rngcmd := GetComs(8, "o")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						str := strings.Replace(cmd, "go ", "", 1)
						numb, _ := strconv.Atoi(str)
						if numb == 0 {
							list := append([]*linetcr.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < 2 {
									o.LeaveGroup(to)
									linetcr.GetRoom(to).RevertGo(o)

								} else {
									break
								}
							}
							room := linetcr.GetRoom(to)
							cls := room.Client
							for _, cl := range cls {
								if !cl.Limited {
									for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
									}
									break
								}
							}
						} else {
							list := append([]*linetcr.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < numb {
									o.LeaveGroup(to)
									linetcr.GetRoom(to).RevertGo(o)
								} else {
									break
								}
							}
							room := linetcr.GetRoom(to)
							cls := room.Cans()
							for _, cl := range cls {
								if !cl.Limited {
									for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
									}
									break
								}
							}
						}
					}
				}
			} else if cmd == "leaveall" {
				rngcmd := GetComs(4, "leaveall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, p := range bk {
							gr, _ := p.GetGroupIdsJoined()
							for _, g := range gr {
								if g != msg.To {
									p.LeaveGroup(g)
									time.Sleep(1 * time.Second)
								}
							}
						}
						LogMode = false
						LogGroup = ""
						newsend += "Leave done"
						linetcr.RoomClear(room)
						SaveBackup()
					}
				}
			} else if strings.HasPrefix(cmd, "bring") {
				rngcmd := GetComs(5, "bring")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "bring ", "", 1)
						numb, _ := strconv.Atoi(str)
						if numb > 0 && numb <= len(ClientBot) {
							GetSquad(client, to)
							all := []string{}
							room := linetcr.GetRoom(to)
							cuk := room.Client
							alls := []*linetcr.Account{}
							for _, x := range ClientBot {
								if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
									alls = append(alls, x)
								}
							}
							sort.Slice(all, func(i, j int) bool {
								return alls[i].KickCount < alls[j].KickCount
							})
							for _, x := range ClientBot {
								if len(all) < numb {
									if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
										all = append(all, x.MID)
									}
								} else {
									break
								}
							}
							cl := linetcr.GetRoom(to).Choose(client)
							if cl.Limited {
								cl.InviteIntoGroupNormal(to, all)
								time.Sleep(1 * time.Second)
								GetSquad(client, to)
							} else {
								client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							}
						} else {
							newsend += "out of range.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "stay ") {
				rngcmd := GetComs(7, "stay")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "stay ", "", 1)
						numb, _ := strconv.Atoi(str)
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						if numb > 0 && numb <= len(ClientBot) {
							GetSquad(client, to)
							room := linetcr.GetRoom(to)
							aa := len(room.Client)
							if aa > numb {
								c := aa - numb
								ca := 0
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for _, o := range list {
									o.LeaveGroup(to)
									ca = ca + 1
									if ca == c {
										break
									}
								}
								GetSquad(client, to)
							} else if aa < numb {
								ti, err := client.ReissueChatTicket(to)
								if err == nil {
									go client.UpdateChatQrV2(to, false)
									all := []*linetcr.Account{}
									room := linetcr.GetRoom(to)
									cuk := room.Client
									for _, x := range ClientBot {
										if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
											all = append(all, x)
										}
									}
									sort.Slice(all, func(i, j int) bool {
										return all[i].KickPoint < all[j].KickPoint
									})
									g := numb - aa
									var wg sync.WaitGroup
									wi := GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
										l := all[i]
										if l != client && !linetcr.InArrayCl(wi, l) {
											wg.Add(1)
											go func() {
												l.AcceptTicket(to, ti)
												wg.Done()
											}()
										}
									}
									wg.Wait()
									client.UpdateChatQrV2(to, true)
									GetSquad(client, to)
									logAccess(client, to, sender, "bringbot", []string{}, 2)
									SaveBackup()
									aa := len(room.Client)
									var name string
									name = fmt.Sprintf("Ready %v bots here", aa)
									newsend += name + "\n"
								}
							}
						} else {
							newsend += "out of range.\n"
						}
					}
				}
			} else if cmd == "lastset" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						a := "Lastset Command: "
						a += "\n\n    lkick"
						a += "\n    lcancel"
						a += "\n    Lqr"
						a += "\n    linvite"
						a += "\n    ljoin"
						a += "\n    lleave"
						a += "\n    lcon"
						a += "\n    ltag"
						a += "\n    lmid"
						a += "\n    lmessage"
						a += "\n    lbanlist"
						a += "\n    @me"
						a += "\n    @all"
						a += "\n    @oa"
						a += "\n    pend"
						a += "\n    pendingall"
						a += "\n    numpend"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickercmd" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						a := "✠ Sticker Command:"
						a += "\n\n  stickerkick"
						a += "\n  stickerinvite"
						a += "\n  stickerkickall"
						a += "\n  stickercancel"
						a += "\n  stickerbypass"
						a += "\n  stickerstayall"
						a += "\n  stickerleave"
						a += "\n  stickerrespon"
						a += "\n  stickerclearban"
						a += "\n  stickerclear"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickerclear" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						stkid = ""
						stkpkgid = ""
						stkid2 = ""
						stkpkgid2 = ""
						stkid3 = ""
						stkpkgid3 = ""
						stkid4 = ""
						stkpkgid4 = ""
						stkid5 = ""
						stkpkgid5 = ""
						stkid6 = ""
						stkpkgid6 = ""
						stkid7 = ""
						stkpkgid7 = ""
						stkid8 = ""
						stkpkgid8 = ""
						stkid9 = ""
						stkpkgid9 = ""
						SaveBackup()
						a := "Clear sticker command"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickerkick" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerKick = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerrespon" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerRespon = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerstayall" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerStayall = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerleave" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerLeave = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerinvite" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerInvite = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerkickall" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerKickall = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerbypass" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerBypass = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerclearban" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerClearban = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickercancel" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						getStickerCancelall = 1
						newsend += "Please send sticker.\n"
					}
				}
			} else if pesan == "sname" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, fancy(Sname))
					}
				}
			} else if pesan == "prefix" {
				if GetCodeprem(7, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, fancy("Rname: "+Rname+"\nSname: "+Sname))
					}
				}
			} else if pesan == "rname" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, fancy(Rname))
					}
				}
			} else if pesan == Sname {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, fancy(MsgRespon))
					}
				}
			} else if pesan == Rname {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						client.SendMessage(to, fancy(MsgRespon))
					}
				}
			} else if cmd == "gadmins" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							list := "𝗴𝗮𝗱𝗺𝗶𝗻𝘀:\n"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bot") && cmd != "botlist" {
				rngcmd := GetComs(5, "bot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gban") && cmd != "gbanlist" {
				rngcmd := GetComs(8, "gban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 10
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "nukeban" {
				rngcmd := GetComs(7, "nukeban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						go func() {
						      KIckbansPurges(client, to)
						}()
					}
				}
			} else if cmd == "nukebot" {
				rngcmd := GetComs(7, "nukebot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						pelaku := op.Param2
						defer panicHandle("groupBackup")
						Room := linetcr.GetRoom(to)
						memlist, _ := client.GetChatListMap(to)
						ban := []string{}
						exe := []*linetcr.Account{}
						for mid, _ := range memlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
							} else if MemBan(to, mid) {
								ban = append(ban, mid)
							}
						}
						if len(exe) != 0 {
							sort.Slice(exe, func(i, j int) bool {
								return exe[i].KickPoint < exe[j].KickPoint
							})
							Room.HaveClient = exe
							chat := client.GetChat([]string{to}, true, true)
							if chat == nil { return }
							memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
							var createdTime int64
							for mid, tt := range memberMids {
								if pelaku == mid {
									createdTime = tt
									break
								}
							}
							for mid, tt := range memberMids {
								ct := float64(createdTime/1000 - tt/1000)
								if valid.Abs(ct) <= 10 {
									if MemUser(to, mid) {
										Banned.AddBan(mid)
										ban = append(ban, mid)
									}
								}
							}
							no := 0
							ah := 0
							for _, target := range ban {
								go func(target string, no int) {
									exe[no].DeleteOtherFromChats(to, []string{target})
								}(target, no)
								if ah >= MaxKick {
									no++
									if no >= len(exe) {
										no = 0
									}
									ah = 0
								}
								ah++
							}
						}
					}
                           	client.SendMessage(to, fancy("Not have bot enemy"))
				}
			} else if strings.HasPrefix(cmd, "spamcallto") {
				rngcmd := GetComs(7, "spamcallto")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						targets := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(targets, lists[i]) {
									targets = append(targets, lists[i])
								}
							}
						}
                    		             if msg.ToType == 2 {
						     nu, _ := strconv.Atoi(CountSpam)
						     for i := 1; i <= nu; i++ {
								call.AcquireGroupCallRoute(to, client.AuthToken)
								call.InviteIntoGroupCall(to, client.AuthToken, targets)
							}
						}
                           		time.Sleep(3 * time.Second)
                           		client.SendMessage(to, fancy("Successful spam invite call group"))
					}
				}
			} else if strings.HasPrefix(cmd, "spamcall") {
				rngcmd := GetComs(7, "spamcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "spamcall"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
                        			CountSpam = str
						result := strings.Split(str,"")
                    		             if msg.ToType == 2 {
						     nu, _ := strconv.Atoi(result[0])
							_, target, _ := client.GetChatList(to)
							targets := []string{}
							for i := range target {
								if !InArray2(checkHaid, target[i]) {
									targets = append(targets, target[i])
								}
							 }
							for i := 1; i <= nu; i++ {
								call.AcquireGroupCallRoute(to, client.AuthToken)
								call.InviteIntoGroupCall(to, client.AuthToken, targets)
						       }
						}
                           		client.SendMessage(to, fancy("Successful spam invite call group"))
					}
				}
			} else if cmd == "fixed" {
				rngcmd := GetComs(7, "fixed")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						SaveBackup()
						newsend += "done.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "autobc") {
				rngcmd := GetComs(7, "autobc")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autobc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "msg" {
							AutoBc = true
							Typebc = str
							count = count + 1
							client.SendMessage(to, fancy("Autobroadcast type message"))
						} else if str == "img" {
							AutoBc = true
							SAVEBcImage = true
							Typebc = str
							count = count + 1
							client.SendMessage(to, fancy("Autobroadcast type image\n Please Send Image"))
						} else if str == "off" {
							AutoBc = false
							Typebc = "none"
							count = count + 1
							client.SendMessage(to, fancy("Autobroadcast is disabled"))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("Autobroadcast type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "msgbc") {
				rngcmd := GetComs(7, "msgbc")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "msgbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						MsgBroadcast = str
						SaveBackup()
						newsend += "Message broadcast set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "groupbc") {
				rngcmd := GetComs(7, "groupbc")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "groupbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							GBcImage = true
							client.SendMessage(to, fancy("Please send image"))
						} else if str == "video" {
							GBcVideo = true
							client.SendMessage(to, fancy("Please send video"))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "friendbc") {
				rngcmd := GetComs(7, "friendbc")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "friendbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							FBcImage = true
							client.SendMessage(to, fancy("Please send image"))
						} else if str == "video" {
							FBcVideo = true
							client.SendMessage(to, fancy("Please send video"))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "startbcmid") {
				rngcmd := GetComs(7, "startbcmid")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "startbcmid"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							BcImage = true
							client.SendMessage(to, fancy("Please send image"))
						} else if str == "video" {
							BcVideo = true
							client.SendMessage(to, fancy("Please send video"))
						}
					}
				}
			} else if cmd == "hentai" {
				rngcmd := GetComs(7, "hentai")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.Sendmedias(to, "hentai","hentai")
					}
				}
			} else if cmd == "pornstart" {
				rngcmd := GetComs(7, "pornstart")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.Sendmedias(to, "pornstart","pornstart")
					}
				}
			} else if strings.HasPrefix(cmd, "videoporn") {
				rngcmd := GetComs(7, "videoporn")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "videoporn"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"pornvideo")
					}
				}
			} else if strings.HasPrefix(cmd, "tiktok") {
				rngcmd := GetComs(7, "tiktok")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "tiktok"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"tiktok")
					}
				}
			} else if strings.HasPrefix(cmd, "smule") {
				rngcmd := GetComs(7, "smule")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "smule"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"smule")
					}
				}
			} else if strings.HasPrefix(cmd, "joox") {
				rngcmd := GetComs(7, "joox")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "joox"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"joox")
					}
				}
			} else if strings.HasPrefix(cmd, "youtube") {
				rngcmd := GetComs(7, "youtube")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "youtube"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"youtube")
					}
				}
			} else if strings.HasPrefix(cmd, "instagram") {
				rngcmd := GetComs(7, "instagram")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "instagram"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"instagram")
					}
				}
			} else if strings.HasPrefix(cmd, "textimage") {
				rngcmd := GetComs(7, "textimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "textimage"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"image_text")
					}
				}
			} else if strings.HasPrefix(cmd, "calculator") {
				rngcmd := GetComs(7, "calculator")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "calculator"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"calculator")
					}
				}
			} else if strings.HasPrefix(cmd, "cuaca") {
				rngcmd := GetComs(7, "cuaca")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "cuaca"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"cuaca")
					}
				}
			} else if strings.HasPrefix(cmd, "simi") {
				rngcmd := GetComs(7, "simi")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "simi"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"simi")
					}
				}
			} else if strings.HasPrefix(cmd, "artinama") {
				rngcmd := GetComs(7, "artinama")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "artinama"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"artinama")
					}
				}
			} else if strings.HasPrefix(cmd, "gimage") {
				rngcmd := GetComs(7, "gimage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "gimage"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"google")
					}
				}
			} else if strings.HasPrefix(cmd, "pinterest") {
				rngcmd := GetComs(7, "pinterest")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "pinterest"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"pinterest")
					}
				}
			} else if cmd == "#gtotal" {
				rngcmd := GetComs(7, "gtotal")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						list := Gtotal(client, to)
						client.SendMessage(to, fancy(list))
					}
				}
			} else if cmd == "tes" {
				rngcmd := GetComs(7, "tes")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						for _, tok := range Data.Authoken {
							sendBigImage(tok, to, "https://thumbs.gfycat.com/SmartDenseBuck.webp")
					       }
					}
				}
			} else if cmd == "getcall" {
				rngcmd := GetComs(7, "getcall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
                    		             if msg.ToType == 2 {
                        		            gcall, _ := call.GetGroupCall(to, client.AuthToken)
                                             Room := linetcr.GetRoom(to)
                        		            res := "Get Call Group:"
                        		            if gcall.MediaType == 1 {
                                                    res += "\n  • Type: Audio Call"
                        		                   res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                    loc, _ := time.LoadLocation("Asia/Jakarta")
	                                             a := time.Now().In(loc)
	                                             yyyy := strconv.Itoa(a.Year())
	                                             MM := a.Month().String()
	                                             dd := strconv.Itoa(a.Day())
	                                             Date := dd + "-" + MM + "-" + yyyy
                                                    cok := gcall.Started / 1000
			                                i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                                tm := time.Unix(i, 0)
			                                ss := time.Since(tm)
			                                sp := fmtDuration(ss)
                                                    res += "\n  • Date: "+Date
                                                    res += "\n  • Started: "+sp
                        		                   res += "\n  • Members:"
                                                    mmk := gcall.MemberMids
						            if len(mmk) != 0 {
							            for num, xd := range mmk {
								            num++
								            rengs := strconv.Itoa(num)
								            x, _ := client.GetContact(xd)
								            res += "\n      " + rengs + ". "+x.DisplayName
							              }
					                     }
                                                       client.SendMessage(to, fancy(res))
                                             }
                        		            if gcall.MediaType == 2 {
                                                    res += "\n  • Type: Video Call"
                        		                   res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                    loc, _ := time.LoadLocation("Asia/Jakarta")
	                                             a := time.Now().In(loc)
	                                             yyyy := strconv.Itoa(a.Year())
	                                             MM := a.Month().String()
	                                             dd := strconv.Itoa(a.Day())
	                                             Date := dd + "-" + MM + "-" + yyyy
                                                    cok := gcall.Started / 1000
			                                i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                                tm := time.Unix(i, 0)
			                                ss := time.Since(tm)
			                                sp := fmtDuration(ss)
                                                    res += "\n  • Date: "+Date
                                                    res += "\n  • Started: "+sp
                        		                   res += "\n  • Members:"
                                                    mmk := gcall.MemberMids
						            if len(mmk) != 0 {
							            for num, xd := range mmk {
								            num++
								            rengs := strconv.Itoa(num)
								            x, _ := client.GetContact(xd)
								            res += "\n      " + rengs + ". "+x.DisplayName
							              }
					                     }
                                                       client.SendMessage(to, fancy(res))
                                               }
					      }
				      }
				}
			} else if cmd == "bans" {
				rngcmd := GetComs(7, "bans")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if SendMybuyer(sender) {
							CheckChatBan()
						}
						toy := len(linetcr.KickBans)
						banchat := len(linetcr.KickBanChat)
						addb := len(linetcr.Waitadd)
						kb := ""
						if toy != 0 {
							kb += fmt.Sprintf("%v/%v bot's kick/inv banned.", toy, len(Squadlist))
							for n, cl := range linetcr.KickBans {
								m := cl.MID
								no := n + 1
								var ta time.Duration
								if _, ok := linetcr.GetBlock.Get(cl.MID); ok {
									t := cl.TimeBan.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.TimeBan.Add(1*time.Hour )
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\n%s\nRemaining %v", no, cl.Namebot, cl.MID, fmtDurations(ta))
							}
						}
						fris := []*linetcr.Account{}
						for _, cl := range ClientBot {
							if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
								if cl.Frez {
									fris = append(fris, cl)
								}
							}
						}
						if len(fris) != 0 {
							no := 1
							mm := kb
							kb += fmt.Sprintf("\n\n%v/%v bot's freeze.", len(fris), len(Squadlist))
							for _, cl := range fris {
								t := cl.TimeBan.Add(1*time.Hour )
								ta := t.Sub(time.Now())
								if ta > 1*time.Second {
									kb += fmt.Sprintf("\n\n%v. %s\n%s\nRemaining %v", no, cl.Namebot, cl.MID, fmtDurations(ta))
									no++
								} else {
									if _, ok := linetcr.GetBlock.Get(cl.MID); !ok {
										linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
										cl.Limited = false
									}
									cl.Frez = false
								}
							}
							if no == 1 {
								kb = mm
							}
						}
						if addb != 0 {
							kb += fmt.Sprintf("\n\n%v/%v bot's add/friend banned.", addb, len(Squadlist))
							for n, cl := range linetcr.Waitadd {
								if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
									m := cl.MID
									no := n + 1
									var ta time.Duration
									if _, ok := linetcr.BlockAdd.Get(cl.MID); ok {
										t := cl.Timeadd.Add(24 * time.Hour)
										ta = t.Sub(time.Now())
									} else {
										t := cl.Timeadd.Add(1 * time.Hour)
										ta = t.Sub(time.Now())
									}
									if cl.Namebot == "" {
										pr, _ := client.GetContact(m)
										cl.Namebot = pr.DisplayName
									}
									kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, fmtDurations(ta))
								}
							}
						}
						if banchat != 0 {
							kb += fmt.Sprintf("\n\n%v/%v bot's banchat.", banchat, len(Squadlist))
							for n, cl := range linetcr.KickBanChat {
								m := cl.MID
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nmid: %s", n+1, cl.Namebot, m)
							}
						}
						if len(fris) == 0 && toy == 0 && addb == 0 && banchat == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}

				}
			} else if cmd == "botlist" {
				rngcmd := GetComs(5, "botlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Bot) != 0 {
							list := "Botlist:\n"
							targets := []string{}
							for _, i := range UserBot.Bot {
								targets = append(targets, i)
							}
							client.SendPollMention(to, list, targets)
						} else {
							newsend += "Botlist is empty.\n"
						}
					}
				}
			} else if cmd == "clearbot" {
				rngcmd := GetComs(5, "clearbot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(UserBot.Bot) != 0 {
							newsend += fmt.Sprintf("Cleared %v botlist\n", len(UserBot.Bot))
							logAccess(client, to, sender, "clearbot", UserBot.Bot, msg.ToType)
							UserBot.ClearBot()
						} else {
							newsend += "Bot is empty.\n"
						}
					}
				}
			} else if cmd == "cleargban" {
				rngcmd := GetComs(8, "cleargban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							logAccess(client, to, sender, "cleargban", room.Gban, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gbanlist", len(room.Gban)) + "\n"
							room.Gban = []string{}
						} else {
							newsend += "Gban is empty.\n"
						}
					}
				}
			} else if cmd == "clears" || cmd == Commands.Clearchat && Commands.Clearchat != "" {
				rngcmd := GetComs(5, "clears")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						for i := range ClientBot {
							if InArray2(memb, ClientBot[i].MID) {
								ClientBot[i].RemoveAllMessage(string(op.Param2))
							}
						}
						newsend += "Cleared all message.\n"
					}
				}
			} else if cmd == "clearall" {
				rngcmd := GetComs(5, "clearall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						for i := range ClientBot {
							if InArray2(memb, ClientBot[i].MID) {
								ClientBot[i].RemoveAllMessage(string(op.Param2))
							}
						}
						exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
						exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
						newsend += "Cleared all message and cache.\n"
					}
				}
			} else if cmd == "clearcache" {
				rngcmd := GetComs(5, "clearcache")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
						exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
						//exec.Command("bash", "-c", "sudo apt update").Output()
						newsend += "Cleared all cache.\n"
					}
				}
			} else if cmd == "enablee2ee" {
				rngcmd := GetComs(1, "enablee2ee")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.EnableE2ee()
						newsend += "Done Succes Enable E2EE.\n"
					}
				}
			} else if cmd == "disablee2ee" {
				rngcmd := GetComs(1, "disablee2ee")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						client.DisableE2ee()
						newsend += "Done Succes Disable E2EE.\n"
					}
				}
			} else if cmd == "gbanlist" {
				rngcmd := GetComs(8, "gbanlist")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							list := "Gbanlist:"
							client.SendPollMention(to, list, room.Gban)
						} else {
							newsend += "Gban list is empty.\n"
						}
					}
				}
			} else if cmd == "infogo" {
				rngcmd := GetComs(8, "infogo")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						o, _ := host.Info()
						r, _ := api.GetRunningVersion()
						Gplat := fmt.Sprintf("%v", r.Platform)
						//Gover := fmt.Sprintf("%v", r.Version)
						Garch := fmt.Sprintf("%v", r.Architecture)
						OS := fmt.Sprintf("Ubuntu %v ", o.PlatformVersion)
                   				a := "✠ INFORMATION:"
						a += "\n\n ⚙️ Platfrom : " + Gplat
						a += "\n ⚙️ OS : " + OS
						a += "\n ⚙️ Executed : Go 1.22.5"
						a += "\n ⚙️ Architecture : " + Garch
                   				a += "\n ⚙️ AppName : ANDROID"
                   				a += "\n ⚙️ UserAgent : Line/14.10.0"
                   				a += "\n ⚙️ Host : legy-jp-addr-long"
                   				a += "\n ⚙️ X-lal : "+client.Locale
						a += "\n ⚙️ Update : 10-07-2024"
						a += "\n ⚙️ Version : Sync5"
						a += "\n ⚙️ Condition : Good"
						a += "\n ⚙️ Team : SELFTCR™"
						a += InfoCreator(client)
						newsend += a
					}
				}
			} else if cmd == "here" || cmd == Commands.Here && Commands.Here != "" {
				rngcmd := GetComs(6, "here")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						aa := len(room.Client)
						//bb := len(room.Client)
						cc := len(room.GoMid)
						var name string
						name = fmt.Sprintf("%v bots here - %v squad", aa, len(ClientBot))
						if cc != 0 {
							name += fmt.Sprintf("\n%v bots on stay.", cc)
						}
						toy := len(linetcr.KickBans)
						if toy != 0 {
							name += fmt.Sprintf("\n%v bots limited", toy)
						}
						newsend += name + "\n"
					}
				}
			} else if cmd == "ourl" {
				rngcmd := GetComs(6, "ourl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						tick, err := client.ReissueChatTicket(to)
						if err == nil {
							var err error
							mes := make(chan bool)
							go func() {
								err = client.UpdateChatQrV2(to, false)
								if err != nil {
									mes <- false
								} else {
									mes <- true
								}
							}()
							newsend += "https://line.me/R/ti/g/" + tick + "\n"
						}
					}
				}
			} else if cmd == "curl" {
				rngcmd := GetComs(6, "curl")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						var err error
						mes := make(chan bool)
						go func() {
							err = client.UpdateChatQrV2(to, true)
							if err != nil {
								mes <- true
							} else {
								mes <- false
							}
						}()
					}
				}
			} else if strings.HasPrefix(cmd, "say ") {
				rngcmd := GetComs(6, "say")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "say ", "", 1)
						client.SendMessage(to, fancy(str))
					}
				}
			} else if cmd == "timeleft" {
				rngcmd := GetComs(6, "timeleft")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", Data.Dalltime)
						has := strings.Split(d, "-")
						hass := strings.Split(has[2], "T")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(hass[0])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							str := fmt.Sprintf("Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\nRemaining:\n  %v", botDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "timenow" {
				rngcmd := GetComs(7, "timenow")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GenerateTimeLog(client, to)
					}
				}
			} else if cmd == "runtime" {
				rngcmd := GetComs(7, "runtime")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						elapsed := time.Since(botStart)
						newsend += "Running Time:\n\n" + botDuration(elapsed) + "\n"
					}
				}
			} else if cmd == "setbot" {
				rngcmd := GetComs(5, "setbot")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "✠ 𝗦𝗲𝘁 𝗠𝗼𝗱𝗲 𝗕𝗼𝘁𝘀:"
						ret += "\n"
						if AutoPro {ret += "\n ⚫ » ᴀᴜᴛᴏ ᴘʀᴏ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴘʀᴏ"}
						if AutoPurge {ret += "\n ⚫ » ᴀᴜᴛᴏ ᴘᴜʀɢᴇ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴘᴜʀɢᴇ"}
						if AutoJointicket {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴛɪᴄᴋᴇᴛ : %s", TypeJoin)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴛɪᴄᴋᴇᴛ"}
						if AutoTranslate {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴛʀᴀɴꜱʟᴀᴛᴇ : %s", TypeTrans)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴛʀᴀɴꜱʟᴀᴛᴇ"}
						if AutoLike {ret += "\n ⚫ » ᴀᴜᴛᴏ ʟɪᴋᴇ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ʟɪᴋᴇ"}
						if AutoBc {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ʙʀᴏᴀᴅᴄᴀꜱᴛ : %s", Typebc)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ʙʀᴏᴀᴅᴄᴀꜱᴛ"}
						if Autojoin != "off" {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴊᴏɪɴ :  %s", Autojoin)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴊᴏɪɴ"}
						if Ajsjoin != "off" {ret += fmt.Sprintf("\n ⚫ » ᴀᴊꜱ ᴊᴏɪɴ :  %s", Ajsjoin)} else {ret += "\n ⚪ » ᴀᴊꜱ ᴊᴏɪɴ"}
						if canceljoin {ret += "\n ⚫ » ᴄᴀɴᴄᴇʟ ᴊᴏɪɴ"} else {ret += "\n ⚪ » ᴄᴀɴᴄᴇʟ ᴊᴏɪɴ"}
						if NukeJoin {ret += "\n ⚫ » ɴᴜᴋᴇ ᴊᴏɪɴ"} else {ret += "\n ⚪ » ɴᴜᴋᴇ ᴊᴏɪɴ"}
						if PowerMode {ret += "\n ⚫ » ᴍᴏᴅᴇ ᴡᴀʀ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ᴡᴀʀ"}
						if ProtectMode {ret += "\n ⚫ » ᴍᴏᴅᴇ ᴘʀᴏᴛᴇᴄᴛ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ᴘʀᴏᴛᴇᴄᴛ"}
						if AutoBackBot {ret += "\n ⚫ » ᴍᴏᴅᴇ ʙᴀᴄᴋ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ʙᴀᴄᴋ"}
						if KickBanQr {ret += "\n ⚫ » ᴋɪᴄᴋ ʙᴀɴQʀ"} else {ret += "\n ⚪ » ᴋɪᴄᴋ ʙᴀɴQʀ"}
						if ModeBackup != "" {ret += fmt.Sprintf("\n ⚫ » ʙᴀᴄᴋᴜᴘ : %s", ModeBackup)} else {ret += "\n ⚪ » ʙᴀᴄᴋᴜᴘ"}
						if linetcr.FlexMode {ret += "\n ⚫ » ꜰʟᴇx ᴍᴏᴅᴇ1"} else if linetcr.FlexMode2 {ret += "\n ⚫ » ꜰʟᴇx ᴍᴏᴅᴇ2"} else if linetcr.FooterMode {ret += "\n ⚫ » ꜰᴏᴏᴛᴇʀ ᴍᴏᴅᴇ"} else {ret += "\n ⚪ » ꜰʟᴇx / ꜰᴏᴏᴛᴇʀ"}
						if MediaDl {ret += "\n ⚫ » ᴍᴇᴅɪᴀ ᴅᴏᴡɴʟᴏᴀᴅ"} else {ret += "\n ⚪ » ᴍᴇᴅɪᴀ ᴅᴏᴡɴʟᴏᴀᴅ"}
						//ret += "\n\n▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎"
						ret += "\n"
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴋɪᴄᴋ : %v", Data.Logobot, MaxKick)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ɪɴᴠɪᴛᴇ : %v", Data.Logobot, MaxInvite)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴄᴀɴꜱ : %v", Data.Logobot, MaxCancel)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴘᴇɴᴅ : %v", Data.Logobot, CancelPend)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ꜱᴘᴀᴍ : %v", Data.Logobot, CountSpam)
						if LockAjs {ret += fmt.Sprintf("\n %v ʟᴏᴄᴋ ᴀɴᴛɪᴊꜱ : %v (true)", Data.Logobot, CountAjs)} else {ret += fmt.Sprintf("\n %v ʟᴏᴄᴋ ᴀɴᴛɪᴊꜱ : %v (false)", Data.Logobot, CountAjs)}
						ret += fmt.Sprintf("\n %v ᴍᴏᴅᴇᴛᴇxᴛ : %v", Data.Logobot, fancy)
						//ret += fmt.Sprintf("\n %v ᴍᴏᴅᴇᴛᴇxᴛ : "+fancy, Data.Logobot)
						rng1 := GetComs(7, "invitebot")
						rng12 := GetComs(4, "remote")
						xx := GETgrade(rng1)
						yy := GETgrade(rng12)
						ret += fmt.Sprintf("\n %v ᴘᴇʀᴍ ɪɴᴠɪᴛᴇʙᴏᴛ : %v ", Data.Logobot, xx)
						ret += fmt.Sprintf("\n %v ᴘᴇʀᴍ ʀᴇᴍᴏᴛᴇ : %v ", Data.Logobot, yy)
						ret += fmt.Sprintf("\n %v ᴄʀᴇᴀᴛɪᴏɴ ᴛᴇᴀᴍ : ꜱᴇʟꜰᴛᴄʀ™", Data.Logobot)
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "setgroup" {
				rngcmd := GetComs(9, "set")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						aa := len(room.Client)
						cc := len(room.GoMid)
						ret := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁 𝗦𝗲𝘁𝘁𝗶𝗻𝗴:\n"
						if op.Message.ToType == 2 {
							if room.ProQr {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							}
							if room.ProKick {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							}
							if room.ProInvite {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							}
							if room.ProCancel {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							}
							if room.ProJoin {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							}
							if room.ProName {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							}
							if room.AntiTag {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							}
							if room.ProPicture {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							}
							if room.ProNote {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							}
							if room.ProAlbum {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							}
							if room.ProLink {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							}
							if room.ProFlex {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							}
							if room.ProImage {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							}
							if room.ProVideo {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							}
							if room.ProCall {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							}
							if room.ProSpam {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							}
							if room.ProSticker {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							}
							if room.ProContact {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							}
							if room.ProPost {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							}
							if room.ProFile {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							}
							if len(room.GoMid) > 0 {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							}
							ret += "\n"
							ret += "\n✠ 𝗚𝗿𝗼𝘂𝗽 𝗦𝗲𝘁𝘁𝗶𝗻𝗴:\n"
							if room.Lurk {
								ret += fmt.Sprintf("\n ⚫ » ʟᴜʀᴋɪɴɢ %s", room.NameLurk)
							} else {
								ret += "\n ⚪ » ʟᴜʀᴋɪɴɢ"
							}
							if room.Automute {
								ret += "\n ⚫ » ᴀᴜᴛᴏᴍᴜᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴀᴜᴛᴏᴍᴜᴛᴇ"
							}
							if room.Welcome {
								ret += "\n ⚫ » ᴡᴇʟᴄᴏᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴡᴇʟᴄᴏᴍᴇ"
							}
							if room.Leavebool {
								ret += "\n ⚫ » ʟᴇᴀᴠᴇ"
							} else {
								ret += "\n ⚪ » ʟᴇᴀᴠᴇ"
							}
							if room.Announce {
								ret += "\n ⚫ » ᴀɴɴᴏᴜɴᴄᴇ"
							} else {
								ret += "\n ⚪ » ᴀɴɴᴏᴜɴᴄᴇ"
							}
							if room.Backleave {
								ret += "\n ⚫ » ʜᴏꜱᴛᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ʜᴏꜱᴛᴀɢᴇ"
							}
							if LogGroup == to {
								ret += "\n ⚫ » ɴᴏᴛɪꜰɪᴄᴀᴛɪᴏɴ"
							} else {
								ret += "\n ⚪ » ɴᴏᴛɪꜰɪᴄᴀᴛɪᴏɴ"
							}
							if room.ImageLurk {
								ret += "\n ⚫ » ʟᴜʀᴋ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ʟᴜʀᴋ ɪᴍᴀɢᴇ"
							}
							if room.Backup {
								ret += "\n ⚫ » ʙᴀᴄᴋᴜᴘ ᴜꜱᴇʀ"
							} else {
								ret += "\n ⚪ » ʙᴀᴄᴋᴜᴘ ᴜꜱᴇʀ"
							}
							if DetectCall {
								ret += "\n ⚫ » ᴅᴇᴛᴇᴄᴛ ɢʀᴏᴜᴘᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴅᴇᴛᴇᴄᴛ ɢʀᴏᴜᴘᴄᴀʟʟ"
							}
						}
						ret += fmt.Sprintf("\n\n %v/%v ʙᴏᴛꜱ ʜᴇʀᴇ.", aa, len(ClientBot))
						if cc != 0 {
							ret += fmt.Sprintf("\n %v ʙᴏᴛꜱ ᴏɴ ꜱᴛᴀʏ.", cc)
						}
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "lurk image" {
				//rngcmd := GetComs(9, "image")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.ImageLurk = true
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk name" {
				//rngcmd := GetComs(8, "name")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk mention" {
				//rngcmd := GetComs(8, "mention")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "mention"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk on" {
				//rngcmd := GetComs(8, "on")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "mediadl on" {
				rngcmd := GetComs(9, "on")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						MediaDl = true
						newsend += "Mediad download enabled.\n"
					}
				}
			} else if cmd == "mediadl off" {
				rngcmd := GetComs(9, "off")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						MediaDl = false
						newsend += "Media download enabled.\n"
					}
				}
			} else if cmd == "autolike on" {
				rngcmd := GetComs(9, "on")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						AutoLike = true
						newsend += "Autolike enabled.\n"
					}
				}
			} else if cmd == "autolike off" {
				rngcmd := GetComs(9, "off")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						AutoLike = false
						newsend += "Autolike enabled.\n"
					}
				}
			} else if cmd == "bomlike 10" {
				rngcmd := GetComs(9, "10")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						BomLike = true
						newsend += "Share you post\n"
					}
				}
			} else if cmd == "bomlike off" {
				rngcmd := GetComs(9, "off")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						BomLike = false
						newsend += "Bomlike enabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "countspam") {
				rngcmd := GetComs(5, "countspam")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "countspam"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
                        			CountSpam = str
                        			client.SendMessage(to, fancy("Success update countspam: "+CountSpam+""))
					}
				}
			} else if strings.HasPrefix(cmd, "countajs") {
				rngcmd := GetComs(5, "countajs")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "countajs"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
                        			CountAjs = str
                        			client.SendMessage(to, fancy("Success update countajs: "+CountAjs+""))
					}
				}
			} else if strings.HasPrefix(cmd, "killmode") {
				rngcmd := GetComs(5, "killmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "killmode"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "kill" {
							Killmode = str
							count = count + 1
						} else if str == "purge" {
							Killmode = str
							count = count + 1
						} else if str == "range" {
							Killmode = str
							count = count + 1
						} else if str == "random" {
							Killmode = str
							count = count + 1
						} else if str == "off" {
							Killmode = "none"
							count = count + 1
						}
						if count != 0 {
							newsend += fmt.Sprintf("Killmode state : %s\nTurn on", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "textmode") {
				rngcmd := GetComs(5, "textmode")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						var su = "textmode"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "1" {
							fancy = unistyle.BoldSans
						} else if str == "2" {
							fancy = unistyle.BoldSerif
						} else if str == "3" {
							fancy = unistyle.ItalicSans
						} else if str == "4" {
							fancy = unistyle.ItalicSerif
						} else if str == "5" {
							fancy = unistyle.BoldItalicSans
						} else if str == "6" {
							fancy = unistyle.BoldItalicSerif
						} else if str == "7" {
							fancy = unistyle.Cursive
						} else if str == "8" {
							fancy = unistyle.Fraktur
						} else if str == "9" {
							fancy = unistyle.BoldMan
						} else if str == "10" {
							fancy = unistyle.BoldMin
						} else if str == "normal" {
							fancy = unistyle.Normal
						}
						newsend += "Update mode text: "+str
					}
				}
			} else if strings.HasPrefix(cmd, "backup") {
				rngcmd := GetComs(5, "backup")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "backup"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							ModeBackup = str
							count = count + 1
						} else if str == "qr" {
							ModeBackup = str
							count = count + 1
						}
						if count != 0 {
							newsend += fmt.Sprintf("Mode Backup state : %s\nTurn on", str)
						}
					}
				}
			} else if cmd == "lurk" {
				rngcmd := GetComs(6, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = true
						room.Userlurk = []string{}
						room.NameLurk = "hide"
						newsend += "Lurking...\n"
					}
				}
			} else if cmd == "lurks" {
				rngcmd := GetComs(6, "lurk")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(room.Userlurk) != 0 {
							list := "✠ Lurkers:\n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"

						} else {
							newsend += "Lurk list empty enable first.\n"
						}
					}
				}
			} else if cmd == "lurk off" {
				rngcmd := GetComs(9, "off")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						room.Lurk = false
						room.ImageLurk = false
						if len(room.Userlurk) != 0 {
							list := " ✠ Lurkers ✠ \n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								new := client.Getcontactuser(xd)
								if new != nil {
									list += "\n   " + rengs + ". Closed Account"
								} else {
									x, _ := client.GetContact(xd)
									list += "\n   " + rengs + ". " + x.DisplayName
								}
							}
							newsend += list + "\n"
						}
						room.Userlurk = []string{}
					}
				}
			} else if cmd == "/status all" {
				rngcmd := GetComs(6, "/statusall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "✠ Status Allbot:"
						ret += "\n"
						for i := range ClientBot {
							ClientBot[i].DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							if ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\n%v. "+ClientBot[i].Namebot+": %s", i+1, Data.Limit)
							} else {
								ret += fmt.Sprintf("\n%v. "+ClientBot[i].Namebot+": %s", i+1, Data.Fresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "status add" {
				rngcmd := GetComs(6, "status add")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "✠ Status Add:"
						ret += "\n"
						for i := range ClientBot {
							ve := "uc52554b082eca0360da013d33df023e0"
							err, _ := ClientBot[i].FindAndAddContactsByMidV2(ve)
							fff := fmt.Sprintf("%v", err)
							er := strings.Contains(fff, "request blocked")
							if er == true {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, Data.Limit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, Data.Fresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "/status" || cmd == Commands.Status && Commands.Status != "" {
				rngcmd := GetComs(6, "/status")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						var a = 0
						ret := "✠ Status Bot:"
						ret += "\n"
						for i := range ClientBot {
							if InArray2(memb, ClientBot[i].MID) {
								a = a + 1
								ClientBot[i].DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
								if ClientBot[i].Limited == true {
									ret += fmt.Sprintf("\n%v. "+ClientBot[i].Namebot+": %s", a, Data.Limit)
								} else {
									ret += fmt.Sprintf("\n%v. "+ClientBot[i].Namebot+": %s", a, Data.Fresh)
								}
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "sp" {
				rngcmd := GetComs(8, "sp")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						for _, p := range bk {
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret := fmt.Sprintf("%v ms", sp)
							p.SendMessage(to, fancy(ret))
						}
					}
				}
			} else if cmd == "speed" || cmd == "speed" || cmd == Commands.Speed && Commands.Speed != "" {
				rngcmd := GetComs(8, "speed")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "✠ Speed Bot:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", a, sp)
						}
						ret += "\n"
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "speed all" {
				rngcmd := GetComs(8, "speed all")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "✠ Speed Profile:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							start := time.Now()
							p.GetProfile()
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", a, sp)
						}
						var b = 0
						ret += "\n\n✠ Speed Contact:"
						ret += "\n"
						for _, p := range bk {
							b = b + 1
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", b, sp)
						}
						var c = 0
						ret += "\n\n✠ Speed Message:"
						ret += "\n"
						for _, p := range bk {
							c = c + 1
							start := time.Now()
							p.SendMessage("u27623a2c021c18746b7aa34e3d2b2220", "sp")
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", c, sp)
						}
						var d = 0
						ret += "\n\n✠ Speed Kick:"
						ret += "\n"
						for _, p := range bk {
							d = d + 1
							start := time.Now()
							p.DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", d, sp)
						}
						var e = 0
						ret += "\n\n✠ Speed Invite:"
						ret += "\n"
						for _, p := range bk {
							e = e + 1
							start := time.Now()
							p.InviteIntoGroupNormal(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", e, sp)
						}
						var f = 0
						ret += "\n\n✠ Speed Cancel:"
						ret += "\n"
						for _, p := range bk {
							f = f + 1
							start := time.Now()
							p.CancelChatInvitations(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", f, sp)
						}
						ret += "\n"
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "status" || cmd == Commands.Status && Commands.Status != "" {
				rngcmd := GetComs(8, "tus")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "✠ Status Bot:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							if p.Limited == true {
								ret += fmt.Sprintf("\n Bots%v: %s", a, MsLimit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", a, MsFresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if cmd == "status all" {
				rngcmd := GetComs(5, "statusall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						ret := "✠ Status Allbot:"
						ret += "\n"
						for i := range ClientBot {
							if ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, MsLimit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, MsFresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, fancy(ret))
					}
				}
			} else if strings.HasPrefix(cmd, "help ") && cmd != "help" {
				if !MemUser(to, sender) {
					if CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "help ", "")
						texts := strings.Split(txt, " ")
						if len(texts) != 0 {
							kata := texts[0]
							if kata == "all" {
								res := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁𝗶𝗼𝗻 𝗠𝗲𝗻𝘂:"
								res += "\n"
								for _, x := range helppro {
									res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
								}
								if InArray2(DEVELOPER,sender) {
									if GetCodeprem(0, sender, to) {
										res += "\n"
										res += "\n✠ 𝗗𝗲𝘃𝗲𝗹𝗼𝗽𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpdeveloper {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMycreator(sender) {
									if GetCodeprem(1, sender, to) {
										res += "\n"
										res += "\n✠ 𝗖𝗿𝗲𝗮𝘁𝗼𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpcreator {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMymaker(sender) {
									if GetCodeprem(2, sender, to) {
										res += "\n"
										res += "\n✠ 𝗠𝗮𝗸𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpmaker {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMyseller(sender) {
									if GetCodeprem(3, sender, to) {
										res += "\n"
										res += "\n✠ 𝗦𝗲𝗹𝗹𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpseller {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMybuyer(sender) {
									if GetCodeprem(4, sender, to) {
										res += "\n"
										res += "\n✠ 𝗕𝘂𝘆𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpbuyer {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMyowner(sender) {
									if GetCodeprem(5, sender, to) {
										res += "\n"
										res += "\n✠ 𝗢𝘄𝗻𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpowner {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMymaster(sender) {
									if GetCodeprem(6, sender, to) {
										res += "\n"
										res += "\n✠ 𝗠𝗮𝘀𝘁𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpmaster {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
									}
								}
								if SendMyadmin(sender) {
									if GetCodeprem(7, sender, to) {
										res += "\n"
										res += "\n✠ 𝗔𝗱𝗺𝗶𝗻 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range helpadmin {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, res + "\n")//fancy(res + "\n"))
									}
								}
							} else if kata == "dev" {
								if GetCodeprem(0, sender, to) {
									if InArray2(DEVELOPER,sender) {
										res := "✠ 𝗗𝗲𝘃𝗲𝗹𝗼𝗽𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpdeveloper {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "creator" {
								if GetCodeprem(1, sender, to) {
									if SendMycreator(sender) {
										res := "✠ 𝗖𝗿𝗲𝗮𝘁𝗼𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpcreator {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "maker" {
								if GetCodeprem(1, sender, to) {
									if SendMymaker(sender) {
										res := "✠ 𝗠𝗮𝗸𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpmaker {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "seller" {
								if GetCodeprem(2, sender, to) {
									if SendMyseller(sender) {
										res := "✠ 𝗦𝗲𝗹𝗹𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpseller {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "buyer" {
								if GetCodeprem(3, sender, to) {
									if SendMybuyer(sender) {
										res := "✠ 𝗕𝘂𝘆𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpbuyer {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "owner" {
								if GetCodeprem(4, sender, to) {
									if SendMyowner(sender) {
										res := "✠ 𝗢𝘄𝗻𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpowner {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "master" {
								if GetCodeprem(5, sender, to) {
									if SendMymaster(sender) {
										res := "✠ 𝗠𝗮𝘀𝘁𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpmaster {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "admin" {
								if GetCodeprem(6, sender, to) {
									if SendMyadmin(sender) {
										res := "✠ 𝗔𝗱𝗺𝗶𝗻 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpadmin {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "gowner" {
								if GetCodeprem(8, sender, to) {
									if SendMygowner(to, sender) {
										res := "✠ 𝗚𝗼𝘄𝗻𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpgowner {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "gadmin" {
								if GetCodeprem(9, sender, to) {
									if SendMygadmin(to, sender) {
										res := "✠ 𝗚𝗮𝗱𝗺𝗶𝗻 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range helpgadmin {
											res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
										}
										client.SendHelp(to, fancy(res + "\n"))
									}
								}
							} else if kata == "protect" {
								res := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
								res += "\n"
								for _, x := range helppro {
									res += fmt.Sprintf("\n  %v %s", Data.Logobot, x)
								}
								client.SendHelp(to, res)
							} else {
								k := getKey(kata)
								det, anu := details[k]
								tt := fmt.Sprintf(det, used, k)
								if anu {
									newsend += tt
								} else {
									newsend += "Not found any command's that's have."
								}
							}
						}
					}
				}
			} else if cmd == "help" {
				if GetCodeprem(9, sender, to) {
					res := "✠ 𝗠𝗲𝗻𝘂 𝗠𝗲𝘀𝘀𝗮𝗴𝗲:"
					res += fmt.Sprintf("\n %v ᴀᴄᴄᴇꜱꜱ",Data.Logobot)
					res += fmt.Sprintf("\n %v ɢᴀᴄᴄᴇꜱꜱ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ (ᴛʏᴘᴇ ᴘᴜʙʟɪᴄ)",Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗔𝗹𝗹:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴀʟʟ",Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗣𝗿𝗼𝘁𝗲𝗰𝘁:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴘʀᴏᴛᴇᴄᴛ",Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗔𝗰𝗰𝗲𝘀𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴅᴇᴠ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴄʀᴇᴀᴛᴏʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴍᴀᴋᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ꜱᴇʟʟᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ʙᴜʏᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴏᴡɴᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴍᴀꜱᴛᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴀᴅᴍɪɴ",Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗚𝗔𝗰𝗰𝗲𝘀𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ɢᴏᴡɴᴇʀ",Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ɢᴀᴅᴍɪɴ",Data.Logobot)
					res += "\n\n✠ 𝗘𝘅𝗮𝗺𝗽𝗹𝗲 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ (ᴄᴏᴍᴍᴀɴᴅ)",Data.Logobot)
					res += "\n\n✠ 𝗣𝗿𝗲𝗳𝗶𝘅:"
					res += fmt.Sprintf("\n %v ꜱɴᴀᴍᴇ : %v",Data.Logobot, Sname)
					res += fmt.Sprintf("\n %v ʀɴᴀᴍᴇ : %v",Data.Logobot, Rname)
					client.SendHelp(to, fancy(res))
				}
			} else if cmd == "about" {
				rngcmd := GetComs(5, "about")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						GetSquad(client, to)
						var a = 0
						ret := "✠ Set Account:"
						for _, p := range bk {
							a = a + 1
							cokk, _ := p.GetSettings()
							ret += fmt.Sprintf("\n\nBot%v:\n", a)
							if cokk.PrivacyReceiveMessagesFromNotFriend == true {
								ret += "   ✓   Filter\n"
							} else {
								ret += "   ✘   Filter\n"
							}
							if cokk.EmailConfirmationStatus == 3 {
								ret += "   ✓   Email\n"
							} else {
								ret += "   ✘   Email\n"
							}
							if cokk.E2eeEnable == true {
								ret += "   ✓   Lsealing\n"
							} else {
								ret += "   ✘   Lsealing\n"
							}
							if cokk.PrivacyAllowSecondaryDeviceLogin == true {
								ret += "   ✓   Secondary\n"
							} else {
								ret += "   ✘   Secondary\n"
							}
						}
						client.SendMessage(to, fancy(ret+"\n\nDevelolper: \nhttps://line.me/ti/p/~code-bot"))
					}
				}
			} else if cmd == "tagall" || cmd == Commands.Tagall && Commands.Tagall != "" {
				//rngcmd := GetComs(8, "all")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						_, target, _ := client.GetChatList(to)
						targets := []string{}
						for i := range target {
							if !InArray2(checkHaid, target[i]) {
								targets = append(targets, target[i])
							}
						}
						client.SendPollMention(to, "Mentions member:\n", targets)
					}
				}
			} else if cmd == "ftagall" {
				rngcmd := GetComs(8, "tagall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						_, target, _ := client.GetChatList(to)
						targets := []string{}
						for i := range target {
							if !InArray2(checkHaid, target[i]) {
								targets = append(targets, target[i])
							}
						}
						client.FakeMention(to, targets)
					}
				}
			} else if strings.HasPrefix(cmd, "unbot") {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unbot"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, UserBot.Bot)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungban") {
				rngcmd := GetComs(8, "ungban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gban)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cancel") && cmd != "cancelall" {
				rngcmd := GetComs(7, "cancel")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								Setcancelto(exe, to, listuser)
								logAccess(client, to, sender, "cancel", listuser, msg.ToType)
							} else {
								client.SendMessage(to, fancy("Please add another bot that has a ban cancel."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invite") {
				rngcmd := GetComs(7, "invite")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								lists := Setinvitetomsg(exe, to, listuser)
								if len(lists) != 0 {
									Cekbanwhois(client, to, lists)
								}
								logAccess(client, to, sender, "invite", listuser, msg.ToType)
							} else {
								client.SendMessage(to, fancy("Please add another bot that has a ban invite."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "addfriends") {
				rngcmd := GetComs(2, "addfriends")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						x := 18
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								addCon(listuser)
								Checklistaccess(client, to, listuser, x, sender)
							} else {
								client.SendMessage(to, fancy("Please add another bot that has a ban addfriends."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "vkick") && cmd != "kickall" || strings.HasPrefix(cmd, Commands.Kick) && Commands.Kick != "" && cmd != "kickall" {
				rngcmd := GetComs(7, "vkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
									if AutoBan {
										Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								Setkickto(exe, to, listuser)
								Setinvitetomsg(exe, to, listuser)
								Setcancelto(exe, to, listuser)
								//AutoproN = true
								logAccess(client, to, sender, "vkick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, fancy("Please add another bot that has a ban kick."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kick") && cmd != "kickall" || strings.HasPrefix(cmd, Commands.Kick) && Commands.Kick != "" && cmd != "kickall" {
				rngcmd := GetComs(7, "kick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if MemUser(to, lists[i]) && !InArray2(listuser, lists[i]) {
									if AutoBan {
										Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := SelectBot(client, to)
							if exe != nil {
								Setkickto(exe, to, listuser)
								//AutoproN = true
								logAccess(client, to, sender, "kick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, fancy("Please add another bot that has a ban kick."))
							}
						}
					}
				}
			
			} else if strings.HasPrefix(cmd, "ban") && cmd != "bans" {
				rngcmd := GetComs(7, "ban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							Checklistaccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "contact") {
				rngcmd := GetComs(5, "contact")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								client.SendContact(to, i)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bio") {
				rngcmd := GetComs(7, "bio")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendMessage(to, fancy(x.StatusMessage))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "tag") {
				rngcmd := GetComs(7, "tag")
				if GetCodeprem(rngcmd, sender, to) {
					listuser := []string{}
					nCount := 0
					fl := strings.Split(cmd, " ")
					typec := strings.Replace(cmd, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !InArray2(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						client.SendPollMention(to, "Tag Users:", listuser)
					}
				}
			} else if strings.HasPrefix(cmd, "story") {
				rngcmd := GetComs(7, "story")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"getstory","getstory")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cover") {
				rngcmd := GetComs(7, "cover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"cover","cover")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "extracover") {
				rngcmd := GetComs(7, "extracover")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"extracover","extracover")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "image") {
				rngcmd := GetComs(7, "image")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendFoto(to, "https://profile.line-scdn.net/"+x.PictureStatus)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "getvideo") {
				rngcmd := GetComs(7, "getvideo")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendVideoWithURL(to, "https://profile.line-scdn.net/"+x.PictureStatus)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "zoom") {
				rngcmd := GetComs(7, "zoom")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								zom := "https://api.minzteam.xyz/zoom?url="+i+"&apikey="+apikey
								client.SendImageWithURL(to, zom)
							}
						}
					}
				}
			} else if strings.HasPrefix(text, Sname+"nunban ") {
				rngcmd := GetComs(7, "nunban ")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						str := strings.Replace(text, Sname+"nunban ", "", 1)
						nm := []string{}
						target := []string{}
						count := strings.Split(str, ",")
						for c, nmr := range count {
							num, _ := strconv.Atoi(nmr)
							if num > 0 && num <= len(Banned.Banlist) {
								target = append(target, Banned.Banlist[num-1])
								pr, _ := client.GetContact(Banned.Banlist[num-1])
								name := pr.DisplayName
								c += 1
								name = fmt.Sprintf(". %s", name)
								nm = append(nm, name)
							}
						}
						if len(target) == 0 {
							newsend += "User not found.\n"
						} else {
	                                       for _, from := range target {
		                                       if Banned.GetBan(from) {
			                                       Banned.DelBan(from)
						            }
						      }
						}
						stx := strings.Join(nm, "\n")
						newsend += "Unban:\n\n"+stx
					}
				}
			} else if strings.HasPrefix(cmd, "unban") {
				rngcmd := GetComs(7, "unban")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !InArray2(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							checkunbanbots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, err := strconv.Atoi(result[1])
								if err != nil {
									client.SendMessage(to, fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := StripOut(str)
										hapuss := linetcr.Archimed(st, Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											checkunbanbots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "pronote on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProNote {
							newsend += "Already enabled.\n"
						} else {
							room.ProNote = true
							newsend += "Protect Note Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "pronote off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProNote {
							newsend += "Already disabled.\n"
						} else {
							room.ProNote = false
							newsend += "Protect Note Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "progpict on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProPicture {
							newsend += "Already enabled.\n"
						} else {
							room.ProPicture = true
							newsend += "Protect Picture Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "progpict off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProPicture {
							newsend += "Already disabled.\n"
						} else {
							room.ProPicture = false
							newsend += "Protect Picture Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proalbum on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProAlbum {
							newsend += "Already enabled.\n"
						} else {
							room.ProAlbum = true
							newsend += "Protect Album Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proalbum off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProAlbum {
							newsend += "Already disabled.\n"
						} else {
							room.ProAlbum = false
							newsend += "Protect Album Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prolink on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProLink {
							newsend += "Already enabled.\n"
						} else {
							room.ProLink = true
							newsend += "Protect Link Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prolink off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProLink {
							newsend += "Already disabled.\n"
						} else {
							room.ProLink = false
							newsend += "Protect Link Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proflex on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProFlex {
							newsend += "Already enabled.\n"
						} else {
							room.ProFlex = true
							newsend += "Protect Flex Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proflex off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProFlex {
							newsend += "Already disabled.\n"
						} else {
							room.ProFlex = false
							newsend += "Protect Flex Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proimage on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProImage {
							newsend += "Already enabled.\n"
						} else {
							room.ProImage = true
							newsend += "Protect Image Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proimage off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProImage {
							newsend += "Already disabled.\n"
						} else {
							room.ProImage = false
							newsend += "Protect Image Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "provideo on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProVideo {
							newsend += "Already enabled.\n"
						} else {
							room.ProVideo = true
							newsend += "Protect Video Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "provideo off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProVideo {
							newsend += "Already disabled.\n"
						} else {
							room.ProVideo = false
							newsend += "Protect Video Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "procall on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProCall {
							newsend += "Already enabled.\n"
						} else {
							room.ProCall = true
							newsend += "Protect Call Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "procall off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProCall {
							newsend += "Already disabled.\n"
						} else {
							room.ProCall = false
							newsend += "Protect Call Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prospam on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProSpam {
							newsend += "Already enabled.\n"
						} else {
							room.ProSpam = true
							newsend += "Protect Spamcall Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prospam off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProSpam {
							newsend += "Already disabled.\n"
						} else {
							room.ProSpam = false
							newsend += "Protect Spamcall Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prosticker on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProSticker {
							newsend += "Already enabled.\n"
						} else {
							room.ProSticker = true
							newsend += "Protect Sticker Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prosticker off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProSticker {
							newsend += "Already disabled.\n"
						} else {
							room.ProSticker = false
							newsend += "Protect Sticker Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "procontact on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProContact {
							newsend += "Already enabled.\n"
						} else {
							room.ProContact = true
							newsend += "Protect Contact Turn on\n"
					       }
						SaveBackup()
					}
				}
			} else if cmd == "procontact off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProContact {
							newsend += "Already disabled.\n"
						} else {
							room.ProContact = false
							newsend += "Protect Contact Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "propost on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProPost {
							newsend += "Already enabled.\n"
						} else {
							room.ProPost = true
							newsend += "Protect Post Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "propost off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProPost {
							newsend += "Already disabled.\n"
						} else {
							room.ProPost = false
							newsend += "Protect Post Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "profile on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProFile {
							newsend += "Already enabled.\n"
						} else {
							room.ProFile = true
							newsend += "Protect File Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "profile off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProFile {
							newsend += "Already disabled.\n"
						} else {
							room.ProFile = false
							newsend += "Protect File Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prokick on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProKick {
							newsend += "Already enabled.\n"

						} else {
							room.ProKick = true
							newsend += "Protect Kick Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "prokick off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProKick {
							newsend += "Already disabled.\n"
						} else {
							room.ProKick = false
							newsend += "Protect Kick Turn off\n"
						}
						SaveBackup()
					}

				}
			} else if cmd == "announce on" {
				//rngcmd := GetComs(8, "on")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Announce {
							newsend += "Already enabled.\n"
						} else {
							room.Announce = true
							newsend += "Announcement is enabled.\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "announce off" {
				//rngcmd := GetComs(8, "off")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Announce {
							room.Announce = false
							newsend += "Announcement is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proqr on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProQr {
							newsend += "Already enabled.\n"
						} else {
							room.ProQr = true
							newsend += "Protect Qr Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proqr off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProQr {
							newsend += "Already disabled.\n"
						} else {
							room.ProQr = false
							newsend += "Protect Qr Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proinvite on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProInvite {
							newsend += "Already enabled.\n"
						} else {
							room.ProInvite = true
							newsend += "Protect Invite Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "proinvite off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProInvite {
							newsend += "Already disabled.\n"
						} else {
							room.ProInvite = false
							newsend += "Protect Invite Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "automute on" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Automute {
							newsend += "Already enabled.\n"
						} else {
							room.Automute = true
							newsend += "Automute enabled.\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "automute off" {
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Automute {
							newsend += "Already disabled.\n"
						} else {
							room.Automute = false
							newsend += "Automute disabled.\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "procancel on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProCancel {
							newsend += "Already enabled.\n"
						} else {
							room.ProCancel = true
							newsend += "Protect Cancel Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "procancel off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProCancel {
							newsend += "Already disabled.\n"
						} else {
							room.ProCancel = false
							newsend += "Protect Cancel Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "projoin on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProJoin {
							newsend += "Already enabled.\n"
						} else {
							room.ProJoin = true
							newsend += "Protect Join Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "projoin off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProJoin {
							newsend += "Already disabled.\n"
						} else {
							room.ProJoin = false
							newsend += "Protect Join Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "progname on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProName {
							newsend += "Already enabled.\n"
						} else {
							room.ProName = true
							newsend += "Protect Group Name Turn on\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "progname off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProName {
							newsend += "Already disabled.\n"
						} else {
							room.ProName = false
							newsend += "Protect Group Name Turn off\n"
						}
						SaveBackup()
					}
				}
			} else if cmd == "leave on" {
				//rngcmd := GetComs(8, "on")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Leavebool {
							newsend += "Already enabled.\n"
						} else {
							room.Leavebool = true
							newsend += "Leave Message Turn on\n"
						}
					}
				}
			} else if cmd == "sendimage on" {
				//rngcmd := GetComs(4, "sendimage")
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ImageLurk {
							newsend += "Already enabled.\n"
						} else {
							room.ImageLurk = true
							newsend += "Sendimage set enabled.\n"
						}
					}
				}
			} else if cmd == "sendimage off" {
				//rngcmd := GetComs(4, "sendimage")
				if GetCodeprem(5, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ImageLurk {
							newsend += "Already disabled.\n"
						} else {
							room.ImageLurk = false
							newsend += "Sendimage set disabled.\n"
						}
					}
				}
			} else if cmd == "leave off" {
				//rngcmd := GetComs(8, "off")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Leavebool {
							newsend += "Already disabled.\n"
						} else {
							room.Leavebool = false
							newsend += "Leave Message Turn off\n"
						}
					}
				}
			} else if cmd == "welcome on" {
				//rngcmd := GetComs(8, "on")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Welcome {
							newsend += "Already enabled.\n"
						} else {
							room.Welcome = true
							newsend += "Welcome Message Turn on\n"
						}
					}
				}
			} else if cmd == "welcome off" {
				//rngcmd := GetComs(8, "off")
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Welcome {
							newsend += "Already disabled.\n"
						} else {
							room.Welcome = false
							newsend += "Welcome Message Turn off\n"
						}
					}
				}
			} else if cmd == "backuser on" {
				rngcmd := GetComs(5, "backuser")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Backup {
							newsend += "Already enabled.\n"
						} else {
							room.Backup = true
							newsend += "backup user set enabled.\n"
						}
					}
				}
			} else if cmd == "backuser off" {
				rngcmd := GetComs(5, "backuser")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Backup {
							newsend += "Already disabled.\n"
						} else {
							room.Backup = false
							newsend += "backup user set disabled.\n"
						}
					}
				}
			} else if cmd == "hostage on" {
				rngcmd := GetComs(5, "hostage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Backleave {
							newsend += "Already enabled.\n"
						} else {
							room.Backleave = true
							newsend += "hostage set enabled.\n"
						}
					}
				}
			} else if cmd == "hostage off" {
				rngcmd := GetComs(5, "hostage")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.Backleave {
							newsend += "Already disabled.\n"
						} else {
							room.Backleave = false
							newsend += "hostage set disabled.\n"
						}
					}
				}
			} else if cmd == "allprotect off" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProName && !room.ProCancel && !room.ProInvite && !room.ProKick && !room.ProQr && !room.ProJoin && !room.ProPicture && !room.ProNote && !room.ProAlbum && !room.ProLink && !room.ProFlex && !room.ProImage && !room.ProVideo && !room.ProCall && !room.ProSpam && !room.ProSticker && !room.ProContact && !room.ProPost && !room.ProFile && !room.AntiTag {
							newsend += "All protection is Already disabled.\n"
						} else {
						      room.ProCancel = false
						      room.ProInvite = false
						      room.ProKick = false
						      room.ProQr = false
						      room.ProName = false
						      room.ProJoin = false
						      room.ProPicture = false
						      room.ProNote = false
						      room.ProAlbum = false
						      room.ProLink = false
						      room.ProFlex = false
						      room.ProImage = false
						      room.ProVideo = false
						      room.ProCall = false
						      room.ProSpam = false
						      room.ProSticker = false
						      room.ProContact = false
						      room.ProPost = false
						      room.ProFile = false
						      room.AntiTag = false
						      SaveBackup()
						      newsend += "All Protect Turn off\n"
						}
					}
				}
			} else if cmd == "allprotect on" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProName && room.ProCancel && room.ProInvite && room.ProKick && room.ProQr && room.ProJoin && room.ProPicture && room.ProNote && room.ProAlbum && room.ProLink && room.ProFlex && room.ProImage && room.ProVideo && room.ProCall && room.ProSpam && room.ProSticker && room.ProContact && room.ProPost && room.ProFile && room.AntiTag {
							newsend += "All protection is Already enabled.\n"
						} else {
						      room.ProCancel = true
						      room.ProInvite = true
						      room.ProKick = true
						      room.ProQr = true
						      room.ProName = true
						      room.ProJoin = true
						      room.ProPicture = true
						      room.ProNote = true
						      room.ProAlbum = true
						      room.ProLink = true
						      room.ProFlex = true
						      room.ProImage = true
						      room.ProVideo = true
						      room.ProCall = true
						      room.ProSpam = true
						      room.ProSticker = true
						      room.ProContact = true
						      room.ProPost = true
						      room.ProFile = true
						      room.AntiTag = true
						      SaveBackup()
						      newsend += "All Protect Turn on\n"
						}
					}
				}
			} else if cmd == "protectmax on" || cmd == Commands.Max && Commands.Max != "" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.ProName && room.ProCancel && room.ProInvite && room.ProKick && room.ProQr {
							newsend += "Max protection is Already enabled.\n"
						} else {
							room.ProName = true
							room.ProCancel = true
							room.ProInvite = true
							room.ProKick = true
							room.ProQr = true
						      SaveBackup()
							newsend += "Max Protect Turn on\n"
						}
					}
				}
			} else if cmd == "protectmax off" || cmd == Commands.None && Commands.None != "" {
				if GetCodeprem(9, sender, to) {
					if CheckExprd(client, to, sender) {
						if !room.ProName && !room.ProCancel && !room.ProInvite && !room.ProKick && !room.ProQr {
							newsend += "Max protection is Already disabled.\n"
						} else {
							room.ProName = false
							room.ProCancel = false
							room.ProInvite = false
							room.ProKick = false
							room.ProQr = false
						      SaveBackup()
							newsend += "Max Protect Turn off\n"
						}
					}
				}
			} else if cmd == "restartperm" {
				Resprem()
				list := PerCheckList()
				newsend += list
			} else if cmd == "nukejs" {
				rngcmd := GetComs(5, "nukejs")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist, _ := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						max := len(exe) * 100
						lkick := []string{}
						for n, v := range memlist {
							if MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0
						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							go func(to string, i int) {
								target := lkick[i]
								cl := nom[i]
								cl.Nodejs(to, target)
							}(to, i)
						}
						logAccess(client, to, sender, "kickall", lkick, msg.ToType)
					}
				}
			} else if cmd == "kickall" || cmd == Commands.Kickall && Commands.Kickall != "" {
				rngcmd := GetComs(5, "kickall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist, _ := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						max := len(exe) * 100
						lkick := []string{}
						for n, v := range memlist {
							if MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0
						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							go func(to string, i int) {
								target := lkick[i]
								cl := nom[i]
								cl.DeleteOtherFromChats(to, []string{target})
							}(to, i)
						}
						logAccess(client, to, sender, "kickall", lkick, msg.ToType)
					}
				}
			} else if cmd == "cancelall" || cmd == Commands.Cancelall && Commands.Cancelall != "" {
				rngcmd := GetComs(5, "cancelall")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist2, memlist := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist2 {
							if InArray2(Squadlist, mid) {
								cl := GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						lkick := []string{}
						max := len(exe) * 10
						for n, v := range memlist {
							if MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0

						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							target := lkick[i]
							cl := nom[i]
							ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
						}
						logAccess(client, to, sender, "cancelall", lkick, msg.ToType)
					}
				}
			} else if strings.HasPrefix(cmd, "joinqr") {
				rngcmd := GetComs(5, "joinqr")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "joinqr"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, "/")
							tkt := client.FindChatByTicket(result[4])
							client.AcceptTicket(tkt.Chat.ChatMid, result[4])
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									err := p.AcceptTicket(tkt.Chat.ChatMid, result[4])
									if err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								newsend += "Succes Accept Group Ticket"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "nukeqr") {
				rngcmd := GetComs(5, "nukeqr")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						su := "nukeqr"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, "/")
							tkt := client.FindChatByTicket(result[4])
							client.AcceptTicket(tkt.Chat.ChatMid, result[4])
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									err := p.AcceptTicket(tkt.Chat.ChatMid, result[4])
									if err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								go Nukjoin(exe[0], op.CreatedTime, tkt.Chat.ChatMid)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "qrjoin") {
				rngcmd := GetComs(5, "qrjoin")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if strings.Contains(pesan, "/ti/g") {
							regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
							links := regex.FindAllString(msg.Text, -1)
							tickets := []string{}
							for _, link := range links {
								if !InArray2(tickets, link) {
									tickets = append(tickets, link)
								}
							}
							for _, tick := range tickets {
								tuk := strings.Split(tick, "/")
								ntk := len(tuk) - 1
								ti := tuk[ntk]
								tkt := client.FindChatByTicket(ti)
								client.AcceptTicket(tkt.Chat.ChatMid, ti)
								exe := []*linetcr.Account{}
								for _, p := range bk {
									if p.Limited == false {
										err := p.AcceptTicket(tkt.Chat.ChatMid, ti)
										if err == nil {
											exe = append(exe, p)
										}
									}
								}
								if len(exe) != 0 {
									newsend += "Succes Accept Group Link"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "qrjoinkick") {
				rngcmd := GetComs(5, "qrjoinkick")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if strings.Contains(pesan, "/ti/g") {
							regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
							links := regex.FindAllString(msg.Text, -1)
							tickets := []string{}
							for _, link := range links {
								if !InArray2(tickets, link) {
									tickets = append(tickets, link)
								}
							}
							for _, tick := range tickets {
								tuk := strings.Split(tick, "/")
								ntk := len(tuk) - 1
								ti := tuk[ntk]
								tkt := client.FindChatByTicket(ti)
								client.AcceptTicket(tkt.Chat.ChatMid, ti)
								exe := []*linetcr.Account{}
								for _, p := range bk {
									if p.Limited == false {
										err := p.AcceptTicket(tkt.Chat.ChatMid, ti)
										if err == nil {
											exe = append(exe, p)
										}
									}
								}
								if len(exe) != 0 {
									go Nukjoin(exe[0], op.CreatedTime, to)
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "clone") {
				rngcmd := GetComs(5, "clone")
				if GetCodeprem(rngcmd, sender, to) {
					if CheckExprd(client, to, sender) {
						if len(mentionlist) == 1 {
							cok := strings.Split((cmd), " ")
							if len(cok) > 1 {
								targets := ""
								var pp, vp, co, cv, name, stats string
								cok := strings.Split((cmd), " ")
								if len(cok) > 1 {
									ann := cok[1]
									var prof *talkservice.Contact
									if ann == "@me" {
										prof, _ = client.GetContact(msg.From_)
										targets = msg.From_
									}
									if prof != nil {
										name = prof.DisplayName
										stats = prof.StatusMessage
										if prof.VideoProfile != "" {
											ps, err := client.Downloads("http://dl.profile.line-cdn.net"+prof.PicturePath+"/vp", "mp4")
											if err != nil {
												client.SendMessage(to, fancy("Download video profile error."))
											} else {
												vp = ps
											}
										}
										if prof.PicturePath != "" {
											ps, err := client.Downloads("http://dl.profile.line.naver.jp"+prof.PicturePath, "jpg")
											if err != nil {
												client.SendMessage(to, fancy("Download picture profile error."))
											} else {
												pp = ps
											}
										}
										profs := client.GetProfileDetail(msg.From_)
										pss, err := client.Downloads("https://obs.line-scdn.net/r/myhome/c/"+gjson.Get(profs, "result.objectId").String(), "jpg")
										if err == nil {
											co = pss
										}
										pss, err = client.Downloads("https://obs.line-scdn.net/r/myhome/vc/"+gjson.Get(profs, "result.objectId").String(), "mp4")
										if err == nil {
											cv = pss
										}
										if len(mentionlist) != 0 {
											clon := false
											for _, target := range mentionlist {
												if target != targets && InArray2(Squadlist, target) {
													idx := GetKorban(target)
													clone(idx, pp, vp, co, cv, name, stats)
													idx.SendMention(to, "Cloning @! profile done.", []string{targets})
													clon = true
												}
											}
											if !clon {
												if pp != "" {
													os.Remove(pp)
												}
												if vp != "" {
													os.Remove(vp)
												}
												if co != "" {
													os.Remove(co)
												}
												if cv != "" {
													os.Remove(cv)
												}
											}
										} else {
											if pp != "" {
												os.Remove(pp)
											}
											if vp != "" {
												os.Remove(vp)
											}
											if co != "" {
												os.Remove(co)
											}
											if cv != "" {
												os.Remove(cv)
											}
										}
									}
								}
							}
						}
					}
				}
			} else if PublicMode {
				 if pesan == "help" {
					res := "✠ 𝗣𝘂𝗯𝗹𝗶𝗰 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
					res += "\n"
					for _, x := range helppublic {
						res += fmt.Sprintf("\n %v %s", Data.Logobot, x)
					}
					res += "\n\n✠ 𝗧𝘆𝗽𝗲 𝗥𝗲𝗽𝗹𝘆:"
					res += "\n"
					for _, x := range helpreply {
						res += fmt.Sprintf("\n %v %s", Data.Logobot, x)
					}
					client.SendHelp(to, res)
				} else if pesan == "mentionall" {
					_, target, _ := client.GetChatList(to)
					targets := []string{}
					for i := range target {
						if !InArray2(checkHaid, target[i]) {
							targets = append(targets, target[i])
						}
					}
					client.SendPollMention(to, "Mentions member:\n", targets)
				} else if pesan == "groupinfo" {
					list := InfoGroup(client, to)
					client.SendMessage(to, fancy(list))
				} else if strings.HasPrefix(pesan, "detectcall ") {
					spl := strings.Replace(pesan, "detectcall ", "", 1)
					if spl == "on" {
						DetectCall = true
						newsend += "Detectgroupcall is enabled.\n"
					} else if spl == "off" {
						DetectCall = false
						newsend += "Detectgroupcall is disabled.\n"
					}
				} else if pesan == "welcome on" {
					if room.Welcome {
						newsend += "Already enabled.\n"
					} else {
						room.Welcome = true
						newsend += "Welcome Message Turn on\n"
					}
				} else if pesan == "welcome off" {
					if !room.Welcome {
						newsend += "Already disabled.\n"
					} else {
						room.Welcome = false
						newsend += "Welcome Message Turn off\n"
					}
				} else if pesan == "lurk on" {
					room.Lurk = true
					room.NameLurk = "mention"
					room.Userlurk = []string{}
					newsend += "Lurking enabled.\n"
				} else if pesan == "lurk off" {
					room.Lurk = false
					room.ImageLurk = false
					if len(room.Userlurk) != 0 {
						list := " ✠ Lurkers ✠ \n"
						for num, xd := range room.Userlurk {
							num++
							rengs := strconv.Itoa(num)
							new := client.Getcontactuser(xd)
							if new != nil {
								list += "\n   " + rengs + ". Closed Account"
							} else {
								x, _ := client.GetContact(xd)
								list += "\n   " + rengs + ". " + x.DisplayName
							}
						}
						newsend += list + "\n"
					}
					room.Userlurk = []string{}
				} else if strings.HasPrefix(pesan, "getpict") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.SendFoto(to, "https://profile.line-scdn.net/"+x.PictureStatus)
						}
					}
				} else if strings.HasPrefix(pesan, "getcontact") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							client.SendContact(to, i)
						}
					}
				} else if strings.HasPrefix(pesan, "getbio") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.SendMessage(to, fancy(x.StatusMessage))
						}
					}
				} else if strings.HasPrefix(pesan, "getname") {
					listuser := []string{}
					nCount := 0
					x := 16
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !InArray2(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						Checklistaccess(client, to, listuser, x, sender)
					}
				} else if strings.HasPrefix(pesan, "getmid") {
					listuser := []string{}
					nCount := 0
					x := 14
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !InArray2(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						Checklistaccess(client, to, listuser, x, sender)
					}
				} else if strings.HasPrefix(pesan, "getcover") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"cover","cover")
						}
					}
				} else if strings.HasPrefix(pesan, "gextracover") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"extracover","extracover")
						}
					}
				} else if strings.HasPrefix(pesan, "getstory") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"getstory","getstory")
						}
					}
				} else if pesan == "getcall" {
                    		         if msg.ToType == 2 {
                        		        gcall, _ := call.GetGroupCall(to, client.AuthToken)
                                         Room := linetcr.GetRoom(to)
                        		        res := "Get Call Group:"
                        		        if gcall.MediaType == 1 {
                                                res += "\n  • Type: Audio Call"
                        		               res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                loc, _ := time.LoadLocation("Asia/Jakarta")
	                                         a := time.Now().In(loc)
	                                         yyyy := strconv.Itoa(a.Year())
	                                         MM := a.Month().String()
	                                         dd := strconv.Itoa(a.Day())
	                                         Date := dd + "-" + MM + "-" + yyyy
                                                cok := gcall.Started / 1000
			                            i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                            tm := time.Unix(i, 0)
			                            ss := time.Since(tm)
			                            sp := fmtDuration(ss)
                                                res += "\n  • Date: "+Date
                                                res += "\n  • Started: "+sp
                        		               res += "\n  • Members:"
                                                mmk := gcall.MemberMids
						        if len(mmk) != 0 {
							        for num, xd := range mmk {
								        num++
								        rengs := strconv.Itoa(num)
								        x, _ := client.GetContact(xd)
								        res += "\n      " + rengs + ". "+x.DisplayName
							          }
					                 }
                                                       client.SendMessage(to, fancy(res))
                                         }
                        		        if gcall.MediaType == 2 {
                                                res += "\n  • Type: Video Call"
                        		               res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                loc, _ := time.LoadLocation("Asia/Jakarta")
	                                         a := time.Now().In(loc)
	                                         yyyy := strconv.Itoa(a.Year())
	                                         MM := a.Month().String()
	                                         dd := strconv.Itoa(a.Day())
	                                         Date := dd + "-" + MM + "-" + yyyy
                                                cok := gcall.Started / 1000
			                            i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                            tm := time.Unix(i, 0)
			                            ss := time.Since(tm)
			                            sp := fmtDuration(ss)
                                                res += "\n  • Date: "+Date
                                                res += "\n  • Started: "+sp
                        		               res += "\n  • Members:"
                                                mmk := gcall.MemberMids
						        if len(mmk) != 0 {
							        for num, xd := range mmk {
								        num++
								        rengs := strconv.Itoa(num)
								        x, _ := client.GetContact(xd)
								        res += "\n      " + rengs + ". "+x.DisplayName
							          }
					                 }
                                                   client.SendMessage(to, fancy(res))
                                           }
					  }
				} else if len(Data.WordbanBack) != 0 {
					for _, selftcr := range Data.WordbanBack {
						if pesan == selftcr {
							if MemUser(msg.To, msg.From_) {
								exe, _ := SelectBot(client, to)
								if exe != nil {
									Setkickto(exe, to, []string{msg.From_})
									Banned.AddBan(msg.From_)
								}
							}
						}					
					}
				}
			}
		}
		if newsend != "" {
			client.SendMessage(to, fancy(newsend))
		}
	}
}
