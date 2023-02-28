package main

import (
    "os"
    "encoding/json"
    "log"
    "net"
    "net/http"
    "net/rpc"
    "fmt"
    "time"
    "strings"

    // These are in local subdirs
    "dclass.ucscx.edu/gorilla/mux"
    "dclass.ucscx.edu/contract"
    "dclass.ucscx.edu/lib"
)

const port = 1234  // RPC
const port2 = 8000 // HTTP REST

var strBackEndPort string
var strBackEndIP string

var  file2IpMap map[string]string   // map: key = filename, value = backendIP
var  Ip2FilesMap map[string][]string // map: key = backendIP, value = list of files

// The file Type (more like an object)
type File struct {
    Filename   string  `json:"filename,omitempty"`
    Type       string  `json:"type,omitempty"` 
    Timestamp  string  `json:"timestamp,omitempty"`
    Length     string  `json:"length,omitempty"`
    Data       string  `json:"data,omitempty"`
}


// Display all from the files var
func GetFiles(w http.ResponseWriter, r *http.Request) {

    fmt.Println("GetFiles")
    log.Println("GetFiles")
    log.Println(r)
   
    for keyIP, _ := range Ip2FilesMap {

        tmpStrBackEndIP := keyIP

        c := CreateClient(tmpStrBackEndIP, strBackEndPort)
        defer c.Close()

        reply := PerformGetFilesRequest(c)

        //fmt.Println(reply)
        log.Println(reply)

        //fmt.Println(reply.Filenames)

        if reply.Status != http.StatusOK {
            w.WriteHeader(reply.Status)
            return
        }

        json.NewEncoder(w).Encode(tmpStrBackEndIP)
        json.NewEncoder(w).Encode(reply.Filenames)

        Ip2FilesMap[tmpStrBackEndIP] = reply.Filenames

    }

}

func isStringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func findStringInSlice(a string, list []string) int {
    for c, b := range list {
        if b == a {
            return c
        }
    }
    return -1
}


