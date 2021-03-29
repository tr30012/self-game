var url = "http://127.0.0.1:8080/"
var pageContent = {}

function UpdataPageContent() {
    console.log('Запрос данных')
    return fetch(url + "get").then((resp) => {return resp.json()}).then((content) => {
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

function AddPlayerPoints(id) {
    let points = parseInt(prompt("Добавить: ", "0"))

    if (isNaN(points)) 
        return 

    pageContent.Players[id].p_points += points

    SetPageContent()
    ReloadPage()
}

function RemovePlayerPoints(id) {
    let points = parseInt(prompt("Отнять: ", "0"))

    if (isNaN(points)) 
        return 
    
    pageContent.Players[id].p_points -= points

    SetPageContent()
    ReloadPage()
}

function AnswerQuestion(id) {
    for (let t in pageContent.Questions) {
        let theme = pageContent.Questions[t]

        for (let p in theme) {
            let question = theme[p]

            if (question.Id == id) {
                if (confirm(question.q_text)) {
                    pageContent.Questions[t][p].q_answered = true
                }
            }
        }
    }

    SetPageContent()
    ReloadPage()
}

function Nothing() {}

function ReloadPage() {
    console.log('Перезагрузка')
    
    for (let p in pageContent.Players) {
        let el = document.getElementById("player" + p)
        el.textContent = pageContent.Players[p].p_points
    }

    for (let t in pageContent.Questions) {
        let theme = pageContent.Questions[t]

        for (let p in theme) {
            let question = theme[p]

            if (question.q_answered) {
                let el = document.getElementById(question.Id)
                el.className = "answered"
                el.href = "javascript:Nothing()"
                el.textContent = "БЫЛ"
            }
        }
    }
}

UpdataPageContent().then(ReloadPage())