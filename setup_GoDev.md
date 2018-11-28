## Step 1
- go get golang.org/x/tools/cmd/goimports
## Step 2
- Open Sublime and press cmd + shift + p 
- Search for and select "Package Control: Add Repository"
- At the bottom of sublime it promots you to enter url, add https://github.com/DisposaBoy/GoSublime.git
- Again press cmd + shift + p
- Search for and select "Install Package"
- Then search for "GoSublime"

# Step 3
Now in Sublime -> Preference -> Package Setting -> GoSUblime -> Setting - User -> add following text
```
{
   "env":{"GOPATH": "$(go env GOPATH)"},
    "fmt_cmd": ["goimports"],
    "on_save": [{
    "cmd": "gs9o_open", "args": {
    //"run": ["sh","go build . && go test -i && go vet && golint ."],
    "focus_view": false
}}],
    "autocomplete_closures": true,
    "complete_builtins": true,
}

```
## Step 5
- Test it, write a basic main.go file and import a package but dont use it (like "fmt") On savingn it should go away
