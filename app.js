// Bloody hack because textarea cannot grow automatically with user input
function resizeTextarea() {
    this.style.height = '12px';
    this.style.height = "calc(" + this.scrollHeight + "px + 2em)";
}


window.addEventListener("load", () => {

    var main = document.getElementById("notes")

    var masonry = new Masonry(main, {
        itemSelector: ".note",
        gutter: 15,
        isFitWidth: true,
        transitionDuration: 0,
    });

    textareas = document.querySelectorAll("textarea")
    textareas.forEach(textarea => {
        textarea.style.height = textarea.scrollHeight + 2 + 'px';
        textarea.addEventListener('input', resizeTextarea);
    });
    masonry.layout()
})
