# -*- coding: utf-8 -*-
# GOLANG AUTONAME
# TYPE: AUTONAME
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

def generateName():
    kelamin = random.choice(["female"])
    a = requests.get(f"https://story-shack-cdn-v2.glitch.me/generators/indonesian-name-generator/{kelamin}?count=6").json()["data"]
    style = random.choice([x for x in range(3)])

    if style == 0:return a[0]['name'].lower()
    if style == 1:return f"{a[0]['name']} {a[1]['name']}"[0:20].lower()
    if style == 2:return a[0]['name']

Tcr = LINE(TOKEN, appName="IOS\t12.6.1\tiOS\t15.5")
name = generateName()
if len(name) <= 99999999:
    dname = Tcr.getProfile()
    dname.displayName = name
    Tcr.updateProfile(dname)
    print(dname)