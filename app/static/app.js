
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
    layout();

    textareas = document.querySelectorAll("textarea");
    textareas.forEach(textarea => {
        resizeTextarea.call(textarea);
        textarea.addEventListener("input", resizeTextarea);
    });

    search = document.getElementById("fts-search");
    search.focus();
}


// JS fallback for masonry
let grid = undefined;
let gridEl = document.getElementById("notes");
if (getComputedStyle(gridEl).gridTemplateRows !== 'masonry') {
    console.log("css masonry grid not supported, falling back to js.");
    grid = {
        _el: gridEl,
        gap: parseFloat(getComputedStyle(gridEl).gap),
        items: [...gridEl.childNodes].filter(c => c.nodeType === 1),
        ncol: 0
    };
}

function layout() {
    if (grid != undefined) {
        // get the post-resize/ load number of columns
        let ncol = getComputedStyle(grid._el).gridTemplateColumns.split(' ').length;
        console.log("layouting", grid);
        if (grid.ncol !== ncol) {
            grid.ncol = ncol;

            // revert to initial positioning, no margin 
            grid.items.forEach(c => c.style.removeProperty('margin-top'));

            // if we have more than one column 
            if (grid.ncol > 1) {
                grid.items.slice(ncol).forEach((c, i) => {
                    let previous_end = grid.items[i].getBoundingClientRect().bottom; // bottom edge of item above 
                    let current_start = c.getBoundingClientRect().top; // top edge of current item

                    c.style.marginTop = `${previous_end + grid.gap - current_start}px`;
                });
            }
        }
    };
}


window.addEventListener("load", () => {
    resizeAllTextarea();
    document.querySelectorAll("textarea").forEach(textarea => {
        textarea.addEventListener("input", resizeTextarea);
    });
});

window.addEventListener('resize', layout, false);

document.addEventListener("htmx:afterSettle", function (evt) {
    if (evt.detail.requestConfig.elt.id == "0" && evt.detail.successful) {
        handleNoteCreated(evt);
    }
    if (evt.detail.requestConfig.elt.id == "fts-search" && evt.detail.successful) {
        handleSearchSuccessfull(evt);
    }
});
