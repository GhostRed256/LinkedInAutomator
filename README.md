<div align="center">

# 🚀 LinkinAutomator

### *Advanced LinkedIn Automation with Stealth Technology*

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)
[![Rod](https://img.shields.io/badge/Powered%20by-Rod-blue?style=for-the-badge)](https://go-rod.github.io/)
[![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)](https://github.com)

*A sophisticated Go-based LinkedIn automation tool showcasing advanced browser automation, anti-detection techniques, and clean architecture.*

[Features](#-features) • [Installation](#-installation) • [Usage](#-usage) • [Architecture](#-architecture) • [Documentation](#-documentation)

---

## 🎥 **Demo Video**

<div align="center">

### Watch LinkinAutomator in Action!

[![LinkinAutomator Demo](https://img.youtube.com/vi/OSeYk-unIsQ/0.jpg)](https://youtu.be/OSeYk-unIsQ)

**Click to watch the full demonstration** ▶️

*See all 8+ stealth techniques working together: Authentication, Smart Search, Bézier Mouse Movement, Human-like Scrolling, and more!*

</div>

</div>

---

## ✨ Features

<table>
<tr>
<td width="50%">

### 🔐 **Authentication System**
- ✅ Environment-based credentials
- ✅ Session cookie persistence
- ✅ 2FA/Captcha detection
- ✅ Graceful failure handling

### 🎯 **Search & Targeting**
- ✅ Keyword-based search
- ✅ Smart pagination
- ✅ Profile URL extraction
- ✅ Duplicate detection

</td>
<td width="50%">

### 🤝 **Connection Management**
- ✅ Automated connection requests
- ✅ Personalized notes
- ✅ Daily limit enforcement
- ✅ Request tracking

### 💬 **Messaging System**
- ✅ Accepted connection detection
- ✅ Auto follow-up messages
- ✅ Template support
- ✅ Comprehensive tracking

</td>
</tr>
</table>

---

## 🛡️ **Advanced Stealth Technology**

<div align="center">

### *8+ Anti-Detection Techniques Implemented*

</div>

| Technique | Implementation | Status |
|-----------|---------------|--------|
| 🖱️ **Bézier Mouse Movement** | Curved trajectories with variable speed | ✅ Active |
| ⏱️ **Randomized Timing** | Human-like delays & pauses | ✅ Active |
| 🎭 **Fingerprint Masking** | User-Agent, viewport, webdriver flags | ✅ Active |
| 📜 **Smart Scrolling** | Variable speed with scroll-back | ✅ Active |
| ⌨️ **Realistic Typing** | Typos, corrections, rhythm variation | ✅ Active |
| 🎯 **Random Hovering** | Natural cursor wandering | ✅ Active |
| 🕐 **Business Hours** | Time-aware operation | ✅ Active |
| 🚦 **Rate Limiting** | Request quotas & throttling | ✅ Active |

---

## 📦 Installation

### Prerequisites

```bash
✓ Go 1.21 or higher
✓ Google Chrome installed
✓ Git
```

### Quick Start

```bash
# Clone the repository
git clone https://github.com/GhostRed256/LinkedInAutomator.git
cd LinkedInAutomator

# Install dependencies
go mod download

# Configure environment
cp .env.example .env
# Edit .env with your credentials

# Build
go build -o linkin-automator.exe cmd/main.go

# Run
./linkin-automator.exe
```

---

## ⚙️ Configuration

### Option 1: Environment Variables (`.env`)

```env
LINKEDIN_USERNAME=your_email@example.com
LINKEDIN_PASSWORD=your_secure_password
HEADLESS=false
LOG_LEVEL=info
```

### Option 2: Configuration File (`config.yaml`)

```yaml
search:
  keywords: "Software Engineer"
  limit: 20

stealth:
  min_delay: 3
  max_delay: 8
```

---

## 🚀 Usage

```go
// The tool runs automatically through these stages:

1. 🔐 Authentication (Login or restore session)
2. 🔍 Search (Find profiles based on keywords)
3. 🤝 Connect (Send personalized connection requests)
4. ✅ Check (Detect accepted connections)
5. 💌 Message (Send follow-up messages)
```

### Sample Output

```
time=2025-12-17 level=INFO msg="Starting authentication flow"
time=2025-12-17 level=INFO msg="Session is valid"
time=2025-12-17 level=INFO msg="Starting search" keywords="Recruiter"
time=2025-12-17 level=INFO msg="Found profiles" count=15
time=2025-12-17 level=INFO msg="Connection request sent"
time=2025-12-17 level=INFO msg="Workflow complete"
```

---

## 🏗️ Architecture

```
LinkinAutomator/
├── 📁 cmd/main.go                    # Application entry point
├── 📁 internal/
│   ├── auth/                         # Authentication logic
│   ├── browser/                      # Rod wrapper with stealth
│   ├── config/                       # Configuration management
│   ├── connection/                   # Connection requests
│   ├── messaging/                    # Follow-up messaging
│   ├── search/                       # Profile search & scraping
│   └── storage/                      # State persistence
├── 📁 pkg/
│   ├── logger/                       # Structured logging
│   └── stealth/                      # Anti-detection engine
│       ├── behavior.go               # Business hours, hovering
│       ├── input.go                  # Realistic typing
│       ├── mouse.go                  # Bézier movement
│       ├── scroll.go                 # Human scrolling
│       └── stealth.go                # Core engine
└── 📁 data/                          # Auto-generated storage
```

---

## 📚 Documentation

| Document | Description |
|----------|-------------|
| 📄 [FEATURES.md](FEATURES.md) | Complete feature checklist |
| 📖 [TECHNICAL_WALKTHROUGH.md](TECHNICAL_WALKTHROUGH.md) | In-depth technical guide |
| 🔧 [.env.example](.env.example) | Configuration template |

---

## 🎯 Key Highlights

<div align="center">

### **Why This Project Stands Out**

</div>

- 🧠 **Smart Architecture**: Clean, modular design with separation of concerns
- 🎭 **Advanced Stealth**: 8+ anti-detection techniques working in harmony
- ⚡ **Performance**: Efficient session reuse and smart caching
- 🛡️ **Robust**: Multiple fallback strategies for reliability
- 📊 **State Management**: Thread-safe storage with comprehensive tracking
- 🎨 **Clean Code**: Professional Go practices with clear documentation

---

## ⚠️ Important Notice

<div align="center">

### **Educational Purpose Only**

This project is a **technical demonstration** of browser automation and anti-detection techniques.

⚡ Using automation tools on LinkedIn violates their Terms of Service  
⚡ May result in account restrictions or permanent bans  
⚡ Not intended for production use  
⚡ For educational and portfolio purposes only  

</div>

---

## 🔒 Security & Privacy

- ✅ All credentials stored locally
- ✅ No external API calls
- ✅ Session cookies saved locally only
- ✅ Zero data transmission to third parties
- ✅ Runs entirely on your machine

---

## 🛠️ Tech Stack

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Chrome](https://img.shields.io/badge/Chrome-4285F4?style=for-the-badge&logo=googlechrome&logoColor=white)
![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)

**Core**: Go 1.21+ | **Automation**: Rod (CDP) | **Stealth**: Custom Engine

</div>

---

## 📈 Performance Metrics

- ⚡ **Session Restore**: < 2 seconds
- 🔍 **Search Speed**: ~10 profiles/minute
- 🤝 **Connection Rate**: 20/day (configurable)
- 💾 **Memory Usage**: < 100MB

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to check the issues page.

---

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">

## 💫 **About the Developer**

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=28&duration=3000&pause=1000&color=00ADD8&center=true&vCenter=true&width=600&lines=Built+with+%E2%9D%A4%EF%B8%8F+by+Ritesh+Dey;Software+Engineer+%7C+Go+Developer;LinkedIn+Automation+Expert" alt="Typing SVG" />

### ✨ **Made by [Ritesh Dey](https://github.com/GhostRed256)** ❤️ ✨

<p align="center">
<a href="https://github.com/GhostRed256">
<img src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white"/>
</a>
<a href="https://www.linkedin.com/in/ritesh-dey-77887a219">
<img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white"/>
</a>
</p>

---

### 💖 **Support Me**

<p align="center">
If you found this project valuable, consider supporting my work!
</p>

<p align="center">
<a href="https://github.com/GhostRed256/LinkedInAutomator">
<img src="https://img.shields.io/badge/Support-Give%20a%20Star%20⭐-yellow?style=for-the-badge"/>
</a>
<a href="https://www.linkedin.com/in/ritesh-dey-77887a219">
<img src="https://img.shields.io/badge/Connect-LinkedIn-0077B5?style=for-the-badge&logo=linkedin"/>
</a>
</p>

---

### 🌟 If you found this project helpful, please give it a star! 🌟

<img src="https://img.shields.io/github/stars/GhostRed256/LinkedInAutomator?style=social" alt="GitHub stars"/>

---

<sub>**Crafted with precision, powered by innovation** 🚀</sub>

<sub>*Showcasing advanced Go development, browser automation mastery, and anti-detection engineering*</sub>

</div>

---

<div align="center">

**© 2025 Ritesh Dey. All Rights Reserved.**

*Building the future, one line of code at a time* ⚡

</div>
