package apns

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"net"
	"strings"
	"time"
)

type Client struct {
	Gateway  string
	KeyFile  string
	CertFile string
}

// MARK: Struct's constructors
func CreateApnsClient(gateway string, certFile string, keyFile string) *Client {
	client := Client{
		Gateway:  gateway,
		KeyFile:  keyFile,
		CertFile: certFile,
	}
	return &client
}

// MARK: Struct's public functions
func (c *Client) Send(apns []*APNs) (resp *APNsResponse) {
	envelop := bytes.NewBuffer([]byte{})
	for _, apn := range apns {
		message, err := apn.Encode()
		if err == nil {
			binary.Write(envelop, binary.BigEndian, uint8(2))
			binary.Write(envelop, binary.BigEndian, uint32(len(message)))
			binary.Write(envelop, binary.BigEndian, message)
		}
	}

	resp, err := c.connectAndWrite(envelop.Bytes())
	if resp == nil {
		resp = &APNsResponse{}
	}

	if err != nil {
		resp.Success = false
		resp.Error = err
	}

	//	resp.Success = true
	//	resp.Error = nil

	return resp
}

// MARK: Struct's private functions
func (c *Client) connectAndWrite(envelop []byte) (*APNsResponse, error) {
	// Load keypair
	kp, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, err
	}

	// Config TLS
	tokens := strings.Split(c.Gateway, ":")
	conf := &tls.Config{
		Certificates: []tls.Certificate{kp},
		ServerName:   tokens[0],
	}

	// Connect to Apple gateway
	conn, err := net.Dial("tcp", c.Gateway)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Handshake
	tlsConn := tls.Client(conn, conf)
	err = tlsConn.Handshake()
	if err != nil {
		return nil, err
	}
	defer tlsConn.Close()

	// Transfer data
	_, err = tlsConn.Write(envelop)
	if err != nil {
		return nil, err
	}

	// This channel that will serve to handle timeouts when the notification succeeds.
	timeoutChannel := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeoutChannel <- true
	}()

	// This channel will contain the binary response from Apple in the event of a failure.
	responseChannel := make(chan []byte, 1)
	go func() {
		buffer := make([]byte, 6, 6)
		tlsConn.Read(buffer)
		responseChannel <- buffer
	}()

	/**
	 * First one back wins! The data structure for an APN response is as follows:
	 * command    -> 1 byte
	 * status     -> 1 byte
	 * identifier -> 4 bytes
	 * The first byte will always be set to 8.
	 */
	var response APNsResponse
	select {
	case r := <-responseChannel:
		response = APNsResponse{
			Success:       true,
			AppleResponse: APNsResponses[r[1]],
			//			Error:         errors.New(resp.AppleResponse),
		}

	case <-timeoutChannel:
		response = APNsResponse{
			Success: false,
		}
	}
	return &response, err
}
