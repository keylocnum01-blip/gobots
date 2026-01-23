# -*- coding: utf-8 -*-
# GOLANG AUTOPICT
# TYPE: AUTOPICT
# CREATE BY: SELFTCR™
# HELPER: 2023
# ID LINE: code-bot
# WHATSAPP: +6282278984821
#==========IMPORT==========#
from MySplit import MySplit
from linepy import *
import sys
import re, ast, httpx, random, html5lib, requests, os, json

#==========LOGIN_MAIN==========#
TOKEN = sys.argv[1]
Tcr = LINE(TOKEN, appName="IOS\t12.6.1\tiOS\t15.5")
path = Tcr.downloadFileURL('http://3650000.xyz/api?mode=2', 'path')
Tcr.updateProfilePicture(path)
print(path)