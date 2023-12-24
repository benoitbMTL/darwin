package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////
// HEALTH CHECK                                                                  //
///////////////////////////////////////////////////////////////////////////////////

func handleHealthCheckAction(c echo.Context) error {
	urls := []string{DVWA_URL, JUICESHOP_URL, PETSTORE_URL, FWB_URL, SPEEDTEST_URL, "https://www.google.com"}

	// Define a custom HTTP client with a redirect policy that returns an error
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Start HTML Table with CSS
	result := `<style>table {width: 100%; border-collapse: collapse; font-size: 0.9em;} td, th {border: 1px solid #ddd; padding: 8px;} .failed {color: red; font-weight: bold;}</style><table>
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
	ip := "https://" + FWB_MGT_IP
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

	// Execute the ping command
	cmd := exec.Command("ping", "-c", "4", "-W", "1", ipFqdn)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Check the error type
		if exitError, ok := err.(*exec.ExitError); ok {
			// Check the exit code
			if exitError.ExitCode() == 1 {
				return c.String(http.StatusInternalServerError, "The destination is not reachable")
			} else if exitError.ExitCode() == 2 {
				return c.String(http.StatusInternalServerError, "The FQDN does not resolve")
			}
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the output of the ping command
	return c.String(http.StatusOK, string(output))
}
