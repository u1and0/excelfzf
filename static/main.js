import { Fzf } from "./node_modules/fzf/dist/fzf.es.js";
main();
async function fetchPath(url) {
    return await fetch(url)
        .then((response) => {
        return response.json();
    })
        .catch((response) => {
        return Promise.reject(new Error(`{${response.status}: ${response.statusText}`));
    });
}
function fzfSearch(list, keyword) {
    const fzf = new Fzf(list);
    const entries = fzf.find(keyword);
    const ranking = entries.map((entry) => entry.item);
    return ranking;
}
async function main() {
    const url = new URL(window.location.href);
    const list = await fetchPath(url.origin + "/list");
    const searchInput = document.getElementById("search-form");
    const resultOutput = document.getElementById("search-result");
    searchInput.addEventListener("keyup", () => {
        while (resultOutput.firstChild) {
            resultOutput.removeChild(resultOutput.firstChild);
        }
        const result = fzfSearch(list, searchInput.value);
        console.log(result);
        result.map((line) => {
            const p = document.createElement("p");
            const text = document.createTextNode(line);
            p.appendChild(text);
            resultOutput.append(p);
        });
    });
}
