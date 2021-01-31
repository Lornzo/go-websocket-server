# Golang 寫成的 websocket server
用官方Api寫成的 WebSocket Server !  

## 相依套件
* golang.org/x/net/websocket

## WebSocket 用途
透過websocket，可以讓網頁用JS的方式，與Server建立一個類Socket的雙向長連線，進而實現即時資料交換與互動的功能

## WebSocket 連線流程
* Client端發起握手請求(Hand Shake)
  * 發起一個http request，並在request中發送：
    * Upgrade
    * Connection Header
    * Sec-WebSocket-Key (base64的key，這個會與Server Response互相對應)
    * Sec-WebSocket-Version (websocket的版本)
* Server端確認並允許握手請求
  * Server收到Hand Shake，並打出Resposne：
    * Sec-WebSocket-Accept (前面從Client拿到的Key，經過處理加密後打回去，讓Client確認這個Response是從這個Server來的)
        * Accept = base64(sha1(Sec-WebSocket-Key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
    * status 101 (101代表跟Client說：我收到了你的請求，請Client端切換到WebSocket Protocal)
* Client 與 Server建立連線
* 連線過程中，WebSocket Client 定時透過 ping(0x9) 和 pong(0xA)來確認Server端是否還維持著連線
    * 要是Server收到 ping 之後沒有回應 pong 給Client，Client會判斷Server斷線，
    * 之後Client端可能會重連或是關閉連線
* 連線結束

## golang.org/x/net/websocket
* 套件已經把HandShake這件事情做完了，所以直接使用連線本體就好

## 注意事項
* 前端連進來的時候，請使用WebSocket協定

