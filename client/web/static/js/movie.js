// change form action last path route (e.g. "smth" in action "/path/to/smth")
function changeFormAction(formElem, newLastRoute) {
	let formAction = formElem.action.split("/")
	formAction[formAction.length - 1] = newLastRoute
	formElem.action = formAction.join("/")
}

document.addEventListener("DOMContentLoaded", function () {
	// color active icons (stared || want || watched) to red color

	// select all forms with icons
	const iconsForms = document.querySelectorAll(".movie-extra form")
	iconsForms.forEach((iconForm) => {
		// if icon data "active" is true
		if (iconForm.dataset.active === "True") {
			const iconElemSvg = iconForm.querySelector("button svg");
			// change styles for svg elem
			iconElemSvg.classList.remove('fill-static-white')
			iconElemSvg.classList.remove('fill-hover-light-red')
			iconElemSvg.classList.add('fill-static-light-red')
			iconElemSvg.classList.add('fill-hover-white')

			// set actions (stared || want || watched) for forms
			if (iconForm.id === "stared") {
				changeFormAction(iconForm, "unstar")
			} else if (iconForm.id === "want" || iconForm.id === "watched") {
				changeFormAction(iconForm, "clear")
			}
		}
	});
});

// +---------------------+
// + Back buttons system +
// +---------------------+

// Returns href for back button
function getBackHref() {
	let history = getPageMovingHistory()

	if (history == null || history.length <= 1) {
		return "/filmer"
	}
	console.log(history[history.length - 2]);
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
