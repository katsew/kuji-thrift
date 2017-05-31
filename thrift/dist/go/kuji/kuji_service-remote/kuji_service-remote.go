// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "kuji"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Response ThRegisterCandidatesWithKey(ReqCandidates req)")
  fmt.Fprintln(os.Stderr, "  Response ThPickOneByKey(ReqPickOneByKey req)")
  fmt.Fprintln(os.Stderr, "  Response ThPickOneByKeyAndIndex(ReqPickOneByKeyAndIndex req)")
  fmt.Fprintln(os.Stderr, "  Response ThPickAndDeleteOneByKey(ReqPickOneByKey req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := kuji.NewKujiServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "ThRegisterCandidatesWithKey":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ThRegisterCandidatesWithKey requires 1 args")
      flag.Usage()
    }
    arg11 := flag.Arg(1)
    mbTrans12 := thrift.NewTMemoryBufferLen(len(arg11))
    defer mbTrans12.Close()
    _, err13 := mbTrans12.WriteString(arg11)
    if err13 != nil {
      Usage()
      return
    }
    factory14 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt15 := factory14.GetProtocol(mbTrans12)
    argvalue0 := kuji.NewReqCandidates()
    err16 := argvalue0.Read(jsProt15)
    if err16 != nil {
      Usage()
      return
    }
    value0 := kuji.ReqCandidates(argvalue0)
    fmt.Print(client.ThRegisterCandidatesWithKey(value0))
    fmt.Print("\n")
    break
  case "ThPickOneByKey":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ThPickOneByKey requires 1 args")
      flag.Usage()
    }
    arg17 := flag.Arg(1)
    mbTrans18 := thrift.NewTMemoryBufferLen(len(arg17))
    defer mbTrans18.Close()
    _, err19 := mbTrans18.WriteString(arg17)
    if err19 != nil {
      Usage()
      return
    }
    factory20 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt21 := factory20.GetProtocol(mbTrans18)
    argvalue0 := kuji.NewReqPickOneByKey()
    err22 := argvalue0.Read(jsProt21)
    if err22 != nil {
      Usage()
      return
    }
    value0 := kuji.ReqPickOneByKey(argvalue0)
    fmt.Print(client.ThPickOneByKey(value0))
    fmt.Print("\n")
    break
  case "ThPickOneByKeyAndIndex":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ThPickOneByKeyAndIndex requires 1 args")
      flag.Usage()
    }
    arg23 := flag.Arg(1)
    mbTrans24 := thrift.NewTMemoryBufferLen(len(arg23))
    defer mbTrans24.Close()
    _, err25 := mbTrans24.WriteString(arg23)
    if err25 != nil {
      Usage()
      return
    }
    factory26 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt27 := factory26.GetProtocol(mbTrans24)
    argvalue0 := kuji.NewReqPickOneByKeyAndIndex()
    err28 := argvalue0.Read(jsProt27)
    if err28 != nil {
      Usage()
      return
    }
    value0 := kuji.ReqPickOneByKeyAndIndex(argvalue0)
    fmt.Print(client.ThPickOneByKeyAndIndex(value0))
    fmt.Print("\n")
    break
  case "ThPickAndDeleteOneByKey":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ThPickAndDeleteOneByKey requires 1 args")
      flag.Usage()
    }
    arg29 := flag.Arg(1)
    mbTrans30 := thrift.NewTMemoryBufferLen(len(arg29))
    defer mbTrans30.Close()
    _, err31 := mbTrans30.WriteString(arg29)
    if err31 != nil {
      Usage()
      return
    }
    factory32 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt33 := factory32.GetProtocol(mbTrans30)
    argvalue0 := kuji.NewReqPickOneByKey()
    err34 := argvalue0.Read(jsProt33)
    if err34 != nil {
      Usage()
      return
    }
    value0 := kuji.ReqPickOneByKey(argvalue0)
    fmt.Print(client.ThPickAndDeleteOneByKey(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
