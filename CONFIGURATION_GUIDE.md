# 🎯 Quick Configuration Guide

This guide helps you quickly configure LinkinAutomator for your needs.

## 📝 Basic Setup (3 Steps)

### Step 1: Add Your Credentials
```env
LINKEDIN_USERNAME=your_email@example.com
LINKEDIN_PASSWORD=your_password
```

### Step 2: Choose What to Search
```env
SEARCH_KEYWORDS=Software Engineer
SEARCH_LIMIT=20
```

### Step 3: Run It!
```bash
./linkin-automator.exe
```

---

## 🎨 Common Use Cases

### 🔍 **Scenario 1: Finding Software Engineers**
```env
SEARCH_KEYWORDS=Software Engineer
SEARCH_LIMIT=50
CONNECTION_MESSAGE=Hi! I'm impressed by your experience. Let's connect!
```

### 🎯 **Scenario 2: Recruiting Product Managers**
```env
SEARCH_KEYWORDS=Product Manager
SEARCH_LIMIT=30
CONNECTION_MESSAGE=I'd love to connect and discuss product opportunities.
```

### 🌍 **Scenario 3: Location-Specific Search**
```env
SEARCH_KEYWORDS=Data Scientist in San Francisco
SEARCH_LIMIT=25
CONNECTION_MESSAGE=Hey! I'm hiring Data Scientists in SF. Let's connect!
```

### 💼 **Scenario 4: Company-Specific**
```env
SEARCH_KEYWORDS=Engineer at Google
SEARCH_LIMIT=15
CONNECTION_MESSAGE=Hi! I'd love to learn more about your work at Google.
```

---

## ⚙️ Configuration Options Explained

### 🔐 **Authentication**
| Variable | Example | Description |
|----------|---------|-------------|
| `LINKEDIN_USERNAME` | `john@example.com` | Your LinkedIn email |
| `LINKEDIN_PASSWORD` | `SecurePass123!` | Your LinkedIn password |

### 🔍 **Search Settings**
| Variable | Example | Description |
|----------|---------|-------------|
| `SEARCH_KEYWORDS` | `"Product Manager"` | What to search for |
| `SEARCH_LIMIT` | `20` | Max profiles to process |

### 💬 **Messages**
| Variable | Example | Description |
|----------|---------|-------------|
| `CONNECTION_MESSAGE` | `Hi! Let's connect!` | Note with connection request |
| `FOLLOWUP_MESSAGE` | `Thanks for connecting!` | Message after they accept |
| `ENABLE_FOLLOWUP` | `true` / `false` | Send follow-ups or not |

### 🎛️ **Operation Mode**
| Variable | Example | Description |
|----------|---------|-------------|
| `HEADLESS` | `false` | `false` = see browser, `true` = hidden |
| `LOG_LEVEL` | `info` | `debug`, `info`, `warn`, `error` |

### ⏱️ **Timing (Stealth)**
| Variable | Example | Description |
|----------|---------|-------------|
| `MIN_DELAY` | `3` | Min seconds between actions |
| `MAX_DELAY` | `8` | Max seconds between actions |
| `CONNECTION_DELAY_MIN` | `30` | Min seconds between connections |
| `CONNECTION_DELAY_MAX` | `60` | Max seconds between connections |
| `DAILY_LIMIT` | `20` | Max connections per day |

### 🛡️ **Safety**
| Variable | Example | Description |
|----------|---------|-------------|
| `RESPECT_BUSINESS_HOURS` | `true` | Only run 9 AM - 6 PM |
| `ENABLE_STEALTH_EXTRAS` | `true` | Extra human-like behavior |

---

## 💡 Tips for Recruiters

### ✅ **Do's**
- ✅ Use professional, friendly connection messages
- ✅ Keep `SEARCH_LIMIT` under 50 for first runs
- ✅ Set `DAILY_LIMIT` to 20-30 (LinkedIn safe zone)
- ✅ Run during business hours for best results
- ✅ Personalize `CONNECTION_MESSAGE` for better acceptance

### ❌ **Don'ts**
- ❌ Don't set `SEARCH_LIMIT` above 100 (too aggressive)
- ❌ Don't use generic/spammy messages
- ❌ Don't set delays too low (looks like bot)
- ❌ Don't run 24/7 (suspicious pattern)

---

## 🚀 Quick Start Examples

### Example 1: Conservative (Safe)
```env
LINKEDIN_USERNAME=recruiter@company.com
LINKEDIN_PASSWORD=SecurePass123
SEARCH_KEYWORDS=Senior Developer
SEARCH_LIMIT=15
DAILY_LIMIT=15
CONNECTION_MESSAGE=Hi! Impressed by your background. Let's connect!
HEADLESS=false
```

### Example 2: Moderate (Balanced)
```env
LINKEDIN_USERNAME=recruiter@company.com
LINKEDIN_PASSWORD=SecurePass123
SEARCH_KEYWORDS=Product Manager
SEARCH_LIMIT=30
DAILY_LIMIT=25
CONNECTION_MESSAGE=I'd love to connect and discuss opportunities.
ENABLE_FOLLOWUP=true
FOLLOWUP_MESSAGE=Thanks for connecting! Would love to chat about roles at our company.
```

### Example 3: Aggressive (Advanced Users)
```env
LINKEDIN_USERNAME=recruiter@company.com
LINKEDIN_PASSWORD=SecurePass123
SEARCH_KEYWORDS=Full Stack Engineer
SEARCH_LIMIT=50
DAILY_LIMIT=40
MIN_DELAY=2
MAX_DELAY=5
CONNECTION_DELAY_MIN=20
CONNECTION_DELAY_MAX=40
ENABLE_STEALTH_EXTRAS=true
```

---

## 🎓 Understanding the Workflow

The tool automatically does this:

```
1. 🔐 Login to LinkedIn
   ↓
2. 🔍 Search for: [SEARCH_KEYWORDS]
   ↓
3. 📋 Collect up to [SEARCH_LIMIT] profiles
   ↓
4. 🤝 Send connection requests with [CONNECTION_MESSAGE]
   ↓
5. ⏸️  Wait [CONNECTION_DELAY] between each
   ↓
6. 🛑 Stop at [DAILY_LIMIT]
   ↓
7. ✅ Check who accepted
   ↓
8. 💬 Send [FOLLOWUP_MESSAGE] (if enabled)
```

---

## 🔧 Troubleshooting

### Problem: "Login failed"
**Solution:** Check your `LINKEDIN_USERNAME` and `LINKEDIN_PASSWORD`

### Problem: "No profiles found"
**Solution:** Try broader `SEARCH_KEYWORDS` like "Engineer" instead of "Senior ML Engineer at Startup"

### Problem: "Too fast / Getting blocked"
**Solution:** Increase `MIN_DELAY`, `MAX_DELAY`, and `CONNECTION_DELAY_MIN`

### Problem: "Want to see what's happening"
**Solution:** Set `HEADLESS=false` and `LOG_LEVEL=debug`

---

## 📞 Need Help?

Check these files:
- `README.md` - Full documentation
- `FEATURES.md` - Feature list
- `TECHNICAL_WALKTHROUGH.md` - Deep dive

---

**Made with ❤️ by Ritesh Dey**
