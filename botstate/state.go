package botstate

import (
	"os"
	"time"

	"../config"
	"../library/hashmap"
	"../library/linetcr"
	talkservice "../library/linethrift"
	"../library/unistyle"
	"../utils"
)

var (
	Fancy            = unistyle.Normal
	GetStickerKick      = 0
	Stkid               = "0"
	Stkpkgid            = "0"
	GetStickerRespon    = 0
	Stkid2              = "0"
	Stkpkgid2           = "0"
	GetStickerStayall   = 0
	Stkid3              = "0"
	Stkpkgid3           = "0"
	GetStickerLeave     = 0
	Stkid4              = "0"
	Stkpkgid4           = "0"
	GetStickerKickall   = 0
	Stkid5              = "0"
	Stkpkgid5           = "0"
	GetStickerBypass    = 0
	Stkid6              = "0"
	Stkpkgid6           = "0"
	GetStickerInvite    = 0
	Stkid7              = "0"
	Stkpkgid7           = "0"
	GetStickerClearban  = 0
	Stkid8              = "0"
	Stkpkgid8           = "0"
	GetStickerCancelall = 0
	Stkid9              = "0"
	Stkpkgid9           = "0"
	GO               = utils.GetArg()
	Whitelist        = []string{}
	SetHelper        = &linetcr.Helper{Rngcmd: make(map[string]int)}
	DB               *config.DATA
	ClientBot        []*linetcr.Account
	Midlist          []string
	Aclear           = time.Now()
	Grupas           []*talkservice.Group
	Poll             *linetcr.Account
	Self             *linetcr.Account
	Cpu              int
	Err              error
	Botleave         = &hashmap.HashMap{}
	Changepic        []*linetcr.Account
	Timeabort        = time.Now()
	TimeSave         = time.Now()
	TimeClear        = time.Now()
	Laststicker      = &hashmap.HashMap{}
	ChangCover       = false
	MsgRespon        = "Ready !!!"
	TimeBc           = time.Now()
	MsgBroadcast     = "tester"
	TimeBroadcast    = 1
	AutoBc           = false
	AutoBackBot      = false
	Timebk           = 6
	MsgLock          = "Success Clear %v locklist."
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
	Cewel            = &hashmap.HashMap{}
	Cleave           = &hashmap.HashMap{}
	ScanTarget       = false
	PowerMode        = false
	PublicMode       = true
	LockMode         = false
	NukeJoin         = false
	AutoBan          = true
	KickBanQr        = false
	Canceljoin       = false
	Autojoin         = "off"
	Ajsjoin          = "qr"
	Backlist         = &hashmap.HashMap{}
	Cekoptime        = []int64{}
	Ceknuke          = &hashmap.HashMap{}
	CekGo            = []int64{}
	FilterMsg        = &hashmap.HashMap{}
	Cekstaybot       = &hashmap.HashMap{}
	Commands         = &linetcr.Command{}
	Waitlistin       = map[string][]string{}
	AutoproN         = false
	LogMode          = false
	DetectCall       = false
	AutoLike         = false
	BomLike          = false
	MediaDl          = false
	LogGroup         = ""
	Delayed          = 10 * time.Second
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
	Tempginv         = []string{}
	Remotegrupidto   = ""
	ModeBackup       = "inv"
	CheckHaid        = []string{}
	BotStart         = time.Now()
	TimeBackup       = time.Now()
	Oplist           = []int64{}
	Oplistinvite     = []int64{}
	PurgeOP          = []int64{}
	Oplistjoin       = []int64{}
	AutoPurge        = true
	BcImage          = false
	StartBc          = false
	BcVideo          = false
	StartBcV         = false
	GBcImage         = false
	GStartBc         = false
	GBcVideo         = false
	GStartBcV        = false
	FBcImage         = false
	FStartBc         = false
	FBcVideo         = false
	FStartBcV        = false
	SAVEBcImage      = false
	StartSaveBc      = false
	ClientMid        = map[string]*linetcr.Account{}
	Squadlist        = []string{}
	ArgsRaw          = os.Args
	Sinderremote     = []string{}
	StartChangeVideo = false
	Tempgroup        = []string{}
	Lastinvite       = &hashmap.HashMap{}
	Lastkick         = &hashmap.HashMap{}
	Lastjoin         = &hashmap.HashMap{}
	Lastcancel       = &hashmap.HashMap{}
	Nkick            = &hashmap.HashMap{}
	Lastupdate       = &hashmap.HashMap{}
	Lastmid          = &hashmap.HashMap{}
	Filterop         = &hashmap.HashMap{}
	Lasttag          = &hashmap.HashMap{}
	Lastcon          = &hashmap.HashMap{}
	Lastmessage      = &hashmap.HashMap{}
	Commandss        = &hashmap.HashMap{}
	Detectjoin       = &linetcr.SaveJoin{}
	Banned           = &linetcr.BanUser{Banlist: []string{}, Fucklist: []string{}, Mutelist: []string{}, Exlist: []string{}, Locklist: []string{}}
	UserBot          = &linetcr.Access{Creator: []string{}, Maker: []string{}, Seller: []string{}, Buyer: []string{}, Owner: []string{}, Master: []string{}, Admin: []string{}, Bot: []string{}}
	TimeSend         = []int64{}
	Opkick           = []int64{}
	Opjoin           = []string{}
	Cekpurge         = []int64{}
	MaxCancel        = 4
	MaxKick          = 3
	MaxInvite        = 5
	CancelPend       = 5
	CountSpam        = "5"
	CountAjs         = "2"
	AllowDoOnce      = 0
	LockAjs          = false
	CekGo            = []int64{}
	UpdatePicture    = map[string]bool{}
	UpdateCover      = map[string]bool{}
	UpdateVProfile   = map[string]bool{}
	UpdateVCover     = map[string]bool{}
	Qrwar            = false
	FilterWar        = &config.Kickop{Kick: []string{}, Inv: []string{}, Opinv: []int64{}}
	ColorCyan        = "\033[36m"
	ColorReset       = "\033[0m"
	Data             config.DATA
	Remotegrupid     = ""
	LastActive       = &hashmap.HashMap{}
	Used             = ""
	IPServer         string
	Killmode         = "none"
	Typebc           = "none"
	AutoJointicket   = false
	TypeJoin         = "none"
	AutoTranslate    = false
	TypeTrans        = "EN"
	Filtermsg        = &hashmap.HashMap{}
	Opinvite         = []int64{}
	StringToInt      = []rune("01")
	DATABASE         = "database/" + utils.GetArg() + ".json"
	CREATOR          = []string{"u0950ac5584daf10c380a53085378775d"}
	DEVELOPER        = []string{"u0950ac5584daf10c380a53085378775d", "u0fbc4b49c6194469781badf7d3194284"}
	TeamNotif        = []string{""}
	AntiJs           = []string{}
	MidBc            = []string{}
	MidRemote        = []string{}
	RemoteOwner      = false
	RemoteMaster     = false
	RemoteAdmin      = false
	RemoteContact    = false
	RemoteBan        = false
	HostName         = []string{
		"legy-jp-addr-long",
	}
	CarierMap = map[string]string{
		"IOSIPAD":     "51089, 1-0",
		"IOS":         "51089, 1-0",
		"ANDROIDLITE": "51000",
		"ANDROID":     "51010",
		"CHROMEOS":    "",
		"DESKTOPMAC":  "",
		"DESKTOPWIN":  "",
		"CHANNELCP":   "51010",
	}
	Helppublic = []string{
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
	Helpreply = []string{
		"ɢᴇᴛᴘɪᴄᴛ",
		"ɢᴇᴛɴᴀᴍᴇ",
		"ɢᴇᴛʙɪᴏ",
		"ɢᴇᴛᴄᴏᴠᴇʀ",
		"ɢᴇxᴛʀᴀᴄᴏᴠᴇʀ",
		"ɢᴇᴛꜱᴛᴏʀʏ",
		"ɢᴇᴛᴍɪᴅ",
		"ɢᴇᴛᴄᴏɴᴛᴀᴄᴛ",
	}
	Helppro = []string{
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
	ListIp        = []string{}
	Helpdeveloper = []string{
		"ᴄʀᴇᴀᴛᴏʀ [@/ʟᴀꜱᴛ]",
		"ᴜɴᴄʀᴇᴀᴛᴏʀ [@/ʟᴀꜱᴛ]",
		"ᴄʀᴇᴀᴛᴏʀꜱ",
		"ᴄʟᴇᴀʀᴄʀᴇᴀᴛᴏʀ",
		"ᴇxᴘᴇʟ [@/ʟᴀꜱᴛ]",
		"ʀᴇʙᴏᴏᴛᴠᴘꜱ",
	}
	Helpcreator = []string{
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
	Hosts = "https://api.vhtear.com/"
	Apikey = "senzu"
)
	Helpmaker = []string{
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
	Helpseller = []string{
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
	Helpbuyer = []string{
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
	}
	Helpowner = []string{
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
	Helpmaster = []string{
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
	Helpadmin = []string{
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
	Helpgowner = []string{
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
	Helpgadmin = []string{
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
	Details = map[string]string{
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
