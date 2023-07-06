const overlapEls = document.querySelectorAll(".overlap") || [];
overlapEls.forEach(el => {
    const chars = [...el.textContent];
    el.innerHTML = "";
    chars.forEach((char, index) => {
        const span = document.createElement("span");
        span.textContent = char;
        span.style.setProperty("--index", index)
        el.append(span)
    })
})
