
window.addEventListener("DOMContentLoaded", () => {
    let elm = document.getElementsByTagName("code");
    for (let i = 0; i < elm.length; i++) elm[i].classList.add("prettyprint");
    prettyPrint();
});