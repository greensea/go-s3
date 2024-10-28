// 用于 Cloudflare Worker 的实现，但尚未成功，暂时注释不用

// package s3

// import (
// 	"context"
// 	"crypto/tls"
// 	"log"
// 	"net"
// 	"net/http"

// 	"github.com/syumai/workers/cloudflare/sockets"

// 	"github.com/rhnvrm/simples3"
// )

// func isHTTPS(addr string) bool {
// 	// 简单判断端口是否为 443，可以根据需要改进
// 	_, port, _ := net.SplitHostPort(addr)
// 	return port == "443"
// }

// func S3() *simples3.S3 {
// 	log.Printf("开始初始化 S3")

// 	s := simples3.New("", S3_AK, S3_SK)

// 	s.SetClient(&http.Client{
// 		Transport: &http.Transport{
// 			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
// 				log.Printf("设置了新的 DialContext: %v %v", network, addr)

// 				conn, err := sockets.Connect(ctx, addr, &sockets.SocketOptions{
// 					SecureTransport: sockets.SecureTransportStartTLS,
// 				})

// 				if err != nil {
// 					return nil, err
// 				}

// 				// 如果是 HTTPS，需要进行 TLS 握手
// 				if false && network == "tcp" && isHTTPS(addr) {
// 					log.Printf("设置 TLS 握手")

// 					// hostArr := strings.Split(addr, ":")
// 					tlsConn := tls.Client(conn, &tls.Config{
// 						InsecureSkipVerify: true, // 根据需要设置是否跳过证书验证
// 						// ServerName:         hostArr[0],
// 					})

// 					// if err := tlsConn.Handshake(); err != nil {
// 					// 	log.Printf("握手出错: %v", err)
// 					// 	conn.Close()
// 					// 	return nil, err
// 					// }
// 					return tlsConn, nil
// 				} else {
// 					return conn, nil
// 				}
// 			},
// 		},
// 	})

// 	s.SetEndpoint(S3_ENDPOINT)

// 	return s
// }
