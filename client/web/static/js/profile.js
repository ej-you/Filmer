const sessionExpiration = document.getElementById("session-expiration");
const hours = document.getElementById("hours");
const minutes = document.getElementById("minutes");
const seconds = document.getElementById("seconds");

let expirationTime = new Date(0, 0, 0, parseInt(hours.textContent), parseInt(minutes.textContent), parseInt(seconds.textContent), 0)

function changeTime() {
	// update date instance
	expirationTime.setSeconds(expirationTime.getSeconds() - 1)

	// get updated values
	let timeHours = expirationTime.getHours()
	let timeMinutes = expirationTime.getMinutes()
	let timeSeconds = expirationTime.getSeconds()

	// check if expiration time is over
	if ((timeHours + timeMinutes + timeSeconds) == 0) {
		sessionExpiration.textContent = "Your session has already expired!"
		clearInterval(interval);
	}

	// update page view for user
	hours.textContent = timeHours
	minutes.textContent = timeMinutes
	seconds.textContent = timeSeconds
};

// setInterval(changeTime, 1000);
const interval = setInterval(changeTime, 1000);
