package app

import (
	"log"
	"os"
	"path"
	"strconv"
)

func WriteMember(ac AppContext, dest string, member Member) {
	f, err := os.Create(dest)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	finalPathToPicture := path.Join(path.Dir(dest), NameToFileName(member.Metadata.NameEn)+".avif")

	// write frontmatter
	writeln(f, "---")
	writeField(f, "nameJa", member.Metadata.NameJa)
	writeField(f, "nameEn", member.Metadata.NameEn)
	writeField(f, "joinYear", strconv.Itoa(member.Metadata.JoinYear))
	writeField(f, "description", member.Metadata.Description)
	writeField(f, "picture", finalPathToPicture)
	writeField(f, "github", member.Metadata.GitHub)
	writeField(f, "twitter", member.Metadata.Twitter)
	writeField(f, "website", member.Metadata.Website)
	writeln(f, "---")

	// write body
	writeln(f, member.Body)
}

func writeField(f *os.File, key string, value string) {
	if value == "" {
		return
	}
	writeln(f, key+": "+value)
}
func writeln(f *os.File, s string) {
	_, err := f.WriteString(s + "\n")
	if err != nil {
		log.Fatalln(err)
	}
}
