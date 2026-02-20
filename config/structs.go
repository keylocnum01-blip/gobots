 package config

import "time"

type (
	Mentions struct {
		MENTIONEES []struct {
			Start string `json:"S"`
			End   string `json:"E"`
			Mid   string `json:"M"`
		} `json:"MENTIONEES"`
	}
	Kickop struct {
		Kick  []string
		Inv   []string
		Opinv []int64
	}
	Emots struct {
		STICON struct {
			RESOURCES []struct {
				PRODUCTID string `json:"productId"`
				STICONID  string `json:"sticonId"`
			} `json:"resources"`
		} `json:"sticon"`
	}
	Stickers struct {
		Id  string
		Pid string
	}
	Clustering struct {
		Mem string
		Tm  int64
		Fr  []string
	}
	DATA struct {
		Authoken           []string             `json:"Authoken"`
		CreatorBack        []string             `json:"CreatorBack"`
		MakerBack          []string             `json:"MakerBack"`
		BuyerBack          []string             `json:"BuyerBack"`
		OwnerBack          []string             `json:"OwnerBack"`
		MasterBack         []string             `json:"MasterBack"`
		AdminBack          []string             `json:"AdminBack"`
		ResponBack         string               `json:"ResponBack"`
		BroadcastBack      string               `json:"BroadcastBack"`
		RnameBack          string               `json:"RnameBack"`
		SnameBack          string               `json:"SnameBack"`
		BotBack            []string             `json:"BotBack"`
		Dalltime           string               `json:"Dalltime"`
		Logobot            string               `json:"Logobot"`
		SellerBack         []string             `json:"SellerBack"`
		BanBack            []string             `json:"BanBack"`
		FuckBack           []string             `json:"FuckBack"`
		LockBack           []string             `json:"LockBack"`
		WordbanBack        []string             `json:"WordbanBack"`
		Limit              string               `json:"MLimit"`
		Fresh              string               `json:"MFfresh"`
		MuteBack           []string             `json:"MuteBack"`
		AnnunceBack        []string             `json:"AnnunceBack"`
		ProNameBack        []string             `json:"ProNameBack"`
		ProPictureBack     []string             `json:"ProPictureBack"`
		ProNoteBack        []string             `json:"ProNoteBack"`
		ProAlbumBack       []string             `json:"ProAlbumBack"`
		ProQrBack          []string             `json:"ProQrBack"`
		ProjoinBack        []string             `json:"ProjoinBack"`
		ProInviteBack      []string             `json:"ProInviteBack"`
		ProCancelBack      []string             `json:"ProCancelBack"`
		ProkickBack        []string             `json:"ProkickBack"`
		GbanBack           map[string][]string  `json:"GbanBack"`
		GadminBack         map[string][]string  `json:"GadminBack"`
		GownerBack         map[string][]string  `json:"GownerBack"`
		TimeBanBack        map[string]time.Time `json:"TimeBanBack"`
		ProLinkBack        []string             `json:"ProLinkBack"`
		ProFlexBack        []string             `json:"ProFlexBack"`
		ProImageBack       []string             `json:"ProImageBack"`
		ProVideoBack       []string             `json:"ProVideoBack"`
		ProCallBack        []string             `json:"ProCallBack"`
		ProSpamBack        []string             `json:"ProSpamBack"`
		ProStickerBack     []string             `json:"ProStickerBack"`
		ProContactBack     []string             `json:"ProContactBack"`
		ProPostBack        []string             `json:"ProPostBack"`
		ProFileBack        []string             `json:"ProFileBack"`
		Kikhistory         int                  `json:"Kikhistory"`
		Invhistory         int                  `json:"Invhistory"`
		Canclhistory       int                  `json:"Canclhistory"`
		Maxkick            int                  `json:"Maxkick"`
		Maxcancel          int                  `json:"Maxcancel"`
		Maxinvite          int                  `json:"Maxinvite"`
		Cancelpend         int                  `json:"Cancelpend"`
		AutoproBack        bool                 `json:"AutoproBack"`
		RestartBack        string               `json:"RestartBack"`
		AutoPurgeBack      bool                 `json:"AutoPurgeBack"`
		ProtectmodeBack    bool                 `json:"ProtectmodeBack"`
		PowermodeBack      bool                 `json:"PowermodeBack"`
		KickbanqrBack      bool                 `json:"KickbanqrBack"`
		MediadlBack        bool                 `json:"MediadlBack"`
		AutolikeBack       bool                 `json:"AutolikeBack"`
		AutobcBack         bool                 `json:"AutobcBack"`
		AutojointicketBack bool                 `json:"AutojointicketBack"`
		AutotranslateBack  bool                 `json:"AutotranslateBack"`
		NukejoinBack       bool                 `json:"NukejoinBack"`
		CanceljoinBack     bool                 `json:"CanceljoinBack"`
		DetectcallBack     bool                 `json:"DetectcallBack"`
		ModebackupBack     string               `json:"ModebackupBack"`
		AutojoinBack       string               `json:"AutojoinBack"`
		AjsjoinBack        string               `json:"AjsjoinBack"`
		TypejoinBack       string               `json:"TypejoinBack"`
		TypebcBack         string               `json:"TypebcBack"`
		TypetransBack      string               `json:"TypetransBack"`

		KickSticker struct {
			Stkid    string `json:"stkid"`
			Stkpkgid string `json:"stkpkgid"`
		} `json:"Command-sticker-kick"`
		ResponSticker struct {
			Stkid2    string `json:"stkid2"`
			Stkpkgid2 string `json:"stkpkgid2"`
		} `json:"Command-sticker-respon"`
		StayallSticker struct {
			Stkid3    string `json:"stkid3"`
			Stkpkgid3 string `json:"stkpkgid3"`
		} `json:"Command-sticker-stayall"`
		LeaveSticker struct {
			Stkid4    string `json:"stkid4"`
			Stkpkgid4 string `json:"stkpkgid4"`
		} `json:"Command-sticker-leave"`
		KickallSticker struct {
			Stkid5    string `json:"stkid5"`
			Stkpkgid5 string `json:"stkpkgid5"`
		} `json:"Command-sticker-kickall"`
		BypassSticker struct {
			Stkid6    string `json:"stkid6"`
			Stkpkgid6 string `json:"stkpkgid6"`
		} `json:"Command-sticker-bypass"`
		InviteSticker struct {
			Stkid7    string `json:"stkid7"`
			Stkpkgid7 string `json:"stkpkgid7"`
		} `json:"Command-sticker-invite"`
		ClearbanSticker struct {
			Stkid8    string `json:"stkid8"`
			Stkpkgid8 string `json:"stkpkgid8"`
		} `json:"Command-sticker-clearban"`
		CancelallSticker struct {
			Stkid9    string `json:"stkid9"`
			Stkpkgid9 string `json:"stkpkgid9"`
		} `json:"Command-sticker-cancelall"`
	}
)
