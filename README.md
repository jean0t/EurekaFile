# 💾 EurekaFile – File Sharing, the 90s Way (But Actually Fast)

**Welcome to EurekaFile**, the raddest file server on the block, coded in **Go** (not the board game, the programming language — we’re not that retro).

This bad boy lets you **upload**, **download**, and **rock out** with your files like it’s 1999. It’s got login pages, session magic, and even a UI so fresh, you’ll think it came off a GeoCities page (but like, classy).

---

## 🔥 Features – Cooler Than a Tamagotchi

* **🕶️ User Login and Sessions**
  Type in a username and password, and BAM — you're in. Never registered? No prob, we make a new account just for you, *automagically*.

* **📼 Upload & Download**
  Drop your files in, pull them out later — it’s like a virtual backpack, but with fewer Pokémon cards.

* **💻 Responsive UI**
  Built with Bootstrap so it adjusts smoother than a Discman on anti-skip. Desktop? Phone? Pager? OK maybe not pager.

---

## 📟 Totally Tubular Routes

| Route     | Function                                                      |
| --------- | ------------------------------------------------------------- |
| `/login`  | Log in, or create a new account if you're a n00b.             |
| `/upload` | File uploader interface. Only for logged-in homies.           |
| `/files`  | Browse your sweet uploads. Details included. Downloads ready. |
| `/logout` | Peace out and log off like it’s the end of an AOL chat room.  |

---

## 🧠 How It’s All Wired Up (Internally, Not Emotionally)

```
internal
├── controllers      # Like your favorite game controller, but for web routes
├── database         # Stores users and files, not Doom save files sadly
├── middleware       # Adds extras like logging, just like Winamp plugins
├── router           # All the URL pathways you need, no MapQuest required
└── views            # HTML templates with Bootstrap – no Comic Sans here
```

---

## 💿 Getting Started – Install It Like It's Shareware

### 1. Clone the code like a true hacker:

```sh
git clone github.com/jean0t/EurekaFile  # or try JMFern01 if you're feeling retro rebellious
cd EurekaFile
```

### 2. Build it like you're installing Quake:

```sh
go build -o eurekafile cmd/main.go
./eurekafile -s
```

### 3. Fire up Netscape Navigator (or just Chrome):

Go to [http://localhost:8080/login](http://localhost:8080/login)
Enter your info, and ride the wave.

---

## 📜 License – MIT, Baby

Like most things in the '90s, you can use it however you want. Just don’t be lame about it.

---

> That’s it! Plug in, power up, and relive the glory days of file sharing — except this time, it’s actually secure and doesn’t crash your parents’ PC.
