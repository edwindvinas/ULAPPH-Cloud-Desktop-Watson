package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strings"
	"html/template"
	"sort"
	"unicode"
	"path/filepath"
)

/*
{
  "name": "99 - General",
  "intents": [
    {
      "intent": "thesaurus",
      "examples": [
        {
          "text": "what are the synonyms for world"
        }, 
*/
//edwinxxx
type WatsonWorkspace struct {
    Name    string `json:"name,omitempty"`
	Intents    []WatsonIntents `json:"intents,omitempty"`
}
type WatsonIntents struct {
    Intent    string `json:"intent,omitempty"`
	Examples    []IntentExamples `json:"examples,omitempty"`
}
type IntentExamples struct {
    Text    string `json:"text,omitempty"`
}

/*
	    <li>
        	<a href="#page" onmouseenter="playAudio();">Cloud</a>   
            <ul class="sublist">
				<li onmouseenter="playAudio();">Google Office    
				  <span class="arrow"></span>
					<ul class="sublist-menu">

						<li>
							<a href="#page" onmouseenter="playAudio();" onclick="window.open('https://gsuite.google.com/','_blank'); return false">G-Suite</a>
						</li>
						<li class="divider"></li>
		
					</ul>
				</li>
				<li class="divider"></li>               
			</ul>
		</li>
*/
type TEMPSTRUCT struct {
	NUM_FILLER1 int
    STR_FILLER1 string
	STR_FILLER2 string
	STR_FILLER3 string
	HTM_FILLER1 template.HTML
	BOOL_FILLER1 bool
}

type WatsonSorter []WatsonIntents

func (a WatsonSorter) Len() int           { return len(a) }
func (a WatsonSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a WatsonSorter) Less(i, j int) bool { return a[i].Intent < a[j].Intent }

type WatsonSorterEx []IntentExamples

func (a WatsonSorterEx) Len() int           { return len(a) }
func (a WatsonSorterEx) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a WatsonSorterEx) Less(i, j int) bool { return a[i].Text < a[j].Text }

