const searchParams = new URLSearchParams(window.location.search);

// +-----------------+
// + burger settings +
// +-----------------+

const modalMinWidth = 1080
let modalIsSet = false

function toModal() {
	document.getElementById("flex-settings").style.display = "none"
	document.getElementById("burger-flex-settings").style.display = "block"
}
function toAside() {
	document.getElementById("flex-settings").style.display = "flex"
	document.getElementById("burger-flex-settings").style.display = "none"
}
// change aside to modal for mobile (small width)
// change modal to aside for desktop (big width)
function changeSortFilterSettings() {
	const width = window.innerWidth

	if (width <= modalMinWidth && !modalIsSet) {
		toModal()
		modalIsSet = true
	} else if (width > modalMinWidth && modalIsSet) {
		toAside()
		modalIsSet = false
	}
}

// execute changeSortFilterSettings func with page resizing
window.addEventListener('resize', changeSortFilterSettings);
// execute changeSortFilterSettings func with page loading
document.addEventListener('DOMContentLoaded', changeSortFilterSettings);


// open and close modal window with sort and filters with clicking buttons
document.addEventListener('DOMContentLoaded', () => {
	// open modal window
    document.getElementById("burger-btn").addEventListener("click", function(event) {
		console.log("open")
		document.getElementById("modal-flex-settings").style.display = "block"

		const sortFilterSettingsElem = document.getElementById("flex-settings")
		const modalContainer = document.querySelector("#modal-flex-settings .modal-window")
		modalContainer.appendChild(sortFilterSettingsElem)
		sortFilterSettingsElem.style = "display: flex"
	});
	// close modal window
    document.querySelector("#modal-flex-settings .modal-close").addEventListener("click", function(event) {
    	console.log("close")
    	document.getElementById("modal-flex-settings").style.display = "none"

		const sortFilterSettingsElem = document.getElementById("flex-settings")
		const asideContainer = document.querySelector(".flex-content")
		asideContainer.appendChild(sortFilterSettingsElem)
		sortFilterSettingsElem.style = "display: none"
    });
});

// +----------------------------+
// + set and select order links +
// +----------------------------+

// add background to selected sort links and setup sort links href
function setupSortLinks() {
	let sortField = searchParams.get("sortField")
	// if sortField is not presented
	if (sortField === null) {
		sortField = "updated_at"
	}
	// sort fields
	document.querySelectorAll(".sort-field-link").forEach((sortFieldLink) => {
		if (sortFieldLink.dataset.linkValue === sortField) {
			sortFieldLink.style.backgroundColor = "#2e2e2e"
		}
	});

	let sortOrder = searchParams.get("sortOrder")
	// if sortOrder is not presented
	if (sortOrder === null && sortField === "updated_at") {
		sortOrder = "desc"
	} else if (sortOrder === null) {
		sortOrder = "asc"
	}
	// sort orders
	document.querySelectorAll(".sort-order-link").forEach((sortOrderLink) => {
		if (sortOrderLink.dataset.linkValue === sortOrder) {
			sortOrderLink.style.backgroundColor = "#2e2e2e"
		}
	});

	updateSortLinksHref(sortField, sortOrder)
}

// update sort links href with selected sort field and order
function updateSortLinksHref(sortField, sortOrder) {
	document.querySelectorAll(".sort-field-link").forEach((sortFieldLink) => {
		sortFieldLink.href = sortFieldLink.href + `&sortOrder=${sortOrder}`
	});
	document.querySelectorAll(".sort-order-link").forEach((sortOrderLink) => {
		sortOrderLink.href = sortOrderLink.href + `&sortField=${sortField}`
	});
}

document.addEventListener('DOMContentLoaded', setupSortLinks);

// +----------------------------------+
// + checkbox input for genres filter +
// +----------------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let genreList = searchParams.getAll("genres")

	document.querySelectorAll("#filter-genre-list .checkbox-btn input").forEach((genreInput) => {
		if (genreList.includes(genreInput.value)) {
			genreInput.checked = true;
		}
	})
});

// +-----------------------------+
// + range input for rating from +
// +-----------------------------+

// set range value ratingFrom with page loading
document.addEventListener('DOMContentLoaded', () => {
	const rangeBox = document.getElementById("filter-rating-from")

	let ratingFrom = searchParams.get("ratingFrom")
	if (ratingFrom === null) {
		ratingFrom = 0
	}
	rangeBox.querySelector("input").value = ratingFrom
	rangeBox.querySelector("span").textContent = ratingFrom
});

// update the current ratingFrom range value (each time you drag the range handle)
document.addEventListener('DOMContentLoaded', () => {
	var ratingFromRangeInput = document.querySelector("#filter-rating-from input");
	var ratingFromRangeValue = document.querySelector("#filter-rating-from span");
	ratingFromRangeInput.oninput = function() {
	    ratingFromRangeValue.textContent = this.value;
	}
});

// +-------------------------------+
// + number input for year filters +
// +-------------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let yearFrom = parseInt(searchParams.get("yearFrom"))
	if (isNaN(yearFrom) || yearFrom < 1500) {
		yearFrom = 1500
	}
	let yearTo = parseInt(searchParams.get("yearTo"))
	if (isNaN(yearTo) || yearTo > 3000) {
		yearTo = 3000
	}
	document.getElementById("filter-year-from").value = yearFrom;
	document.getElementById("filter-year-to").value = yearTo;
});

// +-----------------------------+
// + radio input for type filter +
// +-----------------------------+

document.addEventListener('DOMContentLoaded', () => {
	let type = searchParams.get("type")
	if (type === null) {
		type = "все"
	}
	document.querySelectorAll("#filter-type .radio-button input").forEach((typeInput) => {
		if (typeInput.value === type) {
			typeInput.checked = true;
		}
	})
});

// +--------------+
// + query params +
// +--------------+

// add sort query-params values to filter form
document.addEventListener('DOMContentLoaded', () => {
	if (searchParams.has("sortField")) {
		document.getElementById("sort-field-in-filter").value = searchParams.get("sortField")
	}
	if (searchParams.has("sortOrder")) {
		document.getElementById("sort-order-in-filter").value = searchParams.get("sortOrder")
	}
});

// add filter query params (from page URL) to sort links (then click to follow link)
document.addEventListener('DOMContentLoaded', () => {
	document.querySelectorAll(".sort-field-link, .sort-order-link").forEach(sortLink => {
	    sortLink.addEventListener("click", function(event) {
	    	// abort link following
	        event.preventDefault();
	        // get sort link query params
			const linkSearchParams = new URLSearchParams(new URL(sortLink.href).search);
			// update page query params with sort link query params
	        searchParams.set("sortField", linkSearchParams.get("sortField"));
	        searchParams.set("sortOrder", linkSearchParams.get("sortOrder"));
	        // follow the new link
	        window.location.href = `${window.location.pathname}?${searchParams.toString()}`;
	    });
	});
});
