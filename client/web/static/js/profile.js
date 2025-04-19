// remove all query-params if exists (for removing passwdChangedOK query-param)
document.addEventListener("DOMContentLoaded", function () {
	const url = new URL(window.location.href);
	url.searchParams.delete("passwdChangedOK");
	window.history.replaceState({}, document.title, url.pathname + url.search);
    console.log("Param removed. New URL:", url.pathname + url.search);
});

// token expiration time countdown
const sessionExpiration = document.getElementById("session-expiration");
const days = document.getElementById("days");
const hours = document.getElementById("hours");
const minutes = document.getElementById("minutes");
const seconds = document.getElementById("seconds");

let expirationTime = new Date(0, 0, 0, parseInt(hours.textContent), parseInt(minutes.textContent), parseInt(seconds.textContent), 0)
let timeHours = expirationTime.getHours()
let timeMinutes = expirationTime.getMinutes()
let timeSeconds = expirationTime.getSeconds()

function changeTime() {
	// check if expiration time is over
	if ((timeHours + timeMinutes + timeSeconds) == 0) {
		let daysNumber = parseInt(days.textContent)
		// if days is there
		if (daysNumber != 0) {
			days.textContent = daysNumber - 1
		} else {
			sessionExpiration.textContent = "Your session has already expired!"
			clearInterval(interval);
			// change form and button to login
			const logoutForm = document.getElementById("link-button-form");
			logoutForm.action = "/filmer/user/login"
			logoutForm.method = "get"
			const logoutBtn = document.getElementById("link-button");
			logoutBtn.textContent = "Login"
		}
	}
	// update date instance
	expirationTime.setSeconds(expirationTime.getSeconds() - 1)

	// update time values
	timeHours = expirationTime.getHours()
	timeMinutes = expirationTime.getMinutes()
	timeSeconds = expirationTime.getSeconds()

	// set updated values to html
	hours.textContent = timeHours
	minutes.textContent = timeMinutes
	seconds.textContent = timeSeconds
};

const interval = setInterval(changeTime, 1000);
