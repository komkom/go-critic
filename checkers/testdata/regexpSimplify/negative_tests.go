package checker_test

import (
	"regexp"
)

func shouldNotSimplify() {
	regexp.MustCompile(``)
	regexp.MustCompile(`[*]`)
	regexp.MustCompile(`[|]`)
	regexp.MustCompile(`[?]`)
	regexp.MustCompile(`[)]`)
}

func noCommonPrefixSuffix() {
	regexp.MustCompile(`aa|a`) // Too short to bother

	regexp.MustCompile(`foo|fa`) // Doesn't match our criteria
	regexp.MustCompile(`1path|2path`)

	regexp.MustCompile(`(?:http|https|ftp)://`) // More than 2 alternatives
}

func cantSimplify() {
	// Most of these patterns come from the projects available at GitHub.

	regexp.MustCompile(`\s*\{weight=(\d+)\}\s*`)
	regexp.MustCompile(`\{inherits=(\d+)\}`)
	regexp.MustCompile(`[<>]+`)
	regexp.MustCompile(`[.?,!;:@#$%^&*()]+`)
	regexp.MustCompile(`^\d{1,7}$`)
	regexp.MustCompile(`'([ (\[{<])"`)
	regexp.MustCompile(`^(")`)
	regexp.MustCompile(`( ")`)
	regexp.MustCompile(`unifi_devices{site="Default"} 1`)
	regexp.MustCompile(`unifi_devices_adopted{site="Default"} 1`)
	regexp.MustCompile(`unifi_devices_unadopted{site="Default"} 0`)
	regexp.MustCompile(`^((\").*(\"))`)
	regexp.MustCompile(`/api/internal/login`)
	regexp.MustCompile(`/api/organizations\?limit=100`)
	regexp.MustCompile(`[áàảãạấầẩẫậâăắằẳẵặ]`)
	regexp.MustCompile(`\.(com|com\.\w{2})$`)
	regexp.MustCompile(`\.(gov|gov\.\w{2})$`)
	regexp.MustCompile(`[\xC0-\xC6]`)
	regexp.MustCompile(`[\xE0-\xE6]`)
	regexp.MustCompile(`[\xC8-\xCB]`)
	regexp.MustCompile(`[\xE8-\xEB]`)
	regexp.MustCompile(`\bsample\b`)
	regexp.MustCompile(`\b(720p|1080p|hdtv|x264|dts|bluray)\b.*`)
	regexp.MustCompile(`(?i)\bcode\b`)
	regexp.MustCompile(`\p{Cyrillic}`)
	regexp.MustCompile(`--(?P<var_name>[\\w-]+?):\\s+?(?P<var_val>.+?);`)
	regexp.MustCompile(`^.+_rsa$`)
	regexp.MustCompile(`^.+_dsa.*$`)
	regexp.MustCompile(`^.+_ed25519$`)
	regexp.MustCompile(`^.+_ecdsa$`)
	regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	regexp.MustCompile(`a[^rst]c`)
	regexp.MustCompile(`'''(.*)'''`)
	regexp.MustCompile(`<li class="neirong2">信息标题.*<b>([^<]+)<`)
	regexp.MustCompile(`<li class="neirong2">信息来源[^>]+">([^<]+)<`)
	regexp.MustCompile(`\s*:\s*`)
	regexp.MustCompile(`(?m)^[ \t]*(#+)\s+`)
	regexp.MustCompile(`(?m)^h([0-6])\.(.*)$`)
	regexp.MustCompile(`<==\sPlayerInventory\.GetPlayerCardsV3\(\d*\)`)
	regexp.MustCompile(`^ *(#{1,6}) *([^\n]+?) *#* *(?:\n|$)`)
	regexp.MustCompile(`^[^\n]+`)
	regexp.MustCompile(`^\n+`)
	regexp.MustCompile(`^((?: {4}|\t)[^\n]+\n*)+`)
	regexp.MustCompile(`^(int)`)
	regexp.MustCompile(`^(\+)`)
	regexp.MustCompile(`[^a-zA-Z0-9]`)
	regexp.MustCompile(`^absolute(\.|-).*`)
	regexp.MustCompile(`^metric .*`)
	regexp.MustCompile(`(?m)^([\s\t]*)([\*\-\+]|\d\.)\s+`)
	regexp.MustCompile(`\n={2,}`)
	regexp.MustCompile(`~~`)
	regexp.MustCompile(`[aeiou][^aeiou]`)
	regexp.MustCompile(`.*sses$`)
	regexp.MustCompile(`CIFS Session: (?P<sessions>\d+)`)
	regexp.MustCompile(`Share \(unique mount targets\): (?P<shares>\d+)`)
	regexp.MustCompile(`SMB Request/Response Buffer: (?P<smbBuffer>\d+) Pool size: (?P<smbPoolSize>\d+)`)
	regexp.MustCompile(`(?is)<a.+?</a>`)
	regexp.MustCompile(`(?is)\(.*?\)`)
	regexp.MustCompile(`<a href="#.+?\|`)
	regexp.MustCompile(`^(?:mister )(.*)$`)
	regexp.MustCompile(`(?ims)<!DOCTYPE.*?>`)
	regexp.MustCompile(`(?ims)<!--.*?-->`)
	regexp.MustCompile(`(?ims)<script.*?>.*?</script>`)
	regexp.MustCompile(`[a-z]+`)
	regexp.MustCompile(`[^a-z]+`)
	regexp.MustCompile(`Acl`)
	regexp.MustCompile(`Adm([^i]|$)`)
	regexp.MustCompile(`Aes`)
	regexp.MustCompile(`(es|ed|ing)$`)
	regexp.MustCompile(`[^[:alpha:]]`)
	regexp.MustCompile(`ESCAPE_([[:alnum:]]+)`)
	regexp.MustCompile(`[.]|,|\s/`)
	regexp.MustCompile(`(?imsU)\[quote(?:=[^\]]+)?\](.+)\[/quote\]`)
	regexp.MustCompile(`(?imU)^(.*)$`)
	regexp.MustCompile(`<div class="price_m" id='original-price'>.+\n.+<\/span>(.+)<\/div>`)
	regexp.MustCompile(`开 本：(\d+)开`)
	regexp.MustCompile(`^((inteiro)|(real)|(caractere)|(lógico))(\s*):`)
	regexp.MustCompile(`^\d*H[A-Z0-9]*$`)
	regexp.MustCompile(`^LI[A-Z0-9]*$`)
	regexp.MustCompile(`^C[A-Z0-9]*$`)
	regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	regexp.MustCompile(`}\n+$`)
	regexp.MustCompile(`^\s*events\s*{`)
	regexp.MustCompile(`^\s*http\s*{`)
	regexp.MustCompile(`(?i)windows nt`)
	regexp.MustCompile(`(?i)windows phone`)
	regexp.MustCompile(`^.* ENGINE=.*$/\)`)
	regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z-0-9.]+\.[a-zA-Z0-9]+`)
	regexp.MustCompile("he|ll|o+")
	regexp.MustCompile(`[/.@!~#$%^&*:";?\\+=-_,{}\[\]<>！￥…（）—=、“”：；？。，《》]`)
	regexp.MustCompile(`^(\S*) (\S*) (\d*) (\S*) IP(\d) (\S*)`)
	regexp.MustCompile(`\s*Version:\s*(.+)$`)
	regexp.MustCompile(`Domain Name\.*: *(.+)`)
	regexp.MustCompile("[^A-Z`'\\s]")
	regexp.MustCompile(`(Email|EmailAddress)\(\)`)
	regexp.MustCompile(`\*\*[^*]*\*\*`)
	regexp.MustCompile(`(\*[^ ][^*]*\*)`)
	regexp.MustCompile(`\+[^+]*\+`)
	regexp.MustCompile(`\bMac[A-Za-z]{2,}[^aciozj]\b`)
	regexp.MustCompile(`\bMc`)
	regexp.MustCompile(`\b(Ma?c)([A-Za-z]+)`)
	regexp.MustCompile(`(?m)^Created: *(.*?)$`)
	regexp.MustCompile(`#\+(\w+): (.*)`)
	regexp.MustCompile(`^(\*+)(?: +(CANCELED|DONE|TODO))?(?: +(\[#.\]))?(?: +(.*?))?(?:(:[a-zA-Z0-9_@#%:]+:))??[ \t]*$`)
	regexp.MustCompile(`^[ \t]*#`)
	regexp.MustCompile(`(?i)\+BEGIN_(CENTER|COMMENT|EXAMPLE|QUOTE|SRC|VERSE)`)
	regexp.MustCompile(`(\d+\.\d+) (\w+)(\(.*) <(.+)>`)
	regexp.MustCompile(`(\d+\.\d+) (\w+)(\(.*)`)
	regexp.MustCompile(`^(\s+\S+:\s+.*)$`)
	regexp.MustCompile(`abc{2,5}d?e*f+`)
	regexp.MustCompile(`a+b+`)
	regexp.MustCompile(`(?m:^%(\.\d)?[sdfgtq]$)`)
	regexp.MustCompile(`(?m:^[ \t]*%(\.\d)?[sdfgtq]$)`)
	regexp.MustCompile(`(?i)<html.*/head>`)
	regexp.MustCompile(` COLLATE ([^ ]+)`)
	regexp.MustCompile(` CHARACTER SET ([^ ]+)`)
	regexp.MustCompile(`(?:(.+); )?InnoDB free: .*`)
	regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	regexp.MustCompile(`\ACK2txt\n$`)
	regexp.MustCompile(`^\d{3}[- ]?\d{2}[- ]?\d{4}$`)
	regexp.MustCompile(`<arrmsg1>([\w\W]+?)</arrmsg1>`)
	regexp.MustCompile("<p(.*?)>")
	regexp.MustCompile(`^(?i)(\s*insert\s+into\s+)`)
	regexp.MustCompile(`^(?i)\((\s*\w+(?:\s*,\s*\w+){1,100})\s*\)\s*`)
	regexp.MustCompile(`0\d-[2-9]\d{3}-\d{4}`)
	regexp.MustCompile(`0\d{2}-[2-9]\d{2}-\d{4}`)
	regexp.MustCompile(`0\d{3}-[2-9]\d-\d{4}`)
	regexp.MustCompile(`0\d{4}-[2-9]-\d{4}`)
	regexp.MustCompile(`(?i)(refer):\s+(.*?)(\s|$)`)
	regexp.MustCompile(`(?m:^)(\s+)?(?i)(Whois server|whois):\s+(.*?)(\s|$)`)
	regexp.MustCompile("@(?i:article){.*")
	regexp.MustCompile(`\blimit \?(?:, ?\?| offset \?)?`)
	regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,4}$`)
	regexp.MustCompile(`^(www.|https://|http://)*[A-Za-z0-9._%+\-]+\.[com|org|edu|net]{3}$`)
	regexp.MustCompile(`^4\d{12}(\d{3})?$`)
	regexp.MustCompile(`^(5[1-5]\d{4}|677189)\d{10}$`)
	regexp.MustCompile(`^(6011|65\d{2}|64[4-9]\d)\d{12}|(62\d{14})$`)
	regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
	regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`)
	regexp.MustCompile(`^\d{8}T\d{6}Z$`)
	regexp.MustCompile(`^\w+\(.*\)$`)
	regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
	regexp.MustCompile(`(?i)^d('ye)$`)
	regexp.MustCompile(`^/video/([\w\-]{6,12})\.json$`)
	regexp.MustCompile(`(\n|\r|\r\n)$`)
	regexp.MustCompile(`(?m)^@@@@\n`)
	regexp.MustCompile(`agggtaaa|tttaccct`)
	regexp.MustCompile(`[cgt]gggtaaa|tttaccc[acg]`)
	regexp.MustCompile(`a[act]ggtaaa|tttacc[agt]t`)
	regexp.MustCompile(`(?i:http).*.git`)
	regexp.MustCompile(`kitsu.io/users/(.*?)/library`)
	regexp.MustCompile(`k8s_.*\.metadata\.name$`)
	regexp.MustCompile(`k8s_\w+_\w+_deployment\.spec\.selector\.match_labels$`)
	regexp.MustCompile(`[!&]([^=!&?[)]+)|\[\[(.*?)\]\]`)
	regexp.MustCompile(`^(?:\"|\')\w+(?:\"|\')$`)
	regexp.MustCompile(`\\b[A-Fa-f0-9]{32}\\b`)
	regexp.MustCompile(`(?i)(s?(\d{1,2}))[ex]`)
	regexp.MustCompile(`(?i)([ex](\d{2})(?:\d|$))`)
	regexp.MustCompile(`(-\s+(\d+)(?:\d|$))`)
	regexp.MustCompile(` edge/(\d+)\.(\d+)`)
	regexp.MustCompile(`^.+@.+\..+$`)
	regexp.MustCompile(`^[\w\d]{3,30}$`)
	regexp.MustCompile(`^\$bgm\s+(\[.*\])$`)
	regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	regexp.MustCompile(`;[^=;{}]+;`)
	regexp.MustCompile(`\s(if|else|while|catch)\s*([^{;]+;|;)`)
	regexp.MustCompile(`^[^ ]*apache2`)
	regexp.MustCompile(`^(?:[-+]?(?:0|[1-9]\d*))$`)
	regexp.MustCompile(`<(.|[\r\n])*?>`)
	regexp.MustCompile(`^(::f{4}:)?10\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	regexp.MustCompile(`(?i)\([^)]*remaster[^)]*\)$`)
	regexp.MustCompile(`>([a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+)<`)
	regexp.MustCompile(`(?i)\([^)]*mix[^)]*\)$`)
}