package main

import (
        "os"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/http"
        "time"
        "path/filepath"
        "strconv"
        "io/ioutil"

        // These are in local subdirs
	"dclass.ucscx.edu/contract"
	"dclass.ucscx.edu/lib"
)

const port = 1234


// The file Type (more like an object)
type File struct {
    Filename   string  `json:"filename,omitempty"`
    Type       string  `json:"type,omitempty"`
    Timestamp  string  `json:"timestamp,omitempty"`
    Length     string  `json:"length,omitempty"`
    Data       string  `json:"data,omitempty"`
}


// All handlers are referenced via this object/service
type BackEndHandler struct{}

func (h *BackEndHandler ) GetFiles(request *contract.GetFilesRequest, reply *contract.GetFilesResponse) error {

    log.Println("GetFiles")
    log.Println(request)

    //dirname := "." + string(filepath.Separator)
    //dirname := "./files/" + string(filepath.Separator) // local subdir
    dirname := "/files/" + string(filepath.Separator)    // global volume mounted dir

    d, err := os.Open(dirname)
    if err !=nil {
        reply.Status = http.StatusInternalServerError
        //log.Println("Failed to open ./files directory.") // local subdir
        log.Println("Failed to open /files directory.")    // global volume mounted dir
        return nil
    }
    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
        reply.Status = http.StatusInternalServerError
        //log.Println("Failed to read list of files in ./ directory.")
        log.Println("Failed to read list of files in /files directory.")
        return nil
    }

    for _, file := range files {
        if file.Mode().IsRegular() {
            reply.Filenames = append(reply.Filenames, file.Name())
        }
    }

    reply.Status = http.StatusOK

    log.Println(reply)

    return nil
}


func (h *BackEndHandler) GetFile(request *contract.GetFileRequest, reply *contract.GetFileResponse) error {

    log.Println("GetFile")
    log.Println(request)

    //filename := "./files/" + request.Filename // local subdir
    filename := "/files/" + request.Filename    // global dir in mounted volume

//    file, err := os.Open(request.Filename)
    file, err := os.Open(filename)
    if err !=nil {
        reply.Status = http.StatusNotFound
        log.Println("Failed to open or unknown file %s.", request.Filename)
        return nil
    }
    defer file.Close()

    fileInfo, err := file.Stat()
    if err !=nil {
        reply.Status = http.StatusInternalServerError
        log.Println("Failed to stat file %s.", request.Filename)
        return nil
    }


    if fileInfo.Mode().IsRegular() {
        var buf []byte
        reply.Filename = file.Name()
        reply.Length = strconv.Itoa(int(fileInfo.Size()))
        reply.Timestamp = fileInfo.ModTime().String()
        //cnt, err := file.Read(buf)
        buf, err := ioutil.ReadFile(file.Name())
        if err != nil {
            reply.Status = http.StatusInternalServerError
            log.Println("Failed to read file %s data.", file.Name())
            return nil
        }
        reply.Data = string(buf)
    }

    reply.Status = http.StatusOK

    log.Println(reply)

    return nil
}


func (h *BackEndHandler) CreateFile(request *contract.CreateFileRequest, reply *contract.CreateFileResponse) error {

    log.Println("CreateFile")
    log.Println(request)

    //filename := "./files/" + request.Filename // local subdir
    filename := "/files/" + request.Filename    // global dir in mounted volume

    //file, err := os.Create(request.Filename)
    file, err := os.Create(filename)
    if err !=nil {
        reply.Status = http.StatusInternalServerError
        log.Println("Failed to create file %s.", request.Filename)
        return nil
    }
    defer file.Close()

    _, err = file.WriteString(request.Data)
    if err != nil {
        reply.Status = http.StatusInternalServerError
        log.Println("Failed to write file %s.", request.Filename)
        return nil
    }
    defer file.Close()

    reply.Status = http.StatusCreated

    log.Println(reply)

    return nil
}


func (h *BackEndHandler) DeleteFile(request *contract.DeleteFileRequest, reply *contract.DeleteFileResponse) error {

    log.Println("DeleteFile")
    log.Println(request)

    //filename := "./files/" + request.Filename // local subdir
    filename := "/files/" + request.Filename    // global dir in mounted volume

    //err := os.Remove(request.Filename)
    err := os.Remove(filename)
    if err != nil {
        reply.Status = http.StatusNotFound
        return nil
    }

    reply.Status = http.StatusOK

    log.Println(reply)

    return nil
}


// Get preferred outbound IP of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}


