#!/usr/bin/env python3
"""
LINE Bot API Server
Provides REST API for LINE moderation commands
Works with gobots backend
"""

import json
import os
import sys
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs
import socketserver

# Configuration
PORT = 8080
HOST = "0.0.0.0"

class APIHandler(BaseHTTPRequestHandler):
    
    def do_GET(self):
        parsed = urlparse(self.path)
        path = parsed.path
        params = parse_qs(parsed.query)
        
        if path == "/":
            self.send_json({"service": "LINE Bot API", "status": "running"})
        elif path == "/api/status":
            self.send_json({"status": "ok", "bot_running": True})
        elif path == "/api/ginfo":
            group_id = params.get("group_id", [""])[0]
            if not group_id:
                self.send_error(400, "group_id required")
                return
            result = get_group_info(group_id)
            self.send_json(result)
        else:
            self.send_error(404, "Not found")
    
    def do_POST(self):
        content_length = int(self.headers.get('Content-Length', 0))
        body = self.rfile.read(content_length).decode('utf-8')
        
        try:
            data = json.loads(body) if body else {}
        except:
            data = {}
        
        parsed = urlparse(self.path)
        path = parsed.path
        
        if path == "/api/kick":
            group_id = data.get("group_id", "")
            user_mid = data.get("user_mid", "")
            if not group_id or not user_mid:
                self.send_error(400, "group_id and user_mid required")
                return
            result = kick_user(group_id, user_mid)
            self.send_json(result)
        elif path == "/api/ban":
            group_id = data.get("group_id", "")
            user_mid = data.get("user_mid", "")
            if not group_id or not user_mid:
                self.send_error(400, "group_id and user_mid required")
                return
            result = ban_user(group_id, user_mid)
            self.send_json(result)
        elif path == "/api/unban":
            group_id = data.get("group_id", "")
            user_mid = data.get("user_mid", "")
            if not group_id or not user_mid:
                self.send_error(400, "group_id and user_mid required")
                return
            result = unban_user(group_id, user_mid)
            self.send_json(result)
        elif path == "/api/cancel":
            group_id = data.get("group_id", "")
            if not group_id:
                self.send_error(400, "group_id required")
                return
            result = cancel_invites(group_id)
            self.send_json(result)
        else:
            self.send_error(404, "Not found")
    
    def send_json(self, data):
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(data, indent=2).encode())
    
    def log_message(self, format, *args):
        print(f"[API] {args[0]}")

# Placeholder functions - integrate with gobots
def kick_user(group_id, user_mid):
    """Kick user from group"""
    # TODO: Integrate with gobots
    return {
        "status": "ok",
        "action": "kick",
        "group_id": group_id,
        "user_mid": user_mid,
        "message": "Kick command sent to gobots"
    }

def ban_user(group_id, user_mid):
    """Ban user from group"""
    # TODO: Integrate with gobots
    return {
        "status": "ok",
        "action": "ban",
        "group_id": group_id,
        "user_mid": user_mid,
        "message": "Ban command sent to gobots"
    }

def unban_user(group_id, user_mid):
    """Unban user"""
    # TODO: Integrate with gobots
    return {
        "status": "ok",
        "action": "unban",
        "group_id": group_id,
        "user_mid": user_mid,
        "message": "Unban command sent to gobots"
    }

def cancel_invites(group_id):
    """Cancel all invites"""
    # TODO: Integrate with gobots
    return {
        "status": "ok",
        "action": "cancel",
        "group_id": group_id,
        "message": "Cancel command sent to gobots"
    }

def get_group_info(group_id):
    """Get group info"""
    # TODO: Integrate with gobots
    return {
        "status": "ok",
        "action": "ginfo",
        "group_id": group_id,
        "name": "Unknown Group",
        "member_count": 0
    }

def main():
    server = HTTPServer((HOST, PORT), APIHandler)
    print(f"LINE Bot API Server running on http://{HOST}:{PORT}")
    print(f"Endpoints:")
    print(f"  GET  /                  - API status")
    print(f"  GET  /api/status        - Bot status")
    print(f"  GET  /api/ginfo?group_id=xxx - Group info")
    print(f"  POST /api/kick          - Kick user")
    print(f"  POST /api/ban           - Ban user")
    print(f"  POST /api/unban         - Unban user")
    print(f"  POST /api/cancel        - Cancel invites")
    print()
    server.serve_forever()

if __name__ == "__main__":
    main()
