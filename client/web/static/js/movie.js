// color active icons (stared || want || watched) to red color
document.addEventListener("DOMContentLoaded", function () {
	// select all forms with icons
	const iconsElems = document.querySelectorAll(".movie-extra form")
	iconsElems.forEach((iconElem) => {
		// if icon data "active" is true
		if (iconElem.dataset.active === "true") {
			const iconElemSvg = iconElem.querySelector("button svg");
			// change styles for svg elem
			iconElemSvg.classList.remove('fill-static-white')
			iconElemSvg.classList.remove('fill-hover-light-red')
			iconElemSvg.classList.add('fill-static-light-red')
			iconElemSvg.classList.add('fill-hover-white')
		}
	});
});
