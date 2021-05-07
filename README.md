# gotruder

Blind SQLi script that mimics Burpsuite's intruder functionality

Specifically built for Portswigger's Web Security Academy lab: Blind SQL injection with conditional responses

## Install
`git clone https://github.com/andrewvenson/gotruder.git`
<br/>
`cd gotruder`
<br/>
`go build` or `go run main.go`

- Copy full url to `-h` arg flag
- Copy full cookie header to `-hdr` arg flag
- Copy full cookie header value to `-hdrVal` arg flag
- Set number of characters password is with `-n` arg flag
- Set full path to wordlist with `-wl` arg flag
- Set sqlinjection code within string like so "select sqli from ujustgothacked" to `-sqli` arg flag
- For the lab, set the substring start index to `iter` within your sqli arg

Read code, to understand logic & requests being made

Further iterations pending depending on increase of knowledge and skill level
