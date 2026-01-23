# -*- coding: utf-8 -*-
# GOLANG TEMPLATE
# TYPE: TEMPLATE
# CREATE BY: SELFTCR™
# HELPER: 2023
# ID LINE: code-bot
# WHATSAPP: +6282278984821
#==========IMPORT==========#
from MySplit import MySplit
from linepy import *
from liff.ttypes import LiffChatContext, LiffContext, LiffSquareChatContext, LiffNoneContext, LiffViewRequest
import sys
import re, ast, httpx, random, html5lib, requests, os, json

#==========LOGIN_MAIN==========#
TOKEN = sys.argv[1]
to = sys.argv[2]
text = sys.argv[3]
id = sys.argv[4]
profile = sys.argv[5]

Tcr = LINE(TOKEN, appName="IOS\t12.6.1\tiOS\t15.5")

Tcr.createChat(text, [Tcr.profile.mid])
gids = Tcr.getGroupIdsByName(text)
for gid in gids:
    chat = Tcr.getChats([gid], False, False)
    chat.chats[0].extra.groupExtra.preventedJoinByTicket = False
    Tcr.updateChat(chat.chats[0], 4)