// Get preferred outbound IP of this machine
//func GetOutboundIP() net.IP {
//    ief, err := net.InterfaceByName("docker0")
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    addrs, err := ief.Addrs()
//    if err != nil {
//        log.Fatal(err)
//    }
//
////    tcpAddr := &net.TCPAddr {
////        IP: addrs[0].(*net.IPNet).IP,
////    }
////
////    return tcpAddr
//
//    return addrs[0].(*net.IPNet).IP
//}

func beacon(strRemoteIp string) {

    var localIp, remoteIp net.IP
    var localIpAddr, remoteIpAddr net.IPAddr
    var localUdpAddr, remoteUdpAddr net.UDPAddr

    var strLocalIpAddr string

    remoteIp = net.ParseIP(strRemoteIp)

    localIp = GetOutboundIP()

    localIpAddr = net.IPAddr{IP: localIp}
    remoteIpAddr = net.IPAddr{IP: remoteIp}

    localUdpAddr = net.UDPAddr{IP: localIp, Port: 0}
    remoteUdpAddr = net.UDPAddr{IP: remoteIp, Port: 5555}

    log.Printf("Dialing localIP:remoteIP %s:%s\n", localIpAddr.String(), remoteIpAddr.String())
    fmt.Printf("Dialing localIP:remoteIP %s:%s\n", localIpAddr.String(), remoteIpAddr.String())

    conn, err := net.DialUDP("udp", &localUdpAddr, &remoteUdpAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    strLocalIpAddr = localIpAddr.String()
    //strLocalIpAddr = ""

    buf := []byte(strLocalIpAddr) 

    //log.Printf("Starting IP beaconing on %s with payload %s\n", strRemoteIp, strLocalIpAddr)
    //fmt.Printf("Starting IP beaconing on %s with payload %s\n", strRemoteIp, strLocalIpAddr)
    log.Printf("Starting UDP beaconing on %s with payload %s\n", strRemoteIp, strLocalIpAddr)
    fmt.Printf("Starting UDP beaconing on %s with payload %s\n", strRemoteIp, strLocalIpAddr)
    for { // infinite loop
        cnt := 0
        time.Sleep(5000 * time.Millisecond)
        cnt, err := conn.Write(buf)
        if err != nil {
            continue
            //log.Fatal(err)
        }
        fmt.Printf("Write with payload |%s| (%d bytes)\n", strLocalIpAddr, cnt)
    }

}


func main() {
    StartServer()
}

func StartServer() {

    var strBackEndPort string
    var strBackEndIP   string
    var backEndIP net.IP

    var strRemoteIP   string

    f, err := os.OpenFile("backend.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)

    config, err := lib.ReadConfig(`./config.txt`)
    if err != nil {
        fmt.Println(err)
    }

    //fmt.Println("Config data dump :", config)

    // assign values from config file to variables
    strBackEndPort = config["BACKENDPORT"]
    //strBackEndIP = config["BACKENDIP"]

    strRemoteIP = config["FRONTENDIP"] // replaced by DNS resolve of "frontend"

    backEndIP = GetOutboundIP()
    strBackEndIP = backEndIP.String()

    log.Printf("Backend server IP address is %s\n", strBackEndIP)
    fmt.Printf("Backend server IP address is %s\n", strBackEndIP)
    log.Printf("Backend server starting on port %s\n", strBackEndPort)
    fmt.Printf("Backend server starting on port %s\n", strBackEndPort)

    backEndService := &BackEndHandler{}
    rpc.Register(backEndService)

    strIpList, err := net.LookupHost("frontend")
    if err != nil {
        fmt.Println(err)
    }

    strRemoteIP = strIpList[0]

    log.Printf("Starting IP beaconing on %s\n", strRemoteIP)
    fmt.Printf("Starting IP beaconing on %s\n", strRemoteIP)

    go beacon(strRemoteIP)
    //beacon(strRemoteIP)

    log.Printf("Going to listen on port %s\n", strBackEndPort)
    fmt.Printf("Going to listen on port %s\n", strBackEndPort)

    l, err := net.Listen("tcp4", fmt.Sprintf("%s:%s", strBackEndIP, strBackEndPort))
    if err != nil {
        log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
    }
    defer l.Close()

    log.Printf("Server up and listening on IP:port %s:%s\n", strBackEndIP, strBackEndPort)
    fmt.Printf("Server up and listening on IP:port %s:%s\n", strBackEndIP, strBackEndPort)

    for {
        conn, err := l.Accept()
        if err != nil {
             log.Println(err)
             continue 
        }
        go rpc.ServeConn(conn)
        //conn, _ := l.Accept()
        //go rpc.ServeConn(conn)
    }
}

