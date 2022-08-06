#!/data/data/com.termux/files/usr/bin/bash
#---
PATH=/data/data/com.termux/files/home/go/src/github.com/edwindvinas/ULAPPH-Cloud-Desktop-Watson
CURL=/data/data/com.termux/files/usr/bin/curl
GO=/data/data/com.termux/files/usr/bin/go

echo 'Running ibm watson export...'
#---
echo 'Exporting "00 - Intent Router"....'
cd $PATH
$CURL -u "apiKey":"YOUR-WATSON-API-KEY"  "https://api.jp-tok.assistant.watson.cloud.ibm.com/instances/YOUR-WATSON-INSTANCE-ID/v1/workspaces/YOUR-WATSON-WORKSPACE-ID-1?version=2020-04-01&export=true" > "00 - Intent Router.json"

echo 'Beautifying JSON file...'
echo '*** 00 - Intent Router.json ***'
$PATH/beautifyJson "$PATH/00 - Intent Router.json" "$PATH/00 - Intent Router_(beautify).json"
#---
echo 'Exporting "10 - CloudPlatformAssistant"....'
cd $PATH
$CURL -u "apiKey":"YOUR-WATSON-API-KEY"  "https://api.jp-tok.assistant.watson.cloud.ibm.com/instances/YOUR-WATSON-INSTANCE-ID/v1/workspaces/YOUR-WATSON-WORKSPACE-ID-2?version=2020-04-01&export=true" > "10 - CloudPlatformAssistant.json"

echo 'Beautifying JSON file...'
echo '*** 10 - CloudPlatformAssistant.json ***'
$PATH/beautifyJson "$PATH/10 - CloudPlatformAssistant.json" "$PATH/10 - CloudPlatformAssistant_(beautify).json"
#---
echo 'Exporting "20 - TechnologyArchitect"....'
cd $PATH
$CURL -u "apiKey":"YOUR-WATSON-API-KEY"  "https://api.jp-tok.assistant.watson.cloud.ibm.com/instances/YOUR-WATSON-INSTANCE-ID/v1/workspaces/YOUR-WATSON-WORKSPACE-ID-3?version=2020-04-01&export=true" > "20 - TechnologyArchitect.json"

echo 'Beautifying JSON file...'
echo '*** 20 - TechnologyArchitect.json ***'
$PATH/beautifyJson "$PATH/20 - TechnologyArchitect.json" "$PATH/20 - TechnologyArchitect_(beautify).json"
#---
echo 'Exporting "30 - EnterpriseArchitect"....'
cd $PATH
$CURL -u "apiKey":"YOUR-WATSON-API-KEY"  "https://api.jp-tok.assistant.watson.cloud.ibm.com/instances/YOUR-WATSON-INSTANCE-ID/v1/workspaces/YOUR-WATSON-WORKSPACE-ID-4?version=2020-04-01&export=true" > "30 - EnterpriseArchitect.json"

echo 'Beautifying JSON file...'
echo '*** 30 - EnterpriseArchitect.json ***'
$GO run $PATH/beautifyJson.go "$PATH/30 - EnterpriseArchitect.json" "$PATH/30 - EnterpriseArchitect_(beautify).json"
#---
echo 'Exporting "99 - General"....'
cd $PATH
$CURL -u "apiKey":"YOUR-WATSON-API-KEY"  "https://api.jp-tok.assistant.watson.cloud.ibm.com/instances/YOUR-WATSON-INSTANCE-ID/v1/workspaces/YOUR-WATSON-WORKSPACE-ID-5?version=2020-04-01&export=true" > "99 - General.json"

echo 'Beautifying JSON file...'
echo '*** 99 - General.json ***'
$PATH/beautifyJson "$PATH/99 - General.json" "$PATH/99 - General_(beautify).json"
#---

#echo 'Creating AI Menu...'
#echo 'Saving file: ../ULAPPH-Cloud-Desktop/templates/ulapph-ai-menu.txt'
#$PATH/genUlapphAiMenu.exe "../ULAPPH-Cloud-Desktop/templates/ulapph-ai-menu.txt"
#/c/Go/bin/go run $PATH/genUlapphAiMenu.go --output ../ULAPPH-Cloud-Desktop/templates/ulapph-ai-menu.txt --inputs '00 - Intent Router.json' '10 - CloudPlatformAssistant.json' '20 - TechnologyArchitect.json' '99 - General.json'
