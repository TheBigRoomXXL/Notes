// Bloody hack because textarea cannot grow automatically with user input

window.addEventListener("load", () => {

    var main = document.getElementById("notes")

    var masonry = new Masonry(main, {
        itemSelector: ".note",
        gutter: 15,
        isFitWidth: true,
        transitionDuration: 0,
    });

    function resizeTextarea() {
        this.style.height = '12px';
        this.style.height = this.scrollHeight + 4 + "px";
        masonry.layout()
    }

    textareas = document.querySelectorAll("textarea")
    textareas.forEach(textarea => {
        textarea.style.height = '12px';
        textarea.style.height = textarea.scrollHeight + 4 + 'px';
        textarea.addEventListener('input', resizeTextarea);
    });
    masonry.layout()
})
