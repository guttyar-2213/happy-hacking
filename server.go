package main
import (
  "fmt"
  "net"
  "strconv"
  "github.com/google/uuid"
  "encoding/hex"
)

type Conn net.Conn

func serve(port int, handler func(Conn, error) error) error {
  listen, err := net.Listen("tcp", ":" + strconv.Itoa(port));
  if (err != nil) { return err }
  fmt.Println(listen.Addr())
  for { go handler(listen.Accept()) }
  return nil
}
func handleConnection(buffersize int) func(conn Conn, err error) error {
  return func(conn Conn, err error) error {
    id := uuid.New()
    defer conn.Close()
    defer fmt.Println("Connection closed. id:", id)
    if (err != nil) { return err }
    fmt.Println("Connection open. id:", id)
    for {
      buf := make([]byte, buffersize)
      size, err := conn.Read(buf)
      if err != nil { return err }
      fmt.Println("Data received. id:", id, "size:", size, "data:", hex.EncodeToString(buf[:size]))
    }
    return nil
  }
}
func main() { fmt.Println(serve(8080, handleConnection(128))) }

// 大事なところ
// 13行目 net.Listen(...)
// サーバ起動、第一引数はモード(UDPはない)、第二引数はアドレス
// 16行目 for { go handler(listen.Accept()) }
// 無限ループ(for { ... }) + 並行処理(go handler(...))で複数コネクションに対応
// listen.Accept()でコネクションを確立
// 22行目 defer conn.Close()
// deferで関数がreturnした後コネクションを閉じる
// 28行目 conn.Read(buf)
// bufferに入る分だけ読み込む。送信されたデータがない場合待機