func main() {
	totArgs := len(os.Args)
	fmt.Printf("TOTARGS: %v\n", totArgs)
	if totArgs <= 1 {
		fmt.Printf("ERROR: Unexpected number of params: %v", totArgs)
		fmt.Printf("./genUlapphAiMenu --output ../ULAPPH-Cloud-Desktop/templates/ulapph-ai-menu.txt --inputs '00 - Intent Router.json' '10 - CloudPlatformAssistant.json' '20 - TechnologyArchitect.json' '99 - General.json'")
		panic(fmt.Errorf("ERROR in arguments!"))
	}
	outputFile := os.Args[2] //file
	fmt.Printf("Output: %v\n", outputFile)
	var buffer bytes.Buffer
	for i:=4;i<=totArgs-1;i++ {
		buffer.WriteString(fmt.Sprintf("	    <li>\n"))
		inputFile := os.Args[i] //file
		mName :=  strings.Split(inputFile, " - ")
		mNameText := strings.Replace(mName[1], ".json", "", -1)
		fileOnly := filepath.Base(inputFile)
		sName :=  strings.Split(fileOnly, ".json")
		sNameSkill := sName[0]
		//
		var cs []string
		for _, r := range mNameText {
			if unicode.IsUpper(r) == true {
				cs = append(cs, string(r))	
			}
		}
		mNameText2 := strings.Join(cs,"")
		//
		buffer.WriteString(fmt.Sprintf("        	<a href=\"#page\" onmouseenter=\"playAudio();\" onclick=\"setBotSkillName('%v');return false;\">AI-%v</a>\n", sNameSkill, mNameText2))
		buffer.WriteString(fmt.Sprintf("            <ul class=\"sublist\">\n"))
		g := TEMPSTRUCT {
			NUM_FILLER1: i,
		}
		if err := htmlBotHdr.Execute(&buffer, &g); err != nil {
		  panic(err)
		}
		// Open our jsonFile
		jsonFile, err := os.Open(inputFile)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Processing input file: %v\n", inputFile)	
		var watson WatsonWorkspace
		err = json.Unmarshal(byteValue , &watson)
		if err != nil {
			panic(err)
		}
		name := watson.Name
		intents := watson.Intents
		sort.Sort(WatsonSorter(intents))
		fmt.Printf("Name: %v\n", name)
		fmt.Printf("Intents: %v\n", intents)
		for i := 0; i < len(intents); i++ {
			//fmt.Printf("Intent[%v]: %v\n", i, intents[i].Intent)
			//fmt.Printf("Examples: %v\n", intents[i].Examples)
			buffer.WriteString(fmt.Sprintf("				<li onmouseenter=\"playAudio();\">%v\n", intents[i].Intent))
			buffer.WriteString(fmt.Sprintf("				  <span class=\"arrow\"></span>\n"))
			buffer.WriteString(fmt.Sprintf("					<ul class=\"sublist-menu\">\n"))
			//list examples
			examples := intents[i].Examples
			sort.Sort(WatsonSorterEx(examples))
			for j := 0; j < len(examples); j++ {
				buffer.WriteString(fmt.Sprintf("						<li>\n"))
				buffer.WriteString(fmt.Sprintf("						<a href=\"#page\" onmouseenter=\"playAudio();\" onclick=\"newBotMessage('%v'); return false\">%v</a>\n", examples[j].Text, examples[j].Text))
				buffer.WriteString(fmt.Sprintf("						</li>\n"))
				buffer.WriteString(fmt.Sprintf("						<li class=\"divider\"></li>\n"))
			}
			buffer.WriteString(fmt.Sprintf("					</ul>\n"))
			buffer.WriteString(fmt.Sprintf("				</li>\n"))
			buffer.WriteString(fmt.Sprintf("				<li class=\"divider\"></li>\n"))
		}
		buffer.WriteString(fmt.Sprintf("            </ul>\n"))
		buffer.WriteString(fmt.Sprintf("	    </li>\n"))
	
	}
	//Add hard coded local commands
	g := TEMPSTRUCT {
		NUM_FILLER1: 0,
	}
	if err := htmlBotHdr2.Execute(&buffer, &g); err != nil {
	  panic(err)
	}
	fmt.Printf("\n############################\n")
	fmt.Printf("%v\n", buffer.String())
    f, err := os.Create(outputFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    n2, err := f.Write(buffer.Bytes())
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
	fmt.Println(n2, "bytes written successfully")
}

var htmlBotHdr = template.Must(template.New("htmlBotHdr").Parse(htmlBotHdrA))
const htmlBotHdrA = `
				<li>
					<input type="text" id="newJSWM{{.NUM_FILLER1}}" onfocus="disableKeys();" onkeypress="return searchKeyPress(event);" autocomplete="on" autofocus="autofocus"/>
				</li>
				<li>
					<input type="button" value="Chat" id="btnSearch{{.NUM_FILLER1}}"  onclick="newBotMessage2({{.NUM_FILLER1}});" />
					<script>
					function disableKeys() {
						localStorage["btnSearch{{.NUM_FILLER1}}"] == "active";
					}
					function searchKeyPress(e)
					{
						localStorage["btnSearch{{.NUM_FILLER1}}"] = "inactive";
						//remove event listener
						removeEventListeners();
						// look for window.event in case event isnt passed in
						e = e || window.event;
						if (e.keyCode == 13)
						{
							document.getElementById('btnSearch{{.NUM_FILLER1}}').click();
							return false;
						}
						return true;
					}
					var audio1 = document.getElementById("audioID");
					function playAudio() {
						audio1.play();
					}
					</script>
				</li>
				<li class="divider"></li>
`

var htmlBotHdr2 = template.Must(template.New("htmlBotHdr2").Parse(htmlBotHdrA2))
const htmlBotHdrA2 = `
	    <li>
        	<a href="#page" onmouseenter="playAudio();">AI-Built-in</a>
            <ul class="sublist">
				<li>
					<input type="text" id="newJSWM100" onfocus="disableKeys();" onkeypress="return searchKeyPress(event);" autocomplete="on" autofocus="autofocus"/>
				</li>
				<li>
					<input type="button" value="Chat" id="btnSearch100"  onclick="newBotMessage2( 100 );" />
					<script>
					function disableKeys() {
						localStorage["btnSearch100"] == "active";
					}
					function searchKeyPress(e)
					{
						localStorage["btnSearch100"] = "inactive";
						
						removeEventListeners();
						
						e = e || window.event;
						if (e.keyCode == 13)
						{
							document.getElementById('btnSearch100').click();
							return false;
						}
						return true;
					}
					var audio1 = document.getElementById("audioID");
					function playAudio() {
						audio1.play();
					}
					</script>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">About
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="infoWorld(); return false">Show Info</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="helloWorld(); return false">Say Hello World!</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="hiWorld(); return false">Say Hi!</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Online Markets
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="shopGroceries(); return false">Show Online Groceries</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="shopMeatsSeafoods(); return false">Show Meats/Seafoods Stores</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="shopFruitsVegetables(); return false">Show Fruits & Veg Stores</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="shopStores(); return false">Show Online Stores</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Productivity
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="openWindow('https://www.youtube.com/results?search_query=solfeggio&sp=CAM%253D','Music'); return false">Solfeggio Music</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="worldtimeBuddy(); return false">Compare Timezones</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newCountdownTimer(); return false">New Countdown Timer</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Tips
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="showTips(); return false">Give me some tips!</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Dictation
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="dictationStart(); return false">Dictation Start</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="dictationSave(); return false">Dictation Save</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="dictationSave(); return false">Dictation Stop</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">CCTV
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="cctvStream(); return false">CCTV Stream On/Off Switch</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="cctvReview(); return false">CCTV Review</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="cctvLatest(); return false">CCTV Latest</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="cctvOldest(); return false">CCTV Oldest</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">News
				  <span class="arrow"></span>
					<ul class="sublist-menu">
				<li>
					<input type="text" id="newJSWM101" onfocus="disableKeys();" onkeypress="return searchKeyPress(event);" autocomplete="on" autofocus="autofocus"/>
				</li>
				<li>
					<input type="button" value="SetTopic" id="btnSearch101"  onclick="funcSetTopicFromInput();" />
					<script>
					function disableKeys() {
						localStorage["btnSearch101"] == "active";
					}
					function searchKeyPress(e)
					{
						localStorage["btnSearch101"] = "inactive";
						
						removeEventListeners();
						
						e = e || window.event;
						if (e.keyCode == 13)
						{
							document.getElementById('btnSearch101').click();
							return false;
						}
						return true;
					}
					var audio1 = document.getElementById("audioID");
					function playAudio() {
						audio1.play();
					}
					</script>
				</li>
				<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newsStream(); return false">News Stream On/Off Switch</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newsGlobal(); return false">Global News</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('God'); return false">show news God</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('planet'); return false">show news planet</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('air'); return false">show news air</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('energy'); return false">show news energy</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('education'); return false">show news education</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('philippines'); return false">show news Philippines</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('accenture'); return false">show news accenture</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('asia'); return false">show news asia</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('america'); return false">show news america</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('africa'); return false">show news africa</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('europe'); return false">show news europe</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('software development'); return false">show news software development</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('technology'); return false">show news technology</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('business'); return false">show news business</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('B2B'); return false">show news B2B</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('B2C'); return false">show news B2C</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('conference'); return false">show news conference</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('upcoming events'); return false">show news upcoming events</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="funcSetTopic('opensource'); return false">show news opensource</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Livestream
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newsLiveStream(); return false">Youtube News Livestreams</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Control
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="stopTalking(); return false">Stop Talking</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
				<li onmouseenter="playAudio();">Desktop
				  <span class="arrow"></span>
					<ul class="sublist-menu">
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newUWM(); return false">New Desktop</a>
						</li>
						<li class="divider"></li>
						<li>
						<a href="#page" onmouseenter="playAudio();" onclick="newNote(); return false">New Note</a>
						</li>
						<li class="divider"></li>
					</ul>
				</li>
				<li class="divider"></li>
            </ul>
	    </li>
`
