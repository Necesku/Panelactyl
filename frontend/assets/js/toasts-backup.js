const types = {
    success: {
        icon: {
            class: "material-icons icon",
            text: "check_circle",
            color: "#69ff63"
        }
    },
    error: {
        icon: {
            class: "material-icons icon",
            text: "cancel",
            color: "#ff6363"
        }
    },
    info: {
        icon: {
            class: "material-icons icon",
            text: "information",
            color: "#63acff"
        }
    },
    waiter: {
        icon: {
            class: "spinner",
            text: "",
            color: ""
        }
    }
}

const timeout = 5000;

let totalHeight = 0;
let counter = 0;

function delToast(el, elID, clientHeight) {
    document.getElementById(elID).classList.add("slide-out");
    setTimeout(() => {
        totalHeight = totalHeight - clientHeight - 4; 
        if (el.timeoutId) clearTimeout(el.timeoutId);
        el.remove();
    }, 400);
}

function createToast(id, title, description) {
    const type = types[id];
    if (type) {
        counter++;
        let elID = `toast-${counter}`;
        const el = document.createElement("li");
        const toastHtml = `<div class="toast slide-in" id="${elID}"><span class="${type.icon.class}" style="color: ${type.icon.color}">${type.icon.text}</span><div class="toast-content"><span class="title">${title}</span><p class="description">${description}</p></div></div>`
        el.innerHTML = toastHtml;
        el.style.listStyle = "none";
        document.body.appendChild(el);
        let toastEl = document.getElementById(elID);
        toastEl.style.transform = `translateY(-${totalHeight}px)`;
        totalHeight = totalHeight + toastEl.clientHeight + 4;
        el.timeoutId = setTimeout(() => delToast(el, elID, toastEl.clientHeight), timeout);
    }
}