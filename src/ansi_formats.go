package main

type ansiFormats struct {
	linechange            string
	left                  string
	right                 string
	creset                string
	clearOEL              string
	saveCursorPosition    string
	restoreCursorPosition string
	title                 string
	colorSingle           string
	colorFull             string
	colorTransparent      string
}

func (c *ansiFormats) init(shell string) {
	switch shell {
	case zsh:
		c.linechange = "%%{\x1b[%d%s%%}"
		c.left = "%%{\x1b[%dC%%}"
		c.right = "%%{\x1b[%dD%%}"
		c.creset = "%{\x1b[0m%}"
		c.clearOEL = "%{\x1b[K%}"
		c.saveCursorPosition = "%{\x1b7%}"
		c.restoreCursorPosition = "%{\x1b8%}"
		c.title = "%%{\033]0;%s\007%%}"
		c.colorSingle = "%%{\x1b[%sm%%}%s%%{\x1b[0m%%}"
		c.colorFull = "%%{\x1b[%sm\x1b[%sm%%}%s%%{\x1b[0m%%}"
		c.colorTransparent = "%%{\x1b[%s;49m\x1b[7m%%}%s%%{\x1b[m\x1b[0m%%}"
	case bash:
		c.linechange = "\\[\x1b[%d%s\\]"
		c.left = "\\[\x1b[%dC\\]"
		c.right = "\\[\x1b[%dD\\]"
		c.creset = "\\[\x1b[0m\\]"
		c.clearOEL = "\\[\x1b[K\\]"
		c.saveCursorPosition = "\\[\x1b7\\]"
		c.restoreCursorPosition = "\\[\x1b8\\]"
		c.title = "\\[\033]0;%s\007\\]"
		c.colorSingle = "\\[\x1b[%sm\\]%s\\[\x1b[0m\\]"
		c.colorFull = "\\[\x1b[%sm\x1b[%sm\\]%s\\[\x1b[0m\\]"
		c.colorTransparent = "\\[\x1b[%s;49m\x1b[7m\\]%s\\[\x1b[m\x1b[0m\\]"
	default:
		c.linechange = "\x1b[%d%s"
		c.left = "\x1b[%dC"
		c.right = "\x1b[%dD"
		c.creset = "\x1b[0m"
		c.clearOEL = "\x1b[K"
		c.saveCursorPosition = "\x1b7"
		c.restoreCursorPosition = "\x1b8"
		c.title = "\033]0;%s\007"
		c.colorSingle = "\x1b[%sm%s\x1b[0m"
		c.colorFull = "\x1b[%sm\x1b[%sm%s\x1b[0m"
		c.colorTransparent = "\x1b[%s;49m\x1b[7m%s\x1b[m\x1b[0m"
	}
}
