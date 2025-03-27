// returns int page query param (deafult - 1)
function getPageQueryParam() {
	// parse keyword query param
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);

	if (searchParams.has("page")) {
		let pageParam = parseInt(searchParams.get("page"))
		if (pageParam > 0) {
			return pageParam
		// rewrite URL query param page to 1
		} else {
			searchParams.set("page", 1)
			window.history.replaceState({}, "", `${window.location.pathname}?${searchParams.toString()}`)
			return 1
		}
	} else {
		return 1
	}
}

// returns href link for page button with specified page (int)
function getPageHref(page) {
	// parse keyword query param
	let paramsString = window.location.search;
	let searchParams = new URLSearchParams(paramsString);

	// create and return href
	let href;
	if (searchParams.has("q")) {
		href = `${window.location.pathname}?q=${searchParams.get("q")}&page=${page}`;
	} else {
		href = `${window.location.pathname}?page=${page}`;
	}
	return href
}

// returns new page button elem with specified page (int)
function newPageButton(page) {
	let pageButton = document.createElement("a")
	pageButton.classList.add("page-button")
	pageButton.href = getPageHref(page)

	let innerDiv = document.createElement("div")
	innerDiv.textContent = page
	pageButton.appendChild(innerDiv)
	return pageButton
}

// returns new active page button elem with specified page (int)
function newActivePageButton(page) {
	let pageButton = document.createElement("div")
	pageButton.classList.add("page-button-active")
	pageButton.textContent = page
	return pageButton
}

// returns new three dots elem with
function newThreeDots() {
	let threeDots = document.createElement("div")
	threeDots.classList.add("page-three-dots")
	threeDots.textContent = "..."
	return threeDots
}

// add first and last pages to array of page elems (if not exists)
// pagesArr and num of pages (unt) must be presented
function addFirstLastToArray(pagesArr, pages) {
	// if first elem is not 1 then add it
	if (pagesArr[0].num !== 1) {
		pagesArr.unshift({num: 1, isActive: false})
	}
	// if last elem is not pages then add it
	if (pagesArr[pagesArr.length-1].num !== pages) {
		pagesArr.push({num: pages, isActive: false})
	}
}

// add three dots elems to array of page elems
function addThreeDotsToArray(pagesArr) {
	if (pagesArr[1].num - pagesArr[0].num > 1) {
		pagesArr.splice(1, 0, {num: null, isActive: false})
	}
	let len = pagesArr.length
	if (pagesArr[len-1].num - pagesArr[len-2].num > 1) {
		pagesArr.splice(len-1, 0, {num: null, isActive: false})
	}
}

// set pages buttons into pagination block
document.addEventListener("DOMContentLoaded", function () {
	const paginationBlock = document.getElementById("pagination");
	// get all pages amount and current page
	const pages = parseInt(paginationBlock.dataset.paginationPages);
	const currentPage = getPageQueryParam();

	// if no one page was found
	if (pages === 0) {
		return
	}

	// array of page elems {num: int, isActive: bool}
	// null num value means three dots instead of page button
	let pagesArr = [];
	if (pages < 5) {
		for (let pageNum = 1; pageNum <= pages; pageNum++) {
			if (pageNum === currentPage) {
				pagesArr.push({num: pageNum, isActive: true})
				continue
			}
			pagesArr.push({num: pageNum, isActive: false})
		}
	} else {
		// set upper divide
		let currentPageAbove = currentPage + 2;
		if (currentPageAbove > pages) {
			currentPageAbove = pages
		}
		// set lower divide
		let currentPageBelow = currentPage - 2;
		if (currentPageBelow < 1) {
			currentPageBelow = 1
		}
		// add currentPage elem and elems that above and below from it
		for (let pageNum = currentPageBelow; pageNum <= currentPageAbove; pageNum++) {
			if (pageNum === currentPage) {
				pagesArr.push({num: pageNum, isActive: true})
				continue
			}
			pagesArr.push({num: pageNum, isActive: false})
		}
		addFirstLastToArray(pagesArr, pages)
		addThreeDotsToArray(pagesArr)
	}

	// add page button elems according to pagesArr elems
	pagesArr.forEach((pageElem) => {
		if (pageElem.isActive) {
			paginationBlock.appendChild(newActivePageButton(pageElem.num))
		} else if (pageElem.num === null) {
			paginationBlock.appendChild(newThreeDots())
		} else {
			paginationBlock.appendChild(newPageButton(pageElem.num))
		}
	});
});
