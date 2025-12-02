package connected

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConnection struct {
	IP     net.IP
	Port   int
	Client *grpc.ClientConn
}

var clientConnections []ClientConnection

func CreateClientConnection(ip net.IP, port int) (ClientConnection, error) {
	fullAddress := net.JoinHostPort(ip.String(), fmt.Sprintf("%d", port))
	conn, err := grpc.NewClient(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return ClientConnection{}, err
	}
	return ClientConnection{
		ip,
		port,
		conn,
	}, nil
}
func GetRawClientConnection(ip net.IP, port int) (ClientConnection, error) {
	for _, conn := range clientConnections {
		if conn.IP.String() == ip.String() && conn.Port == port {
			return conn, nil
		}
	}
	conn, err := CreateClientConnection(ip, port)
	if err != nil {
		return ClientConnection{}, err
	}
	clientConnections = append(clientConnections, conn)
	return conn, nil
}
func GetDeviceClientConnection(dev *Connected) (*grpc.ClientConn, error) {
	return GetClientConnection(dev.IP, dev.Port)
}
func GetClientConnection(ip net.IP, port int) (*grpc.ClientConn, error) {
	conn, err := GetRawClientConnection(ip, port)
	if err != nil {
		return nil, err
	} else {
		return conn.Client, nil
	}
}
