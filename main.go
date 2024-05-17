package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const maxWorkers = 100 // chill out goroutines

func main() {
	var IP string
	fmt.Printf("Enter IP address: ")
	fmt.Scan(&IP)

	var wg sync.WaitGroup
	fastAsFuckBoy := make(chan struct{}, maxWorkers) // Semaphore to limit concurrency

	for port := uint16(0); port < 65535; port++ {
		wg.Add(1)
		fastAsFuckBoy <- struct{}{} // Acquire a slot from the fastAsFuckBoy
		go func(p uint16) {
			defer func() {
				wg.Done()
				<-fastAsFuckBoy // Release the slot
			}()
			portScanner(IP, p)
		}(port)
	}

	wg.Wait()
	close(fastAsFuckBoy)
}

func portScanner(ip string, port uint16) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second) // Reduced timeout
	if err != nil {
		return // Port is closed or fucked
	}
	defer conn.Close()
	fmt.Printf("[+] Port %d is open\n", port)
}

// amigo no gooda
//no bomboclat gooda
//package main

// import (
// 	"fmt"
// 	"net"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	ch := make(chan int)
// 	var IP string
// 	fmt.Printf("enter ip address here:")
// 	fmt.Scan(&IP)

// 	//        For most vunurable ports
// 	//
// 	// for port, vuln := range vulnPorts {
// 	// 	go portScanner(IP, vuln, port, ch)
// 	// 	<-ch
// 	// }
// 	//        Scan all existing ports
// 	//
// 	for port := 0; port < 65535; port++ {
// 		go portScanner(IP, "No discription", port, ch)
// 		<-ch

// 	}
// }

// func portScanner(ip string, portDiscription string, port int, ch chan<- int) {
// 	address := ip + ":" + strconv.Itoa(port)
// 	_, err := net.DialTimeout("tcp", address, 10*time.Second)
// 	if err == nil {
// 		fmt.Printf("\n[+] port: %v\n-->  for : %v \n", port, portDiscription)
// 	}
// 	ch <- port
// 	return
// }

// var vulnPorts = map[int]string{
// 	1:     "tcpmux",
// 	7:     "echo",
// 	9:     "discard",
// 	11:    "systat",
// 	13:    "daytime",
// 	15:    "netstat",
// 	17:    "qotd",
// 	19:    "chargen",
// 	20:    "ftp data",
// 	21:    "ftp access",
// 	22:    "ssh",
// 	23:    "telnet",
// 	25:    "smtp",
// 	37:    "time",
// 	42:    "name",
// 	43:    "nicname",
// 	53:    "domain",
// 	69:    "tftp",
// 	77:    "priv-rjs",
// 	79:    "finger",
// 	80:    "http",
// 	87:    "ttylink",
// 	95:    "supdup",
// 	101:   "hostriame",
// 	102:   "iso-tsap",
// 	103:   "gppitnp",
// 	104:   "acr-nema",
// 	109:   "pop2",
// 	110:   "pop3",
// 	111:   "sunrpc",
// 	113:   "auth",
// 	115:   "sftp",
// 	117:   "uucp-path",
// 	119:   "nntp",
// 	123:   "NTP",
// 	135:   "loc-srv /epmap",
// 	137:   "netbios",
// 	139:   "netbios",
// 	143:   "imap2",
// 	161:   "snmp",
// 	179:   "BGP",
// 	389:   "ldap",
// 	427:   "SLP (Also used by Apple Filing Protocol)",
// 	465:   "smtp+ssl",
// 	512:   "print / exec",
// 	513:   "login",
// 	514:   "shell",
// 	515:   "printer",
// 	526:   "tempo",
// 	530:   "courier",
// 	531:   "chat",
// 	532:   "netnews",
// 	540:   "uucp",
// 	548:   "AFP (Apple Filing Protocol)",
// 	554:   "rtsp",
// 	556:   "remotefs",
// 	563:   "nntp+ssl",
// 	587:   "smtp (rfc6409)",
// 	601:   "syslog-conn (rfc3195)",
// 	636:   "ldap+ssl",
// 	989:   "ftps-data",
// 	990:   "ftps",
// 	993:   "ldap+ssl",
// 	995:   "pop3+ssl",
// 	1719:  "h323gatestat",
// 	1720:  "h323hostcall",
// 	1723:  "pptp",
// 	2049:  "nfs",
// 	3000:  "Dev host",
// 	3659:  "apple-sasl / PasswordServer",
// 	4045:  "lockd",
// 	5060:  "sip",
// 	5061:  "sips",
// 	6000:  "X11",
// 	6566:  "sane-port",
// 	6665:  "Alternate IRC [Apple addition]",
// 	6666:  "Alternate IRC [Apple addition]",
// 	6667:  "Standard IRC [Apple addition]",
// 	6668:  "Alternate IRC [Apple addition]",
// 	6669:  "Alternate IRC [Apple addition]",
// 	6697:  "IRC +, TLS",
// 	10080: " Amanda,",
// }
