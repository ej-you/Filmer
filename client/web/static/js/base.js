const minWidth = 750
let iconIsSet = false

// +---------------+
// + Error message +
// +---------------+

document.addEventListener("DOMContentLoaded", function () {
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);

	let errorMessage = searchParams.get("message")
	let errorStatusCode = searchParams.get("statusCode")

	// if error message query-params is presented
	if (errorMessage !== null && errorStatusCode !== null) {
		searchParams.delete("message")
		searchParams.delete("statusCode")

		// if other query-params is presented
		if (searchParams.toString() !== "") {
			window.history.replaceState({}, document.title, `${location.pathname}?${searchParams.toString()}`);
		} else {
			window.history.replaceState({}, document.title, location.pathname);
		}

		const elemErrorMessage = document.getElementById("error-message");
		const elemErrorCode = document.getElementById("error-message-code");
		const elemErrorText = document.getElementById("error-message-text");

		elemErrorCode.textContent = errorStatusCode
		elemErrorText.textContent = errorMessage
		elemErrorMessage.style.display = "block"

		setTimeout(() => {
			console.log("Hide error message")
			elemErrorMessage.style.display = "none"
		}, 3000);
	}
});


// +--------------+
// + Header icons +
// +--------------+

// switch menu buttons from text to icons
function toIcons() {
	console.log(`switch menu buttons from text to icons`);

	document.querySelectorAll(".nav-menu-button-elem a").forEach((menuTextLink) => {
		let linkID = menuTextLink.id;

		if (!linkID) {
			// skip iteration like continue
			return
		}
		switch (linkID) {
			case "menu-search-link":
				menuTextLink.innerHTML = `<svg fill="#ceceb5" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg"><path d="M27.414,24.586l-5.077-5.077C23.386,17.928,24,16.035,24,14c0-5.514-4.486-10-10-10S4,8.486,4,14s4.486,10,10,10c2.035,0,3.928-0.614,5.509-1.663l5.077,5.077c0.78,0.781,2.048,0.781,2.828,0C28.195,26.633,28.195,25.367,27.414,24.586z M7,14c0-3.86,3.14-7,7-7s7,3.14,7,7s-3.14,7-7,7S7,17.86,7,14z"/></svg>`;
				break;
			case "menu-stared-link":
				menuTextLink.innerHTML = `<svg fill="#ceceb5" height="48" viewBox="0 0 48 48" width="48" xmlns="http://www.w3.org/2000/svg"><path d="M34 6H14c-2.21 0-3.98 1.79-3.98 4L10 42l14-6 14 6V10c0-2.21-1.79-4-4-4z"/><path d="M0 0h48v48H0z" fill="none"/></svg>`;
				break;
			case "menu-want-link":
				menuTextLink.innerHTML = `<svg fill="#ceceb5" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="24" height="24" viewBox="0,0,256,256"><g fill-rule="nonzero" stroke="none" stroke-width="1" stroke-linecap="butt" stroke-linejoin="miter" stroke-miterlimit="10" stroke-dasharray="" stroke-dashoffset="0" font-family="none" font-weight="none" font-size="none" text-anchor="none" style="mix-blend-mode: normal"><g transform="scale(10.66667,10.66667)"><path d="M11.98438,-0.01367c-0.55152,0.00862 -0.99193,0.46214 -0.98437,1.01367v1c-0.0051,0.36064 0.18438,0.69608 0.49587,0.87789c0.3115,0.18181 0.69676,0.18181 1.00825,0c0.3115,-0.18181 0.50097,-0.51725 0.49587,-0.87789v-1c0.0037,-0.2703 -0.10218,-0.53059 -0.29351,-0.72155c-0.19133,-0.19097 -0.45182,-0.29634 -0.72212,-0.29212zM4.18945,3.18945c-0.40692,0.00011 -0.77321,0.24676 -0.92633,0.62377c-0.15312,0.37701 -0.06255,0.80921 0.22907,1.09303l0.70117,0.70117c0.25082,0.26123 0.62327,0.36646 0.9737,0.2751c0.35044,-0.09136 0.62411,-0.36503 0.71546,-0.71546c0.09136,-0.35044 -0.01387,-0.72288 -0.2751,-0.9737l-0.70117,-0.70117c-0.18827,-0.19353 -0.4468,-0.30272 -0.7168,-0.30273zM19.78125,3.18945c-0.2598,0.00774 -0.50638,0.11632 -0.6875,0.30273l-0.70117,0.70117c-0.26123,0.25082 -0.36646,0.62326 -0.2751,0.9737c0.09136,0.35044 0.36503,0.6241 0.71546,0.71546c0.35044,0.09136 0.72288,-0.01387 0.9737,-0.2751l0.70117,-0.70117c0.29576,-0.28749 0.38469,-0.72707 0.22393,-1.10691c-0.16075,-0.37985 -0.53821,-0.62204 -0.9505,-0.60988zM12,5c-3.9,0 -7,3.1 -7,7c0,2.8 1.6,5.20078 4,6.30078v2.69922c0,1.1 0.9,2 2,2h2c1.1,0 2,-0.9 2,-2v-2.69922c2.4,-1.1 4,-3.50078 4,-6.30078c0,-3.9 -3.1,-7 -7,-7zM1,11c-0.36064,-0.0051 -0.69608,0.18438 -0.87789,0.49587c-0.18181,0.3115 -0.18181,0.69676 0,1.00825c0.18181,0.3115 0.51725,0.50097 0.87789,0.49587h1c0.36064,0.0051 0.69608,-0.18438 0.87789,-0.49587c0.18181,-0.3115 0.18181,-0.69676 0,-1.00825c-0.18181,-0.3115 -0.51725,-0.50097 -0.87789,-0.49587zM22,11c-0.36064,-0.0051 -0.69608,0.18438 -0.87789,0.49587c-0.18181,0.3115 -0.18181,0.69676 0,1.00825c0.18181,0.3115 0.51725,0.50097 0.87789,0.49587h1c0.36064,0.0051 0.69608,-0.18438 0.87789,-0.49587c0.18181,-0.3115 0.18181,-0.69676 0,-1.00825c-0.18181,-0.3115 -0.51725,-0.50097 -0.87789,-0.49587zM4.88086,18.08984c-0.2598,0.00774 -0.50638,0.11632 -0.6875,0.30273l-0.70117,0.70117c-0.26124,0.25082 -0.36647,0.62327 -0.27511,0.97371c0.09136,0.35044 0.36503,0.62411 0.71547,0.71547c0.35044,0.09136 0.72289,-0.01387 0.97371,-0.27511l0.70117,-0.70117c0.29576,-0.28749 0.38469,-0.72707 0.22393,-1.10691c-0.16075,-0.37985 -0.53821,-0.62204 -0.9505,-0.60988zM19.08984,18.08984c-0.40692,0.00011 -0.77321,0.24676 -0.92633,0.62377c-0.15312,0.37701 -0.06255,0.80922 0.22907,1.09303l0.70117,0.70117c0.25082,0.26124 0.62327,0.36648 0.97371,0.27512c0.35044,-0.09136 0.62411,-0.36503 0.71547,-0.71547c0.09136,-0.35044 -0.01388,-0.72289 -0.27512,-0.97371l-0.70117,-0.70117c-0.18827,-0.19353 -0.4468,-0.30272 -0.7168,-0.30273z"></path></g></g></svg>`;
				break;
			case "menu-watched-link":
				menuTextLink.innerHTML = `<svg fill="#ceceb5" viewBox="0 0 24 24" height="24px" width="24px"><g><g><path d="M12,4C4.063,4-0.012,12-0.012,12S3.063,20,12,20c8.093,0,12.011-7.969,12.011-7.969S20.062,4,12,4z M12.018,17c-2.902,0-5-2.188-5-5c0-2.813,2.098-5,5-5c2.902,0,5,2.187,5,5C17.018,14.812,14.92,17,12.018,17z M12.018,9c-1.658,0.003-3,1.393-3,3c0,1.606,1.342,3,3,3c1.658,0,3-1.395,3-3C15.018,10.392,13.676,8.997,12.018,9z"/></g></g></svg>`;
				break;
		}
	});
}


