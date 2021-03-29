var url = "http://127.0.0.1:8080/"
var pageContent = {}

function UpdataPageContent() {
    fetch(url + "get").then((resp) => {return resp.json()}).then((content) => {
        pageContent = content 
    })
}

function SetPageContent() {
    fetch(url + "set", {
        method: "POST",
        headers: {
            "Content-Type": "application/json" 
        },
        body: JSON.stringify(pageContent)
    })
}

function ReloadPage() {
    for (var i = 0; i < pageContent.Questions.length(); i++) {
        let el = Document.getElementById(String(pageContent.Questions.Id))
        console.log(el)
    }
    
}