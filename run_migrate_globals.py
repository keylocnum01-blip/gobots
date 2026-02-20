import os
import re

# Mappings: Old variable name -> New variable name (with package prefix)
mappings = {
    "filterWar": "botstate.FilterWar",
    "Qrwar": "botstate.Qrwar",
    "botleave": "botstate.Botleave",
    "Laststicker": "botstate.Laststicker",
    "MsgRespon": "botstate.MsgRespon",
    "TimeBc": "botstate.TimeBc",
    "MsgBroadcast": "botstate.MsgBroadcast",
    "TimeBroadcast": "botstate.TimeBroadcast",
    "AutoBc": "botstate.AutoBc",
    "AutoBackBot": "botstate.AutoBackBot",
    "Timebk": "botstate.Timebk",
    "MsgLock": "botstate.MsgLock",
    "MsgBan": "botstate.MsgBan",
    "MsFresh": "botstate.MsFresh",
    "MsLimit": "botstate.MsLimit",
    "MsSname": "botstate.MsSname",
    "MsRname": "botstate.MsRname",
    "AllCheng": "botstate.AllCheng",
    "Lastleave": "botstate.Lastleave",
    "ChangPict": "botstate.ChangPict",
    "ChangName": "botstate.ChangName",
    "ProtectMode": "botstate.ProtectMode",
    "AutokickBan": "botstate.AutokickBan",
    "ChangVpict": "botstate.ChangVpict",
    "ChangVcover": "botstate.ChangVcover",
    "ChangeBio": "botstate.ChangeBio",
    "CmdHelper": "botstate.CmdHelper",
    "cewel": "botstate.Cewel",
    "cleave": "botstate.Cleave",
    "ScanTarget": "botstate.ScanTarget",
    "PowerMode": "botstate.PowerMode",
    "PublicMode": "botstate.PublicMode",
    "LockMode": "botstate.LockMode",
    "NukeJoin": "botstate.NukeJoin",
    "AutoBan": "botstate.AutoBan",
    "KickBanQr": "botstate.KickBanQr",
    "canceljoin": "botstate.Canceljoin",
    "Autojoin": "botstate.Autojoin",
    "Ajsjoin": "botstate.Ajsjoin",
    "backlist": "botstate.Backlist",
    "cekoptime": "botstate.Cekoptime",
    "Ceknuke": "botstate.Ceknuke",
    "Commands": "botstate.Commands",
    "Waitlistin": "botstate.Waitlistin",
    "AutoproN": "botstate.AutoproN",
    "LogMode": "botstate.LogMode",
    "DetectCall": "botstate.DetectCall",
    "AutoLike": "botstate.AutoLike",
    "BomLike": "botstate.BomLike",
    "MediaDl": "botstate.MediaDl",
    "LogGroup": "botstate.LogGroup",
    "delayed": "botstate.Delayed",
    "MsgBio": "botstate.MsgBio",
    "MsgName": "botstate.MsgName",
    "FixedToken": "botstate.FixedToken",
    "From_Token": "botstate.From_Token",
    "Group_Token": "botstate.Group_Token",
    "StartChangeImg": "botstate.StartChangeImg",
    "StartChangevImg": "botstate.StartChangevImg",
    "StartChangevImg2": "botstate.StartChangevImg2",
    "AutoPro": "botstate.AutoPro",
    "Command": "botstate.Command",
    "tempginv": "botstate.Tempginv",
    "remotegrupidto": "botstate.Remotegrupidto",
    "ModeBackup": "botstate.ModeBackup",
    "checkHaid": "botstate.CheckHaid",
    "botStart": "botstate.BotStart",
    "TimeBackup": "botstate.TimeBackup",
    "oplist": "botstate.Oplist",
    "oplistinvite": "botstate.Oplistinvite",
    "PurgeOP": "botstate.PurgeOP",
    "oplistjoin": "botstate.Oplistjoin",
    "AutoPurge": "botstate.AutoPurge",
    "BcImage": "botstate.BcImage",
    "StartBc": "botstate.StartBc",
    "BcVideo": "botstate.BcVideo",
    "StartBcV": "botstate.StartBcV",
    "GBcImage": "botstate.GBcImage",
    "GStartBc": "botstate.GStartBc",
    "GBcVideo": "botstate.GBcVideo",
    "GStartBcV": "botstate.GStartBcV",
    "FBcImage": "botstate.FBcImage",
    "FStartBc": "botstate.FStartBc",
    "FBcVideo": "botstate.FBcVideo",
    "FStartBcV": "botstate.FStartBcV",
    "SAVEBcImage": "botstate.SAVEBcImage",
    "startSaveBc": "botstate.StartSaveBc",
    "ClientMid": "botstate.ClientMid",
    "Squadlist": "botstate.Squadlist",
    "argsRaw": "botstate.ArgsRaw",
    "Sinderremote": "botstate.Sinderremote",
    "StartChangeVideo": "botstate.StartChangeVideo",
    "tempgroup": "botstate.Tempgroup",
    "Lastinvite": "botstate.Lastinvite",
    "Lastkick": "botstate.Lastkick",
    "Lastjoin": "botstate.Lastjoin",
    "Lastcancel": "botstate.Lastcancel",
    "Lastupdate": "botstate.Lastupdate",
    "Lastmid": "botstate.Lastmid",
    "filterop": "botstate.Filterop",
    "Lasttag": "botstate.Lasttag",
    "Lastcon": "botstate.Lastcon",
    "Lastmessage": "botstate.Lastmessage",
    "Commandss": "botstate.Commandss",
    "Detectjoin": "botstate.Detectjoin",
    "Banned": "botstate.Banned",
    "timeSend": "botstate.TimeSend",
    "opkick": "botstate.Opkick",
    "opjoin": "botstate.Opjoin",
    "Cekpurge": "botstate.Cekpurge",
    "MaxCancel": "botstate.MaxCancel",
    "MaxKick": "botstate.MaxKick",
    "MaxInvite": "botstate.MaxInvite",
    "CancelPend": "botstate.CancelPend",
    "CountSpam": "botstate.CountSpam",
    "CountAjs": "botstate.CountAjs",
    "AllowDoOnce": "botstate.AllowDoOnce",
    "LockAjs": "botstate.LockAjs",
    "UpdatePicture": "botstate.UpdatePicture",
    "UpdateCover": "botstate.UpdateCover",
    "UpdateVProfile": "botstate.UpdateVProfile",
    "UpdateVCover": "botstate.UpdateVCover",
    "ColorCyan": "botstate.ColorCyan",
    "ColorReset": "botstate.ColorReset",
    "Data": "botstate.Data",
    "remotegrupid": "botstate.Remotegrupid",
    "LastActive": "botstate.LastActive",
    "used": "botstate.Used",
    "IPServer": "botstate.IPServer",
    "Killmode": "botstate.Killmode",
    "Typebc": "botstate.Typebc",
    "AutoJointicket": "botstate.AutoJointicket",
    "TypeJoin": "botstate.TypeJoin",
    "AutoTranslate": "botstate.AutoTranslate",
    "TypeTrans": "botstate.TypeTrans",
    "Opinvite": "botstate.Opinvite",
    "cekGo": "botstate.CekGo",
    "filtermsg": "botstate.FilterMsg",
    "Nkick": "botstate.Nkick",
    "UserBot": "botstate.UserBot",
    "Cekstaybot": "botstate.Cekstaybot",
    
    # New Mappings
    "getStickerKick": "botstate.GetStickerKick",
    "stkid": "botstate.Stkid",
    "stkpkgid": "botstate.Stkpkgid",
    "getStickerRespon": "botstate.GetStickerRespon",
    "stkid2": "botstate.Stkid2",
    "stkpkgid2": "botstate.Stkpkgid2",
    "getStickerStayall": "botstate.GetStickerStayall",
    "stkid3": "botstate.Stkid3",
    "stkpkgid3": "botstate.Stkpkgid3",
    "getStickerLeave": "botstate.GetStickerLeave",
    "stkid4": "botstate.Stkid4",
    "stkpkgid4": "botstate.Stkpkgid4",
    "getStickerKickall": "botstate.GetStickerKickall",
    "stkid5": "botstate.Stkid5",
    "stkpkgid5": "botstate.Stkpkgid5",
    "getStickerBypass": "botstate.GetStickerBypass",
    "stkid6": "botstate.Stkid6",
    "stkpkgid6": "botstate.Stkpkgid6",
    "getStickerInvite": "botstate.GetStickerInvite",
    "stkid7": "botstate.Stkid7",
    "stkpkgid7": "botstate.Stkpkgid7",
    "getStickerClearban": "botstate.GetStickerClearban",
    "stkid8": "botstate.Stkid8",
    "stkpkgid8": "botstate.Stkpkgid8",
    "getStickerCancelall": "botstate.GetStickerCancelall",
    "stkid9": "botstate.Stkid9",
    "stkpkgid9": "botstate.Stkpkgid9",
    "fancy": "botstate.Fancy",
    "GO": "botstate.GO",
    "Whitelist": "botstate.Whitelist",
    "SetHelper": "botstate.SetHelper",
    "DB": "botstate.DB",
    "ClientBot": "botstate.ClientBot",
    "Midlist": "botstate.Midlist",
    "aclear": "botstate.Aclear",
    "Grupas": "botstate.Grupas",
    "Poll": "botstate.Poll",
    "Self": "botstate.Self",
    "cpu": "botstate.Cpu",
    "err": "botstate.Err",
    "changepic": "botstate.Changepic",
    "timeabort": "botstate.Timeabort",
    "TimeSave": "botstate.TimeSave",
    "TimeClear": "botstate.TimeClear",
    "ChangCover": "botstate.ChangCover",
    "stringToInt": "botstate.StringToInt",
    "DATABASE": "botstate.DATABASE",
    "CREATOR": "botstate.CREATOR",
    "DEVELOPER": "botstate.DEVELOPER",
    "TeamNotif": "botstate.TeamNotif",
    "AntiJs": "botstate.AntiJs",
    "MidBc": "botstate.MidBc",
    "MidRemote": "botstate.MidRemote",
    "RemoteOwner": "botstate.RemoteOwner",
    "RemoteMaster": "botstate.RemoteMaster",
    "RemoteAdmin": "botstate.RemoteAdmin",
    "RemoteContact": "botstate.RemoteContact",
    "RemoteBan": "botstate.RemoteBan",
    "HostName": "botstate.HostName",
    "carierMap": "botstate.CarierMap",
    "helppublic": "botstate.Helppublic",
    "helpreply": "botstate.Helpreply",
    "helppro": "botstate.Helppro",
    "ListIp": "botstate.ListIp",
    "helpdeveloper": "botstate.Helpdeveloper",
    "helpcreator": "botstate.Helpcreator",
    "helpmaker": "botstate.Helpmaker",
    "helpseller": "botstate.Helpseller",
    "helpbuyer": "botstate.Helpbuyer",
    "helpowner": "botstate.Helpowner",
    "helpmaster": "botstate.Helpmaster",
    "helpadmin": "botstate.Helpadmin",
    "helpgowner": "botstate.Helpgowner",
    "helpgadmin": "botstate.Helpgadmin",
    "details": "botstate.Details",
}

