package dat

type CharacterClass []rune

var CharacterClasses = map[string]CharacterClass{
	"xdigit": CharacterClass("abcdefABCDEF0123456789"),
	"upper":  CharacterClass("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"space":  CharacterClass(" \t\r\n\v\f"),
	"punct":  CharacterClass("][!\"#$%&'()*+,./:;<=>?@\\^_`{|}~-"),
	"print":  CharacterClass(""), // TODO: [\x20-\x7E]
	"lower":  CharacterClass("abcdefghijklmnopqrstuvwxyz"),
	"graph":  CharacterClass(""), //TODO: see [:print:]
	"digit":  CharacterClass("0123456789"),
	"cntrl":  CharacterClass(""), // TODO: [\x00-\x1F\x7F]
	"blank":  CharacterClass(" \t"),
	"alpha":  CharacterClass("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"alnum":  CharacterClass("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
}
