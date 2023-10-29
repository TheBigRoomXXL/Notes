
var main = undefined;
var masonry = undefined;

// Hack because textarea cannot grow automatically with user input
function resizeTextarea(updateLayout) {
    this.style.height = "12px";
    this.style.height = this.scrollHeight + 4 + "px";
    if (updateLayout !== false) { layout(); }
}

function resizeAllTextarea() {
    textareas = document.querySelectorAll("textarea");
    textareas.forEach(textarea => {
        resizeTextarea.call(textarea, false);
    });
    layout();
}


// Reset input, focus on new and handle resizing textareas
function handleNoteCreated() {
    create_note_input = document.getElementById("0-content");
    create_note_input.value = "";
    resizeTextarea.call(create_note_input);

    new_note_input = document.getElementById("0").nextSibling.firstElementChild;
    new_note_input.addEventListener("input", resizeTextarea);
    new_note_input.focus();
    resizeTextarea.call(new_note_input);
}

function handleSearchSuccessfull() {
    textareas = document.querySelectorAll("textarea");
    textareas.forEach(textarea => {
        resizeTextarea.call(textarea);
        textarea.addEventListener("input", resizeTextarea);
    });
    layout();

    search = document.getElementById("fts-search");
    search.focus();
}

// function handleUpdateSuccessfull(noteElement) {
//     textarea = noteElement.firstElementChild
//     resizeTextarea.call(textarea);
//     textarea.addEventListener("input", resizeTextarea);
// }


// JS fallback for masonry
let grid = undefined;
let gridEl = document.querySelectorAll("#notes");
gridEl = document.getElementById("notes");

if (getComputedStyle(gridEl).gridTemplateRows !== 'masonry') {
    console.warn("css masonry grid not supported, falling back to js for support.");
    grid = {
        _el: gridEl,
        gap: parseFloat(getComputedStyle(gridEl).gap),
        items: [...gridEl.childNodes].filter(c => c.nodeType === 1),
        ncol: 0
    };
}

function layout() {
    if (grid === undefined) {
        return undefined // no grid to layout
    }

    // get the post-resize/ load number of columns
    let ncol = getComputedStyle(grid._el).gridTemplateColumns.split(' ').length;

    if (grid.ncol === ncol) {
        return undefined // no need to re-calculate layout/ 
    }

    /* update number of columns */
    grid.ncol = ncol;

    // revert to initial positioning, no margin 
    grid.items.forEach(c => c.style.removeProperty('margin-top'));

    // if we have more than one column 
    if (grid.ncol > 1) {
        grid.items.slice(ncol).forEach((current, i) => {
            // Make calculation on textarea, not on form as the expend to 
            // match there neighbours height
            const currentTextarea = grid.items[i].firstElementChild
            const nextTextarea = current.firstElementChild

            // Get the surronding coordinate
            const previous_end = currentTextarea.getBoundingClientRect().bottom; // bottom edge of item above 
            const current_start = nextTextarea.getBoundingClientRect().top; // top edge of current item

            // Move the box
            current.style.marginTop = `${previous_end + grid.gap - current_start}px`;
        });
    }
};



window.addEventListener("load", () => {
    resizeAllTextarea();
    document.querySelectorAll("textarea").forEach(textarea => {
        textarea.addEventListener("input", resizeTextarea);
    });

    layout();
    addEventListener('resize', layout, false);
});


document.addEventListener("htmx:afterSettle", function (evt) {
    if (evt.detail.requestConfig.elt.id == "0" && evt.detail.successful) {
        handleNoteCreated(evt);
    }
    if (evt.detail.requestConfig.elt.id == "fts-search" && evt.detail.successful) {
        handleSearchSuccessfull(evt);
    }
});

document.body.addEventListener('htmx:beforeSwap', function (evt) {
    if (evt.detail.xhr.status === 204) {
        // 204 represent a succesful DELETE and as such the elment should be 
        // swaped in order to be removed. 
        // For details see https://github.com/labstack/echo/issues/241
        evt.detail.shouldSwap = true;
        evt.detail.elt.setAttribute("hx-swap", "outerHTML")
    }
});



