// +--------------------+
// + Person back button +
// +--------------------+

document.addEventListener("DOMContentLoaded", function () {
	let backButtonURL = localStorage.getItem("person-back-url");

	if (backButtonURL == null) {
		backButtonURL = "/filmer"
	}

	let backButton = document.getElementById("back-button")
	backButton.setAttribute("href", backButtonURL)
});
