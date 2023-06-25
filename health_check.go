package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
	"crypto/tls"
	"log"

	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////
// HEALTH CHECK                                                                  //
///////////////////////////////////////////////////////////////////////////////////

func handleHealthCheckAction(c echo.Context) error {
	urls := []string{DVWA_URL, SHOP_URL, FWB_URL, SPEEDTEST_URL, KALI_URL, "https://www.fortinet.com"}

	// Define a custom HTTP client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Start HTML Table with CSS
	result := `<style>
	table {
		width: 100%;
		border-collapse: collapse;
		font-size: 0.8em;
	}
	td, th {
		border: 1px solid #ddd;
		padding: 8px;
	}
	.failed {
		color: red;
		font-weight: bold;
	}
	</style>
	<table>
		<tr>
			<th>URL</th>
			<th>Result</th>
			<th>Code</th>
			<th>Error</th>
		</tr>`

	// Loop over the URLs
	for _, url := range urls {
		res, err := client.Get(url)
		if err != nil {
			shortErr := strings.TrimPrefix(err.Error(), fmt.Sprintf(`Get "%s": `, url))
			result += fmt.Sprintf(`<tr>
				<td>%s</td>
				<td class="failed">Failed</td>
				<td>N/A</td>
				<td>%s</td>
			</tr>`, url, shortErr)
		} else {
			result += fmt.Sprintf(`<tr>
				<td>%s</td>
				<td>Connected</td>
				<td>%d</td>
				<td>N/A</td>
			</tr>`, url, res.StatusCode)
		}
	}

	// Handle FWB_MGT_IP separately because it's only an IP without a scheme
	ip := "http://" + FWB_MGT_IP
	res, err := client.Get(ip)
	if err != nil {
		shortErr := strings.TrimPrefix(err.Error(), fmt.Sprintf(`Get "%s": `, ip))
		result += fmt.Sprintf(`<tr>
			<td>%s</td>
			<td class="failed">Failed</td>
			<td>N/A</td>
			<td>%s</td>
		</tr>`, ip, shortErr)
	} else {
		result += fmt.Sprintf(`<tr>
			<td>%s</td>
			<td>Connected</td>
			<td>%d</td>
			<td>N/A</td>
		</tr>`, ip, res.StatusCode)
	}

	// End HTML Table
	result += `</table>`

	return c.HTML(http.StatusOK, result)
}

///////////////////////////////////////////////////////////////////////////////////
// PING                                                                          //
///////////////////////////////////////////////////////////////////////////////////

func handlePingAction(c echo.Context) error {
	ipFqdn := c.FormValue("ip-fqdn")

	// Sanitize the input to avoid command injection
	if strings.ContainsAny(ipFqdn, ";&|") {
		return c.String(http.StatusBadRequest, "Invalid characters in input")
	}

	// Resolve the IP address
	ipAddr, err := net.ResolveIPAddr("ip", ipFqdn)
	if err != nil {
		log.Println("Failed to resolve the IP address:", err)
		return c.String(http.StatusInternalServerError, "Failed to resolve the IP address")
	}
	log.Println("Resolved IP address:", ipAddr.IP)

	// Create an ICMP connection
	conn, err := net.DialIP("ip4:icmp", nil, ipAddr)
	if err != nil {
		log.Println("Failed to create ICMP connection:", err)
		return c.String(http.StatusInternalServerError, "Failed to create ICMP connection")
	}
	defer conn.Close()
	log.Println("ICMP connection created")

	// Set a deadline for receiving the ICMP reply
	conn.SetDeadline(time.Now().Add(time.Second * 2))

	// Send the ICMP echo request
	echoRequest := make([]byte, 8)

	// Set ICMP message type (8 bits)
	echoRequest[0] = 8 // Echo Request

	// Set ICMP message code (8 bits)
	echoRequest[1] = 0 // Code 0

	// Set ICMP message identifier (16 bits)
	identifier := 1234                       // Example identifier
	echoRequest[4] = byte(identifier >> 8)   // Higher order byte
	echoRequest[5] = byte(identifier & 0xff) // Lower order byte

	// Set ICMP message sequence number (16 bits)
	sequenceNumber := 1                          // Example sequence number
	echoRequest[6] = byte(sequenceNumber >> 8)   // Higher order byte
	echoRequest[7] = byte(sequenceNumber & 0xff) // Lower order byte

	_, err = conn.Write(echoRequest)
	if err != nil {
		log.Println("Failed to send ICMP echo request:", err)
		return c.String(http.StatusInternalServerError, "Failed to send ICMP echo request")
	}
	log.Println("ICMP echo request sent")

	// Receive the ICMP echo reply
	echoReply := make([]byte, 1500)
	_, err = conn.Read(echoReply)
	if err != nil {
		log.Println("Failed to receive ICMP echo reply:", err)
		return c.String(http.StatusInternalServerError, "Failed to receive ICMP echo reply")
	}
	log.Println("ICMP echo reply received")

	return c.String(http.StatusOK, "Ping successful")
}