// switch menu buttons from icons to text
function toText() {
	console.log(`switch menu buttons from icons to text`);

	document.querySelectorAll(".nav-menu-button-elem a").forEach((menuTextLink) => {
		let linkID = menuTextLink.id;

		if (!linkID) {
			// skip iteration like continue
			return
		}
		switch (linkID) {
			case "menu-search-link":
				menuTextLink.innerHTML = `search`;
				break;
			case "menu-stared-link":
				menuTextLink.innerHTML = `stared`;
				break;
			case "menu-want-link":
				menuTextLink.innerHTML = `want`;
				break;
			case "menu-watched-link":
				menuTextLink.innerHTML = `watched`;
				break;
		}
	});
}


// change text to icons for mobile (small width)
// change icons to text for desktop (big width)
function changeMenuButtons() {
	const width = window.innerWidth

	if (width <= minWidth && !iconIsSet) {
		toIcons()
		iconIsSet = true
	} else if (width > minWidth && iconIsSet) {
		toText()
		iconIsSet = false
	}
}


// execute changeMenuButtons func with page resizing
window.addEventListener('resize', changeMenuButtons);

// execute changeMenuButtons func with page loading
document.addEventListener('DOMContentLoaded', changeMenuButtons);

// +-----------+
// + For forms +
// +-----------+

// add page GET query-params to POST request
function addQueryParamsToForm(formElem) {
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);
	formElem.action = formElem.action + "?" + searchParams.toString();
}

// add current page path with query-params like hidden input to form with POST request
function addHiddenNextAttributeToForm(formElem) {
	let currentPage = window.location.pathname + window.location.search;
	// create hidden input with current page path
	let input = document.createElement("input");
	input.type = "hidden";
	input.name = "next";
	input.value = currentPage;
	// add input to form
	formElem.appendChild(input);
}

// +-------------------+
// + Movie back button +
// +-------------------+

document.addEventListener("DOMContentLoaded", function () {
	let currentPath = window.location.pathname;
	let currentQuery = window.location.search;

	// skip for movie page
	if (currentPath.startsWith("/filmer/user-movie/info")) {
		return
	}

	localStorage.setItem("movie-back-url", currentPath + currentQuery);
});
