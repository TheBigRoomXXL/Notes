
var main = undefined;
var masonry = undefined;

// Hack because textarea cannot grow automatically with user input
function resizeTextarea() {
    this.style.height = "12px";
    this.style.height = this.scrollHeight + 4 + "px";
    masonry.layout();
}

function handleNoteCreated() {
    create_note = document.getElementById("create_note");
    resizeTextarea.call(create_note);

    create_note_input = create_note.firstElementChild;
    create_note_input.value = "";

    masonry.reloadItems();
    new_note = create_note.nextSibling;
    new_note.firstElementChild.focus();

    resizeTextarea.call(new_note.firstElementChild);
    new_note.addEventListener("input", resizeTextarea);
}

window.addEventListener("load", () => {
    main = document.getElementById("notes");

    masonry = new Masonry(main, {
        itemSelector: ".note",
        gutter: 15,
        isFitWidth: true,
        horizontalOrder: true,
        transitionDuration: 0,
    });

    textareas = document.querySelectorAll("textarea");
    textareas.forEach(textarea => {
        resizeTextarea.call(textarea);

        textarea.addEventListener("input", resizeTextarea);
    });
    masonry.layout();
});

document.addEventListener("htmx:afterSwap", function (evt) {
    if (evt.detail.requestConfig.elt.id == "create_note" && evt.detail.successful) {
        handleNoteCreated(evt);
    }
});
