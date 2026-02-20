# LINE Bot API

REST API server for LINE moderation commands.

## Quick Start

### 1. Run the API Server

```bash
cd /var/services/homes/kanom/gobots/api
python3 api.py
```

Server will start on `http://0.0.0.0:8080`

### 2. Test the API

```bash
# Get API status
curl http://localhost:8080/

# Get group info
curl "http://localhost:8080/api/ginfo?group_id=C123456789"

# Kick user
curl -X POST http://localhost:8080/api/kick \
  -H "Content-Type: application/json" \
  -d '{"group_id": "C123456789", "user_mid": "U123456789"}'

# Ban user
curl -X POST http://localhost:8080/api/ban \
  -H "Content-Type: application/json" \
  -d '{"group_id": "C123456789", "user_mid": "U123456789"}'

# Unban user
curl -X POST http://localhost:8080/api/unban \
  -H "Content-Type: application/json" \
  -d '{"group_id": "C123456789", "user_mid": "U123456789"}'

# Cancel invites
curl -X POST http://localhost:8080/api/cancel \
  -H "Content-Type: application/json" \
  -d '{"group_id": "C123456789"}'
```

## OpenClaw Integration

In OpenClaw, you can call these APIs using the `exec` tool:

```
curl -X POST http://localhost:8080/api/kick \
  -H "Content-Type: application/json" \
  -d '{"group_id": "Cxxx", "user_mid": "Uxxx"}'
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | API status |
| GET | `/api/status` | Bot status |
| GET | `/api/ginfo?group_id=` | Group info |
| POST | `/api/kick` | Kick user |
| POST | `/api/ban` | Ban user |
| POST | `/api/unban` | Unban user |
| POST | `/api/cancel` | Cancel invites |

## Configuration

Edit `api.py` to change:
- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
