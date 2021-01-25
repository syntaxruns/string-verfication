this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte(fmt.Sprintf("\033]0; TITLE\007")))
    this.conn.Write([]byte("\x1b[0m"))
    verification, err := this.ReadLine(false)
    if err != nil {
        return
    }
    authcode, err := ioutil.ReadFile("FILENAME.txt")
    if err != nil {
        return
    }

    text := string(authcode)

    if verification == text {
        goto login
    } else {
        return
    }
    login: