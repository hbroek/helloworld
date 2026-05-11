# 📦 Render.com Deployment Guide

Render is a Platform-as-a-Service (PaaS) that makes it easy to deploy Go applications.

## 🎯 Quick Start (5 Steps)

### Step 1: Prepare Repository

1. **Add Production Build Script:**

```bash
cd /Users/henry/develop/pi-test/backend

# Create build script
cat > build.sh << 'BUILDSCRIPT'
#!/bin/bash
# Build with target platform for Render
GOOS=linux GOARCH=amd64 go build -o ../frontend_server .
echo "Build complete: frontend_server"
ls -lh ../frontend_server
BUILDSCRIPT

chmod +x build.sh
```

2. **Verify Build:**
```bash
./build.sh
```

### Step 2: Render Account Setup

1. Go to [render.com](https://render.com)
2. Sign up with GitHub account
3. Navigate to **"New +" → "Web Service"**

### Step 3: Connect Repository

In Render Dashboard:
1. Click **"Connect a repository"**
2. Select `hbroek/helloworld`
3. Click **"Create Web Service"**

### Step 4: Configure Build & Runtime

In Render's Web Service settings:

| Setting | Value |
|---------|-------|
| **Name** | `helloworld-api` |
| **Root Directory** | `/` |
| **Build Command** | `go mod download && go build -o frontend_server .` |
| **Start Command** | `./frontend_server` |
| **Instance Type** | Free tier (Standard) |

### Step 5: Environment Variables

Add these in Render settings:

```
PORT=4242
GOOS=linux
GOARCH=amd64

# Optional - enable CORS
CORS_ORIGIN=*
```

### Step 6: Deploy!

Click **"Create Web Service"** and wait ~2-3 minutes.

### 🎉 Get Your URL

After deployment:
```
https://helloworld-api.onrender.com
```

Render provides a free subdomain!

---

## 🔧 Detailed Deployment Process

### Option A: Full Automatic Deployment

1. **On Render Dashboard:**
   - Connect GitHub repo
   - Select branch (default: `main`)
   - Click **"Create Web Service"**

2. **Render auto-detects Go app**
   - Uses `go build` by default
   - Creates `frontend_server` binary
   - Hosts it at your URL

### Option B: Manual Steps (if needed)

If Render doesn't auto-detect:

1. **Add Procfile:**
```bash
# Create Procfile in repo root
cat > Procfile << 'PROCFILE'
web: ./frontend_server
PROCFILE
```

2. **In Render settings:**
   - Start Command: `./frontend_server`
   - No build command needed if binary included

---

## 📁 Required Repository Structure

For Render deployment, ensure your repo contains:

```
├── backend/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── handlers.go
│   ├── names.go
│   ├── random.go
│   ├── server.go
│   └── main_test.go
├── frontend/
│   └── www/
│       ├── index.html
│       └── helloworld_logo.svg
├── .env.example
├── Procfile
└── build.sh (optional)
```

---

## 🌐 What Gets Hosted

| Resource | Endpoint | Accessible? |
|----------|----------|-------------|
| Health Check | `/health` | ✅ Yes |
| Name Generator | `/api/v1/name-generator` | ✅ Yes |
| Main Page | `/` | ✅ Yes |
| Logo | `/helloworld_logo.svg` | ✅ Yes |

---

## 🚨 Important Considerations

### Render Free Tier Limitations:
- **Shuts down** after 15 mins of inactivity
- **Cold start** takes 30-60 seconds
- **3 hours/month** runtime on free tier
- **512 MB RAM**, 1 CPU core
- **5 build minutes/month** free

### Paid Render (Pro):
- Always on
- 750 hours/month
- Custom domains
- Better CPU/RAM
- **$7.50/month**

### Free Alternatives:
1. **Railway.app** - No credit card, easy deployment
2. **Fly.io** - Free tier with constraints
3. **Oracle Cloud Free** - Very generous free tier
4. **Google Cloud Run** - First 65,000 GB-seconds free/month

---

## 📝 Example Render Configuration

### environment.json (Render environment variables):

```json
{
  "env": {
    "PORT": "4242",
    "GOOS": "linux",
    "GOARCH": "amd64",
    "CORS_ENABLED": "false"
  },
  "secrets": [
    "DATABASE_URL" // if you use DB
  ]
}
```

---

## ✅ Deployment Checklist

Before deploying to Render:

- [ ] Build binary works locally
- [ ] `go.mod` and `go.sum` present
- [ ] `.env.example` created
- [ ] `.gitignore` excludes `frontend_server` (or not included)
- [ ] Readme describes environment variables
- [ ] Tests passing
- [ ] CORS configured if needed
- [ ] Health check endpoint verified

---

## 🎯 Alternative: Static Site Hosting

If you want to deploy static HTML on GitHub Pages:

1. Build only the frontend (`frontend/www/`)
2. Host on GitHub Pages
3. Call your backend API separately (e.g., Render + API backend)
4. Or use a **mock API** with hardcoded names

---

## 🔗 Next Steps

After deploying to Render:

1. **Test endpoints:**
   ```bash
   curl https://your-app.onrender.com/health
   curl https://your-app.onrender.com/api/v1/name-generator
   ```

2. **Share URL** with friends
3. **Add to website** portfolio
4. **Configure custom domain** (with Pro plan)

---

## 🆘 Troubleshooting

### Build Fails?
- Check `go.mod` is valid
- Run `go build` locally first
- Try reducing dependencies

### App Crashes on Startup?
- Check `START` command in Render settings
- Verify binary exists at startup path
- Check logs in Render dashboard

### CORS Issues?
- Add Render's domain to allowed origins
- Or use browser testing only

---

## 📚 Resources

- [Render Go Docs](https://render.com/docs/go)
- [Render Free Tier](https://render.com/free)
- [Render Pricing](https://render.com/pricing)

---

**Total Deployment Time:** 5-10 minutes
**Estimated Cost:** Free ($0/month) or $7.50/month (Pro)