# Files to process
files_to_process = [
    "main.go",
    "handler/system.go",
    "handler/helpers.go",
    "handler/info.go",
    "handler/actions.go",
    "handler/background.go",
    "handler/moderation.go",
    "handler/runbot.go",
    "handler/invite.go",
    "handler/contact.go",
    "handler/logging.go",
]

def process_file(file_path):
    if not os.path.exists(file_path):
        print(f"File not found: {file_path}")
        return

    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    original_content = content
    
    # Sort mappings by key length descending to avoid partial matches
    sorted_mappings = sorted(mappings.items(), key=lambda x: len(x[0]), reverse=True)

    # 1. Update references
    for old, new in sorted_mappings:
        # Look for word boundary, OldName, word boundary.
        # Negative lookbehind (?<!\.) to ensure it's not a field access.
        # Negative lookbehind (?<!botstate\.) to ensure it's not already migrated.
        
        pattern = r'(?<!\.)(?<!botstate\.)\b' + re.escape(old) + r'\b'
        content = re.sub(pattern, new, content)

    if content != original_content:
        # Add import botstate if missing and we added botstate references
        if "botstate." in content and '"./botstate"' not in content and '"../botstate"' not in content:
             if "package main" in content:
                 content = content.replace('import (', 'import (\n\t"./botstate"')
             else:
                 content = content.replace('import (', 'import (\n\t"../botstate"')
        
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print(f"Updated references in {file_path}")

def remove_declarations_main():
    file_path = "main.go"
    if not os.path.exists(file_path):
        return

    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    new_lines = []
    # We check against the NEW name because references have been updated in process_file
    vars_to_remove = list(mappings.values())
    
    for line in lines:
        stripped = line.strip()
        remove = False
        for var_full in vars_to_remove:
            # var_full is like "botstate.Botleave"
            # Check if line looks like "var botstate.Botleave = ..." or "botstate.Botleave = ..."
            if stripped.startswith(f"var {var_full} ") or \
               stripped.startswith(f"{var_full} ") or \
               stripped.startswith(f"{var_full}=") or \
               stripped.startswith(f"var {var_full}="):
                remove = True
                print(f"Removing declaration: {stripped}")
                break
        
        if not remove:
            new_lines.append(line)

    with open(file_path, 'w', encoding='utf-8') as f:
        f.writelines(new_lines)
    print("Removed declarations from main.go")

if __name__ == "__main__":
    for fp in files_to_process:
        process_file(fp)
    
    remove_declarations_main()
