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

Tcr = LINE(TOKEN, appName="IOS\t12.6.1\tiOS\t15.5")
#================MAIN_FLEX==============
def allowLiff():
    liffId = id[0:10]
    url = 'https://access.line.me/dialog/api/permissions'
    data = {'on': ['P', 'CM'], 'off': []}
    headers = {'X-Line-Access': Tcr.authToken, 'X-Line-Application': 'IOS\t12.6.1\tiOS\t15.5', 'X-Line-ChannelId': '{}'.format(liffId), 'Content-Type': 'application/json'}
    requests.post(url, json=data, headers=headers)

def sendTemplate(to, data):
    allowLiff()
    xyz = LiffChatContext(to)
    xyzz = LiffContext(chat=xyz)
    view = LiffViewRequest(id, xyzz)
    token = Tcr.liff.issueLiffView(view)
    url = 'https://api.line.me/message/v3/share'
    headers = {'Content-Type': 'application/json','Authorization': 'Bearer %s' % token.accessToken}
    data = {"messages":[data]}
    requests.post(url, headers=headers, data=json.dumps(data))
#================FLEX_MENU==============
def flexMenu(logoname, text, textcolor):
    data = {
  "type": "bubble",
  "size": "micro",
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/2vrfT6P/hlth-adt-menu-c.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1500:2130",
            "aspectMode": "fit"
          }
        ],
        "width": "167px",
        "position": "absolute"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "image",
            "url": "https://i.ibb.co/th7j0nT/hlth-adt-menu-h.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "260:67.5",
            "aspectMode": "fit"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/mcMLDQk/hlth-adt-menu-anim-h.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1665:420",
            "aspectMode": "fit",
            "position": "absolute",
            "animated": True
          },
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "text",
                "text": logoname,
                "size": "7px",
                "color": textcolor,
                "weight": "bold"
              }
            ],
            "height": "10px",
            "width": "90px",
            "position": "absolute",
            "offsetTop": "12px",
            "offsetStart": "50px",
            "justifyContent": "center",
            "alignItems": "center"
          }
        ],
        "width": "167px",
        "height": "43px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text,
            "size": "9px",
            "color": textcolor,
            "lineSpacing": "2px",
            "wrap": True, 
            "action": 
              {
                "type": "uri",
                "label": "action",
                "uri": "https://line.me/ti/p/~code-bot"
             }
          }
        ],
        "paddingStart": "18px",
        "paddingEnd": "5px",
        "paddingTop": "2px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "image",
            "url": "https://i.ibb.co/ctfkH8f/hlth-adt-menu-f.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "255:55",
            "aspectMode": "cover"
          },
          {
            "type": "image",
            "url": "https://i.ibb.co/rs843XQ/hlth-adt-menu-anim-f.png",
            "align": "center",
            "size": "full",
            "aspectRatio": "1665:207",
            "aspectMode": "cover",
            "animated": True,
            "position": "absolute"
          }
        ],
        "width": "167px",
        "height": "18px"
      }
    ],
    "paddingAll": "0px",
    "height": "100%",
    "backgroundColor": "#000000"
  }
}
    return data

#================FLEX_SEND==============
def sendflexON(to, text):
    textcolor = "#FFFFFF"
    logoname = "GOLANG_FLEX"
    data = flexMenu(logoname, text, textcolor)
    sendTemplate(to,{"type": "flex", "altText": "Golang_SELFTCR™", "contents": data})

sendflexON(to, text)
#=====================================
