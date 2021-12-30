/*
This file provides fictional data for making logging functions more realistic and dynamic
*/
package main

import (
	"fmt"
	"math/rand"
)

var data = map[string]interface{}{
	"files": []string{
		"pippo",
		"start",
		"telnet",
		"ftp",
		"ex",
		"impedit",
		"dolor",
		"omnis",
		"excepturi",
		"iusto",
		"magnam",
		"pariatur",
		"repellat",
		"aliquid",
		"fuga",
		"quae",
		"optio",
		"soluta",
		"odit",
		"ad",
		"perspiciatis",
		"occaecati",
		"doloribus",
		"suscipit",
		"officia",
		"modi",
		"unde",
		"nisi",
		"ratione",
		"tempore",
		"nesciunt",
		"alias",
		"voluptatum",
		"optio",
		"fugit",
		"ab",
		"quia",
		"eos",
		"quia",
		"error",
		"sit",
		"fugiat",
		"error",
		"at",
		"reprehenderit",
		"quidem",
		"sequi",
		"repellendus",
		"corrupti",
		"illo",
		"pariatur",
		"itaque",
		"nisi",
		"architecto",
		"ipsum",
		"ipsum",
		"ab",
		"odio",
		"sint",
		"aspernatur",
		"praesentium",
		"quibusdam",
		"dolorem",
		"cum",
		"quod",
		"at",
		"ut",
		"numquam",
		"eos",
		"qui",
		"architecto",
		"ea",
		"qui",
		"tenetur",
		"ullam",
		"excepturi",
		"in",
		"commodi",
		"iure",
		"quisquam",
		"nihil",
		"deserunt",
		"sed",
		"ex",
		"amet",
		"illo",
		"itaque",
		"earum",
		"perferendis",
		"suscipit",
		"beatae",
		"deleniti",
		"alias",
		"ut",
		"cupiditate",
		"laudantium",
		"vero",
		"eligendi",
		"consectetur",
		"perferendis",
		"dolores",
		"facere",
		"alias",
		"delectus",
	},
	"extensions": []string{".exe",
		".dat",
		".bin",
		"",
		".avi",
		".bmp",
		".css",
		".csv",
		".doc",
		".docx",
		".flac",
		".html",
		".jpeg",
		".jpg",
		".js",
		".json",
		".key",
		".mov",
		".mp3",
		".mp4",
		".odp",
		".ods",
		".odt",
		".pages",
		".pdf",
		".png",
		".pptx",
		".tiff",
		".txt",
		".wav",
		".webm",
		".xls",
		".xlsx",
	},
	"disks":       []string{"/dev/sda1", "/dev/sda2", "/dev/sda3", "/dev/sda4", "/dev/sda5"},
	"filesystems": []string{"ext3", "ext4", "vfs", "vfat", "ntfs", "xfs"},
	"urls": []string{
		"http://allen.info",
		"http://amir.com",
		"http://amy.info",
		"http://anastacio.name",
		"http://angela.org",
		"http://blanche.net",
		"http://brenden.org",
		"http://brittany.name",
		"http://cara.biz",
		"http://carmelo.biz",
		"http://clementine.com",
		"http://clinton.name",
		"http://corine.info",
		"http://dario.net",
		"http://darrion.com",
		"http://deborah.org",
		"http://deshaun.com",
		"http://eulalia.info",
		"http://fanny.org",
		"http://felicity.biz",
		"http://glenda.info",
		"http://hayden.info",
		"http://ian.name",
		"http://jarrett.org",
		"http://jewel.biz",
		"http://joe.name",
		"http://johann.com",
		"http://kaylah.org",
		"http://kayley.net",
		"http://kyleigh.info",
		"http://levi.org",
		"http://loren.info",
		"http://lowell.org",
		"http://lucio.org",
		"http://madalyn.biz",
		"http://manuel.org",
		"http://mariela.net",
		"http://melisa.net",
		"http://nettie.biz",
		"http://nichole.net",
		"http://pamela.biz",
		"http://rafael.net",
		"http://ricardo.org",
		"http://richard.biz",
		"http://rocky.net",
		"http://rose.net",
		"http://royal.biz",
		"http://salvatore.biz",
		"http://sandrine.net",
		"http://sedrick.biz",
		"http://sophia.com",
		"http://susanna.net",
		"http://tara.name",
		"http://taya.name",
		"http://tiara.org",
		"http://triston.net",
		"http://yasmin.org",
		"https://adelia.net",
		"https://alysha.biz",
		"https://ari.com",
		"https://blaze.info",
		"https://camden.info",
		"https://camila.net",
		"https://cordie.com",
		"https://daren.info",
		"https://dejon.net",
		"https://delmer.com",
		"https://elliot.biz",
		"https://elliot.org",
		"https://esmeralda.biz",
		"https://fredy.com",
		"https://garfield.com",
		"https://georgette.net",
		"https://ian.biz",
		"https://jedediah.info",
		"https://jewel.net",
		"https://kaia.com",
		"https://kayley.info",
		"https://kennedi.net",
		"https://kody.net",
		"https://leslie.com",
		"https://marjorie.biz",
		"https://maryjane.com",
		"https://mathew.net",
		"https://mia.com",
		"https://misty.name",
		"https://omari.info",
		"https://orie.info",
		"https://orion.org",
		"https://paige.org",
		"https://pascale.name",
		"https://pattie.com",
		"https://reva.biz",
		"https://skyla.name",
		"https://solon.biz",
		"https://tamara.biz",
		"https://terrance.biz",
		"https://terrence.name",
		"https://tremaine.com",
		"https://vella.net",
	},
}

// getFilename returns a random filename
func getFilename() string {
	files := data["files"].([]string)
	extensions := data["extensions"].([]string)
	return fmt.Sprintf("%s%s", files[rand.Intn(len(files))], extensions[rand.Intn(len(extensions))])
}

// getDisk returns a random disk device name
func getDisk() string {
	disks := data["disks"].([]string)
	return disks[rand.Intn(len(disks))]
}

// getFilesystem returns a random filesystem standard name
func getFilesystem() string {
	filesystems := data["filesystems"].([]string)
	return filesystems[rand.Intn(len(filesystems))]
}

// getUrl returns a random URL
func getUrl() string {
	urls := data["urls"].([]string)
	return urls[rand.Intn(len(urls))]
}
