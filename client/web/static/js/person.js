// +---------------------+
// + Back buttons system +
// +---------------------+

// Returns href for back button
function getBackHref() {
    let history = getPageMovingHistory()

    if (history == null || history.length <= 1) {
        return "/filmer"
    }
    return history[history.length - 2]
}

document.addEventListener("DOMContentLoaded", function () {
    let backButton = document.getElementById("back-button")
    backButton.setAttribute("href", getBackHref())
});

// update page moving history list: remove last elem
function onBackButtonClick() {
    let history = getPageMovingHistory()
    history.pop()
    setPageMovingHistory(history)
}