func removeStringInSlice(i int, s []string) []string {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

// Display a single file
func GetFile(w http.ResponseWriter, r *http.Request) {

    var request contract.GetFileRequest
    var tmpFile File
    var tmpFilename string

    fmt.Println("GetFile")
    log.Println("GetFile")
    log.Println(r)

    params := mux.Vars(r)

    tmpFilename = params["filename"] 
    request.Filename = tmpFilename 

    log.Println(request)

    // Determine which backend server to contact
    for keyIP, valueFilenames := range Ip2FilesMap {

        tmpStrBackEndIP := keyIP

        if isStringInSlice(tmpFilename, valueFilenames) {

            c := CreateClient(tmpStrBackEndIP, strBackEndPort)
            defer c.Close()

            reply := PerformGetFileRequest(c, request)

            log.Println(reply)

            if reply.Status != http.StatusOK {
                w.WriteHeader(reply.Status)
                return 
            }
 
            tmpFile.Filename = reply.Filename
            tmpFile.Type = reply.Type
            tmpFile.Timestamp = reply.Timestamp
            tmpFile.Length = reply.Length
            tmpFile.Data = reply.Data

            json.NewEncoder(w).Encode(tmpFile)
            return 

        }
    }


    // if can't find file then lets just get an error response

    c := CreateClient(strBackEndIP, strBackEndPort)
    defer c.Close()

    reply := PerformGetFileRequest(c, request)

    log.Println(reply)

    w.WriteHeader(reply.Status)
    return 


}

// create a new file
func CreateFile(w http.ResponseWriter, r *http.Request) {

    var request contract.CreateFileRequest

    fmt.Println("CreateFile")
    log.Println("CreateFile")
    log.Println(r)

    params := mux.Vars(r)

    var file File

    _ = json.NewDecoder(r.Body).Decode(&file)

//    if file.Filename != params["filename"] {
//        w.WriteHeader(http.StatusBadRequest)
//        return nil
//    }

    request.Filename = params["filename"] 

    tmpFilename := request.Filename

    //request.Filename = file.Filename
    request.Type = file.Type
    request.Timestamp = file.Timestamp
    request.Length = file.Length
    request.Data = file.Data

    log.Println(request)

   
    var minCntFilenames int
    var minStrBackEndIP string

    cnt := 0
    var adjValueFilenames []string


    // Determine which backend server to contact
    for keyIP, valueFilenames := range Ip2FilesMap {

        // check if filename is already stored
        if isStringInSlice(tmpFilename, valueFilenames) {
            minStrBackEndIP = keyIP
            c := CreateClient(minStrBackEndIP, strBackEndPort)
            reply := PerformCreateFileRequest(c, request)
            log.Println(reply)
            w.WriteHeader(reply.Status)
            return  // error exit
        }

        cntFilenames := len(valueFilenames)

        if cnt == 0 {
            minStrBackEndIP = keyIP
            minCntFilenames = cntFilenames
        } else if minCntFilenames > cntFilenames {
            minStrBackEndIP = keyIP
            minCntFilenames = cntFilenames
            adjValueFilenames = append(adjValueFilenames, tmpFilename)
        }

        cnt++

    }

    var c *rpc.Client

    // store new list with added filename or call default backend for error msg
    if cnt>0 {
        Ip2FilesMap[minStrBackEndIP] = adjValueFilenames
        c = CreateClient(minStrBackEndIP, strBackEndPort)
    } else {
        c = CreateClient(strBackEndIP, strBackEndPort)
    }

    reply := PerformCreateFileRequest(c, request)

    log.Println(reply)

    w.WriteHeader(reply.Status)

}

// Delete a file
func DeleteFile(w http.ResponseWriter, r *http.Request) {

    var request contract.DeleteFileRequest

    fmt.Println("DeleteFile")
    log.Println("DeleteFile")
    log.Println(r)

    params := mux.Vars(r)

    request.Filename = params["filename"] 

    tmpFilename := request.Filename

    var foundStrBackEndIP string

    cnt := 0
    var adjValueFilenames []string

    // Determine which backend server to contact
    for keyIP, valueFilenames := range Ip2FilesMap {
       
        if cnt == 0 {
            foundStrBackEndIP = keyIP  // needed for error response if not found
        }

        if isStringInSlice(tmpFilename, valueFilenames) {
            foundStrBackEndIP = keyIP

            // remove the filename
            pos := findStringInSlice(tmpFilename, valueFilenames)

            adjValueFilenames = removeStringInSlice(pos, valueFilenames)

        } 

        cnt++

    }

    var c *rpc.Client

    // store new list without deleted filename or call default backend for error msg
    if cnt>0 {
        Ip2FilesMap[foundStrBackEndIP] = adjValueFilenames
        c = CreateClient(foundStrBackEndIP, strBackEndPort)
    } else {
        c = CreateClient(strBackEndIP, strBackEndPort)
    }
    defer c.Close()

    reply := PerformDeleteFileRequest(c, request)

    log.Println(reply)

    w.WriteHeader(reply.Status)

}

func PerformGetFilesRequest(client *rpc.Client) contract.GetFilesResponse {

        request := &contract.GetFilesRequest{Unused: "Unused"}
        var reply contract.GetFilesResponse

        err := client.Call("BackEndHandler.GetFiles", request, &reply)
        if err != nil {
                log.Fatal("error:", err)
        }

        return reply
}

func PerformGetFileRequest(client *rpc.Client, request contract.GetFileRequest) contract.GetFileResponse {

        var reply contract.GetFileResponse

        err := client.Call("BackEndHandler.GetFile", request, &reply)
        if err != nil {
                log.Fatal("error:", err)
        }

        return reply
}

func PerformCreateFileRequest(client *rpc.Client, request contract.CreateFileRequest) contract.CreateFileResponse {

        var reply contract.CreateFileResponse

        err := client.Call("BackEndHandler.CreateFile", request, &reply)
        if err != nil {
                log.Fatal("error:", err)
        }

        return reply
}

func PerformDeleteFileRequest(client *rpc.Client, request contract.DeleteFileRequest) contract.DeleteFileResponse {

        var reply contract.DeleteFileResponse

        err := client.Call("BackEndHandler.DeleteFile", request, &reply)
        if err != nil {
                log.Fatal("error:", err)
        }

        return reply
}


func CreateClient(inStrBackEndIP string, inStrBackEndPort string) *rpc.Client {

        client, err := rpc.Dial("tcp4", fmt.Sprintf("%s:%s", inStrBackEndIP, inStrBackEndPort))
        if err != nil {
                log.Fatal("dialing:", err)
        }

        return client
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

func beaconReceiver() {

    var localIp net.IP
    var localIpAddr net.IPAddr
    var localUdpAddr net.UDPAddr
    var strLocalIpAddr string
    var strRemoteIpAddr string
    var strRemoteIpAddrPort string

    localIp = GetOutboundIP()

    localIpAddr = net.IPAddr{IP: localIp}
    localUdpAddr = net.UDPAddr{IP: localIp, Port: 5555}

    strLocalIpAddr = localIpAddr.String()

    log.Printf("Listening localIP %s\n", strLocalIpAddr)
    fmt.Printf("Listening localIP %s\n", strLocalIpAddr)

    conn, err := net.ListenUDP("udp", &localUdpAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    var buf []byte

    strRemoteIpAddr = ""
    strRemoteIpAddrPort = ""

    log.Printf("Starting IP beacon receiving on %s\n", strLocalIpAddr)
    fmt.Printf("Starting IP beacon receiving on %s\n", strLocalIpAddr)

    for { // infinite loop
        time.Sleep(1000 * time.Millisecond)
        //cnt, addr, err := conn.ReadFrom(buf)
        _, addr, err := conn.ReadFrom(buf)
        if err != nil {
            log.Fatal(err)
        }

        //tmpStrBuffer := string(buf[:])
        //fmt.Printf("Received backend buffer |%s| (%d bytes, from %s %s)\n", tmpStrBuffer, cnt, addr.Network(), addr.String())

        strRemoteIpAddrPort = addr.String()

        tmpStrArray := strings.Split(strRemoteIpAddrPort, ":")
        strRemoteIpAddr = tmpStrArray[0]


        _, ok := Ip2FilesMap[strRemoteIpAddr]

        if !ok {  // ok = true if key exists, false if key doesn't exist

            log.Printf("Found a new backend server at %s\n", strRemoteIpAddr)
            fmt.Printf("Found a new backend server at %s\n", strRemoteIpAddr)

            // first time we saw this IP address so set up it's key with an empty file list
            Ip2FilesMap[strRemoteIpAddr] = nil 
        }
       
    }

}

// main function to boot up everything
func main() {

    var strFrontEndIP   string
    var strFrontEndPort string

    var frontEndIP net.IP
    //var flag bool

    // These are global
    file2IpMap = make(map[string]string)   // map: key = filename, value = backendIP
    Ip2FilesMap = make(map[string][]string) // map: key = backendIP, value = list of files

    f, err := os.OpenFile("frontend.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
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
    strFrontEndPort = config["FRONTENDPORT"]
    strBackEndPort = config["BACKENDPORT"]
//    strFrontEndIP = config["FRONTENDIP"]
    strBackEndIP = config["BACKENDIP"]

    frontEndIP = GetOutboundIP()
    strFrontEndIP = frontEndIP.String()

    log.Printf("Starting IP beacon receiveing\n")
    fmt.Printf("Starting IP beacon receiveing\n")

    go beaconReceiver()

    log.Printf("Frontend server starting on IP:port %s:%s with backend server\n", strBackEndIP, strBackEndPort)
    fmt.Printf("Frontend server starting on IP:port %s:%s with backend server\n", strBackEndIP, strBackEndPort)

    router := mux.NewRouter()

    router.HandleFunc("/files", GetFiles).Methods("GET")
    router.HandleFunc("/files/{filename}", GetFile).Methods("GET")
    router.HandleFunc("/files/{filename}", CreateFile).Methods("POST")
    router.HandleFunc("/files/{filename}", DeleteFile).Methods("DELETE")

    log.Printf("Frontend server listening for HTTP REST on IP:port %s:%s\n", strFrontEndIP, strFrontEndPort)
    fmt.Printf("Frontend server listening for HTTP REST on IP:port %s:%s\n", strFrontEndIP, strFrontEndPort)

    var strPort = ":" + strFrontEndPort
//    fmt.Printf("strPort = %s", strPort)

    log.Fatal(http.ListenAndServe(strPort, router))

}

